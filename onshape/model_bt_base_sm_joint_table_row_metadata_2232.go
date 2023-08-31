/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.168.21279-402b6292597b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTBaseSMJointTableRowMetadata2232 - struct for BTBaseSMJointTableRowMetadata2232
type BTBaseSMJointTableRowMetadata2232 struct {
	implBTBaseSMJointTableRowMetadata2232 interface{}
}

// BTSMOtherJointTableRowMetadata2640AsBTBaseSMJointTableRowMetadata2232 is a convenience function that returns BTSMOtherJointTableRowMetadata2640 wrapped in BTBaseSMJointTableRowMetadata2232
func (o *BTSMOtherJointTableRowMetadata2640) AsBTBaseSMJointTableRowMetadata2232() *BTBaseSMJointTableRowMetadata2232 {
	return &BTBaseSMJointTableRowMetadata2232{o}
}

// BTSMBendTableRowMetadata1705AsBTBaseSMJointTableRowMetadata2232 is a convenience function that returns BTSMBendTableRowMetadata1705 wrapped in BTBaseSMJointTableRowMetadata2232
func (o *BTSMBendTableRowMetadata1705) AsBTBaseSMJointTableRowMetadata2232() *BTBaseSMJointTableRowMetadata2232 {
	return &BTBaseSMJointTableRowMetadata2232{o}
}

// NewBTBaseSMJointTableRowMetadata2232 instantiates a new BTBaseSMJointTableRowMetadata2232 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBaseSMJointTableRowMetadata2232() *BTBaseSMJointTableRowMetadata2232 {
	this := BTBaseSMJointTableRowMetadata2232{Newbase_BTBaseSMJointTableRowMetadata2232()}
	return &this
}

