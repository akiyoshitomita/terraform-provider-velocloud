/*
 * VMware SD-WAN
 *
 * data_source_
 */

package velocloud

import (
	"context"
	//"strconv"
	"fmt"
	//"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-velocloud/velocloud/vcoclient"
)

func dataSourceVeloProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVeloProfileRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "edge name",
			},
			"profile_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "edge id",
			},
			"profile_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "profile type",
			},
			"bastion_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Bastion state",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "description",
			},
			"logical_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "logical_id",
			},
		},
	}
}

func dataSourceVeloProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	conn := m.(*vcoclient.APIClient).EnterpriseApi

	post := &vcoclient.EnterpriseGetEnterpriseConfigurations{
		Name: d.Get("name").(string),
	}

	res, _, err := conn.EnterpriseGetEnterpriseConfigurations(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(res) == 0 {
		return diag.Errorf("Cloudn't search profile [%s]", d.Get("name"))
	}

	if len(res) != 1 {
		return diag.Errorf("Profile search error [%d]", len(res))
	}

	d.Set("profile_id", res[0].Id)
	d.Set("profile_type", res[0].ConfigurationType)
	d.Set("bastion_state", res[0].BastionState)
	d.Set("description", res[0].Description)
	d.Set("logical_id", res[0].LogicalId)

	d.SetId(fmt.Sprintf("%d", res[0].Id))
	return diags
}
