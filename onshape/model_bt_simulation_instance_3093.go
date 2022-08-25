/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6210-7a182f03d281
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSimulationInstance3093 struct for BTSimulationInstance3093
type BTSimulationInstance3093 struct {
	BtType             *string `json:"btType,omitempty"`
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	NodeId             *string `json:"nodeId,omitempty"`
	IsModal            *bool   `json:"isModal,omitempty"`
	SimulationId       *string `json:"simulationId,omitempty"`
}

// NewBTSimulationInstance3093 instantiates a new BTSimulationInstance3093 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSimulationInstance3093() *BTSimulationInstance3093 {
	this := BTSimulationInstance3093{}
	return &this
}

// NewBTSimulationInstance3093WithDefaults instantiates a new BTSimulationInstance3093 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSimulationInstance3093WithDefaults() *BTSimulationInstance3093 {
	this := BTSimulationInstance3093{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSimulationInstance3093) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationInstance3093) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSimulationInstance3093) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSimulationInstance3093) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTSimulationInstance3093) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationInstance3093) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTSimulationInstance3093) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTSimulationInstance3093) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTSimulationInstance3093) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationInstance3093) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTSimulationInstance3093) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTSimulationInstance3093) SetNodeId(v string) {
	o.NodeId = &v
}

// GetIsModal returns the IsModal field value if set, zero value otherwise.
func (o *BTSimulationInstance3093) GetIsModal() bool {
	if o == nil || o.IsModal == nil {
		var ret bool
		return ret
	}
	return *o.IsModal
}

// GetIsModalOk returns a tuple with the IsModal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationInstance3093) GetIsModalOk() (*bool, bool) {
	if o == nil || o.IsModal == nil {
		return nil, false
	}
	return o.IsModal, true
}

// HasIsModal returns a boolean if a field has been set.
func (o *BTSimulationInstance3093) HasIsModal() bool {
	if o != nil && o.IsModal != nil {
		return true
	}

	return false
}

// SetIsModal gets a reference to the given bool and assigns it to the IsModal field.
func (o *BTSimulationInstance3093) SetIsModal(v bool) {
	o.IsModal = &v
}

// GetSimulationId returns the SimulationId field value if set, zero value otherwise.
func (o *BTSimulationInstance3093) GetSimulationId() string {
	if o == nil || o.SimulationId == nil {
		var ret string
		return ret
	}
	return *o.SimulationId
}

// GetSimulationIdOk returns a tuple with the SimulationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationInstance3093) GetSimulationIdOk() (*string, bool) {
	if o == nil || o.SimulationId == nil {
		return nil, false
	}
	return o.SimulationId, true
}

// HasSimulationId returns a boolean if a field has been set.
func (o *BTSimulationInstance3093) HasSimulationId() bool {
	if o != nil && o.SimulationId != nil {
		return true
	}

	return false
}

// SetSimulationId gets a reference to the given string and assigns it to the SimulationId field.
func (o *BTSimulationInstance3093) SetSimulationId(v string) {
	o.SimulationId = &v
}

func (o BTSimulationInstance3093) MarshalJSON() ([]byte, error) {
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
	if o.IsModal != nil {
		toSerialize["isModal"] = o.IsModal
	}
	if o.SimulationId != nil {
		toSerialize["simulationId"] = o.SimulationId
	}
	return json.Marshal(toSerialize)
}

type NullableBTSimulationInstance3093 struct {
	value *BTSimulationInstance3093
	isSet bool
}

func (v NullableBTSimulationInstance3093) Get() *BTSimulationInstance3093 {
	return v.value
}

func (v *NullableBTSimulationInstance3093) Set(val *BTSimulationInstance3093) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSimulationInstance3093) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSimulationInstance3093) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSimulationInstance3093(val *BTSimulationInstance3093) *NullableBTSimulationInstance3093 {
	return &NullableBTSimulationInstance3093{value: val, isSet: true}
}

func (v NullableBTSimulationInstance3093) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSimulationInstance3093) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
