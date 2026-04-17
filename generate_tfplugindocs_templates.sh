#!/bin/bash

set -e

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== Terraform Provider Documentation Template Generator ===${NC}\n"

# Determine provider name from go.mod or directory name
PROVIDER_NAME=$(grep "^module" go.mod 2>/dev/null | sed 's/module.*terraform-provider-//' || basename $(pwd) | sed 's/terraform-provider-//')
echo -e "${GREEN}Provider name: ${PROVIDER_NAME}${NC}\n"

# Create directories
mkdir -p templates/resources
mkdir -p templates/data-sources
mkdir -p examples/provider
mkdir -p examples/resources
mkdir -p examples/data-sources

# Function to extract resource/data source names from provider code
extract_resources() {
    local map_name=$1  # "allResources" or "DataSourcesMap"

    # Find the provider.go or main provider file
    local provider_file=$(find . -name "provider.go" | grep -E "(internal|chronosphere)" | head -n 1)

    if [ -z "$provider_file" ]; then
        echo -e "${YELLOW}Warning: Could not find provider.go file${NC}"
        return
    fi

    # Extract resource names from the specific map
    # This handles the format: "chronosphere_resource_name": resourceFunction(),
    awk -v map="$map_name" '
        /'"$map"'.*:= map\[string\]/ { in_map=1; next }
        in_map && /^[[:space:]]*}/ { in_map=0 }
        in_map && /"'"${PROVIDER_NAME}"'_[a-z_]+"/ {
            match($0, /"'"${PROVIDER_NAME}"'_[a-z_]+"/, arr)
            gsub(/"/, "", arr[0])
            gsub(/'"${PROVIDER_NAME}"'_/, "", arr[0])
            print arr[0]
        }
    ' "$provider_file" | sort -u
}

# Alternative function using grep for simpler extraction
extract_resources_grep() {
    local map_name=$1
    local provider_file=$(find . -name "provider.go" | grep -E "(internal|chronosphere)" | head -n 1)

    if [ -z "$provider_file" ]; then
        return
    fi

    # Extract the map section and get resource names
    sed -n "/^[[:space:]]*${map_name}.*:= map/,/^[[:space:]]*}/p" "$provider_file" | \
        grep -o "\"${PROVIDER_NAME}_[a-z_]*\"" | \
        sed "s/\"${PROVIDER_NAME}_//" | \
        sed 's/"//' | \
        sort -u
}

# Function to create resource template
create_resource_template() {
    local resource_name=$1
    local template_file="templates/resources/${resource_name}.md.tmpl"

    if [ -f "$template_file" ]; then
        echo -e "${YELLOW}  ⊙ $template_file already exists, skipping${NC}"
        return
    fi

    # Convert snake_case to Title Case for display
    local display_name=$(echo "$resource_name" | sed 's/_/ /g' | awk '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) tolower(substr($i,2));}1')

    cat > "$template_file" <<EOF
---
page_title: "${PROVIDER_NAME}_${resource_name} Resource - ${PROVIDER_NAME}"
subcategory: ""
description: |-
  Manages a Chronosphere ${display_name}.
---

# ${PROVIDER_NAME}_${resource_name} (Resource)

Manages a Chronosphere ${display_name}.

## Example Usage

{{tffile "examples/resources/${PROVIDER_NAME}_${resource_name}/resource.tf"}}

{{ .SchemaMarkdown | trimspace }}

## Import

Import is supported using the following syntax:

{{codefile "shell" "examples/resources/${PROVIDER_NAME}_${resource_name}/import.sh"}}
EOF

    echo -e "${GREEN}  ✓ Created $template_file${NC}"
}

# Function to create data source template
create_data_source_template() {
    local ds_name=$1
    local template_file="templates/data-sources/${ds_name}.md.tmpl"

    if [ -f "$template_file" ]; then
        echo -e "${YELLOW}  ⊙ $template_file already exists, skipping${NC}"
        return
    fi

    # Convert snake_case to Title Case for display
    local display_name=$(echo "$ds_name" | sed 's/_/ /g' | awk '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) tolower(substr($i,2));}1')

    cat > "$template_file" <<EOF
---
page_title: "${PROVIDER_NAME}_${ds_name} Data Source - ${PROVIDER_NAME}"
subcategory: ""
description: |-
  Data source for retrieving a Chronosphere ${display_name}.
---

# ${PROVIDER_NAME}_${ds_name} (Data Source)

Data source for retrieving a Chronosphere ${display_name}.

## Example Usage

{{tffile "examples/data-sources/${PROVIDER_NAME}_${ds_name}/data-source.tf"}}

{{ .SchemaMarkdown | trimspace }}
EOF

    echo -e "${GREEN}  ✓ Created $template_file${NC}"
}

