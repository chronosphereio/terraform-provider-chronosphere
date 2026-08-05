resource "chronosphere_collection" "c" {
  name        = "Collection"
  description = "collection created by terraform examples."
}

resource "chronosphere_notebook" "my_notebook" {
  name          = "Chrono Notebook"
  slug          = "slug"
  collection_id = chronosphere_collection.c.id

  notebook_json = jsonencode({
    version = 1
    cells = [
      {
        kind = "markdown"
        text = "## Investigation notes"
      },
    ]
  })
}
