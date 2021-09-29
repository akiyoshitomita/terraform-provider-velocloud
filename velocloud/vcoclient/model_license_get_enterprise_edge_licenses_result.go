package vcoclient

import (
	"time"
)

type LicenseGetEnterpriseEdgeLicensesResult struct {
	Id                 int           `json:"id,omitempty"`
	Created            time.Time     `json:"created,omitempty"`
	LogicalId          int           `json:"licenseId,omitempty"`
	Sku                string        `json:"sku,omitempty"`
	Name               string        `json:"name,omitempty"`
	Alias              string        `json:"alias,omitempty"`
	Detail             LicenseDetail `json:"detail,omitempty"`
	Quota              interface{}   `json:"quota,omitempty"`
	TermMonths         int           `json:"termMonths,omitempty"`
	Start              string        `json:"start,omitempty"`
	End                string        `json:"end,omitempty"`
	Edition            string        `json:"edition,omitempty"`
	BandwidthTier      string        `json:"bandwidthTier,omitempty"`
	Active             int           `json:"active,omitempty"`
	Modified           time.Time     `json:"modified,omitempty"`
	EdgeCount          int           `json:"edgeCount,omitempty"`
	ActivatedEdgeCount int           `json:"activatedEdgeCount,omitempty"`
}

type LicenseDetail struct {
	Regions []string      `json:"regions,omitempty"`
	AddOns  []interface{} `json:"addOns,omitempty"`
}
