package vcoclient

type EnterpriseGetEnterpriseNetworkSegments struct {
	EnterpriseId int32    `json:"enterpriseId,omitempty"`
	Name         string   `json:"name,omitempty"`
	Type_        string   `json:"type,omitempty"`
	With         []string `json:"with,omitempty"`
}
