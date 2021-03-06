package vcoclient

import (
	"time"
)

type Enterprise struct {
	Id                    int32     `json:"id,omitempty"`
	Created               time.Time `json:"created,omitempty"`
	NetworkId             int32     `json:"networkId,omitempty"`
	GatewayPoolId         int32     `json:"gatewayPoolId,omitempty"`
	AlertsEnabled         *Tinyint  `json:"alertsEnabled,omitempty"`
	OperatorAlertsEnabled *Tinyint  `json:"operatorAlertsEnabled,omitempty"`
	EndpointPkiMode       string    `json:"endpointPkiMode,omitempty"`
	Name                  string    `json:"name,omitempty"`
	Domain                string    `json:"domain,omitempty"`
	Prefix                string    `json:"prefix,omitempty"`
	LogicalId             string    `json:"logicalId,omitempty"`
	AccountNumber         string    `json:"accountNumber,omitempty"`
	Description           string    `json:"description,omitempty"`
	ContactName           string    `json:"contactName,omitempty"`
	ContactPhone          string    `json:"contactPhone,omitempty"`
	ContactMobile         string    `json:"contactMobile,omitempty"`
	ContactEmail          string    `json:"contactEmail,omitempty"`
	StreetAddress         string    `json:"streetAddress,omitempty"`
	StreetAddress2        string    `json:"streetAddress2,omitempty"`
	City                  string    `json:"city,omitempty"`
	State                 string    `json:"state,omitempty"`
	PostalCode            string    `json:"postalCode,omitempty"`
	Country               string    `json:"country,omitempty"`
	Lat                   float64   `json:"lat,omitempty"`
	Lon                   float64   `json:"lon,omitempty"`
	Timezone              string    `json:"timezone,omitempty"`
	Locale                string    `json:"locale,omitempty"`
	Modified              time.Time `json:"modified,omitempty"`
	BastionState          string    `json:"bastionState,omitempty"`
}
