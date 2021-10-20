package vcoclient

type FirewallNetworkProtectionSettings struct {
	DenylistDuration        int                      `json:"denylistDuration,omitempty"`
	NewConnectionThreshold  int                      `json:"newConnectionThreshold,omitempty"`
	Denylist                bool                     `json:"denylist,omitempty"`
	DetectionTime           int                      `json:"detectionTime,omitempty"`
	TcpBasedAttacksEnabled  bool                     `json:"tcpBasedAttacksEnabled,omitempty"`
	TcpBasedAttacks         FirewallTcpBasedAttacks  `json:"tcpBasedAttacks,omitempty"`
	IcmpBasedAttacksEnabled bool                     `json:"icmpBasedAttacksEnabled,omitempty"`
	IcmpBasedAttacks        FirewallIcmpBasedAttacks `json:"icmpBasedAttacks,omitempty"`
	IpBasedAttacksEnabled   bool                     `json:"ipBasedAttacksEnabled,omitempty"`
	IpBasedAttacks          FirewallIpBasedAttacks   `json:"ipBasedAttacks,omitempty"`
}
