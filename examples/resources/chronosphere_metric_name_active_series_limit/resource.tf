resource "chronosphere_metric_name_active_series_limit" "http_requests_total" {
  name              = "http_requests_total active series limit"
  slug              = "http-requests-total-active-series-limit"
  metric_name       = "http_requests_total"
  max_active_series = 100000
}
