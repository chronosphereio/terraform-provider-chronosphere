resource "chronosphere_recording_rule" "up_by_namespace" {
  name        = "up:by_namespace"
  metric_name = "up:by_namespace"
  expr        = "sum by (kubernetes_namespace) (up)"
  interval    = "60s"

  labels = {
    "owner" = "platform"
  }
}
