resource "chronosphere_bucket" "b" {
  name = "Platform"
}

resource "chronosphere_rollup_rule" "rollup_rule" {
  name        = "RollupRule"
  slug        = "rollup-rule"
  bucket_id   = chronosphere_bucket.b.id
  filter      = "__name__:metric_name"
  aggregation = "SUM"
  drop_raw    = true
  group_by    = ["service"]

  metric_type     = "COUNTER"
  metric_type_tag = false
  new_metric      = "new_metric_name"
  permissive      = true

  storage_policies {
    resolution = "30s"
    retention  = "120h"
  }

  mode = "PREVIEW"
}
