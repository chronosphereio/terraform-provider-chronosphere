resource "chronosphere_team" "payments" {
  name = "Payments"
}

resource "chronosphere_blackhole_alert_notifier" "blackhole" {
  name = "Blackhole"
}

resource "chronosphere_notification_policy" "np" {
  name    = "SLO policy"
  team_id = chronosphere_team.payments.id

  route {
    severity  = "warn"
    notifiers = [chronosphere_blackhole_alert_notifier.blackhole.id]
  }
}

resource "chronosphere_collection" "c" {
  name    = "Payments"
  team_id = chronosphere_team.payments.id
}

resource "chronosphere_slo" "request_success" {
  name                   = "payments request success"
  collection_id          = chronosphere_collection.c.id
  notification_policy_id = chronosphere_notification_policy.np.id

  definition {
    objective = 99.95
    time_window {
      duration = "28d"
    }
    enable_burn_rate_alerting = true
  }

  sli {
    custom_indicator {
      bad_query_template   = "sum(rate(http_request_duration_seconds_count{error=\"true\"}[{{ .Window }}]))"
      total_query_template = "sum(rate(http_request_duration_seconds_count[{{ .Window }}]))"
    }
  }
}
