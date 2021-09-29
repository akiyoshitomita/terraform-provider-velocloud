package vcoclient

import (
	"time"
)

type EnterpriseGetEnterpriseConfigurationsResultItem struct {
	ConfigurationType string                `json:"configurationType,omitempty"`
	BastionState      string                `json:"bastionState,omitempty"`
	Created           time.Time             `json:"created,omitempty"`
	Description       string                `json:"description,omitempty"`
	EdgeCount         int32                 `json:"edgeCount,omitempty"`
	Effective         time.Time             `json:"effective,omitempty"`
	Id                int32                 `json:"id,omitempty"`
	LogicalId         string                `json:"logicalId,omitempty"`
	Modified          time.Time             `json:"modified,omitempty"`
	Modules           []ConfigurationModule `json:"modules,omitempty"`
	Name              string                `json:"name,omitempty"`
	SchemaVersion     string                `json:"schemaVersion,omitempty"`
	Version           string                `json:"version,omitempty"`
	IsStaging         *Tinyint              `json:"isStaging,omitempty"`
}
