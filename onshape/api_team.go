/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

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

// TeamAPIService TeamAPI service
type TeamAPIService service

type ApiFindRequest struct {
	ctx                      context.Context
	ApiService               *TeamAPIService
	query                    *string
	filter                   *int32
	uid                      *string
	companyId                *string
	offset                   *int32
	limit                    *int32
	sortColumn               *string
	sortOrder                *string
	includeCompanyOwnedTeams *bool
}

func (r ApiFindRequest) Query(query string) ApiFindRequest {
	r.query = &query
	return r
}

func (r ApiFindRequest) Filter(filter int32) ApiFindRequest {
	r.filter = &filter
	return r
}

func (r ApiFindRequest) Uid(uid string) ApiFindRequest {
	r.uid = &uid
	return r
}

func (r ApiFindRequest) CompanyId(companyId string) ApiFindRequest {
	r.companyId = &companyId
	return r
}

func (r ApiFindRequest) Offset(offset int32) ApiFindRequest {
	r.offset = &offset
	return r
}

func (r ApiFindRequest) Limit(limit int32) ApiFindRequest {
	r.limit = &limit
	return r
}

func (r ApiFindRequest) SortColumn(sortColumn string) ApiFindRequest {
	r.sortColumn = &sortColumn
	return r
}

func (r ApiFindRequest) SortOrder(sortOrder string) ApiFindRequest {
	r.sortOrder = &sortOrder
	return r
}

func (r ApiFindRequest) IncludeCompanyOwnedTeams(includeCompanyOwnedTeams bool) ApiFindRequest {
	r.includeCompanyOwnedTeams = &includeCompanyOwnedTeams
	return r
}

func (r ApiFindRequest) Execute() (*BTGlobalTreeNodeListResponseBTTeamInfo, *http.Response, error) {
	return r.ApiService.FindExecute(r)
}

/*
Find Get a list of all teams the current user belongs to.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFindRequest
*/
func (a *TeamAPIService) Find(ctx context.Context) ApiFindRequest {
	return ApiFindRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BTGlobalTreeNodeListResponseBTTeamInfo
func (a *TeamAPIService) FindExecute(r ApiFindRequest) (*BTGlobalTreeNodeListResponseBTTeamInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTGlobalTreeNodeListResponseBTTeamInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamAPIService.Find")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.query != nil {
		localVarQueryParams.Add("query", parameterToString(*r.query, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.uid != nil {
		localVarQueryParams.Add("uid", parameterToString(*r.uid, ""))
	}
	if r.companyId != nil {
		localVarQueryParams.Add("companyId", parameterToString(*r.companyId, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.sortColumn != nil {
		localVarQueryParams.Add("sortColumn", parameterToString(*r.sortColumn, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sortOrder", parameterToString(*r.sortOrder, ""))
	}
	if r.includeCompanyOwnedTeams != nil {
		localVarQueryParams.Add("includeCompanyOwnedTeams", parameterToString(*r.includeCompanyOwnedTeams, ""))
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
		var v BTGlobalTreeNodeListResponseBTTeamInfo
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

type ApiGetMembersRequest struct {
	ctx        context.Context
	ApiService *TeamAPIService
	tid        string
	sortColumn *string
	sortOrder  *string
	offset     *int32
	limit      *int32
	q          *string
}

func (r ApiGetMembersRequest) SortColumn(sortColumn string) ApiGetMembersRequest {
	r.sortColumn = &sortColumn
	return r
}

func (r ApiGetMembersRequest) SortOrder(sortOrder string) ApiGetMembersRequest {
	r.sortOrder = &sortOrder
	return r
}

func (r ApiGetMembersRequest) Offset(offset int32) ApiGetMembersRequest {
	r.offset = &offset
	return r
}

func (r ApiGetMembersRequest) Limit(limit int32) ApiGetMembersRequest {
	r.limit = &limit
	return r
}

func (r ApiGetMembersRequest) Q(q string) ApiGetMembersRequest {
	r.q = &q
	return r
}

func (r ApiGetMembersRequest) Execute() (*BTListResponseBTTeamMemberInfo, *http.Response, error) {
	return r.ApiService.GetMembersExecute(r)
}

/*
GetMembers Get a list of a team's members.

Returns a maximum of 20 per page.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tid
	@return ApiGetMembersRequest
*/
func (a *TeamAPIService) GetMembers(ctx context.Context, tid string) ApiGetMembersRequest {
	return ApiGetMembersRequest{
		ApiService: a,
		ctx:        ctx,
		tid:        tid,
	}
}

// Execute executes the request
//
//	@return BTListResponseBTTeamMemberInfo
func (a *TeamAPIService) GetMembersExecute(r ApiGetMembersRequest) (*BTListResponseBTTeamMemberInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTListResponseBTTeamMemberInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamAPIService.GetMembers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{tid}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"tid"+"}", url.PathEscape(parameterToString(r.tid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.q != nil {
		localVarQueryParams.Add("q", parameterToString(*r.q, ""))
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
		var v BTListResponseBTTeamMemberInfo
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

type ApiGetTeamRequest struct {
	ctx        context.Context
	ApiService *TeamAPIService
	tid        string
}

func (r ApiGetTeamRequest) Execute() (*BTTeamInfo, *http.Response, error) {
	return r.ApiService.GetTeamExecute(r)
}

/*
GetTeam Get team information by team ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tid
	@return ApiGetTeamRequest
*/
func (a *TeamAPIService) GetTeam(ctx context.Context, tid string) ApiGetTeamRequest {
	return ApiGetTeamRequest{
		ApiService: a,
		ctx:        ctx,
		tid:        tid,
	}
}

// Execute executes the request
//
//	@return BTTeamInfo
func (a *TeamAPIService) GetTeamExecute(r ApiGetTeamRequest) (*BTTeamInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTTeamInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamAPIService.GetTeam")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{tid}"
	localVarPath = strings.Replace(localVarPath, "{"+"tid"+"}", url.PathEscape(parameterToString(r.tid, "")), -1)

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
		var v BTTeamInfo
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
