/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.175.29152-5591e93bd1c1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMSketchCompositeEntity893 struct for BTMSketchCompositeEntity893
type BTMSketchCompositeEntity893 struct {
	BtType                              *string            `json:"btType,omitempty"`
	ImportMicroversion                  *string            `json:"importMicroversion,omitempty"`
	NodeId                              *string            `json:"nodeId,omitempty"`
	EntityId                            *string            `json:"entityId,omitempty"`
	EntityIdAndReplaceInDependentFields *string            `json:"entityIdAndReplaceInDependentFields,omitempty"`
	Namespace                           *string            `json:"namespace,omitempty"`
	Parameters                          []BTMParameter1    `json:"parameters,omitempty"`
	SubEntities                         []BTMSketchEntity3 `json:"subEntities,omitempty"`
}

// NewBTMSketchCompositeEntity893 instantiates a new BTMSketchCompositeEntity893 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMSketchCompositeEntity893() *BTMSketchCompositeEntity893 {
	this := BTMSketchCompositeEntity893{}
	return &this
}

// NewBTMSketchCompositeEntity893WithDefaults instantiates a new BTMSketchCompositeEntity893 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMSketchCompositeEntity893WithDefaults() *BTMSketchCompositeEntity893 {
	this := BTMSketchCompositeEntity893{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMSketchCompositeEntity893) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCompositeEntity893) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMSketchCompositeEntity893) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMSketchCompositeEntity893) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMSketchCompositeEntity893) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCompositeEntity893) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMSketchCompositeEntity893) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMSketchCompositeEntity893) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMSketchCompositeEntity893) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCompositeEntity893) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMSketchCompositeEntity893) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMSketchCompositeEntity893) SetNodeId(v string) {
	o.NodeId = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *BTMSketchCompositeEntity893) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCompositeEntity893) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *BTMSketchCompositeEntity893) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *BTMSketchCompositeEntity893) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityIdAndReplaceInDependentFields returns the EntityIdAndReplaceInDependentFields field value if set, zero value otherwise.
func (o *BTMSketchCompositeEntity893) GetEntityIdAndReplaceInDependentFields() string {
	if o == nil || o.EntityIdAndReplaceInDependentFields == nil {
		var ret string
		return ret
	}
	return *o.EntityIdAndReplaceInDependentFields
}

// GetEntityIdAndReplaceInDependentFieldsOk returns a tuple with the EntityIdAndReplaceInDependentFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCompositeEntity893) GetEntityIdAndReplaceInDependentFieldsOk() (*string, bool) {
	if o == nil || o.EntityIdAndReplaceInDependentFields == nil {
		return nil, false
	}
	return o.EntityIdAndReplaceInDependentFields, true
}

// HasEntityIdAndReplaceInDependentFields returns a boolean if a field has been set.
func (o *BTMSketchCompositeEntity893) HasEntityIdAndReplaceInDependentFields() bool {
	if o != nil && o.EntityIdAndReplaceInDependentFields != nil {
		return true
	}

	return false
}

// SetEntityIdAndReplaceInDependentFields gets a reference to the given string and assigns it to the EntityIdAndReplaceInDependentFields field.
func (o *BTMSketchCompositeEntity893) SetEntityIdAndReplaceInDependentFields(v string) {
	o.EntityIdAndReplaceInDependentFields = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMSketchCompositeEntity893) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCompositeEntity893) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMSketchCompositeEntity893) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMSketchCompositeEntity893) SetNamespace(v string) {
	o.Namespace = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTMSketchCompositeEntity893) GetParameters() []BTMParameter1 {
	if o == nil || o.Parameters == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCompositeEntity893) GetParametersOk() ([]BTMParameter1, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTMSketchCompositeEntity893) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *BTMSketchCompositeEntity893) SetParameters(v []BTMParameter1) {
	o.Parameters = v
}

// GetSubEntities returns the SubEntities field value if set, zero value otherwise.
func (o *BTMSketchCompositeEntity893) GetSubEntities() []BTMSketchEntity3 {
	if o == nil || o.SubEntities == nil {
		var ret []BTMSketchEntity3
		return ret
	}
	return o.SubEntities
}

// GetSubEntitiesOk returns a tuple with the SubEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCompositeEntity893) GetSubEntitiesOk() ([]BTMSketchEntity3, bool) {
	if o == nil || o.SubEntities == nil {
		return nil, false
	}
	return o.SubEntities, true
}

// HasSubEntities returns a boolean if a field has been set.
func (o *BTMSketchCompositeEntity893) HasSubEntities() bool {
	if o != nil && o.SubEntities != nil {
		return true
	}

	return false
}

// SetSubEntities gets a reference to the given []BTMSketchEntity3 and assigns it to the SubEntities field.
func (o *BTMSketchCompositeEntity893) SetSubEntities(v []BTMSketchEntity3) {
	o.SubEntities = v
}

func (o BTMSketchCompositeEntity893) MarshalJSON() ([]byte, error) {
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
	if o.EntityId != nil {
		toSerialize["entityId"] = o.EntityId
	}
	if o.EntityIdAndReplaceInDependentFields != nil {
		toSerialize["entityIdAndReplaceInDependentFields"] = o.EntityIdAndReplaceInDependentFields
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.SubEntities != nil {
		toSerialize["subEntities"] = o.SubEntities
	}
	return json.Marshal(toSerialize)
}

type NullableBTMSketchCompositeEntity893 struct {
	value *BTMSketchCompositeEntity893
	isSet bool
}

func (v NullableBTMSketchCompositeEntity893) Get() *BTMSketchCompositeEntity893 {
	return v.value
}

func (v *NullableBTMSketchCompositeEntity893) Set(val *BTMSketchCompositeEntity893) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMSketchCompositeEntity893) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMSketchCompositeEntity893) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMSketchCompositeEntity893(val *BTMSketchCompositeEntity893) *NullableBTMSketchCompositeEntity893 {
	return &NullableBTMSketchCompositeEntity893{value: val, isSet: true}
}

func (v NullableBTMSketchCompositeEntity893) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMSketchCompositeEntity893) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
