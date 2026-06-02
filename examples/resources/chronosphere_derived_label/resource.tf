resource "chronosphere_derived_label" "tier" {
  name        = "Tier from instance name"
  slug        = "tier-from-instance"
  description = "Derives a 'tier' label (read/write/admin) from the instance label"
  label_name  = "tier"

  metric_label {
    constructed_label {
      value_definitions {
        value = "read"
        filters {
          name       = "instance"
          value_glob = "reader-*"
        }
      }
      value_definitions {
        value = "write"
        filters {
          name       = "instance"
          value_glob = "writer-*"
        }
      }
      value_definitions {
        value = "admin"
        filters {
          name       = "instance"
          value_glob = "admin-*"
        }
      }
    }
  }
}
