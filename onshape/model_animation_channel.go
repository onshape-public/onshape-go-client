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

// AnimationChannel struct for AnimationChannel
type AnimationChannel struct {
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras     *map[string]interface{}           `json:"extras,omitempty"`
	Sampler    *int32                            `json:"sampler,omitempty"`
	Target     *AnimationChannelTarget           `json:"target,omitempty"`
}

// NewAnimationChannel instantiates a new AnimationChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnimationChannel() *AnimationChannel {
	this := AnimationChannel{}
	return &this
}

// NewAnimationChannelWithDefaults instantiates a new AnimationChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnimationChannelWithDefaults() *AnimationChannel {
	this := AnimationChannel{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *AnimationChannel) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnimationChannel) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *AnimationChannel) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *AnimationChannel) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *AnimationChannel) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnimationChannel) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *AnimationChannel) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *AnimationChannel) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetSampler returns the Sampler field value if set, zero value otherwise.
func (o *AnimationChannel) GetSampler() int32 {
	if o == nil || o.Sampler == nil {
		var ret int32
		return ret
	}
	return *o.Sampler
}

// GetSamplerOk returns a tuple with the Sampler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnimationChannel) GetSamplerOk() (*int32, bool) {
	if o == nil || o.Sampler == nil {
		return nil, false
	}
	return o.Sampler, true
}

// HasSampler returns a boolean if a field has been set.
func (o *AnimationChannel) HasSampler() bool {
	if o != nil && o.Sampler != nil {
		return true
	}

	return false
}

// SetSampler gets a reference to the given int32 and assigns it to the Sampler field.
func (o *AnimationChannel) SetSampler(v int32) {
	o.Sampler = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *AnimationChannel) GetTarget() AnimationChannelTarget {
	if o == nil || o.Target == nil {
		var ret AnimationChannelTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnimationChannel) GetTargetOk() (*AnimationChannelTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *AnimationChannel) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AnimationChannelTarget and assigns it to the Target field.
func (o *AnimationChannel) SetTarget(v AnimationChannelTarget) {
	o.Target = &v
}

func (o AnimationChannel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Sampler != nil {
		toSerialize["sampler"] = o.Sampler
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableAnimationChannel struct {
	value *AnimationChannel
	isSet bool
}

func (v NullableAnimationChannel) Get() *AnimationChannel {
	return v.value
}

func (v *NullableAnimationChannel) Set(val *AnimationChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableAnimationChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableAnimationChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnimationChannel(val *AnimationChannel) *NullableAnimationChannel {
	return &NullableAnimationChannel{value: val, isSet: true}
}

func (v NullableAnimationChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnimationChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
