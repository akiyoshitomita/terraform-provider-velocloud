/*
 * VMware SD-WAN
 *
 * resource_
 */

package velocloud

import (
	"context"
	//"encoding/json"
	"fmt"
	//"log"
	"net"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-velocloud/velocloud/vcoclient"
)

func resourceVeloFirewall() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVeloFirewallCreate,
		ReadContext:   resourceVeloFirewallRead,
		UpdateContext: resourceVeloFirewallUpdate,
		DeleteContext: resourceVeloFirewallDelete,
		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"profile_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Default:     0,
				Description: "When you set this attribute, this resource change firewall settigns of Profile. When you set this, don't set edge_id at the same time.",
			},
			"edge_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Default:     0,
				Description: "When you set this attribute, this resource change firewall settings of Edge. When you set this, don't set profile_id at the same time",
			},
			"firewall_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "when true, firewall function is enable",
			},
			"statefull_firewall_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "statefull firewall is enable",
			},
			"edge_overwrite_statefull_firewall_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "if false, use firewall_enalbed of profile, this options need edge_id",
			},
			"syslog_forwarding": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "syslog forwarding is enable.",
			},
			"edge_overwrite_syslog_forwarding": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "if false, use syslog_forwarding of profile. this options need edge_id",
			},
			"segments": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "firewall segment",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
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
							Optional:    true,
							Description: "segment id",
						},
						"segment_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "segment type",
						},
						"firewall_rule": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "fireewall Rules of segment",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Rule Name",
									},
									"source_type": { // any,objectgroup,define
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Match source type [any,objectgroup,none,vlan,interface,ipaddress,macaddress]",
										Default:     "any",
									},
									"source_address_group_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Srouce Address Group id. must set source_type to 'objectgroup'.",
									},
									"source_port_group_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Srouce port Group id. must set source_type to 'objectgroup'.",
									},
									"source_vlan": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "source vlan. must set source_type to 'vlan'",
										Default:     -1,
									},
									"source_interface": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "source interface. must set source_type to 'interface'",
									},
									"source_ipaddress": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "source ip address. must set source_type to 'ipaddress'",
									},
									"source_cidr_prefix": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "source prefix. must set source_type to 'ipaddress'. must not set source_subnet_mask and source_wildcard_mask.",
									},
									"source_subnet_mask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "source mask. must set source_type to 'ipaddress'. must not set source_cidr_prefix and source_wildcard_mask.",
									},
									"source_wildcard_mask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "source mask. must set source_type to 'ipaddress'. must not set source_cidr_prefix and source_subnet_mask.",
									},
									"source_mac": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Source mac address. must set source_type to 'macaddress'",
									},
									"source_port": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Source port range. ex[80-443] . must nost set source_type to 'any' or 'objectgroup'",
									},
									"destination_type": { //any,objectgroup,define
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Match destination type [any,objectgroup,none,vlan,interface,ipaddress,macaddress]",
									},
									"destination_address_group_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Destination Address Group id. must set source_type to 'objectgroup'.",
									},
									"destination_port_group_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Destination port Group id. must set source_type to 'objectgroup'.",
									},
									"destination_vlan": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "source vlan. must set source_type to 'vlan'",
										Default:     -1,
									},
									"destination_interface": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "destination interface. must set source_type to 'interface'",
									},
									"destination_ipaddress": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "destination ip address. must set destination_type to 'ipaddress'",
									},
									"destination_cidr_prefix": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "destination prefix. must set destination_type to 'ipaddress'. must not set destination_subnet_mask and destination_wildcard_mask.",
									},
									"destination_subnet_mask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "destination mask. must set destination_type to 'ipaddress'. must not set destination_cidr_prefix and destination_wildcard_mask.",
									},
									"destination_wildcard_mask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "destination mask. must set destination_type to 'ipaddress'. must not set destination_cidr_prefix and destination_subnet_mask.",
									},
									"destination_mac": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "destination mac address. must set destination_type to 'macaddress'",
									},
									"destination_port": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "destination port range. ex[80-443] . must not set destination_type to 'any' or 'objectgroup'",
									},
									"application_id": { //any or define
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Match application. default is any application(-1).",
										Default:     -1,
									},
									"dscp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "DSCP of Match application",
										Default:     -1,
									},
									"action": { //allow,drop,reject,skip
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Firewall action [allow,drop,reject,skip]",
										Default:     "allow",
									},
									"logging": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "enable firewall log",
										Default:     false,
									},
									"audit_comment": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "audit comment",
									},
								},
							},
						},
					},
				},
			},
			"port_forwarding_rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "port forwarding rule. must set edge_id.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "rule name",
						},
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "protocol name [tcp,udp]",
						},
						"interface": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "interface name [GE1,...]",
						},
						"outside_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "outbound ip address",
							Default:     "any",
						},
						"wan_ports": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "wan ports range",
						},
						"lan_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "lan ip address",
						},
						"lan_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "lan port number",
						},
						"segment_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "segment_id",
						},
						"remote_ip_subnet": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "remote ip/segment [192.168.0.1/24]",
						},
					},
				},
			},
			"nat_rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "1:1 nat rule. must set edge_id.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "port forwarding name",
						},
						"outbound_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "outbound ip address",
						},
						"interface": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "interface name [GE1,...]",
						},
						"inside_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "nside ip address",
						},
						"segment_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "segment_id",
						},
						"outbound_traffic": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "outbound traffic",
						},
						"allow_protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "allow protocol name [all,tcp,udpi,icmp,gre]",
							Default:     "all",
						},
						"allow_ports": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "allow destination ports ",
						},
						"remote_ip_subnet": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "remote ip/segment [192.168.0.1/24]",
						},
					},
				},
			},
			"stateful_firewall": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "firewall segment",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"edge_overwrite": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "if false, use statefull_firewall of profile. this options need edge_id",
						},
						"establieshd_tcp_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Established TCP Flow Timeout of Statefull firewall settings (seconds)",
							Default:     7440,
						},
						"non_established_tcp_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Non Established TCP Flow Timeout of Statefull firewall settings (seconds)",
							Default:     240,
						},
						"udp_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "UDP Flow Timeout of Statefull firewall settings (seconds)",
							Default:     300,
						},
						"other_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Other Flow Timeout of Statefull firewall settings (seconds)",
							Default:     60,
						},
					},
				},
			},
			"network_flood_protection": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "firewall segment",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"edge_overwrite": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "if false, use network_flood_protection of profile. this options need edge_id",
						},
						"new_connection_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "New Connection Threshold(connections per second).",
						},
						"denylist": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "DenyList",
						},
						"detection_time": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Detect Duration time(seconds). must set denylist to true",
						},
						"denylist_time": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Denylist Duration (seconds). must set denylist to true",
						},
						"invalid_tcp_flags": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "DenyList",
						},
						"tcp_land": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "DenyList",
						},
						"tcp_syn_fragment": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "DenyList",
						},
						"icmp_ping_of_death": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "DenyList",
						},
						"icmp_fragment": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "DenyList",
						},
						"ip_unkown_protocol": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "DenyList",
						},
						"ip_options": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "DenyList",
						},
					},
				},
			},
			"edge_access": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "firewall segment",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"edge_overwrite": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "if false, use network_flood_protection of profile. this options need edge_id",
						},
						"ssh": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "enable ssh",
						},
						"ssh_allow": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "allow ip address list of ssh access",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"console": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "enable console port ",
						},
						"usb": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "enable usb port (only applicable for edge models 510 and 6X0)",
						},
						"snmp": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "enable SNMP",
						},
						"snmp_allow": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "allow ip address list of SNMP. if null , allow all LAN",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"webui": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "enable Local web UI access",
						},
						"webui_allow": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "allow ip address list of web UI. if null, allow alli LAN",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"webui_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Local web UI port number",
							Default:     80,
						},
					},
				},
			},
		},
	}
}

func resourceVeloFirewallCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// 確認
	configId := d.Get("profile_id").(int)
	edgeId := d.Get("edge_id").(int)

	if configId == 0 && edgeId == 0 {
		return diag.Errorf("[ERROR] Configuratoin Error. Please set profile_id or edge_id")
	}
	if configId != 0 && edgeId != 0 {
		return diag.Errorf("[ERROR] Configuratoin Error. don't set edge_id and profile_id at same resource")
	}

	if edgeId != 0 {
		conn := m.(*vcoclient.APIClient).EdgeApi
		post := &vcoclient.EdgeGetEdgeConfigurationStack{
			EdgeId: edgeId,
		}

		res, _, err := conn.EdgeGetEdgeConfigurationStack(nil, *post)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(fmt.Sprintf("%d", res[0].Id))

		// データがあるか確認する
		newfirewall := true
		for _, v := range res[0].Modules {
			if v.Name == "firewall" {
				newfirewall = false
				break
			}
		}
		// データが無い場合は作成する
		if newfirewall {
			conn := m.(*vcoclient.APIClient).ConfigurationApi
			firewallData := &vcoclient.FirewallData{
				FirewallEnabled: true,
			}
			post := &vcoclient.ConfigurationInsertConfigurationModule{
				ConfigurationId: res[0].Id,
				Name:            "firewall",
				Data:            firewallData,
			}
			_, _, err := conn.ConfigurationInsertConfigurationModule(nil, *post)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	} else {
		d.SetId(fmt.Sprintf("%d", configId))
	}

	return resourceVeloFirewallUpdate(ctx, d, m)
}

func resourceVeloFirewallRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		diags          diag.Diagnostics
		isEdge         bool
		segments       []interface{}
		portForwarding []interface{}
		natRule        []interface{}
		sfirewall      map[string]interface{}
		nfp            map[string]interface{}
		service        map[string]interface{}
	)

	isEdge = d.Get("edge_id").(int) != 0

	configId, _ := strconv.Atoi(d.Id())
	conn := m.(*vcoclient.APIClient).ConfigurationApi
	post := &vcoclient.ConfigurationGetConfigurationModules{
		ConfigurationId: configId,
		Modules:         []string{"firewall"},
	}

	res1, _, err := conn.ConfigurationGetConfigurationModulesFirewall(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(res1) != 1 {
		return diag.Errorf("[ERROR] firewall configuration times error.")
	}

	firewallData := res1[0].Data
	d.Set("firewall_enabled", firewallData.FirewallEnabled)

	if firewallData.StatefulFirewallEnabled == nil {
		d.Set("edge_overwrite_statefull_firewall_enabled", false)
	} else {
		if isEdge == true {
			d.Set("edge_overwrite_statefull_firewall_enabled", true)
		}
		d.Set("statefull_firewall_enabled", firewallData.StatefulFirewallEnabled)
	}

	if firewallData.SyslogForwarding == nil {
		d.Set("edge_overwrite_syslog_forwarding", false)
	} else {
		if isEdge == true {
			d.Set("edge_overwrite_syslog_forwarding", true)
		}
		d.Set("syslog_forwarding", firewallData.SyslogForwarding)
	}

	// firewall rule
	for _, segment := range firewallData.Segments {
		seg_tmp := map[string]interface{}{}
		seg_tmp["name"] = segment.Segment.Name
		seg_tmp["segment_logical_id"] = segment.Segment.SegmentLogicalId
		seg_tmp["segment_id"] = segment.Segment.SegmentId
		seg_tmp["segment_type"] = segment.Segment.Type
		rules := []interface{}{}
		for _, firewall_rule := range segment.Outbound {
			rule := map[string]interface{}{}
			rule["name"] = firewall_rule.Name
			rule["source_type"] = "any"
			if firewall_rule.Match.SAddressGroup != "" {
				rule["source_address_group_id"] = firewall_rule.Match.SAddressGroup
				rule["source_type"] = "objectgroup"
			}
			if firewall_rule.Match.SPortGroup != "" {
				rule["source_port_group_id"] = firewall_rule.Match.SPortGroup
				rule["source_type"] = "objectgroup"
			}
			if firewall_rule.Match.SportLow != -1 {
				if firewall_rule.Match.SportLow == firewall_rule.Match.SportHigh {
					rule["source_port"] = fmt.Sprintf("%d", firewall_rule.Match.SportLow)
				} else {
					rule["source_port"] = fmt.Sprintf("%d-%d", firewall_rule.Match.SportLow, firewall_rule.Match.SportHigh)
				}
				rule["source_type"] = "none"
			}
			if firewall_rule.Match.Svlan != -1 {
				rule["source_vlan"] = firewall_rule.Match.Svlan
				rule["source_type"] = "vlan"
			} else {
				rule["source_vlan"] = -1
			}
			if firewall_rule.Match.SInterface != "" {
				rule["source_interface"] = firewall_rule.Match.SInterface
				rule["source_type"] = "interface"
			}
			if firewall_rule.Match.Sip != "any" {
				rule["source_ipaddress"] = firewall_rule.Match.Sip
				rule["source_type"] = "ipaddress"
				if firewall_rule.Match.SRuleType == "prefix" {
					mask := net.ParseIP(firewall_rule.Match.Ssm)
					l := 0
					for i := 12; i < 16; i++ {
						for ; mask[i] != 0; mask[i] = mask[i] << 1 {
							l++
						}
					}
					rule["source_cidr_prefix"] = l
				} else if firewall_rule.Match.SRuleType == "netmask" {
					rule["source_subnet_mask"] = firewall_rule.Match.Ssm
				} else if firewall_rule.Match.SRuleType == "wildcard" {
					mask := net.ParseIP(firewall_rule.Match.Ssm)
					for i := 12; i < 16; i++ {
						mask[i] = mask[i] ^ 0xff
					}
					rule["source_wildcard_mask"] = mask.String()
				}
			}
			if firewall_rule.Match.Smac != "any" && firewall_rule.Match.Smac != "" {
				rule["source_mac"] = firewall_rule.Match.Smac
				rule["source_type"] = "macaddress"
			}

			rule["destination_type"] = "any"
			if firewall_rule.Match.DAddressGroup != "" {
				rule["destination_address_group_id"] = firewall_rule.Match.DAddressGroup
				rule["destination_type"] = "objectgroup"
			}
			if firewall_rule.Match.DPortGroup != "" {
				rule["destination_port_group_id"] = firewall_rule.Match.DPortGroup
				rule["destination_type"] = "objectgroup"
			}
			if firewall_rule.Match.DportLow != -1 {
				if firewall_rule.Match.DportLow == firewall_rule.Match.DportHigh {
					rule["destination_port"] = fmt.Sprintf("%d", firewall_rule.Match.DportLow)
				} else {
					rule["destination_port"] = fmt.Sprintf("%d-%d", firewall_rule.Match.DportLow, firewall_rule.Match.DportHigh)
				}
				rule["destination_type"] = "none"
			}
			if firewall_rule.Match.Dvlan != -1 {
				rule["destination_vlan"] = firewall_rule.Match.Dvlan
				rule["destination_type"] = "vlan"
			} else {
				rule["destination_vlan"] = -1
			}
			if firewall_rule.Match.DInterface != "" {
				rule["destination_interface"] = firewall_rule.Match.DInterface
				rule["destination_type"] = "interface"
			}
			if firewall_rule.Match.Dip != "any" {
				rule["destination_ipaddress"] = firewall_rule.Match.Dip
				rule["destination_type"] = "ipaddress"
				if firewall_rule.Match.DRuleType == "prefix" {
					mask := net.ParseIP(firewall_rule.Match.Dsm)
					l := 0
					for i := 12; i < 16; i++ {
						for ; mask[i] != 0; mask[i] = mask[i] << 1 {
							l++
						}
					}
					rule["destination_cidr_prefix"] = l
				} else if firewall_rule.Match.DRuleType == "netmask" {
					rule["destination_subnet_mask"] = firewall_rule.Match.Dsm
				} else if firewall_rule.Match.DRuleType == "wildcard" {
					mask := net.ParseIP(firewall_rule.Match.Dsm)
					for i := 12; i < 16; i++ {
						mask[i] = mask[i] ^ 0xff
					}
					rule["destination_wildcard_mask"] = mask.String()
				}
			}
			if firewall_rule.Match.Dmac != "any" && firewall_rule.Match.Dmac != "" {
				rule["destination_mac"] = firewall_rule.Match.Dmac
				rule["destination_type"] = "macaddress"
			}

			rule["application_id"] = firewall_rule.Match.Appid
			rule["dscp"] = firewall_rule.Match.Dscp

			rule["logging"] = firewall_rule.LoggingEnabled
			rule["action"] = firewall_rule.Action.AllowOrDeny
			rules = append(rules, rule)
		}
		seg_tmp["firewall_rule"] = rules
		segments = append(segments, seg_tmp)
	}
	// Inbound ACLs
	for _, inboundAcl := range firewallData.Inbound {
		if inboundAcl.Action.Type == "port_forwarding" {
			iacl_tmp := map[string]interface{}{}
			iacl_tmp["name"] = inboundAcl.Name
			if inboundAcl.Match.Proto == 6 {
				iacl_tmp["protocol"] = "tcp"
			} else if inboundAcl.Match.Proto == 17 {
				iacl_tmp["protocol"] = "udp"
			}
			iacl_tmp["interface"] = inboundAcl.Action.Interface
			iacl_tmp["outside_ip"] = inboundAcl.Match.Dip
			if inboundAcl.Match.DportLow == inboundAcl.Match.DportHigh {
				iacl_tmp["wan_ports"] = fmt.Sprintf("%d", inboundAcl.Match.DportLow)
			} else {
				iacl_tmp["wan_ports"] = fmt.Sprintf("%d-%d", inboundAcl.Match.DportLow, inboundAcl.Match.DportHigh)
			}
			iacl_tmp["lan_ip"] = inboundAcl.Action.Nat.LanIp
			iacl_tmp["lan_port"] = inboundAcl.Action.Nat.LanPort
			iacl_tmp["segment_id"] = inboundAcl.Action.SegmentId
			if inboundAcl.Match.Sip != "any" {
				mask := net.ParseIP(inboundAcl.Match.Ssm)
				l := 0
				for i := 12; i < 16; i++ {
					for ; mask[i] != 0; mask[i] = mask[i] << 1 {
						l++
					}
				}
				iacl_tmp["remote_ip_subnet"] = fmt.Sprintf("%s/%d", inboundAcl.Match.Sip, l)
			}
			portForwarding = append(portForwarding, iacl_tmp)
		} else if inboundAcl.Action.Type == "one_to_one_nat" {
			iacl_tmp := map[string]interface{}{}
			iacl_tmp["name"] = inboundAcl.Name
			iacl_tmp["outbound_ip"] = inboundAcl.Match.Dip
			iacl_tmp["interface"] = inboundAcl.Action.Interface
			iacl_tmp["inside_ip"] = inboundAcl.Action.Nat.LanIp
			iacl_tmp["segment_id"] = inboundAcl.Action.SegmentId
			iacl_tmp["outbound_traffic"] = inboundAcl.Action.Nat.Outbound
			if inboundAcl.Match.Proto == -1 {
				iacl_tmp["allow_protocol"] = "all"
			} else if inboundAcl.Match.Proto == 1 {
				iacl_tmp["allow_protocol"] = "icmp"
			} else if inboundAcl.Match.Proto == 6 {
				iacl_tmp["allow_protocol"] = "tcp"
			} else if inboundAcl.Match.Proto == 17 {
				iacl_tmp["allow_protocol"] = "udp"
			} else if inboundAcl.Match.Proto == 47 {
				iacl_tmp["allow_protocol"] = "gre"
			}
			if inboundAcl.Match.DportLow == -1 {
			} else if inboundAcl.Match.DportLow == inboundAcl.Match.DportHigh {
				iacl_tmp["allow_ports"] = fmt.Sprintf("%d", inboundAcl.Match.DportLow)
			} else {
				iacl_tmp["allow_ports"] = fmt.Sprintf("%d-%d", inboundAcl.Match.DportLow, inboundAcl.Match.DportHigh)
			}
			if inboundAcl.Match.Sip != "any" {
				mask := net.ParseIP(inboundAcl.Match.Ssm)
				l := 0
				for i := 12; i < 16; i++ {
					for ; mask[i] != 0; mask[i] = mask[i] << 1 {
						l++
					}
				}
				iacl_tmp["remote_ip_subnet"] = fmt.Sprintf("%s/%d", inboundAcl.Match.Sip, l)
			}
			natRule = append(natRule, iacl_tmp)
		}
	}
	// Stateful Firewall Settings
	//service
	sfirewall = map[string]interface{}{}
	if firewallData.StatefulFirewallSettings == nil {
		sfirewall["edge_overwrite"] = false
		sfirewall["establieshd_tcp_timeout"] = 7440
		sfirewall["non_established_tcp_timeout"] = 240
		sfirewall["udp_timeout"] = 300
		sfirewall["other_timeout"] = 60
	} else {
		if isEdge == true {
			sfirewall["edge_overwrite"] = true
		}
		sfirewall["establieshd_tcp_timeout"] = firewallData.StatefulFirewallSettings.EstablishedTcpFlowTimeout
		sfirewall["non_established_tcp_timeout"] = firewallData.StatefulFirewallSettings.NonEstablishedTcpFlowTimeout
		sfirewall["udp_timeout"] = firewallData.StatefulFirewallSettings.UdpFlowTimeout
		sfirewall["other_timeout"] = firewallData.StatefulFirewallSettings.OtherFlowTimeout
	}
	// Network & Flood Protection Settings
	nfp = map[string]interface{}{}
	if firewallData.NetworkProtectionSettings == nil {
		nfp["edge_overwrite"] = false
	} else {
		if isEdge == true {
			nfp["edge_overwrite"] = true
		}
		nfp["new_connection_threshold"] = firewallData.NetworkProtectionSettings.NewConnectionThreshold
		if firewallData.NetworkProtectionSettings.Denylist {
			nfp["denylist"] = true
			nfp["detection_time"] = firewallData.NetworkProtectionSettings.DetectionTime
			nfp["denylist_time"] = firewallData.NetworkProtectionSettings.DenylistDuration
		} else {
			nfp["denylist"] = false
		}
		nfp["invalid_tcp_flags"] = firewallData.NetworkProtectionSettings.TcpBasedAttacks.InvalidFlags
		nfp["tcp_land"] = firewallData.NetworkProtectionSettings.TcpBasedAttacks.EnableLand
		nfp["tcp_syn_fragment"] = firewallData.NetworkProtectionSettings.TcpBasedAttacks.EnableSynFragment
		nfp["icmp_ping_of_death"] = firewallData.NetworkProtectionSettings.IcmpBasedAttacks.EnablePingOfDeath
		nfp["icmp_fragment"] = firewallData.NetworkProtectionSettings.IcmpBasedAttacks.EnableFragment
		nfp["ip_unkown_protocol"] = firewallData.NetworkProtectionSettings.IpBasedAttacks.EnableUnknownProtocol
		nfp["ip_options"] = firewallData.NetworkProtectionSettings.IpBasedAttacks.EnableInsecureOptions
	}
	// Edge Access
	service = map[string]interface{}{}
	if firewallData.Services == nil {
		service["edge_overwrite"] = false
	} else {
		if isEdge == true {
			service["edge_overwrite"] = true
		}
		service["ssh"] = firewallData.Services.Ssh.Enabled
		service["ssh_allow"] = firewallData.Services.Ssh.AllowSelectedIp
		service["console"] = firewallData.Services.Console.Enabled
		service["usb"] = !firewallData.Services.UsbDisabled
		service["snmp"] = firewallData.Services.Snmp.Enabled
		service["snmp_allow"] = firewallData.Services.Snmp.AllowSelectedIp
		service["webui"] = firewallData.Services.LocalUi.Enabled
		service["webui_allow"] = firewallData.Services.LocalUi.AllowSelectedIp
		service["webui_port"] = firewallData.Services.LocalUi.PortNumber
	}

	d.Set("segments", segments)
	d.Set("port_forwarding_rule", portForwarding)
	d.Set("nat_rule", natRule)
	d.Set("stateful_firewall", []interface{}{sfirewall})
	d.Set("network_flood_protection", []interface{}{nfp})
	d.Set("edge_access", []interface{}{service})
	return diags
}

func resourceVeloFirewallUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		newlen  int
		oldlen  int
		oldrule *vcoclient.FirewallOutboundRule
	)
	configId, _ := strconv.Atoi(d.Id())

	isEdge := d.Get("edge_id").(int) != 0

	conn := m.(*vcoclient.APIClient).ConfigurationApi
	post := &vcoclient.ConfigurationGetConfigurationModules{
		ConfigurationId: configId,
		Modules:         []string{"firewall"},
	}
	res1, _, err := conn.ConfigurationGetConfigurationModulesFirewall(nil, *post)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(res1) != 1 {
		return diag.Errorf("[ERROR] firewall configuration times error.")
	}

	firewallData := res1[0].Data

	firewallData.FirewallEnabled = d.Get("firewall_enabled").(bool)

	if isEdge && !d.Get("edge_overwrite_statefull_firewall_enabled").(bool) {
		firewallData.StatefulFirewallEnabled = nil
	} else {
		stetafulFirewall := d.Get("statefull_firewall_enabled").(bool)
		firewallData.StatefulFirewallEnabled = &stetafulFirewall
	}

	if isEdge && !d.Get("edge_overwrite_syslog_forwarding").(bool) {
		firewallData.SyslogForwarding = nil
	} else {
		syslogForwarding := d.Get("syslog_forwarding").(bool)
		firewallData.SyslogForwarding = &syslogForwarding
	}

	//firewall
	segments := d.Get("segments").([]interface{})
	newlen = len(segments)
	oldlen = len(firewallData.Segments)
	for newlen > oldlen {
		firewallData.Segments = append(firewallData.Segments, vcoclient.FirewallSegment{})
		oldlen++
	}
	if newlen < oldlen {
		firewallData.Segments = firewallData.Segments[:newlen]
	}
	for s := 0; s < len(firewallData.Segments); s++ {
		//　セグメントのアップデート
		if firewallData.Segments[s].Segment.SegmentId != int32(d.Get(fmt.Sprintf("segments.%d.segment_id", s)).(int)) {
			post3 := &vcoclient.EnterpriseGetEnterpriseNetworkSegments{}
			segs, _, err := m.(*vcoclient.APIClient).EnterpriseApi.EnterpriseGetEnterpriseNetworkSegments(nil, *post3)
			if err != nil {
				return diag.FromErr(err)
			}
			for _, v := range segs {
				if int32(d.Get(fmt.Sprintf("segments.%d.segment_id", s)).(int)) == v.Data.SegmentId {
					firewallData.Segments[s].Segment.SegmentId = v.Data.SegmentId
					firewallData.Segments[s].Segment.Name = v.Name
					firewallData.Segments[s].Segment.SegmentLogicalId = v.LogicalId
					firewallData.Segments[s].Segment.Type = v.Type_
					break
				}
			}
		}

		// ルールの数をそろえる
		rules := d.Get(fmt.Sprintf("segments.%d.firewall_rule", s)).([]interface{})
		newlen = len(rules)
		oldlen = len(firewallData.Segments[s].Outbound)
		for newlen > oldlen {
			firewallData.Segments[s].Outbound = append(firewallData.Segments[s].Outbound, vcoclient.FirewallOutboundRule{})
			oldlen++
		}
		if newlen < oldlen {
			firewallData.Segments[s].Outbound = firewallData.Segments[s].Outbound[:newlen]
		}

		for r := 0; r < len(rules); r++ {
			newrule := rules[r].(map[string]interface{})
			oldrule = &firewallData.Segments[s].Outbound[r]
			oldrule.Name = newrule["name"].(string)

			stype := newrule["source_type"].(string)

			oldrule.Match.SAddressGroup = ""
			oldrule.Match.SPortGroup = ""
			oldrule.Match.SportLow = -1
			oldrule.Match.SportHigh = -1
			oldrule.Match.Svlan = -1
			oldrule.Match.SInterface = ""
			oldrule.Match.Sip = "any"
			oldrule.Match.SRuleType = "prefix"
			oldrule.Match.Ssm = "255.255.255.255"
			oldrule.Match.Smac = ""
			if stype == "any" {
			} else if stype == "objectgroup" {
				oldrule.Match.SAddressGroup = newrule["source_address_group_id"].(string)
				oldrule.Match.SPortGroup = newrule["source_port_group_id"].(string)
			} else {
				ports := strings.Split(newrule["source_port"].(string), "-")
				plen := len(ports)
				if plen == 1 && ports[0] == "" {
				} else if plen == 1 {
					v, err := strconv.Atoi(ports[0])
					if err != nil {
						return diag.FromErr(err)
					}
					oldrule.Match.SportLow = int32(v)
					oldrule.Match.SportHigh = int32(v)
				} else if plen == 2 {
					v, err := strconv.Atoi(ports[0])
					if err != nil {
						return diag.FromErr(err)
					}
					oldrule.Match.SportLow = int32(v)
					v, err = strconv.Atoi(ports[1])
					if err != nil {
						return diag.FromErr(err)
					}
					oldrule.Match.SportHigh = int32(v)
				} else {
					return diag.Errorf("[ERROR] source_port is error.[%s]", newrule["source_port"].(string))
				}

				if stype == "none" {
				} else if stype == "vlan" {
					oldrule.Match.Svlan = int32(newrule["source_vlan"].(int))
				} else if stype == "interface" {
					oldrule.Match.SInterface = newrule["source_interface"].(string)
				} else if stype == "ipaddress" {
					oldrule.Match.Sip = newrule["source_ipaddress"].(string)
					if newrule["source_cidr_prefix"].(int) > 0 {
						iptmp := net.ParseIP("255.255.255.255")
						mask := net.CIDRMask(newrule["source_cidr_prefix"].(int), 32)
						oldrule.Match.Ssm = iptmp.Mask(mask).String()
						oldrule.Match.SRuleType = "prefix"
					} else if newrule["source_subnet_mask"].(string) != "" {
						oldrule.Match.Ssm = newrule["source_subnet_mask"].(string)
						oldrule.Match.SRuleType = "netmask"
					} else if newrule["source_wildcard_mask"].(string) != "" {
						maskb := net.ParseIP(newrule["source_wildcard_mask"].(string))
						for i := 12; i < 16; i++ {
							maskb[i] = maskb[i] ^ 0xff
						}
						oldrule.Match.Ssm = maskb.String()
						oldrule.Match.SRuleType = "wildcard"
					}
				} else if stype == "macaddress" {
					oldrule.Match.Smac = newrule["source_mac"].(string)
				} else {
					return diag.Errorf("[ERROR] Unknown source type : %s", stype)
				}

			}

			dtype := newrule["destination_type"].(string)

			oldrule.Match.DAddressGroup = ""
			oldrule.Match.DPortGroup = ""
			oldrule.Match.DportLow = -1
			oldrule.Match.DportHigh = -1
			oldrule.Match.Dvlan = -1
			oldrule.Match.DInterface = ""
			oldrule.Match.Dip = "any"
			oldrule.Match.DRuleType = "prefix"
			oldrule.Match.Dsm = "255.255.255.255"
			oldrule.Match.Dmac = ""
			if dtype == "any" {
			} else if dtype == "objectgroup" {
				oldrule.Match.DAddressGroup = newrule["destination_address_group_id"].(string)
				oldrule.Match.DPortGroup = newrule["destination_port_group_id"].(string)
			} else {
				ports := strings.Split(newrule["destination_port"].(string), "-")
				plen := len(ports)
				if plen == 1 && ports[0] == "" {
				} else if plen == 1 {
					v, err := strconv.Atoi(ports[0])
					if err != nil {
						return diag.FromErr(err)
					}
					oldrule.Match.DportLow = int32(v)
					oldrule.Match.DportHigh = int32(v)
				} else if plen == 2 {
					v, err := strconv.Atoi(ports[0])
					if err != nil {
						return diag.FromErr(err)
					}
					oldrule.Match.DportLow = int32(v)
					v, err = strconv.Atoi(ports[1])
					if err != nil {
						return diag.FromErr(err)
					}
					oldrule.Match.DportHigh = int32(v)
				} else {
					return diag.Errorf("[ERROR] source_port is error.[%s]", newrule["destination_port"].(string))
				}

				if dtype == "none" {
				} else if dtype == "vlan" {
					oldrule.Match.Dvlan = int32(newrule["destination_vlan"].(int))
				} else if dtype == "interface" {
					oldrule.Match.DInterface = newrule["destination_interface"].(string)
				} else if dtype == "ipaddress" {
					oldrule.Match.Dip = newrule["destination_ipaddress"].(string)
					if newrule["destination_cidr_prefix"].(int) > 0 {
						iptmp := net.ParseIP("255.255.255.255")
						mask := net.CIDRMask(newrule["destination_cidr_prefix"].(int), 32)
						oldrule.Match.Dsm = iptmp.Mask(mask).String()
						oldrule.Match.DRuleType = "prefix"
					} else if newrule["destination_subnet_mask"].(string) != "" {
						oldrule.Match.Dsm = newrule["destination_subnet_mask"].(string)
						oldrule.Match.DRuleType = "netmask"
					} else if newrule["destination_wildcard_mask"].(string) != "" {
						maskb := net.ParseIP(newrule["destination_wildcard_mask"].(string))
						for i := 12; i < 16; i++ {
							maskb[i] = maskb[i] ^ 0xff
						}
						oldrule.Match.Dsm = maskb.String()
						oldrule.Match.DRuleType = "wildcard"
					}
				} else if dtype == "macaddress" {
					oldrule.Match.Dmac = newrule["destination_mac"].(string)
				} else {
					return diag.Errorf("[ERROR] Unknown destination type : %s", dtype)
				}

			}

			oldrule.Match.Appid = int32(newrule["application_id"].(int))
			oldrule.Match.Dscp = int32(newrule["dscp"].(int))
			oldrule.LoggingEnabled = newrule["logging"].(bool)
			oldrule.Action.AllowOrDeny = newrule["action"].(string)
		}
	}
	// Inbound ACLs
	if isEdge {
		firewallData.Inbound = nil
		for _, protForwardRule := range d.Get("port_forwarding_rule").([]interface{}) {
			tmp_inbound := vcoclient.FirewallInboundRule{}
			rule := protForwardRule.(map[string]interface{})

			tmp_inbound.Action.SubinterfaceId = -1
			tmp_inbound.Action.Type = "port_forwarding"
			tmp_inbound.Match.Appid = -1
			tmp_inbound.Match.Classid = -1
			tmp_inbound.Match.Dscp = -1
			tmp_inbound.Match.Dsm = "any"
			tmp_inbound.Match.Dvlan = -1
			tmp_inbound.Match.Hostname = ""
			tmp_inbound.Match.OsVersion = -1
			tmp_inbound.Match.SportHigh = -1
			tmp_inbound.Match.SportLow = -1
			tmp_inbound.Match.Svlan = -1

			tmp_inbound.Name = rule["name"].(string)
			protocol := rule["protocol"].(string)
			if protocol == "tcp" {
				tmp_inbound.Match.Proto = 6
			} else if protocol == "udp" {
				tmp_inbound.Match.Proto = 17
			}
			tmp_inbound.Action.Interface = rule["interface"].(string)
			if rule["outside_ip"].(string) == "" {
				tmp_inbound.Match.Dip = "any"
			} else {
				tmp_inbound.Match.Dip = rule["outside_ip"].(string)
			}
			ports := strings.Split(rule["wan_ports"].(string), "-")
			plen := len(ports)
			if plen == 1 && ports[0] == "" {
			} else if plen == 1 {
				v, err := strconv.Atoi(ports[0])
				if err != nil {
					return diag.FromErr(err)
				}
				tmp_inbound.Match.DportLow = int32(v)
				tmp_inbound.Match.DportHigh = int32(v)
			} else if plen == 2 {
				v, err := strconv.Atoi(ports[0])
				if err != nil {
					return diag.FromErr(err)
				}
				tmp_inbound.Match.DportLow = int32(v)
				v, err = strconv.Atoi(ports[1])
				if err != nil {
					return diag.FromErr(err)
				}
				tmp_inbound.Match.DportHigh = int32(v)
			} else {
				return diag.Errorf("[ERROR] wan_ports is error.[%s]", rule["wan_port"].(string))
			}
			tmp_inbound.Action.Nat.LanIp = rule["lan_ip"].(string)
			tmp_inbound.Action.Nat.LanPort = rule["lan_port"].(int)
			tmp_inbound.Action.SegmentId = int32(rule["segment_id"].(int))
			if rule["remote_ip_subnet"].(string) == "" {
				tmp_inbound.Match.Sip = "any"
				tmp_inbound.Match.Ssm = "any"
			} else {
				t := strings.Split(rule["remote_ip_subnet"].(string), "/")
				if len(t) != 2 {
					return diag.Errorf("[ERROR] remote_ip_subent is error.[%s]", rule["remote_ip_subnet"].(string))
				}
				tmp_inbound.Match.Sip = t[0]
				mlen, err := strconv.Atoi(t[1])
				if err != nil {
					return diag.Errorf("[ERROR] remote_ip_subent is error.[%s]", rule["remote_ip_subnet"].(string))
				}
				iptmp := net.ParseIP("255.255.255.255")
				mask := net.CIDRMask(mlen, 32)
				tmp_inbound.Match.Ssm = iptmp.Mask(mask).String()
			}
			firewallData.Inbound = append(firewallData.Inbound, tmp_inbound)
		}

		for _, protForwardRule := range d.Get("nat_rule").([]interface{}) {
			tmp_inbound := vcoclient.FirewallInboundRule{}
			rule := protForwardRule.(map[string]interface{})

			tmp_inbound.Action.Nat.LanPort = -1
			tmp_inbound.Action.SubinterfaceId = -1
			tmp_inbound.Action.Type = "one_to_one_nat"
			tmp_inbound.Match.Appid = -1
			tmp_inbound.Match.Classid = -1
			tmp_inbound.Match.Dscp = -1
			tmp_inbound.Match.Dsm = "any"
			tmp_inbound.Match.Dvlan = -1
			tmp_inbound.Match.Hostname = ""
			tmp_inbound.Match.OsVersion = -1
			tmp_inbound.Match.SportHigh = -1
			tmp_inbound.Match.SportLow = -1
			tmp_inbound.Match.Svlan = -1

			tmp_inbound.Name = rule["name"].(string)
			tmp_inbound.Match.Dip = rule["outbound_ip"].(string)
			tmp_inbound.Action.Interface = rule["interface"].(string)
			tmp_inbound.Action.Nat.LanIp = rule["inside_ip"].(string)
			tmp_inbound.Action.SegmentId = int32(rule["segment_id"].(int))
			tmp_inbound.Action.Nat.Outbound = rule["outbound_traffic"].(bool)
			protocol := rule["allow_protocol"].(string)
			if protocol == "icmp" {
				tmp_inbound.Match.Proto = 1
			} else if protocol == "tcp" {
				tmp_inbound.Match.Proto = 6
			} else if protocol == "udp" {
				tmp_inbound.Match.Proto = 17
			} else if protocol == "gre" {
				tmp_inbound.Match.Proto = 47
			} else if protocol == "all" {
				tmp_inbound.Match.Proto = -1
			}
			if rule["allow_ports"].(string) == "" {
				tmp_inbound.Match.DportLow = -1
				tmp_inbound.Match.DportHigh = -1
			} else {
				ports := strings.Split(rule["allow_ports"].(string), "-")
				plen := len(ports)
				if plen == 1 {
					v, err := strconv.Atoi(ports[0])
					if err != nil {
						return diag.FromErr(err)
					}
					tmp_inbound.Match.DportLow = int32(v)
					tmp_inbound.Match.DportHigh = int32(v)
				} else if plen == 2 {
					v, err := strconv.Atoi(ports[0])
					if err != nil {
						return diag.FromErr(err)
					}
					tmp_inbound.Match.DportLow = int32(v)
					v, err = strconv.Atoi(ports[1])
					if err != nil {
						return diag.FromErr(err)
					}
					tmp_inbound.Match.DportHigh = int32(v)
				} else {
					return diag.Errorf("[ERROR] allow_ports is error.[%s]", rule["allow_port"].(string))
				}
			}
			if rule["remote_ip_subnet"].(string) == "" {
				tmp_inbound.Match.Sip = "any"
				tmp_inbound.Match.Ssm = "any"
			} else {
				t := strings.Split(rule["remote_ip_subnet"].(string), "/")
				if len(t) != 2 {
					return diag.Errorf("[ERROR] remote_ip_subent is error.[%s]", rule["remote_ip_subnet"].(string))
				}
				tmp_inbound.Match.Sip = t[0]
				mlen, err := strconv.Atoi(t[1])
				if err != nil {
					return diag.Errorf("[ERROR] remote_ip_subent is error.[%s]", rule["remote_ip_subnet"].(string))
				}
				iptmp := net.ParseIP("255.255.255.255")
				mask := net.CIDRMask(mlen, 32)
				tmp_inbound.Match.Ssm = iptmp.Mask(mask).String()
			}
			firewallData.Inbound = append(firewallData.Inbound, tmp_inbound)
		}
	}
	// Stateful Firewall Settings
	if len(d.Get("stateful_firewall").([]interface{})) == 1 {
		if isEdge && !d.Get("stateful_firewall.0.edge_overwrite").(bool) {
			firewallData.StatefulFirewallSettings = nil
		} else {
			firewallData.StatefulFirewallSettings = &vcoclient.FirewallStatefulFirewallSettings{}
			firewallData.StatefulFirewallSettings.EstablishedTcpFlowTimeout = d.Get("stateful_firewall.0.establieshd_tcp_timeout").(int)
			firewallData.StatefulFirewallSettings.NonEstablishedTcpFlowTimeout = d.Get("stateful_firewall.0.non_established_tcp_timeout").(int)
			firewallData.StatefulFirewallSettings.UdpFlowTimeout = d.Get("stateful_firewall.0.udp_timeout").(int)
			firewallData.StatefulFirewallSettings.OtherFlowTimeout = d.Get("stateful_firewall.0.other_timeout").(int)
		}
	}
	// Network & Flood Protection Settings
	if len(d.Get("network_flood_protection").([]interface{})) == 1 {
		if isEdge && !d.Get("network_flood_protection.0.edge_overwrite").(bool) {
			firewallData.NetworkProtectionSettings = nil
		} else {
			firewallData.NetworkProtectionSettings = &vcoclient.FirewallNetworkProtectionSettings{}
			firewallData.NetworkProtectionSettings.NewConnectionThreshold = d.Get("network_flood_protection.0.new_connection_threshold").(int)
			if d.Get("network_flood_protection.0.denylist").(bool) {
				firewallData.NetworkProtectionSettings.Denylist = true
				firewallData.NetworkProtectionSettings.DetectionTime = d.Get("network_flood_protection.0.detection_time").(int)
				firewallData.NetworkProtectionSettings.DenylistDuration = d.Get("network_flood_protection.0.denylist_time").(int)
			} else {
				firewallData.NetworkProtectionSettings.Denylist = false
			}
			firewallData.NetworkProtectionSettings.TcpBasedAttacks.InvalidFlags = d.Get("network_flood_protection.0.invalid_tcp_flags").(bool)
			firewallData.NetworkProtectionSettings.TcpBasedAttacks.EnableLand = d.Get("network_flood_protection.0.tcp_land").(bool)
			firewallData.NetworkProtectionSettings.TcpBasedAttacks.EnableSynFragment = d.Get("network_flood_protection.0.tcp_syn_fragment").(bool)
			firewallData.NetworkProtectionSettings.IcmpBasedAttacks.EnablePingOfDeath = d.Get("network_flood_protection.0.icmp_ping_of_death").(bool)
			firewallData.NetworkProtectionSettings.IcmpBasedAttacks.EnableFragment = d.Get("network_flood_protection.0.icmp_fragment").(bool)
			firewallData.NetworkProtectionSettings.IpBasedAttacks.EnableUnknownProtocol = d.Get("network_flood_protection.0.ip_unkown_protocol").(bool)
			firewallData.NetworkProtectionSettings.IpBasedAttacks.EnableInsecureOptions = d.Get("network_flood_protection.0.ip_options").(bool)
		}
	}

	// Edge Access
	if len(d.Get("edge_access").([]interface{})) == 1 {
		if isEdge && !d.Get("edge_access.0.edge_overwrite").(bool) {
			firewallData.Services = nil
		} else {
			firewallData.Services = &vcoclient.FirewallServices{}
			firewallData.Services.Ssh.AllowSelectedIp = []string{}
			firewallData.Services.LocalUi.AllowSelectedIp = []string{}
			firewallData.Services.Snmp.AllowSelectedIp = []string{}
			firewallData.Services.Ssh.Enabled = d.Get("edge_access.0.ssh").(bool)
			for _, v := range d.Get("edge_access.0.ssh_allow").([]interface{}) {
				firewallData.Services.Ssh.AllowSelectedIp = append(firewallData.Services.Ssh.AllowSelectedIp, v.(string))
			}
			firewallData.Services.Console.Enabled = d.Get("edge_access.0.console").(bool)
			firewallData.Services.UsbDisabled = d.Get("edge_access.0.usb").(bool)
			firewallData.Services.Snmp.Enabled = d.Get("edge_access.0.snmp").(bool)
			for _, v := range d.Get("edge_access.0.snmp_allow").([]interface{}) {
				firewallData.Services.Snmp.AllowSelectedIp = append(firewallData.Services.Snmp.AllowSelectedIp, v.(string))
			}
			firewallData.Services.LocalUi.Enabled = d.Get("edge_access.0.webui").(bool)
			for _, v := range d.Get("edge_access.0.webui_allow").([]interface{}) {
				firewallData.Services.LocalUi.AllowSelectedIp = append(firewallData.Services.LocalUi.AllowSelectedIp, v.(string))
			}
			firewallData.Services.LocalUi.PortNumber = d.Get("edge_access.0.webui_port").(int)
		}
	}
	// post
	post2 := &vcoclient.ConfigurationUpdateConfigurationModule{
		ConfigurationModuleId: res1[0].Id,
		Update: *&vcoclient.ConfigurationModule{
			Data:        firewallData,
			Name:        res1[0].Name,
			Description: res1[0].Description,
		},
	}
	_, _, err = conn.ConfigurationUpdateConfigurationModule(nil, *post2)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceVeloFirewallRead(ctx, d, m)
}

func resourceVeloFirewallDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// this module is no delete action
	return diags
}
