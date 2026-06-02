# Reference a service that was created out-of-band so its ID can be used as
# the collection_id for monitors, dashboards, and SLOs.
data "chronosphere_service" "gateway" {
  slug = "gateway"
}

resource "chronosphere_monitor" "gateway_up" {
  name          = "Gateway up"
  collection_id = data.chronosphere_service.gateway.id

  query {
    prometheus_expr = "up{service=\"gateway\"}"
  }

  series_conditions {
    condition {
      severity = "warn"
      value    = 1.0
      op       = "LT"
    }
  }
}
