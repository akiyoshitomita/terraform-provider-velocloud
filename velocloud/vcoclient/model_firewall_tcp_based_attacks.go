package vcoclient

type FirewallTcpBasedAttacks struct {
	InvalidFlags      bool `json:"invalidFlags"`
	EnableLand        bool `json:"enableLand"`
	EnableSynFragment bool `json:"enableSynFragment"`
}
