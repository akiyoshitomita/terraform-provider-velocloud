package vcoclient

type ConfigurationUpdateConfiguration struct {
	Id           int                                     `json:"id"`
	EnterpriseId int                                     `json:"enterpriseId,omitempty"`
	Update       *ConfigurationUpdateConfigurationUpdate `json:"_update,omitempty"`
}

type ConfigurationUpdateConfigurationUpdate struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
	Effective   string `json:"effective,omitempty"`
}
