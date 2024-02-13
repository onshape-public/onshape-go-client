/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPExpressionOperator244 struct for BTPExpressionOperator244
type BTPExpressionOperator244 struct {
	Atomic                *bool               `json:"atomic,omitempty"`
	BtType                *string             `json:"btType,omitempty"`
	DocumentationType     *GBTPDefinitionType `json:"documentationType,omitempty"`
	EndSourceLocation     *int32              `json:"endSourceLocation,omitempty"`
	NodeId                *string             `json:"nodeId,omitempty"`
	ShortDescriptor       *string             `json:"shortDescriptor,omitempty"`
	SpaceAfter            *BTPSpace10         `json:"spaceAfter,omitempty"`
	SpaceBefore           *BTPSpace10         `json:"spaceBefore,omitempty"`
	SpaceDefault          *bool               `json:"spaceDefault,omitempty"`
	StartSourceLocation   *int32              `json:"startSourceLocation,omitempty"`
	ForExport             *bool               `json:"forExport,omitempty"`
	GlobalNamespace       *bool               `json:"globalNamespace,omitempty"`
	ImportMicroversion    *string             `json:"importMicroversion,omitempty"`
	Namespace             []BTPIdentifier8    `json:"namespace,omitempty"`
	Operand1              *BTPExpression9     `json:"operand1,omitempty"`
	Operand2              *BTPExpression9     `json:"operand2,omitempty"`
	Operand3              *BTPExpression9     `json:"operand3,omitempty"`
	Operator              *GBTPOperator       `json:"operator,omitempty"`
	SpaceAfterNamespace   *BTPSpace10         `json:"spaceAfterNamespace,omitempty"`
	SpaceAfterOperator    *BTPSpace10         `json:"spaceAfterOperator,omitempty"`
	SpaceBeforeOperator   *BTPSpace10         `json:"spaceBeforeOperator,omitempty"`
	WrittenAsFunctionCall *bool               `json:"writtenAsFunctionCall,omitempty"`
}

// NewBTPExpressionOperator244 instantiates a new BTPExpressionOperator244 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPExpressionOperator244() *BTPExpressionOperator244 {
	this := BTPExpressionOperator244{}
	return &this
}

// NewBTPExpressionOperator244WithDefaults instantiates a new BTPExpressionOperator244 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPExpressionOperator244WithDefaults() *BTPExpressionOperator244 {
	this := BTPExpressionOperator244{}
	return &this
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPExpressionOperator244) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPExpressionOperator244) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetDocumentationType() GBTPDefinitionType {
	if o == nil || o.DocumentationType == nil {
		var ret GBTPDefinitionType
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetDocumentationTypeOk() (*GBTPDefinitionType, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given GBTPDefinitionType and assigns it to the DocumentationType field.
func (o *BTPExpressionOperator244) SetDocumentationType(v GBTPDefinitionType) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPExpressionOperator244) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPExpressionOperator244) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPExpressionOperator244) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPExpressionOperator244) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPExpressionOperator244) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPExpressionOperator244) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPExpressionOperator244) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetForExport returns the ForExport field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetForExport() bool {
	if o == nil || o.ForExport == nil {
		var ret bool
		return ret
	}
	return *o.ForExport
}

// GetForExportOk returns a tuple with the ForExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetForExportOk() (*bool, bool) {
	if o == nil || o.ForExport == nil {
		return nil, false
	}
	return o.ForExport, true
}

// HasForExport returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasForExport() bool {
	if o != nil && o.ForExport != nil {
		return true
	}

	return false
}

// SetForExport gets a reference to the given bool and assigns it to the ForExport field.
func (o *BTPExpressionOperator244) SetForExport(v bool) {
	o.ForExport = &v
}

// GetGlobalNamespace returns the GlobalNamespace field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetGlobalNamespace() bool {
	if o == nil || o.GlobalNamespace == nil {
		var ret bool
		return ret
	}
	return *o.GlobalNamespace
}

// GetGlobalNamespaceOk returns a tuple with the GlobalNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetGlobalNamespaceOk() (*bool, bool) {
	if o == nil || o.GlobalNamespace == nil {
		return nil, false
	}
	return o.GlobalNamespace, true
}

// HasGlobalNamespace returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasGlobalNamespace() bool {
	if o != nil && o.GlobalNamespace != nil {
		return true
	}

	return false
}

