resource "chronosphere_dataset" "production_logs" {
  name        = "Production Application Logs"
  description = "Logs from production services"
  configuration {
    type = "LOGS"

    log_dataset {
      match_criteria {
        query = "service = 'api' AND env = 'production'"
      }
    }
  }
}

resource "chronosphere_dataset" "error_traces" {
  name        = "Error Traces"
  description = "Traces containing errors"
  configuration {
    type = "TRACES"

    trace_dataset {
      match_criteria {
        span {
          match_type = "INCLUDE"
          error {
            value = true
          }
          service {
            value = "api"
            match = "EXACT"
          }
          duration {
            min_secs = 1
            max_secs = 60
          }
        }

        trace {
          error {
            value = true
          }
        }
      }
    }
  }
}
