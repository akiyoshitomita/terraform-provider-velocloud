package vcoclient

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	//"log"
)

type EnterpriseApiService service

func (a *EnterpriseApiService) EnterpriseGetEnterpriseConfigurations(ctx context.Context, data EnterpriseGetEnterpriseConfigurations) ([]EnterpriseGetEnterpriseConfigurationsResultItem, *http.Response, error) {

	localVarHttpMethod := strings.ToUpper("Post")
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	r, err := a.client.prepareRequest(ctx, "/enterprise/getEnterpriseConfigurations", localVarHttpMethod, data, localVarHeaderParams)
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
		return nil, localVarHttpResponse, errors.New("API response error")
	}

	var localVarReturnValue []EnterpriseGetEnterpriseConfigurationsResultItem
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		return nil, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

func (a *EnterpriseApiService) EnterpriseGetEnterpriseEdges(ctx context.Context, data EnterpriseGetEnterpriseEdges) ([]EnterpriseGetEnterpriseEdgesResultItem, *http.Response, error) {

	localVarHttpMethod := strings.ToUpper("Post")
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	r, err := a.client.prepareRequest(ctx, "/enterprise/getEnterpriseEdges", localVarHttpMethod, data, localVarHeaderParams)
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
		return nil, localVarHttpResponse, errors.New("API response error")
	}

	var localVarReturnValue []EnterpriseGetEnterpriseEdgesResultItem
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		return nil, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

func (a *EnterpriseApiService) EnterpriseGetEnterpriseNetworkSegments(ctx context.Context, data EnterpriseGetEnterpriseNetworkSegments) ([]EnterpriseGetEnterpriseNetworkSegmentsResultItem, *http.Response, error) {

	localVarHttpMethod := strings.ToUpper("Post")
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	r, err := a.client.prepareRequest(ctx, "/enterprise/getEnterpriseNetworkSegments", localVarHttpMethod, data, localVarHeaderParams)
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
		return nil, localVarHttpResponse, errors.New("API response error")
	}

	var localVarReturnValue []EnterpriseGetEnterpriseNetworkSegmentsResultItem
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		return nil, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
