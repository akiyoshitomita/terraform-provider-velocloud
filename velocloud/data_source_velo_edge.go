/*
 * VMware SD-WAN
 *
 * data_source_
 */

package velocloud

import (
	"context"
	//"strconv"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-velocloud/velocloud/vcoclient"
)

func dataSourceVeloEdge() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVeloEdgeRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceVeloEdgeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	name := d.Get("name").(string)

	c := m.(*vcoclient.APIClient)

	post := &vcoclient.EnterpriseGetEnterpriseEdges{ 
		Id: 0, 
	}

	a, _, err := c.EnterpriseApi.EnterpriseGetEnterpriseEdges(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}
	//ctx context.Context, data EnterpriseGetEnterpriseEdges

	log.Println("-------------- DEBUG --------------------")
	log.Println(a)
	log.Println(name)
	d.Set("id", 1)

	d.SetId("1")
	return diags
}
