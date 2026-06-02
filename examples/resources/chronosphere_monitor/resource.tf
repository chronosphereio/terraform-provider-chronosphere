resource "chronosphere_collection" "c" {
  name = "Platform"
}

resource "chronosphere_monitor" "namespace_up" {
  name          = "Namespace up"
  collection_id = chronosphere_collection.c.id

  query {
    prometheus_expr = <<-EOF
      sum by (kubernetes_namespace) (
        up{kubernetes_namespace="production"}
      )
    EOF
  }

  signal_grouping {
    label_names = ["kubernetes_namespace"]
  }

  series_conditions {
    condition {
      severity = "warn"
      value    = 20.0
      op       = "GT"
    }
  }
}
