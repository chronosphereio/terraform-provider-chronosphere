resource "chronosphere_dashboard" "my_dashboard" {
  name = "Chrono Dashboard"
  slug          = "slug"
  collection_id = chronosphere_collection.c.id
  dashboard_json = jsonencode({
    kind : "Dashboard",
    spec : {
    }
  })
}
