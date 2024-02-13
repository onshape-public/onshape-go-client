/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSplineHandleFilter2971 struct for BTSplineHandleFilter2971
type BTSplineHandleFilter2971 struct {
	BtType             *string `json:"btType,omitempty"`
	AllowsSplineHandle *bool   `json:"allowsSplineHandle,omitempty"`
}

// NewBTSplineHandleFilter2971 instantiates a new BTSplineHandleFilter2971 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSplineHandleFilter2971() *BTSplineHandleFilter2971 {
	this := BTSplineHandleFilter2971{}
	return &this
}

// NewBTSplineHandleFilter2971WithDefaults instantiates a new BTSplineHandleFilter2971 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSplineHandleFilter2971WithDefaults() *BTSplineHandleFilter2971 {
	this := BTSplineHandleFilter2971{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSplineHandleFilter2971) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSplineHandleFilter2971) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSplineHandleFilter2971) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSplineHandleFilter2971) SetBtType(v string) {
	o.BtType = &v
}

// GetAllowsSplineHandle returns the AllowsSplineHandle field value if set, zero value otherwise.
func (o *BTSplineHandleFilter2971) GetAllowsSplineHandle() bool {
	if o == nil || o.AllowsSplineHandle == nil {
		var ret bool
		return ret
	}
	return *o.AllowsSplineHandle
}

// GetAllowsSplineHandleOk returns a tuple with the AllowsSplineHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSplineHandleFilter2971) GetAllowsSplineHandleOk() (*bool, bool) {
	if o == nil || o.AllowsSplineHandle == nil {
		return nil, false
	}
	return o.AllowsSplineHandle, true
}

// HasAllowsSplineHandle returns a boolean if a field has been set.
func (o *BTSplineHandleFilter2971) HasAllowsSplineHandle() bool {
	if o != nil && o.AllowsSplineHandle != nil {
		return true
	}

	return false
}

// SetAllowsSplineHandle gets a reference to the given bool and assigns it to the AllowsSplineHandle field.
func (o *BTSplineHandleFilter2971) SetAllowsSplineHandle(v bool) {
	o.AllowsSplineHandle = &v
}

func (o BTSplineHandleFilter2971) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.AllowsSplineHandle != nil {
		toSerialize["allowsSplineHandle"] = o.AllowsSplineHandle
	}
	return json.Marshal(toSerialize)
}

type NullableBTSplineHandleFilter2971 struct {
	value *BTSplineHandleFilter2971
	isSet bool
}

func (v NullableBTSplineHandleFilter2971) Get() *BTSplineHandleFilter2971 {
	return v.value
}

func (v *NullableBTSplineHandleFilter2971) Set(val *BTSplineHandleFilter2971) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSplineHandleFilter2971) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSplineHandleFilter2971) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSplineHandleFilter2971(val *BTSplineHandleFilter2971) *NullableBTSplineHandleFilter2971 {
	return &NullableBTSplineHandleFilter2971{value: val, isSet: true}
}

func (v NullableBTSplineHandleFilter2971) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSplineHandleFilter2971) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
