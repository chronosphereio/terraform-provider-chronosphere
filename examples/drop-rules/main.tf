resource "chronosphere_drop_rule" "my-drop-rule" {
  name = "Drop Rule"
  query = [
    "ab=\"123\"",
  ]
  activated_drop_duration = "70s"
  conditional_drop        = true
  rate_limit_threshold    = 50.0
}

resource "chronosphere_drop_rule" "my-drop-rule-active-false" {
  name = "Drop Rule"
  query = [
    "ab=\"123\"",
  ]
  active                  = false
  activated_drop_duration = "70s"
  conditional_drop        = true
  rate_limit_threshold    = 50.0
}

resource "chronosphere_drop_rule" "my-drop-rule-active-true" {
  name = "Drop Rule"
  query = [
    "ab=\"123\"",
  ]
  active                  = false
  activated_drop_duration = "70s"
  conditional_drop        = true
  rate_limit_threshold    = 50.0
}

resource "chronosphere_drop_rule" "my-drop-rule-mode-disabled" {
  name = "Drop Rule"
  query = [
    "ab=\"123\"",
  ]
  mode                    = "DISABLED"
  activated_drop_duration = "70s"
  conditional_drop        = true
  rate_limit_threshold    = 50.0
}

resource "chronosphere_drop_rule" "my-drop-rule-mode-enabled" {
  name = "Drop Rule"
  query = [
    "ab=\"123\"",
  ]
  mode                    = "ENABLED"
  activated_drop_duration = "70s"
  conditional_drop        = true
  rate_limit_threshold    = 50.0
}

resource "chronosphere_drop_rule" "my-drop-rule-mode-preview" {
  name = "Drop Rule"
  query = [
    "ab=\"123\"",
  ]
  mode                    = "PREVIEW"
  activated_drop_duration = "70s"
  conditional_drop        = true
  rate_limit_threshold    = 50.0
}

resource "chronosphere_drop_rule" "my-drop-rule-nan" {
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
