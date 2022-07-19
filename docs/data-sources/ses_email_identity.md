---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ses_email_identity Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::SES::EmailIdentity
---

# awscc_ses_email_identity (Data Source)

Data Source schema for AWS::SES::EmailIdentity



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `configuration_set_attributes` (Attributes) Used to associate a configuration set with an email identity. (see [below for nested schema](#nestedatt--configuration_set_attributes))
- `dkim_attributes` (Attributes) Used to enable or disable DKIM authentication for an email identity. (see [below for nested schema](#nestedatt--dkim_attributes))
- `dkim_dns_token_name_1` (String)
- `dkim_dns_token_name_2` (String)
- `dkim_dns_token_name_3` (String)
- `dkim_dns_token_value_1` (String)
- `dkim_dns_token_value_2` (String)
- `dkim_dns_token_value_3` (String)
- `dkim_signing_attributes` (Attributes) If your request includes this object, Amazon SES configures the identity to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be used for Easy DKIM. (see [below for nested schema](#nestedatt--dkim_signing_attributes))
- `email_identity` (String) The email address or domain to verify.
- `feedback_attributes` (Attributes) Used to enable or disable feedback forwarding for an identity. (see [below for nested schema](#nestedatt--feedback_attributes))
- `mail_from_attributes` (Attributes) Used to enable or disable the custom Mail-From domain configuration for an email identity. (see [below for nested schema](#nestedatt--mail_from_attributes))

<a id="nestedatt--configuration_set_attributes"></a>
### Nested Schema for `configuration_set_attributes`

Read-Only:

- `configuration_set_name` (String) The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.


<a id="nestedatt--dkim_attributes"></a>
### Nested Schema for `dkim_attributes`

Read-Only:

- `signing_enabled` (Boolean) Sets the DKIM signing configuration for the identity. When you set this value true, then the messages that are sent from the identity are signed using DKIM. If you set this value to false, your messages are sent without DKIM signing.


<a id="nestedatt--dkim_signing_attributes"></a>
### Nested Schema for `dkim_signing_attributes`

Read-Only:

- `domain_signing_private_key` (String) [Bring Your Own DKIM] A private key that's used to generate a DKIM signature. The private key must use 1024 or 2048-bit RSA encryption, and must be encoded using base64 encoding.
- `domain_signing_selector` (String) [Bring Your Own DKIM] A string that's used to identify a public key in the DNS configuration for a domain.
- `next_signing_key_length` (String) [Easy DKIM] The key length of the future DKIM key pair to be generated. This can be changed at most once per day.


<a id="nestedatt--feedback_attributes"></a>
### Nested Schema for `feedback_attributes`

Read-Only:

- `email_forwarding_enabled` (Boolean) If the value is true, you receive email notifications when bounce or complaint events occur


<a id="nestedatt--mail_from_attributes"></a>
### Nested Schema for `mail_from_attributes`

Read-Only:

- `behavior_on_mx_failure` (String) The action to take if the required MX record isn't found when you send an email. When you set this value to UseDefaultValue , the mail is sent using amazonses.com as the MAIL FROM domain. When you set this value to RejectMessage , the Amazon SES API v2 returns a MailFromDomainNotVerified error, and doesn't attempt to deliver the email.
- `mail_from_domain` (String) The custom MAIL FROM domain that you want the verified identity to use

