/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23149-3610d8cd750e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSketchEllipseDisplayData712 struct for BTSketchEllipseDisplayData712
type BTSketchEllipseDisplayData712 struct {
	BtType      *string   `json:"btType,omitempty"`
	Points      []float64 `json:"points,omitempty"`
	MinorRadius *float64  `json:"minorRadius,omitempty"`
	Offset      *float64  `json:"offset,omitempty"`
	Radius      *float64  `json:"radius,omitempty"`
}

// NewBTSketchEllipseDisplayData712 instantiates a new BTSketchEllipseDisplayData712 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSketchEllipseDisplayData712() *BTSketchEllipseDisplayData712 {
	this := BTSketchEllipseDisplayData712{}
	return &this
}

// NewBTSketchEllipseDisplayData712WithDefaults instantiates a new BTSketchEllipseDisplayData712 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSketchEllipseDisplayData712WithDefaults() *BTSketchEllipseDisplayData712 {
	this := BTSketchEllipseDisplayData712{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchEllipseDisplayData712) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEllipseDisplayData712) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSketchEllipseDisplayData712) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSketchEllipseDisplayData712) SetBtType(v string) {
	o.BtType = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BTSketchEllipseDisplayData712) GetPoints() []float64 {
	if o == nil || o.Points == nil {
		var ret []float64
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEllipseDisplayData712) GetPointsOk() ([]float64, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *BTSketchEllipseDisplayData712) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []float64 and assigns it to the Points field.
func (o *BTSketchEllipseDisplayData712) SetPoints(v []float64) {
	o.Points = v
}

// GetMinorRadius returns the MinorRadius field value if set, zero value otherwise.
func (o *BTSketchEllipseDisplayData712) GetMinorRadius() float64 {
	if o == nil || o.MinorRadius == nil {
		var ret float64
		return ret
	}
	return *o.MinorRadius
}

// GetMinorRadiusOk returns a tuple with the MinorRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEllipseDisplayData712) GetMinorRadiusOk() (*float64, bool) {
	if o == nil || o.MinorRadius == nil {
		return nil, false
	}
	return o.MinorRadius, true
}

// HasMinorRadius returns a boolean if a field has been set.
func (o *BTSketchEllipseDisplayData712) HasMinorRadius() bool {
	if o != nil && o.MinorRadius != nil {
		return true
	}

	return false
}

// SetMinorRadius gets a reference to the given float64 and assigns it to the MinorRadius field.
func (o *BTSketchEllipseDisplayData712) SetMinorRadius(v float64) {
	o.MinorRadius = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *BTSketchEllipseDisplayData712) GetOffset() float64 {
	if o == nil || o.Offset == nil {
		var ret float64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEllipseDisplayData712) GetOffsetOk() (*float64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *BTSketchEllipseDisplayData712) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given float64 and assigns it to the Offset field.
func (o *BTSketchEllipseDisplayData712) SetOffset(v float64) {
	o.Offset = &v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *BTSketchEllipseDisplayData712) GetRadius() float64 {
	if o == nil || o.Radius == nil {
		var ret float64
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEllipseDisplayData712) GetRadiusOk() (*float64, bool) {
	if o == nil || o.Radius == nil {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *BTSketchEllipseDisplayData712) HasRadius() bool {
	if o != nil && o.Radius != nil {
		return true
	}

	return false
}

// SetRadius gets a reference to the given float64 and assigns it to the Radius field.
func (o *BTSketchEllipseDisplayData712) SetRadius(v float64) {
	o.Radius = &v
}

func (o BTSketchEllipseDisplayData712) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	if o.MinorRadius != nil {
		toSerialize["minorRadius"] = o.MinorRadius
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Radius != nil {
		toSerialize["radius"] = o.Radius
	}
	return json.Marshal(toSerialize)
}

type NullableBTSketchEllipseDisplayData712 struct {
	value *BTSketchEllipseDisplayData712
	isSet bool
}

func (v NullableBTSketchEllipseDisplayData712) Get() *BTSketchEllipseDisplayData712 {
	return v.value
}

func (v *NullableBTSketchEllipseDisplayData712) Set(val *BTSketchEllipseDisplayData712) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSketchEllipseDisplayData712) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSketchEllipseDisplayData712) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSketchEllipseDisplayData712(val *BTSketchEllipseDisplayData712) *NullableBTSketchEllipseDisplayData712 {
	return &NullableBTSketchEllipseDisplayData712{value: val, isSet: true}
}

func (v NullableBTSketchEllipseDisplayData712) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSketchEllipseDisplayData712) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
