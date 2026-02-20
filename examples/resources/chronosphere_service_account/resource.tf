resource "chronosphere_service_account" "api_ingestion" {
  name         = "API Ingestion Service Account"
  unrestricted = true
}

resource "chronosphere_service_account" "readonly_monitoring" {
  name = "Read-Only Monitoring"
  restriction {
    permission = "READ_ONLY"
    labels = {
      environment = "production"
    }
  }
}

resource "chronosphere_service_account" "team_service_account" {
  name = "Team Service Account"
  restriction {
    permission = "READ_AND_WRITE"
  }
}
