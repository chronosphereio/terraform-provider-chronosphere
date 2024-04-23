resource "chronosphere_monitor" "collection_monitor" {
  name          = "${var.prefix} Monitor in Collection"
  collection_id = chronosphere_collection.infra.id

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
      value    = 20.0
      op       = "GT"
    }
  }
}
