package vcoclient

type FirewallInboundRule struct {
	Name          string                         `json:"name,omitempty"`
	Match         FirewallRuleMatch              `json:"match"`
	Action        ModelFirewallInboundRuleAction `json:"action"`
	RuleLogicalId string                         `json:"ruleLogicalId,omitempty"`
}
