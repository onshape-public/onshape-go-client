/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13883-0fccf6f231ad
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTSingleReferenceDisplayData1943 - struct for BTSingleReferenceDisplayData1943
type BTSingleReferenceDisplayData1943 struct {
	implBTSingleReferenceDisplayData1943 interface{}
}

// BTSingleAssemblyReferenceDisplayData1557AsBTSingleReferenceDisplayData1943 is a convenience function that returns BTSingleAssemblyReferenceDisplayData1557 wrapped in BTSingleReferenceDisplayData1943
func (o *BTSingleAssemblyReferenceDisplayData1557) AsBTSingleReferenceDisplayData1943() *BTSingleReferenceDisplayData1943 {
	return &BTSingleReferenceDisplayData1943{o}
}

// NewBTSingleReferenceDisplayData1943 instantiates a new BTSingleReferenceDisplayData1943 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSingleReferenceDisplayData1943() *BTSingleReferenceDisplayData1943 {
	this := BTSingleReferenceDisplayData1943{Newbase_BTSingleReferenceDisplayData1943()}
	return &this
}

// NewBTSingleReferenceDisplayData1943WithDefaults instantiates a new BTSingleReferenceDisplayData1943 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSingleReferenceDisplayData1943WithDefaults() *BTSingleReferenceDisplayData1943 {
	this := BTSingleReferenceDisplayData1943{Newbase_BTSingleReferenceDisplayData1943WithDefaults()}
	return &this
}

// GetContextWorkspaceId returns the ContextWorkspaceId field value if set, zero value otherwise.
func (o *BTSingleReferenceDisplayData1943) GetContextWorkspaceId() string {
	type getResult interface {
		GetContextWorkspaceId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetContextWorkspaceId()
	} else {
		var de string
		return de
	}
}

// GetContextWorkspaceIdOk returns a tuple with the ContextWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleReferenceDisplayData1943) GetContextWorkspaceIdOk() (*string, bool) {
	type getResult interface {
		GetContextWorkspaceIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetContextWorkspaceIdOk()
	} else {
		return nil, false
	}
}

// HasContextWorkspaceId returns a boolean if a field has been set.
func (o *BTSingleReferenceDisplayData1943) HasContextWorkspaceId() bool {
	type getResult interface {
		HasContextWorkspaceId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasContextWorkspaceId()
	} else {
		return false
	}
}

// SetContextWorkspaceId gets a reference to the given string and assigns it to the ContextWorkspaceId field.
func (o *BTSingleReferenceDisplayData1943) SetContextWorkspaceId(v string) {
	type getResult interface {
		SetContextWorkspaceId(v string)
	}

	o.GetActualInstance().(getResult).SetContextWorkspaceId(v)
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTSingleReferenceDisplayData1943) GetDocumentId() string {
	type getResult interface {
		GetDocumentId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDocumentId()
	} else {
		var de string
		return de
	}
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleReferenceDisplayData1943) GetDocumentIdOk() (*string, bool) {
	type getResult interface {
		GetDocumentIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDocumentIdOk()
	} else {
		return nil, false
	}
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTSingleReferenceDisplayData1943) HasDocumentId() bool {
	type getResult interface {
		HasDocumentId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDocumentId()
	} else {
		return false
	}
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTSingleReferenceDisplayData1943) SetDocumentId(v string) {
	type getResult interface {
		SetDocumentId(v string)
	}

	o.GetActualInstance().(getResult).SetDocumentId(v)
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BTSingleReferenceDisplayData1943) GetError() string {
	type getResult interface {
		GetError() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetError()
	} else {
		var de string
		return de
	}
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleReferenceDisplayData1943) GetErrorOk() (*string, bool) {
	type getResult interface {
		GetErrorOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetErrorOk()
	} else {
		return nil, false
	}
}

// HasError returns a boolean if a field has been set.
func (o *BTSingleReferenceDisplayData1943) HasError() bool {
	type getResult interface {
		HasError() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasError()
	} else {
		return false
	}
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *BTSingleReferenceDisplayData1943) SetError(v string) {
	type getResult interface {
		SetError(v string)
	}

	o.GetActualInstance().(getResult).SetError(v)
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *BTSingleReferenceDisplayData1943) GetErrorMessage() string {
	type getResult interface {
		GetErrorMessage() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetErrorMessage()
	} else {
		var de string
		return de
	}
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleReferenceDisplayData1943) GetErrorMessageOk() (*string, bool) {
	type getResult interface {
		GetErrorMessageOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetErrorMessageOk()
	} else {
		return nil, false
	}
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *BTSingleReferenceDisplayData1943) HasErrorMessage() bool {
	type getResult interface {
		HasErrorMessage() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasErrorMessage()
	} else {
		return false
	}
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *BTSingleReferenceDisplayData1943) SetErrorMessage(v string) {
	type getResult interface {
		SetErrorMessage(v string)
	}

	o.GetActualInstance().(getResult).SetErrorMessage(v)
}

