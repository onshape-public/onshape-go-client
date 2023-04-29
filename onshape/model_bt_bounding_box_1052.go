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
	"fmt"
)

// BTBoundingBox1052 - An axis-aligned bounding box indicated by two opposite corners.
type BTBoundingBox1052 struct {
	implBTBoundingBox1052 interface{}
}

// BTNonAlignedBoundingBox4180AsBTBoundingBox1052 is a convenience function that returns BTNonAlignedBoundingBox4180 wrapped in BTBoundingBox1052
func (o *BTNonAlignedBoundingBox4180) AsBTBoundingBox1052() *BTBoundingBox1052 {
	return &BTBoundingBox1052{o}
}

// NewBTBoundingBox1052 instantiates a new BTBoundingBox1052 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBoundingBox1052() *BTBoundingBox1052 {
	this := BTBoundingBox1052{Newbase_BTBoundingBox1052()}
	return &this
}

// NewBTBoundingBox1052WithDefaults instantiates a new BTBoundingBox1052 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBoundingBox1052WithDefaults() *BTBoundingBox1052 {
	this := BTBoundingBox1052{Newbase_BTBoundingBox1052WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBoundingBox1052) GetBtType() string {
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
func (o *BTBoundingBox1052) GetBtTypeOk() (*string, bool) {
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
func (o *BTBoundingBox1052) HasBtType() bool {
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
func (o *BTBoundingBox1052) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetMaxCorner returns the MaxCorner field value if set, zero value otherwise.
func (o *BTBoundingBox1052) GetMaxCorner() BTVector3d389 {
	type getResult interface {
		GetMaxCorner() BTVector3d389
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMaxCorner()
	} else {
		var de BTVector3d389
		return de
	}
}

// GetMaxCornerOk returns a tuple with the MaxCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBoundingBox1052) GetMaxCornerOk() (*BTVector3d389, bool) {
	type getResult interface {
		GetMaxCornerOk() (*BTVector3d389, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMaxCornerOk()
	} else {
		return nil, false
	}
}

// HasMaxCorner returns a boolean if a field has been set.
func (o *BTBoundingBox1052) HasMaxCorner() bool {
	type getResult interface {
		HasMaxCorner() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMaxCorner()
	} else {
		return false
	}
}

// SetMaxCorner gets a reference to the given BTVector3d389 and assigns it to the MaxCorner field.
func (o *BTBoundingBox1052) SetMaxCorner(v BTVector3d389) {
	type getResult interface {
		SetMaxCorner(v BTVector3d389)
	}

	o.GetActualInstance().(getResult).SetMaxCorner(v)
}

// GetMinCorner returns the MinCorner field value if set, zero value otherwise.
func (o *BTBoundingBox1052) GetMinCorner() BTVector3d389 {
	type getResult interface {
		GetMinCorner() BTVector3d389
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMinCorner()
	} else {
		var de BTVector3d389
		return de
	}
}

// GetMinCornerOk returns a tuple with the MinCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBoundingBox1052) GetMinCornerOk() (*BTVector3d389, bool) {
	type getResult interface {
		GetMinCornerOk() (*BTVector3d389, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMinCornerOk()
	} else {
		return nil, false
	}
}

// HasMinCorner returns a boolean if a field has been set.
func (o *BTBoundingBox1052) HasMinCorner() bool {
	type getResult interface {
		HasMinCorner() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMinCorner()
	} else {
		return false
	}
}

// SetMinCorner gets a reference to the given BTVector3d389 and assigns it to the MinCorner field.
func (o *BTBoundingBox1052) SetMinCorner(v BTVector3d389) {
	type getResult interface {
		SetMinCorner(v BTVector3d389)
	}

	o.GetActualInstance().(getResult).SetMinCorner(v)
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *BTBoundingBox1052) GetValid() bool {
	type getResult interface {
		GetValid() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetValid()
	} else {
		var de bool
		return de
	}
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBoundingBox1052) GetValidOk() (*bool, bool) {
	type getResult interface {
		GetValidOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetValidOk()
	} else {
		return nil, false
	}
}

// HasValid returns a boolean if a field has been set.
func (o *BTBoundingBox1052) HasValid() bool {
	type getResult interface {
		HasValid() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasValid()
	} else {
		return false
	}
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *BTBoundingBox1052) SetValid(v bool) {
	type getResult interface {
		SetValid(v bool)
	}

	o.GetActualInstance().(getResult).SetValid(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTBoundingBox1052) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTNonAlignedBoundingBox-4180'
	if jsonDict["btType"] == "BTNonAlignedBoundingBox-4180" {
		// try to unmarshal JSON data into BTNonAlignedBoundingBox4180
		var qr *BTNonAlignedBoundingBox4180
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTBoundingBox1052 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTBoundingBox1052 = nil
			return fmt.Errorf("Failed to unmarshal BTBoundingBox1052 as BTNonAlignedBoundingBox4180: %s", err.Error())
		}
	}

	var qtx *base_BTBoundingBox1052
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTBoundingBox1052 = qtx
		return nil // data stored in dst.base_BTBoundingBox1052, return on the first match
	} else {
		dst.implBTBoundingBox1052 = nil
		return fmt.Errorf("Failed to unmarshal BTBoundingBox1052 as base_BTBoundingBox1052: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTBoundingBox1052) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTBoundingBox1052) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTBoundingBox1052
}

type NullableBTBoundingBox1052 struct {
	value *BTBoundingBox1052
	isSet bool
}

func (v NullableBTBoundingBox1052) Get() *BTBoundingBox1052 {
	return v.value
}

func (v *NullableBTBoundingBox1052) Set(val *BTBoundingBox1052) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBoundingBox1052) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBoundingBox1052) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBoundingBox1052(val *BTBoundingBox1052) *NullableBTBoundingBox1052 {
	return &NullableBTBoundingBox1052{value: val, isSet: true}
}

func (v NullableBTBoundingBox1052) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBoundingBox1052) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTBoundingBox1052 struct {
	BtType    *string        `json:"btType,omitempty"`
	MaxCorner *BTVector3d389 `json:"maxCorner,omitempty"`
	MinCorner *BTVector3d389 `json:"minCorner,omitempty"`
	Valid     *bool          `json:"valid,omitempty"`
}

