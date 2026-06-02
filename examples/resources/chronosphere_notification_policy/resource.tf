resource "chronosphere_team" "platform" {
  name = "Platform"
}

resource "chronosphere_email_alert_notifier" "email" {
  name = "Platform Email"
  to   = "platform@example.com"
}

resource "chronosphere_notification_policy" "platform" {
  name    = "Platform Policy"
  team_id = chronosphere_team.platform.id

  route {
    severity  = "warn"
    notifiers = [chronosphere_email_alert_notifier.email.id]
    group_by {
      label_names = ["service"]
    }
  }
}
