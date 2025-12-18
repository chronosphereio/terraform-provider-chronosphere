resource "chronosphere_opsgenie_alert_notifier" "production" {
  name   = "Opsgenie Production Alerts"
  api_key = var.opsgenie_api_key
}
