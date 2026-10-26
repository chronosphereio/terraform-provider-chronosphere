resource "chronosphere_collection" "c" {
  name        = "Collection"
  description = "collection created by terraform examples."
  team_id     = chronosphere_team.t.id
}
