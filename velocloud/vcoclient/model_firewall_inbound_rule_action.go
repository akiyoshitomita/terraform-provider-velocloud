package vcoclient

type ModelFirewallInboundRuleAction struct {
	Type           string                           `json:"type,omitempty"`
	Nat            ModeFirewallInboundRuleActionNat `json:"nat,omitempty"`
	Interface      string                           `json:"interface,omitempty"`
	SubinterfaceId int                              `json:"subinterfaceId,omitempty"`
}

type ModeFirewallInboundRuleActionNat struct {
	LanIp    string `json:"lan_ip,omitempty"`
	LanPort  int    `json:"lan_port,omitempty"`
	OutBound bool   `json:"outbound,omitempty"`
}
