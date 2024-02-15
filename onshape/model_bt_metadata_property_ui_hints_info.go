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

// BTMetadataPropertyUiHintsInfo struct for BTMetadataPropertyUiHintsInfo
type BTMetadataPropertyUiHintsInfo struct {
	Multiline *bool `json:"multiline,omitempty"`
}

// NewBTMetadataPropertyUiHintsInfo instantiates a new BTMetadataPropertyUiHintsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMetadataPropertyUiHintsInfo() *BTMetadataPropertyUiHintsInfo {
	this := BTMetadataPropertyUiHintsInfo{}
	return &this
}

// NewBTMetadataPropertyUiHintsInfoWithDefaults instantiates a new BTMetadataPropertyUiHintsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMetadataPropertyUiHintsInfoWithDefaults() *BTMetadataPropertyUiHintsInfo {
	this := BTMetadataPropertyUiHintsInfo{}
	return &this
}

// GetMultiline returns the Multiline field value if set, zero value otherwise.
func (o *BTMetadataPropertyUiHintsInfo) GetMultiline() bool {
	if o == nil || o.Multiline == nil {
		var ret bool
		return ret
	}
	return *o.Multiline
}

// GetMultilineOk returns a tuple with the Multiline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyUiHintsInfo) GetMultilineOk() (*bool, bool) {
	if o == nil || o.Multiline == nil {
		return nil, false
	}
	return o.Multiline, true
}

// HasMultiline returns a boolean if a field has been set.
func (o *BTMetadataPropertyUiHintsInfo) HasMultiline() bool {
	if o != nil && o.Multiline != nil {
		return true
	}

	return false
}

// SetMultiline gets a reference to the given bool and assigns it to the Multiline field.
func (o *BTMetadataPropertyUiHintsInfo) SetMultiline(v bool) {
	o.Multiline = &v
}

func (o BTMetadataPropertyUiHintsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Multiline != nil {
		toSerialize["multiline"] = o.Multiline
	}
	return json.Marshal(toSerialize)
}

type NullableBTMetadataPropertyUiHintsInfo struct {
	value *BTMetadataPropertyUiHintsInfo
	isSet bool
}

func (v NullableBTMetadataPropertyUiHintsInfo) Get() *BTMetadataPropertyUiHintsInfo {
	return v.value
}

func (v *NullableBTMetadataPropertyUiHintsInfo) Set(val *BTMetadataPropertyUiHintsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataPropertyUiHintsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataPropertyUiHintsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataPropertyUiHintsInfo(val *BTMetadataPropertyUiHintsInfo) *NullableBTMetadataPropertyUiHintsInfo {
	return &NullableBTMetadataPropertyUiHintsInfo{value: val, isSet: true}
}

func (v NullableBTMetadataPropertyUiHintsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataPropertyUiHintsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
