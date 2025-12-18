resource "chronosphere_drop_rule" "example" {
  name                    = "Drop High Cardinality Metrics"
  query                   = ["__name__:high_cardinality_metric"]
  activated_drop_duration = "70s"
  conditional_drop        = true
  rate_limit_threshold    = 50.0
}

resource "chronosphere_drop_rule" "drop_nan_values" {
  name                    = "Drop NaN Values"
  query                   = ["service=\"api\""]
  active                  = true
  activated_drop_duration = "70s"
  conditional_drop        = true
  drop_nan_value          = true
  rate_limit_threshold    = 50.0
}
