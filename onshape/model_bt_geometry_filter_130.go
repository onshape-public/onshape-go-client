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

// BTGeometryFilter130 struct for BTGeometryFilter130
type BTGeometryFilter130 struct {
	BtType       *string          `json:"btType,omitempty"`
	GeometryType *GBTGeometryType `json:"geometryType,omitempty"`
}

// NewBTGeometryFilter130 instantiates a new BTGeometryFilter130 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTGeometryFilter130() *BTGeometryFilter130 {
	this := BTGeometryFilter130{}
	return &this
}

// NewBTGeometryFilter130WithDefaults instantiates a new BTGeometryFilter130 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTGeometryFilter130WithDefaults() *BTGeometryFilter130 {
	this := BTGeometryFilter130{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTGeometryFilter130) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryFilter130) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTGeometryFilter130) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTGeometryFilter130) SetBtType(v string) {
	o.BtType = &v
}

// GetGeometryType returns the GeometryType field value if set, zero value otherwise.
func (o *BTGeometryFilter130) GetGeometryType() GBTGeometryType {
	if o == nil || o.GeometryType == nil {
		var ret GBTGeometryType
		return ret
	}
	return *o.GeometryType
}

// GetGeometryTypeOk returns a tuple with the GeometryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryFilter130) GetGeometryTypeOk() (*GBTGeometryType, bool) {
	if o == nil || o.GeometryType == nil {
		return nil, false
	}
	return o.GeometryType, true
}

// HasGeometryType returns a boolean if a field has been set.
func (o *BTGeometryFilter130) HasGeometryType() bool {
	if o != nil && o.GeometryType != nil {
		return true
	}

	return false
}

// SetGeometryType gets a reference to the given GBTGeometryType and assigns it to the GeometryType field.
func (o *BTGeometryFilter130) SetGeometryType(v GBTGeometryType) {
	o.GeometryType = &v
}

func (o BTGeometryFilter130) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.GeometryType != nil {
		toSerialize["geometryType"] = o.GeometryType
	}
	return json.Marshal(toSerialize)
}

type NullableBTGeometryFilter130 struct {
	value *BTGeometryFilter130
	isSet bool
}

func (v NullableBTGeometryFilter130) Get() *BTGeometryFilter130 {
	return v.value
}

func (v *NullableBTGeometryFilter130) Set(val *BTGeometryFilter130) {
	v.value = val
	v.isSet = true
}

func (v NullableBTGeometryFilter130) IsSet() bool {
	return v.isSet
}

func (v *NullableBTGeometryFilter130) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTGeometryFilter130(val *BTGeometryFilter130) *NullableBTGeometryFilter130 {
	return &NullableBTGeometryFilter130{value: val, isSet: true}
}

func (v NullableBTGeometryFilter130) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTGeometryFilter130) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
