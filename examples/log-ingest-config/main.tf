resource "chronosphere_log_ingest_config" "my-log-ingest-config" {
  parser {
    name = "syslog"
    regex = <<-EOT
      ^\<(?<pri>[0-9]+)\>(?<time>[^ ]* {1,2}[^ ]* [^ ]*) (?<ident>[a-zA-Z0-9_\/\.\-]*)(?:\[(?<pid>[0-9]+)\])?(?:[^\:]*\:)? *(?<message>.*)$
    EOT
  }

  parser {
    name = "apache_error"
    regex = <<-EOT
      ^\[[^ ]* (?<time>[^\]]*)\] \[(?<level>[^\]]*)\](?: \[pid (?<pid>[^\]]*)\])?( \[client (?<client>[^\]]*)\])? (?<message>.*)$
    EOT
  }
}
