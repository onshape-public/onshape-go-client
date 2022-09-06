/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6344-db89a80956dd
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Versions struct for Versions
type Versions struct {
	AvailableVersions []string `json:"availableVersions,omitempty"`
	SpecifiedVersion  *string  `json:"specifiedVersion,omitempty"`
}

// NewVersions instantiates a new Versions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersions() *Versions {
	this := Versions{}
	return &this
}

// NewVersionsWithDefaults instantiates a new Versions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionsWithDefaults() *Versions {
	this := Versions{}
	return &this
}

// GetAvailableVersions returns the AvailableVersions field value if set, zero value otherwise.
func (o *Versions) GetAvailableVersions() []string {
	if o == nil || o.AvailableVersions == nil {
		var ret []string
		return ret
	}
	return o.AvailableVersions
}

// GetAvailableVersionsOk returns a tuple with the AvailableVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Versions) GetAvailableVersionsOk() ([]string, bool) {
	if o == nil || o.AvailableVersions == nil {
		return nil, false
	}
	return o.AvailableVersions, true
}

// HasAvailableVersions returns a boolean if a field has been set.
func (o *Versions) HasAvailableVersions() bool {
	if o != nil && o.AvailableVersions != nil {
		return true
	}

	return false
}

// SetAvailableVersions gets a reference to the given []string and assigns it to the AvailableVersions field.
func (o *Versions) SetAvailableVersions(v []string) {
	o.AvailableVersions = v
}

// GetSpecifiedVersion returns the SpecifiedVersion field value if set, zero value otherwise.
func (o *Versions) GetSpecifiedVersion() string {
	if o == nil || o.SpecifiedVersion == nil {
		var ret string
		return ret
	}
	return *o.SpecifiedVersion
}

// GetSpecifiedVersionOk returns a tuple with the SpecifiedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Versions) GetSpecifiedVersionOk() (*string, bool) {
	if o == nil || o.SpecifiedVersion == nil {
		return nil, false
	}
	return o.SpecifiedVersion, true
}

// HasSpecifiedVersion returns a boolean if a field has been set.
func (o *Versions) HasSpecifiedVersion() bool {
	if o != nil && o.SpecifiedVersion != nil {
		return true
	}

	return false
}

// SetSpecifiedVersion gets a reference to the given string and assigns it to the SpecifiedVersion field.
func (o *Versions) SetSpecifiedVersion(v string) {
	o.SpecifiedVersion = &v
}

func (o Versions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvailableVersions != nil {
		toSerialize["availableVersions"] = o.AvailableVersions
	}
	if o.SpecifiedVersion != nil {
		toSerialize["specifiedVersion"] = o.SpecifiedVersion
	}
	return json.Marshal(toSerialize)
}

type NullableVersions struct {
	value *Versions
	isSet bool
}

func (v NullableVersions) Get() *Versions {
	return v.value
}

func (v *NullableVersions) Set(val *Versions) {
	v.value = val
	v.isSet = true
}

func (v NullableVersions) IsSet() bool {
	return v.isSet
}

func (v *NullableVersions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersions(val *Versions) *NullableVersions {
	return &NullableVersions{value: val, isSet: true}
}

func (v NullableVersions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
