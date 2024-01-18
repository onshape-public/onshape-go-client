/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.175.29152-5591e93bd1c1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTTextObjectFilter1515 struct for BTTextObjectFilter1515
type BTTextObjectFilter1515 struct {
	BtType *string `json:"btType,omitempty"`
	IsText *bool   `json:"isText,omitempty"`
}

// NewBTTextObjectFilter1515 instantiates a new BTTextObjectFilter1515 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTextObjectFilter1515() *BTTextObjectFilter1515 {
	this := BTTextObjectFilter1515{}
	return &this
}

// NewBTTextObjectFilter1515WithDefaults instantiates a new BTTextObjectFilter1515 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTextObjectFilter1515WithDefaults() *BTTextObjectFilter1515 {
	this := BTTextObjectFilter1515{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTextObjectFilter1515) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTextObjectFilter1515) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTextObjectFilter1515) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTextObjectFilter1515) SetBtType(v string) {
	o.BtType = &v
}

// GetIsText returns the IsText field value if set, zero value otherwise.
func (o *BTTextObjectFilter1515) GetIsText() bool {
	if o == nil || o.IsText == nil {
		var ret bool
		return ret
	}
	return *o.IsText
}

// GetIsTextOk returns a tuple with the IsText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTextObjectFilter1515) GetIsTextOk() (*bool, bool) {
	if o == nil || o.IsText == nil {
		return nil, false
	}
	return o.IsText, true
}

// HasIsText returns a boolean if a field has been set.
func (o *BTTextObjectFilter1515) HasIsText() bool {
	if o != nil && o.IsText != nil {
		return true
	}

	return false
}

// SetIsText gets a reference to the given bool and assigns it to the IsText field.
func (o *BTTextObjectFilter1515) SetIsText(v bool) {
	o.IsText = &v
}

func (o BTTextObjectFilter1515) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsText != nil {
		toSerialize["isText"] = o.IsText
	}
	return json.Marshal(toSerialize)
}

type NullableBTTextObjectFilter1515 struct {
	value *BTTextObjectFilter1515
	isSet bool
}

func (v NullableBTTextObjectFilter1515) Get() *BTTextObjectFilter1515 {
	return v.value
}

func (v *NullableBTTextObjectFilter1515) Set(val *BTTextObjectFilter1515) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTextObjectFilter1515) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTextObjectFilter1515) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTextObjectFilter1515(val *BTTextObjectFilter1515) *NullableBTTextObjectFilter1515 {
	return &NullableBTTextObjectFilter1515{value: val, isSet: true}
}

func (v NullableBTTextObjectFilter1515) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTextObjectFilter1515) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
