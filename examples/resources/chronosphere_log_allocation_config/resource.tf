resource "chronosphere_log_allocation_config" "example" {
  default_dataset {
    allocation {
      percent_of_license = 90.0
    }
    priorities {
      high_priority_filter {
        query = "severity='error'"
      }
      low_priority_filter {
        query = "severity='debug'"
      }
    }
  }

  dataset_allocation {
    dataset_id = chronosphere_dataset.custom_logs.id
    allocation {
      percent_of_license = 10.0
    }
    priorities {
      high_priority_filter {
        query = "severity='error'"
      }
      low_priority_filter {
        query = "severity='info'"
      }
    }
  }
}
