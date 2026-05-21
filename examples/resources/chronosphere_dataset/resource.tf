resource "chronosphere_dataset" "payments_traces" {
  name        = "Production payments traces"
  description = "Traces passing through the payments service in production"

  configuration {
    type = "TRACES"

    trace_dataset {
      match_criteria {
        span {
          match_type = "INCLUDE"

          service {
            match = "EXACT"
            value = "payments"
          }

          duration {
            min_secs = 1
          }

          error {
            value = true
          }
        }
      }
    }
  }
}
