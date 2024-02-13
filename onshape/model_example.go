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

// Example struct for Example
type Example struct {
	Description   *string                           `json:"description,omitempty"`
	Extensions    map[string]map[string]interface{} `json:"extensions,omitempty"`
	ExternalValue *string                           `json:"externalValue,omitempty"`
	Getref        *string                           `json:"get$ref,omitempty"`
	Summary       *string                           `json:"summary,omitempty"`
	Value         *map[string]interface{}           `json:"value,omitempty"`
	ValueSetFlag  *bool                             `json:"valueSetFlag,omitempty"`
}

// NewExample instantiates a new Example object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExample() *Example {
	this := Example{}
	return &this
}

// NewExampleWithDefaults instantiates a new Example object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExampleWithDefaults() *Example {
	this := Example{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Example) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Example) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Example) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Example) SetDescription(v string) {
	o.Description = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Example) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Example) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Example) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Example) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExternalValue returns the ExternalValue field value if set, zero value otherwise.
func (o *Example) GetExternalValue() string {
	if o == nil || o.ExternalValue == nil {
		var ret string
		return ret
	}
	return *o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Example) GetExternalValueOk() (*string, bool) {
	if o == nil || o.ExternalValue == nil {
		return nil, false
	}
	return o.ExternalValue, true
}

// HasExternalValue returns a boolean if a field has been set.
func (o *Example) HasExternalValue() bool {
	if o != nil && o.ExternalValue != nil {
		return true
	}

	return false
}

// SetExternalValue gets a reference to the given string and assigns it to the ExternalValue field.
func (o *Example) SetExternalValue(v string) {
	o.ExternalValue = &v
}

// GetGetref returns the Getref field value if set, zero value otherwise.
func (o *Example) GetGetref() string {
	if o == nil || o.Getref == nil {
		var ret string
		return ret
	}
	return *o.Getref
}

// GetGetrefOk returns a tuple with the Getref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Example) GetGetrefOk() (*string, bool) {
	if o == nil || o.Getref == nil {
		return nil, false
	}
	return o.Getref, true
}

// HasGetref returns a boolean if a field has been set.
func (o *Example) HasGetref() bool {
	if o != nil && o.Getref != nil {
		return true
	}

	return false
}

// SetGetref gets a reference to the given string and assigns it to the Getref field.
func (o *Example) SetGetref(v string) {
	o.Getref = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *Example) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Example) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *Example) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *Example) SetSummary(v string) {
	o.Summary = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Example) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Example) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Example) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *Example) SetValue(v map[string]interface{}) {
	o.Value = &v
}

// GetValueSetFlag returns the ValueSetFlag field value if set, zero value otherwise.
func (o *Example) GetValueSetFlag() bool {
	if o == nil || o.ValueSetFlag == nil {
		var ret bool
		return ret
	}
	return *o.ValueSetFlag
}

// GetValueSetFlagOk returns a tuple with the ValueSetFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Example) GetValueSetFlagOk() (*bool, bool) {
	if o == nil || o.ValueSetFlag == nil {
		return nil, false
	}
	return o.ValueSetFlag, true
}

// HasValueSetFlag returns a boolean if a field has been set.
func (o *Example) HasValueSetFlag() bool {
	if o != nil && o.ValueSetFlag != nil {
		return true
	}

	return false
}

// SetValueSetFlag gets a reference to the given bool and assigns it to the ValueSetFlag field.
func (o *Example) SetValueSetFlag(v bool) {
	o.ValueSetFlag = &v
}

func (o Example) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.ExternalValue != nil {
		toSerialize["externalValue"] = o.ExternalValue
	}
	if o.Getref != nil {
		toSerialize["get$ref"] = o.Getref
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ValueSetFlag != nil {
		toSerialize["valueSetFlag"] = o.ValueSetFlag
	}
	return json.Marshal(toSerialize)
}

type NullableExample struct {
	value *Example
	isSet bool
}

func (v NullableExample) Get() *Example {
	return v.value
}

func (v *NullableExample) Set(val *Example) {
	v.value = val
	v.isSet = true
}

func (v NullableExample) IsSet() bool {
	return v.isSet
}

func (v *NullableExample) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExample(val *Example) *NullableExample {
	return &NullableExample{value: val, isSet: true}
}

func (v NullableExample) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExample) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
