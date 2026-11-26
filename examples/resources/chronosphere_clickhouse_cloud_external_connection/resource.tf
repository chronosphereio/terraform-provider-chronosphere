variable "clickhouse_cloud_key_id" {
  type = string
}

variable "clickhouse_cloud_key_secret" {
  type      = string
  sensitive = true
}

resource "chronosphere_clickhouse_cloud_external_connection" "clickhouse_cloud" {
  name       = "ClickHouse Cloud"
  key_id     = var.clickhouse_cloud_key_id
  key_secret = var.clickhouse_cloud_key_secret
}
