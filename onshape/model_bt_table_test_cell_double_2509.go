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

// BTTableTestCellDouble2509 struct for BTTableTestCellDouble2509
type BTTableTestCellDouble2509 struct {
	BtType        *string  `json:"btType,omitempty"`
	IsEverVisible *bool    `json:"isEverVisible,omitempty"`
	IsReadOnly    *bool    `json:"isReadOnly,omitempty"`
	CellValue     *float64 `json:"cellValue,omitempty"`
}

// NewBTTableTestCellDouble2509 instantiates a new BTTableTestCellDouble2509 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableTestCellDouble2509() *BTTableTestCellDouble2509 {
	this := BTTableTestCellDouble2509{}
	return &this
}

// NewBTTableTestCellDouble2509WithDefaults instantiates a new BTTableTestCellDouble2509 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableTestCellDouble2509WithDefaults() *BTTableTestCellDouble2509 {
	this := BTTableTestCellDouble2509{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableTestCellDouble2509) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableTestCellDouble2509) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTableTestCellDouble2509) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTableTestCellDouble2509) SetBtType(v string) {
	o.BtType = &v
}

// GetIsEverVisible returns the IsEverVisible field value if set, zero value otherwise.
func (o *BTTableTestCellDouble2509) GetIsEverVisible() bool {
	if o == nil || o.IsEverVisible == nil {
		var ret bool
		return ret
	}
	return *o.IsEverVisible
}

// GetIsEverVisibleOk returns a tuple with the IsEverVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableTestCellDouble2509) GetIsEverVisibleOk() (*bool, bool) {
	if o == nil || o.IsEverVisible == nil {
		return nil, false
	}
	return o.IsEverVisible, true
}

// HasIsEverVisible returns a boolean if a field has been set.
func (o *BTTableTestCellDouble2509) HasIsEverVisible() bool {
	if o != nil && o.IsEverVisible != nil {
		return true
	}

	return false
}

// SetIsEverVisible gets a reference to the given bool and assigns it to the IsEverVisible field.
func (o *BTTableTestCellDouble2509) SetIsEverVisible(v bool) {
	o.IsEverVisible = &v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *BTTableTestCellDouble2509) GetIsReadOnly() bool {
	if o == nil || o.IsReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableTestCellDouble2509) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || o.IsReadOnly == nil {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *BTTableTestCellDouble2509) HasIsReadOnly() bool {
	if o != nil && o.IsReadOnly != nil {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *BTTableTestCellDouble2509) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

// GetCellValue returns the CellValue field value if set, zero value otherwise.
func (o *BTTableTestCellDouble2509) GetCellValue() float64 {
	if o == nil || o.CellValue == nil {
		var ret float64
		return ret
	}
	return *o.CellValue
}

// GetCellValueOk returns a tuple with the CellValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableTestCellDouble2509) GetCellValueOk() (*float64, bool) {
	if o == nil || o.CellValue == nil {
		return nil, false
	}
	return o.CellValue, true
}

// HasCellValue returns a boolean if a field has been set.
func (o *BTTableTestCellDouble2509) HasCellValue() bool {
	if o != nil && o.CellValue != nil {
		return true
	}

	return false
}

// SetCellValue gets a reference to the given float64 and assigns it to the CellValue field.
func (o *BTTableTestCellDouble2509) SetCellValue(v float64) {
	o.CellValue = &v
}

func (o BTTableTestCellDouble2509) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsEverVisible != nil {
		toSerialize["isEverVisible"] = o.IsEverVisible
	}
	if o.IsReadOnly != nil {
		toSerialize["isReadOnly"] = o.IsReadOnly
	}
	if o.CellValue != nil {
		toSerialize["cellValue"] = o.CellValue
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableTestCellDouble2509 struct {
	value *BTTableTestCellDouble2509
	isSet bool
}

func (v NullableBTTableTestCellDouble2509) Get() *BTTableTestCellDouble2509 {
	return v.value
}

func (v *NullableBTTableTestCellDouble2509) Set(val *BTTableTestCellDouble2509) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableTestCellDouble2509) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableTestCellDouble2509) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableTestCellDouble2509(val *BTTableTestCellDouble2509) *NullableBTTableTestCellDouble2509 {
	return &NullableBTTableTestCellDouble2509{value: val, isSet: true}
}

func (v NullableBTTableTestCellDouble2509) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableTestCellDouble2509) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
