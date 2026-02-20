resource "chronosphere_mapping_rule" "latency_histogram" {
  name      = "HTTP Request Latency Mapping"
  bucket_id = chronosphere_bucket.example.id
  filter    = "__name__:http_request_duration_seconds k8s_pod:*"

  aggregations = [
    "LAST",
    "P99",
  ]

  storage_policy {
    resolution = "30s"
    retention  = "120h"
  }

  mode = "PREVIEW"
}
