resource "chronosphere_trace_tail_sampling_rules" "my-tail-sampling-rules" {
  default_sample_rate {
    enabled     = true
    sample_rate = 0.5
  }


  rules {
    filter {
      span {
        match_type = "INCLUDE"

        tag {
          key = "tag1k"

          value {
            match = "EXACT"
            value = "tag1v"
          }
        }
        tag {
          key = "tag2k"

          value {
            match = "EXACT"
            value = "tag2v"
          }
        }

        tag {
          key = "tag3k"

          numeric_value {
            comparison = "GREATER_THAN"
            value = 5.3
          }
        }

        duration {
          max_secs = 16
          min_secs = 11
        }

        error {
          value = true
        }

        operation {
          match = "EXACT"
          value = "op1"
        }

        parent_operation {
          match = "EXACT"
          value = "op0"
        }

        parent_service {
          match = "EXACT"
          value = "svc0"
        }

        service {
          match = "EXACT"
          value = "svc1"
        }

        span_count {
          min = 2
          max = 4
        }
      }

      trace {
        duration {
          min_secs = 10
          max_secs = 15
        }

        error {
          value = false
        }
      }
    }

    sample_rate = 0.6
    name = "other name"
    system_name = "some_other_name"
  }

  rules {
    name = "Some Name"
    system_name = "some_system_name"
    filter {
      span {
        match_type = "INCLUDE"

        duration {
          max_secs = 16
          min_secs = 11
        }
      }
    }

    sample_rate = 0.4
  }
}
