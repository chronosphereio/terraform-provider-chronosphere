resource "chronosphere_pagerduty_alert_notifier" "test_pagerduty" {
  name          = "Pagerduty Notifier"
  severity      = "info"
  url           = "https://events.pagerduty.com/v2/enqueue"
  service_key   = "XXXXX"
  send_resolved = true
  details = {
    "runbook" = "http://runbook"
  }
}

