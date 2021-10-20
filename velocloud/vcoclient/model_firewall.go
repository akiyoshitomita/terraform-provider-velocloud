package vcoclient

import (
	"time"
)

type Firewall struct {
	Created         time.Time     `json:"created,omitempty"`
	Effective       string        `json:"effective,omitempty"`
	Modified        time.Time     `json:"modified,omitempty"`
	Id              int           `json:"id,omitempty"`
	Name            string        `json:"name"`
	Type_           string        `json:"type,omitempty"`
	Description     string        `json:"description,omitempty"`
	ConfigurationId int32         `json:"configurationId,omitempty"`
	Data            *FirewallData `json:"data"`
	SchemaVersion   string        `json:"schemaVersion,omitempty"`
	Version         string        `json:"version,omitempty"`
	Metadata        interface{}   `json:"metadata,omitempty"`
	Refs            interface{}   `json:"refs,omitempty"`
}
