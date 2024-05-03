# Releases

## Unreleased

Added:
- Add data source for `chronosphere_collection`
- Add support for `drop_nan_value` in `chronosphere_drop_rule` resource
- Add `gcp_metrics_integration` resource
- Add `name` field to `dashboard` resource.
- Add new v1 resource type, `dataset`, with the first telemetry type supported being tracing.

Updated:
- Renamed `chronosphere_grafana_dashboard` to `chronosphere_classic_dashboard`
- Add new `IN` and `NOT_IN` variations to tracing's `StringFilterMatchType` enum.

Fixed:
- Fix `chronosphere_rollup_rule` migration from `storage_policies` to `interval`.

Deprecated:
- Block unsupported use of duplicate routes with the same severity in `chronosphere_notification_policy`
- Remove unsupported `bucket_id` in `chronosphere_notification_policy`
- Remove unsupported data source for `chronosphere_notification_policy`
- Remove deprecated `rules` field from `chronosphere_notification_policy`
- Remove in-provider validation of rollup rules in favor of server-side dry run validation
- Consolidating tracing schemas: replace usage of `tags` with `tag`, `min_seconds` with `min_secs`, and `max_seconds` with `max_secs`.
- Deprecate `chronosphere_rollup_rule.storage_policies` and `chronsphere_mapping_rule.storage_policy` and recommend `interval`

Internal:
- Validate code generation in pull request CI
