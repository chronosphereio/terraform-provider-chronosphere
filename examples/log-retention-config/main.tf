resource "chronosphere_log_retention_config" "production_logs" {
  name           = "Production Error Logs Long-term Retention"
  mode           = "ENABLED"
  filter         = "severity = 'error' AND env = 'production'"
  retention_days = 365
}

resource "chronosphere_log_retention_config" "audit_logs" {
  name           = "Audit Logs Retention"
  slug           = "audit-logs-retention"
  mode           = "ENABLED"
  filter         = "service = 'audit-service'"
  retention_days = 730
}

resource "chronosphere_log_retention_config" "disabled_example" {
  name           = "Disabled Retention Policy"
  mode           = "DISABLED"
  filter         = "service = 'test-service'"
  retention_days = 90
}
