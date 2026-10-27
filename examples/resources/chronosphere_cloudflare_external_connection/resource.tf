variable "cloudflare_api_token" {
  type      = string
  sensitive = true
}

resource "chronosphere_cloudflare_external_connection" "cloudflare" {
  name                 = "Cloudflare"
  api_token_wo         = var.cloudflare_api_token
  api_token_wo_version = 1
}
