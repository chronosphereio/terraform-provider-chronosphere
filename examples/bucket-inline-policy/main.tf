resource "chronosphere_bucket" "b" {
  name                     = "bucket"
  notification_policy_data = chronosphere_notification_policy.np.notification_policy_data
}

resource "chronosphere_notification_policy" "np" {
  route {
    severity  = "warn"
    notifiers = [chronosphere_email_alert_notifier.email.id]
  }
}
