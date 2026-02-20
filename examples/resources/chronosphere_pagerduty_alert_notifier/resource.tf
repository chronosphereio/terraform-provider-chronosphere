resource "chronosphere_pagerduty_alert_notifier" "oncall" {
  name = "PagerDuty On-Call"
  key  = var.pagerduty_integration_key
}
