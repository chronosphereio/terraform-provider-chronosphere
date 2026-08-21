resource "chronosphere_cloud_integration" "gcp_prod" {
  name = "GCP Prod"
  slug = "gcp-prod"

  provider_type = "gcp"
  provider_config = jsonencode({
    service_account = {
      client_email = "scraper@my-project.iam.gserviceaccount.com"
    }
    metric_groups = [
      {
        project_id = "my-project"
        prefixes   = ["compute.googleapis.com/"]
      }
    ]
  })

  metric_labels = {
    env = "prod"
  }

  state = "ENABLED"
}
