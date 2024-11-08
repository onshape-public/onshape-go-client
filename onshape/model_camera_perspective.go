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

// CameraPerspective struct for CameraPerspective
type CameraPerspective struct {
	AspectRatio *float32                          `json:"aspectRatio,omitempty"`
	Extensions  map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras      map[string]interface{}            `json:"extras,omitempty"`
	Yfov        *float32                          `json:"yfov,omitempty"`
	Zfar        *float32                          `json:"zfar,omitempty"`
	Znear       *float32                          `json:"znear,omitempty"`
}

// NewCameraPerspective instantiates a new CameraPerspective object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCameraPerspective() *CameraPerspective {
	this := CameraPerspective{}
	return &this
}

// NewCameraPerspectiveWithDefaults instantiates a new CameraPerspective object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCameraPerspectiveWithDefaults() *CameraPerspective {
	this := CameraPerspective{}
	return &this
}

// GetAspectRatio returns the AspectRatio field value if set, zero value otherwise.
func (o *CameraPerspective) GetAspectRatio() float32 {
	if o == nil || o.AspectRatio == nil {
		var ret float32
		return ret
	}
	return *o.AspectRatio
}

// GetAspectRatioOk returns a tuple with the AspectRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraPerspective) GetAspectRatioOk() (*float32, bool) {
	if o == nil || o.AspectRatio == nil {
		return nil, false
	}
	return o.AspectRatio, true
}

// HasAspectRatio returns a boolean if a field has been set.
func (o *CameraPerspective) HasAspectRatio() bool {
	if o != nil && o.AspectRatio != nil {
		return true
	}

	return false
}

// SetAspectRatio gets a reference to the given float32 and assigns it to the AspectRatio field.
func (o *CameraPerspective) SetAspectRatio(v float32) {
	o.AspectRatio = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *CameraPerspective) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraPerspective) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *CameraPerspective) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *CameraPerspective) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *CameraPerspective) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraPerspective) GetExtrasOk() (map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *CameraPerspective) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *CameraPerspective) SetExtras(v map[string]interface{}) {
	o.Extras = v
}

// GetYfov returns the Yfov field value if set, zero value otherwise.
func (o *CameraPerspective) GetYfov() float32 {
	if o == nil || o.Yfov == nil {
		var ret float32
		return ret
	}
	return *o.Yfov
}

// GetYfovOk returns a tuple with the Yfov field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraPerspective) GetYfovOk() (*float32, bool) {
	if o == nil || o.Yfov == nil {
		return nil, false
	}
	return o.Yfov, true
}

// HasYfov returns a boolean if a field has been set.
func (o *CameraPerspective) HasYfov() bool {
	if o != nil && o.Yfov != nil {
		return true
	}

	return false
}

// SetYfov gets a reference to the given float32 and assigns it to the Yfov field.
func (o *CameraPerspective) SetYfov(v float32) {
	o.Yfov = &v
}

// GetZfar returns the Zfar field value if set, zero value otherwise.
func (o *CameraPerspective) GetZfar() float32 {
	if o == nil || o.Zfar == nil {
		var ret float32
		return ret
	}
	return *o.Zfar
}

// GetZfarOk returns a tuple with the Zfar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraPerspective) GetZfarOk() (*float32, bool) {
	if o == nil || o.Zfar == nil {
		return nil, false
	}
	return o.Zfar, true
}

// HasZfar returns a boolean if a field has been set.
func (o *CameraPerspective) HasZfar() bool {
	if o != nil && o.Zfar != nil {
		return true
	}

	return false
}

// SetZfar gets a reference to the given float32 and assigns it to the Zfar field.
func (o *CameraPerspective) SetZfar(v float32) {
	o.Zfar = &v
}

// GetZnear returns the Znear field value if set, zero value otherwise.
func (o *CameraPerspective) GetZnear() float32 {
	if o == nil || o.Znear == nil {
		var ret float32
		return ret
	}
	return *o.Znear
}

// GetZnearOk returns a tuple with the Znear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraPerspective) GetZnearOk() (*float32, bool) {
	if o == nil || o.Znear == nil {
		return nil, false
	}
	return o.Znear, true
}

// HasZnear returns a boolean if a field has been set.
func (o *CameraPerspective) HasZnear() bool {
	if o != nil && o.Znear != nil {
		return true
	}

	return false
}

// SetZnear gets a reference to the given float32 and assigns it to the Znear field.
func (o *CameraPerspective) SetZnear(v float32) {
	o.Znear = &v
}

func (o CameraPerspective) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AspectRatio != nil {
		toSerialize["aspectRatio"] = o.AspectRatio
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Yfov != nil {
		toSerialize["yfov"] = o.Yfov
	}
	if o.Zfar != nil {
		toSerialize["zfar"] = o.Zfar
	}
	if o.Znear != nil {
		toSerialize["znear"] = o.Znear
	}
	return json.Marshal(toSerialize)
}

type NullableCameraPerspective struct {
	value *CameraPerspective
	isSet bool
}

func (v NullableCameraPerspective) Get() *CameraPerspective {
	return v.value
}

func (v *NullableCameraPerspective) Set(val *CameraPerspective) {
	v.value = val
	v.isSet = true
}

func (v NullableCameraPerspective) IsSet() bool {
	return v.isSet
}

func (v *NullableCameraPerspective) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCameraPerspective(val *CameraPerspective) *NullableCameraPerspective {
	return &NullableCameraPerspective{value: val, isSet: true}
}

func (v NullableCameraPerspective) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCameraPerspective) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
