#!/usr/bin/env python3

# SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: Apache-2.0

"""
Compare CRDs between two refs and report Kubernetes backward-incompatible changes.

This checker is conservative: anything clearly tightening validation or removing
API surface is treated as breaking.
"""

from __future__ import annotations

import argparse
import json
import os
import subprocess
import sys
from dataclasses import dataclass
from typing import Any, Iterable

@dataclass
class Change:
    severity: str  # FAIL | WARN | INFO
    message: str


def run_git(args: list[str]) -> str:
    cp = subprocess.run(
        ["git", *args],
        check=False,
        text=True,
        capture_output=True,
        env={**os.environ, "GIT_CONFIG_GLOBAL": os.environ.get("GIT_CONFIG_GLOBAL", "/dev/null"), "GIT_CONFIG_NOSYSTEM": os.environ.get("GIT_CONFIG_NOSYSTEM", "1")},
    )
    if cp.returncode != 0:
        raise RuntimeError(cp.stderr.strip() or cp.stdout.strip())
    return cp.stdout


def list_changed_crd_entries(base: str, head: str, crd_dir: str) -> list[tuple[str, str | None, str | None]]:
    raw = run_git(
        [
            "diff",
            "--name-status",
            "--find-renames",
            "--diff-filter=ACDMRT",
            base,
            head,
            "--",
            f"{crd_dir}/*.yaml",
            f"{crd_dir}/*.yml",
        ]
    )
    out: list[tuple[str, str | None, str | None]] = []
    for line in [l for l in raw.splitlines() if l.strip()]:
        parts = line.split("\t")
        status = parts[0]
        if status.startswith("R") and len(parts) >= 3:
            out.append((status, parts[1], parts[2]))
        elif len(parts) >= 2:
            out.append((status, parts[1], parts[1]))
    return out


def load_yaml_at_ref(ref: str, path: str) -> dict[str, Any]:
    text = run_git(["show", f"{ref}:{path}"])
    cp = subprocess.run(
        ["yq", "eval", "-o=json", ".", "-"],
        input=text,
        check=False,
        text=True,
        capture_output=True,
    )
    if cp.returncode != 0:
        raise RuntimeError(f"Failed to parse YAML for {path} at {ref}: {cp.stderr.strip()}")
    crd = json.loads(cp.stdout)
    if not isinstance(crd, dict) or crd.get("kind") != "CustomResourceDefinition":
        raise RuntimeError(f"{path} at {ref} is not a CRD")
    return crd


def version_map(crd: dict[str, Any]) -> dict[str, dict[str, Any]]:
    versions = crd.get("spec", {}).get("versions", []) or []
    out: dict[str, dict[str, Any]] = {}
    for v in versions:
        name = v.get("name")
        if isinstance(name, str):
            out[name] = v
    return out


def compare_crd(base_crd: dict[str, Any], head_crd: dict[str, Any], path: str) -> list[Change]:
    changes: list[Change] = []

    bspec = base_crd.get("spec", {})
    hspec = head_crd.get("spec", {})

    if bspec.get("group") != hspec.get("group"):
        changes.append(Change("FAIL", f"{path}: spec.group changed ({bspec.get('group')} -> {hspec.get('group')})"))

    bnames = bspec.get("names", {})
    hnames = hspec.get("names", {})
    for nkey in ("kind", "plural", "singular"):
        if bnames.get(nkey) != hnames.get(nkey):
            changes.append(Change("FAIL", f"{path}: spec.names.{nkey} changed ({bnames.get(nkey)} -> {hnames.get(nkey)})"))

    bversions = version_map(base_crd)
    hversions = version_map(head_crd)

    for v in sorted(set(bversions) - set(hversions)):
        if bool(bversions[v].get("served", False)):
            changes.append(Change("FAIL", f"{path}: removed served version '{v}'"))
        else:
            changes.append(Change("WARN", f"{path}: removed non-served version '{v}'"))

    for v in sorted(set(bversions) & set(hversions)):
        bv = bversions[v]
        hv = hversions[v]

        if bool(bv.get("served", False)) and not bool(hv.get("served", False)):
            changes.append(Change("FAIL", f"{path}: version '{v}' changed served=true -> false"))

        # storage version flips are not always breaking, but can signal migration risk.
        if bool(bv.get("storage", False)) != bool(hv.get("storage", False)):
            changes.append(Change("WARN", f"{path}: version '{v}' storage flag changed ({bv.get('storage')} -> {hv.get('storage')})"))

        bschema = ((bv.get("schema") or {}).get("openAPIV3Schema"))
        hschema = ((hv.get("schema") or {}).get("openAPIV3Schema"))
        if isinstance(bschema, dict) and isinstance(hschema, dict):
            compare_schema(path, v, "$", bschema, hschema, changes)

    return changes


