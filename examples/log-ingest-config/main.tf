resource "chronosphere_log_ingest_config" "my-log-ingest-config" {
  plaintext_parser {
    name = "syslog"
    mode = "ENABLED"
    keep_original = true
    
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
    keep_original = false
    
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
}
