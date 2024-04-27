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

func (r ApiGetOpenApiRequest) Execute() (*OpenAPI, *http.Response, error) {
	return r.ApiService.GetOpenApiExecute(r)
}

/*
GetOpenApi Get the OpenAPI specification for the Onshape REST API.

The Onshape API OpenAPI specification is returned in the JSON format.

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
//
//	@return OpenAPI
func (a *OpenApiApiService) GetOpenApiExecute(r ApiGetOpenApiRequest) (*OpenAPI, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OpenAPI
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenApiApiService.GetOpenApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
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
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09", "application/yaml;charset=UTF-8; qs=0.09"}

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
		var v OpenAPI
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

type ApiGetTagsRequest struct {
	ctx        context.Context
	ApiService *OpenApiApiService
}

func (r ApiGetTagsRequest) Execute() ([]Tag, *http.Response, error) {
	return r.ApiService.GetTagsExecute(r)
}

/*
GetTags Get the list of tags in the Onshape OpenAPI specification.

Tags are used to group operations. For example, `Document` groups operations on documents.

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
//
//	@return []Tag
func (a *OpenApiApiService) GetTagsExecute(r ApiGetTagsRequest) ([]Tag, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Tag
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenApiApiService.GetTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
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
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09", "application/yaml;charset=UTF-8; qs=0.09"}

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
		var v []Tag
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
