/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.168.21279-402b6292597b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPStatementLoopWhile280 struct for BTPStatementLoopWhile280
type BTPStatementLoopWhile280 struct {
	Annotation          *BTPAnnotation231   `json:"annotation,omitempty"`
	Atomic              *bool               `json:"atomic,omitempty"`
	BtType              *string             `json:"btType,omitempty"`
	DocumentationType   *GBTPDefinitionType `json:"documentationType,omitempty"`
	EndSourceLocation   *int32              `json:"endSourceLocation,omitempty"`
	NodeId              *string             `json:"nodeId,omitempty"`
	ShortDescriptor     *string             `json:"shortDescriptor,omitempty"`
	SpaceAfter          *BTPSpace10         `json:"spaceAfter,omitempty"`
	SpaceBefore         *BTPSpace10         `json:"spaceBefore,omitempty"`
	SpaceDefault        *bool               `json:"spaceDefault,omitempty"`
	StartSourceLocation *int32              `json:"startSourceLocation,omitempty"`
	Body                *BTPStatement269    `json:"body,omitempty"`
	SpaceAfterLoopType  *BTPSpace10         `json:"spaceAfterLoopType,omitempty"`
	Condition           *BTPExpression9     `json:"condition,omitempty"`
}

// NewBTPStatementLoopWhile280 instantiates a new BTPStatementLoopWhile280 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPStatementLoopWhile280() *BTPStatementLoopWhile280 {
	this := BTPStatementLoopWhile280{}
	return &this
}

// NewBTPStatementLoopWhile280WithDefaults instantiates a new BTPStatementLoopWhile280 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPStatementLoopWhile280WithDefaults() *BTPStatementLoopWhile280 {
	this := BTPStatementLoopWhile280{}
	return &this
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetAnnotation() BTPAnnotation231 {
	if o == nil || o.Annotation == nil {
		var ret BTPAnnotation231
		return ret
	}
	return *o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetAnnotationOk() (*BTPAnnotation231, bool) {
	if o == nil || o.Annotation == nil {
		return nil, false
	}
	return o.Annotation, true
}

// HasAnnotation returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasAnnotation() bool {
	if o != nil && o.Annotation != nil {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given BTPAnnotation231 and assigns it to the Annotation field.
func (o *BTPStatementLoopWhile280) SetAnnotation(v BTPAnnotation231) {
	o.Annotation = &v
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPStatementLoopWhile280) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPStatementLoopWhile280) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetDocumentationType() GBTPDefinitionType {
	if o == nil || o.DocumentationType == nil {
		var ret GBTPDefinitionType
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetDocumentationTypeOk() (*GBTPDefinitionType, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given GBTPDefinitionType and assigns it to the DocumentationType field.
func (o *BTPStatementLoopWhile280) SetDocumentationType(v GBTPDefinitionType) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPStatementLoopWhile280) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPStatementLoopWhile280) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPStatementLoopWhile280) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPStatementLoopWhile280) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPStatementLoopWhile280) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPStatementLoopWhile280) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPStatementLoopWhile280) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetBody() BTPStatement269 {
	if o == nil || o.Body == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetBodyOk() (*BTPStatement269, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given BTPStatement269 and assigns it to the Body field.
func (o *BTPStatementLoopWhile280) SetBody(v BTPStatement269) {
	o.Body = &v
}

// GetSpaceAfterLoopType returns the SpaceAfterLoopType field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetSpaceAfterLoopType() BTPSpace10 {
	if o == nil || o.SpaceAfterLoopType == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterLoopType
}

// GetSpaceAfterLoopTypeOk returns a tuple with the SpaceAfterLoopType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetSpaceAfterLoopTypeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterLoopType == nil {
		return nil, false
	}
	return o.SpaceAfterLoopType, true
}

// HasSpaceAfterLoopType returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasSpaceAfterLoopType() bool {
	if o != nil && o.SpaceAfterLoopType != nil {
		return true
	}

	return false
}

// SetSpaceAfterLoopType gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterLoopType field.
func (o *BTPStatementLoopWhile280) SetSpaceAfterLoopType(v BTPSpace10) {
	o.SpaceAfterLoopType = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetCondition() BTPExpression9 {
	if o == nil || o.Condition == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetConditionOk() (*BTPExpression9, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given BTPExpression9 and assigns it to the Condition field.
func (o *BTPStatementLoopWhile280) SetCondition(v BTPExpression9) {
	o.Condition = &v
}

func (o BTPStatementLoopWhile280) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Annotation != nil {
		toSerialize["annotation"] = o.Annotation
	}
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
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.SpaceAfterLoopType != nil {
		toSerialize["spaceAfterLoopType"] = o.SpaceAfterLoopType
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	return json.Marshal(toSerialize)
}

type NullableBTPStatementLoopWhile280 struct {
	value *BTPStatementLoopWhile280
	isSet bool
}

func (v NullableBTPStatementLoopWhile280) Get() *BTPStatementLoopWhile280 {
	return v.value
}

func (v *NullableBTPStatementLoopWhile280) Set(val *BTPStatementLoopWhile280) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPStatementLoopWhile280) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPStatementLoopWhile280) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPStatementLoopWhile280(val *BTPStatementLoopWhile280) *NullableBTPStatementLoopWhile280 {
	return &NullableBTPStatementLoopWhile280{value: val, isSet: true}
}

func (v NullableBTPStatementLoopWhile280) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPStatementLoopWhile280) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
