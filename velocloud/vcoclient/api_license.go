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

type LicenseApiService service

func (a *LicenseApiService) LicenseGetEnterpriseEdge(ctx context.Context, data LicenseGetEnterpriseEdgeLicenses) ([]LicenseGetEnterpriseEdgeLicensesResult, *http.Response, error) {
	//var (
	//        localVarReturnValue EdgeDeleteEdgeResultItem
	//)
	localVarHttpMethod := strings.ToUpper("Post")
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	r, err := a.client.prepareRequest(ctx, "/license/getEnterpriseEdgeLicenses", localVarHttpMethod, data, localVarHeaderParams)
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

	var localVarReturnValue []LicenseGetEnterpriseEdgeLicensesResult
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		return nil, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, nil

}
