/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23367-59c3c9f9feef
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSessionCredentialInfo struct for BTSessionCredentialInfo
type BTSessionCredentialInfo struct {
	Provider *string `json:"provider,omitempty"`
	Type     *string `json:"type,omitempty"`
}

// NewBTSessionCredentialInfo instantiates a new BTSessionCredentialInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSessionCredentialInfo() *BTSessionCredentialInfo {
	this := BTSessionCredentialInfo{}
	return &this
}

// NewBTSessionCredentialInfoWithDefaults instantiates a new BTSessionCredentialInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSessionCredentialInfoWithDefaults() *BTSessionCredentialInfo {
	this := BTSessionCredentialInfo{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *BTSessionCredentialInfo) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSessionCredentialInfo) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *BTSessionCredentialInfo) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *BTSessionCredentialInfo) SetProvider(v string) {
	o.Provider = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTSessionCredentialInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSessionCredentialInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTSessionCredentialInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTSessionCredentialInfo) SetType(v string) {
	o.Type = &v
}

func (o BTSessionCredentialInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTSessionCredentialInfo struct {
	value *BTSessionCredentialInfo
	isSet bool
}

func (v NullableBTSessionCredentialInfo) Get() *BTSessionCredentialInfo {
	return v.value
}

func (v *NullableBTSessionCredentialInfo) Set(val *BTSessionCredentialInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSessionCredentialInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSessionCredentialInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSessionCredentialInfo(val *BTSessionCredentialInfo) *NullableBTSessionCredentialInfo {
	return &NullableBTSessionCredentialInfo{value: val, isSet: true}
}

func (v NullableBTSessionCredentialInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSessionCredentialInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
