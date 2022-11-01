/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7120-bb0c8e12993e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTJPathElement2297 - Identifies a node in the json sturcture, beginning at the specified startNode.
type BTJPathElement2297 struct {
	implBTJPathElement2297 interface{}
}

// BTJPathIndex1871AsBTJPathElement2297 is a convenience function that returns BTJPathIndex1871 wrapped in BTJPathElement2297
func (o *BTJPathIndex1871) AsBTJPathElement2297() *BTJPathElement2297 {
	return &BTJPathElement2297{o}
}

// BTJPathKey3221AsBTJPathElement2297 is a convenience function that returns BTJPathKey3221 wrapped in BTJPathElement2297
func (o *BTJPathKey3221) AsBTJPathElement2297() *BTJPathElement2297 {
	return &BTJPathElement2297{o}
}

// NewBTJPathElement2297 instantiates a new BTJPathElement2297 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTJPathElement2297() *BTJPathElement2297 {
	this := BTJPathElement2297{Newbase_BTJPathElement2297()}
	return &this
}

// NewBTJPathElement2297WithDefaults instantiates a new BTJPathElement2297 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTJPathElement2297WithDefaults() *BTJPathElement2297 {
	this := BTJPathElement2297{Newbase_BTJPathElement2297WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTJPathElement2297) GetBtType() string {
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
func (o *BTJPathElement2297) GetBtTypeOk() (*string, bool) {
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
func (o *BTJPathElement2297) HasBtType() bool {
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
func (o *BTJPathElement2297) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTJPathElement2297) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTJPathIndex-1871'
	if jsonDict["btType"] == "BTJPathIndex-1871" {
		// try to unmarshal JSON data into BTJPathIndex1871
		var qr *BTJPathIndex1871
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTJPathElement2297 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTJPathElement2297 = nil
			return fmt.Errorf("Failed to unmarshal BTJPathElement2297 as BTJPathIndex1871: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTJPathKey-3221'
	if jsonDict["btType"] == "BTJPathKey-3221" {
		// try to unmarshal JSON data into BTJPathKey3221
		var qr *BTJPathKey3221
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTJPathElement2297 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTJPathElement2297 = nil
			return fmt.Errorf("Failed to unmarshal BTJPathElement2297 as BTJPathKey3221: %s", err.Error())
		}
	}

	var qtx *base_BTJPathElement2297
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTJPathElement2297 = qtx
		return nil // data stored in dst.base_BTJPathElement2297, return on the first match
	} else {
		dst.implBTJPathElement2297 = nil
		return fmt.Errorf("Failed to unmarshal BTJPathElement2297 as base_BTJPathElement2297: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTJPathElement2297) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTJPathElement2297) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTJPathElement2297
}

type NullableBTJPathElement2297 struct {
	value *BTJPathElement2297
	isSet bool
}

func (v NullableBTJPathElement2297) Get() *BTJPathElement2297 {
	return v.value
}

func (v *NullableBTJPathElement2297) Set(val *BTJPathElement2297) {
	v.value = val
	v.isSet = true
}

func (v NullableBTJPathElement2297) IsSet() bool {
	return v.isSet
}

func (v *NullableBTJPathElement2297) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTJPathElement2297(val *BTJPathElement2297) *NullableBTJPathElement2297 {
	return &NullableBTJPathElement2297{value: val, isSet: true}
}

func (v NullableBTJPathElement2297) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTJPathElement2297) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTJPathElement2297 struct {
	BtType *string `json:"btType,omitempty"`
}

// Newbase_BTJPathElement2297 instantiates a new base_BTJPathElement2297 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTJPathElement2297() *base_BTJPathElement2297 {
	this := base_BTJPathElement2297{}
	return &this
}

// Newbase_BTJPathElement2297WithDefaults instantiates a new base_BTJPathElement2297 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTJPathElement2297WithDefaults() *base_BTJPathElement2297 {
	this := base_BTJPathElement2297{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTJPathElement2297) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTJPathElement2297) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTJPathElement2297) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTJPathElement2297) SetBtType(v string) {
	o.BtType = &v
}

func (o base_BTJPathElement2297) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}
