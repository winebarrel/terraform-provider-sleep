# terraform-provider-sleep

[![CI](https://github.com/winebarrel/terraform-provider-sleep/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/terraform-provider-sleep/actions/workflows/ci.yml)
[![terraform docs](https://img.shields.io/badge/terraform-docs-%35835CC?logo=terraform)](https://registry.terraform.io/providers/winebarrel/sleep/latest/docs)

Terraform  provider of the function to pause.

## Usage

```tf
terraform {
  required_providers {
    sleep = {
      source  = "winebarrel/sleep"
    }
  }
}

output "with_sleep" {
  value = provider::sleep::with("hello", "3s")
  #=> value = "hello" (with 3s sleep)
}
```

## Run locally for development

```sh
cp sleep.tf.sample sleep.tf
make tf-plan
```
