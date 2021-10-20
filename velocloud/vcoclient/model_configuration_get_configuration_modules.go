package vcoclient

type ConfigurationGetConfigurationModules struct {
	ConfigurationId int      `json:"configurationId,omitempty"`
	EnterpriseId    int      `json:"enterpriseId,omitempty"`
	NoData          bool     `json:"noData,omitempty"`
	Modules         []string `json:"modules,omitempty"`
}
