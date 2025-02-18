resource "chronosphere_victorops_alert_notifier" "test_victorops" {
  name        = "test-victorops"
  api_key     = "00000000-0000-0000-0000-000000000000"
  api_url     = "https://alert.victorops.com/integrations/generic/00000000/alert/"
  routing_key = "test"
}
