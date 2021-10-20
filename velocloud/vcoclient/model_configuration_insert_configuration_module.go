package vcoclient

type ConfigurationInsertConfigurationModule struct {
	EnterpriseId    int         `json:"enterpriseId,omitempty"`
	Name            string      `json:"name,omitempty"`
	Type            string      `json:"type,omitempty"`
	Description     string      `json:"description,omitempty"`
	Data            interface{} `json:"data,omitempty"`
	ConfigurationId int32       `json:"configurationId,omitempty"`
	Version         string      `json:"version,omitempty"`
}
