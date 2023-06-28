/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.18120-f464f720d215
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTVersionInfo struct for BTVersionInfo
type BTVersionInfo struct {
	CreatedAt   *JSONTime               `json:"createdAt,omitempty"`
	Creator     *BTUserBasicSummaryInfo `json:"creator,omitempty"`
	Description *string                 `json:"description,omitempty"`
	DocumentId  *string                 `json:"documentId,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id                  *string                 `json:"id,omitempty"`
	LastModifier        *BTUserBasicSummaryInfo `json:"lastModifier,omitempty"`
	MetadataWorkspaceId *string                 `json:"metadataWorkspaceId,omitempty"`
	Microversion        *string                 `json:"microversion,omitempty"`
	ModifiedAt          *JSONTime               `json:"modifiedAt,omitempty"`
	// Name of the resource.
	Name         *string          `json:"name,omitempty"`
	OverrideDate *JSONTime        `json:"overrideDate,omitempty"`
	Parent       *string          `json:"parent,omitempty"`
	Purpose      *int32           `json:"purpose,omitempty"`
	Thumbnail    *BTThumbnailInfo `json:"thumbnail,omitempty"`
	Type         *string          `json:"type,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTVersionInfo instantiates a new BTVersionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTVersionInfo() *BTVersionInfo {
	this := BTVersionInfo{}
	return &this
}

// NewBTVersionInfoWithDefaults instantiates a new BTVersionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTVersionInfoWithDefaults() *BTVersionInfo {
	this := BTVersionInfo{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTVersionInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTVersionInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTVersionInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *BTVersionInfo) GetCreator() BTUserBasicSummaryInfo {
	if o == nil || o.Creator == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetCreatorOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *BTVersionInfo) HasCreator() bool {
	if o != nil && o.Creator != nil {
		return true
	}

	return false
}

// SetCreator gets a reference to the given BTUserBasicSummaryInfo and assigns it to the Creator field.
func (o *BTVersionInfo) SetCreator(v BTUserBasicSummaryInfo) {
	o.Creator = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTVersionInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTVersionInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTVersionInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTVersionInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTVersionInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTVersionInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTVersionInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTVersionInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTVersionInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTVersionInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTVersionInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTVersionInfo) SetId(v string) {
	o.Id = &v
}

// GetLastModifier returns the LastModifier field value if set, zero value otherwise.
func (o *BTVersionInfo) GetLastModifier() BTUserBasicSummaryInfo {
	if o == nil || o.LastModifier == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.LastModifier
}

// GetLastModifierOk returns a tuple with the LastModifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetLastModifierOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.LastModifier == nil {
		return nil, false
	}
	return o.LastModifier, true
}

// HasLastModifier returns a boolean if a field has been set.
func (o *BTVersionInfo) HasLastModifier() bool {
	if o != nil && o.LastModifier != nil {
		return true
	}

	return false
}

// SetLastModifier gets a reference to the given BTUserBasicSummaryInfo and assigns it to the LastModifier field.
func (o *BTVersionInfo) SetLastModifier(v BTUserBasicSummaryInfo) {
	o.LastModifier = &v
}

// GetMetadataWorkspaceId returns the MetadataWorkspaceId field value if set, zero value otherwise.
func (o *BTVersionInfo) GetMetadataWorkspaceId() string {
	if o == nil || o.MetadataWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.MetadataWorkspaceId
}

// GetMetadataWorkspaceIdOk returns a tuple with the MetadataWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetMetadataWorkspaceIdOk() (*string, bool) {
	if o == nil || o.MetadataWorkspaceId == nil {
		return nil, false
	}
	return o.MetadataWorkspaceId, true
}

// HasMetadataWorkspaceId returns a boolean if a field has been set.
func (o *BTVersionInfo) HasMetadataWorkspaceId() bool {
	if o != nil && o.MetadataWorkspaceId != nil {
		return true
	}

	return false
}

// SetMetadataWorkspaceId gets a reference to the given string and assigns it to the MetadataWorkspaceId field.
func (o *BTVersionInfo) SetMetadataWorkspaceId(v string) {
	o.MetadataWorkspaceId = &v
}

// GetMicroversion returns the Microversion field value if set, zero value otherwise.
func (o *BTVersionInfo) GetMicroversion() string {
	if o == nil || o.Microversion == nil {
		var ret string
		return ret
	}
	return *o.Microversion
}

