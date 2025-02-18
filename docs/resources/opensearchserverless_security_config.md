---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_opensearchserverless_security_config Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Amazon OpenSearchServerless security config resource
---

# awscc_opensearchserverless_security_config (Resource)

Amazon OpenSearchServerless security config resource



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `description` (String) Security config description
- `name` (String) The friendly name of the security config
- `saml_options` (Attributes) Describes saml options in form of key value map (see [below for nested schema](#nestedatt--saml_options))
- `type` (String) Config type for security config

### Read-Only

- `id` (String) The identifier of the security config

<a id="nestedatt--saml_options"></a>
### Nested Schema for `saml_options`

Required:

- `metadata` (String) The XML saml provider metadata document that you want to use

Optional:

- `group_attribute` (String) Group attribute for this saml integration
- `session_timeout` (Number) Defines the session timeout in minutes
- `user_attribute` (String) Custom attribute for this saml integration

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_opensearchserverless_security_config.example <resource ID>
```
