/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.163.15296-122c93d7dbb6
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPlaneOrientationFilter1700 struct for BTPlaneOrientationFilter1700
type BTPlaneOrientationFilter1700 struct {
	BtType *string        `json:"btType,omitempty"`
	Normal *BTVector3d389 `json:"normal,omitempty"`
}

// NewBTPlaneOrientationFilter1700 instantiates a new BTPlaneOrientationFilter1700 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPlaneOrientationFilter1700() *BTPlaneOrientationFilter1700 {
	this := BTPlaneOrientationFilter1700{}
	return &this
}

// NewBTPlaneOrientationFilter1700WithDefaults instantiates a new BTPlaneOrientationFilter1700 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPlaneOrientationFilter1700WithDefaults() *BTPlaneOrientationFilter1700 {
	this := BTPlaneOrientationFilter1700{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPlaneOrientationFilter1700) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlaneOrientationFilter1700) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPlaneOrientationFilter1700) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPlaneOrientationFilter1700) SetBtType(v string) {
	o.BtType = &v
}

// GetNormal returns the Normal field value if set, zero value otherwise.
func (o *BTPlaneOrientationFilter1700) GetNormal() BTVector3d389 {
	if o == nil || o.Normal == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Normal
}

// GetNormalOk returns a tuple with the Normal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlaneOrientationFilter1700) GetNormalOk() (*BTVector3d389, bool) {
	if o == nil || o.Normal == nil {
		return nil, false
	}
	return o.Normal, true
}

// HasNormal returns a boolean if a field has been set.
func (o *BTPlaneOrientationFilter1700) HasNormal() bool {
	if o != nil && o.Normal != nil {
		return true
	}

	return false
}

// SetNormal gets a reference to the given BTVector3d389 and assigns it to the Normal field.
func (o *BTPlaneOrientationFilter1700) SetNormal(v BTVector3d389) {
	o.Normal = &v
}

func (o BTPlaneOrientationFilter1700) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Normal != nil {
		toSerialize["normal"] = o.Normal
	}
	return json.Marshal(toSerialize)
}

type NullableBTPlaneOrientationFilter1700 struct {
	value *BTPlaneOrientationFilter1700
	isSet bool
}

func (v NullableBTPlaneOrientationFilter1700) Get() *BTPlaneOrientationFilter1700 {
	return v.value
}

func (v *NullableBTPlaneOrientationFilter1700) Set(val *BTPlaneOrientationFilter1700) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPlaneOrientationFilter1700) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPlaneOrientationFilter1700) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPlaneOrientationFilter1700(val *BTPlaneOrientationFilter1700) *NullableBTPlaneOrientationFilter1700 {
	return &NullableBTPlaneOrientationFilter1700{value: val, isSet: true}
}

func (v NullableBTPlaneOrientationFilter1700) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPlaneOrientationFilter1700) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
