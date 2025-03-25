resource "chronosphere_opsgenie_alert_notifier" "test_opsgenie" {
  name          = "infra_compute_opsgenie"
  api_key       = "XXXXX"
  api_url       = "https://api.opsgenie.com/"
  send_resolved = true
  priority      = "P1"

  responder {
    name = "Productivity Platform - Compute"
    type = "TEAM" # Or: USER, ESCALATION, SCHEDULE
  }
}
