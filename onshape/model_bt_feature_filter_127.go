/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7050-fa39b5514a7a
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTFeatureFilter127 struct for BTFeatureFilter127
type BTFeatureFilter127 struct {
	BtType    *string `json:"btType,omitempty"`
	Exclusion *string `json:"exclusion,omitempty"`
	FeatureId *string `json:"featureId,omitempty"`
}

// NewBTFeatureFilter127 instantiates a new BTFeatureFilter127 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureFilter127() *BTFeatureFilter127 {
	this := BTFeatureFilter127{}
	return &this
}

// NewBTFeatureFilter127WithDefaults instantiates a new BTFeatureFilter127 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureFilter127WithDefaults() *BTFeatureFilter127 {
	this := BTFeatureFilter127{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFeatureFilter127) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureFilter127) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFeatureFilter127) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFeatureFilter127) SetBtType(v string) {
	o.BtType = &v
}

// GetExclusion returns the Exclusion field value if set, zero value otherwise.
func (o *BTFeatureFilter127) GetExclusion() string {
	if o == nil || o.Exclusion == nil {
		var ret string
		return ret
	}
	return *o.Exclusion
}

// GetExclusionOk returns a tuple with the Exclusion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureFilter127) GetExclusionOk() (*string, bool) {
	if o == nil || o.Exclusion == nil {
		return nil, false
	}
	return o.Exclusion, true
}

// HasExclusion returns a boolean if a field has been set.
func (o *BTFeatureFilter127) HasExclusion() bool {
	if o != nil && o.Exclusion != nil {
		return true
	}

	return false
}

// SetExclusion gets a reference to the given string and assigns it to the Exclusion field.
func (o *BTFeatureFilter127) SetExclusion(v string) {
	o.Exclusion = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTFeatureFilter127) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureFilter127) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTFeatureFilter127) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTFeatureFilter127) SetFeatureId(v string) {
	o.FeatureId = &v
}

func (o BTFeatureFilter127) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Exclusion != nil {
		toSerialize["exclusion"] = o.Exclusion
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	return json.Marshal(toSerialize)
}

type NullableBTFeatureFilter127 struct {
	value *BTFeatureFilter127
	isSet bool
}

func (v NullableBTFeatureFilter127) Get() *BTFeatureFilter127 {
	return v.value
}

func (v *NullableBTFeatureFilter127) Set(val *BTFeatureFilter127) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureFilter127) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureFilter127) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureFilter127(val *BTFeatureFilter127) *NullableBTFeatureFilter127 {
	return &NullableBTFeatureFilter127{value: val, isSet: true}
}

func (v NullableBTFeatureFilter127) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureFilter127) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
