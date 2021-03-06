/*
 * VMware SD-WAN Orchestrator API v1
 *
 * API version: 4.4.0
 */

package vcoclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"strings"
	//"fmt"
	"log"
)

type APIClient struct {
	cfg    *Configuration
	common service

	// API Serives
	LoginApi         *LoginApiService
	EnterpriseApi    *EnterpriseApiService
	EdgeApi          *EdgeApiService
	LicenseApi       *LicenseApiService
	ConfigurationApi *ConfigurationApiService
}

type service struct {
	client *APIClient
}

func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
		jar, err := cookiejar.New(nil)
		if err != nil {
			log.Fatal(err)
		}
		cfg.HTTPClient.Jar = jar
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.LoginApi = (*LoginApiService)(&c.common)
	c.EnterpriseApi = (*EnterpriseApiService)(&c.common)
	c.EdgeApi = (*EdgeApiService)(&c.common)
	c.LicenseApi = (*LicenseApiService)(&c.common)
	c.ConfigurationApi = (*ConfigurationApiService)(&c.common)

	return c
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if strings.Contains(contentType, "application/json") {
		//log.Printf("%s", b)
		//log.Println("####### AAA #######")
		if err = json.Unmarshal(b, v); err != nil {
			//log.Println("####### AAA #######")
			log.Println(err)
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

func (c *APIClient) AddHeader(header string, value string) error {
	c.cfg.DefaultHeader[header] = value
	return nil
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.cfg.HTTPClient.Do(request)
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string,
	method string,
	postBody interface{},
	headerParams map[string]string) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer
	body = &bytes.Buffer{}

	err = json.NewEncoder(body).Encode(postBody)
	if err != nil {
		return nil, err
	}

	var url string
	url = c.cfg.BasePath + "rest" + path

	localVarRequest, err = http.NewRequest(method, url, body)

	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	if len(c.cfg.DefaultHeader) > 0 {
		for h, v := range c.cfg.DefaultHeader {
			localVarRequest.Header.Add(h, v)
		}
	}
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	return localVarRequest, nil
}

func (c *APIClient) restRequest(
	ctx context.Context,
	url string,
	method string,
	postBody interface{},
	headerParams map[string]string) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer
	body = &bytes.Buffer{}

	err = json.NewEncoder(body).Encode(postBody)
	if err != nil {
		return nil, err
	}

	localVarRequest, err = http.NewRequest(method, url, body)

	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers

	}
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	return localVarRequest, nil
}
