package velocloud

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)


func dataSourceTest() *schema.Resource{
	return  &schema.Resource {
		ReadContext: dataSourceOrderRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type: schema.TypeInt,
				Required: true,
			},
			/*"items": &schema.Schema{
				Type:  schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"text" : &schema.Schema{
							Type:  schema.TypeString,
							Computed: true,
						},
					},
				},
			},*/
			"desc" : &schema.Schema{ 
				Type : schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//c := m.(*hc.Client)

	var diags diag.Diagnostics
	orderID := strconv.Itoa(d.Get("id").(int))

	d.Set("desc", "DESCRIPTION")

	//order, err := c.GetOrder(orderID)
	//if err != nil {
	//	return diag.FromErr(err)
	//}

	d.SetId( orderID )
	return  diags
}

