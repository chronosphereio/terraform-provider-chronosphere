---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "chronosphere_opsgenie_alert_notifier Resource - chronosphere"
subcategory: ""
description: |-
  
---

# chronosphere_opsgenie_alert_notifier (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `api_key` (String, Sensitive)
- `name` (String)

### Optional

- `api_url` (String)
- `basic_auth_password` (String, Sensitive)
- `basic_auth_username` (String)
- `bearer_token` (String)
- `description` (String)
- `details` (Map of String)
- `message` (String)
- `note` (String)
- `priority` (String)
- `proxy_url` (String, Deprecated)
- `responder` (Block List) (see [below for nested schema](#nestedblock--responder))
- `send_resolved` (Boolean)
- `slug` (String)
- `source` (String)
- `tags` (Set of String)
- `tls_insecure_skip_verify` (Boolean)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--responder"></a>
### Nested Schema for `responder`

Required:

- `type` (String)

Optional:

- `id` (String)
- `name` (String)
- `username` (String)