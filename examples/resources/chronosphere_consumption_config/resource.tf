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
        # trace_span_filter blocks AND together: a single span must
        # satisfy every block (and every filter within a block) to match
        # the condition.
        trace_span_filter {
          service {
            match = "exact"
            value = "my-service"
          }
          error {
            value = true
          }
          duration {
            min_secs = 0.5
          }
        }
      }
      # Express alternatives with an `in` matcher or with a separate
      # condition block, like this one.
      condition {
        trace_span_filter {
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
