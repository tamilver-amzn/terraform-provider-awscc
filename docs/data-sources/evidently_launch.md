---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_evidently_launch Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Evidently::Launch
---

# awscc_evidently_launch (Data Source)

Data Source schema for AWS::Evidently::Launch



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String)
- **description** (String)
- **groups** (Attributes List) (see [below for nested schema](#nestedatt--groups))
- **metric_monitors** (Attributes List) (see [below for nested schema](#nestedatt--metric_monitors))
- **name** (String)
- **project** (String)
- **randomization_salt** (String)
- **scheduled_splits_config** (Attributes List) (see [below for nested schema](#nestedatt--scheduled_splits_config))
- **tags** (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--groups"></a>
### Nested Schema for `groups`

Read-Only:

- **description** (String)
- **feature** (String)
- **group_name** (String)
- **variation** (String)


<a id="nestedatt--metric_monitors"></a>
### Nested Schema for `metric_monitors`

Read-Only:

- **entity_id_key** (String) The JSON path to reference the entity id in the event.
- **event_pattern** (String) Event patterns have the same structure as the events they match. Rules use event patterns to select events. An event pattern either matches an event or it doesn't.
- **metric_name** (String)
- **unit_label** (String)
- **value_key** (String) The JSON path to reference the numerical metric value in the event.


<a id="nestedatt--scheduled_splits_config"></a>
### Nested Schema for `scheduled_splits_config`

Read-Only:

- **group_weights** (Attributes Set) (see [below for nested schema](#nestedatt--scheduled_splits_config--group_weights))
- **start_time** (String)

<a id="nestedatt--scheduled_splits_config--group_weights"></a>
### Nested Schema for `scheduled_splits_config.group_weights`

Read-Only:

- **group_name** (String)
- **split_weight** (Number)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- **value** (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

