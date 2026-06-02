resource "chronosphere_collection" "c" {
  name = "Platform"
}

resource "chronosphere_dashboard" "platform" {
  name          = "Platform Overview"
  slug          = "platform-overview"
  collection_id = chronosphere_collection.c.id

  labels = {
    "team" = "platform"
  }

  dashboard_json = jsonencode({
    kind = "Dashboard"
    spec = {
      events    = []
      panels    = {}
      layouts   = []
      variables = []
      duration  = "30m"
    }
    spec_version = "1"
  })
}
