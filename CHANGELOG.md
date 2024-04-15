# Releases

## Unreleased

Added:
- Add data source for `chronosphere_collection`
- Add support for `drop_nan_value` in `chronosphere_drop_rule` resource
- Add `gcp_metrics_integration` resource.

Deprecated:
- Remove unsupported `bucket_id` in `chronosphere_notification_policy`
- Remove unsupported data source for `chronosphere_notification_policy`
- Remove in-provider validation of rollup rules in favor of server-side dry run validation
