/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.166.18819-fa27aca4c021
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTVariableInfo Variables in the VariableTable
type BTVariableInfo struct {
	// Variable description
	Description *string `json:"description,omitempty"`
	// Variable expression
	Expression string `json:"expression"`
	// Variable name
	Name string          `json:"name"`
	Type GBTVariableType `json:"type"`
	// Variable formatted value
	Value string `json:"value"`
}

// NewBTVariableInfo instantiates a new BTVariableInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTVariableInfo(expression string, name string, type_ GBTVariableType, value string) *BTVariableInfo {
	this := BTVariableInfo{}
	this.Expression = expression
	this.Name = name
	this.Type = type_
	this.Value = value
	return &this
}

// NewBTVariableInfoWithDefaults instantiates a new BTVariableInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTVariableInfoWithDefaults() *BTVariableInfo {
	this := BTVariableInfo{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTVariableInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTVariableInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTVariableInfo) SetDescription(v string) {
	o.Description = &v
}

// GetExpression returns the Expression field value
func (o *BTVariableInfo) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *BTVariableInfo) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *BTVariableInfo) SetExpression(v string) {
	o.Expression = v
}

// GetName returns the Name field value
func (o *BTVariableInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BTVariableInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BTVariableInfo) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *BTVariableInfo) GetType() GBTVariableType {
	if o == nil {
		var ret GBTVariableType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BTVariableInfo) GetTypeOk() (*GBTVariableType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BTVariableInfo) SetType(v GBTVariableType) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *BTVariableInfo) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *BTVariableInfo) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *BTVariableInfo) SetValue(v string) {
	o.Value = v
}

func (o BTVariableInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["expression"] = o.Expression
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTVariableInfo struct {
	value *BTVariableInfo
	isSet bool
}

func (v NullableBTVariableInfo) Get() *BTVariableInfo {
	return v.value
}

func (v *NullableBTVariableInfo) Set(val *BTVariableInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTVariableInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTVariableInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTVariableInfo(val *BTVariableInfo) *NullableBTVariableInfo {
	return &NullableBTVariableInfo{value: val, isSet: true}
}

func (v NullableBTVariableInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTVariableInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
