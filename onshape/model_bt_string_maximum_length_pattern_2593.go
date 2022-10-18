/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6946-bb9367a9d0cc
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTStringMaximumLengthPattern2593 struct for BTStringMaximumLengthPattern2593
type BTStringMaximumLengthPattern2593 struct {
	BtType                        *string `json:"btType,omitempty"`
	ErrorMessage                  *string `json:"errorMessage,omitempty"`
	ShouldResetValueWhenConfirmed *bool   `json:"shouldResetValueWhenConfirmed,omitempty"`
	MaximumLength                 *int32  `json:"maximumLength,omitempty"`
}

// NewBTStringMaximumLengthPattern2593 instantiates a new BTStringMaximumLengthPattern2593 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTStringMaximumLengthPattern2593() *BTStringMaximumLengthPattern2593 {
	this := BTStringMaximumLengthPattern2593{}
	return &this
}

// NewBTStringMaximumLengthPattern2593WithDefaults instantiates a new BTStringMaximumLengthPattern2593 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTStringMaximumLengthPattern2593WithDefaults() *BTStringMaximumLengthPattern2593 {
	this := BTStringMaximumLengthPattern2593{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTStringMaximumLengthPattern2593) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringMaximumLengthPattern2593) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTStringMaximumLengthPattern2593) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTStringMaximumLengthPattern2593) SetBtType(v string) {
	o.BtType = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *BTStringMaximumLengthPattern2593) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringMaximumLengthPattern2593) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *BTStringMaximumLengthPattern2593) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *BTStringMaximumLengthPattern2593) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetShouldResetValueWhenConfirmed returns the ShouldResetValueWhenConfirmed field value if set, zero value otherwise.
func (o *BTStringMaximumLengthPattern2593) GetShouldResetValueWhenConfirmed() bool {
	if o == nil || o.ShouldResetValueWhenConfirmed == nil {
		var ret bool
		return ret
	}
	return *o.ShouldResetValueWhenConfirmed
}

// GetShouldResetValueWhenConfirmedOk returns a tuple with the ShouldResetValueWhenConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringMaximumLengthPattern2593) GetShouldResetValueWhenConfirmedOk() (*bool, bool) {
	if o == nil || o.ShouldResetValueWhenConfirmed == nil {
		return nil, false
	}
	return o.ShouldResetValueWhenConfirmed, true
}

// HasShouldResetValueWhenConfirmed returns a boolean if a field has been set.
func (o *BTStringMaximumLengthPattern2593) HasShouldResetValueWhenConfirmed() bool {
	if o != nil && o.ShouldResetValueWhenConfirmed != nil {
		return true
	}

	return false
}

// SetShouldResetValueWhenConfirmed gets a reference to the given bool and assigns it to the ShouldResetValueWhenConfirmed field.
func (o *BTStringMaximumLengthPattern2593) SetShouldResetValueWhenConfirmed(v bool) {
	o.ShouldResetValueWhenConfirmed = &v
}

// GetMaximumLength returns the MaximumLength field value if set, zero value otherwise.
func (o *BTStringMaximumLengthPattern2593) GetMaximumLength() int32 {
	if o == nil || o.MaximumLength == nil {
		var ret int32
		return ret
	}
	return *o.MaximumLength
}

// GetMaximumLengthOk returns a tuple with the MaximumLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringMaximumLengthPattern2593) GetMaximumLengthOk() (*int32, bool) {
	if o == nil || o.MaximumLength == nil {
		return nil, false
	}
	return o.MaximumLength, true
}

// HasMaximumLength returns a boolean if a field has been set.
func (o *BTStringMaximumLengthPattern2593) HasMaximumLength() bool {
	if o != nil && o.MaximumLength != nil {
		return true
	}

	return false
}

// SetMaximumLength gets a reference to the given int32 and assigns it to the MaximumLength field.
func (o *BTStringMaximumLengthPattern2593) SetMaximumLength(v int32) {
	o.MaximumLength = &v
}

func (o BTStringMaximumLengthPattern2593) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.ShouldResetValueWhenConfirmed != nil {
		toSerialize["shouldResetValueWhenConfirmed"] = o.ShouldResetValueWhenConfirmed
	}
	if o.MaximumLength != nil {
		toSerialize["maximumLength"] = o.MaximumLength
	}
	return json.Marshal(toSerialize)
}

type NullableBTStringMaximumLengthPattern2593 struct {
	value *BTStringMaximumLengthPattern2593
	isSet bool
}

func (v NullableBTStringMaximumLengthPattern2593) Get() *BTStringMaximumLengthPattern2593 {
	return v.value
}

func (v *NullableBTStringMaximumLengthPattern2593) Set(val *BTStringMaximumLengthPattern2593) {
	v.value = val
	v.isSet = true
}

func (v NullableBTStringMaximumLengthPattern2593) IsSet() bool {
	return v.isSet
}

func (v *NullableBTStringMaximumLengthPattern2593) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTStringMaximumLengthPattern2593(val *BTStringMaximumLengthPattern2593) *NullableBTStringMaximumLengthPattern2593 {
	return &NullableBTStringMaximumLengthPattern2593{value: val, isSet: true}
}

func (v NullableBTStringMaximumLengthPattern2593) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTStringMaximumLengthPattern2593) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
