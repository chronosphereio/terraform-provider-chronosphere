data "chronosphere_bucket" "default" {
  slug = "default"
}

# Look up by display name instead of slug.
data "chronosphere_bucket" "by_name" {
  name = "Default"
}
