resource "chronosphere_pagerduty_external_connection" "pagerduty" {
  name                     = "PagerDuty"
  pagerduty_api_key        = "XXXXX"
  pagerduty_events_version = "PAGERDUTY_EVENTS_VERSION_V2"
}
