package vcoclient

type ConfigurationUpdateConfigurationModule struct {
	ConfigurationModuleId int                 `json:"configurationModuleId"`
	EnterpriseId          int                 `json:"enterpriseId,omitempty"`
	Basic                 bool                `json:"basic,omitempty"`
	Update                ConfigurationModule `json:"_update,omitempty"`
}
