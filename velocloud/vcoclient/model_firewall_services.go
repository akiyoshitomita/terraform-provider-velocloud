package vcoclient

type FirewallServices struct {
	Ssh     FirewallServiceSsh     `json:"ssh,omitempty"`
	LocalUi FirewallServiceLocalUi `json:"localUi,omitempty"`
	Console FirewallServiceConsole `json:"console,omitempty"`
	Snmp    FirewallServiceSnmp    `json:"snmp,omitempty"`
	Icmp    FirewallServiceIcmp    `json:icmp,omitempty"`
}

type FirewallServiceSsh struct {
	Enabled         bool     `json:"enabled,omitempty"`
	AllowSelectedIp []string `json:"allowSelectedIp,omitempty"`
	RuleLogicalId   string   `json:"ruleLogicalId,omitempty"`
}

type FirewallServiceLocalUi struct {
	Enabled         bool     `json:"enabled,omitempty"`
	AllowSelectedIp []string `json:"allowSelectedIp,omitempty"`
	PortNumber      int      `json:"portNumber,omitempty"`
	RuleLogicalId   string   `json:"ruleLogicalId,omitempty"`
}

type FirewallServiceConsole struct {
	Enabled bool `json:"enabled,omitempty"`
}

type FirewallServiceSnmp struct {
	Enabled         bool     `json:"enabled,omitempty"`
	AllowSelectedIp []string `json:"allowSelectedIp,omitempty"`
	RuleLogicalId   string   `json:"ruleLogicalId,omitempty"`
}

type FirewallServiceIcmp struct {
	Enabled         bool     `json:"enabled,omitempty"`
	AllowSelectedIp []string `json:"allowSelectedIp,omitempty"`
	RuleLogicalId   string   `json:"ruleLogicalId,omitempty"`
}
