---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_m2_application Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::M2::Application
---

# awscc_m2_application (Data Source)

Data Source schema for AWS::M2::Application



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `application_arn` (String)
- `application_id` (String)
- `definition` (Attributes) (see [below for nested schema](#nestedatt--definition))
- `description` (String)
- `engine_type` (String)
- `kms_key_id` (String) The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting application-related resources.
- `name` (String)
- `tags` (Map of String)

<a id="nestedatt--definition"></a>
### Nested Schema for `definition`

Read-Only:

- `content` (String)
- `s3_location` (String)


