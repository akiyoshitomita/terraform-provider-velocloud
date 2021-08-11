/*
 * VMware SD-WAN Orchestrator API v1
 *
 * API version: 4.4.0
 */

package vcoclient

type AuthObject struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password"`
	Password2 string `json:"password2,omitempty"`
	Username  string `json:"username"`
}
