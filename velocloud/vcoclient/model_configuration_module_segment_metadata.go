package vcoclient

type ConfigurationModuleSegmentMetadata struct {
	Name             string `json:"name,omitempty"`
	SegmentId        int32  `json:"segmentId"`
	SegmentLogicalId string `json:"segmentLogicalId,omitempty"`
	Type             string `json:"type,omitempty"`
}