def as_set(v: Any) -> set[str]:
    if isinstance(v, list):
        return {str(x) for x in v}
    return set()


def num(v: Any) -> float | None:
    if isinstance(v, (int, float)):
        return float(v)
    return None


def compare_schema(path: str, version: str, spath: str, b: dict[str, Any], h: dict[str, Any], out: list[Change]) -> None:
    loc = f"{path} [{version}] {spath}"

    # Type changes are generally breaking.
    btype = b.get("type")
    htype = h.get("type")
    if btype is not None and htype is not None and btype != htype:
        out.append(Change("FAIL", f"{loc}: type changed ({btype} -> {htype})"))

    # Required list additions are breaking for writers.
    breq = as_set(b.get("required"))
    hreq = as_set(h.get("required"))
    for req in sorted(hreq - breq):
        out.append(Change("FAIL", f"{loc}: new required field '{req}'"))

    # Enum narrowing is breaking.
    benum = as_set(b.get("enum"))
    henum = as_set(h.get("enum"))
    if benum and henum:
        removed = benum - henum
        if removed:
            out.append(Change("FAIL", f"{loc}: enum values removed {sorted(removed)}"))

    # Tightened string constraints can break existing values.
    bmin_len = num(b.get("minLength"))
    hmin_len = num(h.get("minLength"))
    if bmin_len is not None and hmin_len is not None and hmin_len > bmin_len:
        out.append(Change("FAIL", f"{loc}: minLength increased ({bmin_len} -> {hmin_len})"))

    bmax_len = num(b.get("maxLength"))
    hmax_len = num(h.get("maxLength"))
    if bmax_len is not None and hmax_len is not None and hmax_len < bmax_len:
        out.append(Change("FAIL", f"{loc}: maxLength decreased ({bmax_len} -> {hmax_len})"))

    bpattern = b.get("pattern")
    hpattern = h.get("pattern")
    if bpattern is not None and hpattern is not None and bpattern != hpattern:
        out.append(Change("WARN", f"{loc}: pattern changed (manual review needed)"))

    # Numeric constraints tightening.
    bmin = num(b.get("minimum"))
    hmin = num(h.get("minimum"))
    if bmin is not None and hmin is not None and hmin > bmin:
        out.append(Change("FAIL", f"{loc}: minimum increased ({bmin} -> {hmin})"))

    bmax = num(b.get("maximum"))
    hmax = num(h.get("maximum"))
    if bmax is not None and hmax is not None and hmax < bmax:
        out.append(Change("FAIL", f"{loc}: maximum decreased ({bmax} -> {hmax})"))

    # additionalProperties false is more restrictive than true/object.
    bap = b.get("additionalProperties")
    hap = h.get("additionalProperties")
    if bap is not False and hap is False:
        out.append(Change("FAIL", f"{loc}: additionalProperties changed to false"))

    # Recurse object properties.
    bprops = b.get("properties") if isinstance(b.get("properties"), dict) else {}
    hprops = h.get("properties") if isinstance(h.get("properties"), dict) else {}

    for k in sorted(set(bprops) - set(hprops)):
        out.append(Change("FAIL", f"{loc}: property removed '{k}'"))

    for k in sorted(set(hprops) - set(bprops)):
        out.append(Change("INFO", f"{loc}: property added '{k}'"))

    for k in sorted(set(bprops) & set(hprops)):
        bv = bprops[k]
        hv = hprops[k]
        if isinstance(bv, dict) and isinstance(hv, dict):
            compare_schema(path, version, f"{spath}.properties.{k}", bv, hv, out)

    # Recurse arrays.
    bitems = b.get("items")
    hitems = h.get("items")
    if isinstance(bitems, dict) and isinstance(hitems, dict):
        compare_schema(path, version, f"{spath}.items", bitems, hitems, out)


