# provider_type and the structure of provider_config are defined by the
# Chronosphere API for each supported cloud provider; consult the integration's
# documentation for its configuration reference.
resource "chronosphere_cloud_integration" "example" {
  name = "Example Integration"
  slug = "example-integration"

  provider_type = "example_provider"
  provider_config = jsonencode({
    # Provider-specific configuration.
  })

  external_connection_slug = "example-connection"

  metric_labels = {
    env = "prod"
  }

  state = "ENABLED"
}
