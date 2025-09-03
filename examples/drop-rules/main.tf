resource "chronosphere_drop_rule" "my-drop-rule" {
  name = "Drop Rule"
  query = [
    "ab=\"123\"",
  ]
  mode                    = "ENABLED"
  activated_drop_duration = "70s"
  conditional_drop        = true
  rate_limit_threshold    = 50.0
}

resource "chronosphere_drop_rule" "my-drop-rule-nan" {
  name = "Drop Rule (Drop NaN Values)"
  query = [
    "cd=\"321\"",
  ]
  mode                    = "ENABLED"
  activated_drop_duration = "70s"
  conditional_drop        = true
  drop_nan_value          = true
  rate_limit_threshold    = 50.0
}

resource "chronosphere_drop_rule" "my-drop-rule-preview" {
  name = "Drop Rule (Preview Mode)"
  query = [
    "ef=\"456\"",
  ]
  mode                    = "PREVIEW"
  activated_drop_duration = "60s"
  conditional_drop        = true
  rate_limit_threshold    = 30.0
}

