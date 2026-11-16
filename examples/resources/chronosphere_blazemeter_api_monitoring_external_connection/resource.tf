variable "blazemeter_api_monitoring_access_token" {
  type      = string
  sensitive = true
}

resource "chronosphere_blazemeter_api_monitoring_external_connection" "blazemeter_api_monitoring" {
  name         = "BlazeMeter API Monitoring"
  access_token = var.blazemeter_api_monitoring_access_token
}
