resource "chronosphere_derived_metric" "api_success_rate" {
  name        = "API Success Rate"
  slug        = "api-success-rate"
  description = "Derived metric for API success rate across services"
  metric_name = "api_success_rate"

  # Query with label selectors and variables
  queries {
    selector {
      labels = {
        environment = "production"
        service     = "api"
      }
    }
    query {
      expr = "sum(rate(http_requests_total{status=~\"2..\"}[5m])) / sum(rate(http_requests_total[5m]))"
      variables {
        name             = "service"
        default_selector = "service=api"
      }
    }
  }

  # Default query without selector
  queries {
    query {
      expr = "rate(http_requests_total[5m])"
    }
  }
}
