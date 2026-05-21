resource "chronosphere_logscale_action" "example_email_action" {
  repository = "default"
  name       = "email action"

  email_action {
    recipients       = ["email@example.com"]
    subject_template = "hi"
    body_template    = "you should look at this"
    attach_csv       = true
    use_proxy        = false
  }
}

resource "chronosphere_logscale_action" "example_humio_action" {
  repository = "default"
  name       = "log to another repo action"

  humio_action {
    ingest_token = "another-repo-ingest-token"
  }
}

resource "chronosphere_logscale_action" "example_ops_genie_action" {
  repository = "default"
  name       = "ops genie action"

  ops_genie_action {
    api_url = "https://api.opsgenie.com/your-url"
    ops_genie_key = "key"
    use_proxy = false
  }
}

resource "chronosphere_logscale_action" "example_pagerduty_action" {
  repository = "default"
  name       = "pagerduty action"

  pager_duty_action {
    severity = "ERROR"
    routing_key = "routing_key"
    use_proxy = false
  }
}

resource "chronosphere_logscale_action" "example_slack_action" {
  repository = "default"
  name       = "slack action"

  slack_action {
    url = "https://slack.com/your-url"
    fields = {
      "field1": "value1",
    }
    use_proxy = false
  }
}

resource "chronosphere_logscale_action" "example_slack_post_message_action" {
  repository = "default"
  name       = "slack post message action"

  slack_post_message_action {
    api_token = "slack-api-token"
    channels = ["slack-channel"]
    fields = {
      "field1": "value1",
    }
    use_proxy = false
  }
}

resource "chronosphere_logscale_action" "example_victor_ops_action" {
  repository = "default"
  name       = "victor ops action"

  victor_ops_action {
    message_type = "message_type"
    notify_url = "https://victorops.com/your-url"
    use_proxy = false
  }
}

resource "chronosphere_logscale_action" "example_upload_file_action" {
  repository = "default"
  name       = "upload file action"

  upload_file_action {
    file_name = "file.csv"
  }
}

resource "chronosphere_logscale_action" "example_webhook_action" {
  repository = "default"
  name       = "webhook action"

  webhook_action {
    method = "POST"
    url = "https://chronosphere.io/notify/webhook"
    headers = {
      "Header-1": "value1"
    }
    body_template = "this will get posted"
    ignore_ssl = false
    use_proxy = false
  }
}
