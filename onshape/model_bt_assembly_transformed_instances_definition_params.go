/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13883-0fccf6f231ad
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyTransformedInstancesDefinitionParams struct for BTAssemblyTransformedInstancesDefinitionParams
type BTAssemblyTransformedInstancesDefinitionParams struct {
	TransformGroups []TransformGroup `json:"transformGroups,omitempty"`
}

// NewBTAssemblyTransformedInstancesDefinitionParams instantiates a new BTAssemblyTransformedInstancesDefinitionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyTransformedInstancesDefinitionParams() *BTAssemblyTransformedInstancesDefinitionParams {
	this := BTAssemblyTransformedInstancesDefinitionParams{}
	return &this
}

// NewBTAssemblyTransformedInstancesDefinitionParamsWithDefaults instantiates a new BTAssemblyTransformedInstancesDefinitionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyTransformedInstancesDefinitionParamsWithDefaults() *BTAssemblyTransformedInstancesDefinitionParams {
	this := BTAssemblyTransformedInstancesDefinitionParams{}
	return &this
}

// GetTransformGroups returns the TransformGroups field value if set, zero value otherwise.
func (o *BTAssemblyTransformedInstancesDefinitionParams) GetTransformGroups() []TransformGroup {
	if o == nil || o.TransformGroups == nil {
		var ret []TransformGroup
		return ret
	}
	return o.TransformGroups
}

// GetTransformGroupsOk returns a tuple with the TransformGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyTransformedInstancesDefinitionParams) GetTransformGroupsOk() ([]TransformGroup, bool) {
	if o == nil || o.TransformGroups == nil {
		return nil, false
	}
	return o.TransformGroups, true
}

// HasTransformGroups returns a boolean if a field has been set.
func (o *BTAssemblyTransformedInstancesDefinitionParams) HasTransformGroups() bool {
	if o != nil && o.TransformGroups != nil {
		return true
	}

	return false
}

// SetTransformGroups gets a reference to the given []TransformGroup and assigns it to the TransformGroups field.
func (o *BTAssemblyTransformedInstancesDefinitionParams) SetTransformGroups(v []TransformGroup) {
	o.TransformGroups = v
}

func (o BTAssemblyTransformedInstancesDefinitionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TransformGroups != nil {
		toSerialize["transformGroups"] = o.TransformGroups
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyTransformedInstancesDefinitionParams struct {
	value *BTAssemblyTransformedInstancesDefinitionParams
	isSet bool
}

func (v NullableBTAssemblyTransformedInstancesDefinitionParams) Get() *BTAssemblyTransformedInstancesDefinitionParams {
	return v.value
}

func (v *NullableBTAssemblyTransformedInstancesDefinitionParams) Set(val *BTAssemblyTransformedInstancesDefinitionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyTransformedInstancesDefinitionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyTransformedInstancesDefinitionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyTransformedInstancesDefinitionParams(val *BTAssemblyTransformedInstancesDefinitionParams) *NullableBTAssemblyTransformedInstancesDefinitionParams {
	return &NullableBTAssemblyTransformedInstancesDefinitionParams{value: val, isSet: true}
}

func (v NullableBTAssemblyTransformedInstancesDefinitionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyTransformedInstancesDefinitionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
