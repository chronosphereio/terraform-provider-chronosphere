resource "chronosphere_team" "t" {
  name = "Synthetics Team"
}

resource "chronosphere_collection" "c" {
  name    = "Synthetics Collection"
  team_id = chronosphere_team.t.id
}

resource "chronosphere_blackhole_alert_notifier" "blackhole" {
  name = "Blackhole"
}

resource "chronosphere_notification_policy" "np" {
  team_id = chronosphere_team.t.id
  name    = "Synthetics team NP"

  route {
    severity  = "warn"
    notifiers = [chronosphere_blackhole_alert_notifier.blackhole.id]
  }
}

resource "chronosphere_synthetic_test" "checkout_http" {
  name          = "Checkout availability"
  collection_id = chronosphere_collection.c.id
  test_type     = "HTTP"
  status        = "ENABLED"
  locations     = ["GCP_US_OREGON", "GCP_US_VIRGINIA"]
  interval_secs = 60
  timeout_secs  = 30

  labels = {
    service = "checkout"
  }

  retry_config {
    max_retries       = 2
    retry_interval_ms = 500
  }

  monitor_config {
    failing_duration_secs  = 300
    min_failing_locations  = 2
    notification_policy_id = chronosphere_notification_policy.np.id
  }

  http_test {
    url          = "https://example.com/checkout/health"
    method       = "POST"
    http_version = "HTTP_2"
    content_type = "APPLICATION_JSON"
    request_body = jsonencode({ probe = true })

    headers {
      name  = "X-Probe-Source"
      value = "chronosphere-synthetics"
    }

    follow_redirects        = true
    max_redirects           = 3
    max_response_body_bytes = 51200

    authentication {
      basic_auth {
        username            = "probe"
        password_wo         = "example-password"
        password_wo_version = 1
      }
    }

    assertions {
      status_code_assertion {
        operator = "EQUALS"
        target   = "200"
      }
    }

    assertions {
      body_json_path_assertion {
        json_path  = "$.status"
        operator   = "EQUALS"
        target     = "ok"
        match_type = "FIRST_ELEMENT"
      }
    }

    assertions {
      response_time_assertion {
        operator  = "LESS_THAN"
        target_ms = 1500
        scope     = "WITHOUT_DNS"
      }
    }
  }
}

resource "chronosphere_synthetic_test" "checkout_dns" {
  name          = "Checkout DNS"
  collection_id = chronosphere_collection.c.id
  test_type     = "DNS"
  status        = "ENABLED"
  interval_secs = 300
  locations     = ["GCP_US_OREGON"]

  dns_test {
    domain = "example.com"

    assertions {
      dns_record_assertion {
        record_type = "A"
        operator    = "EQUALS"
        target      = "93.184.216.34"
        match_scope = "AT_LEAST_ONE"
      }
    }
  }
}

resource "chronosphere_synthetic_test" "checkout_tls" {
  name          = "Checkout TLS"
  collection_id = chronosphere_collection.c.id
  test_type     = "TLS"
  status        = "ENABLED"
  interval_secs = 300
  locations     = ["GCP_US_OREGON"]

  tls_test {
    host = "example.com"
    port = 443

    assertions {
      certificate_assertion {
        operator    = "EXPIRES_IN_MORE_THAN_DAYS"
        target_days = 14
      }
    }

    assertions {
      tls_version_assertion {
        bound    = "MIN"
        operator = "GREATER_THAN_OR_EQUAL"
        target   = "TLS_1_2"
      }
    }
  }
}
