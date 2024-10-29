resource "chronosphere_dashboard" "my_dashboard" {
  name = "Chrono Dashboard"
  slug          = "slug"
  collection_id = chronosphere_collection.c.id
  labels = {
    "team" = "infra"
    "tracing-context" = ""
  }
  dashboard_json = jsonencode({
    kind: "Dashboard",
    spec: {
      events: [],
      panels: {},
      layouts: [],
      variables: [],
      duration: "30m"
    },
    spec_version: "1"
  })
}
