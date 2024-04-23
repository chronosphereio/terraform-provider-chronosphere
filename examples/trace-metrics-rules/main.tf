resource "chronosphere_trace_metrics_rule" "tmr" {
  name = "${var.prefix} tmr-name"
  slug = "${var.prefix}-tmr-slug"

  histogram_buckets_seconds = [
    1,
    2,
    3,
  ]
  metric_name = "${var.prefix}_tmr_metric_name"

  metric_labels = {
    "tmr_label_1" = "tmr_value_1"
    "tmr_label_2" = "tmr_value_2"
  }
  trace_filter {
    trace {
      duration {
        min_secs = 10
      }
      error {
        value = false
      }
    }
    span {
      match_type = "include"
      service {
        match = "exact"
        value = "svc1"
      }
      operation {
        value = "op1"
      }
      parent_service {
        value = "svc0"
      }
      parent_operation {
        value = "op0"
      }
      duration {
        min_secs = 11
      }
      error {
        value = true
      }
      tag {
        key = "tag1k"
        value {
          value = "tag1v"
        }
      }
      tag {
        key = "tag2k"
        value {
          value = "tag2v"
        }
      }
      tag {
        key = "tag3k"

        numeric_value {
          comparison = "GREATER_THAN"
          value      = 5.3
        }
      }
      span_count {
        min = 2
      }
    }
    span {
      match_type = "exclude"
      service {
        value = "svc2"
      }
      operation {
        value = "op2"
      }
      parent_service {
        value = "svc1"
      }
      parent_operation {
        value = "op1"
      }
      duration {
        max_secs = 2
      }
      tag {
        key = "tag3k"
        value {
          value = "tag3v"
        }
      }
      tag {
        key = "tag4k"
        value {
          value = "tag4v"
        }
      }
      span_count {
        max = 5
      }
    }
  }
}

resource "chronosphere_trace_metrics_rule" "tmr_with_group_by" {
  name = "${var.prefix} tmr-with-group-by-name"
  slug = "${var.prefix}-tmr-with-group-by-slug"

  histogram_buckets_seconds = [
    0.1,
    0.5,
    1.0,
    2.0,
    10.0,
  ]
  metric_name = "${var.prefix}_tmr_with_group_by_metric_name"

  metric_labels = {
    "tmr_wgb_label_1" = "tmr_wbg_value_1"
    "tmr_wbg_label_2" = "tmr_wbg_value_2"
  }
  group_by {
    label = "tmr_wgb_group_by_label_1"
    key {
      type      = "TAG"
      named_key = "foo"
    }
  }
  group_by {
    label = "tmr_wgb_group_by_label_2"
    key {
      type      = "OPERATION"
      named_key = "bar"
    }
  }

  trace_filter {
    trace {
      duration {
        min_secs = 10
      }
      error {
        value = false
      }
    }
    span {
      match_type = "include"
      service {
        match = "exact"
        value = "svc1"
      }
    }
  }
}
