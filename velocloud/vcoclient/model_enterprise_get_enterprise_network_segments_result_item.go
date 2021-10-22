package vcoclient

import (
	"time"
)

type EnterpriseGetEnterpriseNetworkSegmentsResultItem struct {
	Id                    int32               `json:"id"`
	Created               time.Time           `json:"created"`
	OperatorId            int32               `json:"operatorId"`
	NetworkId             int32               `json:"networkId"`
	EnterpriseId          int32               `json:"enterpriseId"`
	EdgeId                int32               `json:"edgeId"`
	GatewayId             int32               `json:"gatewayId"`
	ParentGroupId         int32               `json:"parentGroupId"`
	Description           string              `json:"description"`
	Object                string              `json:"object"`
	Name                  string              `json:"name"`
	Type_                 string              `json:"type"`
	LogicalId             string              `json:"logicalId"`
	AlertsEnabled         *Tinyint            `json:"alertsEnabled"`
	OperatorAlertsEnabled *Tinyint            `json:"operatorAlertsEnabled"`
	Status                string              `json:"status"`
	StatusModified        time.Time           `json:"statusModified"`
	PreviousData          interface{}         `json:"previousData"`
	PreviousCreated       time.Time           `json:"previousCreated"`
	DraftData             string              `json:"draftData"`
	DraftCreated          time.Time           `json:"draftCreated"`
	DraftComment          string              `json:"draftComment"`
	Data                  *NetworkSegmentData `json:"data"`
	LastContact           time.Time           `json:"lastContact"`
	Version               string              `json:"version"`
	Modified              time.Time           `json:"modified"`
	ProfileCount          int32               `json:"profileCount,omitempty"`
	HandoffUsage          []interface{}       `json:"handoffUsage,omitempty"`
	IsAllMPGCDE           bool                `json:"isAllMPGCDE,omitempty"`
	IsAllCtrlCDE          bool                `json:"isAllCtrlCDE,omitempty"`
	EdgeCount             int32               `json:"edgeCount,omitempty"`
	EdgeUsage             []interface{}       `json:"edgeUsage,omitempty"`
	Configuration         []interface{}       `json:"configuration,omitempty"`
}
