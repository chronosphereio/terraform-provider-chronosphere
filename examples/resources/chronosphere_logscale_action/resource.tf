resource "chronosphere_logscale_action" "email_action" {
  repository = "default"
  name       = "Email Alert Action"

  email_action {
    recipients       = ["oncall@example.com"]
    subject_template = "Alert: {{.Query.Name}}"
    body_template    = "Alert triggered at {{.Time}}: {{.Query.QueryString}}"
    attach_csv       = true
    use_proxy        = false
  }
}
