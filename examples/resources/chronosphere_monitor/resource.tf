resource "chronosphere_monitor" "high_cpu_usage" {
  name          = "High CPU Usage Alert"
  collection_id = chronosphere_collection.example.id

  query {
    prometheus_expr = <<-EOF
      sum by (kubernetes_namespace) (
        rate(container_cpu_usage_seconds_total[5m])
      ) > 0.8
    EOF
  }

  signal_grouping {
    label_names = ["kubernetes_namespace"]
  }

  series_conditions {
    condition {
      severity = "critical"
      value    = 0.9
      op       = "GT"
    }
    condition {
      severity = "warn"
      value    = 0.8
      op       = "GT"
    }
  }
}
