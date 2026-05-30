resource "chronosphere_classic_dashboard" "legacy_dashboard" {
  name          = "Legacy Grafana Dashboard"
  collection_id = chronosphere_collection.example.id

  dashboard_json = jsonencode({
    title : "Service Metrics",
    panels : [
      {
        title : "CPU Usage",
        targets : [{
          expr : "rate(container_cpu_usage_seconds_total[5m])"
        }]
      }
    ]
  })
}
