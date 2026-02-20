resource "chronosphere_victorops_alert_notifier" "oncall" {
  name        = "VictorOps On-Call"
  api_key     = var.victorops_api_key
  routing_key = var.victorops_routing_key
}
