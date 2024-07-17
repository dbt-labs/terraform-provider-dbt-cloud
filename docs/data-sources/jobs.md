---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dbtcloud_jobs Data Source - dbtcloud"
subcategory: ""
description: |-
  Retrieve all the jobs for a given dbt Cloud project or environment along with the environment details for the jobs. This will return both the jobs created from Terraform but also the jobs created in the dbt Cloud UI.
---

# dbtcloud_jobs (Data Source)

Retrieve all the jobs for a given dbt Cloud project or environment along with the environment details for the jobs. This will return both the jobs created from Terraform but also the jobs created in the dbt Cloud UI.

## Example Usage

```terraform
// we can search all jobs by project
data dbtcloud_jobs test_all_jobs_in_project {
  project_id = 1234
}

// or by environment
data dbtcloud_jobs test_all_jobs_in_environment {
  environment_id = 1234
}

// we can then retrieve all the jobs from the environment flagged as production
// this would include the jobs created by Terraform and the jobs created from the dbt Cloud UI
locals {
  my_jobs_prod = [for job in data.dbtcloud_jobs.test_all_jobs_in_project.jobs : job if job.environment.deployment_type == "production"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `environment_id` (Number) The ID of the environment for which we want to retrieve the jobs (one of `project_id` or `environment_id` must be set)
- `project_id` (Number) The ID of the project for which we want to retrieve the jobs (one of `project_id` or `environment_id` must be set)

### Read-Only

- `jobs` (Attributes Set) Set of jobs with their details (see [below for nested schema](#nestedatt--jobs))

<a id="nestedatt--jobs"></a>
### Nested Schema for `jobs`

Read-Only:

- `dbt_version` (String) The version of dbt used for the job. If not set, the environment version will be used.
- `deferring_environment_id` (Number) The ID of the environment this job defers to
- `deferring_job_definition_id` (Number) [Deprecated - deferral is now set at the environment level] The ID of the job definition this job defers to
- `description` (String) The description of the job
- `environment` (Attributes) Details of the environment the job is running in (see [below for nested schema](#nestedatt--jobs--environment))
- `environment_id` (Number) The ID of environment
- `execute_steps` (List of String) The list of steps to run in the job
- `execution` (Attributes) (see [below for nested schema](#nestedatt--jobs--execution))
- `generate_docs` (Boolean) Whether the job generate docs
- `id` (Number) The ID of the job
- `job_completion_trigger_condition` (Attributes) Whether the job is triggered by the completion of another job (see [below for nested schema](#nestedatt--jobs--job_completion_trigger_condition))
- `job_type` (String) The type of job (e.g. CI, scheduled)
- `name` (String) The name of the job
- `project_id` (Number) The ID of the project
- `run_generate_sources` (Boolean) Whether the job test source freshness
- `schedule` (Attributes) (see [below for nested schema](#nestedatt--jobs--schedule))
- `settings` (Attributes) (see [below for nested schema](#nestedatt--jobs--settings))
- `triggers` (Attributes) (see [below for nested schema](#nestedatt--jobs--triggers))
- `triggers_on_draft_pr` (Boolean) Whether the CI job should be automatically triggered on draft PRs

<a id="nestedatt--jobs--environment"></a>
### Nested Schema for `jobs.environment`

Read-Only:

- `deployment_type` (String) Type of deployment environment: staging, production
- `id` (Number) ID of the environment
- `name` (String) Name of the environment
- `project_id` (Number)
- `type` (String) Environment type: development or deployment


<a id="nestedatt--jobs--execution"></a>
### Nested Schema for `jobs.execution`

Read-Only:

- `timeout_seconds` (Number) The number of seconds before the job times out


<a id="nestedatt--jobs--job_completion_trigger_condition"></a>
### Nested Schema for `jobs.job_completion_trigger_condition`

Read-Only:

- `condition` (Attributes) (see [below for nested schema](#nestedatt--jobs--job_completion_trigger_condition--condition))

<a id="nestedatt--jobs--job_completion_trigger_condition--condition"></a>
### Nested Schema for `jobs.job_completion_trigger_condition.condition`

Read-Only:

- `job_id` (Number)
- `project_id` (Number)
- `statuses` (Set of String)



<a id="nestedatt--jobs--schedule"></a>
### Nested Schema for `jobs.schedule`

Read-Only:

- `cron` (String) The cron schedule for the job. Only used if triggers.schedule is true


<a id="nestedatt--jobs--settings"></a>
### Nested Schema for `jobs.settings`

Read-Only:

- `target_name` (String) Value for `target.name` in the Jinja context
- `threads` (Number) Number of threads to run dbt with


<a id="nestedatt--jobs--triggers"></a>
### Nested Schema for `jobs.triggers`

Read-Only:

- `git_provider_webhook` (Boolean) Whether the job runs automatically on PR creation
- `github_webhook` (Boolean) Whether the job runs automatically on PR creation
- `on_merge` (Boolean) Whether the job runs automatically once a PR is merged
- `schedule` (Boolean) Whether the job runs on a schedule