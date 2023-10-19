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

// Camera struct for Camera
type Camera struct {
	Extensions   map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras       *map[string]interface{}           `json:"extras,omitempty"`
	Name         *string                           `json:"name,omitempty"`
	Orthographic *CameraOrthographic               `json:"orthographic,omitempty"`
	Perspective  *CameraPerspective                `json:"perspective,omitempty"`
	Type         *string                           `json:"type,omitempty"`
}

// NewCamera instantiates a new Camera object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCamera() *Camera {
	this := Camera{}
	return &this
}

// NewCameraWithDefaults instantiates a new Camera object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCameraWithDefaults() *Camera {
	this := Camera{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Camera) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Camera) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Camera) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Camera) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *Camera) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Camera) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *Camera) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *Camera) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Camera) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Camera) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Camera) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Camera) SetName(v string) {
	o.Name = &v
}

// GetOrthographic returns the Orthographic field value if set, zero value otherwise.
func (o *Camera) GetOrthographic() CameraOrthographic {
	if o == nil || o.Orthographic == nil {
		var ret CameraOrthographic
		return ret
	}
	return *o.Orthographic
}

// GetOrthographicOk returns a tuple with the Orthographic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Camera) GetOrthographicOk() (*CameraOrthographic, bool) {
	if o == nil || o.Orthographic == nil {
		return nil, false
	}
	return o.Orthographic, true
}

// HasOrthographic returns a boolean if a field has been set.
func (o *Camera) HasOrthographic() bool {
	if o != nil && o.Orthographic != nil {
		return true
	}

	return false
}

// SetOrthographic gets a reference to the given CameraOrthographic and assigns it to the Orthographic field.
func (o *Camera) SetOrthographic(v CameraOrthographic) {
	o.Orthographic = &v
}

// GetPerspective returns the Perspective field value if set, zero value otherwise.
func (o *Camera) GetPerspective() CameraPerspective {
	if o == nil || o.Perspective == nil {
		var ret CameraPerspective
		return ret
	}
	return *o.Perspective
}

// GetPerspectiveOk returns a tuple with the Perspective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Camera) GetPerspectiveOk() (*CameraPerspective, bool) {
	if o == nil || o.Perspective == nil {
		return nil, false
	}
	return o.Perspective, true
}

// HasPerspective returns a boolean if a field has been set.
func (o *Camera) HasPerspective() bool {
	if o != nil && o.Perspective != nil {
		return true
	}

	return false
}

// SetPerspective gets a reference to the given CameraPerspective and assigns it to the Perspective field.
func (o *Camera) SetPerspective(v CameraPerspective) {
	o.Perspective = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Camera) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Camera) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Camera) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Camera) SetType(v string) {
	o.Type = &v
}

func (o Camera) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Orthographic != nil {
		toSerialize["orthographic"] = o.Orthographic
	}
	if o.Perspective != nil {
		toSerialize["perspective"] = o.Perspective
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCamera struct {
	value *Camera
	isSet bool
}

func (v NullableCamera) Get() *Camera {
	return v.value
}

func (v *NullableCamera) Set(val *Camera) {
	v.value = val
	v.isSet = true
}

func (v NullableCamera) IsSet() bool {
	return v.isSet
}

func (v *NullableCamera) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCamera(val *Camera) *NullableCamera {
	return &NullableCamera{value: val, isSet: true}
}

func (v NullableCamera) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCamera) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
