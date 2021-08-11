/*
 * VMware SD-WAN Orchestrator API v1
 *
 * API version: 4.4.0
 */

package vcoclient

type EnterpriseGetEnterpriseEdges struct {
	Id           int         `json:"id,omitempty"`
	EnterpriseId int         `json:"enterpriseId,omitempty"`
	With         []string    `json:"with,omitempty"`
	Limit        int         `json:"limit,omitempty"`
	NextPageLink string      `json:"nextPageLink,omitempty"`
	PrevPageLink string      `json:"prevPageLink,omitempty"`
	Count        bool        `json:"_count,omitempty"`
	QuickSearch  string      `json:"quickSearch,omitempty"`
	Filters      interface{} `json:"filters,omitempty"`
	FilterSpec   bool        `json:"_filterSpec,omitempty"`
	FieldNeeded  []string    `json:"fieldsNeeded,omitempty"`
}
