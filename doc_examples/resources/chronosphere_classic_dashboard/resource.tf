resource "chronosphere_team" "t" {
  name = "team"
}

resource "chronosphere_collection" "c" {
  name        = "Bucket"
  team_id     = chronosphere_team.t.id
  description = "collection created by terraform examples"
}

resource "chronosphere_classic_dashboard" "my_dashboard" {
  collection_id = chronosphere_collection.c.id
  dashboard_json = jsonencode({
    title : "Dashboard",
    panels : [{
      "gridPos" : {
        "h" : 12,
        "w" : 24,
        "x" : 0,
        "y" : 0
      },
      id : 2,
      targets : [
        {
          expr : "1",
        }
      ],
      title : "Panel Title",
      type : "graph",
    }],
  })
}
