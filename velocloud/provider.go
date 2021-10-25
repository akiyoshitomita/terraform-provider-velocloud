package velocloud

import (
	"context"
	//"fmt"
	"log"

	//"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-velocloud/velocloud/vcoclient"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCO_USER", nil),
				Description: "Login user name of Velocloud Orchestrator. this user role requires an enterprise superuser.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("VCO_PASS", nil),
				Description: "password of the login user.",
			},
			"vco": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCO_HOST", nil),
				Description: "FQDN of Velocloud Orchestrator.",
			},
			"apikey": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCO_KEY", nil),
				Description: "API Token of Enteprise level. This attirbute cannot be used at same time as 'username'.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"velocloud_edge":     resourceVeloEdge(),
			"velocloud_profile":  resourceVeloProfile(),
			"velocloud_firewall": resourceVeloFirewall(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"velocloud_address_gruop": dataSourceVeloAddressGroup(),
			"velocloud_edge":          dataSourceVeloEdge(),
			"velocloud_license_list":  dataSourceVeloLicenseList(),
			"velocloud_port_gruop":    dataSourceVeloPortGroup(),
			"velocloud_profile":       dataSourceVeloProfile(),
			"velocloud_segment":       dataSourceVeloSegment(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	vco := d.Get("vco").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	apikey := d.Get("apikey").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	c := vcoclient.NewAPIClient(&vcoclient.Configuration{
		UserAgent:     "Terraform velocloud Agent/1.0.0/go",
		BasePath:      "https://" + vco + "/portal/",
		DefaultHeader: make(map[string]string),
		Idcount:       0,
	})

	if (apikey != "") && (username == "") && (password == "") && (vco != "") {
		c.AddHeader("Authorization", "Token "+apikey)
		return c, diags
	}

	if (username != "") && (password != "") && (vco != "") && (apikey == "") {
		log.Println("username authentication")
		auth := &vcoclient.AuthObject{
			Username: username,
			Password: password,
		}
		_, err := c.LoginApi.LoginEnterpriseLogin(nil, *auth)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return c, diags
	}

	return nil, diag.Errorf("Missing credentials")
}
