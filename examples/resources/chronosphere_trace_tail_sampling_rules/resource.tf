resource "chronosphere_trace_tail_sampling_rules" "rules" {
  default_sample_rate {
    enabled     = true
    sample_rate = 0.5
  }

  # Always keep error spans.
  rules {
    filter {
      span {
        match_type = "INCLUDE"
        error {
          value = true
        }
      }
    }
    sample_rate = 1.0
  }

  # Always keep slow spans.
  rules {
    filter {
      span {
        match_type = "INCLUDE"
        duration {
          min_secs = 1
        }
      }
    }
    sample_rate = 1.0
  }
}
