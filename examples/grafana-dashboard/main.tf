resource "chronosphere_grafana_dashboard" "my_dashboard" {
  bucket_id = chronosphere_bucket.b.id
  dashboard_json = jsonencode({
    title : "Dashboard",
    panels : [{
      "gridPos" : {
        "h" : 12,
        "w" : 24,
        "x" : 0,
        "y" : 0
      },
      id : 2,
      targets : [
        {
          expr : "1",
        }
      ],
      title : "Panel Title",
      type : "graph",
    }],
  })
}
