/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13685-0fb99d06bde5
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTTranslationRequestInfo struct for BTTranslationRequestInfo
type BTTranslationRequestInfo struct {
	DocumentId    *string `json:"documentId,omitempty"`
	FailureReason *string `json:"failureReason,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// Name of the resource.
	Name                  *string  `json:"name,omitempty"`
	RequestElementId      *string  `json:"requestElementId,omitempty"`
	RequestState          *string  `json:"requestState,omitempty"`
	ResultDocumentId      *string  `json:"resultDocumentId,omitempty"`
	ResultElementIds      []string `json:"resultElementIds,omitempty"`
	ResultExternalDataIds []string `json:"resultExternalDataIds,omitempty"`
	ResultWorkspaceId     *string  `json:"resultWorkspaceId,omitempty"`
	VersionId             *string  `json:"versionId,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef     *string `json:"viewRef,omitempty"`
	WorkspaceId *string `json:"workspaceId,omitempty"`
}

// NewBTTranslationRequestInfo instantiates a new BTTranslationRequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTranslationRequestInfo() *BTTranslationRequestInfo {
	this := BTTranslationRequestInfo{}
	return &this
}

// NewBTTranslationRequestInfoWithDefaults instantiates a new BTTranslationRequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTranslationRequestInfoWithDefaults() *BTTranslationRequestInfo {
	this := BTTranslationRequestInfo{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTTranslationRequestInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTTranslationRequestInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTTranslationRequestInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *BTTranslationRequestInfo) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestInfo) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *BTTranslationRequestInfo) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *BTTranslationRequestInfo) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTTranslationRequestInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTTranslationRequestInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTTranslationRequestInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTTranslationRequestInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTTranslationRequestInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTTranslationRequestInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTTranslationRequestInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTTranslationRequestInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTTranslationRequestInfo) SetName(v string) {
	o.Name = &v
}

// GetRequestElementId returns the RequestElementId field value if set, zero value otherwise.
func (o *BTTranslationRequestInfo) GetRequestElementId() string {
	if o == nil || o.RequestElementId == nil {
		var ret string
		return ret
	}
	return *o.RequestElementId
}

// GetRequestElementIdOk returns a tuple with the RequestElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestInfo) GetRequestElementIdOk() (*string, bool) {
	if o == nil || o.RequestElementId == nil {
		return nil, false
	}
	return o.RequestElementId, true
}

// HasRequestElementId returns a boolean if a field has been set.
func (o *BTTranslationRequestInfo) HasRequestElementId() bool {
	if o != nil && o.RequestElementId != nil {
		return true
	}

	return false
}

// SetRequestElementId gets a reference to the given string and assigns it to the RequestElementId field.
func (o *BTTranslationRequestInfo) SetRequestElementId(v string) {
	o.RequestElementId = &v
}

// GetRequestState returns the RequestState field value if set, zero value otherwise.
func (o *BTTranslationRequestInfo) GetRequestState() string {
	if o == nil || o.RequestState == nil {
		var ret string
		return ret
	}
	return *o.RequestState
}

// GetRequestStateOk returns a tuple with the RequestState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestInfo) GetRequestStateOk() (*string, bool) {
	if o == nil || o.RequestState == nil {
		return nil, false
	}
	return o.RequestState, true
}

// HasRequestState returns a boolean if a field has been set.
func (o *BTTranslationRequestInfo) HasRequestState() bool {
	if o != nil && o.RequestState != nil {
		return true
	}

	return false
}

// SetRequestState gets a reference to the given string and assigns it to the RequestState field.
func (o *BTTranslationRequestInfo) SetRequestState(v string) {
	o.RequestState = &v
}

// GetResultDocumentId returns the ResultDocumentId field value if set, zero value otherwise.
func (o *BTTranslationRequestInfo) GetResultDocumentId() string {
	if o == nil || o.ResultDocumentId == nil {
		var ret string
		return ret
	}
	return *o.ResultDocumentId
}

// GetResultDocumentIdOk returns a tuple with the ResultDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestInfo) GetResultDocumentIdOk() (*string, bool) {
	if o == nil || o.ResultDocumentId == nil {
		return nil, false
	}
	return o.ResultDocumentId, true
}

// HasResultDocumentId returns a boolean if a field has been set.
func (o *BTTranslationRequestInfo) HasResultDocumentId() bool {
	if o != nil && o.ResultDocumentId != nil {
		return true
	}

	return false
}

// SetResultDocumentId gets a reference to the given string and assigns it to the ResultDocumentId field.
func (o *BTTranslationRequestInfo) SetResultDocumentId(v string) {
	o.ResultDocumentId = &v
}

