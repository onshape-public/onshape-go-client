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

// BTModifiableEntityOnlyFilter1593 struct for BTModifiableEntityOnlyFilter1593
type BTModifiableEntityOnlyFilter1593 struct {
	BtType         *string `json:"btType,omitempty"`
	ModifiableOnly *bool   `json:"modifiableOnly,omitempty"`
}

// NewBTModifiableEntityOnlyFilter1593 instantiates a new BTModifiableEntityOnlyFilter1593 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTModifiableEntityOnlyFilter1593() *BTModifiableEntityOnlyFilter1593 {
	this := BTModifiableEntityOnlyFilter1593{}
	return &this
}

// NewBTModifiableEntityOnlyFilter1593WithDefaults instantiates a new BTModifiableEntityOnlyFilter1593 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTModifiableEntityOnlyFilter1593WithDefaults() *BTModifiableEntityOnlyFilter1593 {
	this := BTModifiableEntityOnlyFilter1593{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTModifiableEntityOnlyFilter1593) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTModifiableEntityOnlyFilter1593) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTModifiableEntityOnlyFilter1593) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTModifiableEntityOnlyFilter1593) SetBtType(v string) {
	o.BtType = &v
}

// GetModifiableOnly returns the ModifiableOnly field value if set, zero value otherwise.
func (o *BTModifiableEntityOnlyFilter1593) GetModifiableOnly() bool {
	if o == nil || o.ModifiableOnly == nil {
		var ret bool
		return ret
	}
	return *o.ModifiableOnly
}

// GetModifiableOnlyOk returns a tuple with the ModifiableOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTModifiableEntityOnlyFilter1593) GetModifiableOnlyOk() (*bool, bool) {
	if o == nil || o.ModifiableOnly == nil {
		return nil, false
	}
	return o.ModifiableOnly, true
}

// HasModifiableOnly returns a boolean if a field has been set.
func (o *BTModifiableEntityOnlyFilter1593) HasModifiableOnly() bool {
	if o != nil && o.ModifiableOnly != nil {
		return true
	}

	return false
}

// SetModifiableOnly gets a reference to the given bool and assigns it to the ModifiableOnly field.
func (o *BTModifiableEntityOnlyFilter1593) SetModifiableOnly(v bool) {
	o.ModifiableOnly = &v
}

func (o BTModifiableEntityOnlyFilter1593) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ModifiableOnly != nil {
		toSerialize["modifiableOnly"] = o.ModifiableOnly
	}
	return json.Marshal(toSerialize)
}

type NullableBTModifiableEntityOnlyFilter1593 struct {
	value *BTModifiableEntityOnlyFilter1593
	isSet bool
}

func (v NullableBTModifiableEntityOnlyFilter1593) Get() *BTModifiableEntityOnlyFilter1593 {
	return v.value
}

func (v *NullableBTModifiableEntityOnlyFilter1593) Set(val *BTModifiableEntityOnlyFilter1593) {
	v.value = val
	v.isSet = true
}

func (v NullableBTModifiableEntityOnlyFilter1593) IsSet() bool {
	return v.isSet
}

func (v *NullableBTModifiableEntityOnlyFilter1593) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTModifiableEntityOnlyFilter1593(val *BTModifiableEntityOnlyFilter1593) *NullableBTModifiableEntityOnlyFilter1593 {
	return &NullableBTModifiableEntityOnlyFilter1593{value: val, isSet: true}
}

func (v NullableBTModifiableEntityOnlyFilter1593) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTModifiableEntityOnlyFilter1593) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
