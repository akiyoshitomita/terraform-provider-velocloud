/*
 * VMware SD-WAN
 *
 * data_source_
 */

package velocloud

import (
	"context"
	//"regexp"
	//"strconv"
	//"fmt"
	//"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-velocloud/velocloud/vcoclient"
)

func dataSourceVeloPortGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVeloPortGroupRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "address group name",
			},
			"port_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "address group id",
			},
			"logical_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "logical_id",
			},
		},
	}
}

func dataSourceVeloPortGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	name := d.Get("name").(string)
	conn := m.(*vcoclient.APIClient).EnterpriseApi

	post := &vcoclient.EnterpriseGetObjectGroups{}
	post.Type_ = []string{"port_group"}

	res, _, err := conn.EnterpriseGetObjectGroups(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, v := range res{
		if v.Name == name {
			d.Set("port_group_id", v.Id)
			d.Set("logical_id", v.LogicalId)
			return diags
		}
	}

	return diag.Errorf("not found address group [%s]", name)
}