// GetIsTransient returns the IsTransient field value if set, zero value otherwise.
func (o *BTSingleReferenceDisplayData1943) GetIsTransient() bool {
	type getResult interface {
		GetIsTransient() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsTransient()
	} else {
		var de bool
		return de
	}
}

// GetIsTransientOk returns a tuple with the IsTransient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleReferenceDisplayData1943) GetIsTransientOk() (*bool, bool) {
	type getResult interface {
		GetIsTransientOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsTransientOk()
	} else {
		return nil, false
	}
}

// HasIsTransient returns a boolean if a field has been set.
func (o *BTSingleReferenceDisplayData1943) HasIsTransient() bool {
	type getResult interface {
		HasIsTransient() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasIsTransient()
	} else {
		return false
	}
}

// SetIsTransient gets a reference to the given bool and assigns it to the IsTransient field.
func (o *BTSingleReferenceDisplayData1943) SetIsTransient(v bool) {
	type getResult interface {
		SetIsTransient(v bool)
	}

	o.GetActualInstance().(getResult).SetIsTransient(v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTSingleReferenceDisplayData1943) GetName() string {
	type getResult interface {
		GetName() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetName()
	} else {
		var de string
		return de
	}
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleReferenceDisplayData1943) GetNameOk() (*string, bool) {
	type getResult interface {
		GetNameOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNameOk()
	} else {
		return nil, false
	}
}

// HasName returns a boolean if a field has been set.
func (o *BTSingleReferenceDisplayData1943) HasName() bool {
	type getResult interface {
		HasName() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasName()
	} else {
		return false
	}
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTSingleReferenceDisplayData1943) SetName(v string) {
	type getResult interface {
		SetName(v string)
	}

	o.GetActualInstance().(getResult).SetName(v)
}

// GetReferenceName returns the ReferenceName field value if set, zero value otherwise.
func (o *BTSingleReferenceDisplayData1943) GetReferenceName() string {
	type getResult interface {
		GetReferenceName() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetReferenceName()
	} else {
		var de string
		return de
	}
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleReferenceDisplayData1943) GetReferenceNameOk() (*string, bool) {
	type getResult interface {
		GetReferenceNameOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetReferenceNameOk()
	} else {
		return nil, false
	}
}

// HasReferenceName returns a boolean if a field has been set.
func (o *BTSingleReferenceDisplayData1943) HasReferenceName() bool {
	type getResult interface {
		HasReferenceName() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasReferenceName()
	} else {
		return false
	}
}

// SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.
func (o *BTSingleReferenceDisplayData1943) SetReferenceName(v string) {
	type getResult interface {
		SetReferenceName(v string)
	}

	o.GetActualInstance().(getResult).SetReferenceName(v)
}

// GetReferenceNodeId returns the ReferenceNodeId field value if set, zero value otherwise.
func (o *BTSingleReferenceDisplayData1943) GetReferenceNodeId() string {
	type getResult interface {
		GetReferenceNodeId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetReferenceNodeId()
	} else {
		var de string
		return de
	}
}

// GetReferenceNodeIdOk returns a tuple with the ReferenceNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleReferenceDisplayData1943) GetReferenceNodeIdOk() (*string, bool) {
	type getResult interface {
		GetReferenceNodeIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetReferenceNodeIdOk()
	} else {
		return nil, false
	}
}

// HasReferenceNodeId returns a boolean if a field has been set.
func (o *BTSingleReferenceDisplayData1943) HasReferenceNodeId() bool {
	type getResult interface {
		HasReferenceNodeId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasReferenceNodeId()
	} else {
		return false
	}
}

