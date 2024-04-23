resource "chronosphere_dataset" "example_prod_payfe" {
  name        = "${var.prefix} Production paymentfe Traces"
  description = "Traces passing through the paymentfe service in production"
  configuration {
    type = "TRACES"

    trace_dataset {
      match_criteria {
        span {
          duration {
            max_secs = 99
            min_secs = 1
          }

          error {
            value = true
          }

          match_type = "INCLUDE"

          operation {
            value = "importantop"
            match = "EXACT"
          }

          parent_operation {
            value = "payments/.*"
            match = "REGEX"
          }

          parent_service {
            value = "frontdoor-[east|west]"
            match = "REGEX"
          }

          service {
            value = "importantsvc"
            match = "EXACT"
          }

          span_count {
            max = 2
            min = 1
          }

          tag {
            key = "cool_tag"

            value {
              value = "coolvalue"
              match = "EXACT"
            }
          }

          tag {
            key = "env_tag"

            value {
              value = "prod.*"
              match = "REGEX"
            }
          }

          tag {
            key = "http.status_code"

            numeric_value {
              comparison = "GREATER_THAN"
              value      = 299
            }
          }
        }

        trace {
          duration {
            max_secs = 10
            min_secs = 5
          }

          error {
            value = true
          }
        }
      }
    }
  }
}
