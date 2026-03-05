#!/usr/bin/env python3

import os
import sys
import glob
from pathlib import Path

# usage: version_diff.py <generated resource list> <base JSON schema path> <bumped JSON schema path>
# example usage: version_diff.py config/generated.lst .work/schema.json.3.38.0 config/schema.json
if __name__ == "__main__":
    directory=".work/databricks/databricks/docs/resources"
    files=glob.glob(f"{directory}/*.md")
    resources = []

    for file_path in files:
        base_name=Path(file_path).stem
        
        # check if folder exists in the config directory
        if os.path.exists(f"config/cluster/{base_name}"):
            print(f"Cluster config found: {base_name} - ✅")
        else:
            print(f"Cluster config not found: {base_name} - ❌")

        if os.path.exists(f"config/namespaced/{base_name}"):
            print(f"Namespaced config found: {base_name} - ✅")
        else:
            print(f"Namespaced config not found: {base_name} - ❌")

    print("------------------------------")

    external_names = ""
    with open("config/external_name.go", 'r') as file:
        external_names = file.read()

    for file_path in files:
        base_name=Path(file_path).stem
        if os.path.exists(f"config/namespaced/{base_name}") and external_names.find(f'databricks_{base_name}') == -1:
            print(f"External name not found: databricks_{base_name} - ❌")
    
    for resource in resources:
        print(f"{resource}")
    
   