// SetReferenceNodeId gets a reference to the given string and assigns it to the ReferenceNodeId field.
func (o *BTSingleReferenceDisplayData1943) SetReferenceNodeId(v string) {
	type getResult interface {
		SetReferenceNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetReferenceNodeId(v)
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *BTSingleReferenceDisplayData1943) GetVisibility() string {
	type getResult interface {
		GetVisibility() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetVisibility()
	} else {
		var de string
		return de
	}
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleReferenceDisplayData1943) GetVisibilityOk() (*string, bool) {
	type getResult interface {
		GetVisibilityOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetVisibilityOk()
	} else {
		return nil, false
	}
}

// HasVisibility returns a boolean if a field has been set.
func (o *BTSingleReferenceDisplayData1943) HasVisibility() bool {
	type getResult interface {
		HasVisibility() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasVisibility()
	} else {
		return false
	}
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *BTSingleReferenceDisplayData1943) SetVisibility(v string) {
	type getResult interface {
		SetVisibility(v string)
	}

	o.GetActualInstance().(getResult).SetVisibility(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTSingleReferenceDisplayData1943) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTSingleAssemblyReferenceDisplayData-1557'
	if jsonDict["btType"] == "BTSingleAssemblyReferenceDisplayData-1557" {
		// try to unmarshal JSON data into BTSingleAssemblyReferenceDisplayData1557
		var qr *BTSingleAssemblyReferenceDisplayData1557
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSingleReferenceDisplayData1943 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSingleReferenceDisplayData1943 = nil
			return fmt.Errorf("Failed to unmarshal BTSingleReferenceDisplayData1943 as BTSingleAssemblyReferenceDisplayData1557: %s", err.Error())
		}
	}

	var qtx *base_BTSingleReferenceDisplayData1943
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTSingleReferenceDisplayData1943 = qtx
		return nil // data stored in dst.base_BTSingleReferenceDisplayData1943, return on the first match
	} else {
		dst.implBTSingleReferenceDisplayData1943 = nil
		return fmt.Errorf("Failed to unmarshal BTSingleReferenceDisplayData1943 as base_BTSingleReferenceDisplayData1943: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTSingleReferenceDisplayData1943) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTSingleReferenceDisplayData1943) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTSingleReferenceDisplayData1943
}

type NullableBTSingleReferenceDisplayData1943 struct {
	value *BTSingleReferenceDisplayData1943
	isSet bool
}

func (v NullableBTSingleReferenceDisplayData1943) Get() *BTSingleReferenceDisplayData1943 {
	return v.value
}

func (v *NullableBTSingleReferenceDisplayData1943) Set(val *BTSingleReferenceDisplayData1943) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSingleReferenceDisplayData1943) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSingleReferenceDisplayData1943) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSingleReferenceDisplayData1943(val *BTSingleReferenceDisplayData1943) *NullableBTSingleReferenceDisplayData1943 {
	return &NullableBTSingleReferenceDisplayData1943{value: val, isSet: true}
}

func (v NullableBTSingleReferenceDisplayData1943) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSingleReferenceDisplayData1943) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTSingleReferenceDisplayData1943 struct {
	ContextWorkspaceId *string `json:"contextWorkspaceId,omitempty"`
	DocumentId         *string `json:"documentId,omitempty"`
	Error              *string `json:"error,omitempty"`
	ErrorMessage       *string `json:"errorMessage,omitempty"`
	IsTransient        *bool   `json:"isTransient,omitempty"`
	Name               *string `json:"name,omitempty"`
	ReferenceName      *string `json:"referenceName,omitempty"`
	ReferenceNodeId    *string `json:"referenceNodeId,omitempty"`
	Visibility         *string `json:"visibility,omitempty"`
}

// Newbase_BTSingleReferenceDisplayData1943 instantiates a new base_BTSingleReferenceDisplayData1943 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTSingleReferenceDisplayData1943() *base_BTSingleReferenceDisplayData1943 {
	this := base_BTSingleReferenceDisplayData1943{}
	return &this
}

// Newbase_BTSingleReferenceDisplayData1943WithDefaults instantiates a new base_BTSingleReferenceDisplayData1943 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTSingleReferenceDisplayData1943WithDefaults() *base_BTSingleReferenceDisplayData1943 {
	this := base_BTSingleReferenceDisplayData1943{}
	return &this
}

// GetContextWorkspaceId returns the ContextWorkspaceId field value if set, zero value otherwise.
func (o *base_BTSingleReferenceDisplayData1943) GetContextWorkspaceId() string {
	if o == nil || o.ContextWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.ContextWorkspaceId
}

// GetContextWorkspaceIdOk returns a tuple with the ContextWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSingleReferenceDisplayData1943) GetContextWorkspaceIdOk() (*string, bool) {
	if o == nil || o.ContextWorkspaceId == nil {
		return nil, false
	}
	return o.ContextWorkspaceId, true
}

// HasContextWorkspaceId returns a boolean if a field has been set.
func (o *base_BTSingleReferenceDisplayData1943) HasContextWorkspaceId() bool {
	if o != nil && o.ContextWorkspaceId != nil {
		return true
	}

	return false
}

