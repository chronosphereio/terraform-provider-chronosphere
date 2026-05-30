resource "chronosphere_email_alert_notifier" "oncall_team" {
  name = "On-Call Team Email"
  to   = "oncall@example.com"
}
