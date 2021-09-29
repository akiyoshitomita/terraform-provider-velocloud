/*
 * VMware SD-WAN
 *
 * data_source_
 */

package velocloud

import (
	"context"
	"regexp"
	//"strconv"
	//"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-velocloud/velocloud/vcoclient"
)

func dataSourceVeloLicenseList() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVeloLicenseListRead,
		Schema: map[string]*schema.Schema{
			"sku": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "sku name of license",
			},
			"name_regexp": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "search license name by regexp",
			},
			"alias_regexp": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "search license alias by regexp",
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "search region",
			},
			"term_months": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "license term (unit month)",
			},
			"edition": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "edition name",
			},
			"bandwidth_tier": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "bandwidthi tier `001GW` `001G` `050M`",
			},
			"active": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "if true(default), active license only. If false, all license.",
			},
			"licenses": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "license list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"license_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "license id",
						},
						"logical_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "logical id of license",
						},
						"sku": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "SKU of license",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "license name",
						},
						"alias": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "license alias name",
						},
						"regions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "region list of license",
						},
						"term_months": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "license term (unit: month)",
						},
						"edition": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "license edition",
						},
						"bandwidth_tier": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "bandwidth tire of license",
						},
						"active": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If true, license is active",
						},
					},
				},
			},
		},
	}
}

func dataSourceVeloLicenseListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	licenses := []map[string]interface{}{}

	conn := m.(*vcoclient.APIClient).LicenseApi

	post := &vcoclient.LicenseGetEnterpriseEdgeLicenses{}

	res, _, err := conn.LicenseGetEnterpriseEdge(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Println("step1")
	sku := d.Get("sku").(string)
	name_regexp := d.Get("name_regexp").(string)
	alias_regexp := d.Get("alias_regexp").(string)
	region := d.Get("region").(string)
	term_months := d.Get("term_months").(int)
	edition := d.Get("edition").(string)
	bandwidth_tier := d.Get("bandwidth_tier").(string)
	log.Println("step2")
	active := d.Get("active").(bool)
	for _, v := range res {
		if sku != "" && sku != v.Sku {
			continue
		}
		if name_regexp != "" && !(regexp.MustCompile(name_regexp).Match([]byte(v.Name))) {
			continue
		}
		if alias_regexp != "" && !(regexp.MustCompile(alias_regexp).Match([]byte(v.Alias))) {
			continue
		}
		if region != "" {
			a := false
			for _, r := range v.Detail.Regions {
				if r == region {
					a = true
					break
				}
			}
			if !a {
				continue
			}
		}
		if term_months > 0 && term_months != v.TermMonths {
			continue
		}
		if edition != "" && edition != v.Edition {
			continue
		}
		if bandwidth_tier != "" && bandwidth_tier != v.BandwidthTier {
			continue
		}
		if active && v.Active != 1 {
			continue
		}

		l2 := map[string]interface{}{
			"license_id":     v.Id,
			"logical_id":     v.LogicalId,
			"sku":            v.Sku,
			"name":           v.Name,
			"alias":          v.Alias,
			"regions":        v.Detail.Regions,
			"term_months":    v.TermMonths,
			"edition":        v.Edition,
			"bandwidth_tier": v.BandwidthTier,
			"active":         false,
		}
		if v.Active == 1 {
			l2["active"] = true
		}
		log.Println(v.Detail.Regions)

		licenses = append(licenses, l2)
	}
	log.Println("step3")
	d.Set("licenses", licenses)
	d.SetId("license")

	return diags
}
