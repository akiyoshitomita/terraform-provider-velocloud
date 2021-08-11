package vcoclient

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"log"
)

type EnterpriseApiService service

func (a *EnterpriseApiService) EnterpriseGetEnterpriseEdges(ctx context.Context, data EnterpriseGetEnterpriseEdges) ([]EnterpriseGetEnterpriseEdgesResultItem, *http.Response, error) {

	localVarPostBody := &JsonRpcRequest{
		Jsonrpc: "2.0",
		Method:  "/enterprise/getEnterpriseEdges",
		Params:  data,
		Id:      1,
	}

	localVarHttpMethod := strings.ToUpper("Post")
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	//log.Println(localVarPostBody)
	//log.Println(localVarHeaderParams)
	r, err := a.client.prepareRequest(ctx, localVarHttpMethod, localVarPostBody, localVarHeaderParams)
	if err != nil {
		return nil, nil, err
	}
	//r, err := a.client.prepareRequest(ctx, localVarHttpMethod, localVarPostBody, localVarHeaderParams)
        //if err != nil {
        //      return nil,err
        //}

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

	log.Printf("%s\n",localVarBody)
	//var localVarReturnValue []EnterpriseGetEnterpriseEdgesResultItem
	var localVarReturnValue JsonRpcResponse
	log.Println("------------ 9 ---------------")
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	log.Println("------------ 10 ---------------")
	log.Println(err)
	if err != nil {
		return nil, localVarHttpResponse, err
	}
	log.Println("------------ 11 ---------------")
	//return localVarReturnValue, localVarHttpResponse, nil
	return nil, localVarHttpResponse, nil
}

/*
func (a *EnterpriseApiService) EnterpriseGetEnterpriseEdges(ctx context.Context, body Body) ([]EnterpriseGetEnterpriseEdgesResultItem, *http.Response, error) {

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v []EnterpriseGetEnterpriseEdgesResultItem
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
*/
