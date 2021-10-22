package vcoclient

type ModelFirewallInboundRuleAction struct {
	Type           string                            `json:"type,omitempty"`
	Nat            ModelFirewallInboundRuleActionNat `json:"nat,omitempty"`
	Interface      string                            `json:"interface,omitempty"`
	SubinterfaceId int32                             `json:"subinterfaceId,omitempty"`
	SegmentId      int32                             `json:"segmentId"`
}

type ModelFirewallInboundRuleActionNat struct {
	LanIp    string `json:"lan_ip,omitempty"`
	LanPort  int    `json:"lan_port,omitempty"`
	Outbound bool   `json:"outbound,omitempty"`
}
