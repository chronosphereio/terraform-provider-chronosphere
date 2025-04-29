resource "chronosphere_team" "t" {
  name        = "Team"
  description = "Optional Team Description"
}

resource "chronosphere_email_alert_notifier" "email" {
  name = "Email Blackhole"
  to   = "blackhole@chronosphere.io"
}

resource "chronosphere_notification_policy" "np" {
  team_id = chronosphere_team.t.id
  name    = "team NP"

  route {
    severity  = "warn"
    notifiers = [chronosphere_email_alert_notifier.email.id]
  }
}

resource "chronosphere_collection" "infra" {
  name        = "Infrastructure Collection"
  team_id     = chronosphere_team.t.id
  description = "Collection of resources related to infrastructure services."

  notification_policy_id = chronosphere_notification_policy.np.id
}

resource "chronosphere_monitor" "collection_monitor" {
  name          = "Monitor in Collection"
  collection_id = chronosphere_collection.infra.id

  query {
    prometheus_expr = <<-EOF
      sum by (kubernetes_namespace) (
        up{kubernetes_namespace="rc"}
      )
    EOF
  }

  signal_grouping {
    label_names = ["kubernetes_namespace"]
  }

  series_conditions {
    condition {
      severity = "warn"
      value    = 20.0
      op       = "GT"
    }
  }
}
