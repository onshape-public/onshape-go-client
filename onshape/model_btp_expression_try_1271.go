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

// BTPExpressionTry1271 struct for BTPExpressionTry1271
type BTPExpressionTry1271 struct {
	BTPExpression9
	BtType              *string             `json:"btType,omitempty"`
	Atomic              *bool               `json:"atomic,omitempty"`
	DocumentationType   *GBTPDefinitionType `json:"documentationType,omitempty"`
	EndSourceLocation   *int32              `json:"endSourceLocation,omitempty"`
	NodeId              *string             `json:"nodeId,omitempty"`
	ShortDescriptor     *string             `json:"shortDescriptor,omitempty"`
	SpaceAfter          *BTPSpace10         `json:"spaceAfter,omitempty"`
	SpaceBefore         *BTPSpace10         `json:"spaceBefore,omitempty"`
	SpaceDefault        *bool               `json:"spaceDefault,omitempty"`
	StartSourceLocation *int32              `json:"startSourceLocation,omitempty"`
	Expression          *BTPExpression9     `json:"expression,omitempty"`
	Silent              *bool               `json:"silent,omitempty"`
	SpaceAfterSilent    *BTPSpace10         `json:"spaceAfterSilent,omitempty"`
	SpaceAfterTry       *BTPSpace10         `json:"spaceAfterTry,omitempty"`
	// Type of JSON object.
	BtType *string `json:"btType,omitempty"`
}

// NewBTPExpressionTry1271 instantiates a new BTPExpressionTry1271 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPExpressionTry1271() *BTPExpressionTry1271 {
	this := BTPExpressionTry1271{}
	return &this
}

// NewBTPExpressionTry1271WithDefaults instantiates a new BTPExpressionTry1271 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPExpressionTry1271WithDefaults() *BTPExpressionTry1271 {
	this := BTPExpressionTry1271{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPExpressionTry1271) SetBtType(v string) {
	o.BtType = &v
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPExpressionTry1271) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetDocumentationType() GBTPDefinitionType {
	if o == nil || o.DocumentationType == nil {
		var ret GBTPDefinitionType
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetDocumentationTypeOk() (*GBTPDefinitionType, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given GBTPDefinitionType and assigns it to the DocumentationType field.
func (o *BTPExpressionTry1271) SetDocumentationType(v GBTPDefinitionType) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPExpressionTry1271) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPExpressionTry1271) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPExpressionTry1271) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPExpressionTry1271) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPExpressionTry1271) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPExpressionTry1271) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPExpressionTry1271) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetExpression() BTPExpression9 {
	if o == nil || o.Expression == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetExpressionOk() (*BTPExpression9, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given BTPExpression9 and assigns it to the Expression field.
func (o *BTPExpressionTry1271) SetExpression(v BTPExpression9) {
	o.Expression = &v
}

// GetSilent returns the Silent field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetSilent() bool {
	if o == nil || o.Silent == nil {
		var ret bool
		return ret
	}
	return *o.Silent
}

// GetSilentOk returns a tuple with the Silent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetSilentOk() (*bool, bool) {
	if o == nil || o.Silent == nil {
		return nil, false
	}
	return o.Silent, true
}

// HasSilent returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasSilent() bool {
	if o != nil && o.Silent != nil {
		return true
	}

	return false
}

// SetSilent gets a reference to the given bool and assigns it to the Silent field.
func (o *BTPExpressionTry1271) SetSilent(v bool) {
	o.Silent = &v
}

// GetSpaceAfterSilent returns the SpaceAfterSilent field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetSpaceAfterSilent() BTPSpace10 {
	if o == nil || o.SpaceAfterSilent == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterSilent
}

// GetSpaceAfterSilentOk returns a tuple with the SpaceAfterSilent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetSpaceAfterSilentOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterSilent == nil {
		return nil, false
	}
	return o.SpaceAfterSilent, true
}

// HasSpaceAfterSilent returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasSpaceAfterSilent() bool {
	if o != nil && o.SpaceAfterSilent != nil {
		return true
	}

	return false
}

// SetSpaceAfterSilent gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterSilent field.
func (o *BTPExpressionTry1271) SetSpaceAfterSilent(v BTPSpace10) {
	o.SpaceAfterSilent = &v
}

// GetSpaceAfterTry returns the SpaceAfterTry field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetSpaceAfterTry() BTPSpace10 {
	if o == nil || o.SpaceAfterTry == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterTry
}

// GetSpaceAfterTryOk returns a tuple with the SpaceAfterTry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetSpaceAfterTryOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterTry == nil {
		return nil, false
	}
	return o.SpaceAfterTry, true
}

// HasSpaceAfterTry returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasSpaceAfterTry() bool {
	if o != nil && o.SpaceAfterTry != nil {
		return true
	}

	return false
}

// SetSpaceAfterTry gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterTry field.
func (o *BTPExpressionTry1271) SetSpaceAfterTry(v BTPSpace10) {
	o.SpaceAfterTry = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPExpressionTry1271) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionTry1271) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPExpressionTry1271) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPExpressionTry1271) SetBtType(v string) {
	o.BtType = &v
}

func (o BTPExpressionTry1271) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPExpression9, errBTPExpression9 := json.Marshal(o.BTPExpression9)
	if errBTPExpression9 != nil {
		return []byte{}, errBTPExpression9
	}
	errBTPExpression9 = json.Unmarshal([]byte(serializedBTPExpression9), &toSerialize)
	if errBTPExpression9 != nil {
		return []byte{}, errBTPExpression9
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Atomic != nil {
		toSerialize["atomic"] = o.Atomic
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
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.Silent != nil {
		toSerialize["silent"] = o.Silent
	}
	if o.SpaceAfterSilent != nil {
		toSerialize["spaceAfterSilent"] = o.SpaceAfterSilent
	}
	if o.SpaceAfterTry != nil {
		toSerialize["spaceAfterTry"] = o.SpaceAfterTry
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTPExpressionTry1271 struct {
	value *BTPExpressionTry1271
	isSet bool
}

func (v NullableBTPExpressionTry1271) Get() *BTPExpressionTry1271 {
	return v.value
}

func (v *NullableBTPExpressionTry1271) Set(val *BTPExpressionTry1271) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPExpressionTry1271) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPExpressionTry1271) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPExpressionTry1271(val *BTPExpressionTry1271) *NullableBTPExpressionTry1271 {
	return &NullableBTPExpressionTry1271{value: val, isSet: true}
}

func (v NullableBTPExpressionTry1271) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPExpressionTry1271) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
