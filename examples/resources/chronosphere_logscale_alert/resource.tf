resource "chronosphere_logscale_alert" "error_alert" {
  repository        = "default"
  name              = "High Error Rate Alert"
  description       = "Alert when error rate exceeds threshold"
  alert_type        = "STANDARD"
  query             = "level = ERROR | count(as=numErrors) | numErrors > 500"
  time_window       = "60s"
  throttle_duration = "300s"
  throttle_field    = "service"
  tags = [
    "production",
    "errors",
  ]
  run_as_user = "alerts@example.com"
  disabled    = false
  action_ids = [
    chronosphere_logscale_action.email_action.id,
  ]
}

resource "chronosphere_logscale_alert" "filter_alert" {
  repository        = "default"
  name              = "Security Event Filter"
  description       = "Alert on security-related log events"
  alert_type        = "FILTER"
  query             = "level = ERROR AND category = security"
  throttle_duration = "60s"
  throttle_field    = "user_id"
  tags = [
    "security",
  ]
  run_as_user = "security@example.com"
  disabled    = false
  action_ids = [
    chronosphere_logscale_action.email_action.id,
  ]
}
