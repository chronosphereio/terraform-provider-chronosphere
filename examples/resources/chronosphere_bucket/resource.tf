resource "chronosphere_bucket" "b" {
  name        = "Bucket"
  description = "bucket created by terraform examples"
  labels      = { "foo" : "bar" }
}
