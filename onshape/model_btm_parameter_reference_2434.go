/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTMParameterReference2434 - struct for BTMParameterReference2434
type BTMParameterReference2434 struct {
	implBTMParameterReference2434 interface{}
}

// BTMParameterReferenceWithConfiguration3028AsBTMParameterReference2434 is a convenience function that returns BTMParameterReferenceWithConfiguration3028 wrapped in BTMParameterReference2434
func (o *BTMParameterReferenceWithConfiguration3028) AsBTMParameterReference2434() *BTMParameterReference2434 {
	return &BTMParameterReference2434{o}
}

// BTMParameterReferenceCADImport2016AsBTMParameterReference2434 is a convenience function that returns BTMParameterReferenceCADImport2016 wrapped in BTMParameterReference2434
func (o *BTMParameterReferenceCADImport2016) AsBTMParameterReference2434() *BTMParameterReference2434 {
	return &BTMParameterReference2434{o}
}

// BTMParameterReferenceTable917AsBTMParameterReference2434 is a convenience function that returns BTMParameterReferenceTable917 wrapped in BTMParameterReference2434
func (o *BTMParameterReferenceTable917) AsBTMParameterReference2434() *BTMParameterReference2434 {
	return &BTMParameterReference2434{o}
}

// BTMParameterReferenceBlob3281AsBTMParameterReference2434 is a convenience function that returns BTMParameterReferenceBlob3281 wrapped in BTMParameterReference2434
func (o *BTMParameterReferenceBlob3281) AsBTMParameterReference2434() *BTMParameterReference2434 {
	return &BTMParameterReference2434{o}
}

// BTMParameterReferenceImage2014AsBTMParameterReference2434 is a convenience function that returns BTMParameterReferenceImage2014 wrapped in BTMParameterReference2434
func (o *BTMParameterReferenceImage2014) AsBTMParameterReference2434() *BTMParameterReference2434 {
	return &BTMParameterReference2434{o}
}

// NewBTMParameterReference2434 instantiates a new BTMParameterReference2434 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterReference2434() *BTMParameterReference2434 {
	this := BTMParameterReference2434{Newbase_BTMParameterReference2434()}
	return &this
}

