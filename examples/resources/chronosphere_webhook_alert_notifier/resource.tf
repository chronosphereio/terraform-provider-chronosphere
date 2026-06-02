resource "chronosphere_webhook_alert_notifier" "webhook" {
  name          = "Webhook"
  url           = "http://example.com/url"
  send_resolved = false
  bearer_token  = "bearer-token"
}
