resource "chronosphere_team" "t" {
  name = "${var.prefix} Team"
  description = "Optional ${var.prefix} Team Description"
}
