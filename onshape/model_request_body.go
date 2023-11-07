/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25478-d4e5ab4765a4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// RequestBody struct for RequestBody
type RequestBody struct {
	Content     *RequestBodyContent               `json:"content,omitempty"`
	Description *string                           `json:"description,omitempty"`
	Extensions  map[string]map[string]interface{} `json:"extensions,omitempty"`
	Getref      *string                           `json:"get$ref,omitempty"`
	Required    *bool                             `json:"required,omitempty"`
}

// NewRequestBody instantiates a new RequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestBody() *RequestBody {
	this := RequestBody{}
	return &this
}

// NewRequestBodyWithDefaults instantiates a new RequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestBodyWithDefaults() *RequestBody {
	this := RequestBody{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *RequestBody) GetContent() RequestBodyContent {
	if o == nil || o.Content == nil {
		var ret RequestBodyContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestBody) GetContentOk() (*RequestBodyContent, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *RequestBody) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given RequestBodyContent and assigns it to the Content field.
func (o *RequestBody) SetContent(v RequestBodyContent) {
	o.Content = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestBody) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestBody) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestBody) SetDescription(v string) {
	o.Description = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *RequestBody) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestBody) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *RequestBody) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *RequestBody) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetGetref returns the Getref field value if set, zero value otherwise.
func (o *RequestBody) GetGetref() string {
	if o == nil || o.Getref == nil {
		var ret string
		return ret
	}
	return *o.Getref
}

// GetGetrefOk returns a tuple with the Getref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestBody) GetGetrefOk() (*string, bool) {
	if o == nil || o.Getref == nil {
		return nil, false
	}
	return o.Getref, true
}

// HasGetref returns a boolean if a field has been set.
func (o *RequestBody) HasGetref() bool {
	if o != nil && o.Getref != nil {
		return true
	}

	return false
}

// SetGetref gets a reference to the given string and assigns it to the Getref field.
func (o *RequestBody) SetGetref(v string) {
	o.Getref = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *RequestBody) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestBody) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *RequestBody) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *RequestBody) SetRequired(v bool) {
	o.Required = &v
}

func (o RequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Getref != nil {
		toSerialize["get$ref"] = o.Getref
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	return json.Marshal(toSerialize)
}

type NullableRequestBody struct {
	value *RequestBody
	isSet bool
}

func (v NullableRequestBody) Get() *RequestBody {
	return v.value
}

func (v *NullableRequestBody) Set(val *RequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestBody(val *RequestBody) *NullableRequestBody {
	return &NullableRequestBody{value: val, isSet: true}
}

func (v NullableRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
