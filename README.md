# terraform-provider-vaultsecure

Maintained fork of [`defreng/terraform-provider-vaultsecure`](https://github.com/defreng/terraform-provider-vaultsecure).

This provider securely configures [AWS secret engines](https://developer.hashicorp.com/vault/docs/secrets/aws) in Vault without storing an AWS access key and its secret in Terraform state. The fork uses a current AWS SDK for Go v2 and supports both modern AWS IAM Identity Center profiles using the `sso_session` format and profiles created by `aws login` using `login_session`.

## Usage

```hcl
terraform {
  required_providers {
    vaultsecure = {
      source = "manojselukar/vaultsecure"
    }
  }
}

provider "vaultsecure" {
  aws_profile     = "my-sso-profile"
  vault_address   = "https://vault.example.com:8200"
  vault_namespace = "example"
}
```

Run `aws sso login --profile my-sso-profile` or `aws --profile <profile> login` before Terraform, depending on the profile type. If `aws_profile` is omitted, the standard AWS SDK credential chain is used, including `AWS_PROFILE`, environment credentials, web identity, and instance or task roles.

Resource names remain under the `vaultsecure_` prefix so existing state can migrate by replacing only the provider source address.

## Upstream attribution

This repository preserves the upstream history and license. Fork-specific changes are maintained independently under the `manojselukar/vaultsecure` provider address.

## Future work

Some ideas for future improvements or features 🤓 

* Improve tests
  * Measure test coverage
  * Add tests that cover import of the resource
  * Add tests that check behavior when the engine_path or iam username is changed
* add support also for other cloud secret backends (gcp, azure, alicloud, ...)
