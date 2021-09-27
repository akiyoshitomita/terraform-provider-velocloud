package vcoclient

type EdgeEdgeProvisionResult struct {
	Id                   int32       `json:"id"`
	ActivationKey        string      `json:"activationKey"`
	GeneratedCertificate interface{} `json:"generatedCertificate,omitempty"`
	Token                interface{} `json:"token,omitempty"`
}
