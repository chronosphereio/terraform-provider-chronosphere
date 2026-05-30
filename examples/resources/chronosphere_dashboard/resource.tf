resource "chronosphere_dashboard" "service_overview" {
  name          = "Service Overview Dashboard"
  slug          = "service-overview"
  collection_id = chronosphere_collection.example.id

  labels = {
    team        = "platform"
    environment = "production"
  }

  dashboard_json = jsonencode({
    kind : "Dashboard",
    spec : {
      events : [],
      panels : {
        cpu_panel : {
          type : "timeseries",
          queries : [{
            expr : "rate(container_cpu_usage_seconds_total[5m])"
          }]
        }
      },
      layouts : [],
      variables : [],
      duration : "1h"
    },
    spec_version : "1"
  })
}
