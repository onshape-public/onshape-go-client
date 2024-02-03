/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSMSpecificMetadata1315 struct for BTSMSpecificMetadata1315
type BTSMSpecificMetadata1315 struct {
	BtType               *string        `json:"btType,omitempty"`
	DefinitionEntityType *GBTEntityType `json:"definitionEntityType,omitempty"`
}

// NewBTSMSpecificMetadata1315 instantiates a new BTSMSpecificMetadata1315 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSMSpecificMetadata1315() *BTSMSpecificMetadata1315 {
	this := BTSMSpecificMetadata1315{}
	return &this
}

// NewBTSMSpecificMetadata1315WithDefaults instantiates a new BTSMSpecificMetadata1315 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSMSpecificMetadata1315WithDefaults() *BTSMSpecificMetadata1315 {
	this := BTSMSpecificMetadata1315{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSMSpecificMetadata1315) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSMSpecificMetadata1315) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSMSpecificMetadata1315) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSMSpecificMetadata1315) SetBtType(v string) {
	o.BtType = &v
}

// GetDefinitionEntityType returns the DefinitionEntityType field value if set, zero value otherwise.
func (o *BTSMSpecificMetadata1315) GetDefinitionEntityType() GBTEntityType {
	if o == nil || o.DefinitionEntityType == nil {
		var ret GBTEntityType
		return ret
	}
	return *o.DefinitionEntityType
}

// GetDefinitionEntityTypeOk returns a tuple with the DefinitionEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSMSpecificMetadata1315) GetDefinitionEntityTypeOk() (*GBTEntityType, bool) {
	if o == nil || o.DefinitionEntityType == nil {
		return nil, false
	}
	return o.DefinitionEntityType, true
}

// HasDefinitionEntityType returns a boolean if a field has been set.
func (o *BTSMSpecificMetadata1315) HasDefinitionEntityType() bool {
	if o != nil && o.DefinitionEntityType != nil {
		return true
	}

	return false
}

// SetDefinitionEntityType gets a reference to the given GBTEntityType and assigns it to the DefinitionEntityType field.
func (o *BTSMSpecificMetadata1315) SetDefinitionEntityType(v GBTEntityType) {
	o.DefinitionEntityType = &v
}

func (o BTSMSpecificMetadata1315) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DefinitionEntityType != nil {
		toSerialize["definitionEntityType"] = o.DefinitionEntityType
	}
	return json.Marshal(toSerialize)
}

type NullableBTSMSpecificMetadata1315 struct {
	value *BTSMSpecificMetadata1315
	isSet bool
}

func (v NullableBTSMSpecificMetadata1315) Get() *BTSMSpecificMetadata1315 {
	return v.value
}

func (v *NullableBTSMSpecificMetadata1315) Set(val *BTSMSpecificMetadata1315) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSMSpecificMetadata1315) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSMSpecificMetadata1315) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSMSpecificMetadata1315(val *BTSMSpecificMetadata1315) *NullableBTSMSpecificMetadata1315 {
	return &NullableBTSMSpecificMetadata1315{value: val, isSet: true}
}

func (v NullableBTSMSpecificMetadata1315) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSMSpecificMetadata1315) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
