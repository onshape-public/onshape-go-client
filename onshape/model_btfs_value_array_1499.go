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

// BTFSValueArray1499 struct for BTFSValueArray1499
type BTFSValueArray1499 struct {
	BtType  string          `json:"btType"`
	TypeTag *string         `json:"typeTag,omitempty"`
	Value   []BTFSValue1888 `json:"value,omitempty"`
}

// NewBTFSValueArray1499 instantiates a new BTFSValueArray1499 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValueArray1499(btType string) *BTFSValueArray1499 {
	this := BTFSValueArray1499{}
	this.BtType = btType
	return &this
}

// NewBTFSValueArray1499WithDefaults instantiates a new BTFSValueArray1499 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValueArray1499WithDefaults() *BTFSValueArray1499 {
	this := BTFSValueArray1499{}
	return &this
}

// GetBtType returns the BtType field value
func (o *BTFSValueArray1499) GetBtType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value
// and a boolean to check if the value has been set.
func (o *BTFSValueArray1499) GetBtTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BtType, true
}

// SetBtType sets field value
func (o *BTFSValueArray1499) SetBtType(v string) {
	o.BtType = v
}

// GetTypeTag returns the TypeTag field value if set, zero value otherwise.
func (o *BTFSValueArray1499) GetTypeTag() string {
	if o == nil || o.TypeTag == nil {
		var ret string
		return ret
	}
	return *o.TypeTag
}

// GetTypeTagOk returns a tuple with the TypeTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueArray1499) GetTypeTagOk() (*string, bool) {
	if o == nil || o.TypeTag == nil {
		return nil, false
	}
	return o.TypeTag, true
}

// HasTypeTag returns a boolean if a field has been set.
func (o *BTFSValueArray1499) HasTypeTag() bool {
	if o != nil && o.TypeTag != nil {
		return true
	}

	return false
}

// SetTypeTag gets a reference to the given string and assigns it to the TypeTag field.
func (o *BTFSValueArray1499) SetTypeTag(v string) {
	o.TypeTag = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTFSValueArray1499) GetValue() []BTFSValue1888 {
	if o == nil || o.Value == nil {
		var ret []BTFSValue1888
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueArray1499) GetValueOk() ([]BTFSValue1888, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTFSValueArray1499) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []BTFSValue1888 and assigns it to the Value field.
func (o *BTFSValueArray1499) SetValue(v []BTFSValue1888) {
	o.Value = v
}

func (o BTFSValueArray1499) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["btType"] = o.BtType
	}
	if o.TypeTag != nil {
		toSerialize["typeTag"] = o.TypeTag
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSValueArray1499 struct {
	value *BTFSValueArray1499
	isSet bool
}

func (v NullableBTFSValueArray1499) Get() *BTFSValueArray1499 {
	return v.value
}

func (v *NullableBTFSValueArray1499) Set(val *BTFSValueArray1499) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValueArray1499) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValueArray1499) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValueArray1499(val *BTFSValueArray1499) *NullableBTFSValueArray1499 {
	return &NullableBTFSValueArray1499{value: val, isSet: true}
}

func (v NullableBTFSValueArray1499) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValueArray1499) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
