resource "chronosphere_notification_policy" "np" {
  team_id = chronosphere_team.t.id
  name    = "${var.prefix} team NP"

  route {
    severity  = "warn"
    notifiers = [chronosphere_email_alert_notifier.email.id]
  }
}

resource "chronosphere_collection" "infra" {
  name        = "${var.prefix} Infrastructure Collection"
  team_id     = chronosphere_team.t.id
  description = "Collection of resources related to infrastructure services."

  notification_policy_id = chronosphere_notification_policy.np.id
}
