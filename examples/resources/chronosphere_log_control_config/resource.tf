resource "chronosphere_log_control_config" "example" {
  rules {
    name   = "Sample Debug Logs"
    mode   = "ENABLED"
    filter = "level = 'debug'"
    type   = "SAMPLE"
    sample {
      rate = 10
    }
  }

  rules {
    name   = "Drop Health Checks"
    mode   = "ENABLED"
    filter = "path = '/health'"
    type   = "DROP"
  }

  rules {
    name   = "Drop Sensitive Fields"
    mode   = "ENABLED"
    filter = "service = 'api'"
    type   = "DROP_FIELD"
    drop_field {
      field_regex = "password|token|secret"
    }
  }

  rules {
    name   = "Hash User IDs"
    mode   = "ENABLED"
    filter = "service = 'user-service'"
    type   = "REPLACE_FIELD"
    replace_field {
      field {
        selector = "user_id"
      }
      replace_regex = ".*"
      replace_all   = true
      replace_mode  = "HASH"
    }
  }

  rules {
    name   = "Emit Error Metrics"
    mode   = "ENABLED"
    filter = "level = 'error'"
    type   = "EMIT_METRICS"
    emit_metrics {
      name = "log_errors_total"
      mode = "COUNTER"
      counter {
        value {
          selector = "count"
        }
      }
      labels {
        key = "service"
        value {
          selector = "service"
        }
      }
    }
  }
}
