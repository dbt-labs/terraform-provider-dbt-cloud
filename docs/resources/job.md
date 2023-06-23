---
page_title: "dbtcloud_job Resource - dbtcloud"
subcategory: ""
description: |-
  
---

# dbtcloud_job (Resource)




## Example Usage

```terraform
// use dbt_cloud_job instead of dbtcloud_job for the legacy resource names
// legacy names will be removed from 0.3 onwards

resource "dbtcloud_job" "test" {
  environment_id = var.dbt_cloud_environment_id
  execute_steps = [
    "dbt test"
  ]
  generate_docs        = false
  is_active            = true
  name                 = "Test"
  num_threads          = 64
  project_id           = data.dbt_cloud_project.test_project.id
  run_generate_sources = false
  target_name          = "default"
  triggers = {
    "custom_branch_only" : true,
    "github_webhook" : false,
    "git_provider_webhook" : false,
    "schedule" : false
  }
  # this is the default that gets set up when modifying jobs in the UI
  schedule_days = [0, 1, 2, 3, 4, 5, 6]
  schedule_type = "days_of_week"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (Number) Environment ID to create the job in
- `execute_steps` (List of String) List of commands to execute for the job
- `name` (String) Job name
- `project_id` (Number) Project ID to create the job in
- `triggers` (Map of Boolean) Flags for which types of triggers to use, keys of github_webhook, git_provider_webhook, schedule, custom_branch_only

### Optional

- `dbt_version` (String) Version number of dbt to use in this job, usually in the format 1.2.0-latest rather than core versions
- `deferring_job_id` (Number) Job identifier that this job defers to
- `generate_docs` (Boolean) Flag for whether the job should generate documentation
- `is_active` (Boolean) Flag for whether the job is marked active or deleted
- `num_threads` (Number) Number of threads to use in the job
- `run_generate_sources` (Boolean) Flag for whether the job should run generate sources
- `schedule_cron` (String) Custom cron expression for schedule
- `schedule_days` (List of Number) List of days of week as numbers (0 = Sunday, 7 = Saturday) to execute the job at if running on a schedule
- `schedule_hours` (List of Number) List of hours to execute the job at if running on a schedule
- `schedule_interval` (Number) Number of hours between job executions if running on a schedule
- `schedule_type` (String) Type of schedule to use, one of every_day/ days_of_week/ custom_cron
- `self_deferring` (Boolean) Whether this job defers on a previous run of itself
- `target_name` (String) Target name for the dbt profile
- `timeout_seconds` (Number) Number of seconds to allow the job to run before timing out

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# Import using a job ID found in the URL or via the API.
terraform import dbtcloud_job.test_job "job_id"
terraform import dbtcloud_job.test_job 12345
```