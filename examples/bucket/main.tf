resource "chronosphere_bucket" "b" {
  name        = "${var.prefix} Bucket"
  description = "${var.prefix} bucket created by terraform examples"
  labels      = { "foo" : "bar" }
}
