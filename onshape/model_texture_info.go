/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.171.24257-687de06de652
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// TextureInfo struct for TextureInfo
type TextureInfo struct {
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras     *map[string]interface{}           `json:"extras,omitempty"`
	Index      *int32                            `json:"index,omitempty"`
	TexCoord   *int32                            `json:"texCoord,omitempty"`
}

// NewTextureInfo instantiates a new TextureInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextureInfo() *TextureInfo {
	this := TextureInfo{}
	return &this
}

// NewTextureInfoWithDefaults instantiates a new TextureInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextureInfoWithDefaults() *TextureInfo {
	this := TextureInfo{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *TextureInfo) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextureInfo) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *TextureInfo) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *TextureInfo) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *TextureInfo) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextureInfo) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *TextureInfo) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *TextureInfo) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *TextureInfo) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextureInfo) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *TextureInfo) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *TextureInfo) SetIndex(v int32) {
	o.Index = &v
}

// GetTexCoord returns the TexCoord field value if set, zero value otherwise.
func (o *TextureInfo) GetTexCoord() int32 {
	if o == nil || o.TexCoord == nil {
		var ret int32
		return ret
	}
	return *o.TexCoord
}

// GetTexCoordOk returns a tuple with the TexCoord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextureInfo) GetTexCoordOk() (*int32, bool) {
	if o == nil || o.TexCoord == nil {
		return nil, false
	}
	return o.TexCoord, true
}

// HasTexCoord returns a boolean if a field has been set.
func (o *TextureInfo) HasTexCoord() bool {
	if o != nil && o.TexCoord != nil {
		return true
	}

	return false
}

// SetTexCoord gets a reference to the given int32 and assigns it to the TexCoord field.
func (o *TextureInfo) SetTexCoord(v int32) {
	o.TexCoord = &v
}

func (o TextureInfo) MarshalJSON() ([]byte, error) {
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
	if o.TexCoord != nil {
		toSerialize["texCoord"] = o.TexCoord
	}
	return json.Marshal(toSerialize)
}

type NullableTextureInfo struct {
	value *TextureInfo
	isSet bool
}

func (v NullableTextureInfo) Get() *TextureInfo {
	return v.value
}

func (v *NullableTextureInfo) Set(val *TextureInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTextureInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTextureInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextureInfo(val *TextureInfo) *NullableTextureInfo {
	return &NullableTextureInfo{value: val, isSet: true}
}

func (v NullableTextureInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextureInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
