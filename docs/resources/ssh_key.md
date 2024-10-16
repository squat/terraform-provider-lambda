---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "lambda_ssh_key Resource - terraform-provider-lambda"
subcategory: ""
description: |-
  SSHKey Resource
---

# lambda_ssh_key (Resource)

SSHKey Resource

## Example Usage

```terraform
resource "lambda_ssh_key" "my_sshkey" {
  name       = "macbook-pro"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDfKpav4ILY54InZe27G user"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the SSH key. Requires replacement if changed.

### Optional

- `public_key` (String) Public key for the SSH key. Requires replacement if changed.

### Read-Only

- `id` (String) Unique identifier (ID) of an SSH key
- `private_key` (String) Private key for the SSH key. Only returned when generating a new key pair.
