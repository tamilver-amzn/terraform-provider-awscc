---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_oam_link Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Oam::Link Resource Type
---

# awscc_oam_link (Resource)

Definition of AWS::Oam::Link Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `resource_types` (Set of String)
- `sink_identifier` (String)

### Optional

- `label_template` (String)
- `tags` (Map of String) Tags to apply to the link

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.
- `label` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_oam_link.example <resource ID>
```
