resource "chronosphere_resource_pools_config" "default" {
  default_pool {
    allocation {
      percent_of_license = 50.1
    }
  }

  # NB: deprecated but equivalent "pools" is also supported
  pool {
    name = "first"
    allocation {
      percent_of_license = 49.9
    }

    # NB: deprecated match_rule is also supported, e.g.
    # match_rule = "foo:bar"
    match_rules = ["foo:bar", "baz:blah"]
  }
}
