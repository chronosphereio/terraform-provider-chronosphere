resource "chronosphere_monitor" "monitor_with_signal" {
  name      = "${var.prefix} Monitor With Signals"
  bucket_id = chronosphere_bucket.b.id
  query {
    prometheus_expr = <<-EOF
      sum by (kubernetes_namespace) (
        up{kubernetes_namespace="rc"}
      )
    EOF
  }
  signal_grouping {
    label_names = ["kubernetes_namespace"]
  }
  series_conditions {
    condition {
      severity = "warn"
      value    = 20
      op       = "GT"
    }

    override {
      label_matcher {
        name  = "app"
        type  = "EXACT_MATCHER_TYPE"
        value = "dbmon"
      }

      condition {
        severity        = "critical"
        value           = 1.0
        op              = "GT"
        sustain         = "60s"
        resolve_sustain = "30s"
      }
    }
  }

  labels = {
    "team" = "my team"
  }
  annotations = {
    "runbook" = "go/runbook"
  }

  schedule {
    timezone = "UTC"

    dynamic "range" {
      for_each = toset(["Monday", "Tuesday", "Wednesday", "Thursday", "Friday"])
      content {
        day   = range.key
        start = "07:00"
        end   = "20:00"
      }
    }
  }
}
