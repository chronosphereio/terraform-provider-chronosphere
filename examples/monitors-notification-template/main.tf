resource "chronosphere_monitor" "monitor_with_notification_template" {
  name      = "Monitor With Notification Template"
  bucket_id = chronosphere_bucket.b.id
  query {
    prometheus_expr = <<-EOF
      sum by (kubernetes_namespace) (
        up{kubernetes_namespace="rc"}
      )
    EOF
  }
  series_conditions {
    condition {
      severity = "warn"
      value    = 20
      op       = "GT"
    }
  }

  notification_template {
    title       = "[{{.Severity}}] {{.Labels.service}} threshold exceeded"
    description = "{{.Labels.env}} crossed {{.ThresholdOp}} {{.Threshold}}"
  }
}
