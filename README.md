# Lambda Terraform Provider

The Terraform provider for Lambda enables the declarative management of resources in [Lambda](https://lambdalabs.com/).

[![Build Status](https://github.com/squat/terraform-provider-lambda/workflows/CI/badge.svg)](https://github.com/squat/terraform-provider-lambda/actions?query=workflow%3ACI)

<!-- Start SDK Installation -->

<!-- End SDK Installation -->



## SDK Example Usage
<!-- Start SDK Example Usage -->

<!-- End SDK Example Usage -->



<!-- Start SDK Available Operations -->

<!-- End SDK Available Operations -->

<!-- Start Summary [summary] -->
## Summary

Lambda Provider: API for interacting with the Lambda GPU Cloud
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [Installation](#installation)
* [Available Resources and Data Sources](#available-resources-and-data-sources)
* [Testing the provider locally](#testing-the-provider-locally)
<!-- End Table of Contents [toc] -->

<!-- Start Installation [installation] -->
## Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    lambda = {
      source  = "squat/lambda"
      version = "0.1.2"
    }
  }
}

provider "lambda" {
  # Configuration options
}
```
<!-- End Installation [installation] -->

<!-- Start Available Resources and Data Sources [operations] -->
## Available Resources and Data Sources

### Resources

* [lambda_ssh_key](docs/resources/ssh_key.md)
### Data Sources

* [lambda_instance](docs/data-sources/instance.md)
* [lambda_instance_types](docs/data-sources/instance_types.md)
<!-- End Available Resources and Data Sources [operations] -->

<!-- Start Testing the provider locally [usage] -->
## Testing the provider locally

#### Local Provider

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```

#### Compiled Provider

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

1. Execute `go build` to construct a binary called `terraform-provider-lambda`
2. Ensure that the `.terraformrc` file is configured with a `dev_overrides` section such that your local copy of terraform can see the provider binary

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/squat/lambda" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```
<!-- End Testing the provider locally [usage] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "squat/lambda" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

Your `<PATH>` may vary depending on how your Go environment variables are configured. Execute `go env GOBIN` to set it, then set the `<PATH>` to the value returned. If nothing is returned, set it to the default location, `$HOME/go/bin`.

### Generation

This project is generated using [Speakeasy](https://github.com/speakeasy-api/speakeasy).
