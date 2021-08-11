package vcoclient

import (
	"time"
)

type EdgeCertificate struct {
	Id             int32     `json:"id,omitempty"`
	Created        time.Time `json:"created,omitempty"`
	CsrId          int32     `json:"csrId,omitempty"`
	EdgeId         int32     `json:"edgeId,omitempty"`
	EnterpriseId   int32     `json:"enterpriseId,omitempty"`
	Certificate    string    `json:"certificate,omitempty"`
	SerialNumber   string    `json:"serialNumber,omitempty"`
	SubjectKeyId   string    `json:"subjectKeyId,omitempty"`
	FingerPrint    string    `json:"fingerPrint,omitempty"`
	FingerPrint256 string    `json:"fingerPrint256,omitempty"`
	ValidFrom      time.Time `json:"validFrom,omitempty"`
	ValidTo        time.Time `json:"validTo,omitempty"`
}
