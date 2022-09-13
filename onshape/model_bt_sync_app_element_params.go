/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6438-5fd2d9755d52
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSyncAppElementParams struct for BTSyncAppElementParams
type BTSyncAppElementParams struct {
	Description *string  `json:"description,omitempty"`
	Elements    []string `json:"elements,omitempty"`
}

// NewBTSyncAppElementParams instantiates a new BTSyncAppElementParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSyncAppElementParams() *BTSyncAppElementParams {
	this := BTSyncAppElementParams{}
	return &this
}

// NewBTSyncAppElementParamsWithDefaults instantiates a new BTSyncAppElementParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSyncAppElementParamsWithDefaults() *BTSyncAppElementParams {
	this := BTSyncAppElementParams{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTSyncAppElementParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSyncAppElementParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTSyncAppElementParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTSyncAppElementParams) SetDescription(v string) {
	o.Description = &v
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *BTSyncAppElementParams) GetElements() []string {
	if o == nil || o.Elements == nil {
		var ret []string
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSyncAppElementParams) GetElementsOk() ([]string, bool) {
	if o == nil || o.Elements == nil {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *BTSyncAppElementParams) HasElements() bool {
	if o != nil && o.Elements != nil {
		return true
	}

	return false
}

// SetElements gets a reference to the given []string and assigns it to the Elements field.
func (o *BTSyncAppElementParams) SetElements(v []string) {
	o.Elements = v
}

func (o BTSyncAppElementParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Elements != nil {
		toSerialize["elements"] = o.Elements
	}
	return json.Marshal(toSerialize)
}

type NullableBTSyncAppElementParams struct {
	value *BTSyncAppElementParams
	isSet bool
}

func (v NullableBTSyncAppElementParams) Get() *BTSyncAppElementParams {
	return v.value
}

func (v *NullableBTSyncAppElementParams) Set(val *BTSyncAppElementParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSyncAppElementParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSyncAppElementParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSyncAppElementParams(val *BTSyncAppElementParams) *NullableBTSyncAppElementParams {
	return &NullableBTSyncAppElementParams{value: val, isSet: true}
}

func (v NullableBTSyncAppElementParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSyncAppElementParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
