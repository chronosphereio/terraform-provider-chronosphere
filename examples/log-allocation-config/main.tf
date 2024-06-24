resource "chronosphere_log_allocation_config" "my_log_allocation_config" {
  default_dataset {
    allocation {
      percent_of_license = 90.1
    }
  }

  dataset_allocation {
    dataset_slug = "example_logs_prod_payfe"
    allocation {
      percent_of_license = 9.9
    }
  }
}
