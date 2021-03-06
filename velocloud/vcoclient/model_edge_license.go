package vcoclient

import (
	"time"
)

type EdgeLicense struct {
	Id            int32     `json:"id,omitempty"`
	Created       time.Time `json:"created,omitempty"`
	LicenseId     int32     `json:"licenseId,omitempty"`
	Sku           string    `json:"sku,omitempty"`
	Name          string    `json:"name,omitempty"`
	Alias         string    `json:"alias,omitempty"`
	Detail        string    `json:"detail,omitempty"`
	Quota         string    `json:"quota,omitempty"`
	TermMonths    int32     `json:"termMonths,omitempty"`
	Start         time.Time `json:"start,omitempty"`
	End           time.Time `json:"end,omitempty"`
	Edition       string    `json:"edition,omitempty"`
	BandwidthTier string    `json:"bandwidthTier,omitempty"`
	Active        *Tinyint  `json:"active,omitempty"`
	Modified      time.Time `json:"modified,omitempty"`
}
