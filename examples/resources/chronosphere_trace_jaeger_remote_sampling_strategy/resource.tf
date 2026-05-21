resource "chronosphere_trace_jaeger_remote_sampling_strategy" "probabilistic" {
  name         = "Probabilistic sampling for service A"
  service_name = "service-a"

  applied_strategy {
    probabilistic_strategy {
      sampling_rate = 0.01
    }
  }
}

resource "chronosphere_trace_jaeger_remote_sampling_strategy" "rate_limited" {
  name         = "Rate-limited sampling for service B"
  service_name = "service-b"

  applied_strategy {
    rate_limiting_strategy {
      max_traces_per_second = 2
    }
  }
}
