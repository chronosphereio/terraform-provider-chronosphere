resource "chronosphere_trace_metrics_rule" "error_traces" {
  name        = "Error Trace Metrics"
  slug        = "error-trace-metrics"
  metric_name = "trace_errors_total"

  histogram_buckets_seconds = [
    0.1,
    0.5,
    1.0,
    5.0,
    10.0,
  ]

  metric_labels = {
    environment = "production"
    team        = "platform"
  }

  trace_filter {
    trace {
      duration {
        min_secs = 1
      }
      error {
        value = true
      }
    }
    span {
      match_type = "include"
      service {
        match = "exact"
        value = "api"
      }
      error {
        value = true
      }
      duration {
        min_secs = 0.5
      }
    }
  }
}

resource "chronosphere_trace_metrics_rule" "latency_by_service" {
  name        = "Service Latency Metrics"
  slug        = "service-latency-metrics"
  metric_name = "trace_duration_seconds"

  histogram_buckets_seconds = [
    0.01,
    0.05,
    0.1,
    0.5,
    1.0,
    2.0,
  ]

  group_by {
    label = "service_name"
    key {
      type = "SERVICE"
    }
  }

  group_by {
    label = "operation"
    key {
      type = "OPERATION"
    }
  }

  trace_filter {
    span {
      match_type = "include"
      service {
        match = "regex"
        value = "api.*"
      }
    }
  }
}
