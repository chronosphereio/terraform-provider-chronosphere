# Releases

## Unreleased

## v1.5.0

Added:
* Add `high_priority_filter` and `low_priority_filter` to unstable `chronosphere_log_allocation_config`.
* Add `fixed_value` in v1 `chronosphere_resource_pools_config.pools[].allocation`.

Fixed:
* Remove invalid fields from `chronosphere_dataset`, `chronosphere_trace_tail_sampling_rules`, `chronosphere_trace_metrics_rule` resources.

Deprecated:
* `allocation` in v1 `chronosphere_resource_pools_config.default_pool`, as this can be derived instead.

## v1.4.0

Added:
* Add v1 `chronosphere_logscale_alert` resource.
* Add v1 `chronosphere_logscale_action` resource.
* Add v1 `chronosphere_otel_metrics_ingestion` resource.

## v1.3.0

Added:
* Add unstable `chronosphere_logscale_alert` resource.
* Add unstable `chronosphere_logscale_action` resource.
* Add unstable `chronosphere_log_allocation_config` resource.
* Add support for `LOGS` type `dataset` resource.

## v1.2.0

Added:
 * Add server-side dry-run validation support for `chronosphere_trace_metrics_rule`.
 * Add `group_by` to `notification_routes` in `notification_policy` resource.

Fixed:
 * Reduce plan errors caused by dry-run concurrency.

## v1.1.0

Added:
 * Add unstable `chronosphere_otel_metrics_ingestion` resource.
 * Add server-side validation support for `chronosphere_trace_tail_sampling_rules`

Deprecated:
 * `proxy_url` is deprecated in alert receivers.
 * `chronosphere_mapping_rule.drop_timestamp` is deprecated.

Bug fixes:
* Fixed bug where `chronosphere_resource_pools_config` would not allow certain combinations
  of allocation percents due to float arithmetic not adding to exactly 100%. TF now relies
  on server-side validation of pools.

## v1.0.0

Initial public release for terraform-provider-chronosphere.