// NewBTMParameterReference2434WithDefaults instantiates a new BTMParameterReference2434 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterReference2434WithDefaults() *BTMParameterReference2434 {
	this := BTMParameterReference2434{Newbase_BTMParameterReference2434WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterReference2434) GetBtType() string {
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
func (o *BTMParameterReference2434) GetBtTypeOk() (*string, bool) {
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
func (o *BTMParameterReference2434) HasBtType() bool {
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
func (o *BTMParameterReference2434) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterReference2434) GetImportMicroversion() string {
	type getResult interface {
		GetImportMicroversion() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversion()
	} else {
		var de string
		return de
	}
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReference2434) GetImportMicroversionOk() (*string, bool) {
	type getResult interface {
		GetImportMicroversionOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversionOk()
	} else {
		return nil, false
	}
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterReference2434) HasImportMicroversion() bool {
	type getResult interface {
		HasImportMicroversion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasImportMicroversion()
	} else {
		return false
	}
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterReference2434) SetImportMicroversion(v string) {
	type getResult interface {
		SetImportMicroversion(v string)
	}

	o.GetActualInstance().(getResult).SetImportMicroversion(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterReference2434) GetNodeId() string {
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
func (o *BTMParameterReference2434) GetNodeIdOk() (*string, bool) {
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
func (o *BTMParameterReference2434) HasNodeId() bool {
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
func (o *BTMParameterReference2434) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterReference2434) GetParameterId() string {
	type getResult interface {
		GetParameterId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameterId()
	} else {
		var de string
		return de
	}
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReference2434) GetParameterIdOk() (*string, bool) {
	type getResult interface {
		GetParameterIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameterIdOk()
	} else {
		return nil, false
	}
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterReference2434) HasParameterId() bool {
	type getResult interface {
		HasParameterId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasParameterId()
	} else {
		return false
	}
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterReference2434) SetParameterId(v string) {
	type getResult interface {
		SetParameterId(v string)
	}

	o.GetActualInstance().(getResult).SetParameterId(v)
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTMParameterReference2434) GetDocumentId() string {
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
func (o *BTMParameterReference2434) GetDocumentIdOk() (*string, bool) {
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
func (o *BTMParameterReference2434) HasDocumentId() bool {
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
func (o *BTMParameterReference2434) SetDocumentId(v string) {
	type getResult interface {
		SetDocumentId(v string)
	}

	o.GetActualInstance().(getResult).SetDocumentId(v)
}

// GetDocumentVersionId returns the DocumentVersionId field value if set, zero value otherwise.
func (o *BTMParameterReference2434) GetDocumentVersionId() string {
	type getResult interface {
		GetDocumentVersionId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDocumentVersionId()
	} else {
		var de string
		return de
	}
}

// GetDocumentVersionIdOk returns a tuple with the DocumentVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReference2434) GetDocumentVersionIdOk() (*string, bool) {
	type getResult interface {
		GetDocumentVersionIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDocumentVersionIdOk()
	} else {
		return nil, false
	}
}

// HasDocumentVersionId returns a boolean if a field has been set.
func (o *BTMParameterReference2434) HasDocumentVersionId() bool {
	type getResult interface {
		HasDocumentVersionId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDocumentVersionId()
	} else {
		return false
	}
}

// SetDocumentVersionId gets a reference to the given string and assigns it to the DocumentVersionId field.
func (o *BTMParameterReference2434) SetDocumentVersionId(v string) {
	type getResult interface {
		SetDocumentVersionId(v string)
	}

	o.GetActualInstance().(getResult).SetDocumentVersionId(v)
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTMParameterReference2434) GetElementId() string {
	type getResult interface {
		GetElementId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetElementId()
	} else {
		var de string
		return de
	}
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReference2434) GetElementIdOk() (*string, bool) {
	type getResult interface {
		GetElementIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetElementIdOk()
	} else {
		return nil, false
	}
}

// HasElementId returns a boolean if a field has been set.
func (o *BTMParameterReference2434) HasElementId() bool {
	type getResult interface {
		HasElementId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasElementId()
	} else {
		return false
	}
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTMParameterReference2434) SetElementId(v string) {
	type getResult interface {
		SetElementId(v string)
	}

	o.GetActualInstance().(getResult).SetElementId(v)
}

// GetElementLibraryData returns the ElementLibraryData field value if set, zero value otherwise.
func (o *BTMParameterReference2434) GetElementLibraryData() BTElementLibraryReferenceData3133 {
	type getResult interface {
		GetElementLibraryData() BTElementLibraryReferenceData3133
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetElementLibraryData()
	} else {
		var de BTElementLibraryReferenceData3133
		return de
	}
}

// GetElementLibraryDataOk returns a tuple with the ElementLibraryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReference2434) GetElementLibraryDataOk() (*BTElementLibraryReferenceData3133, bool) {
	type getResult interface {
		GetElementLibraryDataOk() (*BTElementLibraryReferenceData3133, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetElementLibraryDataOk()
	} else {
		return nil, false
	}
}

// HasElementLibraryData returns a boolean if a field has been set.
func (o *BTMParameterReference2434) HasElementLibraryData() bool {
	type getResult interface {
		HasElementLibraryData() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasElementLibraryData()
	} else {
		return false
	}
}

// SetElementLibraryData gets a reference to the given BTElementLibraryReferenceData3133 and assigns it to the ElementLibraryData field.
func (o *BTMParameterReference2434) SetElementLibraryData(v BTElementLibraryReferenceData3133) {
	type getResult interface {
		SetElementLibraryData(v BTElementLibraryReferenceData3133)
	}

	o.GetActualInstance().(getResult).SetElementLibraryData(v)
}

// GetFeatureScriptType returns the FeatureScriptType field value if set, zero value otherwise.
func (o *BTMParameterReference2434) GetFeatureScriptType() string {
	type getResult interface {
		GetFeatureScriptType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFeatureScriptType()
	} else {
		var de string
		return de
	}
}

// GetFeatureScriptTypeOk returns a tuple with the FeatureScriptType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReference2434) GetFeatureScriptTypeOk() (*string, bool) {
	type getResult interface {
		GetFeatureScriptTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFeatureScriptTypeOk()
	} else {
		return nil, false
	}
}

// HasFeatureScriptType returns a boolean if a field has been set.
func (o *BTMParameterReference2434) HasFeatureScriptType() bool {
	type getResult interface {
		HasFeatureScriptType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasFeatureScriptType()
	} else {
		return false
	}
}

// SetFeatureScriptType gets a reference to the given string and assigns it to the FeatureScriptType field.
func (o *BTMParameterReference2434) SetFeatureScriptType(v string) {
	type getResult interface {
		SetFeatureScriptType(v string)
	}

	o.GetActualInstance().(getResult).SetFeatureScriptType(v)
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *BTMParameterReference2434) GetIds() []string {
	type getResult interface {
		GetIds() []string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIds()
	} else {
		var de []string
		return de
	}
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReference2434) GetIdsOk() ([]string, bool) {
	type getResult interface {
		GetIdsOk() ([]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIdsOk()
	} else {
		return nil, false
	}
}

// HasIds returns a boolean if a field has been set.
func (o *BTMParameterReference2434) HasIds() bool {
	type getResult interface {
		HasIds() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasIds()
	} else {
		return false
	}
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *BTMParameterReference2434) SetIds(v []string) {
	type getResult interface {
		SetIds(v []string)
	}

	o.GetActualInstance().(getResult).SetIds(v)
}

// GetMicroversioId returns the MicroversioId field value if set, zero value otherwise.
func (o *BTMParameterReference2434) GetMicroversioId() string {
	type getResult interface {
		GetMicroversioId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversioId()
	} else {
		var de string
		return de
	}
}

// GetMicroversioIdOk returns a tuple with the MicroversioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReference2434) GetMicroversioIdOk() (*string, bool) {
	type getResult interface {
		GetMicroversioIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversioIdOk()
	} else {
		return nil, false
	}
}

// HasMicroversioId returns a boolean if a field has been set.
func (o *BTMParameterReference2434) HasMicroversioId() bool {
	type getResult interface {
		HasMicroversioId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMicroversioId()
	} else {
		return false
	}
}

// SetMicroversioId gets a reference to the given string and assigns it to the MicroversioId field.
func (o *BTMParameterReference2434) SetMicroversioId(v string) {
	type getResult interface {
		SetMicroversioId(v string)
	}

	o.GetActualInstance().(getResult).SetMicroversioId(v)
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMParameterReference2434) GetNamespace() string {
	type getResult interface {
		GetNamespace() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNamespace()
	} else {
		var de string
		return de
	}
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReference2434) GetNamespaceOk() (*string, bool) {
	type getResult interface {
		GetNamespaceOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNamespaceOk()
	} else {
		return nil, false
	}
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMParameterReference2434) HasNamespace() bool {
	type getResult interface {
		HasNamespace() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNamespace()
	} else {
		return false
	}
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMParameterReference2434) SetNamespace(v string) {
	type getResult interface {
		SetNamespace(v string)
	}

	o.GetActualInstance().(getResult).SetNamespace(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMParameterReference2434) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BTMParameterReferenceBlob-3281'
	if jsonDict["btType"] == "BTMParameterReferenceBlob-3281" {
		// try to unmarshal JSON data into BTMParameterReferenceBlob3281
		var qr *BTMParameterReferenceBlob3281
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameterReference2434 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameterReference2434 = nil
			return fmt.Errorf("failed to unmarshal BTMParameterReference2434 as BTMParameterReferenceBlob3281: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterReferenceCADImport-2016'
	if jsonDict["btType"] == "BTMParameterReferenceCADImport-2016" {
		// try to unmarshal JSON data into BTMParameterReferenceCADImport2016
		var qr *BTMParameterReferenceCADImport2016
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameterReference2434 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameterReference2434 = nil
			return fmt.Errorf("failed to unmarshal BTMParameterReference2434 as BTMParameterReferenceCADImport2016: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterReferenceImage-2014'
	if jsonDict["btType"] == "BTMParameterReferenceImage-2014" {
		// try to unmarshal JSON data into BTMParameterReferenceImage2014
		var qr *BTMParameterReferenceImage2014
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameterReference2434 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameterReference2434 = nil
			return fmt.Errorf("failed to unmarshal BTMParameterReference2434 as BTMParameterReferenceImage2014: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterReferenceTable-917'
	if jsonDict["btType"] == "BTMParameterReferenceTable-917" {
		// try to unmarshal JSON data into BTMParameterReferenceTable917
		var qr *BTMParameterReferenceTable917
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameterReference2434 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameterReference2434 = nil
			return fmt.Errorf("failed to unmarshal BTMParameterReference2434 as BTMParameterReferenceTable917: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterReferenceWithConfiguration-3028'
	if jsonDict["btType"] == "BTMParameterReferenceWithConfiguration-3028" {
		// try to unmarshal JSON data into BTMParameterReferenceWithConfiguration3028
		var qr *BTMParameterReferenceWithConfiguration3028
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameterReference2434 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameterReference2434 = nil
			return fmt.Errorf("failed to unmarshal BTMParameterReference2434 as BTMParameterReferenceWithConfiguration3028: %s", err.Error())
		}
	}

	var qtx *base_BTMParameterReference2434
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMParameterReference2434 = qtx
		return nil // data stored in dst.base_BTMParameterReference2434, return on the first match
	} else {
		dst.implBTMParameterReference2434 = nil
		return fmt.Errorf("failed to unmarshal BTMParameterReference2434 as base_BTMParameterReference2434: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMParameterReference2434) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMParameterReference2434) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMParameterReference2434
}

type NullableBTMParameterReference2434 struct {
	value *BTMParameterReference2434
	isSet bool
}

func (v NullableBTMParameterReference2434) Get() *BTMParameterReference2434 {
	return v.value
}

func (v *NullableBTMParameterReference2434) Set(val *BTMParameterReference2434) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterReference2434) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterReference2434) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterReference2434(val *BTMParameterReference2434) *NullableBTMParameterReference2434 {
	return &NullableBTMParameterReference2434{value: val, isSet: true}
}

func (v NullableBTMParameterReference2434) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterReference2434) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMParameterReference2434 struct {
	BtType *string `json:"btType,omitempty"`
	// Element microversion that is being imported.
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	// ID of the parameter's node.
	NodeId *string `json:"nodeId,omitempty"`
	// Unique ID of the parameter.
	ParameterId        *string                            `json:"parameterId,omitempty"`
	DocumentId         *string                            `json:"documentId,omitempty"`
	DocumentVersionId  *string                            `json:"documentVersionId,omitempty"`
	ElementId          *string                            `json:"elementId,omitempty"`
	ElementLibraryData *BTElementLibraryReferenceData3133 `json:"elementLibraryData,omitempty"`
	FeatureScriptType  *string                            `json:"featureScriptType,omitempty"`
	Ids                []string                           `json:"ids,omitempty"`
	MicroversioId      *string                            `json:"microversioId,omitempty"`
	Namespace          *string                            `json:"namespace,omitempty"`
}

// Newbase_BTMParameterReference2434 instantiates a new base_BTMParameterReference2434 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMParameterReference2434() *base_BTMParameterReference2434 {
	this := base_BTMParameterReference2434{}
	return &this
}

// Newbase_BTMParameterReference2434WithDefaults instantiates a new base_BTMParameterReference2434 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMParameterReference2434WithDefaults() *base_BTMParameterReference2434 {
	this := base_BTMParameterReference2434{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTMParameterReference2434) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterReference2434) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTMParameterReference2434) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTMParameterReference2434) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *base_BTMParameterReference2434) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterReference2434) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *base_BTMParameterReference2434) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *base_BTMParameterReference2434) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTMParameterReference2434) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterReference2434) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTMParameterReference2434) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTMParameterReference2434) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *base_BTMParameterReference2434) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterReference2434) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *base_BTMParameterReference2434) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *base_BTMParameterReference2434) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *base_BTMParameterReference2434) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterReference2434) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *base_BTMParameterReference2434) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *base_BTMParameterReference2434) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentVersionId returns the DocumentVersionId field value if set, zero value otherwise.
