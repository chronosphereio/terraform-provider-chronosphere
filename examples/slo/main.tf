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
    custom_dimension_labels = ["label1", "label2"]
  }
}

resource "chronosphere_slo" "slo_with_signal_grouping_signal_per_series" {
  name                   = "SLO With Signal Grouping (Signal per series)"
  collection_id          = chronosphere_collection.c.id
  notification_policy_id = chronosphere_notification_policy.np.id

  signal_grouping {
    signal_per_series = true
  }

  definition {
    objective = 99.95
    reporting_windows {
      duration = "28d"
    }
    burn_rate_alerting_config {
      window = "1h"
      budget = 99
      severity = "critical"
      labels = {
        "foo": "bar"
      }
    }
    burn_rate_alerting_config {
      window = "6h"
      budget = 99
      severity = "critical"
      labels = {
        "foo": "baz"
      }
    }
    burn_rate_alerting_config {
      window = "24h"
      budget = 99
      severity = "warn"
    }
  }

  sli {
    custom_indicator {
      bad_query_template   = "sum(rate(http_request_duration_seconds_count{error=\"true\"}[{{ .Window }}]))"
      total_query_template = "sum(rate(http_request_duration_seconds_count[{{ .Window }}]))"
    }
  }
}

resource "chronosphere_slo" "slo_with_signal_grouping_labels" {
  name                   = "SLO With Signal Grouping (Labels"
  collection_id          = chronosphere_collection.c.id
  notification_policy_id = chronosphere_notification_policy.np.id

  signal_grouping {
    label_names = ["label1", "label2"]
  }

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

resource "chronosphere_slo" "slo_with_filters" {
  name                   = "SLO with filters"
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
      bad_query_template   = <<-EOT
        sum(rate(http_request_duration_seconds_count{
            error=\"true\",
            {{ .AdditionalFilters }}
        }[{{ .Window }}]))
      EOT
      total_query_template = <<-EOT
        sum(rate(http_request_duration_seconds_count{
              {{ .AdditionalFilters }}
        }[{{ .Window }}]))
      EOT
    }
    custom_dimension_labels = ["label1", "label2"]

    additional_promql_filters{
      name = "env"
      type = "MatchEqual"
      value = "prod"
    }

    additional_promql_filters{
        name = "namespace"
        type = "MatchRegex"
        value = "foo.*"
    }
  }
}