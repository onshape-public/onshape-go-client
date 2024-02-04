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

// MaterialOcclusionTextureInfo struct for MaterialOcclusionTextureInfo
type MaterialOcclusionTextureInfo struct {
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras     *map[string]interface{}           `json:"extras,omitempty"`
	Index      *int32                            `json:"index,omitempty"`
	Strength   *float32                          `json:"strength,omitempty"`
	TexCoord   *int32                            `json:"texCoord,omitempty"`
}

// NewMaterialOcclusionTextureInfo instantiates a new MaterialOcclusionTextureInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaterialOcclusionTextureInfo() *MaterialOcclusionTextureInfo {
	this := MaterialOcclusionTextureInfo{}
	return &this
}

// NewMaterialOcclusionTextureInfoWithDefaults instantiates a new MaterialOcclusionTextureInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaterialOcclusionTextureInfoWithDefaults() *MaterialOcclusionTextureInfo {
	this := MaterialOcclusionTextureInfo{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *MaterialOcclusionTextureInfo) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialOcclusionTextureInfo) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *MaterialOcclusionTextureInfo) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *MaterialOcclusionTextureInfo) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *MaterialOcclusionTextureInfo) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialOcclusionTextureInfo) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *MaterialOcclusionTextureInfo) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *MaterialOcclusionTextureInfo) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *MaterialOcclusionTextureInfo) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialOcclusionTextureInfo) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *MaterialOcclusionTextureInfo) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *MaterialOcclusionTextureInfo) SetIndex(v int32) {
	o.Index = &v
}

// GetStrength returns the Strength field value if set, zero value otherwise.
func (o *MaterialOcclusionTextureInfo) GetStrength() float32 {
	if o == nil || o.Strength == nil {
		var ret float32
		return ret
	}
	return *o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialOcclusionTextureInfo) GetStrengthOk() (*float32, bool) {
	if o == nil || o.Strength == nil {
		return nil, false
	}
	return o.Strength, true
}

// HasStrength returns a boolean if a field has been set.
func (o *MaterialOcclusionTextureInfo) HasStrength() bool {
	if o != nil && o.Strength != nil {
		return true
	}

	return false
}

// SetStrength gets a reference to the given float32 and assigns it to the Strength field.
func (o *MaterialOcclusionTextureInfo) SetStrength(v float32) {
	o.Strength = &v
}

// GetTexCoord returns the TexCoord field value if set, zero value otherwise.
func (o *MaterialOcclusionTextureInfo) GetTexCoord() int32 {
	if o == nil || o.TexCoord == nil {
		var ret int32
		return ret
	}
	return *o.TexCoord
}

// GetTexCoordOk returns a tuple with the TexCoord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialOcclusionTextureInfo) GetTexCoordOk() (*int32, bool) {
	if o == nil || o.TexCoord == nil {
		return nil, false
	}
	return o.TexCoord, true
}

// HasTexCoord returns a boolean if a field has been set.
func (o *MaterialOcclusionTextureInfo) HasTexCoord() bool {
	if o != nil && o.TexCoord != nil {
		return true
	}

	return false
}

// SetTexCoord gets a reference to the given int32 and assigns it to the TexCoord field.
func (o *MaterialOcclusionTextureInfo) SetTexCoord(v int32) {
	o.TexCoord = &v
}

func (o MaterialOcclusionTextureInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Strength != nil {
		toSerialize["strength"] = o.Strength
	}
	if o.TexCoord != nil {
		toSerialize["texCoord"] = o.TexCoord
	}
	return json.Marshal(toSerialize)
}

type NullableMaterialOcclusionTextureInfo struct {
	value *MaterialOcclusionTextureInfo
	isSet bool
}

func (v NullableMaterialOcclusionTextureInfo) Get() *MaterialOcclusionTextureInfo {
	return v.value
}

func (v *NullableMaterialOcclusionTextureInfo) Set(val *MaterialOcclusionTextureInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMaterialOcclusionTextureInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMaterialOcclusionTextureInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaterialOcclusionTextureInfo(val *MaterialOcclusionTextureInfo) *NullableMaterialOcclusionTextureInfo {
	return &NullableMaterialOcclusionTextureInfo{value: val, isSet: true}
}

func (v NullableMaterialOcclusionTextureInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaterialOcclusionTextureInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
