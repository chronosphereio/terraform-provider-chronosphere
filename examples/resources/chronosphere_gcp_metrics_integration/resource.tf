resource "chronosphere_gcp_metrics_integration" "customer" {
  name = "GCP Metrics Integration"
  slug = "gcp-metrics-integration"
  service_account {
    client_email = "chronosphere-collector@my-project.iam.gserviceaccount.com"
  }
}
