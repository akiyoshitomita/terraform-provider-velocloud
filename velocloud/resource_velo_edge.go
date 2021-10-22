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

func resourceVeloEdge() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVeloEdgeCreate,
		ReadContext:   resourceVeloEdgeRead,
		UpdateContext: resourceVeloEdgeUpdate,
		DeleteContext: resourceVeloEdgeDelete,
		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"profile_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Profile id",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "name of edge",
			},
			"serial_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "serial number of edge. this attribute don't update.",
			},
			"model_number": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "model number of edge. If change this attribute, renew this resource and remove configutaions of this resource",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of edge",
			},
			"site": {
				Type:        schema.TypeList,
				Description: "Site description",
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"site_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "site id of edge",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "site name",
						},
						"logical_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "logical id of site",
						},
						"contact_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "name of site administrator",
						},
						"contact_phone": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "phone number of site administrator",
						},
						"contact_mobile": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "cell phone number of site administrator",
						},
						"contact_email": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "cell phone number of site administrator",
						},
						"street_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "address line 1 of location",
						},
						"street_address2": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "address line 2 of location",
						},
						"city": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "city of location",
						},
						"state": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "state of location",
						},
						"postal_code": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "postal code of location",
						},
						"country": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "country of location",
						},
						"lat": {
							Type:        schema.TypeFloat,
							Optional:    true,
							Description: "lat of location",
						},
						"lon": {
							Type:        schema.TypeFloat,
							Optional:    true,
							Description: "lon of location",
						},
						"timezone": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "timezone. this attribute is don't update",
						},
						"locale": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "locale. this attribute is don't update",
						},
						"shipping_same_as_location": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "if true, shipping address is use location address",
						},
						"shipping_contact_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "name of shipping center",
						},
						"shipping_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "address line 1 of shipping center",
						},
						"shipping_address2": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "address line 2 of shipping center",
						},
						"shipping_city": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "city of shipping center",
						},
						"shipping_state": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "state of shipping center",
						},
						"shipping_country": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "country of shipping center",
						},
						"shipping_postal_code": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "postal code of shipping center",
						},
					},
				},
			},
			"ha_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "if true, set ha configuration the edge. This attribute don't update the configuration on VCO. If you need change this, pleaes use configuretion resource. ",
			},
			"endpoint_pki_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "endpoint pki mode",
			},
			"edge_license_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "edge license id",
			},
			"custom_info": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "custom infomation",
			},
			"edge_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "id of edge",
			},
			"activation_key": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "activation key of edge",
			},
			"build_number": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "current build number of edge",
			},
			"device_family": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "device family name of edge",
			},
			"device_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "device id of edge",
			},
			"edge_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "state of edge",
			},
			"enterprise_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "enterprise id",
			},
			"factory_software_version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "current software version of edge",
			},
			"factory_build_number": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "factory build name of edge",
			},
			"ha_previous_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ha prvios state of edge",
			},
			"ha_serial_number": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "serial number of edge",
			},
			"ha_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ha state of edge",
			},
			"self_mac_address": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "MAC address of edge",
			},
			"software_version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "software version of edge",
			},
		},
	}
}

func resourceVeloEdgeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//var diags diag.Diagnostics
	//log.Println("######### DEGUB CREATE ###############")
	conn := m.(*vcoclient.APIClient).EdgeApi

	site_schema := d.Get("site").([]interface{})
	site_map := site_schema[0].(map[string]interface{})
	var site_samelocation int
	if site_map["shipping_same_as_location"].(bool) {
		site_samelocation = 1
	} else {
		site_samelocation = 0
	}
	//log.Println(site_map)
	// siteの内容を入れる

	post := &vcoclient.EdgeEdgeProvision{
		ConfigurationId: d.Get("profile_id").(int),
		Name:            d.Get("name").(string),
		ModelNumber:     d.Get("model_number").(string),
		Site: &vcoclient.Site{
			Name:                   site_map["name"].(string),
			ContactName:            site_map["contact_name"].(string),
			ContactPhone:           site_map["contact_phone"].(string),
			ContactMobile:          site_map["contact_mobile"].(string),
			ContactEmail:           site_map["contact_email"].(string),
			StreetAddress:          site_map["street_address"].(string),
			StreetAddress2:         site_map["street_address2"].(string),
			City:                   site_map["city"].(string),
			State:                  site_map["state"].(string),
			PostalCode:             site_map["postal_code"].(string),
			Country:                site_map["country"].(string),
			Lat:                    site_map["lat"].(float64),
			Lon:                    site_map["lon"].(float64),
			Timezone:               site_map["timezone"].(string),
			Locale:                 site_map["locale"].(string),
			ShippingSameAsLocation: site_samelocation,
			ShippingContactName:    site_map["shipping_contact_name"].(string),
			ShippingAddress:        site_map["shipping_address"].(string),
			ShippingAddress2:       site_map["shipping_address2"].(string),
			ShippingCity:           site_map["shipping_city"].(string),
			ShippingState:          site_map["shipping_state"].(string),
			ShippingCountry:        site_map["shipping_country"].(string),
			ShippingPostalCode:     site_map["shipping_postal_code"].(string),
		},
		HaEnabled:     d.Get("ha_enabled").(bool),
		EdgeLicenseId: d.Get("edge_license_id").(int),
	}

	if v, ok := d.GetOk("serial_number"); ok {
		post.SerialNumber = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		post.Description = v.(string)
	}
	if v, ok := d.GetOk("endpoint_pki_mode"); ok {
		post.EndpointPkiMode = v.(string)
	}
	if v, ok := d.GetOk("custom_info"); ok {
		post.CustomInfo = v.(string)
	}

	res, _, err := conn.EdgeEdgeProvision(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", res.Id))

	return resourceVeloEdgeRead(ctx, d, m)
}

func resourceVeloEdgeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	edgeId, _ := strconv.Atoi(d.Id())
	conn := m.(*vcoclient.APIClient).EdgeApi

	post := &vcoclient.EdgeGetEdge{
		Id:   edgeId,
		With: []string{"configuration", "site"},
	}

	res, _, err := conn.EdgeGetEdge(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("profile_id", res.Configuration.Id)
	d.Set("name", res.Name)
	d.Set("serial_number", res.SerialNumber)
	d.Set("model_number", res.ModelNumber)
	d.Set("description", res.Description)

	site_schema := d.Get("site").([]interface{})
	site_map := site_schema[0].(map[string]interface{})
	site_map["site_id"] = res.Site.Id
	site_map["name"] = res.Site.Name
	site_map["logical_id"] = res.Site.LogicalId
	site_map["contact_name"] = res.Site.ContactName
	site_map["contact_phone"] = res.Site.ContactMobile
	site_map["contact_email"] = res.Site.ContactEmail
	site_map["street_address"] = res.Site.StreetAddress
	site_map["street_address2"] = res.Site.StreetAddress2
	site_map["city"] = res.Site.City
	site_map["state"] = res.Site.State
	site_map["postal_code"] = res.Site.PostalCode
	site_map["country"] = res.Site.Country
	site_map["lat"] = res.Site.Lat
	site_map["lon"] = res.Site.Lon
	site_map["timezone"] = res.Site.Timezone
	site_map["locale"] = res.Site.Locale
	site_map["shipping_contact_name"] = res.Site.ShippingContactName
	site_map["shipping_address"] = res.Site.ShippingAddress
	site_map["shipping_address2"] = res.Site.ShippingAddress2
	site_map["shipping_city"] = res.Site.ShippingCity
	site_map["shipping_state"] = res.Site.ShippingState
	site_map["shipping_country"] = res.Site.ShippingCountry
	site_map["shipping_postal_code"] = res.Site.ShippingPostalCode
	d.Set("site", []interface{}{site_map})

	if res.HaState == "UNCONFIGURED" {
		d.Set("ha_enabled", false)
	} else {
		d.Set("ha_enabled", true)
	}
	//d.Set("endpoint_pki_mode", res.Description)
	//d.Set("edge_license_id", res.EdgeLicenseId)
	d.Set("custom_info", res.CustomInfo)

	// Computed
	d.Set("edge_id", res.Id)
	d.Set("activation_key", res.ActivationKey)
	d.Set("build_number", res.BuildNumber)
	d.Set("device_family", res.DeviceFamily)
	d.Set("device_id", res.DeviceId)
	d.Set("edge_state", res.EdgeState)
	d.Set("enterprise_id", res.EnterpriseId)
	d.Set("factory_software_version", res.FactorySoftwareVersion)
	d.Set("factory_build_number", res.FactoryBuildNumber)
	d.Set("ha_previous_state", res.HaPreviousState)
	d.Set("ha_serial_number", res.HaSerialNumber)
	d.Set("ha_state", res.HaState)
	d.Set("self_mac_address", res.SelfMacAddress)
	d.Set("software_version", res.SoftwareVersion)

	return diags
}

func resourceVeloEdgeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var isUpdate bool

	edgeId, _ := strconv.Atoi(d.Id())
	conn := m.(*vcoclient.APIClient).EdgeApi

	update := &vcoclient.EdgeUpdateEdgeAttributesUpdate{}
	if d.HasChange("site") {
		isUpdate = true

		site_schema := d.Get("site").([]interface{})
		site_map := site_schema[0].(map[string]interface{})
		var site_samelocation int
		if site_map["shipping_same_as_location"].(bool) {
			site_samelocation = 1
		} else {
			site_samelocation = 0
		}

		site := &vcoclient.Site{
			Name:                   site_map["name"].(string),
			ContactName:            site_map["contact_name"].(string),
			ContactPhone:           site_map["contact_phone"].(string),
			ContactMobile:          site_map["contact_mobile"].(string),
			ContactEmail:           site_map["contact_email"].(string),
			StreetAddress:          site_map["street_address"].(string),
			StreetAddress2:         site_map["street_address2"].(string),
			City:                   site_map["city"].(string),
			State:                  site_map["state"].(string),
			PostalCode:             site_map["postal_code"].(string),
			Country:                site_map["country"].(string),
			Lat:                    site_map["lat"].(float64),
			Lon:                    site_map["lon"].(float64),
			Timezone:               site_map["timezone"].(string),
			Locale:                 site_map["locale"].(string),
			ShippingSameAsLocation: site_samelocation,
			ShippingContactName:    site_map["shipping_contact_name"].(string),
			ShippingAddress:        site_map["shipping_address"].(string),
			ShippingAddress2:       site_map["shipping_address2"].(string),
			ShippingCity:           site_map["shipping_city"].(string),
			ShippingState:          site_map["shipping_state"].(string),
			ShippingCountry:        site_map["shipping_country"].(string),
			ShippingPostalCode:     site_map["shipping_postal_code"].(string),
		}
		update.Site = site
	}

	if d.HasChange("name") {
		isUpdate = true
		update.Name = d.Get("name").(string)
	}

	if d.HasChange("description") {
		isUpdate = true
		update.Description = d.Get("description").(string)
	}

	if d.HasChange("serial_number") {
		isUpdate = true
		update.SerialNumber = d.Get("serial_number").(string)
	}

	if d.HasChange("custom_info") {
		isUpdate = true
		update.CustomInfo = d.Get("custom_info").(string)
	}

	if isUpdate {

		post := &vcoclient.EdgeUpdateEdgeAttributes{
			Id:     edgeId,
			Update: update,
		}

		_, _, err := conn.EdgeUpdateEdgeAttributes(nil, *post)
		if err != nil {
			return diag.FromErr(err)
		}

	} else {
		log.Printf("[WARN] couldn't update VCO")
	}
	return resourceVeloEdgeRead(ctx, d, m)

}

func resourceVeloEdgeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	edgeId, _ := strconv.Atoi(d.Id())
	conn := m.(*vcoclient.APIClient).EdgeApi

	post := &vcoclient.EdgeDeleteEdge{
		Id: edgeId,
	}

	res, _, err := conn.EdgeDeleteEdge(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}

	if res.Rows != 1 {
		return diag.Errorf("ERROR couldn't delete edge")
	}

	return diags
}
