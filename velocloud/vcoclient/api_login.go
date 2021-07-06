
/*
 * VMware SD-WAN Orchestrator API v1
 *
 * ## <p><em>**Please Note**: The 4.2.0 VMware SD-WAN Orchestrator release introduces a <a href=\"https://developer.vmware.com/docs/vmware-sdwan-rest/latest/\">new SD-WAN Orchestration API</a>, informally referred to as \"API v2\" because it is intended to gradually replace the API described in this document as the primary interface supported for Partner and Customer development of SD-WAN orchestration solutions. Designed with extensive input from users of the \"v1\" iteration of the API described below, APIv2 offers an improved developer experience to new and experienced users alike. We encourage all users to try APIv2 and welcome any and all feedback regarding its design via VMware SD-WAN Support. We intend to iteratively add functionality to APIv2 with each VCO software release until it achieves feature parity with the Portal API (\"v1\"), at which time we expect to cease further development of the Portal API. According to our current plan, the Portal API will remain supported until at least mid-CY22. As always, feedback and questions regarding our APIs and relevant roadmap initiatives may be directed to VMware SD-WAN Support.</em></p> <p>The VMware SD-WAN Orchestrator (VCO) powers the management plane in the VMware SD-WAN solution. It offers a broad range of configuration, monitoring and troubleshooting functionality to service providers and enterprises alike. The principal web service with which users interact in order to exercise this functionality is called the <strong>VCO Portal</strong>.</p> <h2>The VCO Portal</h2> <p>The VCO Portal allows network administrators (or scripts and applications acting on their behalf) to manage network and device configuration and query current or historical network and device state. API clients may interact with the Portal via a JSON-RPC interface or a REST-like interface. It is possible to invoke all of the methods described in this document using either interface; there is no Portal functionality for which access is constrained exclusively to either JSON-RPC clients or REST-like ones.</p> <p>Both interfaces accept <strong>exclusively HTTP POST</strong> requests. Both also expect that request bodies, when present, are JSON-formatted -- consistent with RFC 2616, clients are furthermore expected to formally assert where this is the case using the `Content-Type` request header, e.g. `Content-Type: application/json`.</p> <h3>JSON-RPC Interface</h3> <p>The JSON-RPC API accepts calls via the `/portal` URL path (e.g. vco.velocloud.net/portal). Consistent with <a href=\"https://www.jsonrpc.org/specification\">v2.0 of the JSON-RPC specification</a>, the API expects JSON-encoded request payloads that consist of a method name (`method`), a parameters object (`params`), a user-specified unique request identifier (`id`, by convention an integer such as a millisecond-precision epoch timestamp), and a JSON-RPC specification version identifier (`jsonrpc`). The VCO supports only the 2.0 iteration of the JSON-RPC specification, and so the value of the `jsonrpc` parameter should always be the string `\"2.0\"`. A sample request follows:</p> <pre>curl --header 'Content-Type: application/json' --data '{\"jsonrpc\":\"2.0\",\"method\":\"event/getEnterpriseEvents\",\"params\":{\"enterpriseId\":1},\"id\":1}' --cookie cookies.txt -X POST https://vco.velocloud.net/portal/</pre> <h3>REST-like Interface</h3> <p>The REST-like interface eliminates some of the protocol \"overhead\" imposed by the JSON-RPC interface, and may feel more familiar to those familiar with URL-based REST semantics. It also offers a greater degree of interoperability with a range of client-side tools designed for use with traditional REST APIs. The interface is accessible via the `/portal/rest/` base path. In processing REST-like requests, the VCO parses the method name from the portion of the URL path that follows the base path. The request body need contain only the method parameters, e.g.:</p> <pre>curl --header 'Content-Type: application/json' --data '{\"enterpriseId\":1}' --cookie cookies.txt -X POST https://vco.velocloud.net/portal/rest/event/getEnterpriseEvents</pre> <h2>Authentication</h2> <h3> API Tokens </h3> <p>As of the 3.4.0 release the Orchestrator supports a token-based authentication scheme based on the HTTP `Authorization` header, in addition to the cookie-based authentication scheme it has historically supported. Where privileges allow, users may provision and download API tokens via the \"Account\" page on the Orchestrator Web UI. API tokens may be downloaded only once, and should be treated as sensitive, just as you would treat a password. API tokens are typically longer-lived than session cookies, but clients may refresh them when required using either the Orchestrator Web UI or the underlying token management API methods.</p> <p>Tokens are passed to the server in an HTTP `Authorization` header. For example:</p> <pre>curl --header 'Content-Type: application/json' --header 'Authorization: Token &lt;token&gt;' --data '{\"enterpriseId\":1}' -X POST https://vco.velocloud.net/portal/rest/event/getEnterpriseEvents</pre> <h3>Cookie-Based Authentication</h3> <p>The VCO API supports cookie-based authentication. Most programming languages and HTTP client applications expose libraries or options that facilitate the management and use of session cookies, which clients are free to leverage in working with the VCO (e.g. curl exposes the `--cookie-jar` and `--cookie` options, Python's `requests` library <a href=\"https://requests.readthedocs.io/en/master/user/advanced/#session-objects\">exposes a Session interface</a>, etc.). Numerous code samples, wherein authentication is demonstrated in a variety of programming languages, are available via <a href=\"https://code.vmware.com/samples?categories=Sample&keywords=velocloud\">VMware Sample Exchange</a>.</p> <p>Clients initiate sessions by invoking either the `login/enterpriseLogin` or the `login/operatorLogin` method, depending on the user type associated with the client's credentials (Partner and Customer Admins should use the former method, and Operator Admins the latter). In the event of a successful authentication call, the API responds with an HTTP 200 status code and embeds a `velocloud.session` cookie in a `Set-Cookie` response header. When authentication is unsuccessful, the API responds with an HTTP 302 status code and includes a short message elaborating on the failure in a `velocloud.message` cookie. A sample authentication call is demonstrated with the curl command-line utility below (response truncated for brevity):</p> <pre>curl --cookie-jar /tmp/cookie.txt -i -X POST https://vco.velocloud.net/portal/rest/login/enterpriseLogin --data '{\"username\":\"admin@velocloud.net\",\"password\":\"'$SECRET'\"}'<br/>&lt; HTTP/1.1 200 OK<br/>&lt; Set-Cookie: velocloud.session=&lt;token&gt;; &lt;attributes&gt;</pre> <p>Once a client has successfully retrieved a session cookie, it may begin to make API calls to API methods that require authentication by embedding the `velocloud.session` cookie in a `Cookie` request header (programming languages and other client utilities typically provide interfaces that simplify this).</p> <p>Session cookies typically expire after a period of 24 hours (though liftetimes are configurable and may vary across VCO deployments). It is considered best practice to invalidate cookies whenever they are no longer required by initiating a call to the `logout` API method:</p> <pre>curl --cookie /tmp/cookie.txt -X POST https://vco.velocloud.net/portal/rest/logout</pre> <h2>Data Model & Terminology</h2> <p>The terminology of the VCO API schema doesn't always align with the terminology of the Web Console. Consider this a \"cheat sheat\" to aid in interpreting API constructs:</p> <ul><li><strong>Enterprise</strong>: Customer</li><li><strong>Enterprise Proxy</strong>: Partner</li><li><strong>Network</strong>: The Network construct encapsulates all resources in the VCO Operator scope. In typical environments, each VCO has exactly one Network (to which all Partners and Customers belong), and the ID of that Network is 1. Methods in the `/network` namespace are generally reserved exclusively for Operator use.</li><li><strong>Configuration</strong>: Device configurations are modeled in the API schema as a composition of \"Configuration\" entities. There are effectively three distinct types of Configurations: Operator Profiles (also referred to as \"Software Images\"), Customers Profiles (referred to in the API schema as \"Enterprise Configurations\"), and Edge-Specific Profiles.</li><li><strong>Configuration Module</strong>: Each configuration is composed of a set of modules (e.g. `deviceSettings`, `QOS`, `firewall`, `controlPlane`, etc.), wherein the actual configuration `data` resides. In the current version of the API, configuration changes must always be applied at the module level (i.e. via calls to the `configuration/updateConfigurationModule` API method). Partial updates on specific sections of Configuration Module `data` are not (yet) supported.</li><li><strong>Refs</strong>: `refs` are associations between a Network Service (e.g. DNS providers, authentication services, VPN hubs, etc.) and a Configuration (more precisely, a Configuration Module). They should generally be treated as read-only.</li></ul> <h2>Common Parameters</h2> <p>A few parameters appear repeatedly throughout the API schema:</p> <h3>`enterpriseId`</h3> <p>The Portal API enforces that an `enterpriseId` parameter is <strong>required on any request initiated by an Operator or Partner Administrator that accesses, or operates upon, a Customer-managed resource</strong> (e.g. Edges, Profiles, network services). `enterpriseId` is never required for API calls initiated by Customer Administrators (in such cases it is inferred based on the user's credential).</p> <h3>`enterpriseProxyId`</h3> <p>Similar to the `enterpriseId` parameter, the Portal API enforces that an `enterpriseProxyId` parameter is <strong>required on any request initiated by an Operator Administrator that accesses, or operates upon, a Partner-managed resource</strong> (e.g. Partner Events, Partner Gateway Pools, etc.). `enterpriseProxyId` is never required for API calls initiated by Partner Administrators (in such cases it is inferred based on the user's credential).</p> <h3>`networkId`</h3> <p>Some API methods accept a `networkId` parameter, which determines the Network context in which a request is processed. <strong>Partner and Customer Administrators need never use this parameter</strong>, as those user types exist within the context of a single Network which is trivially inferred by the VCO. Meanwhile, Operator Administrators need only specify a value for this parameter in highly-atypical multi-Network VCO deployments (e.g. test environments).</p> <h3>`with`</h3> <p>Many \"fetch\" API methods support a `with` parameter, which allows the user to optionally resolve related entities. `recentLinks` is a special instance of one such option that is supported by methods that fetch Edges, which will cause the API to resolve WAN links for which activity has been recorded in the last 24 hours. This should generally be preferred to the `links` option on methods where it is supported.</p> <h3>`interval`</h3> <p>Many methods, such as those that query events or volumetric flow data, support a query `interval`. The default query interval, inferred by the server when none is otherwise specified, is the most recent 12 hour period.</p> <p>The VCO exposes time series data (e.g. device system health metrics such as CPU and memory usage, network metrics such as latency/jitter/loss, volumetric traffic flow data) via various API methods that accept query intervals. By default, Edges and Gateways report new statistics to the Orchestrator every five minutes. Due to various factors (clock drift, network jitter, server-side processing delays), statistics associated with a given interval beginning at time `t` are often not reflected in API output until time `t + 10 minutes`. As such, we do not recommend using query intervals smaller than 10 minutes in time for these methods.</p> <h2>Datetimes</h2> <p>The Orchestrator API uses UTC time universally. Whenever a method request schema calls for a datetime value, and whenever a response includes a datetime value, the timezone should be inferred to be UTC.</p> <p>The VCO accepts the following datetime formats:</p> <ul><li>13-digit millisecond-precision epoch timestamps (e.g. `1500000000000`)</li><li>Datetime strings formatted consistently with RFC 3339. (e.g. `\"2017-01-01T00:00:00.000Z\"`)</li></ul>  <h2>Rate Limits</h2> <p>The VCO makes use of a rate-limiting mechanism to ensure an equitable allocation of server resources among API clients and safeguard overall system stability. The VCO supports two kinds of API rate limits:</p> <ul><li><strong>Concurrency limits</strong> govern the maximum number of requests that the server processes on behalf of a client or tenant at any given moment.</li><li><strong>Leaky-bucket limits</strong> prevent bursty client activity by limiting the number of requests processed by the server on behalf of a particular client or tenant within a short time period (5 seconds, by default).</li></ul> <p>Rate limiting policies are configurable at the discretion of the VCO Operator, and as such can vary from one environment to the next. <strong>Questions regarding the policies applicable to your environment should be directed to your VCO Operator.</strong></p> <p>When an API client is rate limited, it incurs a \"penalty\" whereby the server refuses to process new requests on its behalf for a brief period of time (5 seconds, by default). During this time, the server issues responses bearing the HTTP 429 status code (Too Many Requests) to the rate-limited client.</p> <p>Developers are advised to take the following steps to (1) avoid triggering defensive rate limiter responses, and (2) mitigate the impact of rate limiter penalties.</p> <ul><li>Handle HTTP 429 responses by implementing a constant-time \"backoff\" (i.e. wait 5 seconds) before attempting another request.</li><li>Limit the number of requests in flight at any given time to no more than a handful.</li><li>Limit polling frequency based on the rate at which source data is refreshed on the server (5 minutes for device statistics, 30s for most other data points).</li><li>When a use case requires querying multiple sites, prefer \"aggregate\" query methods (e.g. `monitoring/getAggregateEdgeLinkMetrics`) to those that query one site at a time (e.g. `metrics/getEdgeLinkMetrics`), wherever possible.</li></ul> <h2>Backward Compatibility</h2> <p>We endeavor not to introduce backward-incompatible changes in the API unless we consider such a change to be essential to the security or stability of the product. We view the following classes of changes as backward-compatible:</p> <ul><li>Changes that introduce new API methods</li><li>Changes that introduce new optional request parameters</li><li>Changes that introduce new response parameters</li><li>Changes that add support for new values to existing enumerations (i.e. properties that have a discrete range of possible values)</li><li>Changes to the format of opaque (i.e. server-generated) strings (e.g. entity identifiers)</li><li>Changes to API error messages</li><li>Changes to Event `message`s and `detail`s</li><li>Changes to the content of alert messages (except where that content is specifically dictated by the user, e.g. for Webhooks)</li></ul> <br> 
 *
 * API version: 4.4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vcoclient

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type LoginApiService service

/*
LoginApiService Authenticate enterprise or partner (MSP) user
Authenticates an enterprise or partner (MSP) user and, upon successful login, returns a velocloud.session cookie. Pass this session cookie in the authentication header in subsequent VCO calls.  If you are using an HTTP client (e.g. Postman) that is configured to automatically follow HTTP redirects, a successful authentication request will cause your client to follow an HTTP 302 redirect to the portal &#39;Home&#39; web page. Your session cookie can then be used to make VCO API calls.  Note that session cookies expire after a period of time specified in the VCO configuration (default is 24 hours).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization


*/
func (a *LoginApiService) LoginEnterpriseLogin(ctx context.Context, authorization AuthObject) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/login/enterpriseLogin"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &authorization
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
LoginApiService Authenticate operator user
Authenticates an operator user and, upon successful login, returns a velocloud.session cookie. Pass this session cookie in the authentication header in subsequent VCO calls.  If you are using an HTTP client (e.g. Postman) that is configured to automatically follow HTTP redirects, a successful authentication request will cause your client to follow an HTTP 302 redirect to the portal &#39;Home&#39; web page. Your session cookie can then be used to make VCO API calls.   Note that session cookies expire after a period of time specified in the VCO configuration (default is 24 hours).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization


*/
func (a *LoginApiService) LoginOperatorLogin(ctx context.Context, authorization AuthObject) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/login/operatorLogin"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &authorization
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
LoginApiService Logout and invalidate authorization session cookie
Logs out the VCO API user and invalidates the session cookie.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().


*/
func (a *LoginApiService) Logout(ctx context.Context) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/logout"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

