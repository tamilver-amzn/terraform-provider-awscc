---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iotfleetwise_campaign Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::IoTFleetWise::Campaign
---

# awscc_iotfleetwise_campaign (Data Source)

Data Source schema for AWS::IoTFleetWise::Campaign



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `action` (String)
- `arn` (String)
- `collection_scheme` (Attributes) (see [below for nested schema](#nestedatt--collection_scheme))
- `compression` (String)
- `creation_time` (String)
- `data_extra_dimensions` (List of String)
- `description` (String)
- `diagnostics_mode` (String)
- `expiry_time` (String)
- `last_modification_time` (String)
- `name` (String)
- `post_trigger_collection_duration` (Number)
- `priority` (Number)
- `signal_catalog_arn` (String)
- `signals_to_collect` (Attributes List) (see [below for nested schema](#nestedatt--signals_to_collect))
- `spooling_mode` (String)
- `start_time` (String)
- `status` (String)
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))
- `target_arn` (String)

<a id="nestedatt--collection_scheme"></a>
### Nested Schema for `collection_scheme`

Read-Only:

- `condition_based_collection_scheme` (Attributes) (see [below for nested schema](#nestedatt--collection_scheme--condition_based_collection_scheme))
- `time_based_collection_scheme` (Attributes) (see [below for nested schema](#nestedatt--collection_scheme--time_based_collection_scheme))

<a id="nestedatt--collection_scheme--condition_based_collection_scheme"></a>
### Nested Schema for `collection_scheme.condition_based_collection_scheme`

Read-Only:

- `condition_language_version` (Number)
- `expression` (String)
- `minimum_trigger_interval_ms` (Number)
- `trigger_mode` (String)


<a id="nestedatt--collection_scheme--time_based_collection_scheme"></a>
### Nested Schema for `collection_scheme.time_based_collection_scheme`

Read-Only:

- `period_ms` (Number)



<a id="nestedatt--signals_to_collect"></a>
### Nested Schema for `signals_to_collect`

Read-Only:

- `max_sample_count` (Number)
- `minimum_sampling_interval_ms` (Number)
- `name` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)


