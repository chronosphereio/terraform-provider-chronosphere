resource "chronosphere_trace_metrics_rule" "payments_latency" {
  name = "Payments service latency"
  slug = "payments-latency"

  metric_name = "payments_request_duration"
  histogram_buckets_seconds = [
    0.1,
    0.5,
    1,
    2,
    5,
  ]

  metric_labels = {
    "service" = "payments"
  }

  trace_filter {
    span {
      match_type = "include"
      service {
        match = "exact"
        value = "payments"
      }
    }
  }

  group_by = ["operation"]
}
