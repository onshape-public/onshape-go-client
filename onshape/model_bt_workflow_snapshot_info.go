/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.159.11153-01265f451c50
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTWorkflowSnapshotInfo struct for BTWorkflowSnapshotInfo
type BTWorkflowSnapshotInfo struct {
	Actions             []BTWorkflowActionInfo `json:"actions,omitempty"`
	ApproverIds         []string               `json:"approverIds,omitempty"`
	CanBeDiscarded      *bool                  `json:"canBeDiscarded,omitempty"`
	DebugMicroversionId *string                `json:"debugMicroversionId,omitempty"`
	ErrorMessage        *string                `json:"errorMessage,omitempty"`
	IsCreator           *bool                  `json:"isCreator,omitempty"`
	IsDiscarded         *bool                  `json:"isDiscarded,omitempty"`
	IsFrozen            *bool                  `json:"isFrozen,omitempty"`
	IsSetup             *bool                  `json:"isSetup,omitempty"`
	MetadataState       *string                `json:"metadataState,omitempty"`
	NotifierIds         []string               `json:"notifierIds,omitempty"`
	State               *BTWorkflowStateInfo   `json:"state,omitempty"`
}

// NewBTWorkflowSnapshotInfo instantiates a new BTWorkflowSnapshotInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWorkflowSnapshotInfo() *BTWorkflowSnapshotInfo {
	this := BTWorkflowSnapshotInfo{}
	return &this
}

// NewBTWorkflowSnapshotInfoWithDefaults instantiates a new BTWorkflowSnapshotInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWorkflowSnapshotInfoWithDefaults() *BTWorkflowSnapshotInfo {
	this := BTWorkflowSnapshotInfo{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *BTWorkflowSnapshotInfo) GetActions() []BTWorkflowActionInfo {
	if o == nil || o.Actions == nil {
		var ret []BTWorkflowActionInfo
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowSnapshotInfo) GetActionsOk() ([]BTWorkflowActionInfo, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *BTWorkflowSnapshotInfo) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []BTWorkflowActionInfo and assigns it to the Actions field.
func (o *BTWorkflowSnapshotInfo) SetActions(v []BTWorkflowActionInfo) {
	o.Actions = v
}

// GetApproverIds returns the ApproverIds field value if set, zero value otherwise.
func (o *BTWorkflowSnapshotInfo) GetApproverIds() []string {
	if o == nil || o.ApproverIds == nil {
		var ret []string
		return ret
	}
	return o.ApproverIds
}

// GetApproverIdsOk returns a tuple with the ApproverIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowSnapshotInfo) GetApproverIdsOk() ([]string, bool) {
	if o == nil || o.ApproverIds == nil {
		return nil, false
	}
	return o.ApproverIds, true
}

// HasApproverIds returns a boolean if a field has been set.
func (o *BTWorkflowSnapshotInfo) HasApproverIds() bool {
	if o != nil && o.ApproverIds != nil {
		return true
	}

	return false
}

// SetApproverIds gets a reference to the given []string and assigns it to the ApproverIds field.
func (o *BTWorkflowSnapshotInfo) SetApproverIds(v []string) {
	o.ApproverIds = v
}

// GetCanBeDiscarded returns the CanBeDiscarded field value if set, zero value otherwise.
func (o *BTWorkflowSnapshotInfo) GetCanBeDiscarded() bool {
	if o == nil || o.CanBeDiscarded == nil {
		var ret bool
		return ret
	}
	return *o.CanBeDiscarded
}

// GetCanBeDiscardedOk returns a tuple with the CanBeDiscarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowSnapshotInfo) GetCanBeDiscardedOk() (*bool, bool) {
	if o == nil || o.CanBeDiscarded == nil {
		return nil, false
	}
	return o.CanBeDiscarded, true
}

// HasCanBeDiscarded returns a boolean if a field has been set.
func (o *BTWorkflowSnapshotInfo) HasCanBeDiscarded() bool {
	if o != nil && o.CanBeDiscarded != nil {
		return true
	}

	return false
}

// SetCanBeDiscarded gets a reference to the given bool and assigns it to the CanBeDiscarded field.
func (o *BTWorkflowSnapshotInfo) SetCanBeDiscarded(v bool) {
	o.CanBeDiscarded = &v
}

// GetDebugMicroversionId returns the DebugMicroversionId field value if set, zero value otherwise.
func (o *BTWorkflowSnapshotInfo) GetDebugMicroversionId() string {
	if o == nil || o.DebugMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.DebugMicroversionId
}

