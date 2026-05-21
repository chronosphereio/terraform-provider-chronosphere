resource "chronosphere_log_ingest_config" "config" {
  # Parse raw JSON log lines into structured fields.
  field_parser {
    mode = "ENABLED"

    source {
      selector = "raw_message"
    }

    parser {
      parser_type = "JSON"
    }
  }

  # Normalize common fields across heterogeneous log sources.
  field_normalization {
    timestamp {
      source {
        selector = "timestamp"
      }
      source {
        selector = "ts"
      }
      source {
        selector = "@timestamp"
      }
    }

    service {
      source {
        selector = "service"
      }
      source {
        selector = "app"
      }
      default_value = "UNKNOWN"
    }

    severity {
      source {
        selector = "level"
      }
      source {
        selector = "severity"
      }
      value_map = {
        "debug" = "DEBUG"
        "info"  = "INFO"
        "warn"  = "WARNING"
        "error" = "ERROR"
        "fatal" = "CRITICAL"
      }
      default_value = "INFO"
    }

    message {
      source {
        selector = "message"
      }
      source {
        selector = "msg"
      }
      default_value = "no message"
    }
  }
}
