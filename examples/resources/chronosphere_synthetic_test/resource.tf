resource "chronosphere_team" "platform" {
  name = "Platform"
}

resource "chronosphere_collection" "checkout" {
  name    = "Checkout"
  team_id = chronosphere_team.platform.id
}

variable "checkout_probe_password" {
  type      = string
  sensitive = true
}

resource "chronosphere_synthetic_test" "checkout_availability" {
  name          = "Checkout availability"
  collection_id = chronosphere_collection.checkout.id
  test_type     = "HTTP"
  status        = "ENABLED"
  locations     = ["GCP_US_OREGON", "GCP_US_VIRGINIA"]
  interval_secs = 60
  timeout_secs  = 30

  retry_config {
    max_retries       = 2
    retry_interval_ms = 500
  }

  http_test {
    url          = "https://example.com/checkout/health"
    method       = "GET"
    http_version = "HTTP_2"

    authentication {
      basic_auth {
        username = "probe"
        # Write-only: supplied on every apply, never written to state. Bump
        # password_wo_version to roll out a new password.
        password_wo         = var.checkout_probe_password
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
      response_time_assertion {
        operator  = "LESS_THAN"
        target_ms = 1500
        scope     = "WITHOUT_DNS"
      }
    }
  }
}
