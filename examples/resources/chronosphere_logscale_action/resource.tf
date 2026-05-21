resource "chronosphere_logscale_action" "email" {
  repository = "default"
  name       = "Email on-call"

  email_action {
    recipients       = ["oncall@example.com"]
    subject_template = "Logscale alert: {{alert.name}}"
    body_template    = "{{query.results}}"
    attach_csv       = true
    use_proxy        = false
  }
}

resource "chronosphere_logscale_action" "pagerduty" {
  repository = "default"
  name       = "PagerDuty page"

  pager_duty_action {
    severity    = "ERROR"
    routing_key = "XXXXX"
    use_proxy   = false
  }
}
