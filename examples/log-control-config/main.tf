resource "chronosphere_log_control_config" "my-log-control-config" {
  rules {
    name   = "sample"
    mode   = "ENABLED"
    filter = "service = 'sample_service'"
    type   = "SAMPLE"

    sample {
      rate = 1
    }
  }

  rules {
    name   = "drop logs"
    mode   = "ENABLED"
    filter = "service = 'another-service'"
    type   = "DROP"
  }

  rules {
    name   = "drop field"
    mode   = "ENABLED"
    filter = "service = 'third-service'"
    type   = "DROP_FIELD"

    drop_field {
      field_regex = "not|important|key"
    }
  }

  rules {
    name   = "drop field with parent path"
    mode   = "ENABLED"
    filter = "service = 'fourth-service'"
    type   = "DROP_FIELD"

    drop_field {
      field_regex = "sensitive|secret|password"
      parent_path {
        selector = "kubernetes['container']['pod']"
      }
    }
  }

  rules {
    name   = "hash-long-request-ids"
    mode   = "ENABLED"
    filter = "service = 'api-gateway'"
    type   = "REPLACE_FIELD"
    replace_field {
      field {
        selector = "request_id"
      }
      replace_regex = ".*"
      replace_all   = true
      replace_mode  = "HASH"
    }
  }

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
      replace_mode  = "FIXED_VALUE"
      fixed_value {
        value = "[trace-id]"
      }
    }
  }

  rules {
    name   = "shorten-region-names"
    mode   = "ENABLED"
    filter = "service = 'cloud-service'"
    type   = "REPLACE_FIELD"
    replace_field {
      field {
        selector = "region"
      }
      replace_regex = ".*"
      replace_all   = true
      replace_mode  = "MAPPED_VALUE"
      mapped_value {
        use_default   = true
        default_value = "other"
        pairs {
          key   = "us-east-1"
          value = "use1"
        }
        pairs {
          key   = "us-west-2"
          value = "usw2"
        }
        pairs {
          key   = "eu-central-1"
          value = "euc1"
        }
      }
    }
  }
}
