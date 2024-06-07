# Releases

## Unreleased

Bug fixes:
 * Fixed bug where `chronosphere_resource_pools_config` would not allow certain combinations
   of allocation percents due to float arithmetic not adding to exactly 100%.

Added:
 * Add unstable `otel_metrics_ingestion` resource.

Deprecated:
 * `proxy_url` is deprecated in alert receivers.
 * `chronosphere_mapping_rule.drop_timestamp` is deprecated.

## v1.0.0

Initial public release for terraform-provider-chronosphere.
