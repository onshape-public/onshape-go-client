/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6017-d06351faf6f2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTNextPartNumbersParam struct for BTNextPartNumbersParam
type BTNextPartNumbersParam struct {
	ItemPartNumbers []BTNextPartNumberParam `json:"itemPartNumbers,omitempty"`
	SkipPartNumbers []string                `json:"skipPartNumbers,omitempty"`
}

// NewBTNextPartNumbersParam instantiates a new BTNextPartNumbersParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTNextPartNumbersParam() *BTNextPartNumbersParam {
	this := BTNextPartNumbersParam{}
	return &this
}

// NewBTNextPartNumbersParamWithDefaults instantiates a new BTNextPartNumbersParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTNextPartNumbersParamWithDefaults() *BTNextPartNumbersParam {
	this := BTNextPartNumbersParam{}
	return &this
}

// GetItemPartNumbers returns the ItemPartNumbers field value if set, zero value otherwise.
func (o *BTNextPartNumbersParam) GetItemPartNumbers() []BTNextPartNumberParam {
	if o == nil || o.ItemPartNumbers == nil {
		var ret []BTNextPartNumberParam
		return ret
	}
	return o.ItemPartNumbers
}

// GetItemPartNumbersOk returns a tuple with the ItemPartNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNextPartNumbersParam) GetItemPartNumbersOk() ([]BTNextPartNumberParam, bool) {
	if o == nil || o.ItemPartNumbers == nil {
		return nil, false
	}
	return o.ItemPartNumbers, true
}

// HasItemPartNumbers returns a boolean if a field has been set.
func (o *BTNextPartNumbersParam) HasItemPartNumbers() bool {
	if o != nil && o.ItemPartNumbers != nil {
		return true
	}

	return false
}

// SetItemPartNumbers gets a reference to the given []BTNextPartNumberParam and assigns it to the ItemPartNumbers field.
func (o *BTNextPartNumbersParam) SetItemPartNumbers(v []BTNextPartNumberParam) {
	o.ItemPartNumbers = v
}

// GetSkipPartNumbers returns the SkipPartNumbers field value if set, zero value otherwise.
func (o *BTNextPartNumbersParam) GetSkipPartNumbers() []string {
	if o == nil || o.SkipPartNumbers == nil {
		var ret []string
		return ret
	}
	return o.SkipPartNumbers
}

// GetSkipPartNumbersOk returns a tuple with the SkipPartNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNextPartNumbersParam) GetSkipPartNumbersOk() ([]string, bool) {
	if o == nil || o.SkipPartNumbers == nil {
		return nil, false
	}
	return o.SkipPartNumbers, true
}

// HasSkipPartNumbers returns a boolean if a field has been set.
func (o *BTNextPartNumbersParam) HasSkipPartNumbers() bool {
	if o != nil && o.SkipPartNumbers != nil {
		return true
	}

	return false
}

// SetSkipPartNumbers gets a reference to the given []string and assigns it to the SkipPartNumbers field.
func (o *BTNextPartNumbersParam) SetSkipPartNumbers(v []string) {
	o.SkipPartNumbers = v
}

func (o BTNextPartNumbersParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ItemPartNumbers != nil {
		toSerialize["itemPartNumbers"] = o.ItemPartNumbers
	}
	if o.SkipPartNumbers != nil {
		toSerialize["skipPartNumbers"] = o.SkipPartNumbers
	}
	return json.Marshal(toSerialize)
}

type NullableBTNextPartNumbersParam struct {
	value *BTNextPartNumbersParam
	isSet bool
}

func (v NullableBTNextPartNumbersParam) Get() *BTNextPartNumbersParam {
	return v.value
}

func (v *NullableBTNextPartNumbersParam) Set(val *BTNextPartNumbersParam) {
	v.value = val
	v.isSet = true
}

func (v NullableBTNextPartNumbersParam) IsSet() bool {
	return v.isSet
}

func (v *NullableBTNextPartNumbersParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTNextPartNumbersParam(val *BTNextPartNumbersParam) *NullableBTNextPartNumbersParam {
	return &NullableBTNextPartNumbersParam{value: val, isSet: true}
}

func (v NullableBTNextPartNumbersParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTNextPartNumbersParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
