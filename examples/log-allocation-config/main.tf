resource "chronosphere_log_allocation_config" "my_log_allocation_config" {
  default_dataset {
    allocation {
      percent_of_license = 90.1
    }
  }

  dataset_allocation {
    dataset_slug = chronosphere_dataset.example_logs_dataset.slug
    allocation {
      percent_of_license = 9.9
    }
  }
}
