# This assumes the service is created out-of-band (see create-service.sh)
data "chronosphere_service" "gateway" {
  slug = "gateway"
}

resource "chronosphere_monitor" "m" {
  name = "${var.prefix} Monitor in Gateway Service"
  collection_id = data.chronosphere_service.gateway.id
  notification_policy_id = chronosphere_notification_policy.np.id

  query {
    prometheus_expr = "up{foo=\"bar\"}"
  }

	series_conditions {
		condition {
			severity = "warn"
			value = 2.0
			op = "GT"
		}
	}
}


resource "chronosphere_dashboard" "dash_in_svc" {
  collection_id = data.chronosphere_service.gateway.id
  dashboard_json = jsonencode({
    kind : "Dashboard",
    metadata : {
      name : "${var.prefix} Native Dashboard in Gateway Service"
    }
    spec : {
    }
  })
}


resource "chronosphere_grafana_dashboard" "dash_in_svc" {
  collection_id = data.chronosphere_service.gateway.id
  dashboard_json = jsonencode({
    title : "${var.prefix} Dashboard In Gateway Service",
    panels : [{
      "gridPos" : {
        "h" : 12,
        "w" : 24,
        "x" : 0,
        "y" : 0
      },
      id : 2,
      targets : [
        {
          expr : "1",
        }
      ],
      title : "Panel Title",
      type : "graph",
    }],
  })
}
