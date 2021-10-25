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

func dataSourceVeloAddressGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVeloAddressGroupRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "address group name",
			},
			"address_group_id": &schema.Schema{
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

func dataSourceVeloAddressGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	name := d.Get("name").(string)
	conn := m.(*vcoclient.APIClient).EnterpriseApi

	post := &vcoclient.EnterpriseGetObjectGroups{}
	post.Type_ = []string{"address_group"}

	res, _, err := conn.EnterpriseGetObjectGroups(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, v := range res {
		if v.Name == name {
			d.Set("adderss_group_id", v.Id)
			d.Set("logical_id", v.LogicalId)
			return diags
		}
	}

	return diag.Errorf("not found address group [%s]", name)
}