// Newbase_BTBoundingBox1052 instantiates a new base_BTBoundingBox1052 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTBoundingBox1052() *base_BTBoundingBox1052 {
	this := base_BTBoundingBox1052{}
	return &this
}

// Newbase_BTBoundingBox1052WithDefaults instantiates a new base_BTBoundingBox1052 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTBoundingBox1052WithDefaults() *base_BTBoundingBox1052 {
	this := base_BTBoundingBox1052{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTBoundingBox1052) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTBoundingBox1052) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTBoundingBox1052) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTBoundingBox1052) SetBtType(v string) {
	o.BtType = &v
}

// GetMaxCorner returns the MaxCorner field value if set, zero value otherwise.
func (o *base_BTBoundingBox1052) GetMaxCorner() BTVector3d389 {
	if o == nil || o.MaxCorner == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.MaxCorner
}

// GetMaxCornerOk returns a tuple with the MaxCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTBoundingBox1052) GetMaxCornerOk() (*BTVector3d389, bool) {
	if o == nil || o.MaxCorner == nil {
		return nil, false
	}
	return o.MaxCorner, true
}

// HasMaxCorner returns a boolean if a field has been set.
func (o *base_BTBoundingBox1052) HasMaxCorner() bool {
	if o != nil && o.MaxCorner != nil {
		return true
	}

	return false
}

// SetMaxCorner gets a reference to the given BTVector3d389 and assigns it to the MaxCorner field.
func (o *base_BTBoundingBox1052) SetMaxCorner(v BTVector3d389) {
	o.MaxCorner = &v
}

// GetMinCorner returns the MinCorner field value if set, zero value otherwise.
func (o *base_BTBoundingBox1052) GetMinCorner() BTVector3d389 {
	if o == nil || o.MinCorner == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.MinCorner
}

// GetMinCornerOk returns a tuple with the MinCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTBoundingBox1052) GetMinCornerOk() (*BTVector3d389, bool) {
	if o == nil || o.MinCorner == nil {
		return nil, false
	}
	return o.MinCorner, true
}

// HasMinCorner returns a boolean if a field has been set.
func (o *base_BTBoundingBox1052) HasMinCorner() bool {
	if o != nil && o.MinCorner != nil {
		return true
	}

	return false
}

// SetMinCorner gets a reference to the given BTVector3d389 and assigns it to the MinCorner field.
func (o *base_BTBoundingBox1052) SetMinCorner(v BTVector3d389) {
	o.MinCorner = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *base_BTBoundingBox1052) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTBoundingBox1052) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *base_BTBoundingBox1052) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *base_BTBoundingBox1052) SetValid(v bool) {
	o.Valid = &v
}

func (o base_BTBoundingBox1052) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.MaxCorner != nil {
		toSerialize["maxCorner"] = o.MaxCorner
	}
	if o.MinCorner != nil {
		toSerialize["minCorner"] = o.MinCorner
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	return json.Marshal(toSerialize)
}
