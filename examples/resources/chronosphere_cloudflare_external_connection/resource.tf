variable "cloudflare_api_token" {
  type      = string
  sensitive = true
}

resource "chronosphere_cloudflare_external_connection" "cloudflare" {
  name      = "Cloudflare"
  api_token = var.cloudflare_api_token
}
