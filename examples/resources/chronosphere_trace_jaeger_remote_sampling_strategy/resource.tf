resource "chronosphere_trace_jaeger_remote_sampling_strategy" "api_service" {
  name         = "API Service Sampling Strategy"
  service_name = "api"

  applied_strategy {
    probabilistic_strategy {
      sampling_rate = 0.1
    }
  }
}

resource "chronosphere_trace_jaeger_remote_sampling_strategy" "rate_limited_service" {
  name         = "Rate Limited Service"
  service_name = "high_volume_service"

  applied_strategy {
    rate_limiting_strategy {
      max_traces_per_second = 100
    }
  }
}

resource "chronosphere_trace_jaeger_remote_sampling_strategy" "per_operation" {
  name         = "Per-Operation Sampling"
  service_name = "complex_service"

  applied_strategy {
    per_operation_strategies {
      default_sampling_rate                 = 0.01
      default_lower_bound_traces_per_second = 1
      default_upper_bound_traces_per_second = 1000

      per_operation_strategies {
        operation = "health_check"
        probabilistic_strategy {
          sampling_rate = 0.001
        }
      }

      per_operation_strategies {
        operation = "checkout"
        probabilistic_strategy {
          sampling_rate = 1.0
        }
      }

      per_operation_strategies {
        operation = "search"
        probabilistic_strategy {
          sampling_rate = 0.1
        }
      }
    }
  }
}
