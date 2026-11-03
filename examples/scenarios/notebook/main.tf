resource "chronosphere_notebook" "my_notebook" {
  name          = "On-call runbook"
  slug          = "on-call-runbook"
  collection_id = chronosphere_collection.c.id
  labels = {
    "team" = "infra"
  }
  notebook_json = jsonencode({
    version = 1
    document = {
      type = "doc"
      content = [
        {
          type    = "heading"
          attrs   = { level = 1 }
          content = [{ type = "text", text = "On-call runbook" }]
        },
        {
          type    = "paragraph"
          content = [{ type = "text", text = "Start with the service dashboard, then check recent deploys." }]
        },
      ]
    }
  })
}

# A notebook in a user's personal collection, addressed as USER:<slug>.
resource "chronosphere_notebook" "personal_notes" {
  name          = "Personal notes"
  collection_id = "USER:alice"
  notebook_json = jsonencode({
    version  = 1
    document = { type = "doc", content = [] }
  })
}
