/*
 * VMware SD-WAN Orchestrator API v1
 *
 * API version: 4.4.0
 */

package vcoclient

type JsonRpcRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int         `json:"id"`
}

type JsonRpcResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Meta    interface{} `json:"meta"`
	Result  interface{} `json:"result"`
	Id      int         `json:"id"`
}
