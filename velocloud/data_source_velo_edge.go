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
			"activation_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"build_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_info": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_family": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"edge_state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"endpoint_pki_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterprise_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"factory_software_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"factory_build_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ha_previous_state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ha_serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ha_state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"model_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"self_mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"site_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"software_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_hub": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_software_version_supported_by_vco": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceVeloEdgeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var edge *vcoclient.EnterpriseGetEnterpriseEdgesResultItem
	edge = nil
	name := d.Get("name").(string)

	c := m.(*vcoclient.APIClient)

	post := &vcoclient.EnterpriseGetEnterpriseEdges{
		Id: 0,
	}

	a, _, err := c.EnterpriseApi.EnterpriseGetEnterpriseEdges(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, v := range a {
		if v.Name == name {
			edge = &v
			break
		}
	}
	if edge == nil {
		return diag.Errorf("Not Found Edge[" + name + "]")
	}

	d.Set("id", edge.Id)
	d.Set("activation_key", edge.ActivationKey)
	d.Set("build_number", edge.BuildNumber)
	d.Set("custom_info", edge.CustomInfo)
	d.Set("description", edge.Description)
	d.Set("device_family", edge.DeviceFamily)
	d.Set("device_id", edge.DeviceId)
	d.Set("edge_state", edge.EdgeState)
	d.Set("endpoint_pki_mode", edge.EndpointPkiMode)
	d.Set("enterprise_id", edge.EnterpriseId)
	d.Set("factory_software_version", edge.FactorySoftwareVersion)
	d.Set("factory_build_number", edge.FactoryBuildNumber)
	d.Set("ha_previous_state", edge.HaPreviousState)
	d.Set("ha_serial_number", edge.HaSerialNumber)
	d.Set("ha_state", edge.HaState)
	d.Set("model_number", edge.ModelNumber)
	d.Set("self_mac_address", edge.SelfMacAddress)
	d.Set("serial_number", edge.SerialNumber)
	d.Set("site_id", edge.SiteId)
	d.Set("software_version", edge.SoftwareVersion)
	d.Set("is_hub", edge.IsHub)
	d.Set("is_software_version_supported_by_vco", edge.IsSoftwareVersionSupportedByVco)

	d.SetId(fmt.Sprintf("%d", edge.Id))
	return diags
}
