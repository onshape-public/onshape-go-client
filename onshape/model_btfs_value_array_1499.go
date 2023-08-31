/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.168.21279-402b6292597b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTFSValueArray1499 struct for BTFSValueArray1499
type BTFSValueArray1499 struct {
	BtType  *string         `json:"btType,omitempty"`
	TypeTag *string         `json:"typeTag,omitempty"`
	Value   []BTFSValue1888 `json:"value,omitempty"`
}

// NewBTFSValueArray1499 instantiates a new BTFSValueArray1499 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValueArray1499() *BTFSValueArray1499 {
	this := BTFSValueArray1499{}
	return &this
}

// NewBTFSValueArray1499WithDefaults instantiates a new BTFSValueArray1499 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValueArray1499WithDefaults() *BTFSValueArray1499 {
	this := BTFSValueArray1499{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFSValueArray1499) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueArray1499) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFSValueArray1499) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFSValueArray1499) SetBtType(v string) {
	o.BtType = &v
}

// GetTypeTag returns the TypeTag field value if set, zero value otherwise.
func (o *BTFSValueArray1499) GetTypeTag() string {
	if o == nil || o.TypeTag == nil {
		var ret string
		return ret
	}
	return *o.TypeTag
}

// GetTypeTagOk returns a tuple with the TypeTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueArray1499) GetTypeTagOk() (*string, bool) {
	if o == nil || o.TypeTag == nil {
		return nil, false
	}
	return o.TypeTag, true
}

// HasTypeTag returns a boolean if a field has been set.
func (o *BTFSValueArray1499) HasTypeTag() bool {
	if o != nil && o.TypeTag != nil {
		return true
	}

	return false
}

// SetTypeTag gets a reference to the given string and assigns it to the TypeTag field.
func (o *BTFSValueArray1499) SetTypeTag(v string) {
	o.TypeTag = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTFSValueArray1499) GetValue() []BTFSValue1888 {
	if o == nil || o.Value == nil {
		var ret []BTFSValue1888
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueArray1499) GetValueOk() ([]BTFSValue1888, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTFSValueArray1499) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []BTFSValue1888 and assigns it to the Value field.
func (o *BTFSValueArray1499) SetValue(v []BTFSValue1888) {
	o.Value = v
}

func (o BTFSValueArray1499) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.TypeTag != nil {
		toSerialize["typeTag"] = o.TypeTag
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSValueArray1499 struct {
	value *BTFSValueArray1499
	isSet bool
}

func (v NullableBTFSValueArray1499) Get() *BTFSValueArray1499 {
	return v.value
}

func (v *NullableBTFSValueArray1499) Set(val *BTFSValueArray1499) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValueArray1499) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValueArray1499) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValueArray1499(val *BTFSValueArray1499) *NullableBTFSValueArray1499 {
	return &NullableBTFSValueArray1499{value: val, isSet: true}
}

func (v NullableBTFSValueArray1499) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValueArray1499) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
