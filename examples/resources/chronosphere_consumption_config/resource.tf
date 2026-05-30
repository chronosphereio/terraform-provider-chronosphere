resource "chronosphere_consumption_config" "example" {
  # Consumption config is typically a singleton resource
  # that configures consumption tracking for the organization

  trace_collection {
    enabled = true
  }

  metric_collection {
    enabled = true
  }
}
