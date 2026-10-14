resource "chronosphere_metric_name_default_active_series_limit" "default" {
  default_max_active_standard_series  = 50000
  default_max_active_histogram_series = 5000
}
