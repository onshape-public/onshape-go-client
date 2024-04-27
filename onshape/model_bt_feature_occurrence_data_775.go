/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTFeatureOccurrenceData775 - struct for BTFeatureOccurrenceData775
type BTFeatureOccurrenceData775 struct {
	implBTFeatureOccurrenceData775 interface{}
}

// BTMateOccurrenceData1671AsBTFeatureOccurrenceData775 is a convenience function that returns BTMateOccurrenceData1671 wrapped in BTFeatureOccurrenceData775
func (o *BTMateOccurrenceData1671) AsBTFeatureOccurrenceData775() *BTFeatureOccurrenceData775 {
	return &BTFeatureOccurrenceData775{o}
}

// NewBTFeatureOccurrenceData775 instantiates a new BTFeatureOccurrenceData775 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureOccurrenceData775() *BTFeatureOccurrenceData775 {
	this := BTFeatureOccurrenceData775{Newbase_BTFeatureOccurrenceData775()}
	return &this
}

// NewBTFeatureOccurrenceData775WithDefaults instantiates a new BTFeatureOccurrenceData775 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureOccurrenceData775WithDefaults() *BTFeatureOccurrenceData775 {
	this := BTFeatureOccurrenceData775{Newbase_BTFeatureOccurrenceData775WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFeatureOccurrenceData775) GetBtType() string {
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
func (o *BTFeatureOccurrenceData775) GetBtTypeOk() (*string, bool) {
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
func (o *BTFeatureOccurrenceData775) HasBtType() bool {
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
func (o *BTFeatureOccurrenceData775) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *BTFeatureOccurrenceData775) GetVisibility() GBTBSFeatureVisibility {
	type getResult interface {
		GetVisibility() GBTBSFeatureVisibility
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetVisibility()
	} else {
		var de GBTBSFeatureVisibility
		return de
	}
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureOccurrenceData775) GetVisibilityOk() (*GBTBSFeatureVisibility, bool) {
	type getResult interface {
		GetVisibilityOk() (*GBTBSFeatureVisibility, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetVisibilityOk()
	} else {
		return nil, false
	}
}

// HasVisibility returns a boolean if a field has been set.
func (o *BTFeatureOccurrenceData775) HasVisibility() bool {
	type getResult interface {
		HasVisibility() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasVisibility()
	} else {
		return false
	}
}

// SetVisibility gets a reference to the given GBTBSFeatureVisibility and assigns it to the Visibility field.
func (o *BTFeatureOccurrenceData775) SetVisibility(v GBTBSFeatureVisibility) {
	type getResult interface {
		SetVisibility(v GBTBSFeatureVisibility)
	}

	o.GetActualInstance().(getResult).SetVisibility(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTFeatureOccurrenceData775) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BTMateOccurrenceData-1671'
	if jsonDict["btType"] == "BTMateOccurrenceData-1671" {
		// try to unmarshal JSON data into BTMateOccurrenceData1671
		var qr *BTMateOccurrenceData1671
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureOccurrenceData775 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureOccurrenceData775 = nil
			return fmt.Errorf("failed to unmarshal BTFeatureOccurrenceData775 as BTMateOccurrenceData1671: %s", err.Error())
		}
	}

	var qtx *base_BTFeatureOccurrenceData775
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTFeatureOccurrenceData775 = qtx
		return nil // data stored in dst.base_BTFeatureOccurrenceData775, return on the first match
	} else {
		dst.implBTFeatureOccurrenceData775 = nil
		return fmt.Errorf("failed to unmarshal BTFeatureOccurrenceData775 as base_BTFeatureOccurrenceData775: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTFeatureOccurrenceData775) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTFeatureOccurrenceData775) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTFeatureOccurrenceData775
}

type NullableBTFeatureOccurrenceData775 struct {
	value *BTFeatureOccurrenceData775
	isSet bool
}

func (v NullableBTFeatureOccurrenceData775) Get() *BTFeatureOccurrenceData775 {
	return v.value
}

func (v *NullableBTFeatureOccurrenceData775) Set(val *BTFeatureOccurrenceData775) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureOccurrenceData775) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureOccurrenceData775) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureOccurrenceData775(val *BTFeatureOccurrenceData775) *NullableBTFeatureOccurrenceData775 {
	return &NullableBTFeatureOccurrenceData775{value: val, isSet: true}
}

func (v NullableBTFeatureOccurrenceData775) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureOccurrenceData775) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTFeatureOccurrenceData775 struct {
	// Type of JSON object.
	BtType     *string                 `json:"btType,omitempty"`
	Visibility *GBTBSFeatureVisibility `json:"visibility,omitempty"`
}

// Newbase_BTFeatureOccurrenceData775 instantiates a new base_BTFeatureOccurrenceData775 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTFeatureOccurrenceData775() *base_BTFeatureOccurrenceData775 {
	this := base_BTFeatureOccurrenceData775{}
	return &this
}

// Newbase_BTFeatureOccurrenceData775WithDefaults instantiates a new base_BTFeatureOccurrenceData775 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTFeatureOccurrenceData775WithDefaults() *base_BTFeatureOccurrenceData775 {
	this := base_BTFeatureOccurrenceData775{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTFeatureOccurrenceData775) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureOccurrenceData775) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTFeatureOccurrenceData775) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTFeatureOccurrenceData775) SetBtType(v string) {
	o.BtType = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *base_BTFeatureOccurrenceData775) GetVisibility() GBTBSFeatureVisibility {
	if o == nil || o.Visibility == nil {
		var ret GBTBSFeatureVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureOccurrenceData775) GetVisibilityOk() (*GBTBSFeatureVisibility, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *base_BTFeatureOccurrenceData775) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given GBTBSFeatureVisibility and assigns it to the Visibility field.
func (o *base_BTFeatureOccurrenceData775) SetVisibility(v GBTBSFeatureVisibility) {
	o.Visibility = &v
}

func (o base_BTFeatureOccurrenceData775) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}
