package vcoclient

type FirewallRuleMatch struct {
	// Integer ID corresponding to an application in the network-level application map
	Appid int32 `json:"appid,omitempty"`
	// Integer ID corresponding to an application class in the network-level application map
	Classid int32 `json:"classid,omitempty"`
	// Integer ID indicating DSCP classification, where mappings are as follows: [EF:46,VA:44,AF11:10,AF12:12,AF13:14,AF21:18,AF22:20,AF23:22,AF31:26,AF32:28,AF33:30,AF41:34,AF42:36,AF43:38,CS0:0,CS1:8,CS2:16,CS3:24,CS4:32,CS5:40,CS6:48,CS7:56]
	Dscp int32 `json:"dscp"`
	// Source IP address
	Sip string `json:"sip,omitempty"`
	// Upper bound of a source port range
	SportHigh int32 `json:"sport_high,omitempty"`
	// Lower bound of a source port range
	SportLow int32 `json:"sport_low,omitempty"`
	// Source address group reference
	SAddressGroup string `json:"sAddressGroup,omitempty"`
	// Source port group reference
	SPortGroup string `json:"sPortGroup,omitempty"`
	// Source subnet mask, e.g. 255.255.255.0
	Ssm string `json:"ssm,omitempty"`
	// Source MAC address
	Smac string `json:"smac,omitempty"`
	// Integer ID for the source VLAN
	Svlan      int32  `json:"svlan,omitempty"`
	SInterface string `json:"sInterface"`
	// Index corresponding to the OS in the array: [OTHER,WINDOWS,LINUX,MACOS,IOS,ANDROID,EDGE]
	OsVersion int32  `json:"os_version,omitempty"`
	Hostname  string `json:"hostname"`
	// Destination IP address
	Dip string `json:"dip,omitempty"`
	// Lower bound of a destination port range
	DportLow int32 `json:"dport_low,omitempty"`
	// Upper bound of a destination port range
	DportHigh int32 `json:"dport_high,omitempty"`
	// Destination address group reference
	DAddressGroup string `json:"dAddressGroup,omitempty"`
	// Destination port group reference
	DPortGroup string `json:"dPortGroup,omitempty"`
	// Destination subnet mask e.g. 255.255.255.0
	Dsm string `json:"dsm,omitempty"`
	// Destination MAC address
	Dmac string `json:"dmac,omitempty"`
	// Integer ID for the destination VLAN
	Dvlan      int32  `json:"dvlan,omitempty"`
	DInterface string `json:"dInterface"`
	// Integer ID corresponding to a protocol
	Proto int32 `json:"proto,omitempty"`
	// Source rule type
	SRuleType string `json:"s_rule_type,omitempty"`
	// Destination rule type
	DRuleType string `json:"d_rule_type,omitempty"`
}
