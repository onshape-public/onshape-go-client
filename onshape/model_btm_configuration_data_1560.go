/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.17497-411bb6b98e6b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMConfigurationData1560 struct for BTMConfigurationData1560
type BTMConfigurationData1560 struct {
	BtType                     *string                   `json:"btType,omitempty"`
	ImportMicroversion         *string                   `json:"importMicroversion,omitempty"`
	NodeId                     *string                   `json:"nodeId,omitempty"`
	CurrentConfiguration       []BTMParameter1           `json:"currentConfiguration,omitempty"`
	CurrentFSValues            *map[string]BTFSValue1888 `json:"currentFSValues,omitempty"`
	DefaultConfigurationValues *map[string]BTFSValue1888 `json:"defaultConfigurationValues,omitempty"`
}

// NewBTMConfigurationData1560 instantiates a new BTMConfigurationData1560 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMConfigurationData1560() *BTMConfigurationData1560 {
	this := BTMConfigurationData1560{}
	return &this
}

// NewBTMConfigurationData1560WithDefaults instantiates a new BTMConfigurationData1560 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMConfigurationData1560WithDefaults() *BTMConfigurationData1560 {
	this := BTMConfigurationData1560{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMConfigurationData1560) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationData1560) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMConfigurationData1560) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMConfigurationData1560) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMConfigurationData1560) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationData1560) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMConfigurationData1560) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMConfigurationData1560) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMConfigurationData1560) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationData1560) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMConfigurationData1560) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMConfigurationData1560) SetNodeId(v string) {
	o.NodeId = &v
}

// GetCurrentConfiguration returns the CurrentConfiguration field value if set, zero value otherwise.
func (o *BTMConfigurationData1560) GetCurrentConfiguration() []BTMParameter1 {
	if o == nil || o.CurrentConfiguration == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.CurrentConfiguration
}

// GetCurrentConfigurationOk returns a tuple with the CurrentConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationData1560) GetCurrentConfigurationOk() ([]BTMParameter1, bool) {
	if o == nil || o.CurrentConfiguration == nil {
		return nil, false
	}
	return o.CurrentConfiguration, true
}

// HasCurrentConfiguration returns a boolean if a field has been set.
func (o *BTMConfigurationData1560) HasCurrentConfiguration() bool {
	if o != nil && o.CurrentConfiguration != nil {
		return true
	}

	return false
}

// SetCurrentConfiguration gets a reference to the given []BTMParameter1 and assigns it to the CurrentConfiguration field.
func (o *BTMConfigurationData1560) SetCurrentConfiguration(v []BTMParameter1) {
	o.CurrentConfiguration = v
}

// GetCurrentFSValues returns the CurrentFSValues field value if set, zero value otherwise.
func (o *BTMConfigurationData1560) GetCurrentFSValues() map[string]BTFSValue1888 {
	if o == nil || o.CurrentFSValues == nil {
		var ret map[string]BTFSValue1888
		return ret
	}
	return *o.CurrentFSValues
}

// GetCurrentFSValuesOk returns a tuple with the CurrentFSValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationData1560) GetCurrentFSValuesOk() (*map[string]BTFSValue1888, bool) {
	if o == nil || o.CurrentFSValues == nil {
		return nil, false
	}
	return o.CurrentFSValues, true
}

// HasCurrentFSValues returns a boolean if a field has been set.
func (o *BTMConfigurationData1560) HasCurrentFSValues() bool {
	if o != nil && o.CurrentFSValues != nil {
		return true
	}

	return false
}

// SetCurrentFSValues gets a reference to the given map[string]BTFSValue1888 and assigns it to the CurrentFSValues field.
func (o *BTMConfigurationData1560) SetCurrentFSValues(v map[string]BTFSValue1888) {
	o.CurrentFSValues = &v
}

// GetDefaultConfigurationValues returns the DefaultConfigurationValues field value if set, zero value otherwise.
func (o *BTMConfigurationData1560) GetDefaultConfigurationValues() map[string]BTFSValue1888 {
	if o == nil || o.DefaultConfigurationValues == nil {
		var ret map[string]BTFSValue1888
		return ret
	}
	return *o.DefaultConfigurationValues
}

// GetDefaultConfigurationValuesOk returns a tuple with the DefaultConfigurationValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationData1560) GetDefaultConfigurationValuesOk() (*map[string]BTFSValue1888, bool) {
	if o == nil || o.DefaultConfigurationValues == nil {
		return nil, false
	}
	return o.DefaultConfigurationValues, true
}

// HasDefaultConfigurationValues returns a boolean if a field has been set.
func (o *BTMConfigurationData1560) HasDefaultConfigurationValues() bool {
	if o != nil && o.DefaultConfigurationValues != nil {
		return true
	}

	return false
}

// SetDefaultConfigurationValues gets a reference to the given map[string]BTFSValue1888 and assigns it to the DefaultConfigurationValues field.
func (o *BTMConfigurationData1560) SetDefaultConfigurationValues(v map[string]BTFSValue1888) {
	o.DefaultConfigurationValues = &v
}

func (o BTMConfigurationData1560) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.CurrentConfiguration != nil {
		toSerialize["currentConfiguration"] = o.CurrentConfiguration
	}
	if o.CurrentFSValues != nil {
		toSerialize["currentFSValues"] = o.CurrentFSValues
	}
	if o.DefaultConfigurationValues != nil {
		toSerialize["defaultConfigurationValues"] = o.DefaultConfigurationValues
	}
	return json.Marshal(toSerialize)
}

type NullableBTMConfigurationData1560 struct {
	value *BTMConfigurationData1560
	isSet bool
}

func (v NullableBTMConfigurationData1560) Get() *BTMConfigurationData1560 {
	return v.value
}

func (v *NullableBTMConfigurationData1560) Set(val *BTMConfigurationData1560) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMConfigurationData1560) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMConfigurationData1560) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMConfigurationData1560(val *BTMConfigurationData1560) *NullableBTMConfigurationData1560 {
	return &NullableBTMConfigurationData1560{value: val, isSet: true}
}

func (v NullableBTMConfigurationData1560) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMConfigurationData1560) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
