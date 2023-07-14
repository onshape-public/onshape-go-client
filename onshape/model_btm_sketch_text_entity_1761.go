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

// BTMSketchTextEntity1761 struct for BTMSketchTextEntity1761
type BTMSketchTextEntity1761 struct {
	BtType                              *string         `json:"btType,omitempty"`
	ControlBoxIds                       []string        `json:"controlBoxIds,omitempty"`
	EntityId                            *string         `json:"entityId,omitempty"`
	EntityIdAndReplaceInDependentFields *string         `json:"entityIdAndReplaceInDependentFields,omitempty"`
	ImportMicroversion                  *string         `json:"importMicroversion,omitempty"`
	IsConstruction                      *bool           `json:"isConstruction,omitempty"`
	IsFromEndpointSplineHandle          *bool           `json:"isFromEndpointSplineHandle,omitempty"`
	IsFromSplineControlPolygon          *bool           `json:"isFromSplineControlPolygon,omitempty"`
	IsFromSplineHandle                  *bool           `json:"isFromSplineHandle,omitempty"`
	Namespace                           *string         `json:"namespace,omitempty"`
	NodeId                              *string         `json:"nodeId,omitempty"`
	Parameters                          []BTMParameter1 `json:"parameters,omitempty"`
	Ascent                              *float64        `json:"ascent,omitempty"`
	BaselineDirectionX                  *float64        `json:"baselineDirectionX,omitempty"`
	BaselineDirectionY                  *float64        `json:"baselineDirectionY,omitempty"`
	BaselineStartX                      *float64        `json:"baselineStartX,omitempty"`
	BaselineStartY                      *float64        `json:"baselineStartY,omitempty"`
	FontName                            *string         `json:"fontName,omitempty"`
	Text                                *string         `json:"text,omitempty"`
}

// NewBTMSketchTextEntity1761 instantiates a new BTMSketchTextEntity1761 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMSketchTextEntity1761() *BTMSketchTextEntity1761 {
	this := BTMSketchTextEntity1761{}
	return &this
}

// NewBTMSketchTextEntity1761WithDefaults instantiates a new BTMSketchTextEntity1761 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMSketchTextEntity1761WithDefaults() *BTMSketchTextEntity1761 {
	this := BTMSketchTextEntity1761{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMSketchTextEntity1761) SetBtType(v string) {
	o.BtType = &v
}

// GetControlBoxIds returns the ControlBoxIds field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetControlBoxIds() []string {
	if o == nil || o.ControlBoxIds == nil {
		var ret []string
		return ret
	}
	return o.ControlBoxIds
}

// GetControlBoxIdsOk returns a tuple with the ControlBoxIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetControlBoxIdsOk() ([]string, bool) {
	if o == nil || o.ControlBoxIds == nil {
		return nil, false
	}
	return o.ControlBoxIds, true
}

// HasControlBoxIds returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasControlBoxIds() bool {
	if o != nil && o.ControlBoxIds != nil {
		return true
	}

	return false
}

// SetControlBoxIds gets a reference to the given []string and assigns it to the ControlBoxIds field.
func (o *BTMSketchTextEntity1761) SetControlBoxIds(v []string) {
	o.ControlBoxIds = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *BTMSketchTextEntity1761) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityIdAndReplaceInDependentFields returns the EntityIdAndReplaceInDependentFields field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetEntityIdAndReplaceInDependentFields() string {
	if o == nil || o.EntityIdAndReplaceInDependentFields == nil {
		var ret string
		return ret
	}
	return *o.EntityIdAndReplaceInDependentFields
}

// GetEntityIdAndReplaceInDependentFieldsOk returns a tuple with the EntityIdAndReplaceInDependentFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetEntityIdAndReplaceInDependentFieldsOk() (*string, bool) {
	if o == nil || o.EntityIdAndReplaceInDependentFields == nil {
		return nil, false
	}
	return o.EntityIdAndReplaceInDependentFields, true
}

