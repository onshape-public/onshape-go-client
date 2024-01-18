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

// BTSketchArcDisplayData349 struct for BTSketchArcDisplayData349
type BTSketchArcDisplayData349 struct {
	BtType *string   `json:"btType,omitempty"`
	Points []float64 `json:"points,omitempty"`
}

// NewBTSketchArcDisplayData349 instantiates a new BTSketchArcDisplayData349 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSketchArcDisplayData349() *BTSketchArcDisplayData349 {
	this := BTSketchArcDisplayData349{}
	return &this
}

// NewBTSketchArcDisplayData349WithDefaults instantiates a new BTSketchArcDisplayData349 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSketchArcDisplayData349WithDefaults() *BTSketchArcDisplayData349 {
	this := BTSketchArcDisplayData349{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchArcDisplayData349) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchArcDisplayData349) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSketchArcDisplayData349) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSketchArcDisplayData349) SetBtType(v string) {
	o.BtType = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BTSketchArcDisplayData349) GetPoints() []float64 {
	if o == nil || o.Points == nil {
		var ret []float64
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchArcDisplayData349) GetPointsOk() ([]float64, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *BTSketchArcDisplayData349) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []float64 and assigns it to the Points field.
func (o *BTSketchArcDisplayData349) SetPoints(v []float64) {
	o.Points = v
}

func (o BTSketchArcDisplayData349) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	return json.Marshal(toSerialize)
}

type NullableBTSketchArcDisplayData349 struct {
	value *BTSketchArcDisplayData349
	isSet bool
}

func (v NullableBTSketchArcDisplayData349) Get() *BTSketchArcDisplayData349 {
	return v.value
}

func (v *NullableBTSketchArcDisplayData349) Set(val *BTSketchArcDisplayData349) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSketchArcDisplayData349) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSketchArcDisplayData349) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSketchArcDisplayData349(val *BTSketchArcDisplayData349) *NullableBTSketchArcDisplayData349 {
	return &NullableBTSketchArcDisplayData349{value: val, isSet: true}
}

func (v NullableBTSketchArcDisplayData349) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSketchArcDisplayData349) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
