/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.175.29152-5591e93bd1c1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBodyTypeFilter112 struct for BTBodyTypeFilter112
type BTBodyTypeFilter112 struct {
	BtType   *string      `json:"btType,omitempty"`
	BodyType *GBTBodyType `json:"bodyType,omitempty"`
}

// NewBTBodyTypeFilter112 instantiates a new BTBodyTypeFilter112 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBodyTypeFilter112() *BTBodyTypeFilter112 {
	this := BTBodyTypeFilter112{}
	return &this
}

// NewBTBodyTypeFilter112WithDefaults instantiates a new BTBodyTypeFilter112 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBodyTypeFilter112WithDefaults() *BTBodyTypeFilter112 {
	this := BTBodyTypeFilter112{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBodyTypeFilter112) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBodyTypeFilter112) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTBodyTypeFilter112) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTBodyTypeFilter112) SetBtType(v string) {
	o.BtType = &v
}

// GetBodyType returns the BodyType field value if set, zero value otherwise.
func (o *BTBodyTypeFilter112) GetBodyType() GBTBodyType {
	if o == nil || o.BodyType == nil {
		var ret GBTBodyType
		return ret
	}
	return *o.BodyType
}

// GetBodyTypeOk returns a tuple with the BodyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBodyTypeFilter112) GetBodyTypeOk() (*GBTBodyType, bool) {
	if o == nil || o.BodyType == nil {
		return nil, false
	}
	return o.BodyType, true
}

// HasBodyType returns a boolean if a field has been set.
func (o *BTBodyTypeFilter112) HasBodyType() bool {
	if o != nil && o.BodyType != nil {
		return true
	}

	return false
}

// SetBodyType gets a reference to the given GBTBodyType and assigns it to the BodyType field.
func (o *BTBodyTypeFilter112) SetBodyType(v GBTBodyType) {
	o.BodyType = &v
}

func (o BTBodyTypeFilter112) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.BodyType != nil {
		toSerialize["bodyType"] = o.BodyType
	}
	return json.Marshal(toSerialize)
}

type NullableBTBodyTypeFilter112 struct {
	value *BTBodyTypeFilter112
	isSet bool
}

func (v NullableBTBodyTypeFilter112) Get() *BTBodyTypeFilter112 {
	return v.value
}

func (v *NullableBTBodyTypeFilter112) Set(val *BTBodyTypeFilter112) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBodyTypeFilter112) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBodyTypeFilter112) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBodyTypeFilter112(val *BTBodyTypeFilter112) *NullableBTBodyTypeFilter112 {
	return &NullableBTBodyTypeFilter112{value: val, isSet: true}
}

func (v NullableBTBodyTypeFilter112) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBodyTypeFilter112) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