// HasEntityIdAndReplaceInDependentFields returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasEntityIdAndReplaceInDependentFields() bool {
	if o != nil && o.EntityIdAndReplaceInDependentFields != nil {
		return true
	}

	return false
}

// SetEntityIdAndReplaceInDependentFields gets a reference to the given string and assigns it to the EntityIdAndReplaceInDependentFields field.
func (o *BTMSketchTextEntity1761) SetEntityIdAndReplaceInDependentFields(v string) {
	o.EntityIdAndReplaceInDependentFields = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMSketchTextEntity1761) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetIsConstruction returns the IsConstruction field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetIsConstruction() bool {
	if o == nil || o.IsConstruction == nil {
		var ret bool
		return ret
	}
	return *o.IsConstruction
}

// GetIsConstructionOk returns a tuple with the IsConstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetIsConstructionOk() (*bool, bool) {
	if o == nil || o.IsConstruction == nil {
		return nil, false
	}
	return o.IsConstruction, true
}

// HasIsConstruction returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasIsConstruction() bool {
	if o != nil && o.IsConstruction != nil {
		return true
	}

	return false
}

// SetIsConstruction gets a reference to the given bool and assigns it to the IsConstruction field.
func (o *BTMSketchTextEntity1761) SetIsConstruction(v bool) {
	o.IsConstruction = &v
}

// GetIsFromEndpointSplineHandle returns the IsFromEndpointSplineHandle field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetIsFromEndpointSplineHandle() bool {
	if o == nil || o.IsFromEndpointSplineHandle == nil {
		var ret bool
		return ret
	}
	return *o.IsFromEndpointSplineHandle
}

// GetIsFromEndpointSplineHandleOk returns a tuple with the IsFromEndpointSplineHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetIsFromEndpointSplineHandleOk() (*bool, bool) {
	if o == nil || o.IsFromEndpointSplineHandle == nil {
		return nil, false
	}
	return o.IsFromEndpointSplineHandle, true
}

// HasIsFromEndpointSplineHandle returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasIsFromEndpointSplineHandle() bool {
	if o != nil && o.IsFromEndpointSplineHandle != nil {
		return true
	}

	return false
}

// SetIsFromEndpointSplineHandle gets a reference to the given bool and assigns it to the IsFromEndpointSplineHandle field.
func (o *BTMSketchTextEntity1761) SetIsFromEndpointSplineHandle(v bool) {
	o.IsFromEndpointSplineHandle = &v
}

// GetIsFromSplineControlPolygon returns the IsFromSplineControlPolygon field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetIsFromSplineControlPolygon() bool {
	if o == nil || o.IsFromSplineControlPolygon == nil {
		var ret bool
		return ret
	}
	return *o.IsFromSplineControlPolygon
}

// GetIsFromSplineControlPolygonOk returns a tuple with the IsFromSplineControlPolygon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetIsFromSplineControlPolygonOk() (*bool, bool) {
	if o == nil || o.IsFromSplineControlPolygon == nil {
		return nil, false
	}
	return o.IsFromSplineControlPolygon, true
}

// HasIsFromSplineControlPolygon returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasIsFromSplineControlPolygon() bool {
	if o != nil && o.IsFromSplineControlPolygon != nil {
		return true
	}

	return false
}

// SetIsFromSplineControlPolygon gets a reference to the given bool and assigns it to the IsFromSplineControlPolygon field.
func (o *BTMSketchTextEntity1761) SetIsFromSplineControlPolygon(v bool) {
	o.IsFromSplineControlPolygon = &v
}

// GetIsFromSplineHandle returns the IsFromSplineHandle field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetIsFromSplineHandle() bool {
	if o == nil || o.IsFromSplineHandle == nil {
		var ret bool
		return ret
	}
	return *o.IsFromSplineHandle
}

// GetIsFromSplineHandleOk returns a tuple with the IsFromSplineHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetIsFromSplineHandleOk() (*bool, bool) {
	if o == nil || o.IsFromSplineHandle == nil {
		return nil, false
	}
	return o.IsFromSplineHandle, true
}

