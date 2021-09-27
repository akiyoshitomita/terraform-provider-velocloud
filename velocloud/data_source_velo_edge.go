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
				Type:        schema.TypeString,
				Required:    true,
				Description: "edge name",
			},
			"edge_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "edge id",
			},
			"activation_key": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "activation key",
			},
			"build_number": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "software build number",
			},
			"custom_info": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "custom info",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "description",
			},
			"device_family": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "device family",
			},
			"device_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "device id",
			},
			"edge_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "edge state",
			},
			"endpoint_pki_mode": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "endpoint_pki_mode",
			},
			"enterprise_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "enterprise_id",
			},
			"factory_software_version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "factory software version",
			},
			"factory_build_number": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "factory software build number",
			},
			"ha_previous_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ha_previous_state",
			},
			"ha_serial_number": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "serial number of HA secundry device",
			},
			"ha_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "HA status",
			},
			"model_number": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "model number",
			},
			"self_mac_address": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "mac address of device",
			},
			"serial_number": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "serial number of device",
			},
			"site_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "site id of edge",
			},
			"software_version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "software version ",
			},
			"is_hub": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "true is configurated hub",
			},
			"is_software_version_supported_by_vco": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If true, the device is using supoerted software version ",
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

	d.Set("edge_id", edge.Id)
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
