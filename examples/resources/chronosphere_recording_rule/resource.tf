resource "chronosphere_recording_rule" "aggregated_requests" {
  bucket_id   = chronosphere_bucket.example.id
  name        = "Aggregated HTTP Requests by Namespace"
  metric_name = "http_requests:by_namespace:rate5m"
  expr        = "sum by (kubernetes_namespace) (rate(http_requests_total[5m]))"
  interval    = "60s"
  labels = {
    team  = "platform"
    type  = "aggregation"
  }
}

resource "chronosphere_recording_rule" "minimal" {
  name            = "CPU Usage Aggregation"
  execution_group = "infrastructure"
  expr            = "sum by (cluster) (rate(cpu_usage_seconds_total[5m]))"
}
