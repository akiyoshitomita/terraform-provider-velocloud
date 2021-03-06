package vcoclient

//import (
//	"time"
//)

type ConfigurationModule struct {
	Created         string      `json:"created,omitempty"`
	Effective       string      `json:"effective,omitempty"`
	Modified        string      `json:"modified,omitempty"`
	Id              int32       `json:"id,omitempty"`
	Name            string      `json:"name"`
	Type_           string      `json:"type,omitempty"`
	Description     string      `json:"description,omitempty"`
	ConfigurationId int32       `json:"configurationId,omitempty"`
	Data            interface{} `json:"data,omitempty"`
	SchemaVersion   string      `json:"schemaVersion,omitempty"`
	Version         string      `json:"version,omitempty"`
	Metadata        interface{} `json:"metadata,omitempty"`
	Refs            interface{} `json:"refs,omitempty"`
}
