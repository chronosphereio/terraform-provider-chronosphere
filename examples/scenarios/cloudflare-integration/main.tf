resource "chronosphere_cloudflare_external_connection" "production" {
  name      = "Cloudflare production"
  slug      = "cloudflare-production"
  api_token = "XXXXX"
}

resource "chronosphere_cloudflare_external_connection" "generated_slug" {
  name      = "Cloudflare staging"
  api_token = "XXXXX"
}

resource "chronosphere_cloud_integration" "production" {
  name                     = "Cloudflare production"
  provider_type            = "cloudflare"
  external_connection_slug = chronosphere_cloudflare_external_connection.production.slug
  provider_config = jsonencode({
    zone_names = ["example.com"]
  })
}

resource "chronosphere_cloud_integration" "staging" {
  name                     = "Cloudflare staging"
  provider_type            = "cloudflare"
  external_connection_slug = chronosphere_cloudflare_external_connection.generated_slug.slug
  provider_config          = jsonencode({})
}
