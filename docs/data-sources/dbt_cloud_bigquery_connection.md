---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dbt_cloud_bigquery_connection Data Source - terraform-provider-dbt-cloud"
subcategory: ""
description: |-
  
---

# dbt_cloud_bigquery_connection (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connection_id` (Number) Connection Identifier
- `project_id` (Number) Project ID to create the connection in

### Read-Only

- `auth_provider_x509_cert_url` (String) Auth Provider X509 Cert URL for the Service Account
- `auth_uri` (String) Auth URI for the Service Account
- `client_email` (String) Service Account email
- `client_id` (String) Client ID of the Service Account
- `client_x509_cert_url` (String) Client X509 Cert URL for the Service Account
- `dataproc_cluster_name` (String) Dataproc cluster name for PySpark workloads
- `dataproc_region` (String) Google Cloud region for PySpark workloads on Dataproc
- `gcp_project_id` (String) GCP project ID
- `gcs_bucket` (String) URI for a Google Cloud Storage bucket to host Python code executed via Datapro
- `id` (String) The ID of this resource.
- `is_active` (Boolean) Whether the connection is active
- `location` (String) Location to create new Datasets in
- `maximum_bytes_billed` (Number) Max number of bytes that can be billed for a given BigQuery query
- `name` (String) Connection name
- `private_key` (String) Private key of the Service Account
- `private_key_id` (String) Private key ID of the Service Account
- `retries` (Number) Number of retries for queries
- `timeout_seconds` (Number) Timeout in seconds for queries
- `token_uri` (String) Token URI for the Service Account
- `type` (String) The type of connection

