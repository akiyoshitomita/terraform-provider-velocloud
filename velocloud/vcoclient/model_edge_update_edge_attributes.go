package vcoclient

type EdgeUpdateEdgeAttributes struct {
	Id           int                             `json:"id,omitempty"`
	EnterpriseId int                             `json:"enterpriseId,omitempty"`
	Update       *EdgeUpdateEdgeAttributesUpdate `json:"_update"`
}

type EdgeUpdateEdgeAttributesUpdate struct {
	Site         *Site  `json:"site,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
	CustomInfo   string `json:"customInfo,omitempty"`
}
