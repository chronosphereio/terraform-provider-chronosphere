resource "chronosphere_derived_metric" "my-derived-metric" {
  name        = "my-derived-metric"
  slug        = "my-metric"
  description = "this is my derived metric"
  metric_name = "my-name"

  // There are two underlying queries for this derived metric.
  // When querying my-name{label1="value1"} the underlying "query1_expr" will be executed.
  // In all other cases, the "query_default" will be executed, without any variable support.
  queries {
    selector {
      labels = {
        label1 = "value1"
        label2 = "value2"
      }
    }
    query {
      expr = "query1_expr"
      variables {
        name             = "service"
        default_selector = "service=default"
      }
      variables {
        name             = "app"
        default_selector = "app!=production"
      }
    }
  }

  # default, no selector
  queries {
    query {
      expr = "query_default"
    }
  }
}
