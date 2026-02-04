resource "chronosphere_recording_rule" "bucket_owned" {
  bucket_id   = chronosphere_bucket.b.id
  name        = "bucket owned rule"
  metric_name = "up:by_cluster:bucket_owned"
  expr        = "sum by (kubernetes_namespace) (up)"
  interval    = "60s"

  labels = {
    "owner" = "infra"
  }
}

resource "chronosphere_recording_rule" "minimal_no_bucket" {
  name            = "up:by_cluster:minimal"
  execution_group = "foo"
  expr            = "sum by (kubernetes_namespace) (up)"
}

resource "chronosphere_recording_rule" "synchronized" {
  name            = "up:by_cluster:synchronized"
  execution_group = "my-group"
  expr            = "sum by (namespace) (up)"
  execution_mode  = "SYNCHRONIZED"
}
