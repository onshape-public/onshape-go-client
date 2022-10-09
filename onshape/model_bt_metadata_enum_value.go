/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6826-13f16e212af4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMetadataEnumValue struct for BTMetadataEnumValue
type BTMetadataEnumValue struct {
	Label *string `json:"label,omitempty"`
	State *int32  `json:"state,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewBTMetadataEnumValue instantiates a new BTMetadataEnumValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMetadataEnumValue() *BTMetadataEnumValue {
	this := BTMetadataEnumValue{}
	return &this
}

// NewBTMetadataEnumValueWithDefaults instantiates a new BTMetadataEnumValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMetadataEnumValueWithDefaults() *BTMetadataEnumValue {
	this := BTMetadataEnumValue{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BTMetadataEnumValue) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataEnumValue) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BTMetadataEnumValue) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BTMetadataEnumValue) SetLabel(v string) {
	o.Label = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTMetadataEnumValue) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataEnumValue) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTMetadataEnumValue) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTMetadataEnumValue) SetState(v int32) {
	o.State = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMetadataEnumValue) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataEnumValue) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTMetadataEnumValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BTMetadataEnumValue) SetValue(v string) {
	o.Value = &v
}

func (o BTMetadataEnumValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTMetadataEnumValue struct {
	value *BTMetadataEnumValue
	isSet bool
}

func (v NullableBTMetadataEnumValue) Get() *BTMetadataEnumValue {
	return v.value
}

func (v *NullableBTMetadataEnumValue) Set(val *BTMetadataEnumValue) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataEnumValue) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataEnumValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataEnumValue(val *BTMetadataEnumValue) *NullableBTMetadataEnumValue {
	return &NullableBTMetadataEnumValue{value: val, isSet: true}
}

func (v NullableBTMetadataEnumValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataEnumValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
