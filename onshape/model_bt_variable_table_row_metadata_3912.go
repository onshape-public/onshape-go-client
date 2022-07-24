/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5686-86cede96e73b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTVariableTableRowMetadata3912 struct for BTVariableTableRowMetadata3912
type BTVariableTableRowMetadata3912 struct {
	CrossHighlightDataIfAny  *BTTableBaseCrossHighlightData2609 `json:"crossHighlightDataIfAny,omitempty"`
	BtType                   *string                            `json:"btType,omitempty"`
	CrossHighlightData       *BTTableBaseCrossHighlightData2609 `json:"crossHighlightData,omitempty"`
	Info                     *string                            `json:"info,omitempty"`
	IsFullyEditable          *bool                              `json:"isFullyEditable,omitempty"`
	IsRecursiveImport        *bool                              `json:"isRecursiveImport,omitempty"`
	LastWritingFeatureNodeId *string                            `json:"lastWritingFeatureNodeId,omitempty"`
}

// NewBTVariableTableRowMetadata3912 instantiates a new BTVariableTableRowMetadata3912 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTVariableTableRowMetadata3912() *BTVariableTableRowMetadata3912 {
	this := BTVariableTableRowMetadata3912{}
	return &this
}

// NewBTVariableTableRowMetadata3912WithDefaults instantiates a new BTVariableTableRowMetadata3912 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTVariableTableRowMetadata3912WithDefaults() *BTVariableTableRowMetadata3912 {
	this := BTVariableTableRowMetadata3912{}
	return &this
}

// GetCrossHighlightDataIfAny returns the CrossHighlightDataIfAny field value if set, zero value otherwise.
func (o *BTVariableTableRowMetadata3912) GetCrossHighlightDataIfAny() BTTableBaseCrossHighlightData2609 {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		var ret BTTableBaseCrossHighlightData2609
		return ret
	}
	return *o.CrossHighlightDataIfAny
}

// GetCrossHighlightDataIfAnyOk returns a tuple with the CrossHighlightDataIfAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableTableRowMetadata3912) GetCrossHighlightDataIfAnyOk() (*BTTableBaseCrossHighlightData2609, bool) {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		return nil, false
	}
	return o.CrossHighlightDataIfAny, true
}

// HasCrossHighlightDataIfAny returns a boolean if a field has been set.
func (o *BTVariableTableRowMetadata3912) HasCrossHighlightDataIfAny() bool {
	if o != nil && o.CrossHighlightDataIfAny != nil {
		return true
	}

	return false
}

// SetCrossHighlightDataIfAny gets a reference to the given BTTableBaseCrossHighlightData2609 and assigns it to the CrossHighlightDataIfAny field.
func (o *BTVariableTableRowMetadata3912) SetCrossHighlightDataIfAny(v BTTableBaseCrossHighlightData2609) {
	o.CrossHighlightDataIfAny = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTVariableTableRowMetadata3912) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableTableRowMetadata3912) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTVariableTableRowMetadata3912) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTVariableTableRowMetadata3912) SetBtType(v string) {
	o.BtType = &v
}

// GetCrossHighlightData returns the CrossHighlightData field value if set, zero value otherwise.
func (o *BTVariableTableRowMetadata3912) GetCrossHighlightData() BTTableBaseCrossHighlightData2609 {
	if o == nil || o.CrossHighlightData == nil {
		var ret BTTableBaseCrossHighlightData2609
		return ret
	}
	return *o.CrossHighlightData
}

// GetCrossHighlightDataOk returns a tuple with the CrossHighlightData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableTableRowMetadata3912) GetCrossHighlightDataOk() (*BTTableBaseCrossHighlightData2609, bool) {
	if o == nil || o.CrossHighlightData == nil {
		return nil, false
	}
	return o.CrossHighlightData, true
}

// HasCrossHighlightData returns a boolean if a field has been set.
func (o *BTVariableTableRowMetadata3912) HasCrossHighlightData() bool {
	if o != nil && o.CrossHighlightData != nil {
		return true
	}

	return false
}

