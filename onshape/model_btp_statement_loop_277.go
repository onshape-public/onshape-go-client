/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.150.5633-5ed6b38daa6b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTPStatementLoop277 - struct for BTPStatementLoop277
type BTPStatementLoop277 struct {
	implBTPStatementLoop277 interface{}
}

// BTPStatementLoopForIn279AsBTPStatementLoop277 is a convenience function that returns BTPStatementLoopForIn279 wrapped in BTPStatementLoop277
func (o *BTPStatementLoopForIn279) AsBTPStatementLoop277() *BTPStatementLoop277 {
	return &BTPStatementLoop277{o}
}

// BTPStatementLoopFor3278AsBTPStatementLoop277 is a convenience function that returns BTPStatementLoopFor3278 wrapped in BTPStatementLoop277
func (o *BTPStatementLoopFor3278) AsBTPStatementLoop277() *BTPStatementLoop277 {
	return &BTPStatementLoop277{o}
}

// BTPStatementLoopWhile280AsBTPStatementLoop277 is a convenience function that returns BTPStatementLoopWhile280 wrapped in BTPStatementLoop277
func (o *BTPStatementLoopWhile280) AsBTPStatementLoop277() *BTPStatementLoop277 {
	return &BTPStatementLoop277{o}
}

// NewBTPStatementLoop277 instantiates a new BTPStatementLoop277 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPStatementLoop277() *BTPStatementLoop277 {
	this := BTPStatementLoop277{Newbase_BTPStatementLoop277()}
	return &this
}

// NewBTPStatementLoop277WithDefaults instantiates a new BTPStatementLoop277 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPStatementLoop277WithDefaults() *BTPStatementLoop277 {
	this := BTPStatementLoop277{Newbase_BTPStatementLoop277WithDefaults()}
	return &this
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *BTPStatementLoop277) GetAnnotation() BTPAnnotation231 {
	type getResult interface {
		GetAnnotation() BTPAnnotation231
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetAnnotation()
	} else {
		var de BTPAnnotation231
		return de
	}
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277) GetAnnotationOk() (*BTPAnnotation231, bool) {
	type getResult interface {
		GetAnnotationOk() (*BTPAnnotation231, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetAnnotationOk()
	} else {
		return nil, false
	}
}

// HasAnnotation returns a boolean if a field has been set.
func (o *BTPStatementLoop277) HasAnnotation() bool {
	type getResult interface {
		HasAnnotation() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasAnnotation()
	} else {
		return false
	}
}

// SetAnnotation gets a reference to the given BTPAnnotation231 and assigns it to the Annotation field.
func (o *BTPStatementLoop277) SetAnnotation(v BTPAnnotation231) {
	type getResult interface {
		SetAnnotation(v BTPAnnotation231)
	}

	o.GetActualInstance().(getResult).SetAnnotation(v)
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPStatementLoop277) GetAtomic() bool {
	type getResult interface {
		GetAtomic() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetAtomic()
	} else {
		var de bool
		return de
	}
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277) GetAtomicOk() (*bool, bool) {
	type getResult interface {
		GetAtomicOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetAtomicOk()
	} else {
		return nil, false
	}
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPStatementLoop277) HasAtomic() bool {
	type getResult interface {
		HasAtomic() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasAtomic()
	} else {
		return false
	}
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPStatementLoop277) SetAtomic(v bool) {
	type getResult interface {
		SetAtomic(v bool)
	}

	o.GetActualInstance().(getResult).SetAtomic(v)
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPStatementLoop277) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPStatementLoop277) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPStatementLoop277) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPStatementLoop277) GetDocumentationType() string {
	type getResult interface {
		GetDocumentationType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDocumentationType()
	} else {
		var de string
		return de
	}
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277) GetDocumentationTypeOk() (*string, bool) {
	type getResult interface {
		GetDocumentationTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDocumentationTypeOk()
	} else {
		return nil, false
	}
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPStatementLoop277) HasDocumentationType() bool {
	type getResult interface {
		HasDocumentationType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDocumentationType()
	} else {
		return false
	}
}

