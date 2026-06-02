resource "chronosphere_opsgenie_external_connection" "opsgenie" {
  name    = "OpsGenie"
  api_key = "XXXXX"
  api_url = "https://api.opsgenie.com/"
}
