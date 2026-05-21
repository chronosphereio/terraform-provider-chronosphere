resource "chronosphere_webhook_external_connection" "webhook" {
  name         = "Webhook"
  url          = "https://example.com/notify"
  bearer_token = "XXXXX"
}
