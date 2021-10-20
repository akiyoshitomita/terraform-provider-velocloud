package vcoclient

type FirewallData struct {
	FirewallEnabled           bool                               `json:"firewall_enabled"`
	InboundLoggingEnabled     *bool                              `json:"inboundLoggingEnabled,omitempty"`
	StatefulFirewallEnabled   *bool                              `json:"stateful_firewall_enabled,omitempty"`
	FirewallLoggingEnabled    *bool                              `json:"firewall_logging_enabled,omitempty"`
	SyslogForwarding          *bool                              `json:"syslog_forwarding,omitempty"`
	Inbound                   []FirewallInboundRule              `json:"inbound,omitempty"`
	StatefulFirewallSettings  *FirewallStatefulFirewallSettings  `json:"statefulFirewallSettings,omitempty"`
	NetworkProtectionSettings *FirewallNetworkProtectionSettings `json:"networkProtectionSettings,omitempty"`
	Segments                  []FirewallSegment                  `json:"segments,omitempty"`
	Services                  *FirewallServices                  `json:"services,omitempty"`
}
