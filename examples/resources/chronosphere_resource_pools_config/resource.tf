resource "chronosphere_resource_pools_config" "example" {
  default_pool {
    allocation {
      percent_of_license = 70
    }
    priorities {
      high_priority_match_rules {
        name   = "Critical Services"
        filter = "service:~(api|checkout|payment)"
      }
      low_priority_match_rules {
        name   = "Development"
        filter = "env:~(dev|test)"
      }
    }
  }

  pool {
    name = "Team A Pool"
    allocation {
      percent_of_license = 20
    }
    match_rules {
      name   = "Team A Services"
      filter = "team:a"
    }
    priorities {
      high_priority_match_rules {
        name   = "Team A Production"
        filter = "env:prod"
      }
    }
  }

  pool {
    name = "Team B Pool"
    allocation {
      percent_of_license = 10
    }
    match_rules {
      name   = "Team B Services"
      filter = "team:b"
    }
  }
}
