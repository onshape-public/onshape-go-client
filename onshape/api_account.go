/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5768-0013f50d23d2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AccountApiService AccountApi service
type AccountApiService service

type ApiCancelPurchaseNewRequest struct {
	ctx               context.Context
	ApiService        *AccountApiService
	aid               string
	pid               string
	cancelImmediately *bool
}

func (r ApiCancelPurchaseNewRequest) CancelImmediately(cancelImmediately bool) ApiCancelPurchaseNewRequest {
	r.cancelImmediately = &cancelImmediately
	return r
}

func (r ApiCancelPurchaseNewRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CancelPurchaseNewExecute(r)
}

/*
CancelPurchaseNew Cancel recurring subscription for the account by account ID and purchase ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param aid
 @param pid
 @return ApiCancelPurchaseNewRequest
*/
func (a *AccountApiService) CancelPurchaseNew(ctx context.Context, aid string, pid string) ApiCancelPurchaseNewRequest {
	return ApiCancelPurchaseNewRequest{
		ApiService: a,
		ctx:        ctx,
		aid:        aid,
		pid:        pid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AccountApiService) CancelPurchaseNewExecute(r ApiCancelPurchaseNewRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountApiService.CancelPurchaseNew")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{aid}/purchases/{pid}"
	localVarPath = strings.Replace(localVarPath, "{"+"aid"+"}", url.PathEscape(parameterToString(r.aid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pid"+"}", url.PathEscape(parameterToString(r.pid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cancelImmediately != nil {
		localVarQueryParams.Add("cancelImmediately", parameterToString(*r.cancelImmediately, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader
	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v map[string]interface{}
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiConsumePurchaseRequest struct {
	ctx                  context.Context
	ApiService           *AccountApiService
	pid                  string
	bTPurchaseUserParams *BTPurchaseUserParams
}

func (r ApiConsumePurchaseRequest) BTPurchaseUserParams(bTPurchaseUserParams BTPurchaseUserParams) ApiConsumePurchaseRequest {
	r.bTPurchaseUserParams = &bTPurchaseUserParams
	return r
}

func (r ApiConsumePurchaseRequest) Execute() (*BTPurchaseInfo, *http.Response, error) {
	return r.ApiService.ConsumePurchaseExecute(r)
}

/*
ConsumePurchase Mark purchase consumed by the user by purchase ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pid
 @return ApiConsumePurchaseRequest
*/
func (a *AccountApiService) ConsumePurchase(ctx context.Context, pid string) ApiConsumePurchaseRequest {
	return ApiConsumePurchaseRequest{
		ApiService: a,
		ctx:        ctx,
		pid:        pid,
	}
}

// Execute executes the request
//  @return BTPurchaseInfo
func (a *AccountApiService) ConsumePurchaseExecute(r ApiConsumePurchaseRequest) (*BTPurchaseInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTPurchaseInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountApiService.ConsumePurchase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/purchases/{pid}/consume"
	localVarPath = strings.Replace(localVarPath, "{"+"pid"+"}", url.PathEscape(parameterToString(r.pid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bTPurchaseUserParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader
	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTPurchaseInfo
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPlanPurchasesRequest struct {
	ctx        context.Context
	ApiService *AccountApiService
	planId     string
	offset     *int32
	limit      *int32
}

func (r ApiGetPlanPurchasesRequest) Offset(offset int32) ApiGetPlanPurchasesRequest {
	r.offset = &offset
	return r
}

func (r ApiGetPlanPurchasesRequest) Limit(limit int32) ApiGetPlanPurchasesRequest {
	r.limit = &limit
	return r
}

func (r ApiGetPlanPurchasesRequest) Execute() (*BTListResponseBTPurchaseInfo, *http.Response, error) {
	return r.ApiService.GetPlanPurchasesExecute(r)
}

/*
GetPlanPurchases Method for GetPlanPurchases

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param planId
 @return ApiGetPlanPurchasesRequest
*/
func (a *AccountApiService) GetPlanPurchases(ctx context.Context, planId string) ApiGetPlanPurchasesRequest {
	return ApiGetPlanPurchasesRequest{
		ApiService: a,
		ctx:        ctx,
		planId:     planId,
	}
}

// Execute executes the request
//  @return BTListResponseBTPurchaseInfo
func (a *AccountApiService) GetPlanPurchasesExecute(r ApiGetPlanPurchasesRequest) (*BTListResponseBTPurchaseInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTListResponseBTPurchaseInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountApiService.GetPlanPurchases")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/plans/{planId}/purchases"
	localVarPath = strings.Replace(localVarPath, "{"+"planId"+"}", url.PathEscape(parameterToString(r.planId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader
	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTListResponseBTPurchaseInfo
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPurchasesRequest struct {
	ctx             context.Context
	ApiService      *AccountApiService
	all             *bool
	ownPurchaseOnly *bool
}

func (r ApiGetPurchasesRequest) All(all bool) ApiGetPurchasesRequest {
	r.all = &all
	return r
}

func (r ApiGetPurchasesRequest) OwnPurchaseOnly(ownPurchaseOnly bool) ApiGetPurchasesRequest {
	r.ownPurchaseOnly = &ownPurchaseOnly
	return r
}

func (r ApiGetPurchasesRequest) Execute() ([]BTPurchaseInfo, *http.Response, error) {
	return r.ApiService.GetPurchasesExecute(r)
}

/*
GetPurchases Retrieve an array of the user’s App Store purchases.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPurchasesRequest
*/
func (a *AccountApiService) GetPurchases(ctx context.Context) ApiGetPurchasesRequest {
	return ApiGetPurchasesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []BTPurchaseInfo
func (a *AccountApiService) GetPurchasesExecute(r ApiGetPurchasesRequest) ([]BTPurchaseInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []BTPurchaseInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountApiService.GetPurchases")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/purchases"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.all != nil {
		localVarQueryParams.Add("all", parameterToString(*r.all, ""))
	}
	if r.ownPurchaseOnly != nil {
		localVarQueryParams.Add("ownPurchaseOnly", parameterToString(*r.ownPurchaseOnly, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader
	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v []BTPurchaseInfo
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
