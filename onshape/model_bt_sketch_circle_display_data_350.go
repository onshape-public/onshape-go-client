/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.158.9878-ff51e1211d95
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSketchCircleDisplayData350 struct for BTSketchCircleDisplayData350
type BTSketchCircleDisplayData350 struct {
	Points []float64 `json:"points,omitempty"`
	BtType *string   `json:"btType,omitempty"`
	Radius *float64  `json:"radius,omitempty"`
}

// NewBTSketchCircleDisplayData350 instantiates a new BTSketchCircleDisplayData350 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSketchCircleDisplayData350() *BTSketchCircleDisplayData350 {
	this := BTSketchCircleDisplayData350{}
	return &this
}

// NewBTSketchCircleDisplayData350WithDefaults instantiates a new BTSketchCircleDisplayData350 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSketchCircleDisplayData350WithDefaults() *BTSketchCircleDisplayData350 {
	this := BTSketchCircleDisplayData350{}
	return &this
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BTSketchCircleDisplayData350) GetPoints() []float64 {
	if o == nil || o.Points == nil {
		var ret []float64
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchCircleDisplayData350) GetPointsOk() ([]float64, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *BTSketchCircleDisplayData350) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []float64 and assigns it to the Points field.
func (o *BTSketchCircleDisplayData350) SetPoints(v []float64) {
	o.Points = v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchCircleDisplayData350) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchCircleDisplayData350) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSketchCircleDisplayData350) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSketchCircleDisplayData350) SetBtType(v string) {
	o.BtType = &v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *BTSketchCircleDisplayData350) GetRadius() float64 {
	if o == nil || o.Radius == nil {
		var ret float64
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchCircleDisplayData350) GetRadiusOk() (*float64, bool) {
	if o == nil || o.Radius == nil {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *BTSketchCircleDisplayData350) HasRadius() bool {
	if o != nil && o.Radius != nil {
		return true
	}

	return false
}

// SetRadius gets a reference to the given float64 and assigns it to the Radius field.
func (o *BTSketchCircleDisplayData350) SetRadius(v float64) {
	o.Radius = &v
}

func (o BTSketchCircleDisplayData350) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Radius != nil {
		toSerialize["radius"] = o.Radius
	}
	return json.Marshal(toSerialize)
}

type NullableBTSketchCircleDisplayData350 struct {
	value *BTSketchCircleDisplayData350
	isSet bool
}

func (v NullableBTSketchCircleDisplayData350) Get() *BTSketchCircleDisplayData350 {
	return v.value
}

func (v *NullableBTSketchCircleDisplayData350) Set(val *BTSketchCircleDisplayData350) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSketchCircleDisplayData350) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSketchCircleDisplayData350) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSketchCircleDisplayData350(val *BTSketchCircleDisplayData350) *NullableBTSketchCircleDisplayData350 {
	return &NullableBTSketchCircleDisplayData350{value: val, isSet: true}
}

func (v NullableBTSketchCircleDisplayData350) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSketchCircleDisplayData350) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
