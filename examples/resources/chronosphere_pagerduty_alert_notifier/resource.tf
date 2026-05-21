resource "chronosphere_pagerduty_alert_notifier" "pagerduty" {
  name          = "PagerDuty Notifier"
  severity      = "info"
  url           = "https://events.pagerduty.com/v2/enqueue"
  routing_key   = "XXXXX"
  send_resolved = true

  details = {
    "runbook" = "http://runbook"
  }
}
