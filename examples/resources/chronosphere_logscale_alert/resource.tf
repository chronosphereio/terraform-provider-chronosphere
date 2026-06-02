resource "chronosphere_logscale_action" "email" {
  repository = "default"
  name       = "Email on-call"

  email_action {
    recipients       = ["oncall@example.com"]
    subject_template = "Logscale alert: {{alert.name}}"
    body_template    = "{{query.results}}"
  }
}

resource "chronosphere_logscale_alert" "high_error_rate" {
  repository  = "default"
  name        = "High error rate"
  description = "More than 500 errors in a 60s window"
  alert_type  = "STANDARD"

  query             = "level = ERROR | count(as=numErrors) | numErrors > 500"
  time_window       = "60s"
  throttle_duration = "60s"
  throttle_field    = "service"

  tags     = ["errors", "platform"]
  disabled = false

  action_ids = [
    chronosphere_logscale_action.email.id,
  ]
}
