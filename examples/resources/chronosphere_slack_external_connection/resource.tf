variable "slack_bot_token" {
  type      = string
  sensitive = true
}

resource "chronosphere_slack_external_connection" "slack" {
  name  = "Slack"
  token = var.slack_bot_token # xoxb-..., with the chat:write scope
}

resource "chronosphere_team" "platform" {
  name = "Platform"
}

resource "chronosphere_notification_policy" "platform" {
  name    = "Platform Policy"
  team_id = chronosphere_team.platform.id

  route {
    severity = "warn"
    destination {
      slack {
        external_connection_slug = chronosphere_slack_external_connection.slack.slug
        channels                 = ["#platform-alerts"]
      }
    }
  }
}
