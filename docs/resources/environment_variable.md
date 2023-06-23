---
page_title: "dbtcloud_environment_variable Resource - dbtcloud"
subcategory: ""
description: |-
  
---

# dbtcloud_environment_variable (Resource)

*Note*: Some upstream resources can be slow to create, so if creating a project or environment at
the same time as the environment variables, it's recommended to use the `depends_on` meta argument.

## Example Usage

```terraform
// use dbt_cloud_environment_variable instead of dbtcloud_environment_variable for the legacy resource names
// legacy names will be removed from 0.3 onwards

resource "dbtcloud_environment_variable" "my_env_var" {
  name       = "DBT_MY_ENV_VAR"
  project_id = dbtcloud_project.my_project.id
  environment_values = {
    "project" : "my_project_level_value",
    "My Env" : "my_env_level_value"
  }
  depends_on = [
    dbtcloud_project.my_project,
    dbtcloud_environment.my_env
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_values` (Map of String) Map from environment names to respective variable value, a special key `project` should be set for the project default variable value
- `name` (String) Name for the variable, must be unique within a project, must be prefixed with 'DBT_'
- `project_id` (Number) Project for the variable to be created in

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# Import using a project ID and environment variable name found in the URL and UI or via the API.
terraform import dbtcloud_environment_variable.test_environment_variable "project_id:environment_variable_name"
terraform import dbtcloud_environment_variable.test_environment_variable 12345:DBT_ENV_VAR
```