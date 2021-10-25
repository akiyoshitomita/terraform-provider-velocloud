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

func dataSourceVeloSegment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVeloSegmentRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "segment name",
				//Default:     "Global Segment",
			},
			"segment_logical_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "segment logical id",
			},
			"segment_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "segment id",
			},
			"segment_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "segment type",
			},
		},
	}
}

func dataSourceVeloSegmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	name := d.Get("name").(string)
	conn := m.(*vcoclient.APIClient).EnterpriseApi

	post := &vcoclient.EnterpriseGetEnterpriseNetworkSegments{}
	segs, _, err := conn.EnterpriseGetEnterpriseNetworkSegments(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}
	for _, v := range segs {
		if name == v.Name {
			d.Set("segment_logical_id", v.LogicalId)
			d.Set("segment_id", v.Data.SegmentId)
			d.Set("segment_type", v.Type_)
			return diags
		}
	}

	return diag.Errorf("not found segment  [%s]", name)
}
