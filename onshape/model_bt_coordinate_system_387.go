/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCoordinateSystem387 struct for BTCoordinateSystem387
type BTCoordinateSystem387 struct {
	// Type of JSON object.
	BtType *string        `json:"btType,omitempty"`
	Matrix *BTBSMatrix386 `json:"matrix,omitempty"`
	Origin *BTVector3d389 `json:"origin,omitempty"`
	XAxis  *BTVector3d389 `json:"xAxis,omitempty"`
	YAxis  *BTVector3d389 `json:"yAxis,omitempty"`
	ZAxis  *BTVector3d389 `json:"zAxis,omitempty"`
}

// NewBTCoordinateSystem387 instantiates a new BTCoordinateSystem387 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCoordinateSystem387() *BTCoordinateSystem387 {
	this := BTCoordinateSystem387{}
	return &this
}

// NewBTCoordinateSystem387WithDefaults instantiates a new BTCoordinateSystem387 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCoordinateSystem387WithDefaults() *BTCoordinateSystem387 {
	this := BTCoordinateSystem387{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCoordinateSystem387) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCoordinateSystem387) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTCoordinateSystem387) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTCoordinateSystem387) SetBtType(v string) {
	o.BtType = &v
}

// GetMatrix returns the Matrix field value if set, zero value otherwise.
func (o *BTCoordinateSystem387) GetMatrix() BTBSMatrix386 {
	if o == nil || o.Matrix == nil {
		var ret BTBSMatrix386
		return ret
	}
	return *o.Matrix
}

// GetMatrixOk returns a tuple with the Matrix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCoordinateSystem387) GetMatrixOk() (*BTBSMatrix386, bool) {
	if o == nil || o.Matrix == nil {
		return nil, false
	}
	return o.Matrix, true
}

// HasMatrix returns a boolean if a field has been set.
func (o *BTCoordinateSystem387) HasMatrix() bool {
	if o != nil && o.Matrix != nil {
		return true
	}

	return false
}

// SetMatrix gets a reference to the given BTBSMatrix386 and assigns it to the Matrix field.
func (o *BTCoordinateSystem387) SetMatrix(v BTBSMatrix386) {
	o.Matrix = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTCoordinateSystem387) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCoordinateSystem387) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTCoordinateSystem387) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTCoordinateSystem387) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

// GetXAxis returns the XAxis field value if set, zero value otherwise.
func (o *BTCoordinateSystem387) GetXAxis() BTVector3d389 {
	if o == nil || o.XAxis == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.XAxis
}

// GetXAxisOk returns a tuple with the XAxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCoordinateSystem387) GetXAxisOk() (*BTVector3d389, bool) {
	if o == nil || o.XAxis == nil {
		return nil, false
	}
	return o.XAxis, true
}

// HasXAxis returns a boolean if a field has been set.
func (o *BTCoordinateSystem387) HasXAxis() bool {
	if o != nil && o.XAxis != nil {
		return true
	}

	return false
}

// SetXAxis gets a reference to the given BTVector3d389 and assigns it to the XAxis field.
func (o *BTCoordinateSystem387) SetXAxis(v BTVector3d389) {
	o.XAxis = &v
}

// GetYAxis returns the YAxis field value if set, zero value otherwise.
func (o *BTCoordinateSystem387) GetYAxis() BTVector3d389 {
	if o == nil || o.YAxis == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.YAxis
}

// GetYAxisOk returns a tuple with the YAxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCoordinateSystem387) GetYAxisOk() (*BTVector3d389, bool) {
	if o == nil || o.YAxis == nil {
		return nil, false
	}
	return o.YAxis, true
}

// HasYAxis returns a boolean if a field has been set.
func (o *BTCoordinateSystem387) HasYAxis() bool {
	if o != nil && o.YAxis != nil {
		return true
	}

	return false
}

// SetYAxis gets a reference to the given BTVector3d389 and assigns it to the YAxis field.
func (o *BTCoordinateSystem387) SetYAxis(v BTVector3d389) {
	o.YAxis = &v
}

// GetZAxis returns the ZAxis field value if set, zero value otherwise.
func (o *BTCoordinateSystem387) GetZAxis() BTVector3d389 {
	if o == nil || o.ZAxis == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.ZAxis
}

// GetZAxisOk returns a tuple with the ZAxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCoordinateSystem387) GetZAxisOk() (*BTVector3d389, bool) {
	if o == nil || o.ZAxis == nil {
		return nil, false
	}
	return o.ZAxis, true
}

// HasZAxis returns a boolean if a field has been set.
func (o *BTCoordinateSystem387) HasZAxis() bool {
	if o != nil && o.ZAxis != nil {
		return true
	}

	return false
}

// SetZAxis gets a reference to the given BTVector3d389 and assigns it to the ZAxis field.
func (o *BTCoordinateSystem387) SetZAxis(v BTVector3d389) {
	o.ZAxis = &v
}

func (o BTCoordinateSystem387) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Matrix != nil {
		toSerialize["matrix"] = o.Matrix
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.XAxis != nil {
		toSerialize["xAxis"] = o.XAxis
	}
	if o.YAxis != nil {
		toSerialize["yAxis"] = o.YAxis
	}
	if o.ZAxis != nil {
		toSerialize["zAxis"] = o.ZAxis
	}
	return json.Marshal(toSerialize)
}

type NullableBTCoordinateSystem387 struct {
	value *BTCoordinateSystem387
	isSet bool
}

func (v NullableBTCoordinateSystem387) Get() *BTCoordinateSystem387 {
	return v.value
}

func (v *NullableBTCoordinateSystem387) Set(val *BTCoordinateSystem387) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCoordinateSystem387) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCoordinateSystem387) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCoordinateSystem387(val *BTCoordinateSystem387) *NullableBTCoordinateSystem387 {
	return &NullableBTCoordinateSystem387{value: val, isSet: true}
}

func (v NullableBTCoordinateSystem387) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCoordinateSystem387) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
