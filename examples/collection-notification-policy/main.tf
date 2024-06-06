resource "chronosphere_notification_policy" "np" {
  team_id = chronosphere_team.t.id
  name    = "team NP"

  route {
    severity  = "warn"
    notifiers = [chronosphere_email_alert_notifier.email.id]
    group_by  = ["example-attribute"]
  }
}

resource "chronosphere_collection" "infra" {
  name        = "Infrastructure Collection"
  team_id     = chronosphere_team.t.id
  description = "Collection of resources related to infrastructure services."

  notification_policy_id = chronosphere_notification_policy.np.id
}
