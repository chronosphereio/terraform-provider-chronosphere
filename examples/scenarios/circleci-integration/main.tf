resource "chronosphere_circleci_external_connection" "production" {
  name      = "CircleCI production"
  slug      = "circleci-production"
  api_token = "XXXXX"
}

resource "chronosphere_circleci_external_connection" "generated_slug" {
  name      = "CircleCI staging"
  api_token = "XXXXX"
}

resource "chronosphere_cloud_integration" "production" {
  name                     = "CircleCI production"
  provider_type            = "circleci"
  state                    = "ENABLED"
  external_connection_slug = chronosphere_circleci_external_connection.production.slug
  provider_config = jsonencode({
    project_slugs       = ["gh/example-org/api"]
    flaky_tests_enabled = true
  })
}

resource "chronosphere_cloud_integration" "staging" {
  name                     = "CircleCI staging"
  provider_type            = "circleci"
  state                    = "ENABLED"
  external_connection_slug = chronosphere_circleci_external_connection.generated_slug.slug
  provider_config = jsonencode({
    project_slugs = ["gh/example-org/web"]
  })
}