// SetDocumentationType gets a reference to the given string and assigns it to the DocumentationType field.
func (o *BTPStatementLoop277) SetDocumentationType(v string) {
	type getResult interface {
		SetDocumentationType(v string)
	}

	o.GetActualInstance().(getResult).SetDocumentationType(v)
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPStatementLoop277) GetEndSourceLocation() int32 {
	type getResult interface {
		GetEndSourceLocation() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEndSourceLocation()
	} else {
		var de int32
		return de
	}
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277) GetEndSourceLocationOk() (*int32, bool) {
	type getResult interface {
		GetEndSourceLocationOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEndSourceLocationOk()
	} else {
		return nil, false
	}
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPStatementLoop277) HasEndSourceLocation() bool {
	type getResult interface {
		HasEndSourceLocation() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasEndSourceLocation()
	} else {
		return false
	}
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPStatementLoop277) SetEndSourceLocation(v int32) {
	type getResult interface {
		SetEndSourceLocation(v int32)
	}

	o.GetActualInstance().(getResult).SetEndSourceLocation(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPStatementLoop277) GetNodeId() string {
	type getResult interface {
		GetNodeId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeId()
	} else {
		var de string
		return de
	}
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277) GetNodeIdOk() (*string, bool) {
	type getResult interface {
		GetNodeIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeIdOk()
	} else {
		return nil, false
	}
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPStatementLoop277) HasNodeId() bool {
	type getResult interface {
		HasNodeId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNodeId()
	} else {
		return false
	}
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPStatementLoop277) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPStatementLoop277) GetShortDescriptor() string {
	type getResult interface {
		GetShortDescriptor() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetShortDescriptor()
	} else {
		var de string
		return de
	}
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277) GetShortDescriptorOk() (*string, bool) {
	type getResult interface {
		GetShortDescriptorOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetShortDescriptorOk()
	} else {
		return nil, false
	}
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPStatementLoop277) HasShortDescriptor() bool {
	type getResult interface {
		HasShortDescriptor() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasShortDescriptor()
	} else {
		return false
	}
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPStatementLoop277) SetShortDescriptor(v string) {
	type getResult interface {
		SetShortDescriptor(v string)
	}

	o.GetActualInstance().(getResult).SetShortDescriptor(v)
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPStatementLoop277) GetSpaceAfter() BTPSpace10 {
	type getResult interface {
		GetSpaceAfter() BTPSpace10
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSpaceAfter()
	} else {
		var de BTPSpace10
		return de
	}
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277) GetSpaceAfterOk() (*BTPSpace10, bool) {
	type getResult interface {
		GetSpaceAfterOk() (*BTPSpace10, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSpaceAfterOk()
	} else {
		return nil, false
	}
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPStatementLoop277) HasSpaceAfter() bool {
	type getResult interface {
		HasSpaceAfter() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasSpaceAfter()
	} else {
		return false
	}
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPStatementLoop277) SetSpaceAfter(v BTPSpace10) {
	type getResult interface {
		SetSpaceAfter(v BTPSpace10)
	}

	o.GetActualInstance().(getResult).SetSpaceAfter(v)
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPStatementLoop277) GetSpaceBefore() BTPSpace10 {
	type getResult interface {
		GetSpaceBefore() BTPSpace10
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSpaceBefore()
	} else {
		var de BTPSpace10
		return de
	}
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	type getResult interface {
		GetSpaceBeforeOk() (*BTPSpace10, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSpaceBeforeOk()
	} else {
		return nil, false
	}
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPStatementLoop277) HasSpaceBefore() bool {
	type getResult interface {
		HasSpaceBefore() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasSpaceBefore()
	} else {
		return false
	}
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPStatementLoop277) SetSpaceBefore(v BTPSpace10) {
	type getResult interface {
		SetSpaceBefore(v BTPSpace10)
	}

	o.GetActualInstance().(getResult).SetSpaceBefore(v)
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPStatementLoop277) GetSpaceDefault() bool {
	type getResult interface {
		GetSpaceDefault() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSpaceDefault()
	} else {
		var de bool
		return de
	}
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277) GetSpaceDefaultOk() (*bool, bool) {
	type getResult interface {
		GetSpaceDefaultOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSpaceDefaultOk()
	} else {
		return nil, false
	}
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPStatementLoop277) HasSpaceDefault() bool {
	type getResult interface {
		HasSpaceDefault() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasSpaceDefault()
	} else {
		return false
	}
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPStatementLoop277) SetSpaceDefault(v bool) {
	type getResult interface {
		SetSpaceDefault(v bool)
	}

	o.GetActualInstance().(getResult).SetSpaceDefault(v)
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPStatementLoop277) GetStartSourceLocation() int32 {
	type getResult interface {
		GetStartSourceLocation() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetStartSourceLocation()
	} else {
		var de int32
		return de
	}
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277) GetStartSourceLocationOk() (*int32, bool) {
	type getResult interface {
		GetStartSourceLocationOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetStartSourceLocationOk()
	} else {
		return nil, false
	}
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPStatementLoop277) HasStartSourceLocation() bool {
	type getResult interface {
		HasStartSourceLocation() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasStartSourceLocation()
	} else {
		return false
	}
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPStatementLoop277) SetStartSourceLocation(v int32) {
	type getResult interface {
		SetStartSourceLocation(v int32)
	}

	o.GetActualInstance().(getResult).SetStartSourceLocation(v)
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BTPStatementLoop277) GetBody() BTPStatement269 {
	type getResult interface {
		GetBody() BTPStatement269
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBody()
	} else {
		var de BTPStatement269
		return de
	}
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277) GetBodyOk() (*BTPStatement269, bool) {
	type getResult interface {
		GetBodyOk() (*BTPStatement269, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBodyOk()
	} else {
		return nil, false
	}
}

// HasBody returns a boolean if a field has been set.
func (o *BTPStatementLoop277) HasBody() bool {
	type getResult interface {
		HasBody() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBody()
	} else {
		return false
	}
}

// SetBody gets a reference to the given BTPStatement269 and assigns it to the Body field.
func (o *BTPStatementLoop277) SetBody(v BTPStatement269) {
	type getResult interface {
		SetBody(v BTPStatement269)
	}

	o.GetActualInstance().(getResult).SetBody(v)
}

// GetSpaceAfterLoopType returns the SpaceAfterLoopType field value if set, zero value otherwise.
func (o *BTPStatementLoop277) GetSpaceAfterLoopType() BTPSpace10 {
	type getResult interface {
		GetSpaceAfterLoopType() BTPSpace10
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSpaceAfterLoopType()
	} else {
		var de BTPSpace10
		return de
	}
}

// GetSpaceAfterLoopTypeOk returns a tuple with the SpaceAfterLoopType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277) GetSpaceAfterLoopTypeOk() (*BTPSpace10, bool) {
	type getResult interface {
		GetSpaceAfterLoopTypeOk() (*BTPSpace10, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSpaceAfterLoopTypeOk()
	} else {
		return nil, false
	}
}

// HasSpaceAfterLoopType returns a boolean if a field has been set.
func (o *BTPStatementLoop277) HasSpaceAfterLoopType() bool {
	type getResult interface {
		HasSpaceAfterLoopType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasSpaceAfterLoopType()
	} else {
		return false
	}
}

// SetSpaceAfterLoopType gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterLoopType field.
func (o *BTPStatementLoop277) SetSpaceAfterLoopType(v BTPSpace10) {
	type getResult interface {
		SetSpaceAfterLoopType(v BTPSpace10)
	}

	o.GetActualInstance().(getResult).SetSpaceAfterLoopType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTPStatementLoop277) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTPStatementLoopFor3-278'
	if jsonDict["btType"] == "BTPStatementLoopFor3-278" {
		// try to unmarshal JSON data into BTPStatementLoopFor3278
		var qr *BTPStatementLoopFor3278
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTPStatementLoop277 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTPStatementLoop277 = nil
			return fmt.Errorf("Failed to unmarshal BTPStatementLoop277 as BTPStatementLoopFor3278: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTPStatementLoopForIn-279'
	if jsonDict["btType"] == "BTPStatementLoopForIn-279" {
		// try to unmarshal JSON data into BTPStatementLoopForIn279
		var qr *BTPStatementLoopForIn279
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTPStatementLoop277 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTPStatementLoop277 = nil
			return fmt.Errorf("Failed to unmarshal BTPStatementLoop277 as BTPStatementLoopForIn279: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTPStatementLoopWhile-280'
	if jsonDict["btType"] == "BTPStatementLoopWhile-280" {
		// try to unmarshal JSON data into BTPStatementLoopWhile280
		var qr *BTPStatementLoopWhile280
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTPStatementLoop277 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTPStatementLoop277 = nil
			return fmt.Errorf("Failed to unmarshal BTPStatementLoop277 as BTPStatementLoopWhile280: %s", err.Error())
		}
	}

	var qtx *base_BTPStatementLoop277
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTPStatementLoop277 = qtx
		return nil // data stored in dst.base_BTPStatementLoop277, return on the first match
	} else {
		dst.implBTPStatementLoop277 = nil
		return fmt.Errorf("Failed to unmarshal BTPStatementLoop277 as base_BTPStatementLoop277: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTPStatementLoop277) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTPStatementLoop277) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTPStatementLoop277
}

type NullableBTPStatementLoop277 struct {
	value *BTPStatementLoop277
	isSet bool
}

func (v NullableBTPStatementLoop277) Get() *BTPStatementLoop277 {
	return v.value
}

func (v *NullableBTPStatementLoop277) Set(val *BTPStatementLoop277) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPStatementLoop277) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPStatementLoop277) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPStatementLoop277(val *BTPStatementLoop277) *NullableBTPStatementLoop277 {
	return &NullableBTPStatementLoop277{value: val, isSet: true}
}

func (v NullableBTPStatementLoop277) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPStatementLoop277) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTPStatementLoop277 struct {
	Annotation          *BTPAnnotation231 `json:"annotation,omitempty"`
	Atomic              *bool             `json:"atomic,omitempty"`
	BtType              *string           `json:"btType,omitempty"`
	DocumentationType   *string           `json:"documentationType,omitempty"`
	EndSourceLocation   *int32            `json:"endSourceLocation,omitempty"`
	NodeId              *string           `json:"nodeId,omitempty"`
	ShortDescriptor     *string           `json:"shortDescriptor,omitempty"`
	SpaceAfter          *BTPSpace10       `json:"spaceAfter,omitempty"`
	SpaceBefore         *BTPSpace10       `json:"spaceBefore,omitempty"`
	SpaceDefault        *bool             `json:"spaceDefault,omitempty"`
	StartSourceLocation *int32            `json:"startSourceLocation,omitempty"`
	Body                *BTPStatement269  `json:"body,omitempty"`
	SpaceAfterLoopType  *BTPSpace10       `json:"spaceAfterLoopType,omitempty"`
}

// Newbase_BTPStatementLoop277 instantiates a new base_BTPStatementLoop277 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTPStatementLoop277() *base_BTPStatementLoop277 {
	this := base_BTPStatementLoop277{}
	return &this
}

// Newbase_BTPStatementLoop277WithDefaults instantiates a new base_BTPStatementLoop277 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTPStatementLoop277WithDefaults() *base_BTPStatementLoop277 {
	this := base_BTPStatementLoop277{}
	return &this
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *base_BTPStatementLoop277) GetAnnotation() BTPAnnotation231 {
	if o == nil || o.Annotation == nil {
		var ret BTPAnnotation231
		return ret
	}
	return *o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTPStatementLoop277) GetAnnotationOk() (*BTPAnnotation231, bool) {
	if o == nil || o.Annotation == nil {
		return nil, false
	}
	return o.Annotation, true
}

// HasAnnotation returns a boolean if a field has been set.
func (o *base_BTPStatementLoop277) HasAnnotation() bool {
	if o != nil && o.Annotation != nil {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given BTPAnnotation231 and assigns it to the Annotation field.
func (o *base_BTPStatementLoop277) SetAnnotation(v BTPAnnotation231) {
	o.Annotation = &v
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *base_BTPStatementLoop277) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTPStatementLoop277) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *base_BTPStatementLoop277) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *base_BTPStatementLoop277) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTPStatementLoop277) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTPStatementLoop277) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTPStatementLoop277) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTPStatementLoop277) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *base_BTPStatementLoop277) GetDocumentationType() string {
	if o == nil || o.DocumentationType == nil {
		var ret string
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTPStatementLoop277) GetDocumentationTypeOk() (*string, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *base_BTPStatementLoop277) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given string and assigns it to the DocumentationType field.
func (o *base_BTPStatementLoop277) SetDocumentationType(v string) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *base_BTPStatementLoop277) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTPStatementLoop277) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *base_BTPStatementLoop277) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *base_BTPStatementLoop277) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTPStatementLoop277) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTPStatementLoop277) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTPStatementLoop277) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTPStatementLoop277) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *base_BTPStatementLoop277) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTPStatementLoop277) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *base_BTPStatementLoop277) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *base_BTPStatementLoop277) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *base_BTPStatementLoop277) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTPStatementLoop277) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *base_BTPStatementLoop277) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *base_BTPStatementLoop277) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *base_BTPStatementLoop277) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTPStatementLoop277) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *base_BTPStatementLoop277) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *base_BTPStatementLoop277) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *base_BTPStatementLoop277) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTPStatementLoop277) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *base_BTPStatementLoop277) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *base_BTPStatementLoop277) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *base_BTPStatementLoop277) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTPStatementLoop277) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *base_BTPStatementLoop277) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *base_BTPStatementLoop277) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *base_BTPStatementLoop277) GetBody() BTPStatement269 {
	if o == nil || o.Body == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTPStatementLoop277) GetBodyOk() (*BTPStatement269, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *base_BTPStatementLoop277) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given BTPStatement269 and assigns it to the Body field.
func (o *base_BTPStatementLoop277) SetBody(v BTPStatement269) {
	o.Body = &v
}

// GetSpaceAfterLoopType returns the SpaceAfterLoopType field value if set, zero value otherwise.
func (o *base_BTPStatementLoop277) GetSpaceAfterLoopType() BTPSpace10 {
	if o == nil || o.SpaceAfterLoopType == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterLoopType
}

// GetSpaceAfterLoopTypeOk returns a tuple with the SpaceAfterLoopType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTPStatementLoop277) GetSpaceAfterLoopTypeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterLoopType == nil {
		return nil, false
	}
	return o.SpaceAfterLoopType, true
}

// HasSpaceAfterLoopType returns a boolean if a field has been set.
func (o *base_BTPStatementLoop277) HasSpaceAfterLoopType() bool {
	if o != nil && o.SpaceAfterLoopType != nil {
		return true
	}

	return false
}

// SetSpaceAfterLoopType gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterLoopType field.
func (o *base_BTPStatementLoop277) SetSpaceAfterLoopType(v BTPSpace10) {
	o.SpaceAfterLoopType = &v
}

func (o base_BTPStatementLoop277) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}
