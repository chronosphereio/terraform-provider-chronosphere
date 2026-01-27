provider "chronosphere" {
  api_token = var.chronosphere_api_token
  org       = var.chronosphere_org

  # Optional: Override API endpoint
  # api_url = "https://api.chronosphere.io"
}

variable "chronosphere_api_token" {
  description = "Chronosphere API token"
  type        = string
  sensitive   = true
}

variable "chronosphere_org" {
  description = "Chronosphere organization"
  type        = string
}
