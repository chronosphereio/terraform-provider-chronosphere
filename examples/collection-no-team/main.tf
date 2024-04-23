resource "chronosphere_collection" "infra" {
  name        = "${var.prefix} Infrastructure Collection"
  description = "Collection of resources related to infrastructure services."
}
