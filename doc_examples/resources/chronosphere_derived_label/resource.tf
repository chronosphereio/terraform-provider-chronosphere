resource "chronosphere_derived_label" "my-constructed-derived-label" {
  name        = "my-constructed-label"
  slug        = "my-constructed-label"
  description = "this is my derived label"
  label_name  = "tier"

  metric_label {
    constructed_label {
      value_definitions {
        value = "read"
        filters {
          name       = "instance"
          value_glob = "reader-*"
        }
      }
      value_definitions {
        value = "write"
        filters {
          name       = "instance"
          value_glob = "writer-*"
        }
      }
      value_definitions {
        value = "admin"
        filters {
          name       = "instance"
          value_glob = "admin-*"
        }
      }
    }
  }
}

resource "chronosphere_derived_label" "my-mapping-derived-label" {
  name        = "my-mapping-label"
  slug        = "my-mapping-label"
  description = "this is my derived label"
  label_name  = "my_derived_label"

  metric_label {
    mapping_label {
      name_mappings {
        source_label = "grpc_service"
        filters {
          name       = "__name__"
          value_glob = "grpc_*"
        }
        value_mappings {
          target_value       = "gateway"
          source_value_globs = ["rpcgateway, gateway-service"]
        }
      }
      name_mappings {
        source_label = "backend_service"
        filters {
          name       = "__name__"
          value_glob = "envoy_*"
        }
      }
      value_mappings {
        target_value       = "auth"
        source_value_globs = ["Auth, auth-service"]
      }
    }
  }

  existing_label_policy = "OVERRIDE"
}
