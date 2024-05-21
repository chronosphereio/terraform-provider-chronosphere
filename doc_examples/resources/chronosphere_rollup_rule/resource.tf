resource "chronosphere_bucket" "b" {
  name        = "Bucket"
  description = "bucket created by terraform examples"
  labels      = { "foo" : "bar" }
}

resource "chronosphere_rollup_rule" "rollup_rule" {
  aggregation = "SUM"
  bucket_id   = chronosphere_bucket.b.id
  drop_raw    = true
  filter      = "__name__:metric_name"
  group_by = [
    "service",
  ]
  metric_type     = "COUNTER"
  metric_type_tag = false
  name            = "RollupRule"
  new_metric      = "new_metric_name"
  permissive      = true
  slug            = "rollup-rule"
  storage_policies {
    resolution = "30s"
    retention  = "120h"
  }
  mode = "PREVIEW"
}