func (o *base_BTMParameterReference2434) GetDocumentVersionId() string {
	if o == nil || o.DocumentVersionId == nil {
		var ret string
		return ret
	}
	return *o.DocumentVersionId
}

// GetDocumentVersionIdOk returns a tuple with the DocumentVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterReference2434) GetDocumentVersionIdOk() (*string, bool) {
	if o == nil || o.DocumentVersionId == nil {
		return nil, false
	}
	return o.DocumentVersionId, true
}

// HasDocumentVersionId returns a boolean if a field has been set.
func (o *base_BTMParameterReference2434) HasDocumentVersionId() bool {
	if o != nil && o.DocumentVersionId != nil {
		return true
	}

	return false
}

// SetDocumentVersionId gets a reference to the given string and assigns it to the DocumentVersionId field.
func (o *base_BTMParameterReference2434) SetDocumentVersionId(v string) {
	o.DocumentVersionId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *base_BTMParameterReference2434) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterReference2434) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *base_BTMParameterReference2434) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *base_BTMParameterReference2434) SetElementId(v string) {
	o.ElementId = &v
}

// GetElementLibraryData returns the ElementLibraryData field value if set, zero value otherwise.
func (o *base_BTMParameterReference2434) GetElementLibraryData() BTElementLibraryReferenceData3133 {
	if o == nil || o.ElementLibraryData == nil {
		var ret BTElementLibraryReferenceData3133
		return ret
	}
	return *o.ElementLibraryData
}

