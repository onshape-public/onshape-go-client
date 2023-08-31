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

// BTNotFilter165 struct for BTNotFilter165
type BTNotFilter165 struct {
	BtType  *string           `json:"btType,omitempty"`
	Operand *BTQueryFilter183 `json:"operand,omitempty"`
}

// NewBTNotFilter165 instantiates a new BTNotFilter165 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTNotFilter165() *BTNotFilter165 {
	this := BTNotFilter165{}
	return &this
}

// NewBTNotFilter165WithDefaults instantiates a new BTNotFilter165 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTNotFilter165WithDefaults() *BTNotFilter165 {
	this := BTNotFilter165{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTNotFilter165) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNotFilter165) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTNotFilter165) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTNotFilter165) SetBtType(v string) {
	o.BtType = &v
}

// GetOperand returns the Operand field value if set, zero value otherwise.
func (o *BTNotFilter165) GetOperand() BTQueryFilter183 {
	if o == nil || o.Operand == nil {
		var ret BTQueryFilter183
		return ret
	}
	return *o.Operand
}

// GetOperandOk returns a tuple with the Operand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNotFilter165) GetOperandOk() (*BTQueryFilter183, bool) {
	if o == nil || o.Operand == nil {
		return nil, false
	}
	return o.Operand, true
}

// HasOperand returns a boolean if a field has been set.
func (o *BTNotFilter165) HasOperand() bool {
	if o != nil && o.Operand != nil {
		return true
	}

	return false
}

// SetOperand gets a reference to the given BTQueryFilter183 and assigns it to the Operand field.
func (o *BTNotFilter165) SetOperand(v BTQueryFilter183) {
	o.Operand = &v
}

func (o BTNotFilter165) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Operand != nil {
		toSerialize["operand"] = o.Operand
	}
	return json.Marshal(toSerialize)
}

type NullableBTNotFilter165 struct {
	value *BTNotFilter165
	isSet bool
}

func (v NullableBTNotFilter165) Get() *BTNotFilter165 {
	return v.value
}

func (v *NullableBTNotFilter165) Set(val *BTNotFilter165) {
	v.value = val
	v.isSet = true
}

func (v NullableBTNotFilter165) IsSet() bool {
	return v.isSet
}

func (v *NullableBTNotFilter165) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTNotFilter165(val *BTNotFilter165) *NullableBTNotFilter165 {
	return &NullableBTNotFilter165{value: val, isSet: true}
}

func (v NullableBTNotFilter165) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTNotFilter165) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
