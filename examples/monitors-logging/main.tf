resource "chronosphere_monitor" "monitor_with_logging_query" {
  name      = "Monitor With Logging Query"
  collection_id = chronosphere_collection.infra.id
  query {
    logging_expr = <<-EOF
      service = "nginx" | make-series by chronosphere_namespace step 1m
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
