resource "chronosphere_trace_jaeger_remote_sampling_strategy" "tjrss_svc_a" {
  name         = "svc-a JRS strategy"
  service_name = "tjrss_svc_a"

  applied_strategy {
    probabilistic_strategy {
      sampling_rate = 0.01
    }
  }
}

resource "chronosphere_trace_jaeger_remote_sampling_strategy" "tjrss_svc_b" {
  name         = "svc-b JRS strategy"
  service_name = "tjrss_svc_b"

  applied_strategy {
    rate_limiting_strategy {
      max_traces_per_second = 2
    }
  }
}

resource "chronosphere_trace_jaeger_remote_sampling_strategy" "tjrss_svc_c" {
  name         = "svc-c JRS strategy"
  service_name = "tjrss_svc_c"

  applied_strategy {
    per_operation_strategies {
      default_sampling_rate                 = 0.01
      default_lower_bound_traces_per_second = 1
      default_upper_bound_traces_per_second = 1000
      per_operation_strategies {
        operation = "noisyop"
        probabilistic_strategy {
          sampling_rate = 0.0
        }
      }
      per_operation_strategies {
        operation = "interestinghighvolumeop"
        probabilistic_strategy {
          sampling_rate = 0.1
        }
      }
      per_operation_strategies {
        operation = "importantop"
        probabilistic_strategy {
          sampling_rate = 1.0
        }
      }
    }
  }
}
