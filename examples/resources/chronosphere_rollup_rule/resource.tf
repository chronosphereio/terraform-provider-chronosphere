resource "chronosphere_rollup_rule" "request_rollup" {
  name        = "HTTP Request Counter Rollup"
  slug        = "http-request-rollup"
  bucket_id   = chronosphere_bucket.example.id
  filter      = "__name__:http_requests_total"
  aggregation = "SUM"
  group_by = [
    "service",
    "status_code",
  ]
  metric_type     = "COUNTER"
  metric_type_tag = false
  new_metric      = "http_requests_total_rollup"
  drop_raw        = true
  permissive      = true
  storage_policies {
    resolution = "30s"
    retention  = "120h"
  }
  mode = "PREVIEW"
}
