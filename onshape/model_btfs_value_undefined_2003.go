/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25478-d4e5ab4765a4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTFSValueUndefined2003 struct for BTFSValueUndefined2003
type BTFSValueUndefined2003 struct {
	BtType  string  `json:"btType"`
	TypeTag *string `json:"typeTag,omitempty"`
}

// NewBTFSValueUndefined2003 instantiates a new BTFSValueUndefined2003 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValueUndefined2003(btType string) *BTFSValueUndefined2003 {
	this := BTFSValueUndefined2003{}
	this.BtType = btType
	return &this
}

// NewBTFSValueUndefined2003WithDefaults instantiates a new BTFSValueUndefined2003 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValueUndefined2003WithDefaults() *BTFSValueUndefined2003 {
	this := BTFSValueUndefined2003{}
	return &this
}

// GetBtType returns the BtType field value
func (o *BTFSValueUndefined2003) GetBtType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value
// and a boolean to check if the value has been set.
func (o *BTFSValueUndefined2003) GetBtTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BtType, true
}

// SetBtType sets field value
func (o *BTFSValueUndefined2003) SetBtType(v string) {
	o.BtType = v
}

// GetTypeTag returns the TypeTag field value if set, zero value otherwise.
func (o *BTFSValueUndefined2003) GetTypeTag() string {
	if o == nil || o.TypeTag == nil {
		var ret string
		return ret
	}
	return *o.TypeTag
}

// GetTypeTagOk returns a tuple with the TypeTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueUndefined2003) GetTypeTagOk() (*string, bool) {
	if o == nil || o.TypeTag == nil {
		return nil, false
	}
	return o.TypeTag, true
}

// HasTypeTag returns a boolean if a field has been set.
func (o *BTFSValueUndefined2003) HasTypeTag() bool {
	if o != nil && o.TypeTag != nil {
		return true
	}

	return false
}

// SetTypeTag gets a reference to the given string and assigns it to the TypeTag field.
func (o *BTFSValueUndefined2003) SetTypeTag(v string) {
	o.TypeTag = &v
}

func (o BTFSValueUndefined2003) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["btType"] = o.BtType
	}
	if o.TypeTag != nil {
		toSerialize["typeTag"] = o.TypeTag
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSValueUndefined2003 struct {
	value *BTFSValueUndefined2003
	isSet bool
}

func (v NullableBTFSValueUndefined2003) Get() *BTFSValueUndefined2003 {
	return v.value
}

func (v *NullableBTFSValueUndefined2003) Set(val *BTFSValueUndefined2003) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValueUndefined2003) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValueUndefined2003) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValueUndefined2003(val *BTFSValueUndefined2003) *NullableBTFSValueUndefined2003 {
	return &NullableBTFSValueUndefined2003{value: val, isSet: true}
}

func (v NullableBTFSValueUndefined2003) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValueUndefined2003) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