// GetElementLibraryDataOk returns a tuple with the ElementLibraryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterReference2434) GetElementLibraryDataOk() (*BTElementLibraryReferenceData3133, bool) {
	if o == nil || o.ElementLibraryData == nil {
		return nil, false
	}
	return o.ElementLibraryData, true
}

// HasElementLibraryData returns a boolean if a field has been set.
func (o *base_BTMParameterReference2434) HasElementLibraryData() bool {
	if o != nil && o.ElementLibraryData != nil {
		return true
	}

	return false
}

// SetElementLibraryData gets a reference to the given BTElementLibraryReferenceData3133 and assigns it to the ElementLibraryData field.
func (o *base_BTMParameterReference2434) SetElementLibraryData(v BTElementLibraryReferenceData3133) {
	o.ElementLibraryData = &v
}

// GetFeatureScriptType returns the FeatureScriptType field value if set, zero value otherwise.
func (o *base_BTMParameterReference2434) GetFeatureScriptType() string {
	if o == nil || o.FeatureScriptType == nil {
		var ret string
		return ret
	}
	return *o.FeatureScriptType
}

// GetFeatureScriptTypeOk returns a tuple with the FeatureScriptType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterReference2434) GetFeatureScriptTypeOk() (*string, bool) {
	if o == nil || o.FeatureScriptType == nil {
		return nil, false
	}
	return o.FeatureScriptType, true
}

