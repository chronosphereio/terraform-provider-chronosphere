resource "chronosphere_dashboard" "my_dashboard" {
  slug          = "slug"
  collection_id = chronosphere_collection.c.id
  dashboard_json = jsonencode({
    kind : "Dashboard",
    metadata : {
      name : "Chrono Dashboard"
    }
    spec : {
    }
  })
}
