package vcoclient

type FirewallInboundRule struct {
	Name          string                            `json:"name,omitempty"`
	Match         *FirewallRuleMatch                `json:"match"`
	Action        *ModeFirewallInboundRuleActionNat `json:"action"`
	RuleLogicalId string                            `json:"ruleLogicalId,omitempty"`
}
