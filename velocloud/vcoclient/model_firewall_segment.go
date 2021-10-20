package vcoclient

type FirewallSegment struct {
	FirewallLoggingEnabled  bool                                `json:"firewall_logging_enabled"`
	StatefulFirewallEnabled bool                                `json:"stateful_firewall_enabled,omitempty"`
	Outbound                []FirewallOutboundRule              `json:"outbound"`
	Segment                 *ConfigurationModuleSegmentMetadata `json:"segment"`
}
