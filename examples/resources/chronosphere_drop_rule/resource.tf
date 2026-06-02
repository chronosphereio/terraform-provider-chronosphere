resource "chronosphere_drop_rule" "noisy_metric" {
  name = "Drop noisy metric"
  query = [
    "__name__:noisy_metric_name",
  ]
  mode = "ENABLED"
}
