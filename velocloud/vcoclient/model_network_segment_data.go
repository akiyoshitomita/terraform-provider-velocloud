package vcoclient

type NetworkSegmentData struct {
	SegmentId                 int32 `json:"segmentId,omitempty"`
	ServiceVlan               int32 `json:"serviceVlan,omitempty"`
	DelegateToEnterprise      bool  `json:"delegateToEnterprise,omitempty"`
	DelegateToEnterpriseProxy bool  `json:"delegateToEnterpriseProxy,omitempty"`
}
