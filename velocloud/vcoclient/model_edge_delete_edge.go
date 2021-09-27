/*
 * VMware SD-WAN Orchestrator API v1
 *
 * API version: 4.4.0
 */

package vcoclient

type EdgeDeleteEdge struct {
	EnterpriseId int   `json:"enterpriseId,omitempty"`
	Id           int   `json:"id,omitempty"`
	Ids          []int `json:"ids,omitempty"`
}