// NewBTBaseSMJointTableRowMetadata2232WithDefaults instantiates a new BTBaseSMJointTableRowMetadata2232 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBaseSMJointTableRowMetadata2232WithDefaults() *BTBaseSMJointTableRowMetadata2232 {
	this := BTBaseSMJointTableRowMetadata2232{Newbase_BTBaseSMJointTableRowMetadata2232WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBaseSMJointTableRowMetadata2232) GetBtType() string {
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
func (o *BTBaseSMJointTableRowMetadata2232) GetBtTypeOk() (*string, bool) {
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
func (o *BTBaseSMJointTableRowMetadata2232) HasBtType() bool {
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
func (o *BTBaseSMJointTableRowMetadata2232) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetCrossHighlightDataIfAny returns the CrossHighlightDataIfAny field value if set, zero value otherwise.
func (o *BTBaseSMJointTableRowMetadata2232) GetCrossHighlightDataIfAny() BTTableCrossHighlightData1753 {
	type getResult interface {
		GetCrossHighlightDataIfAny() BTTableCrossHighlightData1753
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetCrossHighlightDataIfAny()
	} else {
		var de BTTableCrossHighlightData1753
		return de
	}
}

// GetCrossHighlightDataIfAnyOk returns a tuple with the CrossHighlightDataIfAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBaseSMJointTableRowMetadata2232) GetCrossHighlightDataIfAnyOk() (*BTTableCrossHighlightData1753, bool) {
	type getResult interface {
		GetCrossHighlightDataIfAnyOk() (*BTTableCrossHighlightData1753, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetCrossHighlightDataIfAnyOk()
	} else {
		return nil, false
	}
}

// HasCrossHighlightDataIfAny returns a boolean if a field has been set.
func (o *BTBaseSMJointTableRowMetadata2232) HasCrossHighlightDataIfAny() bool {
	type getResult interface {
		HasCrossHighlightDataIfAny() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasCrossHighlightDataIfAny()
	} else {
		return false
	}
}

// SetCrossHighlightDataIfAny gets a reference to the given BTTableCrossHighlightData1753 and assigns it to the CrossHighlightDataIfAny field.
func (o *BTBaseSMJointTableRowMetadata2232) SetCrossHighlightDataIfAny(v BTTableCrossHighlightData1753) {
	type getResult interface {
		SetCrossHighlightDataIfAny(v BTTableCrossHighlightData1753)
	}

	o.GetActualInstance().(getResult).SetCrossHighlightDataIfAny(v)
}

// GetCrossHighlightData returns the CrossHighlightData field value if set, zero value otherwise.
func (o *BTBaseSMJointTableRowMetadata2232) GetCrossHighlightData() BTTableCrossHighlightData1753 {
	type getResult interface {
		GetCrossHighlightData() BTTableCrossHighlightData1753
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetCrossHighlightData()
	} else {
		var de BTTableCrossHighlightData1753
		return de
	}
}

// GetCrossHighlightDataOk returns a tuple with the CrossHighlightData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBaseSMJointTableRowMetadata2232) GetCrossHighlightDataOk() (*BTTableCrossHighlightData1753, bool) {
	type getResult interface {
		GetCrossHighlightDataOk() (*BTTableCrossHighlightData1753, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetCrossHighlightDataOk()
	} else {
		return nil, false
	}
}

// HasCrossHighlightData returns a boolean if a field has been set.
func (o *BTBaseSMJointTableRowMetadata2232) HasCrossHighlightData() bool {
	type getResult interface {
		HasCrossHighlightData() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasCrossHighlightData()
	} else {
		return false
	}
}

// SetCrossHighlightData gets a reference to the given BTTableCrossHighlightData1753 and assigns it to the CrossHighlightData field.
func (o *BTBaseSMJointTableRowMetadata2232) SetCrossHighlightData(v BTTableCrossHighlightData1753) {
	type getResult interface {
		SetCrossHighlightData(v BTTableCrossHighlightData1753)
	}

	o.GetActualInstance().(getResult).SetCrossHighlightData(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTBaseSMJointTableRowMetadata2232) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTSMBendTableRowMetadata-1705'
	if jsonDict["btType"] == "BTSMBendTableRowMetadata-1705" {
		// try to unmarshal JSON data into BTSMBendTableRowMetadata1705
		var qr *BTSMBendTableRowMetadata1705
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTBaseSMJointTableRowMetadata2232 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTBaseSMJointTableRowMetadata2232 = nil
			return fmt.Errorf("Failed to unmarshal BTBaseSMJointTableRowMetadata2232 as BTSMBendTableRowMetadata1705: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSMOtherJointTableRowMetadata-2640'
	if jsonDict["btType"] == "BTSMOtherJointTableRowMetadata-2640" {
		// try to unmarshal JSON data into BTSMOtherJointTableRowMetadata2640
		var qr *BTSMOtherJointTableRowMetadata2640
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTBaseSMJointTableRowMetadata2232 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTBaseSMJointTableRowMetadata2232 = nil
			return fmt.Errorf("Failed to unmarshal BTBaseSMJointTableRowMetadata2232 as BTSMOtherJointTableRowMetadata2640: %s", err.Error())
		}
	}

	var qtx *base_BTBaseSMJointTableRowMetadata2232
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTBaseSMJointTableRowMetadata2232 = qtx
		return nil // data stored in dst.base_BTBaseSMJointTableRowMetadata2232, return on the first match
	} else {
		dst.implBTBaseSMJointTableRowMetadata2232 = nil
		return fmt.Errorf("Failed to unmarshal BTBaseSMJointTableRowMetadata2232 as base_BTBaseSMJointTableRowMetadata2232: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTBaseSMJointTableRowMetadata2232) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTBaseSMJointTableRowMetadata2232) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTBaseSMJointTableRowMetadata2232
}

type NullableBTBaseSMJointTableRowMetadata2232 struct {
	value *BTBaseSMJointTableRowMetadata2232
	isSet bool
}

func (v NullableBTBaseSMJointTableRowMetadata2232) Get() *BTBaseSMJointTableRowMetadata2232 {
	return v.value
}

func (v *NullableBTBaseSMJointTableRowMetadata2232) Set(val *BTBaseSMJointTableRowMetadata2232) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBaseSMJointTableRowMetadata2232) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBaseSMJointTableRowMetadata2232) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBaseSMJointTableRowMetadata2232(val *BTBaseSMJointTableRowMetadata2232) *NullableBTBaseSMJointTableRowMetadata2232 {
	return &NullableBTBaseSMJointTableRowMetadata2232{value: val, isSet: true}
}

func (v NullableBTBaseSMJointTableRowMetadata2232) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBaseSMJointTableRowMetadata2232) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTBaseSMJointTableRowMetadata2232 struct {
	BtType                  *string                        `json:"btType,omitempty"`
	CrossHighlightDataIfAny *BTTableCrossHighlightData1753 `json:"crossHighlightDataIfAny,omitempty"`
	CrossHighlightData      *BTTableCrossHighlightData1753 `json:"crossHighlightData,omitempty"`
}

// Newbase_BTBaseSMJointTableRowMetadata2232 instantiates a new base_BTBaseSMJointTableRowMetadata2232 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTBaseSMJointTableRowMetadata2232() *base_BTBaseSMJointTableRowMetadata2232 {
	this := base_BTBaseSMJointTableRowMetadata2232{}
	return &this
}

// Newbase_BTBaseSMJointTableRowMetadata2232WithDefaults instantiates a new base_BTBaseSMJointTableRowMetadata2232 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTBaseSMJointTableRowMetadata2232WithDefaults() *base_BTBaseSMJointTableRowMetadata2232 {
	this := base_BTBaseSMJointTableRowMetadata2232{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTBaseSMJointTableRowMetadata2232) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTBaseSMJointTableRowMetadata2232) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTBaseSMJointTableRowMetadata2232) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTBaseSMJointTableRowMetadata2232) SetBtType(v string) {
	o.BtType = &v
}

