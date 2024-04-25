# Releases

## Unreleased

Added:
- Add data source for `chronosphere_collection`
- Add support for `drop_nan_value` in `chronosphere_drop_rule` resource
- Add `gcp_metrics_integration` resource
- Add new resource type, `dataset`, with the first telemetry type supported being tracing.
- Add new `IN` and `NOT_IN` variations to tracing's `StringFilterMatchType` enum.

Deprecated:
- Remove unsupported `bucket_id` in `chronosphere_notification_policy`
- Remove unsupported data source for `chronosphere_notification_policy`
- Remove in-provider validation of rollup rules in favor of server-side dry run validation
- Consolidating tracing schemas: replace usage of `tags` with `tag`, `min_seconds` with `min_secs`, and `max_seconds` with `max_secs`.

Internal:
- Validate code generation in pull request CI