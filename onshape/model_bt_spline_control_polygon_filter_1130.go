/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.5998-d3227e94fd7e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSplineControlPolygonFilter1130 struct for BTSplineControlPolygonFilter1130
type BTSplineControlPolygonFilter1130 struct {
	BtType                     *string `json:"btType,omitempty"`
	AllowsSplineControlPolygon *bool   `json:"allowsSplineControlPolygon,omitempty"`
}

// NewBTSplineControlPolygonFilter1130 instantiates a new BTSplineControlPolygonFilter1130 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSplineControlPolygonFilter1130() *BTSplineControlPolygonFilter1130 {
	this := BTSplineControlPolygonFilter1130{}
	return &this
}

// NewBTSplineControlPolygonFilter1130WithDefaults instantiates a new BTSplineControlPolygonFilter1130 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSplineControlPolygonFilter1130WithDefaults() *BTSplineControlPolygonFilter1130 {
	this := BTSplineControlPolygonFilter1130{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSplineControlPolygonFilter1130) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSplineControlPolygonFilter1130) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSplineControlPolygonFilter1130) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSplineControlPolygonFilter1130) SetBtType(v string) {
	o.BtType = &v
}

// GetAllowsSplineControlPolygon returns the AllowsSplineControlPolygon field value if set, zero value otherwise.
func (o *BTSplineControlPolygonFilter1130) GetAllowsSplineControlPolygon() bool {
	if o == nil || o.AllowsSplineControlPolygon == nil {
		var ret bool
		return ret
	}
	return *o.AllowsSplineControlPolygon
}

// GetAllowsSplineControlPolygonOk returns a tuple with the AllowsSplineControlPolygon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSplineControlPolygonFilter1130) GetAllowsSplineControlPolygonOk() (*bool, bool) {
	if o == nil || o.AllowsSplineControlPolygon == nil {
		return nil, false
	}
	return o.AllowsSplineControlPolygon, true
}

// HasAllowsSplineControlPolygon returns a boolean if a field has been set.
func (o *BTSplineControlPolygonFilter1130) HasAllowsSplineControlPolygon() bool {
	if o != nil && o.AllowsSplineControlPolygon != nil {
		return true
	}

	return false
}

// SetAllowsSplineControlPolygon gets a reference to the given bool and assigns it to the AllowsSplineControlPolygon field.
func (o *BTSplineControlPolygonFilter1130) SetAllowsSplineControlPolygon(v bool) {
	o.AllowsSplineControlPolygon = &v
}

func (o BTSplineControlPolygonFilter1130) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.AllowsSplineControlPolygon != nil {
		toSerialize["allowsSplineControlPolygon"] = o.AllowsSplineControlPolygon
	}
	return json.Marshal(toSerialize)
}

type NullableBTSplineControlPolygonFilter1130 struct {
	value *BTSplineControlPolygonFilter1130
	isSet bool
}

func (v NullableBTSplineControlPolygonFilter1130) Get() *BTSplineControlPolygonFilter1130 {
	return v.value
}

func (v *NullableBTSplineControlPolygonFilter1130) Set(val *BTSplineControlPolygonFilter1130) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSplineControlPolygonFilter1130) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSplineControlPolygonFilter1130) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSplineControlPolygonFilter1130(val *BTSplineControlPolygonFilter1130) *NullableBTSplineControlPolygonFilter1130 {
	return &NullableBTSplineControlPolygonFilter1130{value: val, isSet: true}
}

func (v NullableBTSplineControlPolygonFilter1130) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSplineControlPolygonFilter1130) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
