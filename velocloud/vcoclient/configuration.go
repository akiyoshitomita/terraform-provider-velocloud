/*
 * VMware SD-WAN Orchestrator API v1
 *
 * API version: 4.4.0
 */

package vcoclient

import (
	"net/http"
)

type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	Idcount       int
	HTTPClient    *http.Client
}
