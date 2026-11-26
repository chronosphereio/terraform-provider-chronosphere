resource "chronosphere_clickhouse_cloud_external_connection" "production" {
  name       = "ClickHouse Cloud production"
  slug       = "clickhouse-cloud-production"
  key_id     = "XXXXX"
  key_secret = "XXXXX"
}

resource "chronosphere_clickhouse_cloud_external_connection" "generated_slug" {
  name       = "ClickHouse Cloud staging"
  key_id     = "XXXXX"
  key_secret = "XXXXX"
}

resource "chronosphere_cloud_integration" "production" {
  name                     = "ClickHouse Cloud production"
  provider_type            = "clickhouse_cloud"
  state                    = "ENABLED"
  external_connection_slug = chronosphere_clickhouse_cloud_external_connection.production.slug
  provider_config = jsonencode({
    organization_id     = "0c6f2b8e-3d1a-4e7b-9f5c-8a2d4e6b1c73"
    include_service_ids = ["5b7e0a11-0000-4000-8000-0000000000aa"]
  })
}

resource "chronosphere_cloud_integration" "staging" {
  name                     = "ClickHouse Cloud staging"
  provider_type            = "clickhouse_cloud"
  state                    = "ENABLED"
  external_connection_slug = chronosphere_clickhouse_cloud_external_connection.generated_slug.slug
  provider_config = jsonencode({
    organization_id = "7d3a9f42-1b6e-4c8d-a5f0-2e9b8c4d6a17"
  })
}
