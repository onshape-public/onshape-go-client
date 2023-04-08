/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.14179-e31166ce6183
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMSuppressionStateConfigured2598 struct for BTMSuppressionStateConfigured2598
type BTMSuppressionStateConfigured2598 struct {
	BtType                             *string                  `json:"btType,omitempty"`
	ImportMicroversion                 *string                  `json:"importMicroversion,omitempty"`
	NodeId                             *string                  `json:"nodeId,omitempty"`
	ConfigurationParameterId           *string                  `json:"configurationParameterId,omitempty"`
	ConfigurationParameterIdFieldIndex *int32                   `json:"configurationParameterIdFieldIndex,omitempty"`
	Values                             []BTMConfiguredValue1341 `json:"values,omitempty"`
	ValuesFieldIndex                   *int32                   `json:"valuesFieldIndex,omitempty"`
}

// NewBTMSuppressionStateConfigured2598 instantiates a new BTMSuppressionStateConfigured2598 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMSuppressionStateConfigured2598() *BTMSuppressionStateConfigured2598 {
	this := BTMSuppressionStateConfigured2598{}
	return &this
}

// NewBTMSuppressionStateConfigured2598WithDefaults instantiates a new BTMSuppressionStateConfigured2598 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMSuppressionStateConfigured2598WithDefaults() *BTMSuppressionStateConfigured2598 {
	this := BTMSuppressionStateConfigured2598{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMSuppressionStateConfigured2598) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSuppressionStateConfigured2598) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMSuppressionStateConfigured2598) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMSuppressionStateConfigured2598) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMSuppressionStateConfigured2598) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSuppressionStateConfigured2598) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMSuppressionStateConfigured2598) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMSuppressionStateConfigured2598) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMSuppressionStateConfigured2598) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSuppressionStateConfigured2598) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMSuppressionStateConfigured2598) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMSuppressionStateConfigured2598) SetNodeId(v string) {
	o.NodeId = &v
}

// GetConfigurationParameterId returns the ConfigurationParameterId field value if set, zero value otherwise.
func (o *BTMSuppressionStateConfigured2598) GetConfigurationParameterId() string {
	if o == nil || o.ConfigurationParameterId == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationParameterId
}

// GetConfigurationParameterIdOk returns a tuple with the ConfigurationParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSuppressionStateConfigured2598) GetConfigurationParameterIdOk() (*string, bool) {
	if o == nil || o.ConfigurationParameterId == nil {
		return nil, false
	}
	return o.ConfigurationParameterId, true
}

// HasConfigurationParameterId returns a boolean if a field has been set.
func (o *BTMSuppressionStateConfigured2598) HasConfigurationParameterId() bool {
	if o != nil && o.ConfigurationParameterId != nil {
		return true
	}

	return false
}

// SetConfigurationParameterId gets a reference to the given string and assigns it to the ConfigurationParameterId field.
func (o *BTMSuppressionStateConfigured2598) SetConfigurationParameterId(v string) {
	o.ConfigurationParameterId = &v
}

// GetConfigurationParameterIdFieldIndex returns the ConfigurationParameterIdFieldIndex field value if set, zero value otherwise.
func (o *BTMSuppressionStateConfigured2598) GetConfigurationParameterIdFieldIndex() int32 {
	if o == nil || o.ConfigurationParameterIdFieldIndex == nil {
		var ret int32
		return ret
	}
	return *o.ConfigurationParameterIdFieldIndex
}

// GetConfigurationParameterIdFieldIndexOk returns a tuple with the ConfigurationParameterIdFieldIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSuppressionStateConfigured2598) GetConfigurationParameterIdFieldIndexOk() (*int32, bool) {
	if o == nil || o.ConfigurationParameterIdFieldIndex == nil {
		return nil, false
	}
	return o.ConfigurationParameterIdFieldIndex, true
}

// HasConfigurationParameterIdFieldIndex returns a boolean if a field has been set.
func (o *BTMSuppressionStateConfigured2598) HasConfigurationParameterIdFieldIndex() bool {
	if o != nil && o.ConfigurationParameterIdFieldIndex != nil {
		return true
	}

	return false
}

// SetConfigurationParameterIdFieldIndex gets a reference to the given int32 and assigns it to the ConfigurationParameterIdFieldIndex field.
func (o *BTMSuppressionStateConfigured2598) SetConfigurationParameterIdFieldIndex(v int32) {
	o.ConfigurationParameterIdFieldIndex = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *BTMSuppressionStateConfigured2598) GetValues() []BTMConfiguredValue1341 {
	if o == nil || o.Values == nil {
		var ret []BTMConfiguredValue1341
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSuppressionStateConfigured2598) GetValuesOk() ([]BTMConfiguredValue1341, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *BTMSuppressionStateConfigured2598) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []BTMConfiguredValue1341 and assigns it to the Values field.
func (o *BTMSuppressionStateConfigured2598) SetValues(v []BTMConfiguredValue1341) {
	o.Values = v
}

// GetValuesFieldIndex returns the ValuesFieldIndex field value if set, zero value otherwise.
func (o *BTMSuppressionStateConfigured2598) GetValuesFieldIndex() int32 {
	if o == nil || o.ValuesFieldIndex == nil {
		var ret int32
		return ret
	}
	return *o.ValuesFieldIndex
}

// GetValuesFieldIndexOk returns a tuple with the ValuesFieldIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSuppressionStateConfigured2598) GetValuesFieldIndexOk() (*int32, bool) {
	if o == nil || o.ValuesFieldIndex == nil {
		return nil, false
	}
	return o.ValuesFieldIndex, true
}

// HasValuesFieldIndex returns a boolean if a field has been set.
func (o *BTMSuppressionStateConfigured2598) HasValuesFieldIndex() bool {
	if o != nil && o.ValuesFieldIndex != nil {
		return true
	}

	return false
}

// SetValuesFieldIndex gets a reference to the given int32 and assigns it to the ValuesFieldIndex field.
func (o *BTMSuppressionStateConfigured2598) SetValuesFieldIndex(v int32) {
	o.ValuesFieldIndex = &v
}

func (o BTMSuppressionStateConfigured2598) MarshalJSON() ([]byte, error) {
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
	if o.ConfigurationParameterId != nil {
		toSerialize["configurationParameterId"] = o.ConfigurationParameterId
	}
	if o.ConfigurationParameterIdFieldIndex != nil {
		toSerialize["configurationParameterIdFieldIndex"] = o.ConfigurationParameterIdFieldIndex
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	if o.ValuesFieldIndex != nil {
		toSerialize["valuesFieldIndex"] = o.ValuesFieldIndex
	}
	return json.Marshal(toSerialize)
}

type NullableBTMSuppressionStateConfigured2598 struct {
	value *BTMSuppressionStateConfigured2598
	isSet bool
}

func (v NullableBTMSuppressionStateConfigured2598) Get() *BTMSuppressionStateConfigured2598 {
	return v.value
}

func (v *NullableBTMSuppressionStateConfigured2598) Set(val *BTMSuppressionStateConfigured2598) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMSuppressionStateConfigured2598) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMSuppressionStateConfigured2598) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMSuppressionStateConfigured2598(val *BTMSuppressionStateConfigured2598) *NullableBTMSuppressionStateConfigured2598 {
	return &NullableBTMSuppressionStateConfigured2598{value: val, isSet: true}
}

func (v NullableBTMSuppressionStateConfigured2598) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMSuppressionStateConfigured2598) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