// SetCrossHighlightData gets a reference to the given BTTableBaseCrossHighlightData2609 and assigns it to the CrossHighlightData field.
func (o *BTVariableTableRowMetadata3912) SetCrossHighlightData(v BTTableBaseCrossHighlightData2609) {
	o.CrossHighlightData = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *BTVariableTableRowMetadata3912) GetInfo() string {
	if o == nil || o.Info == nil {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableTableRowMetadata3912) GetInfoOk() (*string, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *BTVariableTableRowMetadata3912) HasInfo() bool {
	if o != nil && o.Info != nil {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *BTVariableTableRowMetadata3912) SetInfo(v string) {
	o.Info = &v
}

// GetIsFullyEditable returns the IsFullyEditable field value if set, zero value otherwise.
func (o *BTVariableTableRowMetadata3912) GetIsFullyEditable() bool {
	if o == nil || o.IsFullyEditable == nil {
		var ret bool
		return ret
	}
	return *o.IsFullyEditable
}

// GetIsFullyEditableOk returns a tuple with the IsFullyEditable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableTableRowMetadata3912) GetIsFullyEditableOk() (*bool, bool) {
	if o == nil || o.IsFullyEditable == nil {
		return nil, false
	}
	return o.IsFullyEditable, true
}

// HasIsFullyEditable returns a boolean if a field has been set.
func (o *BTVariableTableRowMetadata3912) HasIsFullyEditable() bool {
	if o != nil && o.IsFullyEditable != nil {
		return true
	}

	return false
}

// SetIsFullyEditable gets a reference to the given bool and assigns it to the IsFullyEditable field.
func (o *BTVariableTableRowMetadata3912) SetIsFullyEditable(v bool) {
	o.IsFullyEditable = &v
}

// GetIsRecursiveImport returns the IsRecursiveImport field value if set, zero value otherwise.
func (o *BTVariableTableRowMetadata3912) GetIsRecursiveImport() bool {
	if o == nil || o.IsRecursiveImport == nil {
		var ret bool
		return ret
	}
	return *o.IsRecursiveImport
}

// GetIsRecursiveImportOk returns a tuple with the IsRecursiveImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableTableRowMetadata3912) GetIsRecursiveImportOk() (*bool, bool) {
	if o == nil || o.IsRecursiveImport == nil {
		return nil, false
	}
	return o.IsRecursiveImport, true
}

// HasIsRecursiveImport returns a boolean if a field has been set.
func (o *BTVariableTableRowMetadata3912) HasIsRecursiveImport() bool {
	if o != nil && o.IsRecursiveImport != nil {
		return true
	}

	return false
}

// SetIsRecursiveImport gets a reference to the given bool and assigns it to the IsRecursiveImport field.
func (o *BTVariableTableRowMetadata3912) SetIsRecursiveImport(v bool) {
	o.IsRecursiveImport = &v
}

// GetLastWritingFeatureNodeId returns the LastWritingFeatureNodeId field value if set, zero value otherwise.
func (o *BTVariableTableRowMetadata3912) GetLastWritingFeatureNodeId() string {
	if o == nil || o.LastWritingFeatureNodeId == nil {
		var ret string
		return ret
	}
	return *o.LastWritingFeatureNodeId
}

// GetLastWritingFeatureNodeIdOk returns a tuple with the LastWritingFeatureNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableTableRowMetadata3912) GetLastWritingFeatureNodeIdOk() (*string, bool) {
	if o == nil || o.LastWritingFeatureNodeId == nil {
		return nil, false
	}
	return o.LastWritingFeatureNodeId, true
}

// HasLastWritingFeatureNodeId returns a boolean if a field has been set.
func (o *BTVariableTableRowMetadata3912) HasLastWritingFeatureNodeId() bool {
	if o != nil && o.LastWritingFeatureNodeId != nil {
		return true
	}

	return false
}

// SetLastWritingFeatureNodeId gets a reference to the given string and assigns it to the LastWritingFeatureNodeId field.
func (o *BTVariableTableRowMetadata3912) SetLastWritingFeatureNodeId(v string) {
	o.LastWritingFeatureNodeId = &v
}

func (o BTVariableTableRowMetadata3912) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CrossHighlightDataIfAny != nil {
		toSerialize["crossHighlightDataIfAny"] = o.CrossHighlightDataIfAny
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CrossHighlightData != nil {
		toSerialize["crossHighlightData"] = o.CrossHighlightData
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.IsFullyEditable != nil {
		toSerialize["isFullyEditable"] = o.IsFullyEditable
	}
	if o.IsRecursiveImport != nil {
		toSerialize["isRecursiveImport"] = o.IsRecursiveImport
	}
	if o.LastWritingFeatureNodeId != nil {
		toSerialize["lastWritingFeatureNodeId"] = o.LastWritingFeatureNodeId
	}
	return json.Marshal(toSerialize)
}

type NullableBTVariableTableRowMetadata3912 struct {
	value *BTVariableTableRowMetadata3912
	isSet bool
}

func (v NullableBTVariableTableRowMetadata3912) Get() *BTVariableTableRowMetadata3912 {
	return v.value
}

func (v *NullableBTVariableTableRowMetadata3912) Set(val *BTVariableTableRowMetadata3912) {
	v.value = val
	v.isSet = true
}

func (v NullableBTVariableTableRowMetadata3912) IsSet() bool {
	return v.isSet
}

func (v *NullableBTVariableTableRowMetadata3912) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTVariableTableRowMetadata3912(val *BTVariableTableRowMetadata3912) *NullableBTVariableTableRowMetadata3912 {
	return &NullableBTVariableTableRowMetadata3912{value: val, isSet: true}
}

func (v NullableBTVariableTableRowMetadata3912) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTVariableTableRowMetadata3912) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
