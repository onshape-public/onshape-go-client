/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.158.9878-ff51e1211d95
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCardInfo struct for BTCardInfo
type BTCardInfo struct {
	BillingAddress *BTAddressInfo `json:"billingAddress,omitempty"`
	ExpMonth       *int32         `json:"expMonth,omitempty"`
	ExpYear        *int32         `json:"expYear,omitempty"`
	Last4          *string        `json:"last4,omitempty"`
	Type           *string        `json:"type,omitempty"`
}

// NewBTCardInfo instantiates a new BTCardInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCardInfo() *BTCardInfo {
	this := BTCardInfo{}
	return &this
}

// NewBTCardInfoWithDefaults instantiates a new BTCardInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCardInfoWithDefaults() *BTCardInfo {
	this := BTCardInfo{}
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *BTCardInfo) GetBillingAddress() BTAddressInfo {
	if o == nil || o.BillingAddress == nil {
		var ret BTAddressInfo
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCardInfo) GetBillingAddressOk() (*BTAddressInfo, bool) {
	if o == nil || o.BillingAddress == nil {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *BTCardInfo) HasBillingAddress() bool {
	if o != nil && o.BillingAddress != nil {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given BTAddressInfo and assigns it to the BillingAddress field.
func (o *BTCardInfo) SetBillingAddress(v BTAddressInfo) {
	o.BillingAddress = &v
}

// GetExpMonth returns the ExpMonth field value if set, zero value otherwise.
func (o *BTCardInfo) GetExpMonth() int32 {
	if o == nil || o.ExpMonth == nil {
		var ret int32
		return ret
	}
	return *o.ExpMonth
}

// GetExpMonthOk returns a tuple with the ExpMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCardInfo) GetExpMonthOk() (*int32, bool) {
	if o == nil || o.ExpMonth == nil {
		return nil, false
	}
	return o.ExpMonth, true
}

// HasExpMonth returns a boolean if a field has been set.
func (o *BTCardInfo) HasExpMonth() bool {
	if o != nil && o.ExpMonth != nil {
		return true
	}

	return false
}

// SetExpMonth gets a reference to the given int32 and assigns it to the ExpMonth field.
func (o *BTCardInfo) SetExpMonth(v int32) {
	o.ExpMonth = &v
}

// GetExpYear returns the ExpYear field value if set, zero value otherwise.
func (o *BTCardInfo) GetExpYear() int32 {
	if o == nil || o.ExpYear == nil {
		var ret int32
		return ret
	}
	return *o.ExpYear
}

// GetExpYearOk returns a tuple with the ExpYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCardInfo) GetExpYearOk() (*int32, bool) {
	if o == nil || o.ExpYear == nil {
		return nil, false
	}
	return o.ExpYear, true
}

// HasExpYear returns a boolean if a field has been set.
func (o *BTCardInfo) HasExpYear() bool {
	if o != nil && o.ExpYear != nil {
		return true
	}

	return false
}

// SetExpYear gets a reference to the given int32 and assigns it to the ExpYear field.
func (o *BTCardInfo) SetExpYear(v int32) {
	o.ExpYear = &v
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *BTCardInfo) GetLast4() string {
	if o == nil || o.Last4 == nil {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCardInfo) GetLast4Ok() (*string, bool) {
	if o == nil || o.Last4 == nil {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *BTCardInfo) HasLast4() bool {
	if o != nil && o.Last4 != nil {
		return true
	}

	return false
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *BTCardInfo) SetLast4(v string) {
	o.Last4 = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTCardInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCardInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTCardInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTCardInfo) SetType(v string) {
	o.Type = &v
}

func (o BTCardInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillingAddress != nil {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if o.ExpMonth != nil {
		toSerialize["expMonth"] = o.ExpMonth
	}
	if o.ExpYear != nil {
		toSerialize["expYear"] = o.ExpYear
	}
	if o.Last4 != nil {
		toSerialize["last4"] = o.Last4
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTCardInfo struct {
	value *BTCardInfo
	isSet bool
}

func (v NullableBTCardInfo) Get() *BTCardInfo {
	return v.value
}

func (v *NullableBTCardInfo) Set(val *BTCardInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCardInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCardInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCardInfo(val *BTCardInfo) *NullableBTCardInfo {
	return &NullableBTCardInfo{value: val, isSet: true}
}

func (v NullableBTCardInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCardInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
