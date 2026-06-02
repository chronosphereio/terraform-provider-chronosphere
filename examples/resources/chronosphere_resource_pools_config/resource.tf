resource "chronosphere_resource_pools_config" "default" {
  default_pool {
    priorities {
      high_priority_match_rules = ["cluster:production*"]
      low_priority_match_rules  = ["cluster:test*"]
    }
  }

  pool {
    name = "first"
    allocation {
      # Allocation specified as a fraction of the license; fixed_value takes
      # precedence when both are set.
      percent_of_license = 49.9

      fixed_value {
        license = "PERSISTED_WRITES_STANDARD"
        value   = 1000
      }
    }

    match_rules = ["team:platform", "env:production"]

    priorities {
      high_priority_match_rules = ["cluster:production*"]
      low_priority_match_rules  = ["cluster:test*"]
    }
  }
}
