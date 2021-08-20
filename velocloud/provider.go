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

//func dataSourceTest() *schema.Resource{
//        return nil
//}

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCO_USER", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("VCO_PASS", nil),
			},
			"vco": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCO_HOST", nil),
			},
			"apikey": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCO_KEY", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			//"hashicups_order": resourceOrder(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"velocloud_edge":  dataSourceVeloEdge(),
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
	//log.Fatal(c.LoginApi)

	// diags = append(diags, diag.Diagnostic{
	//      Severity: diag.Warning,
	//      Summary:  "Warning Message Summary",
	//      Detail:   "This is the detailed warning message from providerConfigure",
	// })

	//if (username != "") && (password != "") {
	//c, err := hashicups.NewClient(nil, &username, &password)
	//if err != nil {
	//diags = append(diags, diag.Diagnostic{
	//Severity: diag.Error,
	//Summary:  "Unable to create HashiCups client",
	//Detail:   "Unable to auth user for authenticated HashiCups client",
	//})
	//return nil, diags
	//}

	//return c, diags
	//}

	//c, err := hashicups.NewClient(nil, nil, nil)
	//if err != nil {
	//diags = append(diags, diag.Diagnostic{
	//Severity: diag.Error,
	//Summary:  "Unable to create HashiCups client",
	//Detail:   "Unable to auth user for authenticated HashiCups client",
	//})
	//return nil, diags
	//}

	//return c, diags
	return nil, diag.Errorf("Missing credentials")
}
