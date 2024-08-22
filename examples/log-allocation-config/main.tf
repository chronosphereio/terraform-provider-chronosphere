resource "chronosphere_log_allocation_config" "my_log_allocation_config" {
  default_dataset {
    allocation {
      percent_of_license = 90.1
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
    dataset_id = chronosphere_dataset.example_logs_dataset.id
    allocation {
      percent_of_license = 9.9
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
}
