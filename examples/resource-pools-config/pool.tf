resource "chronosphere_resource_pools_config" "default" {
  default_pool {
    priorities {
      high_priority_match_rules = ["cluster:production*"]
      low_priority_match_rules  = ["cluster:test*"]
    }
  }

  # NB: deprecated but equivalent "pools" is also supported
  pool {
    name = "first"
    allocation {
      # Allocation specified as a % of the license.
      percent_of_license = 49.9
      # Allocation specified as a fixed value for a specific license.
      # fixed_values take precedence over percent_of_license.
      fixed_values = [
        {
          license = "PERSISTED_WRITES_STANDARD"
          value = 1000
        }
      ]
    }

    # NB: deprecated match_rule is also supported, e.g.
    # match_rule = "foo:bar"
    match_rules = ["foo:bar", "baz:blah"]

    priorities {
      high_priority_match_rules = ["cluster:production*"]
      low_priority_match_rules  = ["cluster:test*"]
    }
  }
}