// HasIsFromSplineHandle returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasIsFromSplineHandle() bool {
	if o != nil && o.IsFromSplineHandle != nil {
		return true
	}

	return false
}

// SetIsFromSplineHandle gets a reference to the given bool and assigns it to the IsFromSplineHandle field.
func (o *BTMSketchTextEntity1761) SetIsFromSplineHandle(v bool) {
	o.IsFromSplineHandle = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMSketchTextEntity1761) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMSketchTextEntity1761) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetParameters() []BTMParameter1 {
	if o == nil || o.Parameters == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetParametersOk() ([]BTMParameter1, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *BTMSketchTextEntity1761) SetParameters(v []BTMParameter1) {
	o.Parameters = v
}

// GetAscent returns the Ascent field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetAscent() float64 {
	if o == nil || o.Ascent == nil {
		var ret float64
		return ret
	}
	return *o.Ascent
}

// GetAscentOk returns a tuple with the Ascent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetAscentOk() (*float64, bool) {
	if o == nil || o.Ascent == nil {
		return nil, false
	}
	return o.Ascent, true
}

// HasAscent returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasAscent() bool {
	if o != nil && o.Ascent != nil {
		return true
	}

	return false
}

// SetAscent gets a reference to the given float64 and assigns it to the Ascent field.
func (o *BTMSketchTextEntity1761) SetAscent(v float64) {
	o.Ascent = &v
}

// GetBaselineDirectionX returns the BaselineDirectionX field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetBaselineDirectionX() float64 {
	if o == nil || o.BaselineDirectionX == nil {
		var ret float64
		return ret
	}
	return *o.BaselineDirectionX
}

// GetBaselineDirectionXOk returns a tuple with the BaselineDirectionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetBaselineDirectionXOk() (*float64, bool) {
	if o == nil || o.BaselineDirectionX == nil {
		return nil, false
	}
	return o.BaselineDirectionX, true
}

// HasBaselineDirectionX returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasBaselineDirectionX() bool {
	if o != nil && o.BaselineDirectionX != nil {
		return true
	}

	return false
}

// SetBaselineDirectionX gets a reference to the given float64 and assigns it to the BaselineDirectionX field.
func (o *BTMSketchTextEntity1761) SetBaselineDirectionX(v float64) {
	o.BaselineDirectionX = &v
}

// GetBaselineDirectionY returns the BaselineDirectionY field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetBaselineDirectionY() float64 {
	if o == nil || o.BaselineDirectionY == nil {
		var ret float64
		return ret
	}
	return *o.BaselineDirectionY
}

// GetBaselineDirectionYOk returns a tuple with the BaselineDirectionY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetBaselineDirectionYOk() (*float64, bool) {
	if o == nil || o.BaselineDirectionY == nil {
		return nil, false
	}
	return o.BaselineDirectionY, true
}

// HasBaselineDirectionY returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasBaselineDirectionY() bool {
	if o != nil && o.BaselineDirectionY != nil {
		return true
	}

	return false
}

// SetBaselineDirectionY gets a reference to the given float64 and assigns it to the BaselineDirectionY field.
func (o *BTMSketchTextEntity1761) SetBaselineDirectionY(v float64) {
	o.BaselineDirectionY = &v
}

// GetBaselineStartX returns the BaselineStartX field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetBaselineStartX() float64 {
	if o == nil || o.BaselineStartX == nil {
		var ret float64
		return ret
	}
	return *o.BaselineStartX
}

// GetBaselineStartXOk returns a tuple with the BaselineStartX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetBaselineStartXOk() (*float64, bool) {
	if o == nil || o.BaselineStartX == nil {
		return nil, false
	}
	return o.BaselineStartX, true
}

// HasBaselineStartX returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasBaselineStartX() bool {
	if o != nil && o.BaselineStartX != nil {
		return true
	}

	return false
}

// SetBaselineStartX gets a reference to the given float64 and assigns it to the BaselineStartX field.
func (o *BTMSketchTextEntity1761) SetBaselineStartX(v float64) {
	o.BaselineStartX = &v
}

