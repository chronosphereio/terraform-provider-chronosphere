resource "chronosphere_logscale_action" "example_email_action" {
  repository = "default"
  name       = "email action 2"

  email_action {
    recipients       = ["email2@example.com"]
    subject_template = "hi"
    body_template    = "you should look at this"
    attach_csv       = true
    use_proxy        = false
  }
}

resource "chronosphere_logscale_alert" "example_standard_alert" {
  repository = "default"
  name = "Standard Alert"
  description = "Standard Alert"
  alert_type = "STANDARD"
  query = "level = ERROR | count(as=numErrors) | numErrors > 500"
  time_window = "60s"
  throttle_duration = "60s"
  throttle_field = "some_field_to_throttle_by"
  tags = [
    "tag1",
    "tag2",
  ]
  run_as_user = "example@chronosphere.io"
  disabled = false
  action_ids = [
    chronosphere_logscale_action.example_email_action.id,
  ]
}

resource "chronosphere_logscale_alert" "example_filter_alert" {
  repository = "default"
  name = "Filter Alert"
  description = "Filter Alert"
  alert_type = "FILTER"
  query = "level = ERROR"
  throttle_duration = "60s"
  throttle_field = "some_field_to_throttle_by"
  tags = [
    "tag1",
    "tag2",
  ]
  run_as_user = "example@chronosphere.io"
  disabled = false
  action_ids = [
    chronosphere_logscale_action.example_email_action.id,
  ]
}