// GetCrossHighlightDataIfAny returns the CrossHighlightDataIfAny field value if set, zero value otherwise.
func (o *base_BTBaseSMJointTableRowMetadata2232) GetCrossHighlightDataIfAny() BTTableCrossHighlightData1753 {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		var ret BTTableCrossHighlightData1753
		return ret
	}
	return *o.CrossHighlightDataIfAny
}

// GetCrossHighlightDataIfAnyOk returns a tuple with the CrossHighlightDataIfAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTBaseSMJointTableRowMetadata2232) GetCrossHighlightDataIfAnyOk() (*BTTableCrossHighlightData1753, bool) {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		return nil, false
	}
	return o.CrossHighlightDataIfAny, true
}

// HasCrossHighlightDataIfAny returns a boolean if a field has been set.
func (o *base_BTBaseSMJointTableRowMetadata2232) HasCrossHighlightDataIfAny() bool {
	if o != nil && o.CrossHighlightDataIfAny != nil {
		return true
	}

	return false
}

// SetCrossHighlightDataIfAny gets a reference to the given BTTableCrossHighlightData1753 and assigns it to the CrossHighlightDataIfAny field.
func (o *base_BTBaseSMJointTableRowMetadata2232) SetCrossHighlightDataIfAny(v BTTableCrossHighlightData1753) {
	o.CrossHighlightDataIfAny = &v
}

// GetCrossHighlightData returns the CrossHighlightData field value if set, zero value otherwise.
func (o *base_BTBaseSMJointTableRowMetadata2232) GetCrossHighlightData() BTTableCrossHighlightData1753 {
	if o == nil || o.CrossHighlightData == nil {
		var ret BTTableCrossHighlightData1753
		return ret
	}
	return *o.CrossHighlightData
}

// GetCrossHighlightDataOk returns a tuple with the CrossHighlightData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTBaseSMJointTableRowMetadata2232) GetCrossHighlightDataOk() (*BTTableCrossHighlightData1753, bool) {
	if o == nil || o.CrossHighlightData == nil {
		return nil, false
	}
	return o.CrossHighlightData, true
}

// HasCrossHighlightData returns a boolean if a field has been set.
func (o *base_BTBaseSMJointTableRowMetadata2232) HasCrossHighlightData() bool {
	if o != nil && o.CrossHighlightData != nil {
		return true
	}

	return false
}

// SetCrossHighlightData gets a reference to the given BTTableCrossHighlightData1753 and assigns it to the CrossHighlightData field.
func (o *base_BTBaseSMJointTableRowMetadata2232) SetCrossHighlightData(v BTTableCrossHighlightData1753) {
	o.CrossHighlightData = &v
}

func (o base_BTBaseSMJointTableRowMetadata2232) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CrossHighlightDataIfAny != nil {
		toSerialize["crossHighlightDataIfAny"] = o.CrossHighlightDataIfAny
	}
	if o.CrossHighlightData != nil {
		toSerialize["crossHighlightData"] = o.CrossHighlightData
	}
	return json.Marshal(toSerialize)
}
