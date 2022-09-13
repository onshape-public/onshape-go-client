/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6438-5fd2d9755d52
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// NextRecurringCharge struct for NextRecurringCharge
type NextRecurringCharge struct {
	Amount *int64  `json:"amount,omitempty"`
	Date   *string `json:"date,omitempty"`
}

// NewNextRecurringCharge instantiates a new NextRecurringCharge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNextRecurringCharge() *NextRecurringCharge {
	this := NextRecurringCharge{}
	return &this
}

// NewNextRecurringChargeWithDefaults instantiates a new NextRecurringCharge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNextRecurringChargeWithDefaults() *NextRecurringCharge {
	this := NextRecurringCharge{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *NextRecurringCharge) GetAmount() int64 {
	if o == nil || o.Amount == nil {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextRecurringCharge) GetAmountOk() (*int64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *NextRecurringCharge) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *NextRecurringCharge) SetAmount(v int64) {
	o.Amount = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *NextRecurringCharge) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextRecurringCharge) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *NextRecurringCharge) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *NextRecurringCharge) SetDate(v string) {
	o.Date = &v
}

func (o NextRecurringCharge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	return json.Marshal(toSerialize)
}

type NullableNextRecurringCharge struct {
	value *NextRecurringCharge
	isSet bool
}

func (v NullableNextRecurringCharge) Get() *NextRecurringCharge {
	return v.value
}

func (v *NullableNextRecurringCharge) Set(val *NextRecurringCharge) {
	v.value = val
	v.isSet = true
}

func (v NullableNextRecurringCharge) IsSet() bool {
	return v.isSet
}

func (v *NullableNextRecurringCharge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNextRecurringCharge(val *NextRecurringCharge) *NullableNextRecurringCharge {
	return &NullableNextRecurringCharge{value: val, isSet: true}
}

func (v NullableNextRecurringCharge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNextRecurringCharge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
