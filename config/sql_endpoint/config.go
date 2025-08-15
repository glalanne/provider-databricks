package sql_endpoint

import (
	"strconv"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_sql_endpoint", func(r *config.Resource) {
		r.ShortGroup = "sql"

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}

			if jdbc_url, ok := attr["jdbc_url"].(string); ok {
				conn["jdbc_url"] = []byte(jdbc_url)
			}

			if instances, ok := attr["odbc_params"].([]any); ok {
				if len(instances) > 0 {
					instance := instances[0].(map[string]any)

					if host, ok := instance["hostname"].(string); ok {
						conn["odbc_params_hostname"] = []byte(host)
					}
					if path, ok := instance["path"].(string); ok {
						conn["odbc_params_path"] = []byte(path)
					}
					if port, ok := instance["port"].(float64); ok {
						conn["odbc_params_port"] = []byte(strconv.Itoa(int(port)))
					}
					if protocol, ok := instance["protocol"].(string); ok {
						conn["odbc_params_protocol"] = []byte(protocol)
					}
				}
			}
			return conn, nil
		}
	})
}
