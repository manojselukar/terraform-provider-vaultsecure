---
page_title: "Provider: VaultSecure"
subcategory: ""
description: |-
  Set Vault AWS secret backend root credentials without leaking them into the Terraform state
---

# VaultSecure Provider

This provider is used to securely setup AWS secret backends in Vault **without storing an AWS access key and its secret in the Terraform state**.

It does this by creating and managing an AWS access key, and directly storing the secret key in the Vault secret backend configuration. On top of that, after passing the secret key to Vault, it also uses the [key rotation feature in Vault](https://www.vaultproject.io/api-docs/secret/aws#rotate-root-iam-credentials) to ensure that the secret key is not known to any entity outside of Vault and AWS.

## Example Usage

```terraform
provider "vaultsecure" {
  // Optional. Supports IAM Identity Center sso_session profiles and
  // profiles created by aws login. AWS_PROFILE is used when omitted.
  aws_profile = "my-sso-profile"
  
  // The Vault token must be provided
  // in the VAULT_TOKEN environment variable
  vault_address = "https://myvaultserver.com:8200"
}
```

## Schema

### Optional

- **aws_profile** (String, Optional) AWS shared configuration profile. Modern IAM Identity Center profiles using `sso_session` and profiles created by `aws login` using `login_session` are supported. Defaults to `AWS_PROFILE` and then the AWS SDK default credential chain.
- **vault_address** (String, Optional) The URL of the Vault server (defaults to `https://127.0.0.1:8200`), can also be set via the `VAULT_ADDR` environment variable.
- **vault_namespace** (String, Optional) Vault namespace that should be used (defaults to `null`), can also be set via the `VAULT_NAMESPACE` environment variable.
