/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSMDefinitionEntityTypeFilter1651 struct for BTSMDefinitionEntityTypeFilter1651
type BTSMDefinitionEntityTypeFilter1651 struct {
	BTQueryFilter183
	BtType                 *string        `json:"btType,omitempty"`
	SmDefinitionEntityType *GBTEntityType `json:"smDefinitionEntityType,omitempty"`
}

// NewBTSMDefinitionEntityTypeFilter1651 instantiates a new BTSMDefinitionEntityTypeFilter1651 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSMDefinitionEntityTypeFilter1651() *BTSMDefinitionEntityTypeFilter1651 {
	this := BTSMDefinitionEntityTypeFilter1651{}
	return &this
}

// NewBTSMDefinitionEntityTypeFilter1651WithDefaults instantiates a new BTSMDefinitionEntityTypeFilter1651 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSMDefinitionEntityTypeFilter1651WithDefaults() *BTSMDefinitionEntityTypeFilter1651 {
	this := BTSMDefinitionEntityTypeFilter1651{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSMDefinitionEntityTypeFilter1651) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSMDefinitionEntityTypeFilter1651) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSMDefinitionEntityTypeFilter1651) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSMDefinitionEntityTypeFilter1651) SetBtType(v string) {
	o.BtType = &v
}

// GetSmDefinitionEntityType returns the SmDefinitionEntityType field value if set, zero value otherwise.
func (o *BTSMDefinitionEntityTypeFilter1651) GetSmDefinitionEntityType() GBTEntityType {
	if o == nil || o.SmDefinitionEntityType == nil {
		var ret GBTEntityType
		return ret
	}
	return *o.SmDefinitionEntityType
}

// GetSmDefinitionEntityTypeOk returns a tuple with the SmDefinitionEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSMDefinitionEntityTypeFilter1651) GetSmDefinitionEntityTypeOk() (*GBTEntityType, bool) {
	if o == nil || o.SmDefinitionEntityType == nil {
		return nil, false
	}
	return o.SmDefinitionEntityType, true
}

// HasSmDefinitionEntityType returns a boolean if a field has been set.
func (o *BTSMDefinitionEntityTypeFilter1651) HasSmDefinitionEntityType() bool {
	if o != nil && o.SmDefinitionEntityType != nil {
		return true
	}

	return false
}

// SetSmDefinitionEntityType gets a reference to the given GBTEntityType and assigns it to the SmDefinitionEntityType field.
func (o *BTSMDefinitionEntityTypeFilter1651) SetSmDefinitionEntityType(v GBTEntityType) {
	o.SmDefinitionEntityType = &v
}

func (o BTSMDefinitionEntityTypeFilter1651) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTQueryFilter183, errBTQueryFilter183 := json.Marshal(o.BTQueryFilter183)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	errBTQueryFilter183 = json.Unmarshal([]byte(serializedBTQueryFilter183), &toSerialize)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.SmDefinitionEntityType != nil {
		toSerialize["smDefinitionEntityType"] = o.SmDefinitionEntityType
	}
	return json.Marshal(toSerialize)
}

type NullableBTSMDefinitionEntityTypeFilter1651 struct {
	value *BTSMDefinitionEntityTypeFilter1651
	isSet bool
}

func (v NullableBTSMDefinitionEntityTypeFilter1651) Get() *BTSMDefinitionEntityTypeFilter1651 {
	return v.value
}

func (v *NullableBTSMDefinitionEntityTypeFilter1651) Set(val *BTSMDefinitionEntityTypeFilter1651) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSMDefinitionEntityTypeFilter1651) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSMDefinitionEntityTypeFilter1651) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSMDefinitionEntityTypeFilter1651(val *BTSMDefinitionEntityTypeFilter1651) *NullableBTSMDefinitionEntityTypeFilter1651 {
	return &NullableBTSMDefinitionEntityTypeFilter1651{value: val, isSet: true}
}

func (v NullableBTSMDefinitionEntityTypeFilter1651) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSMDefinitionEntityTypeFilter1651) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
