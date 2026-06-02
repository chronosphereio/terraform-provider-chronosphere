resource "chronosphere_victorops_external_connection" "victorops" {
  name    = "VictorOps"
  api_key = "00000000-0000-0000-0000-000000000000"
  api_url = "https://alert.victorops.com/integrations/generic/00000000/alert/"
}
