resource "chronosphere_team" "t" {
  name = "team"
}

resource "chronosphere_collection" "c" {
  name        = "Bucket"
  team_id     = chronosphere_team.t.id
  description = "collection created by terraform examples"
}
