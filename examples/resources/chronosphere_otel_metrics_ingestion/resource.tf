resource "chronosphere_otel_metrics_ingestion" "config" {
  resource_attributes {
    flatten_mode = "MERGE"
    filter_mode  = "APPEND_DEFAULT_EXCLUDE_KEYS"
    exclude_keys = ["host.id", "process.pid"]

    generate_target_info = false
  }
}
