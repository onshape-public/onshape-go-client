/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.159.11153-01265f451c50
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTQuantityRange181 - struct for BTQuantityRange181
type BTQuantityRange181 struct {
	implBTQuantityRange181 interface{}
}

// BTNullableQuantityRange1340AsBTQuantityRange181 is a convenience function that returns BTNullableQuantityRange1340 wrapped in BTQuantityRange181
func (o *BTNullableQuantityRange1340) AsBTQuantityRange181() *BTQuantityRange181 {
	return &BTQuantityRange181{o}
}

// NewBTQuantityRange181 instantiates a new BTQuantityRange181 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTQuantityRange181() *BTQuantityRange181 {
	this := BTQuantityRange181{Newbase_BTQuantityRange181()}
	return &this
}

// NewBTQuantityRange181WithDefaults instantiates a new BTQuantityRange181 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTQuantityRange181WithDefaults() *BTQuantityRange181 {
	this := BTQuantityRange181{Newbase_BTQuantityRange181WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTQuantityRange181) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTQuantityRange181) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTQuantityRange181) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTQuantityRange181) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *BTQuantityRange181) GetDefaultValue() float64 {
	type getResult interface {
		GetDefaultValue() float64
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDefaultValue()
	} else {
		var de float64
		return de
	}
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTQuantityRange181) GetDefaultValueOk() (*float64, bool) {
	type getResult interface {
		GetDefaultValueOk() (*float64, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDefaultValueOk()
	} else {
		return nil, false
	}
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BTQuantityRange181) HasDefaultValue() bool {
	type getResult interface {
		HasDefaultValue() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDefaultValue()
	} else {
		return false
	}
}

// SetDefaultValue gets a reference to the given float64 and assigns it to the DefaultValue field.
func (o *BTQuantityRange181) SetDefaultValue(v float64) {
	type getResult interface {
		SetDefaultValue(v float64)
	}

	o.GetActualInstance().(getResult).SetDefaultValue(v)
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *BTQuantityRange181) GetLocation() BTLocationInfo226 {
	type getResult interface {
		GetLocation() BTLocationInfo226
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLocation()
	} else {
		var de BTLocationInfo226
		return de
	}
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTQuantityRange181) GetLocationOk() (*BTLocationInfo226, bool) {
	type getResult interface {
		GetLocationOk() (*BTLocationInfo226, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLocationOk()
	} else {
		return nil, false
	}
}

// HasLocation returns a boolean if a field has been set.
func (o *BTQuantityRange181) HasLocation() bool {
	type getResult interface {
		HasLocation() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasLocation()
	} else {
		return false
	}
}

// SetLocation gets a reference to the given BTLocationInfo226 and assigns it to the Location field.
func (o *BTQuantityRange181) SetLocation(v BTLocationInfo226) {
	type getResult interface {
		SetLocation(v BTLocationInfo226)
	}

	o.GetActualInstance().(getResult).SetLocation(v)
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *BTQuantityRange181) GetMaxValue() float64 {
	type getResult interface {
		GetMaxValue() float64
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMaxValue()
	} else {
		var de float64
		return de
	}
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTQuantityRange181) GetMaxValueOk() (*float64, bool) {
	type getResult interface {
		GetMaxValueOk() (*float64, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMaxValueOk()
	} else {
		return nil, false
	}
}

// HasMaxValue returns a boolean if a field has been set.
func (o *BTQuantityRange181) HasMaxValue() bool {
	type getResult interface {
		HasMaxValue() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMaxValue()
	} else {
		return false
	}
}

// SetMaxValue gets a reference to the given float64 and assigns it to the MaxValue field.
func (o *BTQuantityRange181) SetMaxValue(v float64) {
	type getResult interface {
		SetMaxValue(v float64)
	}

	o.GetActualInstance().(getResult).SetMaxValue(v)
}

// GetMinValue returns the MinValue field value if set, zero value otherwise.
func (o *BTQuantityRange181) GetMinValue() float64 {
	type getResult interface {
		GetMinValue() float64
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMinValue()
	} else {
		var de float64
		return de
	}
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTQuantityRange181) GetMinValueOk() (*float64, bool) {
	type getResult interface {
		GetMinValueOk() (*float64, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMinValueOk()
	} else {
		return nil, false
	}
}

// HasMinValue returns a boolean if a field has been set.
func (o *BTQuantityRange181) HasMinValue() bool {
	type getResult interface {
		HasMinValue() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMinValue()
	} else {
		return false
	}
}

// SetMinValue gets a reference to the given float64 and assigns it to the MinValue field.
func (o *BTQuantityRange181) SetMinValue(v float64) {
	type getResult interface {
		SetMinValue(v float64)
	}

	o.GetActualInstance().(getResult).SetMinValue(v)
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BTQuantityRange181) GetUnits() string {
	type getResult interface {
		GetUnits() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetUnits()
	} else {
		var de string
		return de
	}
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTQuantityRange181) GetUnitsOk() (*string, bool) {
	type getResult interface {
		GetUnitsOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetUnitsOk()
	} else {
		return nil, false
	}
}

// HasUnits returns a boolean if a field has been set.
func (o *BTQuantityRange181) HasUnits() bool {
	type getResult interface {
		HasUnits() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasUnits()
	} else {
		return false
	}
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *BTQuantityRange181) SetUnits(v string) {
	type getResult interface {
		SetUnits(v string)
	}

	o.GetActualInstance().(getResult).SetUnits(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTQuantityRange181) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTNullableQuantityRange-1340'
	if jsonDict["btType"] == "BTNullableQuantityRange-1340" {
		// try to unmarshal JSON data into BTNullableQuantityRange1340
		var qr *BTNullableQuantityRange1340
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQuantityRange181 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQuantityRange181 = nil
			return fmt.Errorf("Failed to unmarshal BTQuantityRange181 as BTNullableQuantityRange1340: %s", err.Error())
		}
	}

	var qtx *base_BTQuantityRange181
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTQuantityRange181 = qtx
		return nil // data stored in dst.base_BTQuantityRange181, return on the first match
	} else {
		dst.implBTQuantityRange181 = nil
		return fmt.Errorf("Failed to unmarshal BTQuantityRange181 as base_BTQuantityRange181: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTQuantityRange181) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTQuantityRange181) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTQuantityRange181
}

type NullableBTQuantityRange181 struct {
	value *BTQuantityRange181
	isSet bool
}

func (v NullableBTQuantityRange181) Get() *BTQuantityRange181 {
	return v.value
}

func (v *NullableBTQuantityRange181) Set(val *BTQuantityRange181) {
	v.value = val
	v.isSet = true
}

func (v NullableBTQuantityRange181) IsSet() bool {
	return v.isSet
}

func (v *NullableBTQuantityRange181) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTQuantityRange181(val *BTQuantityRange181) *NullableBTQuantityRange181 {
	return &NullableBTQuantityRange181{value: val, isSet: true}
}

func (v NullableBTQuantityRange181) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTQuantityRange181) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTQuantityRange181 struct {
	BtType       *string            `json:"btType,omitempty"`
	DefaultValue *float64           `json:"defaultValue,omitempty"`
	Location     *BTLocationInfo226 `json:"location,omitempty"`
	MaxValue     *float64           `json:"maxValue,omitempty"`
	MinValue     *float64           `json:"minValue,omitempty"`
	Units        *string            `json:"units,omitempty"`
}

// Newbase_BTQuantityRange181 instantiates a new base_BTQuantityRange181 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTQuantityRange181() *base_BTQuantityRange181 {
	this := base_BTQuantityRange181{}
	return &this
}

// Newbase_BTQuantityRange181WithDefaults instantiates a new base_BTQuantityRange181 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTQuantityRange181WithDefaults() *base_BTQuantityRange181 {
	this := base_BTQuantityRange181{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTQuantityRange181) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTQuantityRange181) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTQuantityRange181) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTQuantityRange181) SetBtType(v string) {
	o.BtType = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *base_BTQuantityRange181) GetDefaultValue() float64 {
	if o == nil || o.DefaultValue == nil {
		var ret float64
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTQuantityRange181) GetDefaultValueOk() (*float64, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *base_BTQuantityRange181) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given float64 and assigns it to the DefaultValue field.
func (o *base_BTQuantityRange181) SetDefaultValue(v float64) {
	o.DefaultValue = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *base_BTQuantityRange181) GetLocation() BTLocationInfo226 {
	if o == nil || o.Location == nil {
		var ret BTLocationInfo226
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTQuantityRange181) GetLocationOk() (*BTLocationInfo226, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *base_BTQuantityRange181) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given BTLocationInfo226 and assigns it to the Location field.
func (o *base_BTQuantityRange181) SetLocation(v BTLocationInfo226) {
	o.Location = &v
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *base_BTQuantityRange181) GetMaxValue() float64 {
	if o == nil || o.MaxValue == nil {
		var ret float64
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTQuantityRange181) GetMaxValueOk() (*float64, bool) {
	if o == nil || o.MaxValue == nil {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *base_BTQuantityRange181) HasMaxValue() bool {
	if o != nil && o.MaxValue != nil {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given float64 and assigns it to the MaxValue field.
func (o *base_BTQuantityRange181) SetMaxValue(v float64) {
	o.MaxValue = &v
}

// GetMinValue returns the MinValue field value if set, zero value otherwise.
func (o *base_BTQuantityRange181) GetMinValue() float64 {
	if o == nil || o.MinValue == nil {
		var ret float64
		return ret
	}
	return *o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTQuantityRange181) GetMinValueOk() (*float64, bool) {
	if o == nil || o.MinValue == nil {
		return nil, false
	}
	return o.MinValue, true
}

// HasMinValue returns a boolean if a field has been set.
func (o *base_BTQuantityRange181) HasMinValue() bool {
	if o != nil && o.MinValue != nil {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given float64 and assigns it to the MinValue field.
func (o *base_BTQuantityRange181) SetMinValue(v float64) {
	o.MinValue = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *base_BTQuantityRange181) GetUnits() string {
	if o == nil || o.Units == nil {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTQuantityRange181) GetUnitsOk() (*string, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *base_BTQuantityRange181) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *base_BTQuantityRange181) SetUnits(v string) {
	o.Units = &v
}

func (o base_BTQuantityRange181) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.MaxValue != nil {
		toSerialize["maxValue"] = o.MaxValue
	}
	if o.MinValue != nil {
		toSerialize["minValue"] = o.MinValue
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	return json.Marshal(toSerialize)
}
