# Releases

## Unreleased

## v1.13.0

Added:
* Add `chronosphere_log_control_config` resource for configuring log control rules.
* Replace `drop_original` with `keep_original` in LogIngestConfig parser.
* Promote `chronosphere_consumption_config` resource from unstable to `/v1`.
* Promote `chronosphere_consumption_budget` resource from unstable to `/v1`.
* Add `span_tag` field to `chronosphere_derived_label` resource.

## v1.12.0

Added:
* Add unstable `chronosphere_consumption_config` resource for configuring telemetry data partitions.
* Add unstable `chronosphere_consumption_budget` resource for configuring telemetry data partitions.
* Update parsing support in `chronosphere_log_ingest_config` resource.

Deprecated:
* Deprecate `permissive` in `chronosphere_rollup_rule` resource.

## v1.11.0

Added:
* Add timeslice indicator support to `chronosphere_slo` resource with `custom_timeslice_indicator` field.

## v1.10.0

Added:
* Add enable burn rate alerting field to the unstable `chronosphere_slo` resource.
* Add time window to the unstable `chronosphere_slo` resource.
* Moved `chronosphere_slo` resource from stable to v1 API.
* Add `log_allocation_config` from model conversion
* Add `chronosphere_log_ingest_config` resource.

Removed:
* Remove query less SLO fields from unstable `chronosphere_slo` resource.
* Remove `reporting_windows` from unstable `chronosphere_slo` resource.

## 1.9.0

Removed:
* Remove deprecated low value field from unstable `chronosphere_slo` resource.
* Deprecated queryless additional_promql_filters in unstable `chronosphere_slo` resource.

Added:
* Add burn rate configs to the unstable `chronosphere_slo` resource.
* Add SLI level additional_promql_filters to the unstable `chronosphere_slo` resource.

## v1.8.0

Added:
* Add `priority_thresholds` to `chronosphere_resource_pools_config` pools.

## v1.7.0

Added:
* Adds support for logging query in the `chronosphere_monitor` resource.
* Adds support for custom dimension labels in the `chronosphere_slo` resource.
* Add v1 `chronosphere_logging_allocation_config` resource
* Add support for `is_root_span` field in common trace span filter type.
* Allow disabling dry-run via provider configuration `disable_dryrun`.


Fixed:
* Fix dry-run for `chronosphere_recording_rule` with `bucket_slug` and `execution_group`.

## 1.6.2

Added:
* Implemented support for signal grouping in the `chronosphere_slo` resource.

## v1.6.1

Added:
* Add unstable `chronosphere_slo` resource.

Fixed:
* Fix `chronosphere_resource_pools_config` error when default pool is not set.

## v1.6.0

Added:
* Adds support for `labels` in `chronosphere_dashboard`
* Allow `chronosphere_notification_policy` resources without a `team_id`.

Fixed:
* Support dry-run validation of monitor prometheus queries that contain
  dynamic expressions that aren't known at plan time.

## v1.5.1

Fixed:
 * Allow creating `chronosphere_slack_notifier` with actions without `action_confirm_text`.
 * Fix `chronosphere_rollup_rule` and `chronosphere_mapping_rule` diff after applying
   a rule without an interval or storage policy.

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
