/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.166.18417-1bd990c6fbaa
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

// OpenApiApiService OpenApiApi service
type OpenApiApiService service

type ApiGetOpenApiRequest struct {
	ctx                   context.Context
	ApiService            *OpenApiApiService
	forceReload           *bool
	version               *string
	versionAlias          *VersionAlias
	noFilter              *bool
	includedTags          *[]string
	excludedTags          *[]string
	includeDeprecated     *bool
	onlyDeprecated        *bool
	documentationStatuses *[]Status
	restUserRole          *BTRestUserRole
	operationIds          *[]string
	excludedOperationIds  *[]string
}

// Force reload the OpenApi definition. Only works when asking for the latest version.
func (r ApiGetOpenApiRequest) ForceReload(forceReload bool) ApiGetOpenApiRequest {
	r.forceReload = &forceReload
	return r
}

// Specify a version of Onshape from which the OpenAPI is generated. If &#39;*&#39; is specified in any of the version fields, that indicates any version if acceptable.
func (r ApiGetOpenApiRequest) Version(version string) ApiGetOpenApiRequest {
	r.version = &version
	return r
}

// Version aliases based on the currently released version.
func (r ApiGetOpenApiRequest) VersionAlias(versionAlias VersionAlias) ApiGetOpenApiRequest {
	r.versionAlias = &versionAlias
	return r
}

// Do not filter the specification at all.
func (r ApiGetOpenApiRequest) NoFilter(noFilter bool) ApiGetOpenApiRequest {
	r.noFilter = &noFilter
	return r
}

// Return only operations with tags included in includedTags.
func (r ApiGetOpenApiRequest) IncludedTags(includedTags []string) ApiGetOpenApiRequest {
	r.includedTags = &includedTags
	return r
}

// If an operation contains an excluded tag, it is not returned from this endpoint.
func (r ApiGetOpenApiRequest) ExcludedTags(excludedTags []string) ApiGetOpenApiRequest {
	r.excludedTags = &excludedTags
	return r
}

// Include deprecated endpoints.
func (r ApiGetOpenApiRequest) IncludeDeprecated(includeDeprecated bool) ApiGetOpenApiRequest {
	r.includeDeprecated = &includeDeprecated
	return r
}

// Only include deprecated endpoints.
func (r ApiGetOpenApiRequest) OnlyDeprecated(onlyDeprecated bool) ApiGetOpenApiRequest {
	r.onlyDeprecated = &onlyDeprecated
	return r
}

// Only return endpoints that have the specified documentation status. Default is to return all the endpoints the user should have access to.
func (r ApiGetOpenApiRequest) DocumentationStatuses(documentationStatuses []Status) ApiGetOpenApiRequest {
	r.documentationStatuses = &documentationStatuses
	return r
}

// The REST user role for which this spec is requested.
func (r ApiGetOpenApiRequest) RestUserRole(restUserRole BTRestUserRole) ApiGetOpenApiRequest {
	r.restUserRole = &restUserRole
	return r
}

// Only return operations with specified ids.
func (r ApiGetOpenApiRequest) OperationIds(operationIds []string) ApiGetOpenApiRequest {
	r.operationIds = &operationIds
	return r
}

// Do not return operations with specified ids.
func (r ApiGetOpenApiRequest) ExcludedOperationIds(excludedOperationIds []string) ApiGetOpenApiRequest {
	r.excludedOperationIds = &excludedOperationIds
	return r
}

func (r ApiGetOpenApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetOpenApiExecute(r)
}

/*
GetOpenApi Method for GetOpenApi

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetOpenApiRequest
*/
func (a *OpenApiApiService) GetOpenApi(ctx context.Context) ApiGetOpenApiRequest {
	return ApiGetOpenApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *OpenApiApiService) GetOpenApiExecute(r ApiGetOpenApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenApiApiService.GetOpenApi")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openapi"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.forceReload != nil {
		localVarQueryParams.Add("forceReload", parameterToString(*r.forceReload, ""))
	}
	if r.version != nil {
		localVarQueryParams.Add("version", parameterToString(*r.version, ""))
	}
	if r.versionAlias != nil {
		localVarQueryParams.Add("versionAlias", parameterToString(*r.versionAlias, ""))
	}
	if r.noFilter != nil {
		localVarQueryParams.Add("noFilter", parameterToString(*r.noFilter, ""))
	}
	if r.includedTags != nil {
		t := *r.includedTags
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("includedTags", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("includedTags", parameterToString(t, "multi"))
		}
	}
	if r.excludedTags != nil {
		t := *r.excludedTags
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("excludedTags", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("excludedTags", parameterToString(t, "multi"))
		}
	}
	if r.includeDeprecated != nil {
		localVarQueryParams.Add("includeDeprecated", parameterToString(*r.includeDeprecated, ""))
	}
	if r.onlyDeprecated != nil {
		localVarQueryParams.Add("onlyDeprecated", parameterToString(*r.onlyDeprecated, ""))
	}
	if r.documentationStatuses != nil {
		t := *r.documentationStatuses
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("documentationStatuses", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("documentationStatuses", parameterToString(t, "multi"))
		}
	}
	if r.restUserRole != nil {
		localVarQueryParams.Add("restUserRole", parameterToString(*r.restUserRole, ""))
	}
	if r.operationIds != nil {
		t := *r.operationIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("operationIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("operationIds", parameterToString(t, "multi"))
		}
	}
	if r.excludedOperationIds != nil {
		t := *r.excludedOperationIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("excludedOperationIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("excludedOperationIds", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetTagsRequest struct {
	ctx        context.Context
	ApiService *OpenApiApiService
}

func (r ApiGetTagsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetTagsExecute(r)
}

/*
GetTags Method for GetTags

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTagsRequest
*/
func (a *OpenApiApiService) GetTags(ctx context.Context) ApiGetTagsRequest {
	return ApiGetTagsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *OpenApiApiService) GetTagsExecute(r ApiGetTagsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenApiApiService.GetTags")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openapi/tags"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
