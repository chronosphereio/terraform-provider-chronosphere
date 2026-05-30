resource "chronosphere_consumption_budget" "monthly_budget" {
  name        = "Monthly Consumption Budget"
  budget_type = "MONTHLY"
  amount      = 1000000

  alert_threshold_percent = 80

  notification_email_addresses = [
    "finance@example.com",
    "platform-team@example.com"
  ]
}
