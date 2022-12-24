/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.9191-43c781405890
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTFSValueMap2062 struct for BTFSValueMap2062
type BTFSValueMap2062 struct {
	TypeTag *string                 `json:"typeTag,omitempty"`
	Value   []BTFSValueMapEntry2077 `json:"value,omitempty"`
}

// NewBTFSValueMap2062 instantiates a new BTFSValueMap2062 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValueMap2062() *BTFSValueMap2062 {
	this := BTFSValueMap2062{}
	return &this
}

// NewBTFSValueMap2062WithDefaults instantiates a new BTFSValueMap2062 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValueMap2062WithDefaults() *BTFSValueMap2062 {
	this := BTFSValueMap2062{}
	return &this
}

// GetTypeTag returns the TypeTag field value if set, zero value otherwise.
func (o *BTFSValueMap2062) GetTypeTag() string {
	if o == nil || o.TypeTag == nil {
		var ret string
		return ret
	}
	return *o.TypeTag
}

// GetTypeTagOk returns a tuple with the TypeTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueMap2062) GetTypeTagOk() (*string, bool) {
	if o == nil || o.TypeTag == nil {
		return nil, false
	}
	return o.TypeTag, true
}

// HasTypeTag returns a boolean if a field has been set.
func (o *BTFSValueMap2062) HasTypeTag() bool {
	if o != nil && o.TypeTag != nil {
		return true
	}

	return false
}

// SetTypeTag gets a reference to the given string and assigns it to the TypeTag field.
func (o *BTFSValueMap2062) SetTypeTag(v string) {
	o.TypeTag = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTFSValueMap2062) GetValue() []BTFSValueMapEntry2077 {
	if o == nil || o.Value == nil {
		var ret []BTFSValueMapEntry2077
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueMap2062) GetValueOk() ([]BTFSValueMapEntry2077, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTFSValueMap2062) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []BTFSValueMapEntry2077 and assigns it to the Value field.
func (o *BTFSValueMap2062) SetValue(v []BTFSValueMapEntry2077) {
	o.Value = v
}

func (o BTFSValueMap2062) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TypeTag != nil {
		toSerialize["typeTag"] = o.TypeTag
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSValueMap2062 struct {
	value *BTFSValueMap2062
	isSet bool
}

func (v NullableBTFSValueMap2062) Get() *BTFSValueMap2062 {
	return v.value
}

func (v *NullableBTFSValueMap2062) Set(val *BTFSValueMap2062) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValueMap2062) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValueMap2062) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValueMap2062(val *BTFSValueMap2062) *NullableBTFSValueMap2062 {
	return &NullableBTFSValueMap2062{value: val, isSet: true}
}

func (v NullableBTFSValueMap2062) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValueMap2062) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
