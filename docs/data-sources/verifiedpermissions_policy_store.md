---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_verifiedpermissions_policy_store Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::VerifiedPermissions::PolicyStore
---

# awscc_verifiedpermissions_policy_store (Data Source)

Data Source schema for AWS::VerifiedPermissions::PolicyStore



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `policy_store_id` (String)
- `schema` (Attributes) (see [below for nested schema](#nestedatt--schema))
- `validation_settings` (Attributes) (see [below for nested schema](#nestedatt--validation_settings))

<a id="nestedatt--schema"></a>
### Nested Schema for `schema`

Read-Only:

- `cedar_json` (String)


<a id="nestedatt--validation_settings"></a>
### Nested Schema for `validation_settings`

Read-Only:

- `mode` (String)