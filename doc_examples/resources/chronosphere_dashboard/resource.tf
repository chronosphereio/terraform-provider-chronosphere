resource "chronosphere_team" "t" {
  name        = "Team"
  description = "Optional Team Description"
}

resource "chronosphere_collection" "c" {
  name        = "Collection"
  description = "collection created by terraform examples."
  team_id     = chronosphere_team.t.id
}

resource "chronosphere_dashboard" "my_dashboard" {
  name          = "Chrono Dashboard"
  slug          = "slug"
  collection_id = chronosphere_collection.c.id
  dashboard_json = jsonencode({
    kind : "Dashboard",
    spec : {
    }
  })
}
