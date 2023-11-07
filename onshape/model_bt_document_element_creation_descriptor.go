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

// BTDocumentElementCreationDescriptor struct for BTDocumentElementCreationDescriptor
type BTDocumentElementCreationDescriptor struct {
	ElementParams *BTAppElementParams `json:"elementParams,omitempty"`
	ElementType   *int32              `json:"elementType,omitempty"`
}

// NewBTDocumentElementCreationDescriptor instantiates a new BTDocumentElementCreationDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentElementCreationDescriptor() *BTDocumentElementCreationDescriptor {
	this := BTDocumentElementCreationDescriptor{}
	return &this
}

// NewBTDocumentElementCreationDescriptorWithDefaults instantiates a new BTDocumentElementCreationDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentElementCreationDescriptorWithDefaults() *BTDocumentElementCreationDescriptor {
	this := BTDocumentElementCreationDescriptor{}
	return &this
}

// GetElementParams returns the ElementParams field value if set, zero value otherwise.
func (o *BTDocumentElementCreationDescriptor) GetElementParams() BTAppElementParams {
	if o == nil || o.ElementParams == nil {
		var ret BTAppElementParams
		return ret
	}
	return *o.ElementParams
}

// GetElementParamsOk returns a tuple with the ElementParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementCreationDescriptor) GetElementParamsOk() (*BTAppElementParams, bool) {
	if o == nil || o.ElementParams == nil {
		return nil, false
	}
	return o.ElementParams, true
}

// HasElementParams returns a boolean if a field has been set.
func (o *BTDocumentElementCreationDescriptor) HasElementParams() bool {
	if o != nil && o.ElementParams != nil {
		return true
	}

	return false
}

// SetElementParams gets a reference to the given BTAppElementParams and assigns it to the ElementParams field.
func (o *BTDocumentElementCreationDescriptor) SetElementParams(v BTAppElementParams) {
	o.ElementParams = &v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *BTDocumentElementCreationDescriptor) GetElementType() int32 {
	if o == nil || o.ElementType == nil {
		var ret int32
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementCreationDescriptor) GetElementTypeOk() (*int32, bool) {
	if o == nil || o.ElementType == nil {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *BTDocumentElementCreationDescriptor) HasElementType() bool {
	if o != nil && o.ElementType != nil {
		return true
	}

	return false
}

// SetElementType gets a reference to the given int32 and assigns it to the ElementType field.
func (o *BTDocumentElementCreationDescriptor) SetElementType(v int32) {
	o.ElementType = &v
}

func (o BTDocumentElementCreationDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ElementParams != nil {
		toSerialize["elementParams"] = o.ElementParams
	}
	if o.ElementType != nil {
		toSerialize["elementType"] = o.ElementType
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentElementCreationDescriptor struct {
	value *BTDocumentElementCreationDescriptor
	isSet bool
}

func (v NullableBTDocumentElementCreationDescriptor) Get() *BTDocumentElementCreationDescriptor {
	return v.value
}

func (v *NullableBTDocumentElementCreationDescriptor) Set(val *BTDocumentElementCreationDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentElementCreationDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentElementCreationDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentElementCreationDescriptor(val *BTDocumentElementCreationDescriptor) *NullableBTDocumentElementCreationDescriptor {
	return &NullableBTDocumentElementCreationDescriptor{value: val, isSet: true}
}

func (v NullableBTDocumentElementCreationDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentElementCreationDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
