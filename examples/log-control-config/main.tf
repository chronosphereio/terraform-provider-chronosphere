resource "chronosphere_log_control_config" "my-log-control-config" {
  rules {
    name   = "drop-debug-logs"
    mode   = "ENABLED"
    filter = "{level=\"debug\"}"
    type   = "DROP"
  }

  rules {
    name   = "sample-high-volume-logs"
    mode   = "ENABLED"
    filter = "{app=\"high-volume-app\"}"
    type   = "SAMPLE"
    sample {
      rate = 0.1
    }
  }

  rules {
    name   = "drop-sensitive-fields"
    mode   = "ENABLED"
    filter = "{app=\"sensitive-app\"}"
    type   = "DROP_FIELD"
    drop_field {
      field_regex = "password|secret|token"
      parent_path {
        selector = "metadata"
      }
    }
  }
}
