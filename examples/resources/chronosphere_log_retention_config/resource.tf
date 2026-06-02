resource "chronosphere_log_retention_config" "production_errors" {
  name           = "Production Error Logs Long-term Retention"
  mode           = "ENABLED"
  filter         = "severity = 'error' AND env = 'production'"
  retention_days = 365
}
