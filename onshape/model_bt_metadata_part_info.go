/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.17369-82f2ed5d514e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMetadataPartInfo struct for BTMetadataPartInfo
type BTMetadataPartInfo struct {
	Href            *string                  `json:"href,omitempty"`
	JsonType        string                   `json:"jsonType"`
	Properties      []BTMetadataPropertyInfo `json:"properties,omitempty"`
	Thumbnail       *BTThumbnailInfo         `json:"thumbnail,omitempty"`
	IsFlattenedBody *bool                    `json:"isFlattenedBody,omitempty"`
	MeshState       *int32                   `json:"meshState,omitempty"`
	PartId          *string                  `json:"partId,omitempty"`
	PartIdentity    *string                  `json:"partIdentity,omitempty"`
	PartType        *string                  `json:"partType,omitempty"`
}

// NewBTMetadataPartInfo instantiates a new BTMetadataPartInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMetadataPartInfo(jsonType string) *BTMetadataPartInfo {
	this := BTMetadataPartInfo{}
	this.JsonType = jsonType
	return &this
}

// NewBTMetadataPartInfoWithDefaults instantiates a new BTMetadataPartInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMetadataPartInfoWithDefaults() *BTMetadataPartInfo {
	this := BTMetadataPartInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTMetadataPartInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPartInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTMetadataPartInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTMetadataPartInfo) SetHref(v string) {
	o.Href = &v
}

// GetJsonType returns the JsonType field value
func (o *BTMetadataPartInfo) GetJsonType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *BTMetadataPartInfo) GetJsonTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonType, true
}

// SetJsonType sets field value
func (o *BTMetadataPartInfo) SetJsonType(v string) {
	o.JsonType = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTMetadataPartInfo) GetProperties() []BTMetadataPropertyInfo {
	if o == nil || o.Properties == nil {
		var ret []BTMetadataPropertyInfo
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPartInfo) GetPropertiesOk() ([]BTMetadataPropertyInfo, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTMetadataPartInfo) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []BTMetadataPropertyInfo and assigns it to the Properties field.
func (o *BTMetadataPartInfo) SetProperties(v []BTMetadataPropertyInfo) {
	o.Properties = v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *BTMetadataPartInfo) GetThumbnail() BTThumbnailInfo {
	if o == nil || o.Thumbnail == nil {
		var ret BTThumbnailInfo
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPartInfo) GetThumbnailOk() (*BTThumbnailInfo, bool) {
	if o == nil || o.Thumbnail == nil {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *BTMetadataPartInfo) HasThumbnail() bool {
	if o != nil && o.Thumbnail != nil {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given BTThumbnailInfo and assigns it to the Thumbnail field.
func (o *BTMetadataPartInfo) SetThumbnail(v BTThumbnailInfo) {
	o.Thumbnail = &v
}

// GetIsFlattenedBody returns the IsFlattenedBody field value if set, zero value otherwise.
func (o *BTMetadataPartInfo) GetIsFlattenedBody() bool {
	if o == nil || o.IsFlattenedBody == nil {
		var ret bool
		return ret
	}
	return *o.IsFlattenedBody
}

// GetIsFlattenedBodyOk returns a tuple with the IsFlattenedBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPartInfo) GetIsFlattenedBodyOk() (*bool, bool) {
	if o == nil || o.IsFlattenedBody == nil {
		return nil, false
	}
	return o.IsFlattenedBody, true
}

// HasIsFlattenedBody returns a boolean if a field has been set.
func (o *BTMetadataPartInfo) HasIsFlattenedBody() bool {
	if o != nil && o.IsFlattenedBody != nil {
		return true
	}

	return false
}

// SetIsFlattenedBody gets a reference to the given bool and assigns it to the IsFlattenedBody field.
func (o *BTMetadataPartInfo) SetIsFlattenedBody(v bool) {
	o.IsFlattenedBody = &v
}

// GetMeshState returns the MeshState field value if set, zero value otherwise.
func (o *BTMetadataPartInfo) GetMeshState() int32 {
	if o == nil || o.MeshState == nil {
		var ret int32
		return ret
	}
	return *o.MeshState
}

// GetMeshStateOk returns a tuple with the MeshState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPartInfo) GetMeshStateOk() (*int32, bool) {
	if o == nil || o.MeshState == nil {
		return nil, false
	}
	return o.MeshState, true
}

// HasMeshState returns a boolean if a field has been set.
func (o *BTMetadataPartInfo) HasMeshState() bool {
	if o != nil && o.MeshState != nil {
		return true
	}

	return false
}

// SetMeshState gets a reference to the given int32 and assigns it to the MeshState field.
func (o *BTMetadataPartInfo) SetMeshState(v int32) {
	o.MeshState = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTMetadataPartInfo) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPartInfo) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTMetadataPartInfo) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTMetadataPartInfo) SetPartId(v string) {
	o.PartId = &v
}

// GetPartIdentity returns the PartIdentity field value if set, zero value otherwise.
func (o *BTMetadataPartInfo) GetPartIdentity() string {
	if o == nil || o.PartIdentity == nil {
		var ret string
		return ret
	}
	return *o.PartIdentity
}

// GetPartIdentityOk returns a tuple with the PartIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPartInfo) GetPartIdentityOk() (*string, bool) {
	if o == nil || o.PartIdentity == nil {
		return nil, false
	}
	return o.PartIdentity, true
}

// HasPartIdentity returns a boolean if a field has been set.
func (o *BTMetadataPartInfo) HasPartIdentity() bool {
	if o != nil && o.PartIdentity != nil {
		return true
	}

	return false
}

// SetPartIdentity gets a reference to the given string and assigns it to the PartIdentity field.
func (o *BTMetadataPartInfo) SetPartIdentity(v string) {
	o.PartIdentity = &v
}

// GetPartType returns the PartType field value if set, zero value otherwise.
func (o *BTMetadataPartInfo) GetPartType() string {
	if o == nil || o.PartType == nil {
		var ret string
		return ret
	}
	return *o.PartType
}

// GetPartTypeOk returns a tuple with the PartType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPartInfo) GetPartTypeOk() (*string, bool) {
	if o == nil || o.PartType == nil {
		return nil, false
	}
	return o.PartType, true
}

// HasPartType returns a boolean if a field has been set.
func (o *BTMetadataPartInfo) HasPartType() bool {
	if o != nil && o.PartType != nil {
		return true
	}

	return false
}

// SetPartType gets a reference to the given string and assigns it to the PartType field.
func (o *BTMetadataPartInfo) SetPartType(v string) {
	o.PartType = &v
}

func (o BTMetadataPartInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["jsonType"] = o.JsonType
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Thumbnail != nil {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if o.IsFlattenedBody != nil {
		toSerialize["isFlattenedBody"] = o.IsFlattenedBody
	}
	if o.MeshState != nil {
		toSerialize["meshState"] = o.MeshState
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PartIdentity != nil {
		toSerialize["partIdentity"] = o.PartIdentity
	}
	if o.PartType != nil {
		toSerialize["partType"] = o.PartType
	}
	return json.Marshal(toSerialize)
}

type NullableBTMetadataPartInfo struct {
	value *BTMetadataPartInfo
	isSet bool
}

func (v NullableBTMetadataPartInfo) Get() *BTMetadataPartInfo {
	return v.value
}

func (v *NullableBTMetadataPartInfo) Set(val *BTMetadataPartInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataPartInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataPartInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataPartInfo(val *BTMetadataPartInfo) *NullableBTMetadataPartInfo {
	return &NullableBTMetadataPartInfo{value: val, isSet: true}
}

func (v NullableBTMetadataPartInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataPartInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
