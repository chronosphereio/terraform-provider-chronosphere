resource "chronosphere_bucket" "production" {
  name        = "Production Metrics"
  description = "Bucket for production environment metrics"
  labels = {
    environment = "production"
    team        = "platform"
  }
}
