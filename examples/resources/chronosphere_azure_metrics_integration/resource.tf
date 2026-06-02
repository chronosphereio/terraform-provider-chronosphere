resource "chronosphere_azure_metrics_integration" "subscription" {
  name = "Azure Metrics"
  slug = "azure-metrics"

  principal {
    tenant_id = "00000000-0000-0000-0000-000000000000"
    client_id = "00000000-0000-0000-0000-000000000000"
  }

  scrape_config {
    subscription_ids = ["00000000-0000-0000-0000-000000000000"]
    locations        = ["eastus", "westus"]

    resource_type {
      name = "Microsoft.Compute/virtualMachines"
    }

    resource_type {
      name         = "Microsoft.Storage/storageAccounts"
      metric_names = ["UsedCapacity"]
    }
  }

  count_metrics_enabled = true
  propagate_tags        = true
}
