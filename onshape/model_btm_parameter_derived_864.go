/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23196-ae5f57bec467
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMParameterDerived864 struct for BTMParameterDerived864
type BTMParameterDerived864 struct {
	BtType             *string         `json:"btType,omitempty"`
	ImportMicroversion *string         `json:"importMicroversion,omitempty"`
	NodeId             *string         `json:"nodeId,omitempty"`
	ParameterId        *string         `json:"parameterId,omitempty"`
	Imports            []BTMImport136  `json:"imports,omitempty"`
	ModuleId           *BTPModuleId235 `json:"moduleId,omitempty"`
	Namespace          *string         `json:"namespace,omitempty"`
}

// NewBTMParameterDerived864 instantiates a new BTMParameterDerived864 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterDerived864() *BTMParameterDerived864 {
	this := BTMParameterDerived864{}
	return &this
}

// NewBTMParameterDerived864WithDefaults instantiates a new BTMParameterDerived864 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterDerived864WithDefaults() *BTMParameterDerived864 {
	this := BTMParameterDerived864{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterDerived864) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterDerived864) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterDerived864) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterDerived864) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterDerived864) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterDerived864) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterDerived864) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterDerived864) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterDerived864) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterDerived864) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameterDerived864) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameterDerived864) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterDerived864) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterDerived864) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterDerived864) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterDerived864) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetImports returns the Imports field value if set, zero value otherwise.
func (o *BTMParameterDerived864) GetImports() []BTMImport136 {
	if o == nil || o.Imports == nil {
		var ret []BTMImport136
		return ret
	}
	return o.Imports
}

// GetImportsOk returns a tuple with the Imports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterDerived864) GetImportsOk() ([]BTMImport136, bool) {
	if o == nil || o.Imports == nil {
		return nil, false
	}
	return o.Imports, true
}

// HasImports returns a boolean if a field has been set.
func (o *BTMParameterDerived864) HasImports() bool {
	if o != nil && o.Imports != nil {
		return true
	}

	return false
}

// SetImports gets a reference to the given []BTMImport136 and assigns it to the Imports field.
func (o *BTMParameterDerived864) SetImports(v []BTMImport136) {
	o.Imports = v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *BTMParameterDerived864) GetModuleId() BTPModuleId235 {
	if o == nil || o.ModuleId == nil {
		var ret BTPModuleId235
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterDerived864) GetModuleIdOk() (*BTPModuleId235, bool) {
	if o == nil || o.ModuleId == nil {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *BTMParameterDerived864) HasModuleId() bool {
	if o != nil && o.ModuleId != nil {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given BTPModuleId235 and assigns it to the ModuleId field.
func (o *BTMParameterDerived864) SetModuleId(v BTPModuleId235) {
	o.ModuleId = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMParameterDerived864) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterDerived864) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMParameterDerived864) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMParameterDerived864) SetNamespace(v string) {
	o.Namespace = &v
}

func (o BTMParameterDerived864) MarshalJSON() ([]byte, error) {
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
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.Imports != nil {
		toSerialize["imports"] = o.Imports
	}
	if o.ModuleId != nil {
		toSerialize["moduleId"] = o.ModuleId
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterDerived864 struct {
	value *BTMParameterDerived864
	isSet bool
}

func (v NullableBTMParameterDerived864) Get() *BTMParameterDerived864 {
	return v.value
}

func (v *NullableBTMParameterDerived864) Set(val *BTMParameterDerived864) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterDerived864) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterDerived864) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterDerived864(val *BTMParameterDerived864) *NullableBTMParameterDerived864 {
	return &NullableBTMParameterDerived864{value: val, isSet: true}
}

func (v NullableBTMParameterDerived864) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterDerived864) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
