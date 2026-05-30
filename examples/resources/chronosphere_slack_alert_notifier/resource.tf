resource "chronosphere_slack_alert_notifier" "alerts_channel" {
  name    = "Production Alerts Slack"
  channel = "#prod-alerts"
  token   = var.slack_token
}