# Function to create example resource file
create_resource_example() {
    local resource_name=$1
    local example_dir="examples/resources/${PROVIDER_NAME}_${resource_name}"
    local example_file="${example_dir}/resource.tf"
    local import_file="${example_dir}/import.sh"

    mkdir -p "$example_dir"

    if [ ! -f "$example_file" ]; then
        cat > "$example_file" <<EOF
resource "${PROVIDER_NAME}_${resource_name}" "example" {
  # Add example configuration here
  name = "example-${resource_name}"
}
EOF
        echo -e "${GREEN}  ✓ Created $example_file${NC}"
    else
        echo -e "${YELLOW}  ⊙ $example_file already exists, skipping${NC}"
    fi

    if [ ! -f "$import_file" ]; then
        cat > "$import_file" <<EOF
# Replace 'example-id' with the actual resource ID
terraform import ${PROVIDER_NAME}_${resource_name}.example example-id
EOF
        echo -e "${GREEN}  ✓ Created $import_file${NC}"
    else
        echo -e "${YELLOW}  ⊙ $import_file already exists, skipping${NC}"
    fi
}

# Function to create example data source file
create_data_source_example() {
    local ds_name=$1
    local example_dir="examples/data-sources/${PROVIDER_NAME}_${ds_name}"
    local example_file="${example_dir}/data-source.tf"

    mkdir -p "$example_dir"

    if [ ! -f "$example_file" ]; then
        cat > "$example_file" <<EOF
data "${PROVIDER_NAME}_${ds_name}" "example" {
  # Add example configuration here
  id = "example-id"
}
EOF
        echo -e "${GREEN}  ✓ Created $example_file${NC}"
    else
        echo -e "${YELLOW}  ⊙ $example_file already exists, skipping${NC}"
    fi
}

# Create provider index template
create_provider_index() {
    local template_file="templates/index.md.tmpl"

    if [ -f "$template_file" ]; then
        echo -e "${YELLOW}⊙ $template_file already exists, skipping${NC}\n"
        return
    fi

    cat > "$template_file" <<EOF
---
page_title: "Provider: ${PROVIDER_NAME}"
description: |-
  The ${PROVIDER_NAME} provider is used to interact with Chronosphere resources.
---

# ${PROVIDER_NAME} Provider

The ${PROVIDER_NAME} provider allows you to manage Chronosphere resources using Terraform.

## Example Usage

{{tffile "examples/provider/provider.tf"}}

## Authentication

The provider supports authentication via API token. You can provide credentials via:

- Provider configuration (shown above)
- Environment variables

### Environment Variables

- \`CHRONOSPHERE_API_TOKEN\` - API token for authentication
- \`CHRONOSPHERE_ORG\` - Organization name

{{ .SchemaMarkdown | trimspace }}
EOF

    echo -e "${GREEN}✓ Created $template_file${NC}\n"
}

# Create provider example
create_provider_example() {
    local example_file="examples/provider/provider.tf"

    if [ -f "$example_file" ]; then
        echo -e "${YELLOW}⊙ $example_file already exists, skipping${NC}\n"
        return
    fi

    cat > "$example_file" <<EOF
provider "${PROVIDER_NAME}" {
  api_token = var.chronosphere_api_token
  org       = var.chronosphere_org

  # Optional: Override API endpoint
  # api_url = "https://api.chronosphere.io"
}

variable "chronosphere_api_token" {
  description = "Chronosphere API token"
  type        = string
  sensitive   = true
}

variable "chronosphere_org" {
  description = "Chronosphere organization"
  type        = string
}
EOF

    echo -e "${GREEN}✓ Created $example_file${NC}\n"
}

# Main execution
echo -e "${BLUE}Creating provider templates...${NC}"
create_provider_index
create_provider_example

# Process resources
echo -e "${BLUE}Discovering resources...${NC}"
RESOURCES=$(extract_resources_grep "allResources")

if [ -z "$RESOURCES" ]; then
    echo -e "${YELLOW}No resources found in provider code${NC}\n"
else
    echo -e "${GREEN}Found $(echo "$RESOURCES" | wc -l) resources${NC}\n"
    echo -e "${BLUE}Creating resource templates and examples...${NC}"

    while IFS= read -r resource; do
        [ -z "$resource" ] && continue
        echo "Processing resource: ${resource}"
        create_resource_template "$resource"
        create_resource_example "$resource"
        echo ""
    done <<< "$RESOURCES"
fi

# Process data sources
echo -e "${BLUE}Discovering data sources...${NC}"
DATA_SOURCES=$(extract_resources_grep "DataSourcesMap")

if [ -z "$DATA_SOURCES" ]; then
    echo -e "${YELLOW}No data sources found in provider code${NC}\n"
else
    echo -e "${GREEN}Found $(echo "$DATA_SOURCES" | wc -l) data sources${NC}\n"
    echo -e "${BLUE}Creating data source templates and examples...${NC}"

    while IFS= read -r ds; do
        [ -z "$ds" ] && continue
        echo "Processing data source: ${ds}"
        create_data_source_template "$ds"
        create_data_source_example "$ds"
        echo ""
    done <<< "$DATA_SOURCES"
fi

echo -e "${GREEN}=== Done! ===${NC}"
echo -e "\nNext steps:"
echo -e "1. Review and customize the generated templates in ${BLUE}templates/${NC}"
echo -e "2. Fill in the example configurations in ${BLUE}examples/${NC}"
echo -e "3. Run ${BLUE}tfplugindocs generate${NC} to create the documentation"
echo -e "4. Review the generated docs in ${BLUE}docs/${NC}"