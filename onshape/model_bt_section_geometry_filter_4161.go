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

// BTSectionGeometryFilter4161 struct for BTSectionGeometryFilter4161
type BTSectionGeometryFilter4161 struct {
	BTQueryFilter183
	BtType            *string `json:"btType,omitempty"`
	IsSectionGeometry *bool   `json:"isSectionGeometry,omitempty"`
	// Type of JSON object.
	BtType *string `json:"btType,omitempty"`
}

// NewBTSectionGeometryFilter4161 instantiates a new BTSectionGeometryFilter4161 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSectionGeometryFilter4161() *BTSectionGeometryFilter4161 {
	this := BTSectionGeometryFilter4161{}
	return &this
}

// NewBTSectionGeometryFilter4161WithDefaults instantiates a new BTSectionGeometryFilter4161 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSectionGeometryFilter4161WithDefaults() *BTSectionGeometryFilter4161 {
	this := BTSectionGeometryFilter4161{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSectionGeometryFilter4161) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSectionGeometryFilter4161) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSectionGeometryFilter4161) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSectionGeometryFilter4161) SetBtType(v string) {
	o.BtType = &v
}

// GetIsSectionGeometry returns the IsSectionGeometry field value if set, zero value otherwise.
func (o *BTSectionGeometryFilter4161) GetIsSectionGeometry() bool {
	if o == nil || o.IsSectionGeometry == nil {
		var ret bool
		return ret
	}
	return *o.IsSectionGeometry
}

// GetIsSectionGeometryOk returns a tuple with the IsSectionGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSectionGeometryFilter4161) GetIsSectionGeometryOk() (*bool, bool) {
	if o == nil || o.IsSectionGeometry == nil {
		return nil, false
	}
	return o.IsSectionGeometry, true
}

// HasIsSectionGeometry returns a boolean if a field has been set.
func (o *BTSectionGeometryFilter4161) HasIsSectionGeometry() bool {
	if o != nil && o.IsSectionGeometry != nil {
		return true
	}

	return false
}

// SetIsSectionGeometry gets a reference to the given bool and assigns it to the IsSectionGeometry field.
func (o *BTSectionGeometryFilter4161) SetIsSectionGeometry(v bool) {
	o.IsSectionGeometry = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSectionGeometryFilter4161) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSectionGeometryFilter4161) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSectionGeometryFilter4161) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSectionGeometryFilter4161) SetBtType(v string) {
	o.BtType = &v
}

func (o BTSectionGeometryFilter4161) MarshalJSON() ([]byte, error) {
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
	if o.IsSectionGeometry != nil {
		toSerialize["isSectionGeometry"] = o.IsSectionGeometry
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTSectionGeometryFilter4161 struct {
	value *BTSectionGeometryFilter4161
	isSet bool
}

func (v NullableBTSectionGeometryFilter4161) Get() *BTSectionGeometryFilter4161 {
	return v.value
}

func (v *NullableBTSectionGeometryFilter4161) Set(val *BTSectionGeometryFilter4161) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSectionGeometryFilter4161) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSectionGeometryFilter4161) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSectionGeometryFilter4161(val *BTSectionGeometryFilter4161) *NullableBTSectionGeometryFilter4161 {
	return &NullableBTSectionGeometryFilter4161{value: val, isSet: true}
}

func (v NullableBTSectionGeometryFilter4161) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSectionGeometryFilter4161) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