// GetDebugMicroversionIdOk returns a tuple with the DebugMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowSnapshotInfo) GetDebugMicroversionIdOk() (*string, bool) {
	if o == nil || o.DebugMicroversionId == nil {
		return nil, false
	}
	return o.DebugMicroversionId, true
}

// HasDebugMicroversionId returns a boolean if a field has been set.
func (o *BTWorkflowSnapshotInfo) HasDebugMicroversionId() bool {
	if o != nil && o.DebugMicroversionId != nil {
		return true
	}

	return false
}

// SetDebugMicroversionId gets a reference to the given string and assigns it to the DebugMicroversionId field.
func (o *BTWorkflowSnapshotInfo) SetDebugMicroversionId(v string) {
	o.DebugMicroversionId = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *BTWorkflowSnapshotInfo) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowSnapshotInfo) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *BTWorkflowSnapshotInfo) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *BTWorkflowSnapshotInfo) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetIsCreator returns the IsCreator field value if set, zero value otherwise.
func (o *BTWorkflowSnapshotInfo) GetIsCreator() bool {
	if o == nil || o.IsCreator == nil {
		var ret bool
		return ret
	}
	return *o.IsCreator
}

// GetIsCreatorOk returns a tuple with the IsCreator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowSnapshotInfo) GetIsCreatorOk() (*bool, bool) {
	if o == nil || o.IsCreator == nil {
		return nil, false
	}
	return o.IsCreator, true
}

// HasIsCreator returns a boolean if a field has been set.
func (o *BTWorkflowSnapshotInfo) HasIsCreator() bool {
	if o != nil && o.IsCreator != nil {
		return true
	}

	return false
}

// SetIsCreator gets a reference to the given bool and assigns it to the IsCreator field.
func (o *BTWorkflowSnapshotInfo) SetIsCreator(v bool) {
	o.IsCreator = &v
}

// GetIsDiscarded returns the IsDiscarded field value if set, zero value otherwise.
func (o *BTWorkflowSnapshotInfo) GetIsDiscarded() bool {
	if o == nil || o.IsDiscarded == nil {
		var ret bool
		return ret
	}
	return *o.IsDiscarded
}

// GetIsDiscardedOk returns a tuple with the IsDiscarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowSnapshotInfo) GetIsDiscardedOk() (*bool, bool) {
	if o == nil || o.IsDiscarded == nil {
		return nil, false
	}
	return o.IsDiscarded, true
}

// HasIsDiscarded returns a boolean if a field has been set.
func (o *BTWorkflowSnapshotInfo) HasIsDiscarded() bool {
	if o != nil && o.IsDiscarded != nil {
		return true
	}

	return false
}

// SetIsDiscarded gets a reference to the given bool and assigns it to the IsDiscarded field.
func (o *BTWorkflowSnapshotInfo) SetIsDiscarded(v bool) {
	o.IsDiscarded = &v
}

// GetIsFrozen returns the IsFrozen field value if set, zero value otherwise.
func (o *BTWorkflowSnapshotInfo) GetIsFrozen() bool {
	if o == nil || o.IsFrozen == nil {
		var ret bool
		return ret
	}
	return *o.IsFrozen
}

// GetIsFrozenOk returns a tuple with the IsFrozen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowSnapshotInfo) GetIsFrozenOk() (*bool, bool) {
	if o == nil || o.IsFrozen == nil {
		return nil, false
	}
	return o.IsFrozen, true
}

// HasIsFrozen returns a boolean if a field has been set.
func (o *BTWorkflowSnapshotInfo) HasIsFrozen() bool {
	if o != nil && o.IsFrozen != nil {
		return true
	}

	return false
}

// SetIsFrozen gets a reference to the given bool and assigns it to the IsFrozen field.
func (o *BTWorkflowSnapshotInfo) SetIsFrozen(v bool) {
	o.IsFrozen = &v
}

// GetIsSetup returns the IsSetup field value if set, zero value otherwise.
func (o *BTWorkflowSnapshotInfo) GetIsSetup() bool {
	if o == nil || o.IsSetup == nil {
		var ret bool
		return ret
	}
	return *o.IsSetup
}

// GetIsSetupOk returns a tuple with the IsSetup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowSnapshotInfo) GetIsSetupOk() (*bool, bool) {
	if o == nil || o.IsSetup == nil {
		return nil, false
	}
	return o.IsSetup, true
}

