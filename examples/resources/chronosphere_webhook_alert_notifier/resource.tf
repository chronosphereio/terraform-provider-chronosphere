resource "chronosphere_webhook_alert_notifier" "custom_endpoint" {
  name = "Custom Webhook Notifier"
  url  = "https://api.example.com/webhooks/alerts"
}
