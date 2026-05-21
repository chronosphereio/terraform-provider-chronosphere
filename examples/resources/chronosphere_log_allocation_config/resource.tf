resource "chronosphere_log_allocation_config" "config" {
  default_dataset {
    allocation {
      percent_of_license = 100
    }

    priorities {
      high_priority_filter {
        query = "severity = 'error'"
      }
      low_priority_filter {
        query = "severity = 'debug'"
      }
    }
  }
}