// GetResultElementIds returns the ResultElementIds field value if set, zero value otherwise.
func (o *BTTranslationRequestInfo) GetResultElementIds() []string {
	if o == nil || o.ResultElementIds == nil {
		var ret []string
		return ret
	}
	return o.ResultElementIds
}

// GetResultElementIdsOk returns a tuple with the ResultElementIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestInfo) GetResultElementIdsOk() ([]string, bool) {
	if o == nil || o.ResultElementIds == nil {
		return nil, false
	}
	return o.ResultElementIds, true
}

// HasResultElementIds returns a boolean if a field has been set.
func (o *BTTranslationRequestInfo) HasResultElementIds() bool {
	if o != nil && o.ResultElementIds != nil {
		return true
	}

	return false
}

// SetResultElementIds gets a reference to the given []string and assigns it to the ResultElementIds field.
func (o *BTTranslationRequestInfo) SetResultElementIds(v []string) {
	o.ResultElementIds = v
}

// GetResultExternalDataIds returns the ResultExternalDataIds field value if set, zero value otherwise.
func (o *BTTranslationRequestInfo) GetResultExternalDataIds() []string {
	if o == nil || o.ResultExternalDataIds == nil {
		var ret []string
		return ret
	}
	return o.ResultExternalDataIds
}

// GetResultExternalDataIdsOk returns a tuple with the ResultExternalDataIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestInfo) GetResultExternalDataIdsOk() ([]string, bool) {
	if o == nil || o.ResultExternalDataIds == nil {
		return nil, false
	}
	return o.ResultExternalDataIds, true
}

// HasResultExternalDataIds returns a boolean if a field has been set.
func (o *BTTranslationRequestInfo) HasResultExternalDataIds() bool {
	if o != nil && o.ResultExternalDataIds != nil {
		return true
	}

	return false
}

// SetResultExternalDataIds gets a reference to the given []string and assigns it to the ResultExternalDataIds field.
func (o *BTTranslationRequestInfo) SetResultExternalDataIds(v []string) {
	o.ResultExternalDataIds = v
}

// GetResultWorkspaceId returns the ResultWorkspaceId field value if set, zero value otherwise.
func (o *BTTranslationRequestInfo) GetResultWorkspaceId() string {
	if o == nil || o.ResultWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.ResultWorkspaceId
}

// GetResultWorkspaceIdOk returns a tuple with the ResultWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestInfo) GetResultWorkspaceIdOk() (*string, bool) {
	if o == nil || o.ResultWorkspaceId == nil {
		return nil, false
	}
	return o.ResultWorkspaceId, true
}

// HasResultWorkspaceId returns a boolean if a field has been set.
func (o *BTTranslationRequestInfo) HasResultWorkspaceId() bool {
	if o != nil && o.ResultWorkspaceId != nil {
		return true
	}

	return false
}

// SetResultWorkspaceId gets a reference to the given string and assigns it to the ResultWorkspaceId field.
func (o *BTTranslationRequestInfo) SetResultWorkspaceId(v string) {
	o.ResultWorkspaceId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTTranslationRequestInfo) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestInfo) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTTranslationRequestInfo) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTTranslationRequestInfo) SetVersionId(v string) {
	o.VersionId = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTTranslationRequestInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTTranslationRequestInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTTranslationRequestInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTTranslationRequestInfo) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestInfo) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTTranslationRequestInfo) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTTranslationRequestInfo) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTTranslationRequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.FailureReason != nil {
		toSerialize["failureReason"] = o.FailureReason
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RequestElementId != nil {
		toSerialize["requestElementId"] = o.RequestElementId
	}
	if o.RequestState != nil {
		toSerialize["requestState"] = o.RequestState
	}
	if o.ResultDocumentId != nil {
		toSerialize["resultDocumentId"] = o.ResultDocumentId
	}
	if o.ResultElementIds != nil {
		toSerialize["resultElementIds"] = o.ResultElementIds
	}
	if o.ResultExternalDataIds != nil {
		toSerialize["resultExternalDataIds"] = o.ResultExternalDataIds
	}
	if o.ResultWorkspaceId != nil {
		toSerialize["resultWorkspaceId"] = o.ResultWorkspaceId
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTTranslationRequestInfo struct {
	value *BTTranslationRequestInfo
	isSet bool
}

func (v NullableBTTranslationRequestInfo) Get() *BTTranslationRequestInfo {
	return v.value
}

func (v *NullableBTTranslationRequestInfo) Set(val *BTTranslationRequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTranslationRequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTranslationRequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTranslationRequestInfo(val *BTTranslationRequestInfo) *NullableBTTranslationRequestInfo {
	return &NullableBTTranslationRequestInfo{value: val, isSet: true}
}

func (v NullableBTTranslationRequestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTranslationRequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