// SetGlobalNamespace gets a reference to the given bool and assigns it to the GlobalNamespace field.
func (o *BTPExpressionOperator244) SetGlobalNamespace(v bool) {
	o.GlobalNamespace = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTPExpressionOperator244) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetNamespace() []BTPIdentifier8 {
	if o == nil || o.Namespace == nil {
		var ret []BTPIdentifier8
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetNamespaceOk() ([]BTPIdentifier8, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given []BTPIdentifier8 and assigns it to the Namespace field.
func (o *BTPExpressionOperator244) SetNamespace(v []BTPIdentifier8) {
	o.Namespace = v
}

// GetOperand1 returns the Operand1 field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetOperand1() BTPExpression9 {
	if o == nil || o.Operand1 == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Operand1
}

// GetOperand1Ok returns a tuple with the Operand1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetOperand1Ok() (*BTPExpression9, bool) {
	if o == nil || o.Operand1 == nil {
		return nil, false
	}
	return o.Operand1, true
}

// HasOperand1 returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasOperand1() bool {
	if o != nil && o.Operand1 != nil {
		return true
	}

	return false
}

// SetOperand1 gets a reference to the given BTPExpression9 and assigns it to the Operand1 field.
func (o *BTPExpressionOperator244) SetOperand1(v BTPExpression9) {
	o.Operand1 = &v
}

// GetOperand2 returns the Operand2 field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetOperand2() BTPExpression9 {
	if o == nil || o.Operand2 == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Operand2
}

// GetOperand2Ok returns a tuple with the Operand2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetOperand2Ok() (*BTPExpression9, bool) {
	if o == nil || o.Operand2 == nil {
		return nil, false
	}
	return o.Operand2, true
}

// HasOperand2 returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasOperand2() bool {
	if o != nil && o.Operand2 != nil {
		return true
	}

	return false
}

// SetOperand2 gets a reference to the given BTPExpression9 and assigns it to the Operand2 field.
func (o *BTPExpressionOperator244) SetOperand2(v BTPExpression9) {
	o.Operand2 = &v
}

// GetOperand3 returns the Operand3 field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetOperand3() BTPExpression9 {
	if o == nil || o.Operand3 == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Operand3
}

// GetOperand3Ok returns a tuple with the Operand3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetOperand3Ok() (*BTPExpression9, bool) {
	if o == nil || o.Operand3 == nil {
		return nil, false
	}
	return o.Operand3, true
}

// HasOperand3 returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasOperand3() bool {
	if o != nil && o.Operand3 != nil {
		return true
	}

	return false
}

// SetOperand3 gets a reference to the given BTPExpression9 and assigns it to the Operand3 field.
func (o *BTPExpressionOperator244) SetOperand3(v BTPExpression9) {
	o.Operand3 = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetOperator() GBTPOperator {
	if o == nil || o.Operator == nil {
		var ret GBTPOperator
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetOperatorOk() (*GBTPOperator, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given GBTPOperator and assigns it to the Operator field.
func (o *BTPExpressionOperator244) SetOperator(v GBTPOperator) {
	o.Operator = &v
}

// GetSpaceAfterNamespace returns the SpaceAfterNamespace field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetSpaceAfterNamespace() BTPSpace10 {
	if o == nil || o.SpaceAfterNamespace == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterNamespace
}

// GetSpaceAfterNamespaceOk returns a tuple with the SpaceAfterNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetSpaceAfterNamespaceOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterNamespace == nil {
		return nil, false
	}
	return o.SpaceAfterNamespace, true
}

// HasSpaceAfterNamespace returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasSpaceAfterNamespace() bool {
	if o != nil && o.SpaceAfterNamespace != nil {
		return true
	}

	return false
}

// SetSpaceAfterNamespace gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterNamespace field.
func (o *BTPExpressionOperator244) SetSpaceAfterNamespace(v BTPSpace10) {
	o.SpaceAfterNamespace = &v
}

// GetSpaceAfterOperator returns the SpaceAfterOperator field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetSpaceAfterOperator() BTPSpace10 {
	if o == nil || o.SpaceAfterOperator == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterOperator
}

// GetSpaceAfterOperatorOk returns a tuple with the SpaceAfterOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetSpaceAfterOperatorOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterOperator == nil {
		return nil, false
	}
	return o.SpaceAfterOperator, true
}

// HasSpaceAfterOperator returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasSpaceAfterOperator() bool {
	if o != nil && o.SpaceAfterOperator != nil {
		return true
	}

	return false
}

// SetSpaceAfterOperator gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterOperator field.
func (o *BTPExpressionOperator244) SetSpaceAfterOperator(v BTPSpace10) {
	o.SpaceAfterOperator = &v
}

