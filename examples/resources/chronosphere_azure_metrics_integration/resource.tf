resource "chronosphere_azure_metrics_integration" "production" {
  name          = "Azure Production Metrics"
  slug          = "azure-production"
  tenant_id     = "00000000-0000-0000-0000-000000000000"
  client_id     = "11111111-1111-1111-1111-111111111111"
  client_secret = var.azure_client_secret
}
