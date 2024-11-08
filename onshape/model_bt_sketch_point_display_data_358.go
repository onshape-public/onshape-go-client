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

// BTSketchPointDisplayData358 struct for BTSketchPointDisplayData358
type BTSketchPointDisplayData358 struct {
	BTSketchEntityDisplayData354
	BtType *string   `json:"btType,omitempty"`
	Points []float64 `json:"points,omitempty"`
	// Type of JSON object.
	BtType *string `json:"btType,omitempty"`
}

// NewBTSketchPointDisplayData358 instantiates a new BTSketchPointDisplayData358 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSketchPointDisplayData358() *BTSketchPointDisplayData358 {
	this := BTSketchPointDisplayData358{}
	return &this
}

// NewBTSketchPointDisplayData358WithDefaults instantiates a new BTSketchPointDisplayData358 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSketchPointDisplayData358WithDefaults() *BTSketchPointDisplayData358 {
	this := BTSketchPointDisplayData358{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchPointDisplayData358) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchPointDisplayData358) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSketchPointDisplayData358) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSketchPointDisplayData358) SetBtType(v string) {
	o.BtType = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BTSketchPointDisplayData358) GetPoints() []float64 {
	if o == nil || o.Points == nil {
		var ret []float64
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchPointDisplayData358) GetPointsOk() ([]float64, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *BTSketchPointDisplayData358) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []float64 and assigns it to the Points field.
func (o *BTSketchPointDisplayData358) SetPoints(v []float64) {
	o.Points = v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchPointDisplayData358) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchPointDisplayData358) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSketchPointDisplayData358) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSketchPointDisplayData358) SetBtType(v string) {
	o.BtType = &v
}

func (o BTSketchPointDisplayData358) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTSketchEntityDisplayData354, errBTSketchEntityDisplayData354 := json.Marshal(o.BTSketchEntityDisplayData354)
	if errBTSketchEntityDisplayData354 != nil {
		return []byte{}, errBTSketchEntityDisplayData354
	}
	errBTSketchEntityDisplayData354 = json.Unmarshal([]byte(serializedBTSketchEntityDisplayData354), &toSerialize)
	if errBTSketchEntityDisplayData354 != nil {
		return []byte{}, errBTSketchEntityDisplayData354
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTSketchPointDisplayData358 struct {
	value *BTSketchPointDisplayData358
	isSet bool
}

func (v NullableBTSketchPointDisplayData358) Get() *BTSketchPointDisplayData358 {
	return v.value
}

func (v *NullableBTSketchPointDisplayData358) Set(val *BTSketchPointDisplayData358) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSketchPointDisplayData358) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSketchPointDisplayData358) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSketchPointDisplayData358(val *BTSketchPointDisplayData358) *NullableBTSketchPointDisplayData358 {
	return &NullableBTSketchPointDisplayData358{value: val, isSet: true}
}

func (v NullableBTSketchPointDisplayData358) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSketchPointDisplayData358) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
