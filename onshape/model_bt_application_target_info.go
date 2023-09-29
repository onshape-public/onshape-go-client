/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23149-3610d8cd750e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTApplicationTargetInfo struct for BTApplicationTargetInfo
type BTApplicationTargetInfo struct {
	BaseHref              *string `json:"baseHref,omitempty"`
	ClientId              *string `json:"clientId,omitempty"`
	SupportsCollaboration *bool   `json:"supportsCollaboration,omitempty"`
	TabIconHref           *string `json:"tabIconHref,omitempty"`
}

// NewBTApplicationTargetInfo instantiates a new BTApplicationTargetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTApplicationTargetInfo() *BTApplicationTargetInfo {
	this := BTApplicationTargetInfo{}
	return &this
}

// NewBTApplicationTargetInfoWithDefaults instantiates a new BTApplicationTargetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTApplicationTargetInfoWithDefaults() *BTApplicationTargetInfo {
	this := BTApplicationTargetInfo{}
	return &this
}

// GetBaseHref returns the BaseHref field value if set, zero value otherwise.
func (o *BTApplicationTargetInfo) GetBaseHref() string {
	if o == nil || o.BaseHref == nil {
		var ret string
		return ret
	}
	return *o.BaseHref
}

// GetBaseHrefOk returns a tuple with the BaseHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApplicationTargetInfo) GetBaseHrefOk() (*string, bool) {
	if o == nil || o.BaseHref == nil {
		return nil, false
	}
	return o.BaseHref, true
}

// HasBaseHref returns a boolean if a field has been set.
func (o *BTApplicationTargetInfo) HasBaseHref() bool {
	if o != nil && o.BaseHref != nil {
		return true
	}

	return false
}

// SetBaseHref gets a reference to the given string and assigns it to the BaseHref field.
func (o *BTApplicationTargetInfo) SetBaseHref(v string) {
	o.BaseHref = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BTApplicationTargetInfo) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApplicationTargetInfo) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BTApplicationTargetInfo) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BTApplicationTargetInfo) SetClientId(v string) {
	o.ClientId = &v
}

// GetSupportsCollaboration returns the SupportsCollaboration field value if set, zero value otherwise.
func (o *BTApplicationTargetInfo) GetSupportsCollaboration() bool {
	if o == nil || o.SupportsCollaboration == nil {
		var ret bool
		return ret
	}
	return *o.SupportsCollaboration
}

// GetSupportsCollaborationOk returns a tuple with the SupportsCollaboration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApplicationTargetInfo) GetSupportsCollaborationOk() (*bool, bool) {
	if o == nil || o.SupportsCollaboration == nil {
		return nil, false
	}
	return o.SupportsCollaboration, true
}

// HasSupportsCollaboration returns a boolean if a field has been set.
func (o *BTApplicationTargetInfo) HasSupportsCollaboration() bool {
	if o != nil && o.SupportsCollaboration != nil {
		return true
	}

	return false
}

// SetSupportsCollaboration gets a reference to the given bool and assigns it to the SupportsCollaboration field.
func (o *BTApplicationTargetInfo) SetSupportsCollaboration(v bool) {
	o.SupportsCollaboration = &v
}

// GetTabIconHref returns the TabIconHref field value if set, zero value otherwise.
func (o *BTApplicationTargetInfo) GetTabIconHref() string {
	if o == nil || o.TabIconHref == nil {
		var ret string
		return ret
	}
	return *o.TabIconHref
}

// GetTabIconHrefOk returns a tuple with the TabIconHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApplicationTargetInfo) GetTabIconHrefOk() (*string, bool) {
	if o == nil || o.TabIconHref == nil {
		return nil, false
	}
	return o.TabIconHref, true
}

// HasTabIconHref returns a boolean if a field has been set.
func (o *BTApplicationTargetInfo) HasTabIconHref() bool {
	if o != nil && o.TabIconHref != nil {
		return true
	}

	return false
}

// SetTabIconHref gets a reference to the given string and assigns it to the TabIconHref field.
func (o *BTApplicationTargetInfo) SetTabIconHref(v string) {
	o.TabIconHref = &v
}

func (o BTApplicationTargetInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseHref != nil {
		toSerialize["baseHref"] = o.BaseHref
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.SupportsCollaboration != nil {
		toSerialize["supportsCollaboration"] = o.SupportsCollaboration
	}
	if o.TabIconHref != nil {
		toSerialize["tabIconHref"] = o.TabIconHref
	}
	return json.Marshal(toSerialize)
}

type NullableBTApplicationTargetInfo struct {
	value *BTApplicationTargetInfo
	isSet bool
}

func (v NullableBTApplicationTargetInfo) Get() *BTApplicationTargetInfo {
	return v.value
}

func (v *NullableBTApplicationTargetInfo) Set(val *BTApplicationTargetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTApplicationTargetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTApplicationTargetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTApplicationTargetInfo(val *BTApplicationTargetInfo) *NullableBTApplicationTargetInfo {
	return &NullableBTApplicationTargetInfo{value: val, isSet: true}
}

func (v NullableBTApplicationTargetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTApplicationTargetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
