package vcoclient

type ObjectGroupData struct {
	// IP address
	Ip string `json:"ip,omitempty"`
	// Netmask, e.g. 255.255.255.0
	Mask string `json:"mask,omitempty"`
	// Rule type
	RuleType string `json:"rule_type,omitempty"`
	// Domain name
	Domain string `json:"domain,omitempty"`
	// Integer ID corresponding to a protocol
	Proto int32 `json:"proto,omitempty"`
	// Lower bound of a port range
	PortLow int32 `json:"port_low,omitempty"`
	// Upper bound of a port range
	PortHigh int32 `json:"port_high,omitempty"`
}
