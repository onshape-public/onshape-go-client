/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6826-13f16e212af4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMateOccurrenceData1671 struct for BTMateOccurrenceData1671
type BTMateOccurrenceData1671 struct {
	BtType     *string             `json:"btType,omitempty"`
	Visibility *string             `json:"visibility,omitempty"`
	ValueMap   *map[string]float64 `json:"valueMap,omitempty"`
	Values     []float64           `json:"values,omitempty"`
}

// NewBTMateOccurrenceData1671 instantiates a new BTMateOccurrenceData1671 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMateOccurrenceData1671() *BTMateOccurrenceData1671 {
	this := BTMateOccurrenceData1671{}
	return &this
}

// NewBTMateOccurrenceData1671WithDefaults instantiates a new BTMateOccurrenceData1671 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMateOccurrenceData1671WithDefaults() *BTMateOccurrenceData1671 {
	this := BTMateOccurrenceData1671{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMateOccurrenceData1671) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateOccurrenceData1671) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMateOccurrenceData1671) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMateOccurrenceData1671) SetBtType(v string) {
	o.BtType = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *BTMateOccurrenceData1671) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateOccurrenceData1671) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *BTMateOccurrenceData1671) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *BTMateOccurrenceData1671) SetVisibility(v string) {
	o.Visibility = &v
}

// GetValueMap returns the ValueMap field value if set, zero value otherwise.
func (o *BTMateOccurrenceData1671) GetValueMap() map[string]float64 {
	if o == nil || o.ValueMap == nil {
		var ret map[string]float64
		return ret
	}
	return *o.ValueMap
}

// GetValueMapOk returns a tuple with the ValueMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateOccurrenceData1671) GetValueMapOk() (*map[string]float64, bool) {
	if o == nil || o.ValueMap == nil {
		return nil, false
	}
	return o.ValueMap, true
}

// HasValueMap returns a boolean if a field has been set.
func (o *BTMateOccurrenceData1671) HasValueMap() bool {
	if o != nil && o.ValueMap != nil {
		return true
	}

	return false
}

// SetValueMap gets a reference to the given map[string]float64 and assigns it to the ValueMap field.
func (o *BTMateOccurrenceData1671) SetValueMap(v map[string]float64) {
	o.ValueMap = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *BTMateOccurrenceData1671) GetValues() []float64 {
	if o == nil || o.Values == nil {
		var ret []float64
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateOccurrenceData1671) GetValuesOk() ([]float64, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *BTMateOccurrenceData1671) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []float64 and assigns it to the Values field.
func (o *BTMateOccurrenceData1671) SetValues(v []float64) {
	o.Values = v
}

func (o BTMateOccurrenceData1671) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.ValueMap != nil {
		toSerialize["valueMap"] = o.ValueMap
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableBTMateOccurrenceData1671 struct {
	value *BTMateOccurrenceData1671
	isSet bool
}

func (v NullableBTMateOccurrenceData1671) Get() *BTMateOccurrenceData1671 {
	return v.value
}

func (v *NullableBTMateOccurrenceData1671) Set(val *BTMateOccurrenceData1671) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMateOccurrenceData1671) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMateOccurrenceData1671) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMateOccurrenceData1671(val *BTMateOccurrenceData1671) *NullableBTMateOccurrenceData1671 {
	return &NullableBTMateOccurrenceData1671{value: val, isSet: true}
}

func (v NullableBTMateOccurrenceData1671) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMateOccurrenceData1671) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
