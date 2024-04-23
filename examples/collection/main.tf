resource "chronosphere_collection" "infra" {
  name        = "${var.prefix} Infrastructure Collection"
  description = "Collection of resources related to infrastructure services."
  team_id     = chronosphere_team.t.id
}
