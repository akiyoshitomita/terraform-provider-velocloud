package vcoclient

import (
	"time"
)

type Link struct {
	Id          int32     `json:"id"`
	Created     time.Time `json:"created"`
	EdgeId      int32     `json:"edgeId"`
	LogicalId   string    `json:"logicalId"`
	InternalId  string    `json:"internalId"`
	Interface_  string    `json:"interface"`
	MacAddress  string    `json:"macAddress"`
	IpAddress   string    `json:"ipAddress"`
	IpV6Address string    `json:"ipV6Address,omitempty"`
	Netmask     string    `json:"netmask"`
	NetworkSide string    `json:"networkSide"`
	NetworkType string    `json:"networkType"`
	DisplayName string    `json:"displayName"`
	Isp         string    `json:"isp"`
	Org         string    `json:"org"`
	Lat         float64   `json:"lat"`
	Lon         float64   `json:"lon"`
	LastActive  time.Time `json:"lastActive"`
	State       string    `json:"state"`
	BackupState string    `json:"backupState"`
	LinkMode    string    `json:"linkMode,omitempty"`
	// *Deprecated* - Do not use
	VpnState              string             `json:"vpnState"`
	LastEvent             time.Time          `json:"lastEvent"`
	LastEventState        string             `json:"lastEventState"`
	AlertsEnabled         *Tinyint           `json:"alertsEnabled"`
	OperatorAlertsEnabled *Tinyint           `json:"operatorAlertsEnabled"`
	ServiceState          string             `json:"serviceState"`
	Modified              time.Time          `json:"modified"`
	ServiceGroups         *LinkServiceGroups `json:"serviceGroups,omitempty"`
}
