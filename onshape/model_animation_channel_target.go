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

// AnimationChannelTarget struct for AnimationChannelTarget
type AnimationChannelTarget struct {
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras     *map[string]interface{}           `json:"extras,omitempty"`
	Node       *int32                            `json:"node,omitempty"`
	Path       *string                           `json:"path,omitempty"`
}

// NewAnimationChannelTarget instantiates a new AnimationChannelTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnimationChannelTarget() *AnimationChannelTarget {
	this := AnimationChannelTarget{}
	return &this
}

// NewAnimationChannelTargetWithDefaults instantiates a new AnimationChannelTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnimationChannelTargetWithDefaults() *AnimationChannelTarget {
	this := AnimationChannelTarget{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *AnimationChannelTarget) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnimationChannelTarget) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *AnimationChannelTarget) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *AnimationChannelTarget) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *AnimationChannelTarget) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnimationChannelTarget) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *AnimationChannelTarget) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *AnimationChannelTarget) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *AnimationChannelTarget) GetNode() int32 {
	if o == nil || o.Node == nil {
		var ret int32
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnimationChannelTarget) GetNodeOk() (*int32, bool) {
	if o == nil || o.Node == nil {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *AnimationChannelTarget) HasNode() bool {
	if o != nil && o.Node != nil {
		return true
	}

	return false
}

// SetNode gets a reference to the given int32 and assigns it to the Node field.
func (o *AnimationChannelTarget) SetNode(v int32) {
	o.Node = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *AnimationChannelTarget) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnimationChannelTarget) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *AnimationChannelTarget) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *AnimationChannelTarget) SetPath(v string) {
	o.Path = &v
}

func (o AnimationChannelTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Node != nil {
		toSerialize["node"] = o.Node
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableAnimationChannelTarget struct {
	value *AnimationChannelTarget
	isSet bool
}

func (v NullableAnimationChannelTarget) Get() *AnimationChannelTarget {
	return v.value
}

func (v *NullableAnimationChannelTarget) Set(val *AnimationChannelTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableAnimationChannelTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableAnimationChannelTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnimationChannelTarget(val *AnimationChannelTarget) *NullableAnimationChannelTarget {
	return &NullableAnimationChannelTarget{value: val, isSet: true}
}

func (v NullableAnimationChannelTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnimationChannelTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
