/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.171.24257-687de06de652
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Callback struct for Callback
type Callback struct {
	Empty      *bool                             `json:"empty,omitempty"`
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Getref     *string                           `json:"get$ref,omitempty"`
}

// NewCallback instantiates a new Callback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallback() *Callback {
	this := Callback{}
	return &this
}

// NewCallbackWithDefaults instantiates a new Callback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallbackWithDefaults() *Callback {
	this := Callback{}
	return &this
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *Callback) GetEmpty() bool {
	if o == nil || o.Empty == nil {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Callback) GetEmptyOk() (*bool, bool) {
	if o == nil || o.Empty == nil {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *Callback) HasEmpty() bool {
	if o != nil && o.Empty != nil {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *Callback) SetEmpty(v bool) {
	o.Empty = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Callback) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Callback) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Callback) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Callback) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetGetref returns the Getref field value if set, zero value otherwise.
func (o *Callback) GetGetref() string {
	if o == nil || o.Getref == nil {
		var ret string
		return ret
	}
	return *o.Getref
}

// GetGetrefOk returns a tuple with the Getref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Callback) GetGetrefOk() (*string, bool) {
	if o == nil || o.Getref == nil {
		return nil, false
	}
	return o.Getref, true
}

// HasGetref returns a boolean if a field has been set.
func (o *Callback) HasGetref() bool {
	if o != nil && o.Getref != nil {
		return true
	}

	return false
}

// SetGetref gets a reference to the given string and assigns it to the Getref field.
func (o *Callback) SetGetref(v string) {
	o.Getref = &v
}

func (o Callback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Empty != nil {
		toSerialize["empty"] = o.Empty
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Getref != nil {
		toSerialize["get$ref"] = o.Getref
	}
	return json.Marshal(toSerialize)
}

type NullableCallback struct {
	value *Callback
	isSet bool
}

func (v NullableCallback) Get() *Callback {
	return v.value
}

func (v *NullableCallback) Set(val *Callback) {
	v.value = val
	v.isSet = true
}

func (v NullableCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallback(val *Callback) *NullableCallback {
	return &NullableCallback{value: val, isSet: true}
}

func (v NullableCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
