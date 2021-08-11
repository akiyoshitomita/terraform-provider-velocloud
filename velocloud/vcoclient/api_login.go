/*
 * VMware SD-WAN Orchestrator API v1
 *
 * API version: 4.4.0
 */

package vcoclient

import (
	"context"
	"net/http"
	"strings"
	//"log"
)

type LoginApiService service

func (a *LoginApiService) LoginEnterpriseLogin(ctx context.Context, authorization AuthObject) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarUrl        string
	)

	//localVarPostBody := &JsonRpcRequest {
	//	Jsonrpc : "2.0",
	//	Method : "/login/enterpriseLogin",
	//	Params : authorization,
	//	Id : 1,
	//}
	localVarPostBody := authorization
	localVarUrl = a.client.cfg.BasePath + "rest/login/enterpriseLogin"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	//r, err := a.client.prepareRequest(ctx, localVarHttpMethod, localVarPostBody, localVarHeaderParams)
	//if err != nil {
	//	return nil,err
	//}
	r, err := a.client.restRequest(ctx, localVarUrl, localVarHttpMethod, localVarPostBody, localVarHeaderParams)

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse.Body.Close()
	//log.Println(localVarHttpResponse)

	return localVarHttpResponse, err
}