// GetSpaceBeforeOperator returns the SpaceBeforeOperator field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetSpaceBeforeOperator() BTPSpace10 {
	if o == nil || o.SpaceBeforeOperator == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBeforeOperator
}

// GetSpaceBeforeOperatorOk returns a tuple with the SpaceBeforeOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetSpaceBeforeOperatorOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBeforeOperator == nil {
		return nil, false
	}
	return o.SpaceBeforeOperator, true
}

// HasSpaceBeforeOperator returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasSpaceBeforeOperator() bool {
	if o != nil && o.SpaceBeforeOperator != nil {
		return true
	}

	return false
}

// SetSpaceBeforeOperator gets a reference to the given BTPSpace10 and assigns it to the SpaceBeforeOperator field.
func (o *BTPExpressionOperator244) SetSpaceBeforeOperator(v BTPSpace10) {
	o.SpaceBeforeOperator = &v
}

// GetWrittenAsFunctionCall returns the WrittenAsFunctionCall field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetWrittenAsFunctionCall() bool {
	if o == nil || o.WrittenAsFunctionCall == nil {
		var ret bool
		return ret
	}
	return *o.WrittenAsFunctionCall
}

// GetWrittenAsFunctionCallOk returns a tuple with the WrittenAsFunctionCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetWrittenAsFunctionCallOk() (*bool, bool) {
	if o == nil || o.WrittenAsFunctionCall == nil {
		return nil, false
	}
	return o.WrittenAsFunctionCall, true
}

// HasWrittenAsFunctionCall returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasWrittenAsFunctionCall() bool {
	if o != nil && o.WrittenAsFunctionCall != nil {
		return true
	}

	return false
}

// SetWrittenAsFunctionCall gets a reference to the given bool and assigns it to the WrittenAsFunctionCall field.
func (o *BTPExpressionOperator244) SetWrittenAsFunctionCall(v bool) {
	o.WrittenAsFunctionCall = &v
}

func (o BTPExpressionOperator244) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Atomic != nil {
		toSerialize["atomic"] = o.Atomic
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DocumentationType != nil {
		toSerialize["documentationType"] = o.DocumentationType
	}
	if o.EndSourceLocation != nil {
		toSerialize["endSourceLocation"] = o.EndSourceLocation
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ShortDescriptor != nil {
		toSerialize["shortDescriptor"] = o.ShortDescriptor
	}
	if o.SpaceAfter != nil {
		toSerialize["spaceAfter"] = o.SpaceAfter
	}
	if o.SpaceBefore != nil {
		toSerialize["spaceBefore"] = o.SpaceBefore
	}
	if o.SpaceDefault != nil {
		toSerialize["spaceDefault"] = o.SpaceDefault
	}
	if o.StartSourceLocation != nil {
		toSerialize["startSourceLocation"] = o.StartSourceLocation
	}
	if o.ForExport != nil {
		toSerialize["forExport"] = o.ForExport
	}
	if o.GlobalNamespace != nil {
		toSerialize["globalNamespace"] = o.GlobalNamespace
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Operand1 != nil {
		toSerialize["operand1"] = o.Operand1
	}
	if o.Operand2 != nil {
		toSerialize["operand2"] = o.Operand2
	}
	if o.Operand3 != nil {
		toSerialize["operand3"] = o.Operand3
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.SpaceAfterNamespace != nil {
		toSerialize["spaceAfterNamespace"] = o.SpaceAfterNamespace
	}
	if o.SpaceAfterOperator != nil {
		toSerialize["spaceAfterOperator"] = o.SpaceAfterOperator
	}
	if o.SpaceBeforeOperator != nil {
		toSerialize["spaceBeforeOperator"] = o.SpaceBeforeOperator
	}
	if o.WrittenAsFunctionCall != nil {
		toSerialize["writtenAsFunctionCall"] = o.WrittenAsFunctionCall
	}
	return json.Marshal(toSerialize)
}

type NullableBTPExpressionOperator244 struct {
	value *BTPExpressionOperator244
	isSet bool
}

func (v NullableBTPExpressionOperator244) Get() *BTPExpressionOperator244 {
	return v.value
}

func (v *NullableBTPExpressionOperator244) Set(val *BTPExpressionOperator244) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPExpressionOperator244) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPExpressionOperator244) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPExpressionOperator244(val *BTPExpressionOperator244) *NullableBTPExpressionOperator244 {
	return &NullableBTPExpressionOperator244{value: val, isSet: true}
}

func (v NullableBTPExpressionOperator244) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPExpressionOperator244) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
