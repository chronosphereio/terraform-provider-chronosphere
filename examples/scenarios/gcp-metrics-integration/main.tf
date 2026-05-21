resource "chronosphere_gcp_metrics_integration" "customer" {
  name = "GCP Metrics Integration"
  slug = "test-gcp-metrics-integration"
  service_account {
        client_email = "test_cid"
  }
}
