#!/bin/bash

# Directory containing the files
resources=$(jq --raw-output '.provider_schemas."registry.terraform.io/databricks/databricks".resource_schemas | keys | .[]' ./config/schema.json)
root_folder="./config"

rm -f ./.work/external_name.txt || true
rm -f ./.work/provider.txt || true

for resource in $resources
do
  resource_name=$(echo $resource | sed 's/databricks_//')

  if [ ! -d "$root_folder/$resource_name" ] | [ ! -f  "$root_folder/$resource_name/config.go" ]; then
    echo "Creating: $resource_name"
    mkdir "$root_folder/$resource_name"
    cat > "$root_folder/$resource_name/config.go"  << EOF
package $resource_name

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("$resource", func(r *config.Resource) {
		r.ShortGroup = "databricks"
	})
}
EOF
    echo "\"$resource\":               config.IdentifierFromProvider," >> ./.work/external_name.txt
    echo "$resource_name.Configure," >> ./.work/provider.txt

  else 
    echo "Config exists for: $resource_name"
  fi
done

echo "Open ./.work/external_name.txt and add it into ./config/external_name.go"
echo "Open ./.work/provider.txt and add it into ./config/provider.go"
