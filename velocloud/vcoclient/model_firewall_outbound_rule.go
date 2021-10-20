package vcoclient

type FirewallOutboundRule struct {
	Name           string                     `json:"name,omitempty"`
	Match          FirewallRuleMatch         `json:"match"`
	Action         FirewallOutboundRuleAction `json:"action"`
	RuleLogicalId  string                     `json:"ruleLogicalId,omitempty"`
	LoggingEnabled bool                       `json:"loggingEnabled,omitempty"`
	Comments       []interface{}              `json:"comments,omitempty"`
}
