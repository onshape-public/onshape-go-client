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

// BTJPathIndex1871 Identifies a value in a json array. For insert and move edit destinations, -1 can be used to indicate 'end'.
type BTJPathIndex1871 struct {
	BTJPathElement2297
	BtType *string `json:"btType,omitempty"`
	Index  *int32  `json:"index,omitempty"`
	// Type of JSON object.
	BtType *string `json:"btType,omitempty"`
}

// NewBTJPathIndex1871 instantiates a new BTJPathIndex1871 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTJPathIndex1871() *BTJPathIndex1871 {
	this := BTJPathIndex1871{}
	return &this
}

// NewBTJPathIndex1871WithDefaults instantiates a new BTJPathIndex1871 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTJPathIndex1871WithDefaults() *BTJPathIndex1871 {
	this := BTJPathIndex1871{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTJPathIndex1871) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJPathIndex1871) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTJPathIndex1871) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTJPathIndex1871) SetBtType(v string) {
	o.BtType = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *BTJPathIndex1871) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJPathIndex1871) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *BTJPathIndex1871) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *BTJPathIndex1871) SetIndex(v int32) {
	o.Index = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTJPathIndex1871) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJPathIndex1871) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTJPathIndex1871) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTJPathIndex1871) SetBtType(v string) {
	o.BtType = &v
}

func (o BTJPathIndex1871) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTJPathElement2297, errBTJPathElement2297 := json.Marshal(o.BTJPathElement2297)
	if errBTJPathElement2297 != nil {
		return []byte{}, errBTJPathElement2297
	}
	errBTJPathElement2297 = json.Unmarshal([]byte(serializedBTJPathElement2297), &toSerialize)
	if errBTJPathElement2297 != nil {
		return []byte{}, errBTJPathElement2297
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTJPathIndex1871 struct {
	value *BTJPathIndex1871
	isSet bool
}

func (v NullableBTJPathIndex1871) Get() *BTJPathIndex1871 {
	return v.value
}

func (v *NullableBTJPathIndex1871) Set(val *BTJPathIndex1871) {
	v.value = val
	v.isSet = true
}

func (v NullableBTJPathIndex1871) IsSet() bool {
	return v.isSet
}

func (v *NullableBTJPathIndex1871) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTJPathIndex1871(val *BTJPathIndex1871) *NullableBTJPathIndex1871 {
	return &NullableBTJPathIndex1871{value: val, isSet: true}
}

func (v NullableBTJPathIndex1871) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTJPathIndex1871) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
