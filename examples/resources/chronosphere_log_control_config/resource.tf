resource "chronosphere_log_control_config" "config" {
  # Sample 1% of debug-level logs from the sample service.
  rules {
    name   = "sample-debug"
    mode   = "ENABLED"
    filter = "service = 'sample-service' AND severity = 'debug'"
    type   = "SAMPLE"

    sample {
      rate = 0.01
    }
  }

  # Drop logs from a deprecated service entirely.
  rules {
    name   = "drop-deprecated"
    mode   = "ENABLED"
    filter = "service = 'deprecated-service'"
    type   = "DROP"
  }

  # Redact sensitive fields nested under kubernetes labels.
  rules {
    name   = "drop-sensitive-fields"
    mode   = "ENABLED"
    filter = "service = 'api-gateway'"
    type   = "DROP_FIELD"

    drop_field {
      field_regex = "password|secret|api_key"
      parent_path {
        selector = "kubernetes['labels']"
      }
    }
  }

  # Replace long trace IDs with a placeholder to reduce log volume.
  rules {
    name   = "shorten-trace-ids"
    mode   = "ENABLED"
    filter = "service = 'api-gateway'"
    type   = "REPLACE_FIELD"

    replace_field {
      field {
        selector = "trace_id"
      }
      replace_regex = "[0-9a-f]{32}"
      replace_all   = false
      replace_mode  = "STATIC_VALUE"
      static_value {
        value = "[trace-id]"
      }
    }
  }
}