// GetBaselineStartY returns the BaselineStartY field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetBaselineStartY() float64 {
	if o == nil || o.BaselineStartY == nil {
		var ret float64
		return ret
	}
	return *o.BaselineStartY
}

// GetBaselineStartYOk returns a tuple with the BaselineStartY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetBaselineStartYOk() (*float64, bool) {
	if o == nil || o.BaselineStartY == nil {
		return nil, false
	}
	return o.BaselineStartY, true
}

// HasBaselineStartY returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasBaselineStartY() bool {
	if o != nil && o.BaselineStartY != nil {
		return true
	}

	return false
}

// SetBaselineStartY gets a reference to the given float64 and assigns it to the BaselineStartY field.
func (o *BTMSketchTextEntity1761) SetBaselineStartY(v float64) {
	o.BaselineStartY = &v
}

// GetFontName returns the FontName field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetFontName() string {
	if o == nil || o.FontName == nil {
		var ret string
		return ret
	}
	return *o.FontName
}

// GetFontNameOk returns a tuple with the FontName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetFontNameOk() (*string, bool) {
	if o == nil || o.FontName == nil {
		return nil, false
	}
	return o.FontName, true
}

// HasFontName returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasFontName() bool {
	if o != nil && o.FontName != nil {
		return true
	}

	return false
}

// SetFontName gets a reference to the given string and assigns it to the FontName field.
func (o *BTMSketchTextEntity1761) SetFontName(v string) {
	o.FontName = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *BTMSketchTextEntity1761) SetText(v string) {
	o.Text = &v
}

func (o BTMSketchTextEntity1761) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ControlBoxIds != nil {
		toSerialize["controlBoxIds"] = o.ControlBoxIds
	}
	if o.EntityId != nil {
		toSerialize["entityId"] = o.EntityId
	}
	if o.EntityIdAndReplaceInDependentFields != nil {
		toSerialize["entityIdAndReplaceInDependentFields"] = o.EntityIdAndReplaceInDependentFields
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.IsConstruction != nil {
		toSerialize["isConstruction"] = o.IsConstruction
	}
	if o.IsFromEndpointSplineHandle != nil {
		toSerialize["isFromEndpointSplineHandle"] = o.IsFromEndpointSplineHandle
	}
	if o.IsFromSplineControlPolygon != nil {
		toSerialize["isFromSplineControlPolygon"] = o.IsFromSplineControlPolygon
	}
	if o.IsFromSplineHandle != nil {
		toSerialize["isFromSplineHandle"] = o.IsFromSplineHandle
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Ascent != nil {
		toSerialize["ascent"] = o.Ascent
	}
	if o.BaselineDirectionX != nil {
		toSerialize["baselineDirectionX"] = o.BaselineDirectionX
	}
	if o.BaselineDirectionY != nil {
		toSerialize["baselineDirectionY"] = o.BaselineDirectionY
	}
	if o.BaselineStartX != nil {
		toSerialize["baselineStartX"] = o.BaselineStartX
	}
	if o.BaselineStartY != nil {
		toSerialize["baselineStartY"] = o.BaselineStartY
	}
	if o.FontName != nil {
		toSerialize["fontName"] = o.FontName
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableBTMSketchTextEntity1761 struct {
	value *BTMSketchTextEntity1761
	isSet bool
}

func (v NullableBTMSketchTextEntity1761) Get() *BTMSketchTextEntity1761 {
	return v.value
}

func (v *NullableBTMSketchTextEntity1761) Set(val *BTMSketchTextEntity1761) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMSketchTextEntity1761) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMSketchTextEntity1761) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMSketchTextEntity1761(val *BTMSketchTextEntity1761) *NullableBTMSketchTextEntity1761 {
	return &NullableBTMSketchTextEntity1761{value: val, isSet: true}
}

func (v NullableBTMSketchTextEntity1761) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMSketchTextEntity1761) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
