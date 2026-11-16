resource "chronosphere_blazemeter_api_monitoring_external_connection" "production" {
  name         = "BlazeMeter API Monitoring production"
  slug         = "blazemeter-api-monitoring-production"
  access_token = "XXXXX"
}

resource "chronosphere_blazemeter_api_monitoring_external_connection" "generated_slug" {
  name         = "BlazeMeter API Monitoring staging"
  access_token = "XXXXX"
}

resource "chronosphere_cloud_integration" "production" {
  name                     = "BlazeMeter API Monitoring production"
  provider_type            = "blazemeter_api_monitoring"
  state                    = "ENABLED"
  external_connection_slug = chronosphere_blazemeter_api_monitoring_external_connection.production.slug
  provider_config = jsonencode({
    include_bucket_keys = ["bk1example"]
  })
}

resource "chronosphere_cloud_integration" "staging" {
  name                     = "BlazeMeter API Monitoring staging"
  provider_type            = "blazemeter_api_monitoring"
  state                    = "ENABLED"
  external_connection_slug = chronosphere_blazemeter_api_monitoring_external_connection.generated_slug.slug
  provider_config          = jsonencode({})
}
