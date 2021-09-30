package vcoclient

import (
	"context"
	"io/ioutil"
	"net/http"
	//"net/url"
	"fmt"
	"strings"
)

type ConfigurationApiService service

func (a *ConfigurationApiService) ConfigurationCloneEnterpriseTemplate(ctx context.Context, data ConfigurationCloneEnterpriseTemplate) (ConfigurationCloneEnterpriseTemplateResult, *http.Response, error) {
	var (
		localVarReturnValue ConfigurationCloneEnterpriseTemplateResult
	)
	localVarHttpMethod := strings.ToUpper("Post")
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	r, err := a.client.prepareRequest(ctx, "/configuration/cloneEnterpriseTemplate", localVarHttpMethod, data, localVarHeaderParams)
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

func (a *ConfigurationApiService) ConfigurationDeleteConfiguration(ctx context.Context, data ConfigurationDeleteConfiguration) (ConfigurationDeleteConfigurationResult, *http.Response, error) {
	var (
		localVarReturnValue ConfigurationDeleteConfigurationResult
	)
	localVarHttpMethod := strings.ToUpper("Post")
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	r, err := a.client.prepareRequest(ctx, "/configuration/deleteConfiguration", localVarHttpMethod, data, localVarHeaderParams)
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

func (a *ConfigurationApiService) ConfigurationGetConfiguration(ctx context.Context, data ConfigurationGetConfiguration) (ConfigurationGetConfigurationResult, *http.Response, error) {
	var (
		localVarReturnValue ConfigurationGetConfigurationResult
	)
	localVarHttpMethod := strings.ToUpper("Post")
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	r, err := a.client.prepareRequest(ctx, "/configuration/getConfiguration", localVarHttpMethod, data, localVarHeaderParams)
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

func (a *ConfigurationApiService) ConfigurationUpdateConfiguration(ctx context.Context, data ConfigurationUpdateConfiguration) (RowsModifiedConfirmation, *http.Response, error) {
	var (
		localVarReturnValue RowsModifiedConfirmation
	)
	localVarHttpMethod := strings.ToUpper("Post")
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	r, err := a.client.prepareRequest(ctx, "/configuration/updateConfiguration", localVarHttpMethod, data, localVarHeaderParams)
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
