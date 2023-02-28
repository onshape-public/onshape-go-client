/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.160.12117-533984354e9a
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMParameterReferenceTable917 struct for BTMParameterReferenceTable917
type BTMParameterReferenceTable917 struct {
	ImportMicroversion *string                            `json:"importMicroversion,omitempty"`
	NodeId             *string                            `json:"nodeId,omitempty"`
	ParameterId        *string                            `json:"parameterId,omitempty"`
	BtType             *string                            `json:"btType,omitempty"`
	DocumentId         *string                            `json:"documentId,omitempty"`
	DocumentVersionId  *string                            `json:"documentVersionId,omitempty"`
	ElementId          *string                            `json:"elementId,omitempty"`
	ElementLibraryData *BTElementLibraryReferenceData3133 `json:"elementLibraryData,omitempty"`
	FeatureScriptType  *string                            `json:"featureScriptType,omitempty"`
	Ids                []string                           `json:"ids,omitempty"`
	MicroversioId      *string                            `json:"microversioId,omitempty"`
	Namespace          *string                            `json:"namespace,omitempty"`
}

// NewBTMParameterReferenceTable917 instantiates a new BTMParameterReferenceTable917 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterReferenceTable917() *BTMParameterReferenceTable917 {
	this := BTMParameterReferenceTable917{}
	return &this
}

// NewBTMParameterReferenceTable917WithDefaults instantiates a new BTMParameterReferenceTable917 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterReferenceTable917WithDefaults() *BTMParameterReferenceTable917 {
	this := BTMParameterReferenceTable917{}
	return &this
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterReferenceTable917) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReferenceTable917) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterReferenceTable917) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterReferenceTable917) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterReferenceTable917) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReferenceTable917) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameterReferenceTable917) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameterReferenceTable917) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterReferenceTable917) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReferenceTable917) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterReferenceTable917) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterReferenceTable917) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterReferenceTable917) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReferenceTable917) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterReferenceTable917) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterReferenceTable917) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTMParameterReferenceTable917) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReferenceTable917) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTMParameterReferenceTable917) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTMParameterReferenceTable917) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentVersionId returns the DocumentVersionId field value if set, zero value otherwise.
func (o *BTMParameterReferenceTable917) GetDocumentVersionId() string {
	if o == nil || o.DocumentVersionId == nil {
		var ret string
		return ret
	}
	return *o.DocumentVersionId
}

// GetDocumentVersionIdOk returns a tuple with the DocumentVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReferenceTable917) GetDocumentVersionIdOk() (*string, bool) {
	if o == nil || o.DocumentVersionId == nil {
		return nil, false
	}
	return o.DocumentVersionId, true
}

// HasDocumentVersionId returns a boolean if a field has been set.
func (o *BTMParameterReferenceTable917) HasDocumentVersionId() bool {
	if o != nil && o.DocumentVersionId != nil {
		return true
	}

	return false
}

// SetDocumentVersionId gets a reference to the given string and assigns it to the DocumentVersionId field.
func (o *BTMParameterReferenceTable917) SetDocumentVersionId(v string) {
	o.DocumentVersionId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTMParameterReferenceTable917) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReferenceTable917) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTMParameterReferenceTable917) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTMParameterReferenceTable917) SetElementId(v string) {
	o.ElementId = &v
}

// GetElementLibraryData returns the ElementLibraryData field value if set, zero value otherwise.
func (o *BTMParameterReferenceTable917) GetElementLibraryData() BTElementLibraryReferenceData3133 {
	if o == nil || o.ElementLibraryData == nil {
		var ret BTElementLibraryReferenceData3133
		return ret
	}
	return *o.ElementLibraryData
}

// GetElementLibraryDataOk returns a tuple with the ElementLibraryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReferenceTable917) GetElementLibraryDataOk() (*BTElementLibraryReferenceData3133, bool) {
	if o == nil || o.ElementLibraryData == nil {
		return nil, false
	}
	return o.ElementLibraryData, true
}

// HasElementLibraryData returns a boolean if a field has been set.
func (o *BTMParameterReferenceTable917) HasElementLibraryData() bool {
	if o != nil && o.ElementLibraryData != nil {
		return true
	}

	return false
}

// SetElementLibraryData gets a reference to the given BTElementLibraryReferenceData3133 and assigns it to the ElementLibraryData field.
func (o *BTMParameterReferenceTable917) SetElementLibraryData(v BTElementLibraryReferenceData3133) {
	o.ElementLibraryData = &v
}

// GetFeatureScriptType returns the FeatureScriptType field value if set, zero value otherwise.
func (o *BTMParameterReferenceTable917) GetFeatureScriptType() string {
	if o == nil || o.FeatureScriptType == nil {
		var ret string
		return ret
	}
	return *o.FeatureScriptType
}

// GetFeatureScriptTypeOk returns a tuple with the FeatureScriptType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReferenceTable917) GetFeatureScriptTypeOk() (*string, bool) {
	if o == nil || o.FeatureScriptType == nil {
		return nil, false
	}
	return o.FeatureScriptType, true
}

// HasFeatureScriptType returns a boolean if a field has been set.
func (o *BTMParameterReferenceTable917) HasFeatureScriptType() bool {
	if o != nil && o.FeatureScriptType != nil {
		return true
	}

	return false
}

// SetFeatureScriptType gets a reference to the given string and assigns it to the FeatureScriptType field.
func (o *BTMParameterReferenceTable917) SetFeatureScriptType(v string) {
	o.FeatureScriptType = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *BTMParameterReferenceTable917) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReferenceTable917) GetIdsOk() ([]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *BTMParameterReferenceTable917) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *BTMParameterReferenceTable917) SetIds(v []string) {
	o.Ids = v
}

// GetMicroversioId returns the MicroversioId field value if set, zero value otherwise.
func (o *BTMParameterReferenceTable917) GetMicroversioId() string {
	if o == nil || o.MicroversioId == nil {
		var ret string
		return ret
	}
	return *o.MicroversioId
}

// GetMicroversioIdOk returns a tuple with the MicroversioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReferenceTable917) GetMicroversioIdOk() (*string, bool) {
	if o == nil || o.MicroversioId == nil {
		return nil, false
	}
	return o.MicroversioId, true
}

// HasMicroversioId returns a boolean if a field has been set.
func (o *BTMParameterReferenceTable917) HasMicroversioId() bool {
	if o != nil && o.MicroversioId != nil {
		return true
	}

	return false
}

// SetMicroversioId gets a reference to the given string and assigns it to the MicroversioId field.
func (o *BTMParameterReferenceTable917) SetMicroversioId(v string) {
	o.MicroversioId = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMParameterReferenceTable917) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterReferenceTable917) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMParameterReferenceTable917) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMParameterReferenceTable917) SetNamespace(v string) {
	o.Namespace = &v
}

func (o BTMParameterReferenceTable917) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
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

type NullableBTMParameterReferenceTable917 struct {
	value *BTMParameterReferenceTable917
	isSet bool
}

func (v NullableBTMParameterReferenceTable917) Get() *BTMParameterReferenceTable917 {
	return v.value
}

func (v *NullableBTMParameterReferenceTable917) Set(val *BTMParameterReferenceTable917) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterReferenceTable917) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterReferenceTable917) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterReferenceTable917(val *BTMParameterReferenceTable917) *NullableBTMParameterReferenceTable917 {
	return &NullableBTMParameterReferenceTable917{value: val, isSet: true}
}

func (v NullableBTMParameterReferenceTable917) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterReferenceTable917) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
