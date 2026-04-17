resource "chronosphere_notification_policy" "team_alerts" {
  team_id = chronosphere_team.example.id
  name    = "Team Alert Policy"

  route {
    severity  = "critical"
    notifiers = [chronosphere_pagerduty_alert_notifier.oncall.id]
    group_by {
      label_names = ["service", "environment"]
    }
  }

  route {
    severity  = "warn"
    notifiers = [chronosphere_slack_alert_notifier.alerts_channel.id]
    group_by {
      label_names = ["service"]
    }
  }
}
