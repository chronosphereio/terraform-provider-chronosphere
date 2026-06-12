resource "chronosphere_consumption_config" "config" {
  partition {
    name = "metric partition"
    slug = "metric-part"
    filter {
      operator = "IN"
      condition {
        metric_filter {
          name       = "job"
          value_glob = "myservice*"
        }
      }
    }
  }

  partition {
    name = "trace partition"
    slug = "trace-part"
    filter {
      operator = "IN"
      condition {
        # A span matches the trace_filter if it satisfies ANY one
        # span_filter block: blocks OR together, while the conditions
        # within a block must all hold on the same span (AND).
        trace_filter {
          span_filter {
            service {
              match = "exact"
              value = "my-service"
            }
            error {
              value = true
            }
          }
          span_filter {
            parent_service {
              match     = "in"
              in_values = ["gateway", "frontdoor"]
            }
            tag {
              key = "env"
              value {
                match = "exact"
                value = "prod"
              }
            }
          }
        }
      }
    }
  }
}
