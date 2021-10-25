package vcoclient

type EnterpriseGetObjectGroups struct {
	EnterpriseId int32    `json:"enterpriseId,omitempty"`
	Id           int32    `json:"id,omitempty"`
	Type_        []string `json:"type,omitempty"`
}
