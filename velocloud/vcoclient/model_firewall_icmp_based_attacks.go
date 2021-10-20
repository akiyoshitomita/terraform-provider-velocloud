package vcoclient

type FirewallIcmpBasedAttacks struct {
	EnablePingOfDeath bool `json:"enablePingOfDeath"`
	EnableFragment    bool `json:"enableFragment"`
}
