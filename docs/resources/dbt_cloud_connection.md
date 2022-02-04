---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dbt_cloud_connection Resource - terraform-provider-dbt-cloud"
subcategory: ""
description: |-
  
---

# dbt_cloud_connection (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **account** (String) Account name for the connection
- **allow_keep_alive** (Boolean) Whether or not the connection should allow client session keep alive
- **allow_sso** (Boolean) Whether or not the connection should allow SSO
- **database** (String) Database name for the connection
- **name** (String) Connection name
- **project_id** (Number) Project ID to create the environment in
- **role** (String) Role name for the connection
- **type** (String) The type of environment (must be either development or deployment)
- **warehouse** (String) Warehouse name for the connection

### Optional

- **id** (String) The ID of this resource.
- **is_active** (Boolean) Whether the connection is active

