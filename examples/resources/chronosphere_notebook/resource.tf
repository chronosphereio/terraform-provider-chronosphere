resource "chronosphere_collection" "c" {
  name = "Platform"
}

resource "chronosphere_notebook" "checkout_latency" {
  name          = "Checkout Latency Investigation"
  slug          = "checkout-latency-investigation"
  collection_id = chronosphere_collection.c.id

  notebook_json = jsonencode({
    version = 1
    cells = [
      {
        kind = "markdown"
        text = "## Checkout latency\nWhat changed at 14:00?"
      },
      {
        kind  = "query"
        query = "sum(rate(checkout_requests_total[5m]))"
      },
    ]
  })
}
