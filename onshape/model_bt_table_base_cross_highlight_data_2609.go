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
	"fmt"
)

// BTTableBaseCrossHighlightData2609 - struct for BTTableBaseCrossHighlightData2609
type BTTableBaseCrossHighlightData2609 struct {
	implBTTableBaseCrossHighlightData2609 interface{}
}

// BTTableCrossHighlightData1753AsBTTableBaseCrossHighlightData2609 is a convenience function that returns BTTableCrossHighlightData1753 wrapped in BTTableBaseCrossHighlightData2609
func (o *BTTableCrossHighlightData1753) AsBTTableBaseCrossHighlightData2609() *BTTableBaseCrossHighlightData2609 {
	return &BTTableBaseCrossHighlightData2609{o}
}

// BTTableAssemblyCrossHighlightData2675AsBTTableBaseCrossHighlightData2609 is a convenience function that returns BTTableAssemblyCrossHighlightData2675 wrapped in BTTableBaseCrossHighlightData2609
func (o *BTTableAssemblyCrossHighlightData2675) AsBTTableBaseCrossHighlightData2609() *BTTableBaseCrossHighlightData2609 {
	return &BTTableBaseCrossHighlightData2609{o}
}

// NewBTTableBaseCrossHighlightData2609 instantiates a new BTTableBaseCrossHighlightData2609 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableBaseCrossHighlightData2609() *BTTableBaseCrossHighlightData2609 {
	this := BTTableBaseCrossHighlightData2609{Newbase_BTTableBaseCrossHighlightData2609()}
	return &this
}

// NewBTTableBaseCrossHighlightData2609WithDefaults instantiates a new BTTableBaseCrossHighlightData2609 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableBaseCrossHighlightData2609WithDefaults() *BTTableBaseCrossHighlightData2609 {
	this := BTTableBaseCrossHighlightData2609{Newbase_BTTableBaseCrossHighlightData2609WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableBaseCrossHighlightData2609) GetBtType() string {
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
func (o *BTTableBaseCrossHighlightData2609) GetBtTypeOk() (*string, bool) {
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
func (o *BTTableBaseCrossHighlightData2609) HasBtType() bool {
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
func (o *BTTableBaseCrossHighlightData2609) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTTableBaseCrossHighlightData2609) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTTableAssemblyCrossHighlightData-2675'
	if jsonDict["btType"] == "BTTableAssemblyCrossHighlightData-2675" {
		// try to unmarshal JSON data into BTTableAssemblyCrossHighlightData2675
		var qr *BTTableAssemblyCrossHighlightData2675
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableBaseCrossHighlightData2609 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableBaseCrossHighlightData2609 = nil
			return fmt.Errorf("Failed to unmarshal BTTableBaseCrossHighlightData2609 as BTTableAssemblyCrossHighlightData2675: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTTableCrossHighlightData-1753'
	if jsonDict["btType"] == "BTTableCrossHighlightData-1753" {
		// try to unmarshal JSON data into BTTableCrossHighlightData1753
		var qr *BTTableCrossHighlightData1753
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableBaseCrossHighlightData2609 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableBaseCrossHighlightData2609 = nil
			return fmt.Errorf("Failed to unmarshal BTTableBaseCrossHighlightData2609 as BTTableCrossHighlightData1753: %s", err.Error())
		}
	}

	var qtx *base_BTTableBaseCrossHighlightData2609
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTTableBaseCrossHighlightData2609 = qtx
		return nil // data stored in dst.base_BTTableBaseCrossHighlightData2609, return on the first match
	} else {
		dst.implBTTableBaseCrossHighlightData2609 = nil
		return fmt.Errorf("Failed to unmarshal BTTableBaseCrossHighlightData2609 as base_BTTableBaseCrossHighlightData2609: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTTableBaseCrossHighlightData2609) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTTableBaseCrossHighlightData2609) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTTableBaseCrossHighlightData2609
}

type NullableBTTableBaseCrossHighlightData2609 struct {
	value *BTTableBaseCrossHighlightData2609
	isSet bool
}

func (v NullableBTTableBaseCrossHighlightData2609) Get() *BTTableBaseCrossHighlightData2609 {
	return v.value
}

func (v *NullableBTTableBaseCrossHighlightData2609) Set(val *BTTableBaseCrossHighlightData2609) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableBaseCrossHighlightData2609) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableBaseCrossHighlightData2609) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableBaseCrossHighlightData2609(val *BTTableBaseCrossHighlightData2609) *NullableBTTableBaseCrossHighlightData2609 {
	return &NullableBTTableBaseCrossHighlightData2609{value: val, isSet: true}
}

func (v NullableBTTableBaseCrossHighlightData2609) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableBaseCrossHighlightData2609) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTTableBaseCrossHighlightData2609 struct {
	BtType *string `json:"btType,omitempty"`
}

// Newbase_BTTableBaseCrossHighlightData2609 instantiates a new base_BTTableBaseCrossHighlightData2609 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTTableBaseCrossHighlightData2609() *base_BTTableBaseCrossHighlightData2609 {
	this := base_BTTableBaseCrossHighlightData2609{}
	return &this
}

// Newbase_BTTableBaseCrossHighlightData2609WithDefaults instantiates a new base_BTTableBaseCrossHighlightData2609 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTTableBaseCrossHighlightData2609WithDefaults() *base_BTTableBaseCrossHighlightData2609 {
	this := base_BTTableBaseCrossHighlightData2609{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTTableBaseCrossHighlightData2609) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableBaseCrossHighlightData2609) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTTableBaseCrossHighlightData2609) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTTableBaseCrossHighlightData2609) SetBtType(v string) {
	o.BtType = &v
}

func (o base_BTTableBaseCrossHighlightData2609) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}