// SetContextWorkspaceId gets a reference to the given string and assigns it to the ContextWorkspaceId field.
func (o *base_BTSingleReferenceDisplayData1943) SetContextWorkspaceId(v string) {
	o.ContextWorkspaceId = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *base_BTSingleReferenceDisplayData1943) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSingleReferenceDisplayData1943) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *base_BTSingleReferenceDisplayData1943) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *base_BTSingleReferenceDisplayData1943) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *base_BTSingleReferenceDisplayData1943) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSingleReferenceDisplayData1943) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *base_BTSingleReferenceDisplayData1943) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *base_BTSingleReferenceDisplayData1943) SetError(v string) {
	o.Error = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *base_BTSingleReferenceDisplayData1943) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSingleReferenceDisplayData1943) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *base_BTSingleReferenceDisplayData1943) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *base_BTSingleReferenceDisplayData1943) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetIsTransient returns the IsTransient field value if set, zero value otherwise.
func (o *base_BTSingleReferenceDisplayData1943) GetIsTransient() bool {
	if o == nil || o.IsTransient == nil {
		var ret bool
		return ret
	}
	return *o.IsTransient
}

// GetIsTransientOk returns a tuple with the IsTransient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSingleReferenceDisplayData1943) GetIsTransientOk() (*bool, bool) {
	if o == nil || o.IsTransient == nil {
		return nil, false
	}
	return o.IsTransient, true
}

// HasIsTransient returns a boolean if a field has been set.
func (o *base_BTSingleReferenceDisplayData1943) HasIsTransient() bool {
	if o != nil && o.IsTransient != nil {
		return true
	}

	return false
}

// SetIsTransient gets a reference to the given bool and assigns it to the IsTransient field.
func (o *base_BTSingleReferenceDisplayData1943) SetIsTransient(v bool) {
	o.IsTransient = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *base_BTSingleReferenceDisplayData1943) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSingleReferenceDisplayData1943) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *base_BTSingleReferenceDisplayData1943) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *base_BTSingleReferenceDisplayData1943) SetName(v string) {
	o.Name = &v
}

// GetReferenceName returns the ReferenceName field value if set, zero value otherwise.
func (o *base_BTSingleReferenceDisplayData1943) GetReferenceName() string {
	if o == nil || o.ReferenceName == nil {
		var ret string
		return ret
	}
	return *o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSingleReferenceDisplayData1943) GetReferenceNameOk() (*string, bool) {
	if o == nil || o.ReferenceName == nil {
		return nil, false
	}
	return o.ReferenceName, true
}

// HasReferenceName returns a boolean if a field has been set.
func (o *base_BTSingleReferenceDisplayData1943) HasReferenceName() bool {
	if o != nil && o.ReferenceName != nil {
		return true
	}

	return false
}

// SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.
func (o *base_BTSingleReferenceDisplayData1943) SetReferenceName(v string) {
	o.ReferenceName = &v
}

// GetReferenceNodeId returns the ReferenceNodeId field value if set, zero value otherwise.
func (o *base_BTSingleReferenceDisplayData1943) GetReferenceNodeId() string {
	if o == nil || o.ReferenceNodeId == nil {
		var ret string
		return ret
	}
	return *o.ReferenceNodeId
}

// GetReferenceNodeIdOk returns a tuple with the ReferenceNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSingleReferenceDisplayData1943) GetReferenceNodeIdOk() (*string, bool) {
	if o == nil || o.ReferenceNodeId == nil {
		return nil, false
	}
	return o.ReferenceNodeId, true
}

// HasReferenceNodeId returns a boolean if a field has been set.
func (o *base_BTSingleReferenceDisplayData1943) HasReferenceNodeId() bool {
	if o != nil && o.ReferenceNodeId != nil {
		return true
	}

	return false
}

// SetReferenceNodeId gets a reference to the given string and assigns it to the ReferenceNodeId field.
func (o *base_BTSingleReferenceDisplayData1943) SetReferenceNodeId(v string) {
	o.ReferenceNodeId = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *base_BTSingleReferenceDisplayData1943) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSingleReferenceDisplayData1943) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *base_BTSingleReferenceDisplayData1943) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *base_BTSingleReferenceDisplayData1943) SetVisibility(v string) {
	o.Visibility = &v
}

func (o base_BTSingleReferenceDisplayData1943) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContextWorkspaceId != nil {
		toSerialize["contextWorkspaceId"] = o.ContextWorkspaceId
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.IsTransient != nil {
		toSerialize["isTransient"] = o.IsTransient
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ReferenceName != nil {
		toSerialize["referenceName"] = o.ReferenceName
	}
	if o.ReferenceNodeId != nil {
		toSerialize["referenceNodeId"] = o.ReferenceNodeId
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}
