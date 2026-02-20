resource "chronosphere_log_ingest_config" "example" {
  plaintext_parser {
    name = "syslog"
    mode = "ENABLED"

    parser {
      parser_type = "REGEX"
      regex_parser {
        regex = "^\\<(?<pri>[0-9]+)\\>(?<time>[^ ]* {1,2}[^ ]* [^ ]*) (?<ident>[a-zA-Z0-9_\\/\\.\\-]*)(?:\\[(?<pid>[0-9]+)\\])?(?:[^\\:]*\\:)? *(?<message>.*)$"
      }
    }
  }

  field_parser {
    mode = "ENABLED"

    source {
      selector = "raw_message"
    }

    parser {
      parser_type = "JSON"
    }
  }

  field_normalization {
    primary_key {
      normalization {
        source {
          selector = "service_name"
        }
        source {
          selector = "app"
        }
        default_value = "unknown"
      }
      target = "service"
    }

    timestamp {
      source {
        selector = "timestamp"
      }
      source {
        selector = "@timestamp"
      }
    }

    severity {
      source {
        selector = "level"
      }
      value_map = {
        "debug" = "DEBUG"
        "info"  = "INFO"
        "warn"  = "WARNING"
        "error" = "ERROR"
      }
      default_value = "INFO"
    }

    message {
      source {
        selector = "message"
      }
      default_value = "no message"
    }
  }
}
