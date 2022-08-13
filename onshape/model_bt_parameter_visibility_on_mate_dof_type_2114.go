/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6017-d06351faf6f2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTParameterVisibilityOnMateDOFType2114 struct for BTParameterVisibilityOnMateDOFType2114
type BTParameterVisibilityOnMateDOFType2114 struct {
	BtType      *string `json:"btType,omitempty"`
	InArray     *bool   `json:"inArray,omitempty"`
	ParameterId *string `json:"parameterId,omitempty"`
	Value       *string `json:"value,omitempty"`
}

// NewBTParameterVisibilityOnMateDOFType2114 instantiates a new BTParameterVisibilityOnMateDOFType2114 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterVisibilityOnMateDOFType2114() *BTParameterVisibilityOnMateDOFType2114 {
	this := BTParameterVisibilityOnMateDOFType2114{}
	return &this
}

// NewBTParameterVisibilityOnMateDOFType2114WithDefaults instantiates a new BTParameterVisibilityOnMateDOFType2114 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterVisibilityOnMateDOFType2114WithDefaults() *BTParameterVisibilityOnMateDOFType2114 {
	this := BTParameterVisibilityOnMateDOFType2114{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterVisibilityOnMateDOFType2114) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityOnMateDOFType2114) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterVisibilityOnMateDOFType2114) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterVisibilityOnMateDOFType2114) SetBtType(v string) {
	o.BtType = &v
}

// GetInArray returns the InArray field value if set, zero value otherwise.
func (o *BTParameterVisibilityOnMateDOFType2114) GetInArray() bool {
	if o == nil || o.InArray == nil {
		var ret bool
		return ret
	}
	return *o.InArray
}

// GetInArrayOk returns a tuple with the InArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityOnMateDOFType2114) GetInArrayOk() (*bool, bool) {
	if o == nil || o.InArray == nil {
		return nil, false
	}
	return o.InArray, true
}

// HasInArray returns a boolean if a field has been set.
func (o *BTParameterVisibilityOnMateDOFType2114) HasInArray() bool {
	if o != nil && o.InArray != nil {
		return true
	}

	return false
}

// SetInArray gets a reference to the given bool and assigns it to the InArray field.
func (o *BTParameterVisibilityOnMateDOFType2114) SetInArray(v bool) {
	o.InArray = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTParameterVisibilityOnMateDOFType2114) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityOnMateDOFType2114) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTParameterVisibilityOnMateDOFType2114) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTParameterVisibilityOnMateDOFType2114) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTParameterVisibilityOnMateDOFType2114) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityOnMateDOFType2114) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTParameterVisibilityOnMateDOFType2114) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BTParameterVisibilityOnMateDOFType2114) SetValue(v string) {
	o.Value = &v
}

func (o BTParameterVisibilityOnMateDOFType2114) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.InArray != nil {
		toSerialize["inArray"] = o.InArray
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterVisibilityOnMateDOFType2114 struct {
	value *BTParameterVisibilityOnMateDOFType2114
	isSet bool
}

func (v NullableBTParameterVisibilityOnMateDOFType2114) Get() *BTParameterVisibilityOnMateDOFType2114 {
	return v.value
}

func (v *NullableBTParameterVisibilityOnMateDOFType2114) Set(val *BTParameterVisibilityOnMateDOFType2114) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterVisibilityOnMateDOFType2114) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterVisibilityOnMateDOFType2114) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterVisibilityOnMateDOFType2114(val *BTParameterVisibilityOnMateDOFType2114) *NullableBTParameterVisibilityOnMateDOFType2114 {
	return &NullableBTParameterVisibilityOnMateDOFType2114{value: val, isSet: true}
}

func (v NullableBTParameterVisibilityOnMateDOFType2114) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterVisibilityOnMateDOFType2114) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
