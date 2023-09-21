package data_sources

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dbt-labs/terraform-provider-dbtcloud/pkg/dbt_cloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var projectSchema = map[string]*schema.Schema{
	"project_id": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "ID of the project to represent",
	},
	"name": &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Optional:    true,
		Description: "Given name for project",
	},
	"connection_id": &schema.Schema{
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "ID of the connection associated with the project",
	},
	"repository_id": &schema.Schema{
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "ID of the repository associated with the project",
	},
	"freshness_job_id": &schema.Schema{
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "ID of Job for source freshness",
	},
	"docs_job_id": &schema.Schema{
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "ID of Job for the documentation",
	},
	"state": &schema.Schema{
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Project state should be 1 = active, as 2 = deleted",
	},
}

func DatasourceProject() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceProjectRead,
		Schema:      projectSchema,
	}
}

func datasourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dbt_cloud.Client)

	var diags diag.Diagnostics
	var project *dbt_cloud.Project

	if _, ok := d.GetOk("project_id"); ok {
		projectId := strconv.Itoa(d.Get("project_id").(int))

		if _, ok := d.GetOk("name"); ok {
			return diag.FromErr(fmt.Errorf("Both project_id and name were provided, only one is allowed"))
		}

		var err error
		project, err = c.GetProject(projectId)
		if err != nil {
			return diag.FromErr(err)
		}

	} else if _, ok := d.GetOk("name"); ok {
		projectName := d.Get("name").(string)

		var err error
		project, err = c.GetProjectByName(projectName)
		if err != nil {
			return diag.FromErr(err)
		}

	} else {
		return diag.FromErr(fmt.Errorf("Either project_id or name must be provided"))
	}

	if err := d.Set("project_id", project.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", project.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("connection_id", project.ConnectionID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repository_id", project.RepositoryID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("freshness_job_id", project.FreshnessJobId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("docs_job_id", project.DocsJobId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("state", project.State); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(*project.ID))

	return diags
}
