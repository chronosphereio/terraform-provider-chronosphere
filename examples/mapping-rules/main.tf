
resource "chronosphere_mapping_rule" "http_request_duration" {
  name      = "http request duration"
  bucket_id = chronosphere_bucket.b.id
  filter    = "__name__:http_request_duration k8s_pod:*"

  # See https://docs.chronosphere.io/documentation/user/metrics/aggregation-rules#supported-aggregation-operations for supported values.
  aggregations = [
    "LAST",
  ]

  storage_policy {
    resolution = "30s"
    retention  = "120h"
  }

  mode = "PREVIEW"
}