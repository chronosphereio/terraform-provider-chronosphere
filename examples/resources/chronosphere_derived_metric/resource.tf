resource "chronosphere_derived_metric" "request_rate" {
  name        = "request_rate"
  slug        = "request-rate"
  description = "Per-service request rate, with selector-aware variants"
  metric_name = "request_rate"

  # Specialized query when the label1 selector matches.
  queries {
    selector {
      labels = {
        label1 = "value1"
      }
    }
    query {
      expr = "sum by (service) (rate(http_requests_total{label1=\"value1\"}[5m]))"
      variables {
        name             = "service"
        default_selector = "service=default"
      }
    }
  }

  # Default query when no selector matches.
  queries {
    query {
      expr = "sum by (service) (rate(http_requests_total[5m]))"
    }
  }
}