// HasFeatureScriptType returns a boolean if a field has been set.
func (o *base_BTMParameterReference2434) HasFeatureScriptType() bool {
	if o != nil && o.FeatureScriptType != nil {
		return true
	}

	return false
}

// SetFeatureScriptType gets a reference to the given string and assigns it to the FeatureScriptType field.
func (o *base_BTMParameterReference2434) SetFeatureScriptType(v string) {
	o.FeatureScriptType = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *base_BTMParameterReference2434) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterReference2434) GetIdsOk() ([]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *base_BTMParameterReference2434) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *base_BTMParameterReference2434) SetIds(v []string) {
	o.Ids = v
}

// GetMicroversioId returns the MicroversioId field value if set, zero value otherwise.
func (o *base_BTMParameterReference2434) GetMicroversioId() string {
	if o == nil || o.MicroversioId == nil {
		var ret string
		return ret
	}
	return *o.MicroversioId
}

// GetMicroversioIdOk returns a tuple with the MicroversioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterReference2434) GetMicroversioIdOk() (*string, bool) {
	if o == nil || o.MicroversioId == nil {
		return nil, false
	}
	return o.MicroversioId, true
}

// HasMicroversioId returns a boolean if a field has been set.
func (o *base_BTMParameterReference2434) HasMicroversioId() bool {
	if o != nil && o.MicroversioId != nil {
		return true
	}

	return false
}

// SetMicroversioId gets a reference to the given string and assigns it to the MicroversioId field.
func (o *base_BTMParameterReference2434) SetMicroversioId(v string) {
	o.MicroversioId = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *base_BTMParameterReference2434) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterReference2434) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *base_BTMParameterReference2434) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *base_BTMParameterReference2434) SetNamespace(v string) {
	o.Namespace = &v
}

func (o base_BTMParameterReference2434) MarshalJSON() ([]byte, error) {
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
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.DocumentVersionId != nil {
		toSerialize["documentVersionId"] = o.DocumentVersionId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ElementLibraryData != nil {
		toSerialize["elementLibraryData"] = o.ElementLibraryData
	}
	if o.FeatureScriptType != nil {
		toSerialize["featureScriptType"] = o.FeatureScriptType
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.MicroversioId != nil {
		toSerialize["microversioId"] = o.MicroversioId
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}
