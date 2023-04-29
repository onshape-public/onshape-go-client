/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.163.15296-122c93d7dbb6
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTFSValueOther1124 struct for BTFSValueOther1124
type BTFSValueOther1124 struct {
	BtType  string    `json:"btType"`
	TypeTag *string   `json:"typeTag,omitempty"`
	Type    *GBTPType `json:"type,omitempty"`
}

// NewBTFSValueOther1124 instantiates a new BTFSValueOther1124 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValueOther1124(btType string) *BTFSValueOther1124 {
	this := BTFSValueOther1124{}
	this.BtType = btType
	return &this
}

// NewBTFSValueOther1124WithDefaults instantiates a new BTFSValueOther1124 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValueOther1124WithDefaults() *BTFSValueOther1124 {
	this := BTFSValueOther1124{}
	return &this
}

// GetBtType returns the BtType field value
func (o *BTFSValueOther1124) GetBtType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value
// and a boolean to check if the value has been set.
func (o *BTFSValueOther1124) GetBtTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BtType, true
}

// SetBtType sets field value
func (o *BTFSValueOther1124) SetBtType(v string) {
	o.BtType = v
}

// GetTypeTag returns the TypeTag field value if set, zero value otherwise.
func (o *BTFSValueOther1124) GetTypeTag() string {
	if o == nil || o.TypeTag == nil {
		var ret string
		return ret
	}
	return *o.TypeTag
}

// GetTypeTagOk returns a tuple with the TypeTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueOther1124) GetTypeTagOk() (*string, bool) {
	if o == nil || o.TypeTag == nil {
		return nil, false
	}
	return o.TypeTag, true
}

// HasTypeTag returns a boolean if a field has been set.
func (o *BTFSValueOther1124) HasTypeTag() bool {
	if o != nil && o.TypeTag != nil {
		return true
	}

	return false
}

// SetTypeTag gets a reference to the given string and assigns it to the TypeTag field.
func (o *BTFSValueOther1124) SetTypeTag(v string) {
	o.TypeTag = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTFSValueOther1124) GetType() GBTPType {
	if o == nil || o.Type == nil {
		var ret GBTPType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueOther1124) GetTypeOk() (*GBTPType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTFSValueOther1124) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given GBTPType and assigns it to the Type field.
func (o *BTFSValueOther1124) SetType(v GBTPType) {
	o.Type = &v
}

func (o BTFSValueOther1124) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["btType"] = o.BtType
	}
	if o.TypeTag != nil {
		toSerialize["typeTag"] = o.TypeTag
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSValueOther1124 struct {
	value *BTFSValueOther1124
	isSet bool
}

func (v NullableBTFSValueOther1124) Get() *BTFSValueOther1124 {
	return v.value
}

func (v *NullableBTFSValueOther1124) Set(val *BTFSValueOther1124) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValueOther1124) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValueOther1124) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValueOther1124(val *BTFSValueOther1124) *NullableBTFSValueOther1124 {
	return &NullableBTFSValueOther1124{value: val, isSet: true}
}

func (v NullableBTFSValueOther1124) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValueOther1124) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