// GetMicroversionOk returns a tuple with the Microversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetMicroversionOk() (*string, bool) {
	if o == nil || o.Microversion == nil {
		return nil, false
	}
	return o.Microversion, true
}

// HasMicroversion returns a boolean if a field has been set.
func (o *BTVersionInfo) HasMicroversion() bool {
	if o != nil && o.Microversion != nil {
		return true
	}

	return false
}

// SetMicroversion gets a reference to the given string and assigns it to the Microversion field.
func (o *BTVersionInfo) SetMicroversion(v string) {
	o.Microversion = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *BTVersionInfo) GetModifiedAt() JSONTime {
	if o == nil || o.ModifiedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetModifiedAtOk() (*JSONTime, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *BTVersionInfo) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given JSONTime and assigns it to the ModifiedAt field.
func (o *BTVersionInfo) SetModifiedAt(v JSONTime) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTVersionInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTVersionInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTVersionInfo) SetName(v string) {
	o.Name = &v
}

// GetOverrideDate returns the OverrideDate field value if set, zero value otherwise.
func (o *BTVersionInfo) GetOverrideDate() JSONTime {
	if o == nil || o.OverrideDate == nil {
		var ret JSONTime
		return ret
	}
	return *o.OverrideDate
}

// GetOverrideDateOk returns a tuple with the OverrideDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetOverrideDateOk() (*JSONTime, bool) {
	if o == nil || o.OverrideDate == nil {
		return nil, false
	}
	return o.OverrideDate, true
}

// HasOverrideDate returns a boolean if a field has been set.
func (o *BTVersionInfo) HasOverrideDate() bool {
	if o != nil && o.OverrideDate != nil {
		return true
	}

	return false
}

// SetOverrideDate gets a reference to the given JSONTime and assigns it to the OverrideDate field.
func (o *BTVersionInfo) SetOverrideDate(v JSONTime) {
	o.OverrideDate = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *BTVersionInfo) GetParent() string {
	if o == nil || o.Parent == nil {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetParentOk() (*string, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *BTVersionInfo) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *BTVersionInfo) SetParent(v string) {
	o.Parent = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *BTVersionInfo) GetPurpose() int32 {
	if o == nil || o.Purpose == nil {
		var ret int32
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetPurposeOk() (*int32, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *BTVersionInfo) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given int32 and assigns it to the Purpose field.
func (o *BTVersionInfo) SetPurpose(v int32) {
	o.Purpose = &v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *BTVersionInfo) GetThumbnail() BTThumbnailInfo {
	if o == nil || o.Thumbnail == nil {
		var ret BTThumbnailInfo
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetThumbnailOk() (*BTThumbnailInfo, bool) {
	if o == nil || o.Thumbnail == nil {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *BTVersionInfo) HasThumbnail() bool {
	if o != nil && o.Thumbnail != nil {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given BTThumbnailInfo and assigns it to the Thumbnail field.
func (o *BTVersionInfo) SetThumbnail(v BTThumbnailInfo) {
	o.Thumbnail = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTVersionInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTVersionInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTVersionInfo) SetType(v string) {
	o.Type = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTVersionInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTVersionInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTVersionInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTVersionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastModifier != nil {
		toSerialize["lastModifier"] = o.LastModifier
	}
	if o.MetadataWorkspaceId != nil {
		toSerialize["metadataWorkspaceId"] = o.MetadataWorkspaceId
	}
	if o.Microversion != nil {
		toSerialize["microversion"] = o.Microversion
	}
	if o.ModifiedAt != nil {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OverrideDate != nil {
		toSerialize["overrideDate"] = o.OverrideDate
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.Purpose != nil {
		toSerialize["purpose"] = o.Purpose
	}
	if o.Thumbnail != nil {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTVersionInfo struct {
	value *BTVersionInfo
	isSet bool
}

func (v NullableBTVersionInfo) Get() *BTVersionInfo {
	return v.value
}

func (v *NullableBTVersionInfo) Set(val *BTVersionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTVersionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTVersionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTVersionInfo(val *BTVersionInfo) *NullableBTVersionInfo {
	return &NullableBTVersionInfo{value: val, isSet: true}
}

func (v NullableBTVersionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTVersionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
