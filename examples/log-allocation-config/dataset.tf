resource "chronosphere_dataset" "example_logs_dataset" {
  name        = "Production Logs"
  slug        = "example_logs_dataset"
  configuration {
    type = "LOGS"

    log_dataset {
      match_criteria {
        query = "env = 'prod'"
      }
    }
  }
}
