resource "chronosphere_webhook_alert_notifier" "webhook" {
  name          = "${var.prefix} Webhook"
  url           = "http://example.com/url"
  send_resolved = false
  bearer_token  = "bearer-token"
}
