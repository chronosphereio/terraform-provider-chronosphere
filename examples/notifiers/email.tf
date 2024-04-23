resource "chronosphere_email_alert_notifier" "email" {
  name = "${var.prefix} Email Blackhole"
  to   = "blackhole@chronosphere.io"
}
