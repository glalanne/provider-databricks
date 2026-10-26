package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	diff "github.com/crossplane/upjet/v2/proto/diff/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestDiffServerWithoutKubernetes(t *testing.T) {
	dir := t.TempDir()
	binary := filepath.Join(dir, "provider")
	build := exec.CommandContext(t.Context(), "go", "build", "-o", binary, ".")
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build provider: %v\n%s", err, output)
	}
	var requests atomic.Int64
	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests.Add(1)
		http.Error(w, "network forbidden in offline integration test", http.StatusForbidden)
	}))
	defer proxy.Close()
	// Short Unix paths also work on macOS, whose socket path limit is 104 bytes.
	socketDir, err := os.MkdirTemp("", "dbdiff-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(socketDir)
	socket := filepath.Join(socketDir, "s")
	logPath := filepath.Join(dir, "server.log")
	log, err := os.Create(logPath)
	if err != nil {
		t.Fatal(err)
	}
	defer log.Close()
	cmd := exec.CommandContext(t.Context(), binary, "internal", "diff-server", "--network", "unix", "--address", socket)
	for _, env := range os.Environ() {
		key, _, _ := strings.Cut(env, "=")
		upper := strings.ToUpper(key)
		if strings.HasPrefix(upper, "KUBE") || strings.HasPrefix(upper, "DATABRICKS_") ||
			upper == "HTTP_PROXY" || upper == "HTTPS_PROXY" || upper == "NO_PROXY" || upper == "ALL_PROXY" {
			continue
		}
		cmd.Env = append(cmd.Env, env)
	}
	cmd.Env = append(cmd.Env, "KUBECONFIG="+filepath.Join(dir, "missing-kubeconfig"),
		"HTTP_PROXY="+proxy.URL, "HTTPS_PROXY="+proxy.URL, "NO_PROXY=", "ALL_PROXY="+proxy.URL,
		"DATABRICKS_AUTH_TYPE=azure-cli", "DATABRICKS_HOST=https://environment.invalid",
		"DATABRICKS_CLIENT_ID=environment-client", "DATABRICKS_CLIENT_SECRET=environment-secret")
	cmd.Stdout, cmd.Stderr = log, log
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	defer func() {
		_ = cmd.Process.Signal(os.Interrupt)
		select {
		case <-done:
		case <-time.After(10 * time.Second):
			_ = cmd.Process.Kill()
			<-done
			t.Error("diff server did not stop on interrupt")
		}
		if n := requests.Load(); n != 0 {
			t.Errorf("offline server leaked %d HTTP requests", n)
		}
	}()
	conn, err := grpc.NewClient("unix://"+socket, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(t.Context(), 30*time.Second)
	defer cancel()
	for {
		if _, err := os.Stat(socket); err == nil {
			break
		}
		select {
		case err := <-done:
			done <- err
			output, _ := os.ReadFile(logPath)
			t.Fatalf("server exited before listening: %v\n%s", err, output)
		case <-ctx.Done():
			t.Fatal("server did not listen without Kubernetes config")
		case <-time.After(50 * time.Millisecond):
		}
	}
	c := diff.NewPlanServiceClient(conn)
	for _, scope := range []string{"legacy", "ProviderConfig", "ClusterProviderConfig"} {
		for _, resource := range []struct {
			kind, group string
			params      map[string]any
		}{
			{kind: "Notebook", group: "workspace", params: map[string]any{"path": "/Shared/offline", "contentBase64": "cHJpbnQoMSk=", "language": "PYTHON"}},
			{kind: "BudgetPolicy", group: "billing", params: map[string]any{"policyName": "offline-policy"}},
		} {
			t.Run(scope+"/"+resource.kind, func(t *testing.T) {
				root, pcKind := "databricks.m.crossplane.io", scope
				meta := map[string]any{"name": "offline", "namespace": "test"}
				pcMeta := map[string]any{"name": "default"}
				ref := map[string]any{"name": "default"}
				if scope == "legacy" {
					root, pcKind = "databricks.crossplane.io", "ProviderConfig"
					delete(meta, "namespace")
				} else {
					ref["kind"] = scope
					if scope == "ProviderConfig" {
						pcMeta["namespace"] = "test"
					}
				}
				object := func(v map[string]any) *structpb.Struct {
					t.Helper()
					s, err := structpb.NewStruct(v)
					if err != nil {
						t.Fatal(err)
					}
					return s
				}
				pc := object(map[string]any{"apiVersion": root + "/v1beta1", "kind": pcKind, "metadata": pcMeta,
					"spec": map[string]any{"credentials": map[string]any{"source": "None"}}})
				desired := object(map[string]any{"apiVersion": resource.group + "." + root + "/v1beta1", "kind": resource.kind,
					"metadata": meta, "spec": map[string]any{"providerConfigRef": ref, "forProvider": resource.params}})
				request := &diff.PlanRequest{DesiredResource: desired, KubernetesObjectStore: []*structpb.Struct{pc}}
				for i := 0; i < 2; i++ {
					callCtx, stop := context.WithTimeout(t.Context(), 10*time.Second)
					response, err := c.Plan(callCtx, request, grpc.WaitForReady(true))
					stop()
					if err != nil {
						t.Fatal(err)
					}
					if response.Action != diff.Action_ACTION_CREATE {
						t.Fatalf("action=%v, want CREATE", response.Action)
					}
				}
				if resource.kind == "BudgetPolicy" {
					for _, tc := range []struct {
						actualName string
						action     diff.Action
					}{
						{actualName: "old-policy", action: diff.Action_ACTION_UPDATE},
						{actualName: "offline-policy", action: diff.Action_ACTION_NO_OP},
					} {
						actualMeta := map[string]any{"name": "offline", "annotations": map[string]any{"crossplane.io/external-name": "policy-123"}}
						if scope != "legacy" {
							actualMeta["namespace"] = "test"
						}
						request.ActualResource = object(map[string]any{
							"apiVersion": resource.group + "." + root + "/v1beta1", "kind": resource.kind, "metadata": actualMeta,
							"spec":   map[string]any{"providerConfigRef": ref, "forProvider": map[string]any{"policyName": tc.actualName}},
							"status": map[string]any{"atProvider": map[string]any{"policyId": "policy-123", "policyName": tc.actualName}},
						})
						callCtx, stop := context.WithTimeout(t.Context(), 10*time.Second)
						response, err := c.Plan(callCtx, request)
						stop()
						if err != nil {
							t.Fatal(err)
						}
						if response.Action != tc.action {
							t.Fatalf("existing policy action=%v, want %v: %v", response.Action, tc.action, response)
						}
					}
				}
			})
		}
	}
}
