resource "chronosphere_team" "t" {
  name = "${var.prefix} team"
}

resource "chronosphere_collection" "c" {
  name        = "${var.prefix} Bucket"
  team_id     = chronosphere_team.t.id
  description = "${var.prefix} collection created by terraform examples"
}
