resource "chronosphere_service_account" "unrestricted" {
  name         = "ci-deployer"
  unrestricted = true
}

resource "chronosphere_service_account" "restricted_read_only" {
  name = "metrics-reader"
  restriction {
    permission = "READ_ONLY"
    labels     = { "team" = "platform" }
  }
}
