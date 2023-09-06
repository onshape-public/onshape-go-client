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

// BTConfigurationInfo struct for BTConfigurationInfo
type BTConfigurationInfo struct {
	IsStandardContent *bool                    `json:"isStandardContent,omitempty"`
	Parameters        []ConfigurationInfoEntry `json:"parameters,omitempty"`
}

// NewBTConfigurationInfo instantiates a new BTConfigurationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfigurationInfo() *BTConfigurationInfo {
	this := BTConfigurationInfo{}
	return &this
}

// NewBTConfigurationInfoWithDefaults instantiates a new BTConfigurationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfigurationInfoWithDefaults() *BTConfigurationInfo {
	this := BTConfigurationInfo{}
	return &this
}

// GetIsStandardContent returns the IsStandardContent field value if set, zero value otherwise.
func (o *BTConfigurationInfo) GetIsStandardContent() bool {
	if o == nil || o.IsStandardContent == nil {
		var ret bool
		return ret
	}
	return *o.IsStandardContent
}

// GetIsStandardContentOk returns a tuple with the IsStandardContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationInfo) GetIsStandardContentOk() (*bool, bool) {
	if o == nil || o.IsStandardContent == nil {
		return nil, false
	}
	return o.IsStandardContent, true
}

// HasIsStandardContent returns a boolean if a field has been set.
func (o *BTConfigurationInfo) HasIsStandardContent() bool {
	if o != nil && o.IsStandardContent != nil {
		return true
	}

	return false
}

// SetIsStandardContent gets a reference to the given bool and assigns it to the IsStandardContent field.
func (o *BTConfigurationInfo) SetIsStandardContent(v bool) {
	o.IsStandardContent = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTConfigurationInfo) GetParameters() []ConfigurationInfoEntry {
	if o == nil || o.Parameters == nil {
		var ret []ConfigurationInfoEntry
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationInfo) GetParametersOk() ([]ConfigurationInfoEntry, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTConfigurationInfo) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ConfigurationInfoEntry and assigns it to the Parameters field.
func (o *BTConfigurationInfo) SetParameters(v []ConfigurationInfoEntry) {
	o.Parameters = v
}

func (o BTConfigurationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsStandardContent != nil {
		toSerialize["isStandardContent"] = o.IsStandardContent
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfigurationInfo struct {
	value *BTConfigurationInfo
	isSet bool
}

func (v NullableBTConfigurationInfo) Get() *BTConfigurationInfo {
	return v.value
}

func (v *NullableBTConfigurationInfo) Set(val *BTConfigurationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfigurationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfigurationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfigurationInfo(val *BTConfigurationInfo) *NullableBTConfigurationInfo {
	return &NullableBTConfigurationInfo{value: val, isSet: true}
}

func (v NullableBTConfigurationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfigurationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
