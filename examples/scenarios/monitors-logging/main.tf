resource "chronosphere_monitor" "monitor_with_logging_query" {
  name      = "Monitor With Logging Query"
  collection_id = chronosphere_collection.infra.id
  query {
    logging_expr = <<-EOF
      service = "nginx" | make-series step 1m by chronosphere_namespace
    EOF
  }
  series_conditions {
    condition {
      severity = "warn"
      value    = 20
      op       = "GT"
    }
  }
}
