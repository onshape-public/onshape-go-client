/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AliasApiService AliasApi service
type AliasApiService service

type ApiCreateAliasRequest struct {
	ctx           context.Context
	ApiService    *AliasApiService
	bTAliasParams *BTAliasParams
}

func (r ApiCreateAliasRequest) BTAliasParams(bTAliasParams BTAliasParams) ApiCreateAliasRequest {
	r.bTAliasParams = &bTAliasParams
	return r
}

func (r ApiCreateAliasRequest) Execute() (*BTAliasInfo, *http.Response, error) {
	return r.ApiService.CreateAliasExecute(r)
}

/*
CreateAlias Create an alias in your enterprise.

`Manage users and teams` global permission is required to call this API.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateAliasRequest
*/
func (a *AliasApiService) CreateAlias(ctx context.Context) ApiCreateAliasRequest {
	return ApiCreateAliasRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BTAliasInfo
func (a *AliasApiService) CreateAliasExecute(r ApiCreateAliasRequest) (*BTAliasInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTAliasInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AliasApiService.CreateAlias")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/aliases"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTAliasParams == nil {
		return localVarReturnValue, nil, reportError("bTAliasParams is required and must be specified")
	}

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
	localVarPostBody = r.bTAliasParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTAliasInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteAliasRequest struct {
	ctx        context.Context
	ApiService *AliasApiService
	aid        string
}

func (r ApiDeleteAliasRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteAliasExecute(r)
}

/*
DeleteAlias Delete an alias from your enterprise.

`Manage users and teams` global permission is required to call this API.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param aid
	@return ApiDeleteAliasRequest
*/
func (a *AliasApiService) DeleteAlias(ctx context.Context, aid string) ApiDeleteAliasRequest {
	return ApiDeleteAliasRequest{
		ApiService: a,
		ctx:        ctx,
		aid:        aid,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AliasApiService) DeleteAliasExecute(r ApiDeleteAliasRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AliasApiService.DeleteAlias")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/aliases/{aid}"
	localVarPath = strings.Replace(localVarPath, "{"+"aid"+"}", url.PathEscape(parameterToString(r.aid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v map[string]interface{}
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAliasRequest struct {
	ctx        context.Context
	ApiService *AliasApiService
	aid        string
}

func (r ApiGetAliasRequest) Execute() (*BTAliasInfo, *http.Response, error) {
	return r.ApiService.GetAliasExecute(r)
}

/*
GetAlias Get an alias by ID.

Get the information for an alias ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param aid
	@return ApiGetAliasRequest
*/
func (a *AliasApiService) GetAlias(ctx context.Context, aid string) ApiGetAliasRequest {
	return ApiGetAliasRequest{
		ApiService: a,
		ctx:        ctx,
		aid:        aid,
	}
}

// Execute executes the request
//
//	@return BTAliasInfo
func (a *AliasApiService) GetAliasExecute(r ApiGetAliasRequest) (*BTAliasInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTAliasInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AliasApiService.GetAlias")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/aliases/{aid}"
	localVarPath = strings.Replace(localVarPath, "{"+"aid"+"}", url.PathEscape(parameterToString(r.aid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTAliasInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAliasMembersRequest struct {
	ctx        context.Context
	ApiService *AliasApiService
	aid        string
	prefix     *string
	sortColumn *string
	sortOrder  *string
	offset     *int32
	limit      *int32
}

func (r ApiGetAliasMembersRequest) Prefix(prefix string) ApiGetAliasMembersRequest {
	r.prefix = &prefix
	return r
}

func (r ApiGetAliasMembersRequest) SortColumn(sortColumn string) ApiGetAliasMembersRequest {
	r.sortColumn = &sortColumn
	return r
}

func (r ApiGetAliasMembersRequest) SortOrder(sortOrder string) ApiGetAliasMembersRequest {
	r.sortOrder = &sortOrder
	return r
}

func (r ApiGetAliasMembersRequest) Offset(offset int32) ApiGetAliasMembersRequest {
	r.offset = &offset
	return r
}

func (r ApiGetAliasMembersRequest) Limit(limit int32) ApiGetAliasMembersRequest {
	r.limit = &limit
	return r
}

func (r ApiGetAliasMembersRequest) Execute() (*BTListResponseBTAliasEntryInfo, *http.Response, error) {
	return r.ApiService.GetAliasMembersExecute(r)
}

/*
GetAliasMembers Get all users and teams assigned to an alias.

This is a search-like endpoint that returns a subset of the member list. Use `getAlias` to return all members every time it's called.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param aid
	@return ApiGetAliasMembersRequest
*/
func (a *AliasApiService) GetAliasMembers(ctx context.Context, aid string) ApiGetAliasMembersRequest {
	return ApiGetAliasMembersRequest{
		ApiService: a,
		ctx:        ctx,
		aid:        aid,
	}
}

// Execute executes the request
//
//	@return BTListResponseBTAliasEntryInfo
func (a *AliasApiService) GetAliasMembersExecute(r ApiGetAliasMembersRequest) (*BTListResponseBTAliasEntryInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTListResponseBTAliasEntryInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AliasApiService.GetAliasMembers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/aliases/{aid}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"aid"+"}", url.PathEscape(parameterToString(r.aid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.prefix != nil {
		localVarQueryParams.Add("prefix", parameterToString(*r.prefix, ""))
	}
	if r.sortColumn != nil {
		localVarQueryParams.Add("sortColumn", parameterToString(*r.sortColumn, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sortOrder", parameterToString(*r.sortOrder, ""))
	}
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

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTListResponseBTAliasEntryInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAliasesInCompanyRequest struct {
	ctx        context.Context
	ApiService *AliasApiService
	prefix     *string
	sortColumn *string
	sortOrder  *string
	offset     *int32
	limit      *int32
}

func (r ApiGetAliasesInCompanyRequest) Prefix(prefix string) ApiGetAliasesInCompanyRequest {
	r.prefix = &prefix
	return r
}

func (r ApiGetAliasesInCompanyRequest) SortColumn(sortColumn string) ApiGetAliasesInCompanyRequest {
	r.sortColumn = &sortColumn
	return r
}

func (r ApiGetAliasesInCompanyRequest) SortOrder(sortOrder string) ApiGetAliasesInCompanyRequest {
	r.sortOrder = &sortOrder
	return r
}

func (r ApiGetAliasesInCompanyRequest) Offset(offset int32) ApiGetAliasesInCompanyRequest {
	r.offset = &offset
	return r
}

func (r ApiGetAliasesInCompanyRequest) Limit(limit int32) ApiGetAliasesInCompanyRequest {
	r.limit = &limit
	return r
}

func (r ApiGetAliasesInCompanyRequest) Execute() (*BTListResponseBTAliasInfo, *http.Response, error) {
	return r.ApiService.GetAliasesInCompanyExecute(r)
}

/*
GetAliasesInCompany Get a list of all aliases that exist for your enterprise.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAliasesInCompanyRequest
*/
func (a *AliasApiService) GetAliasesInCompany(ctx context.Context) ApiGetAliasesInCompanyRequest {
	return ApiGetAliasesInCompanyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BTListResponseBTAliasInfo
func (a *AliasApiService) GetAliasesInCompanyExecute(r ApiGetAliasesInCompanyRequest) (*BTListResponseBTAliasInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTListResponseBTAliasInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AliasApiService.GetAliasesInCompany")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/aliases"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.prefix != nil {
		localVarQueryParams.Add("prefix", parameterToString(*r.prefix, ""))
	}
	if r.sortColumn != nil {
		localVarQueryParams.Add("sortColumn", parameterToString(*r.sortColumn, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sortOrder", parameterToString(*r.sortOrder, ""))
	}
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

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTListResponseBTAliasInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateAliasRequest struct {
	ctx           context.Context
	ApiService    *AliasApiService
	aid           string
	bTAliasParams *BTAliasParams
}

func (r ApiUpdateAliasRequest) BTAliasParams(bTAliasParams BTAliasParams) ApiUpdateAliasRequest {
	r.bTAliasParams = &bTAliasParams
	return r
}

func (r ApiUpdateAliasRequest) Execute() (*BTAliasInfo, *http.Response, error) {
	return r.ApiService.UpdateAliasExecute(r)
}

/*
UpdateAlias Add, remove, replace, or rename entries in an alias list.

`Manage users and teams` global permission is required to call this API.
* Add new users in the `additions` array.
* Remove existing users in the `removals` array. Attempts to remove a user that does not exist in the Alias list will have no effect.
* Replace the entire Alias list with the `entries` array.
* You can also update the alias' `name` and `description`.
For example, given an Alias with members userA and userB:
* `additions: [userC]` results in [userA, userB, userC]
* `removals: [userB]` results in [userA]
* `entries: [userC, user D]` results in [userC, userD]

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param aid
	@return ApiUpdateAliasRequest
*/
func (a *AliasApiService) UpdateAlias(ctx context.Context, aid string) ApiUpdateAliasRequest {
	return ApiUpdateAliasRequest{
		ApiService: a,
		ctx:        ctx,
		aid:        aid,
	}
}

// Execute executes the request
//
//	@return BTAliasInfo
func (a *AliasApiService) UpdateAliasExecute(r ApiUpdateAliasRequest) (*BTAliasInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTAliasInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AliasApiService.UpdateAlias")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/aliases/{aid}"
	localVarPath = strings.Replace(localVarPath, "{"+"aid"+"}", url.PathEscape(parameterToString(r.aid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTAliasParams == nil {
		return localVarReturnValue, nil, reportError("bTAliasParams is required and must be specified")
	}

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
	localVarPostBody = r.bTAliasParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTAliasInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
