/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.163.15457-d8ebaa9b9e42
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTVariableTableInfo struct for BTVariableTableInfo
type BTVariableTableInfo struct {
	VariableStudioReference *BTVariableStudioReferenceInfo `json:"variableStudioReference,omitempty"`
	// Variables in the VariableTable
	Variables []BTVariableInfo `json:"variables"`
}

// NewBTVariableTableInfo instantiates a new BTVariableTableInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTVariableTableInfo(variables []BTVariableInfo) *BTVariableTableInfo {
	this := BTVariableTableInfo{}
	this.Variables = variables
	return &this
}

// NewBTVariableTableInfoWithDefaults instantiates a new BTVariableTableInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTVariableTableInfoWithDefaults() *BTVariableTableInfo {
	this := BTVariableTableInfo{}
	return &this
}

// GetVariableStudioReference returns the VariableStudioReference field value if set, zero value otherwise.
func (o *BTVariableTableInfo) GetVariableStudioReference() BTVariableStudioReferenceInfo {
	if o == nil || o.VariableStudioReference == nil {
		var ret BTVariableStudioReferenceInfo
		return ret
	}
	return *o.VariableStudioReference
}

// GetVariableStudioReferenceOk returns a tuple with the VariableStudioReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableTableInfo) GetVariableStudioReferenceOk() (*BTVariableStudioReferenceInfo, bool) {
	if o == nil || o.VariableStudioReference == nil {
		return nil, false
	}
	return o.VariableStudioReference, true
}

// HasVariableStudioReference returns a boolean if a field has been set.
func (o *BTVariableTableInfo) HasVariableStudioReference() bool {
	if o != nil && o.VariableStudioReference != nil {
		return true
	}

	return false
}

// SetVariableStudioReference gets a reference to the given BTVariableStudioReferenceInfo and assigns it to the VariableStudioReference field.
func (o *BTVariableTableInfo) SetVariableStudioReference(v BTVariableStudioReferenceInfo) {
	o.VariableStudioReference = &v
}

// GetVariables returns the Variables field value
func (o *BTVariableTableInfo) GetVariables() []BTVariableInfo {
	if o == nil {
		var ret []BTVariableInfo
		return ret
	}

	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *BTVariableTableInfo) GetVariablesOk() ([]BTVariableInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variables, true
}

// SetVariables sets field value
func (o *BTVariableTableInfo) SetVariables(v []BTVariableInfo) {
	o.Variables = v
}

func (o BTVariableTableInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VariableStudioReference != nil {
		toSerialize["variableStudioReference"] = o.VariableStudioReference
	}
	if true {
		toSerialize["variables"] = o.Variables
	}
	return json.Marshal(toSerialize)
}

type NullableBTVariableTableInfo struct {
	value *BTVariableTableInfo
	isSet bool
}

func (v NullableBTVariableTableInfo) Get() *BTVariableTableInfo {
	return v.value
}

func (v *NullableBTVariableTableInfo) Set(val *BTVariableTableInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTVariableTableInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTVariableTableInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTVariableTableInfo(val *BTVariableTableInfo) *NullableBTVariableTableInfo {
	return &NullableBTVariableTableInfo{value: val, isSet: true}
}

func (v NullableBTVariableTableInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTVariableTableInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
