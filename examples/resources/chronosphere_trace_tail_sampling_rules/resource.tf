resource "chronosphere_trace_tail_sampling_rules" "production_sampling" {
  default_sample_rate {
    enabled     = true
    sample_rate = 0.1
  }

  rules {
    name        = "Sample All Errors"
    system_name = "sample_all_errors"
    sample_rate = 1.0

    filter {
      trace {
        error {
          value = true
        }
      }
    }
  }

  rules {
    name        = "Sample Slow Requests"
    system_name = "sample_slow_requests"
    sample_rate = 0.5

    filter {
      span {
        match_type = "INCLUDE"
        service {
          match = "EXACT"
          value = "api"
        }
        duration {
          min_secs = 5
        }
      }
    }
  }

  rules {
    name        = "Sample Specific Operations"
    system_name = "sample_specific_operations"
    sample_rate = 0.25

    filter {
      span {
        match_type = "INCLUDE"
        operation {
          match = "REGEX"
          value = "checkout.*"
        }
        tag {
          key = "http.status_code"
          numeric_value {
            comparison = "GREATER_THAN_OR_EQUAL"
            value      = 500
          }
        }
      }
    }
  }
}
