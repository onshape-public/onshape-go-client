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

// Material struct for Material
type Material struct {
	AlphaCutoff          *float32                          `json:"alphaCutoff,omitempty"`
	AlphaMode            *string                           `json:"alphaMode,omitempty"`
	DoubleSided          *bool                             `json:"doubleSided,omitempty"`
	EmissiveFactor       []float32                         `json:"emissiveFactor,omitempty"`
	EmissiveTexture      *TextureInfo                      `json:"emissiveTexture,omitempty"`
	Extensions           map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras               *map[string]interface{}           `json:"extras,omitempty"`
	Name                 *string                           `json:"name,omitempty"`
	NormalTexture        *MaterialNormalTextureInfo        `json:"normalTexture,omitempty"`
	OcclusionTexture     *MaterialOcclusionTextureInfo     `json:"occlusionTexture,omitempty"`
	PbrMetallicRoughness *MaterialPbrMetallicRoughness     `json:"pbrMetallicRoughness,omitempty"`
}

// NewMaterial instantiates a new Material object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaterial() *Material {
	this := Material{}
	return &this
}

// NewMaterialWithDefaults instantiates a new Material object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaterialWithDefaults() *Material {
	this := Material{}
	return &this
}

// GetAlphaCutoff returns the AlphaCutoff field value if set, zero value otherwise.
func (o *Material) GetAlphaCutoff() float32 {
	if o == nil || o.AlphaCutoff == nil {
		var ret float32
		return ret
	}
	return *o.AlphaCutoff
}

// GetAlphaCutoffOk returns a tuple with the AlphaCutoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Material) GetAlphaCutoffOk() (*float32, bool) {
	if o == nil || o.AlphaCutoff == nil {
		return nil, false
	}
	return o.AlphaCutoff, true
}

// HasAlphaCutoff returns a boolean if a field has been set.
func (o *Material) HasAlphaCutoff() bool {
	if o != nil && o.AlphaCutoff != nil {
		return true
	}

	return false
}

// SetAlphaCutoff gets a reference to the given float32 and assigns it to the AlphaCutoff field.
func (o *Material) SetAlphaCutoff(v float32) {
	o.AlphaCutoff = &v
}

// GetAlphaMode returns the AlphaMode field value if set, zero value otherwise.
func (o *Material) GetAlphaMode() string {
	if o == nil || o.AlphaMode == nil {
		var ret string
		return ret
	}
	return *o.AlphaMode
}

// GetAlphaModeOk returns a tuple with the AlphaMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Material) GetAlphaModeOk() (*string, bool) {
	if o == nil || o.AlphaMode == nil {
		return nil, false
	}
	return o.AlphaMode, true
}

// HasAlphaMode returns a boolean if a field has been set.
func (o *Material) HasAlphaMode() bool {
	if o != nil && o.AlphaMode != nil {
		return true
	}

	return false
}

// SetAlphaMode gets a reference to the given string and assigns it to the AlphaMode field.
func (o *Material) SetAlphaMode(v string) {
	o.AlphaMode = &v
}

// GetDoubleSided returns the DoubleSided field value if set, zero value otherwise.
func (o *Material) GetDoubleSided() bool {
	if o == nil || o.DoubleSided == nil {
		var ret bool
		return ret
	}
	return *o.DoubleSided
}

// GetDoubleSidedOk returns a tuple with the DoubleSided field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Material) GetDoubleSidedOk() (*bool, bool) {
	if o == nil || o.DoubleSided == nil {
		return nil, false
	}
	return o.DoubleSided, true
}

// HasDoubleSided returns a boolean if a field has been set.
func (o *Material) HasDoubleSided() bool {
	if o != nil && o.DoubleSided != nil {
		return true
	}

	return false
}

// SetDoubleSided gets a reference to the given bool and assigns it to the DoubleSided field.
func (o *Material) SetDoubleSided(v bool) {
	o.DoubleSided = &v
}

// GetEmissiveFactor returns the EmissiveFactor field value if set, zero value otherwise.
func (o *Material) GetEmissiveFactor() []float32 {
	if o == nil || o.EmissiveFactor == nil {
		var ret []float32
		return ret
	}
	return o.EmissiveFactor
}

// GetEmissiveFactorOk returns a tuple with the EmissiveFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Material) GetEmissiveFactorOk() ([]float32, bool) {
	if o == nil || o.EmissiveFactor == nil {
		return nil, false
	}
	return o.EmissiveFactor, true
}

// HasEmissiveFactor returns a boolean if a field has been set.
func (o *Material) HasEmissiveFactor() bool {
	if o != nil && o.EmissiveFactor != nil {
		return true
	}

	return false
}

// SetEmissiveFactor gets a reference to the given []float32 and assigns it to the EmissiveFactor field.
func (o *Material) SetEmissiveFactor(v []float32) {
	o.EmissiveFactor = v
}

// GetEmissiveTexture returns the EmissiveTexture field value if set, zero value otherwise.
func (o *Material) GetEmissiveTexture() TextureInfo {
	if o == nil || o.EmissiveTexture == nil {
		var ret TextureInfo
		return ret
	}
	return *o.EmissiveTexture
}

// GetEmissiveTextureOk returns a tuple with the EmissiveTexture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Material) GetEmissiveTextureOk() (*TextureInfo, bool) {
	if o == nil || o.EmissiveTexture == nil {
		return nil, false
	}
	return o.EmissiveTexture, true
}

