/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23367-59c3c9f9feef
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
	"strings"
)

// InsertableApiService InsertableApi service
type InsertableApiService service

type ApiGetLatestInDocumentRequest struct {
	ctx                                    context.Context
	ApiService                             *InsertableApiService
	did                                    string
	includeParts                           *bool
	includeSurfaces                        *bool
	includeSketches                        *bool
	includeReferenceFeatures               *bool
	includeAssemblies                      *bool
	includeFeatureStudios                  *bool
	includeBlobs                           *bool
	allowedBlobMimeTypes                   *string
	excludeNewerFSVersions                 *bool
	maxFeatureScriptVersion                *int32
	includePartStudios                     *bool
	includeFeatures                        *bool
	includeMeshes                          *bool
	includeWires                           *bool
	includeFlattenedBodies                 *bool
	includeApplications                    *bool
	allowedApplicationMimeTypes            *string
	includeCompositeParts                  *bool
	includeFSTables                        *bool
	includeFSComputedPartPropertyFunctions *bool
	includeVariables                       *bool
	includeVariableStudios                 *bool
	allowedBlobExtensions                  *string
}

func (r ApiGetLatestInDocumentRequest) IncludeParts(includeParts bool) ApiGetLatestInDocumentRequest {
	r.includeParts = &includeParts
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeSurfaces(includeSurfaces bool) ApiGetLatestInDocumentRequest {
	r.includeSurfaces = &includeSurfaces
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeSketches(includeSketches bool) ApiGetLatestInDocumentRequest {
	r.includeSketches = &includeSketches
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeReferenceFeatures(includeReferenceFeatures bool) ApiGetLatestInDocumentRequest {
	r.includeReferenceFeatures = &includeReferenceFeatures
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeAssemblies(includeAssemblies bool) ApiGetLatestInDocumentRequest {
	r.includeAssemblies = &includeAssemblies
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeFeatureStudios(includeFeatureStudios bool) ApiGetLatestInDocumentRequest {
	r.includeFeatureStudios = &includeFeatureStudios
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeBlobs(includeBlobs bool) ApiGetLatestInDocumentRequest {
	r.includeBlobs = &includeBlobs
	return r
}

func (r ApiGetLatestInDocumentRequest) AllowedBlobMimeTypes(allowedBlobMimeTypes string) ApiGetLatestInDocumentRequest {
	r.allowedBlobMimeTypes = &allowedBlobMimeTypes
	return r
}

func (r ApiGetLatestInDocumentRequest) ExcludeNewerFSVersions(excludeNewerFSVersions bool) ApiGetLatestInDocumentRequest {
	r.excludeNewerFSVersions = &excludeNewerFSVersions
	return r
}

func (r ApiGetLatestInDocumentRequest) MaxFeatureScriptVersion(maxFeatureScriptVersion int32) ApiGetLatestInDocumentRequest {
	r.maxFeatureScriptVersion = &maxFeatureScriptVersion
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludePartStudios(includePartStudios bool) ApiGetLatestInDocumentRequest {
	r.includePartStudios = &includePartStudios
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeFeatures(includeFeatures bool) ApiGetLatestInDocumentRequest {
	r.includeFeatures = &includeFeatures
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeMeshes(includeMeshes bool) ApiGetLatestInDocumentRequest {
	r.includeMeshes = &includeMeshes
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeWires(includeWires bool) ApiGetLatestInDocumentRequest {
	r.includeWires = &includeWires
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeFlattenedBodies(includeFlattenedBodies bool) ApiGetLatestInDocumentRequest {
	r.includeFlattenedBodies = &includeFlattenedBodies
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeApplications(includeApplications bool) ApiGetLatestInDocumentRequest {
	r.includeApplications = &includeApplications
	return r
}

func (r ApiGetLatestInDocumentRequest) AllowedApplicationMimeTypes(allowedApplicationMimeTypes string) ApiGetLatestInDocumentRequest {
	r.allowedApplicationMimeTypes = &allowedApplicationMimeTypes
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeCompositeParts(includeCompositeParts bool) ApiGetLatestInDocumentRequest {
	r.includeCompositeParts = &includeCompositeParts
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeFSTables(includeFSTables bool) ApiGetLatestInDocumentRequest {
	r.includeFSTables = &includeFSTables
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeFSComputedPartPropertyFunctions(includeFSComputedPartPropertyFunctions bool) ApiGetLatestInDocumentRequest {
	r.includeFSComputedPartPropertyFunctions = &includeFSComputedPartPropertyFunctions
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeVariables(includeVariables bool) ApiGetLatestInDocumentRequest {
	r.includeVariables = &includeVariables
	return r
}

func (r ApiGetLatestInDocumentRequest) IncludeVariableStudios(includeVariableStudios bool) ApiGetLatestInDocumentRequest {
	r.includeVariableStudios = &includeVariableStudios
	return r
}

func (r ApiGetLatestInDocumentRequest) AllowedBlobExtensions(allowedBlobExtensions string) ApiGetLatestInDocumentRequest {
	r.allowedBlobExtensions = &allowedBlobExtensions
	return r
}

func (r ApiGetLatestInDocumentRequest) Execute() (*BTListResponseBTInsertableInfo, *http.Response, error) {
	return r.ApiService.GetLatestInDocumentExecute(r)
}

/*
GetLatestInDocument Get a list of things in this document that can be inserted elsewhere.

* Returns only the latest revision of released insertables.
* Use the document ID (`did`) parameter to specify the source document, not the insertion target.
* For example, you can insert a custom Feature library into another custom Feature library, insert Parts into an Assembly or a Drawing, etc.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @return ApiGetLatestInDocumentRequest
*/
func (a *InsertableApiService) GetLatestInDocument(ctx context.Context, did string) ApiGetLatestInDocumentRequest {
	return ApiGetLatestInDocumentRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
	}
}

// Execute executes the request
//  @return BTListResponseBTInsertableInfo
func (a *InsertableApiService) GetLatestInDocumentExecute(r ApiGetLatestInDocumentRequest) (*BTListResponseBTInsertableInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTListResponseBTInsertableInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InsertableApiService.GetLatestInDocument")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/insertables/d/{did}/latest"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeParts != nil {
		localVarQueryParams.Add("includeParts", parameterToString(*r.includeParts, ""))
	}
	if r.includeSurfaces != nil {
		localVarQueryParams.Add("includeSurfaces", parameterToString(*r.includeSurfaces, ""))
	}
	if r.includeSketches != nil {
		localVarQueryParams.Add("includeSketches", parameterToString(*r.includeSketches, ""))
	}
	if r.includeReferenceFeatures != nil {
		localVarQueryParams.Add("includeReferenceFeatures", parameterToString(*r.includeReferenceFeatures, ""))
	}
	if r.includeAssemblies != nil {
		localVarQueryParams.Add("includeAssemblies", parameterToString(*r.includeAssemblies, ""))
	}
	if r.includeFeatureStudios != nil {
		localVarQueryParams.Add("includeFeatureStudios", parameterToString(*r.includeFeatureStudios, ""))
	}
	if r.includeBlobs != nil {
		localVarQueryParams.Add("includeBlobs", parameterToString(*r.includeBlobs, ""))
	}
	if r.allowedBlobMimeTypes != nil {
		localVarQueryParams.Add("allowedBlobMimeTypes", parameterToString(*r.allowedBlobMimeTypes, ""))
	}
	if r.excludeNewerFSVersions != nil {
		localVarQueryParams.Add("excludeNewerFSVersions", parameterToString(*r.excludeNewerFSVersions, ""))
	}
	if r.maxFeatureScriptVersion != nil {
		localVarQueryParams.Add("maxFeatureScriptVersion", parameterToString(*r.maxFeatureScriptVersion, ""))
	}
	if r.includePartStudios != nil {
		localVarQueryParams.Add("includePartStudios", parameterToString(*r.includePartStudios, ""))
	}
	if r.includeFeatures != nil {
		localVarQueryParams.Add("includeFeatures", parameterToString(*r.includeFeatures, ""))
	}
	if r.includeMeshes != nil {
		localVarQueryParams.Add("includeMeshes", parameterToString(*r.includeMeshes, ""))
	}
	if r.includeWires != nil {
		localVarQueryParams.Add("includeWires", parameterToString(*r.includeWires, ""))
	}
	if r.includeFlattenedBodies != nil {
		localVarQueryParams.Add("includeFlattenedBodies", parameterToString(*r.includeFlattenedBodies, ""))
	}
	if r.includeApplications != nil {
		localVarQueryParams.Add("includeApplications", parameterToString(*r.includeApplications, ""))
	}
	if r.allowedApplicationMimeTypes != nil {
		localVarQueryParams.Add("allowedApplicationMimeTypes", parameterToString(*r.allowedApplicationMimeTypes, ""))
	}
	if r.includeCompositeParts != nil {
		localVarQueryParams.Add("includeCompositeParts", parameterToString(*r.includeCompositeParts, ""))
	}
	if r.includeFSTables != nil {
		localVarQueryParams.Add("includeFSTables", parameterToString(*r.includeFSTables, ""))
	}
	if r.includeFSComputedPartPropertyFunctions != nil {
		localVarQueryParams.Add("includeFSComputedPartPropertyFunctions", parameterToString(*r.includeFSComputedPartPropertyFunctions, ""))
	}
	if r.includeVariables != nil {
		localVarQueryParams.Add("includeVariables", parameterToString(*r.includeVariables, ""))
	}
	if r.includeVariableStudios != nil {
		localVarQueryParams.Add("includeVariableStudios", parameterToString(*r.includeVariableStudios, ""))
	}
	if r.allowedBlobExtensions != nil {
		localVarQueryParams.Add("allowedBlobExtensions", parameterToString(*r.allowedBlobExtensions, ""))
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
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTListResponseBTInsertableInfo
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
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
