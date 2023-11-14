---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dbtcloud_azure_dev_ops_repository Data Source - dbtcloud"
subcategory: ""
description: |-
  Use this data source to retrieve the ID and details of an Azure Dev Ops repository
  based on its name and the ID of the Azure Dev Ops project it belongs to.
  This data source requires connecting with a user token and doesn't work with a service token.
---

# dbtcloud_azure_dev_ops_repository (Data Source)

Use this data source to retrieve the ID and details of an Azure Dev Ops repository 
based on its name and the ID of the Azure Dev Ops project it belongs to.
		
This data source requires connecting with a user token and doesn't work with a service token.

## Example Usage

```terraform
data "dbtcloud_azure_dev_ops_repository" "my_ado_repository" {
  name = "my-repo-name"
  azure_dev_ops_project_id = data.dbtcloud_azure_dev_ops_project.my_ado_project.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `azure_dev_ops_project_id` (String) The internal Azure Dev Ops ID of the ADO Project. Can be retrieved using the data source dbtcloud_azure_dev_ops_project and the project name
- `name` (String) The name of the ADO repository

### Read-Only

- `default_branch` (String) The default branch of the ADO repository
- `details_url` (String) The URL of the ADO repository showing details about the repository and its attributes
- `id` (String) The internal Azure Dev Ops ID of the ADO Repository
- `remote_url` (String) The HTTP URL of the ADO repository used to connect to dbt Cloud
- `web_url` (String) The URL of the ADO repository accessible in the browser