// HasEmissiveTexture returns a boolean if a field has been set.
func (o *Material) HasEmissiveTexture() bool {
	if o != nil && o.EmissiveTexture != nil {
		return true
	}

	return false
}

// SetEmissiveTexture gets a reference to the given TextureInfo and assigns it to the EmissiveTexture field.
func (o *Material) SetEmissiveTexture(v TextureInfo) {
	o.EmissiveTexture = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Material) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Material) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Material) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Material) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *Material) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Material) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *Material) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *Material) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Material) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Material) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Material) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Material) SetName(v string) {
	o.Name = &v
}

// GetNormalTexture returns the NormalTexture field value if set, zero value otherwise.
func (o *Material) GetNormalTexture() MaterialNormalTextureInfo {
	if o == nil || o.NormalTexture == nil {
		var ret MaterialNormalTextureInfo
		return ret
	}
	return *o.NormalTexture
}

// GetNormalTextureOk returns a tuple with the NormalTexture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Material) GetNormalTextureOk() (*MaterialNormalTextureInfo, bool) {
	if o == nil || o.NormalTexture == nil {
		return nil, false
	}
	return o.NormalTexture, true
}

// HasNormalTexture returns a boolean if a field has been set.
func (o *Material) HasNormalTexture() bool {
	if o != nil && o.NormalTexture != nil {
		return true
	}

	return false
}

// SetNormalTexture gets a reference to the given MaterialNormalTextureInfo and assigns it to the NormalTexture field.
func (o *Material) SetNormalTexture(v MaterialNormalTextureInfo) {
	o.NormalTexture = &v
}

// GetOcclusionTexture returns the OcclusionTexture field value if set, zero value otherwise.
func (o *Material) GetOcclusionTexture() MaterialOcclusionTextureInfo {
	if o == nil || o.OcclusionTexture == nil {
		var ret MaterialOcclusionTextureInfo
		return ret
	}
	return *o.OcclusionTexture
}

// GetOcclusionTextureOk returns a tuple with the OcclusionTexture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Material) GetOcclusionTextureOk() (*MaterialOcclusionTextureInfo, bool) {
	if o == nil || o.OcclusionTexture == nil {
		return nil, false
	}
	return o.OcclusionTexture, true
}

// HasOcclusionTexture returns a boolean if a field has been set.
func (o *Material) HasOcclusionTexture() bool {
	if o != nil && o.OcclusionTexture != nil {
		return true
	}

	return false
}

// SetOcclusionTexture gets a reference to the given MaterialOcclusionTextureInfo and assigns it to the OcclusionTexture field.
func (o *Material) SetOcclusionTexture(v MaterialOcclusionTextureInfo) {
	o.OcclusionTexture = &v
}

// GetPbrMetallicRoughness returns the PbrMetallicRoughness field value if set, zero value otherwise.
func (o *Material) GetPbrMetallicRoughness() MaterialPbrMetallicRoughness {
	if o == nil || o.PbrMetallicRoughness == nil {
		var ret MaterialPbrMetallicRoughness
		return ret
	}
	return *o.PbrMetallicRoughness
}

// GetPbrMetallicRoughnessOk returns a tuple with the PbrMetallicRoughness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Material) GetPbrMetallicRoughnessOk() (*MaterialPbrMetallicRoughness, bool) {
	if o == nil || o.PbrMetallicRoughness == nil {
		return nil, false
	}
	return o.PbrMetallicRoughness, true
}

// HasPbrMetallicRoughness returns a boolean if a field has been set.
func (o *Material) HasPbrMetallicRoughness() bool {
	if o != nil && o.PbrMetallicRoughness != nil {
		return true
	}

	return false
}

// SetPbrMetallicRoughness gets a reference to the given MaterialPbrMetallicRoughness and assigns it to the PbrMetallicRoughness field.
func (o *Material) SetPbrMetallicRoughness(v MaterialPbrMetallicRoughness) {
	o.PbrMetallicRoughness = &v
}

func (o Material) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AlphaCutoff != nil {
		toSerialize["alphaCutoff"] = o.AlphaCutoff
	}
	if o.AlphaMode != nil {
		toSerialize["alphaMode"] = o.AlphaMode
	}
	if o.DoubleSided != nil {
		toSerialize["doubleSided"] = o.DoubleSided
	}
	if o.EmissiveFactor != nil {
		toSerialize["emissiveFactor"] = o.EmissiveFactor
	}
	if o.EmissiveTexture != nil {
		toSerialize["emissiveTexture"] = o.EmissiveTexture
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NormalTexture != nil {
		toSerialize["normalTexture"] = o.NormalTexture
	}
	if o.OcclusionTexture != nil {
		toSerialize["occlusionTexture"] = o.OcclusionTexture
	}
	if o.PbrMetallicRoughness != nil {
		toSerialize["pbrMetallicRoughness"] = o.PbrMetallicRoughness
	}
	return json.Marshal(toSerialize)
}

type NullableMaterial struct {
	value *Material
	isSet bool
}

func (v NullableMaterial) Get() *Material {
	return v.value
}

func (v *NullableMaterial) Set(val *Material) {
	v.value = val
	v.isSet = true
}

func (v NullableMaterial) IsSet() bool {
	return v.isSet
}

func (v *NullableMaterial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaterial(val *Material) *NullableMaterial {
	return &NullableMaterial{value: val, isSet: true}
}

func (v NullableMaterial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaterial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
