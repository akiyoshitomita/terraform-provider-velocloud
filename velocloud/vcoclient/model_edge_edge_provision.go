package vcoclient

type EdgeEdgeProvision struct {
	EnterpriseId    int    `json:"enterpriseId,omitempty"`
	ConfigurationId int    `json:"configurationId"`
	Name            string `json:"name,omitempty"`
	SerialNumber    string `json:"serialNumber,omitempty"`
	ModelNumber     string `json:"modelNumber"`
	Description     string `json:"description,omitempty"`
	Site            *Site  `json:"site,omitempty"`
	HaEnabled       bool   `json:"haEnabled,omitempty"`
	//GenerateCertificate bool `json:"generateCertificate,omitempty"`
	//SubjectCN string `json:"subjectCN,omitempty"`
	//SubjectO string `json:"subjectO,omitempty"`
	//SubjectOU string `json:"subjectOU,omitempty"`
	//ChallengePassword string `json:"challengePassword,omitempty"`
	//PrivateKeyPassword string `json:"privateKeyPassword,omitempty"`
	EndpointPkiMode string `json:"endpointPkiMode,omitempty"`
	EdgeLicenseId   int    `json:"edgeLicenseId,omitempty"`
	CustomInfo      string `json:"customInfo,omitempty"`
	AnalyticsMode   string `json:"analyticsMode,omitempty"`
	GenerateToken   bool   `json:"generateToken,omitempty"`
}
