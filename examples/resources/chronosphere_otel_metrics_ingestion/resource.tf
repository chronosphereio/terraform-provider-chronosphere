resource "chronosphere_otel_metrics_ingestion" "default_config" {
  resource_attributes {
    flatten_mode               = "MERGE"
    filter_mode                = "APPEND_DEFAULT_EXCLUDE_KEYS"
    exclude_keys               = ["telemetry.sdk.version", "telemetry.sdk.language"]
    generate_target_info       = true
  }
}
