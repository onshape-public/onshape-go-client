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

// BTConfigurationParams struct for BTConfigurationParams
type BTConfigurationParams struct {
	Parameters                  []ConfigurationEntry `json:"parameters,omitempty"`
	StandardContentParametersId *string              `json:"standardContentParametersId,omitempty"`
}

// NewBTConfigurationParams instantiates a new BTConfigurationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfigurationParams() *BTConfigurationParams {
	this := BTConfigurationParams{}
	return &this
}

// NewBTConfigurationParamsWithDefaults instantiates a new BTConfigurationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfigurationParamsWithDefaults() *BTConfigurationParams {
	this := BTConfigurationParams{}
	return &this
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTConfigurationParams) GetParameters() []ConfigurationEntry {
	if o == nil || o.Parameters == nil {
		var ret []ConfigurationEntry
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationParams) GetParametersOk() ([]ConfigurationEntry, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTConfigurationParams) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ConfigurationEntry and assigns it to the Parameters field.
func (o *BTConfigurationParams) SetParameters(v []ConfigurationEntry) {
	o.Parameters = v
}

// GetStandardContentParametersId returns the StandardContentParametersId field value if set, zero value otherwise.
func (o *BTConfigurationParams) GetStandardContentParametersId() string {
	if o == nil || o.StandardContentParametersId == nil {
		var ret string
		return ret
	}
	return *o.StandardContentParametersId
}

// GetStandardContentParametersIdOk returns a tuple with the StandardContentParametersId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationParams) GetStandardContentParametersIdOk() (*string, bool) {
	if o == nil || o.StandardContentParametersId == nil {
		return nil, false
	}
	return o.StandardContentParametersId, true
}

// HasStandardContentParametersId returns a boolean if a field has been set.
func (o *BTConfigurationParams) HasStandardContentParametersId() bool {
	if o != nil && o.StandardContentParametersId != nil {
		return true
	}

	return false
}

// SetStandardContentParametersId gets a reference to the given string and assigns it to the StandardContentParametersId field.
func (o *BTConfigurationParams) SetStandardContentParametersId(v string) {
	o.StandardContentParametersId = &v
}

func (o BTConfigurationParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.StandardContentParametersId != nil {
		toSerialize["standardContentParametersId"] = o.StandardContentParametersId
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfigurationParams struct {
	value *BTConfigurationParams
	isSet bool
}

func (v NullableBTConfigurationParams) Get() *BTConfigurationParams {
	return v.value
}

func (v *NullableBTConfigurationParams) Set(val *BTConfigurationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfigurationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfigurationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfigurationParams(val *BTConfigurationParams) *NullableBTConfigurationParams {
	return &NullableBTConfigurationParams{value: val, isSet: true}
}

func (v NullableBTConfigurationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfigurationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
