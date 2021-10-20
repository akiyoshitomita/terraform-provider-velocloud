package vcoclient

import (
	"context"
	//"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	//"log"
)

type EdgeApiService service

func (a *EdgeApiService) EdgeDeleteEdge(ctx context.Context, data EdgeDeleteEdge) (EdgeDeleteEdgeResultItem, *http.Response, error) {
	var (
		localVarReturnValue EdgeDeleteEdgeResultItem
	)
	localVarHttpMethod := strings.ToUpper("Post")
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	r, err := a.client.prepareRequest(ctx, "/edge/deleteEdge", localVarHttpMethod, data, localVarHeaderParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode != 200 {
		return localVarReturnValue, localVarHttpResponse, fmt.Errorf("ERROR velocloud api access error [%d]", localVarHttpResponse.StatusCode)
	}

	//var localVarReturnValue []EdgeDeleteEdgeResultItem
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, nil

}

func (a *EdgeApiService) EdgeEdgeProvision(ctx context.Context, data EdgeEdgeProvision) (EdgeEdgeProvisionResult, *http.Response, error) {
	var (
		localVarReturnValue EdgeEdgeProvisionResult
	)

	localVarHttpMethod := strings.ToUpper("Post")
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	r, err := a.client.prepareRequest(ctx, "/edge/edgeProvision", localVarHttpMethod, data, localVarHeaderParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode != 200 {
		return localVarReturnValue, localVarHttpResponse, fmt.Errorf("ERROR velocloud api access error [%d]", localVarHttpResponse.StatusCode)
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, nil

}

func (a *EdgeApiService) EdgeGetEdge(ctx context.Context, data EdgeGetEdge) (EdgeGetEdgeResult, *http.Response, error) {
	var (
		localVarReturnValue EdgeGetEdgeResult
	)

	localVarHttpMethod := strings.ToUpper("Post")
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	r, err := a.client.prepareRequest(ctx, "/edge/getEdge", localVarHttpMethod, data, localVarHeaderParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode != 200 {
		return localVarReturnValue, localVarHttpResponse, fmt.Errorf("ERROR velocloud api access error [%d]", localVarHttpResponse.StatusCode)
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

func (a *EdgeApiService) EdgeGetEdgeConfigurationStack(ctx context.Context, data EdgeGetEdgeConfigurationStack) ([]EdgeGetEdgeConfigurationStackResultItem, *http.Response, error) {

	localVarHttpMethod := strings.ToUpper("Post")
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	r, err := a.client.prepareRequest(ctx, "/edge/getEdgeConfigurationStack", localVarHttpMethod, data, localVarHeaderParams)
	if err != nil {
		return nil, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil {
		return nil, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return nil, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode != 200 {
		return nil, localVarHttpResponse, fmt.Errorf("ERROR velocloud api access error [%d]", localVarHttpResponse.StatusCode)
	}

	var localVarReturnValue []EdgeGetEdgeConfigurationStackResultItem
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		return nil, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

func (a *EdgeApiService) EdgeUpdateEdgeAttributes(ctx context.Context, data EdgeUpdateEdgeAttributes) (EdgeUpdateEdgeAttributesResult, *http.Response, error) {
	var (
		localVarReturnValue EdgeUpdateEdgeAttributesResult
	)

	localVarHttpMethod := strings.ToUpper("Post")
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	r, err := a.client.prepareRequest(ctx, "/edge/updateEdgeAttributes", localVarHttpMethod, data, localVarHeaderParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode != 200 {
		return localVarReturnValue, localVarHttpResponse, fmt.Errorf("ERROR velocloud api access error [%d]", localVarHttpResponse.StatusCode)
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
