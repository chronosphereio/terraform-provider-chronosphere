variable "circleci_api_token" {
  type      = string
  sensitive = true
}

resource "chronosphere_circleci_external_connection" "circleci" {
  name      = "CircleCI"
  api_token = var.circleci_api_token
}
