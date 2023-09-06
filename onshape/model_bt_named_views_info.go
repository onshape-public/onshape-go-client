/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.169.21759-9ddbba9fdfb8
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTNamedViewsInfo struct for BTNamedViewsInfo
type BTNamedViewsInfo struct {
	NamedViews *map[string]BTNamedViewInfo `json:"namedViews,omitempty"`
}

// NewBTNamedViewsInfo instantiates a new BTNamedViewsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTNamedViewsInfo() *BTNamedViewsInfo {
	this := BTNamedViewsInfo{}
	return &this
}

// NewBTNamedViewsInfoWithDefaults instantiates a new BTNamedViewsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTNamedViewsInfoWithDefaults() *BTNamedViewsInfo {
	this := BTNamedViewsInfo{}
	return &this
}

// GetNamedViews returns the NamedViews field value if set, zero value otherwise.
func (o *BTNamedViewsInfo) GetNamedViews() map[string]BTNamedViewInfo {
	if o == nil || o.NamedViews == nil {
		var ret map[string]BTNamedViewInfo
		return ret
	}
	return *o.NamedViews
}

// GetNamedViewsOk returns a tuple with the NamedViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNamedViewsInfo) GetNamedViewsOk() (*map[string]BTNamedViewInfo, bool) {
	if o == nil || o.NamedViews == nil {
		return nil, false
	}
	return o.NamedViews, true
}

// HasNamedViews returns a boolean if a field has been set.
func (o *BTNamedViewsInfo) HasNamedViews() bool {
	if o != nil && o.NamedViews != nil {
		return true
	}

	return false
}

// SetNamedViews gets a reference to the given map[string]BTNamedViewInfo and assigns it to the NamedViews field.
func (o *BTNamedViewsInfo) SetNamedViews(v map[string]BTNamedViewInfo) {
	o.NamedViews = &v
}

func (o BTNamedViewsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NamedViews != nil {
		toSerialize["namedViews"] = o.NamedViews
	}
	return json.Marshal(toSerialize)
}

type NullableBTNamedViewsInfo struct {
	value *BTNamedViewsInfo
	isSet bool
}

func (v NullableBTNamedViewsInfo) Get() *BTNamedViewsInfo {
	return v.value
}

func (v *NullableBTNamedViewsInfo) Set(val *BTNamedViewsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTNamedViewsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTNamedViewsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTNamedViewsInfo(val *BTNamedViewsInfo) *NullableBTNamedViewsInfo {
	return &NullableBTNamedViewsInfo{value: val, isSet: true}
}

func (v NullableBTNamedViewsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTNamedViewsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
