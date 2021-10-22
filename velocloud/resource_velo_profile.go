/*
 * VMware SD-WAN
 *
 * resource_
 */

package velocloud

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-velocloud/velocloud/vcoclient"
)

func resourceVeloProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVeloProfileCreate,
		ReadContext:   resourceVeloProfileRead,
		UpdateContext: resourceVeloProfileUpdate,
		DeleteContext: resourceVeloProfileDelete,
		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"profile_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Profile id",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "name of profile",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of profile",
			},
			"logical_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "logical id",
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
		},
	}
}

func resourceVeloProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conn := m.(*vcoclient.APIClient).ConfigurationApi

	post := &vcoclient.ConfigurationCloneEnterpriseTemplate{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	res, _, err := conn.ConfigurationCloneEnterpriseTemplate(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", res.Id))

	return resourceVeloProfileRead(ctx, d, m)
}

func resourceVeloProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	configId, _ := strconv.Atoi(d.Id())
	conn := m.(*vcoclient.APIClient).ConfigurationApi

	post := &vcoclient.ConfigurationGetConfiguration{
		Id: configId,
	}

	res, _, err := conn.ConfigurationGetConfiguration(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("profile_id", res.Id)
	d.Set("name", res.Name)
	d.Set("description", res.Description)
	d.Set("logical_id", res.LogicalId)
	d.Set("profile_type", res.ConfigurationType)
	d.Set("bastion_state", res.BastionState)

	return diags
}

func resourceVeloProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var isUpdate bool

	configId, _ := strconv.Atoi(d.Id())
	conn := m.(*vcoclient.APIClient).ConfigurationApi

	update := &vcoclient.ConfigurationUpdateConfigurationUpdate{}

	if d.HasChange("name") {
		isUpdate = true
		update.Name = d.Get("name").(string)
	}

	if d.HasChange("description") {
		isUpdate = true
		update.Description = d.Get("description").(string)
	}

	if isUpdate {

		post := &vcoclient.ConfigurationUpdateConfiguration{
			Id:     configId,
			Update: update,
		}

		_, _, err := conn.ConfigurationUpdateConfiguration(nil, *post)
		if err != nil {
			return diag.FromErr(err)
		}

	} else {
		log.Printf("[WARN] couldn't update VCO")
	}
	return resourceVeloProfileRead(ctx, d, m)

}

func resourceVeloProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	configId, _ := strconv.Atoi(d.Id())
	conn := m.(*vcoclient.APIClient).ConfigurationApi

	post := &vcoclient.ConfigurationDeleteConfiguration{
		Id: configId,
	}

	res, _, err := conn.ConfigurationDeleteConfiguration(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}

	if res.Rows != 1 {
		return diag.Errorf("ERROR couldn't delete profile")
	}
	return diags
}
