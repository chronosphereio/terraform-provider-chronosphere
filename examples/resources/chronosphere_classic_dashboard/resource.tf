resource "chronosphere_collection" "c" {
  name = "Platform"
}

resource "chronosphere_classic_dashboard" "platform" {
  collection_id = chronosphere_collection.c.id

  dashboard_json = jsonencode({
    title = "Dashboard"
    panels = [
      {
        gridPos = {
          h = 12
          w = 24
          x = 0
          y = 0
        }
        id = 2
        targets = [
          {
            expr = "up"
          },
        ]
        title = "Up by instance"
        type  = "graph"
      },
    ]
  })
}
