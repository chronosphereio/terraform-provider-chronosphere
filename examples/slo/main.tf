resource "chronosphere_blackhole_alert_notifier" "blackhole" {
  name = "Blackhole"
}

resource "chronosphere_notification_policy" "np" {
  team_id = chronosphere_team.t.id
  name    = "SLO team NP"

  route {
    severity  = "warn"
    notifiers = [chronosphere_blackhole_alert_notifier.blackhole.id]
  }
}

resource "chronosphere_team" "t" {
  name        = "SLO Team"
  description = "SLO Team"
}

resource "chronosphere_collection" "c" {
  name        = "SLO Collection"
  description = "Collection to put an SLO in"
  team_id     = chronosphere_team.t.id
}

resource "chronosphere_slo" "slo" {
  name                   = "SLO GO"
  collection_id          = chronosphere_collection.c.id
  notification_policy_id = chronosphere_notification_policy.np.id

  definition {
    objective = 99.95
    reporting_windows {
      duration = "28d"
    }
  }

  sli {
    custom_indicator {
      bad_query_template   = "sum(rate(http_request_duration_seconds_count{error=\"true\"}[{{ .Window }}]))"
      total_query_template = "sum(rate(http_request_duration_seconds_count[{{ .Window }}]))"
    }
  }
}