def summarize(changes: Iterable[Change]) -> tuple[int, int, int]:
    fail = warn = info = 0
    for c in changes:
        if c.severity == "FAIL":
            fail += 1
        elif c.severity == "WARN":
            warn += 1
        else:
            info += 1
    return fail, warn, info


def main() -> int:
    ap = argparse.ArgumentParser(description="Check CRD backward compatibility between refs")
    ap.add_argument("--base-ref", default="main")
    ap.add_argument("--head-ref", default="HEAD")
    ap.add_argument("--crd-dir", default="package/crds")
    ap.add_argument("--json", action="store_true", help="Emit JSON report")
    args = ap.parse_args()

    # Use merge-base for branch comparison.
    try:
        merge_base = run_git(["merge-base", args.base_ref, args.head_ref]).strip()
    except Exception as e:
        print(f"Unable to resolve merge-base for '{args.base_ref}' and '{args.head_ref}': {e}")
        return 2

    try:
        changes = list_changed_crd_entries(merge_base, args.head_ref, args.crd_dir)
    except Exception as e:
        print(f"Unable to list CRD changes: {e}")
        return 2

    if not changes:
        print(f"No CRD changes found between {args.base_ref} and {args.head_ref}.")
        return 0

    report: dict[str, list[dict[str, str]]] = {}
    all_changes: list[Change] = []

    for status, old_path, new_path in changes:
        if status == "A":
            c = Change("INFO", f"{new_path}: CRD added")
            report[new_path or "<unknown>"] = [{"severity": c.severity, "message": c.message}]
            all_changes.append(c)
            continue

        if status == "D":
            c = Change("FAIL", f"{old_path}: CRD deleted")
            report[old_path or "<unknown>"] = [{"severity": c.severity, "message": c.message}]
            all_changes.append(c)
            continue

        if not old_path or not new_path:
            continue

        try:
            b = load_yaml_at_ref(merge_base, old_path)
            h = load_yaml_at_ref(args.head_ref, new_path)
            per = compare_crd(b, h, new_path)
            if not per:
                per = [Change("INFO", f"{new_path}: no breaking changes detected")]
        except Exception as e:
            per = [Change("FAIL", f"{new_path}: comparison failed: {e}")]

        report[new_path] = [{"severity": c.severity, "message": c.message} for c in per]
        all_changes.extend(per)

    fail, warn, info = summarize(all_changes)

    if args.json:
        print(
            json.dumps(
                {
                    "base_ref": args.base_ref,
                    "head_ref": args.head_ref,
                    "merge_base": merge_base,
                    "summary": {"fail": fail, "warn": warn, "info": info},
                    "report": report,
                },
                indent=2,
                sort_keys=True,
            )
        )
    else:
        print(f"Comparing CRDs using merge-base {merge_base} ({args.base_ref}...{args.head_ref})")
        print(f"CRD directory: {args.crd_dir}")
        for crd, entries in report.items():
            print(f"\n[CRD] {crd}")
            for e in entries:
                print(f"  [{e['severity']}] {e['message']}")
        print(f"\nSummary: FAIL={fail} WARN={warn} INFO={info}")

    return 1 if fail > 0 else 0


if __name__ == "__main__":
    sys.exit(main())