// HasIsSetup returns a boolean if a field has been set.
func (o *BTWorkflowSnapshotInfo) HasIsSetup() bool {
	if o != nil && o.IsSetup != nil {
		return true
	}

	return false
}

// SetIsSetup gets a reference to the given bool and assigns it to the IsSetup field.
func (o *BTWorkflowSnapshotInfo) SetIsSetup(v bool) {
	o.IsSetup = &v
}

// GetMetadataState returns the MetadataState field value if set, zero value otherwise.
func (o *BTWorkflowSnapshotInfo) GetMetadataState() string {
	if o == nil || o.MetadataState == nil {
		var ret string
		return ret
	}
	return *o.MetadataState
}

// GetMetadataStateOk returns a tuple with the MetadataState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowSnapshotInfo) GetMetadataStateOk() (*string, bool) {
	if o == nil || o.MetadataState == nil {
		return nil, false
	}
	return o.MetadataState, true
}

// HasMetadataState returns a boolean if a field has been set.
func (o *BTWorkflowSnapshotInfo) HasMetadataState() bool {
	if o != nil && o.MetadataState != nil {
		return true
	}

	return false
}

// SetMetadataState gets a reference to the given string and assigns it to the MetadataState field.
func (o *BTWorkflowSnapshotInfo) SetMetadataState(v string) {
	o.MetadataState = &v
}

// GetNotifierIds returns the NotifierIds field value if set, zero value otherwise.
func (o *BTWorkflowSnapshotInfo) GetNotifierIds() []string {
	if o == nil || o.NotifierIds == nil {
		var ret []string
		return ret
	}
	return o.NotifierIds
}

// GetNotifierIdsOk returns a tuple with the NotifierIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowSnapshotInfo) GetNotifierIdsOk() ([]string, bool) {
	if o == nil || o.NotifierIds == nil {
		return nil, false
	}
	return o.NotifierIds, true
}

// HasNotifierIds returns a boolean if a field has been set.
func (o *BTWorkflowSnapshotInfo) HasNotifierIds() bool {
	if o != nil && o.NotifierIds != nil {
		return true
	}

	return false
}

// SetNotifierIds gets a reference to the given []string and assigns it to the NotifierIds field.
func (o *BTWorkflowSnapshotInfo) SetNotifierIds(v []string) {
	o.NotifierIds = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTWorkflowSnapshotInfo) GetState() BTWorkflowStateInfo {
	if o == nil || o.State == nil {
		var ret BTWorkflowStateInfo
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowSnapshotInfo) GetStateOk() (*BTWorkflowStateInfo, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTWorkflowSnapshotInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given BTWorkflowStateInfo and assigns it to the State field.
func (o *BTWorkflowSnapshotInfo) SetState(v BTWorkflowStateInfo) {
	o.State = &v
}

func (o BTWorkflowSnapshotInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.ApproverIds != nil {
		toSerialize["approverIds"] = o.ApproverIds
	}
	if o.CanBeDiscarded != nil {
		toSerialize["canBeDiscarded"] = o.CanBeDiscarded
	}
	if o.DebugMicroversionId != nil {
		toSerialize["debugMicroversionId"] = o.DebugMicroversionId
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.IsCreator != nil {
		toSerialize["isCreator"] = o.IsCreator
	}
	if o.IsDiscarded != nil {
		toSerialize["isDiscarded"] = o.IsDiscarded
	}
	if o.IsFrozen != nil {
		toSerialize["isFrozen"] = o.IsFrozen
	}
	if o.IsSetup != nil {
		toSerialize["isSetup"] = o.IsSetup
	}
	if o.MetadataState != nil {
		toSerialize["metadataState"] = o.MetadataState
	}
	if o.NotifierIds != nil {
		toSerialize["notifierIds"] = o.NotifierIds
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableBTWorkflowSnapshotInfo struct {
	value *BTWorkflowSnapshotInfo
	isSet bool
}

func (v NullableBTWorkflowSnapshotInfo) Get() *BTWorkflowSnapshotInfo {
	return v.value
}

func (v *NullableBTWorkflowSnapshotInfo) Set(val *BTWorkflowSnapshotInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkflowSnapshotInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkflowSnapshotInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkflowSnapshotInfo(val *BTWorkflowSnapshotInfo) *NullableBTWorkflowSnapshotInfo {
	return &NullableBTWorkflowSnapshotInfo{value: val, isSet: true}
}

func (v NullableBTWorkflowSnapshotInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkflowSnapshotInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
