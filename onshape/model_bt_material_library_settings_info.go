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

// BTMaterialLibrarySettingsInfo struct for BTMaterialLibrarySettingsInfo
type BTMaterialLibrarySettingsInfo struct {
	CompanyLibraries []BTMaterialLibraryMetadataInfo `json:"companyLibraries,omitempty"`
	Libraries        []BTMaterialLibraryMetadataInfo `json:"libraries,omitempty"`
}

// NewBTMaterialLibrarySettingsInfo instantiates a new BTMaterialLibrarySettingsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMaterialLibrarySettingsInfo() *BTMaterialLibrarySettingsInfo {
	this := BTMaterialLibrarySettingsInfo{}
	return &this
}

// NewBTMaterialLibrarySettingsInfoWithDefaults instantiates a new BTMaterialLibrarySettingsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMaterialLibrarySettingsInfoWithDefaults() *BTMaterialLibrarySettingsInfo {
	this := BTMaterialLibrarySettingsInfo{}
	return &this
}

// GetCompanyLibraries returns the CompanyLibraries field value if set, zero value otherwise.
func (o *BTMaterialLibrarySettingsInfo) GetCompanyLibraries() []BTMaterialLibraryMetadataInfo {
	if o == nil || o.CompanyLibraries == nil {
		var ret []BTMaterialLibraryMetadataInfo
		return ret
	}
	return o.CompanyLibraries
}

// GetCompanyLibrariesOk returns a tuple with the CompanyLibraries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialLibrarySettingsInfo) GetCompanyLibrariesOk() ([]BTMaterialLibraryMetadataInfo, bool) {
	if o == nil || o.CompanyLibraries == nil {
		return nil, false
	}
	return o.CompanyLibraries, true
}

// HasCompanyLibraries returns a boolean if a field has been set.
func (o *BTMaterialLibrarySettingsInfo) HasCompanyLibraries() bool {
	if o != nil && o.CompanyLibraries != nil {
		return true
	}

	return false
}

// SetCompanyLibraries gets a reference to the given []BTMaterialLibraryMetadataInfo and assigns it to the CompanyLibraries field.
func (o *BTMaterialLibrarySettingsInfo) SetCompanyLibraries(v []BTMaterialLibraryMetadataInfo) {
	o.CompanyLibraries = v
}

// GetLibraries returns the Libraries field value if set, zero value otherwise.
func (o *BTMaterialLibrarySettingsInfo) GetLibraries() []BTMaterialLibraryMetadataInfo {
	if o == nil || o.Libraries == nil {
		var ret []BTMaterialLibraryMetadataInfo
		return ret
	}
	return o.Libraries
}

// GetLibrariesOk returns a tuple with the Libraries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialLibrarySettingsInfo) GetLibrariesOk() ([]BTMaterialLibraryMetadataInfo, bool) {
	if o == nil || o.Libraries == nil {
		return nil, false
	}
	return o.Libraries, true
}

// HasLibraries returns a boolean if a field has been set.
func (o *BTMaterialLibrarySettingsInfo) HasLibraries() bool {
	if o != nil && o.Libraries != nil {
		return true
	}

	return false
}

// SetLibraries gets a reference to the given []BTMaterialLibraryMetadataInfo and assigns it to the Libraries field.
func (o *BTMaterialLibrarySettingsInfo) SetLibraries(v []BTMaterialLibraryMetadataInfo) {
	o.Libraries = v
}

func (o BTMaterialLibrarySettingsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyLibraries != nil {
		toSerialize["companyLibraries"] = o.CompanyLibraries
	}
	if o.Libraries != nil {
		toSerialize["libraries"] = o.Libraries
	}
	return json.Marshal(toSerialize)
}

type NullableBTMaterialLibrarySettingsInfo struct {
	value *BTMaterialLibrarySettingsInfo
	isSet bool
}

func (v NullableBTMaterialLibrarySettingsInfo) Get() *BTMaterialLibrarySettingsInfo {
	return v.value
}

func (v *NullableBTMaterialLibrarySettingsInfo) Set(val *BTMaterialLibrarySettingsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMaterialLibrarySettingsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMaterialLibrarySettingsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMaterialLibrarySettingsInfo(val *BTMaterialLibrarySettingsInfo) *NullableBTMaterialLibrarySettingsInfo {
	return &NullableBTMaterialLibrarySettingsInfo{value: val, isSet: true}
}

func (v NullableBTMaterialLibrarySettingsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMaterialLibrarySettingsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
