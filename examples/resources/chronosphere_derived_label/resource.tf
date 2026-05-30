resource "chronosphere_derived_label" "service_tier" {
  name        = "Service Tier Label"
  slug        = "service-tier"
  description = "Derives service tier from instance name patterns"
  label_name  = "tier"

  metric_label {
    constructed_label {
      value_definitions {
        value = "production"
        filters {
          name       = "instance"
          value_glob = "prod-*"
        }
      }
      value_definitions {
        value = "staging"
        filters {
          name       = "instance"
          value_glob = "staging-*"
        }
      }
      value_definitions {
        value = "development"
        filters {
          name       = "instance"
          value_glob = "dev-*"
        }
      }
    }
  }
}

resource "chronosphere_derived_label" "service_mapping" {
  name        = "Service Name Mapping"
  slug        = "service-mapping"
  description = "Maps various service labels to standard service names"
  label_name  = "service"

  metric_label {
    mapping_label {
      name_mappings {
        source_label = "k8s_service"
        filters {
          name       = "__name__"
          value_glob = "k8s_*"
        }
        value_mappings {
          target_value       = "api"
          source_value_globs = ["api-service", "api-gateway"]
        }
      }
      name_mappings {
        source_label = "backend_service"
        filters {
          name       = "__name__"
          value_glob = "http_*"
        }
      }
    }
  }

  existing_label_policy = "OVERRIDE"
}
