resource "chronosphere_gcp_metrics_integration" "production" {
  name = "GCP Production Metrics"
  slug = "gcp-production"
  service_account {
    client_email = "chronosphere-metrics@my-project.iam.gserviceaccount.com"
  }
}
