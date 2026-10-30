resource "chronosphere_collection" "c" {
  name = "Platform"
}

resource "chronosphere_notebook" "oncall_runbook" {
  name          = "On-call runbook"
  slug          = "on-call-runbook"
  collection_id = chronosphere_collection.c.id

  labels = {
    "team" = "platform"
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
