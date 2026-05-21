resource "chronosphere_log_ingest_config" "my-log-ingest-config" {
  plaintext_parser {
    name = "syslog"
    mode = "ENABLED"

    parser {
      parser_type = "REGEX"
      regex_parser {
        regex = <<-EOT
          ^\<(?<pri>[0-9]+)\>(?<time>[^ ]* {1,2}[^ ]* [^ ]*) (?<ident>[a-zA-Z0-9_\/\.\-]*)(?:\[(?<pid>[0-9]+)\])?(?:[^\:]*\:)? *(?<message>.*)$
        EOT
      }
    }
  }

  plaintext_parser {
    name = "apache_error"
    mode = "ENABLED"
    keep_original = true

    parser {
      parser_type = "REGEX"
      regex_parser {
        regex = <<-EOT
          ^\[[^ ]* (?<time>[^\]]*)\] \[(?<level>[^\]]*)\](?: \[pid (?<pid>[^\]]*)\])?( \[client (?<client>[^\]]*)\])? (?<message>.*)$
        EOT
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

  field_parser {
    mode = "ENABLED"

    source {
      selector = "kv_details"
    }

    destination {
      selector = "structured_details"
    }

    parser {
      parser_type = "KEY_VALUE"
      key_value_parser {
        pair_separator = ":"
        delimiter = ","
        trim_set = " {}"
      }
    }
  }

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
        selector = "app"
      }
      source {
        selector = "service"
      }
      default_value = "UNKNOWN"
      sanitize_patterns = ["^[a-zA-Z0-9_-]+$"]
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
      sanitize_patterns = ["^\\[.*\\] (.*)$"]
      default_value = "no message"
    }

    custom_field_normalization {
      normalization {
        source {
          selector = "env"
        }
        source {
          selector = "environment"
        }
        value_map = {
          "dev"  = "development"
          "stg"  = "staging"
          "prod" = "production"
        }
        default_value = "development"
      }
      target = "environment"
    }

    custom_field_normalization {
      normalization {
        source {
          selector = "region"
        }
        source {
          selector = "datacenter"
        }
        sanitize_patterns = ["^dc-(.*)$"]
        default_value = "us-east-1"
      }
      target = "region"
    }
  }
}
