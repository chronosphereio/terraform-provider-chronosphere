resource "chronosphere_service_account" "unrestricted" {
  name         = "- Unrestricted"
  unrestricted = true
}

resource "chronosphere_service_account" "restricted_readwrite" {
  name = "- Restricted, Labeled"
  restriction {
    permission = "READ_AND_WRITE"
  }
}

resource "chronosphere_service_account" "restricted_read_labeled" {
  name = "- Restricted, READ_ONLY, labeled"
  restriction {
    permission = "READ_ONLY"
    labels     = { "foo" : "bar" }
  }
}

resource "chronosphere_team" "my_team" {
  name        = "My Team With SA"
  description = "team with a service account"
  user_emails = [
    chronosphere_service_account.unrestricted.email,
  ]
}
