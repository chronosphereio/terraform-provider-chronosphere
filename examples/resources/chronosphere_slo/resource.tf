resource "chronosphere_slo" "api_availability" {
  name                   = "API Availability SLO"
  collection_id          = chronosphere_collection.example.id
  notification_policy_id = chronosphere_notification_policy.example.id

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
    custom_dimension_labels = ["service", "environment"]
  }
}

resource "chronosphere_slo" "api_latency" {
  name                   = "API Latency SLO"
  collection_id          = chronosphere_collection.example.id
  notification_policy_id = chronosphere_notification_policy.example.id

  signal_grouping {
    label_names = ["service"]
  }

  definition {
    objective = 95.0
    time_window {
      duration = "7d"
    }
    burn_rate_alerting_config {
      window   = "1h"
      budget   = 2
      severity = "critical"
    }
    burn_rate_alerting_config {
      window   = "6h"
      budget   = 5
      severity = "warn"
    }
    enable_burn_rate_alerting = true
  }

  sli {
    custom_timeslice_indicator {
      query_template = <<-EOT
        histogram_quantile(0.95, sum(rate(http_request_duration_seconds_bucket{
          {{ .AdditionalFilters }}
        }[{{ .TimeSlice }}])) by (le, {{ .GroupBy }}))
      EOT
      timeslice_size = "FIVE_MINUTES"
      condition {
        op    = "LEQ"
        value = 0.5
      }
    }
    custom_dimension_labels = ["service"]

    additional_promql_filters {
      name  = "env"
      type  = "MatchEqual"
      value = "prod"
    }
  }
}
