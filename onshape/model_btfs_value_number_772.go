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

// BTFSValueNumber772 struct for BTFSValueNumber772
type BTFSValueNumber772 struct {
	BTFSValue1888
	BtType  string   `json:"btType"`
	TypeTag *string  `json:"typeTag,omitempty"`
	Value   *float64 `json:"value,omitempty"`
}

// NewBTFSValueNumber772 instantiates a new BTFSValueNumber772 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValueNumber772(btType string) *BTFSValueNumber772 {
	this := BTFSValueNumber772{}
	this.BtType = btType
	return &this
}

// NewBTFSValueNumber772WithDefaults instantiates a new BTFSValueNumber772 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValueNumber772WithDefaults() *BTFSValueNumber772 {
	this := BTFSValueNumber772{}
	return &this
}

// GetBtType returns the BtType field value
func (o *BTFSValueNumber772) GetBtType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value
// and a boolean to check if the value has been set.
func (o *BTFSValueNumber772) GetBtTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BtType, true
}

// SetBtType sets field value
func (o *BTFSValueNumber772) SetBtType(v string) {
	o.BtType = v
}

// GetTypeTag returns the TypeTag field value if set, zero value otherwise.
func (o *BTFSValueNumber772) GetTypeTag() string {
	if o == nil || o.TypeTag == nil {
		var ret string
		return ret
	}
	return *o.TypeTag
}

// GetTypeTagOk returns a tuple with the TypeTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueNumber772) GetTypeTagOk() (*string, bool) {
	if o == nil || o.TypeTag == nil {
		return nil, false
	}
	return o.TypeTag, true
}

// HasTypeTag returns a boolean if a field has been set.
func (o *BTFSValueNumber772) HasTypeTag() bool {
	if o != nil && o.TypeTag != nil {
		return true
	}

	return false
}

// SetTypeTag gets a reference to the given string and assigns it to the TypeTag field.
func (o *BTFSValueNumber772) SetTypeTag(v string) {
	o.TypeTag = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTFSValueNumber772) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueNumber772) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTFSValueNumber772) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *BTFSValueNumber772) SetValue(v float64) {
	o.Value = &v
}

func (o BTFSValueNumber772) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTFSValue1888, errBTFSValue1888 := json.Marshal(o.BTFSValue1888)
	if errBTFSValue1888 != nil {
		return []byte{}, errBTFSValue1888
	}
	errBTFSValue1888 = json.Unmarshal([]byte(serializedBTFSValue1888), &toSerialize)
	if errBTFSValue1888 != nil {
		return []byte{}, errBTFSValue1888
	}
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

type NullableBTFSValueNumber772 struct {
	value *BTFSValueNumber772
	isSet bool
}

func (v NullableBTFSValueNumber772) Get() *BTFSValueNumber772 {
	return v.value
}

func (v *NullableBTFSValueNumber772) Set(val *BTFSValueNumber772) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValueNumber772) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValueNumber772) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValueNumber772(val *BTFSValueNumber772) *NullableBTFSValueNumber772 {
	return &NullableBTFSValueNumber772{value: val, isSet: true}
}

func (v NullableBTFSValueNumber772) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValueNumber772) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
