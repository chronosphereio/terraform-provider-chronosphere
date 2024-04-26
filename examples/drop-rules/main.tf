resource "chronosphere_drop_rule" "mah-drop-rule" {
  name = "Drop Rule"
  query = [
    "ab=\"123\"",
  ]
  active                  = true
  activated_drop_duration = "70s"
  conditional_drop        = true
  rate_limit_threshold    = 50.0
}

resource "chronosphere_drop_rule" "mah-drop-rule-nan" {
  name = "Drop Rule (Drop NaN Values)"
  query = [
    "cd=\"321\"",
  ]
  active                  = true
  activated_drop_duration = "70s"
  conditional_drop        = true
  drop_nan_value          = true
  rate_limit_threshold    = 50.0
}

