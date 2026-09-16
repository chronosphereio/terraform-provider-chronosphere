resource "chronosphere_collection" "c" {
  name = "Checkout"
}

resource "chronosphere_slo" "checkout_availability" {
  name          = "checkout availability"
  collection_id = chronosphere_collection.c.id
  definition {
    objective = 99.9
    time_window {
      duration = "28d"
    }
  }
  sli {
    custom_indicator {
      bad_query_template   = "sum(rate(checkout_errors_total[{{ .Window }}]))"
      total_query_template = "sum(rate(checkout_requests_total[{{ .Window }}]))"
    }
  }
}

resource "chronosphere_slo" "checkout_latency" {
  name          = "checkout latency"
  collection_id = chronosphere_collection.c.id
  definition {
    objective = 99
    time_window {
      duration = "28d"
    }
  }
  sli {
    custom_indicator {
      bad_query_template   = "sum(rate(checkout_slow_requests_total[{{ .Window }}]))"
      total_query_template = "sum(rate(checkout_requests_total[{{ .Window }}]))"
    }
  }
}

resource "chronosphere_slo" "checkout_payment_success" {
  name          = "checkout payment success"
  collection_id = chronosphere_collection.c.id
  definition {
    objective = 99.95
    time_window {
      duration = "28d"
    }
  }
  sli {
    custom_indicator {
      bad_query_template   = "sum(rate(payment_errors_total[{{ .Window }}]))"
      total_query_template = "sum(rate(payment_requests_total[{{ .Window }}]))"
    }
  }
}

resource "chronosphere_command_center_group" "checkout" {
  name = "Checkout"
  primary_slo_reference {
    slug = chronosphere_slo.checkout_availability.slug
  }
  related_slo_references {
    slug = chronosphere_slo.checkout_latency.slug
  }
  related_slo_references {
    slug = chronosphere_slo.checkout_payment_success.slug
  }
}
