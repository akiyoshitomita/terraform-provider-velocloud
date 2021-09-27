package vcoclient

type EdgeGetEdge struct {
	Id            int      `json:"id,omitempty"`
	EnterpriseId  int      `json:"enterpriseId,omitempty"`
	LogicalId     string   `json:"logicalId,omitempty"`
	ActivationKey string   `json:"activationKey,omitempty"`
	Name          string   `json:"string,omitempty"`
	With          []string `json:"with,omitempty"`
}
