resource "chronosphere_otel_metrics_ingestion" "my-otel-metrics-ingestion" {
  resource_attributes {
    flatten_mode = "MERGE"
    filter_mode  = "APPEND_DEFAULT_EXCLUDE_KEYS"
    exclude_keys = ["key1", "key2"]

    generate_target_info = false
  }
}
