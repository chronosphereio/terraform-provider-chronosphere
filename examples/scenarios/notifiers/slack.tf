resource "chronosphere_slack_alert_notifier" "test_slack" {
  name    = "Slack Notifier"
  api_url = "https://hooks.slack.com/services/XXXXX/XXXXX/XXXXX"
  channel = "alerts"
}
