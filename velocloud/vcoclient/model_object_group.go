package vcoclient

import (
	"time"
)

type ObjectGroup struct {
	Id                    int32             `json:"id"`
	Created               time.Time         `json:"created"`
	OperatorId            int32             `json:"operatorId"`
	NetworkId             int32             `json:"networkId"`
	EnterpriseId          int32             `json:"enterpriseId"`
	EdgeId                int32             `json:"edgeId"`
	GatewayId             int32             `json:"gatewayId"`
	ParentGroupId         int32             `json:"parentGroupId"`
	Description           string            `json:"description"`
	Object                string            `json:"object"`
	Name                  string            `json:"name"`
	Type_                 string            `json:"type"`
	LogicalId             string            `json:"logicalId"`
	AlertsEnabled         *Tinyint          `json:"alertsEnabled"`
	OperatorAlertsEnabled *Tinyint          `json:"operatorAlertsEnabled"`
	Status                string            `json:"status"`
	StatusModified        string         `json:"statusModified"`
	PreviousData          interface{}       `json:"previousData"`
	PreviousCreated       string         `json:"previousCreated"`
	DraftData             string            `json:"draftData"`
	DraftCreated          string         `json:"draftCreated"`
	DraftComment          string            `json:"draftComment"`
	Data                  []ObjectGroupData `json:"data"`
	LastContact           string         `json:"lastContact"`
	Version               string            `json:"version"`
	Modified              time.Time         `json:"modified"`
}
