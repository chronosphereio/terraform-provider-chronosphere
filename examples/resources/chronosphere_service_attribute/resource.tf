resource "chronosphere_service_attribute" "environment" {
  name        = "Environment"
  slug        = "environment"
  description = "Service environment (production, staging, development)"

  values = [
    "production",
    "staging",
    "development"
  ]
}

resource "chronosphere_service_attribute" "tier" {
  name        = "Service Tier"
  slug        = "tier"
  description = "Service tier classification"

  values = [
    "critical",
    "high",
    "medium",
    "low"
  ]
}
