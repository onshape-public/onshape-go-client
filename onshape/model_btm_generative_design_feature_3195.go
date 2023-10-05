/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23367-59c3c9f9feef
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMGenerativeDesignFeature3195 struct for BTMGenerativeDesignFeature3195
type BTMGenerativeDesignFeature3195 struct {
	BtType                                 *string                                   `json:"btType,omitempty"`
	FeatureId                              *string                                   `json:"featureId,omitempty"`
	FeatureType                            *string                                   `json:"featureType,omitempty"`
	ImportMicroversion                     *string                                   `json:"importMicroversion,omitempty"`
	Name                                   *string                                   `json:"name,omitempty"`
	Namespace                              *string                                   `json:"namespace,omitempty"`
	NodeId                                 *string                                   `json:"nodeId,omitempty"`
	Parameters                             []BTMParameter1                           `json:"parameters,omitempty"`
	ReturnAfterSubfeatures                 *bool                                     `json:"returnAfterSubfeatures,omitempty"`
	SubFeatures                            []BTMFeature134                           `json:"subFeatures,omitempty"`
	Suppressed                             *bool                                     `json:"suppressed,omitempty"`
	SuppressionConfigured                  *bool                                     `json:"suppressionConfigured,omitempty"`
	VariableStudioReference                *bool                                     `json:"variableStudioReference,omitempty"`
	AuxiliaryTreeFeature                   *bool                                     `json:"auxiliaryTreeFeature,omitempty"`
	FeatureFolder                          *bool                                     `json:"featureFolder,omitempty"`
	FeatureListFieldIndex                  *int32                                    `json:"featureListFieldIndex,omitempty"`
	FieldIndexForOwnedMateConnectors       *int32                                    `json:"fieldIndexForOwnedMateConnectors,omitempty"`
	OccurrenceQueriesFromAllConfigurations []BTMIndividualQueryWithOccurrenceBase904 `json:"occurrenceQueriesFromAllConfigurations,omitempty"`
	ParametricInstanceFeature              *bool                                     `json:"parametricInstanceFeature,omitempty"`
	Version                                *int32                                    `json:"version,omitempty"`
	AdditiveSupportsOccurrences            []BTOccurrence74                          `json:"additiveSupportsOccurrences,omitempty"`
	DraftOccurrences                       []BTOccurrence74                          `json:"draftOccurrences,omitempty"`
	ExtrusionOccurrences                   []BTOccurrence74                          `json:"extrusionOccurrences,omitempty"`
	JsonRecipe                             *BTMParameterReferenceJSON790             `json:"jsonRecipe,omitempty"`
	MillingOccurrences                     []BTOccurrence74                          `json:"millingOccurrences,omitempty"`
	Occurrences                            []BTOccurrence74                          `json:"occurrences,omitempty"`
}

// NewBTMGenerativeDesignFeature3195 instantiates a new BTMGenerativeDesignFeature3195 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMGenerativeDesignFeature3195() *BTMGenerativeDesignFeature3195 {
	this := BTMGenerativeDesignFeature3195{}
	return &this
}

// NewBTMGenerativeDesignFeature3195WithDefaults instantiates a new BTMGenerativeDesignFeature3195 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMGenerativeDesignFeature3195WithDefaults() *BTMGenerativeDesignFeature3195 {
	this := BTMGenerativeDesignFeature3195{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMGenerativeDesignFeature3195) SetBtType(v string) {
	o.BtType = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTMGenerativeDesignFeature3195) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetFeatureType() string {
	if o == nil || o.FeatureType == nil {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetFeatureTypeOk() (*string, bool) {
	if o == nil || o.FeatureType == nil {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasFeatureType() bool {
	if o != nil && o.FeatureType != nil {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *BTMGenerativeDesignFeature3195) SetFeatureType(v string) {
	o.FeatureType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMGenerativeDesignFeature3195) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMGenerativeDesignFeature3195) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMGenerativeDesignFeature3195) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMGenerativeDesignFeature3195) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetParameters() []BTMParameter1 {
	if o == nil || o.Parameters == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetParametersOk() ([]BTMParameter1, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *BTMGenerativeDesignFeature3195) SetParameters(v []BTMParameter1) {
	o.Parameters = v
}

// GetReturnAfterSubfeatures returns the ReturnAfterSubfeatures field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetReturnAfterSubfeatures() bool {
	if o == nil || o.ReturnAfterSubfeatures == nil {
		var ret bool
		return ret
	}
	return *o.ReturnAfterSubfeatures
}

// GetReturnAfterSubfeaturesOk returns a tuple with the ReturnAfterSubfeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetReturnAfterSubfeaturesOk() (*bool, bool) {
	if o == nil || o.ReturnAfterSubfeatures == nil {
		return nil, false
	}
	return o.ReturnAfterSubfeatures, true
}

// HasReturnAfterSubfeatures returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasReturnAfterSubfeatures() bool {
	if o != nil && o.ReturnAfterSubfeatures != nil {
		return true
	}

	return false
}

// SetReturnAfterSubfeatures gets a reference to the given bool and assigns it to the ReturnAfterSubfeatures field.
func (o *BTMGenerativeDesignFeature3195) SetReturnAfterSubfeatures(v bool) {
	o.ReturnAfterSubfeatures = &v
}

// GetSubFeatures returns the SubFeatures field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetSubFeatures() []BTMFeature134 {
	if o == nil || o.SubFeatures == nil {
		var ret []BTMFeature134
		return ret
	}
	return o.SubFeatures
}

// GetSubFeaturesOk returns a tuple with the SubFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetSubFeaturesOk() ([]BTMFeature134, bool) {
	if o == nil || o.SubFeatures == nil {
		return nil, false
	}
	return o.SubFeatures, true
}

// HasSubFeatures returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasSubFeatures() bool {
	if o != nil && o.SubFeatures != nil {
		return true
	}

	return false
}

// SetSubFeatures gets a reference to the given []BTMFeature134 and assigns it to the SubFeatures field.
func (o *BTMGenerativeDesignFeature3195) SetSubFeatures(v []BTMFeature134) {
	o.SubFeatures = v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetSuppressed() bool {
	if o == nil || o.Suppressed == nil {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetSuppressedOk() (*bool, bool) {
	if o == nil || o.Suppressed == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasSuppressed() bool {
	if o != nil && o.Suppressed != nil {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *BTMGenerativeDesignFeature3195) SetSuppressed(v bool) {
	o.Suppressed = &v
}

// GetSuppressionConfigured returns the SuppressionConfigured field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetSuppressionConfigured() bool {
	if o == nil || o.SuppressionConfigured == nil {
		var ret bool
		return ret
	}
	return *o.SuppressionConfigured
}

// GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetSuppressionConfiguredOk() (*bool, bool) {
	if o == nil || o.SuppressionConfigured == nil {
		return nil, false
	}
	return o.SuppressionConfigured, true
}

// HasSuppressionConfigured returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasSuppressionConfigured() bool {
	if o != nil && o.SuppressionConfigured != nil {
		return true
	}

	return false
}

// SetSuppressionConfigured gets a reference to the given bool and assigns it to the SuppressionConfigured field.
func (o *BTMGenerativeDesignFeature3195) SetSuppressionConfigured(v bool) {
	o.SuppressionConfigured = &v
}

// GetVariableStudioReference returns the VariableStudioReference field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetVariableStudioReference() bool {
	if o == nil || o.VariableStudioReference == nil {
		var ret bool
		return ret
	}
	return *o.VariableStudioReference
}

// GetVariableStudioReferenceOk returns a tuple with the VariableStudioReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetVariableStudioReferenceOk() (*bool, bool) {
	if o == nil || o.VariableStudioReference == nil {
		return nil, false
	}
	return o.VariableStudioReference, true
}

// HasVariableStudioReference returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasVariableStudioReference() bool {
	if o != nil && o.VariableStudioReference != nil {
		return true
	}

	return false
}

// SetVariableStudioReference gets a reference to the given bool and assigns it to the VariableStudioReference field.
func (o *BTMGenerativeDesignFeature3195) SetVariableStudioReference(v bool) {
	o.VariableStudioReference = &v
}

// GetAuxiliaryTreeFeature returns the AuxiliaryTreeFeature field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetAuxiliaryTreeFeature() bool {
	if o == nil || o.AuxiliaryTreeFeature == nil {
		var ret bool
		return ret
	}
	return *o.AuxiliaryTreeFeature
}

// GetAuxiliaryTreeFeatureOk returns a tuple with the AuxiliaryTreeFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetAuxiliaryTreeFeatureOk() (*bool, bool) {
	if o == nil || o.AuxiliaryTreeFeature == nil {
		return nil, false
	}
	return o.AuxiliaryTreeFeature, true
}

// HasAuxiliaryTreeFeature returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasAuxiliaryTreeFeature() bool {
	if o != nil && o.AuxiliaryTreeFeature != nil {
		return true
	}

	return false
}

// SetAuxiliaryTreeFeature gets a reference to the given bool and assigns it to the AuxiliaryTreeFeature field.
func (o *BTMGenerativeDesignFeature3195) SetAuxiliaryTreeFeature(v bool) {
	o.AuxiliaryTreeFeature = &v
}

// GetFeatureFolder returns the FeatureFolder field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetFeatureFolder() bool {
	if o == nil || o.FeatureFolder == nil {
		var ret bool
		return ret
	}
	return *o.FeatureFolder
}

// GetFeatureFolderOk returns a tuple with the FeatureFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetFeatureFolderOk() (*bool, bool) {
	if o == nil || o.FeatureFolder == nil {
		return nil, false
	}
	return o.FeatureFolder, true
}

// HasFeatureFolder returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasFeatureFolder() bool {
	if o != nil && o.FeatureFolder != nil {
		return true
	}

	return false
}

// SetFeatureFolder gets a reference to the given bool and assigns it to the FeatureFolder field.
func (o *BTMGenerativeDesignFeature3195) SetFeatureFolder(v bool) {
	o.FeatureFolder = &v
}

// GetFeatureListFieldIndex returns the FeatureListFieldIndex field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetFeatureListFieldIndex() int32 {
	if o == nil || o.FeatureListFieldIndex == nil {
		var ret int32
		return ret
	}
	return *o.FeatureListFieldIndex
}

// GetFeatureListFieldIndexOk returns a tuple with the FeatureListFieldIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetFeatureListFieldIndexOk() (*int32, bool) {
	if o == nil || o.FeatureListFieldIndex == nil {
		return nil, false
	}
	return o.FeatureListFieldIndex, true
}

// HasFeatureListFieldIndex returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasFeatureListFieldIndex() bool {
	if o != nil && o.FeatureListFieldIndex != nil {
		return true
	}

	return false
}

// SetFeatureListFieldIndex gets a reference to the given int32 and assigns it to the FeatureListFieldIndex field.
func (o *BTMGenerativeDesignFeature3195) SetFeatureListFieldIndex(v int32) {
	o.FeatureListFieldIndex = &v
}

// GetFieldIndexForOwnedMateConnectors returns the FieldIndexForOwnedMateConnectors field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetFieldIndexForOwnedMateConnectors() int32 {
	if o == nil || o.FieldIndexForOwnedMateConnectors == nil {
		var ret int32
		return ret
	}
	return *o.FieldIndexForOwnedMateConnectors
}

// GetFieldIndexForOwnedMateConnectorsOk returns a tuple with the FieldIndexForOwnedMateConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetFieldIndexForOwnedMateConnectorsOk() (*int32, bool) {
	if o == nil || o.FieldIndexForOwnedMateConnectors == nil {
		return nil, false
	}
	return o.FieldIndexForOwnedMateConnectors, true
}

// HasFieldIndexForOwnedMateConnectors returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasFieldIndexForOwnedMateConnectors() bool {
	if o != nil && o.FieldIndexForOwnedMateConnectors != nil {
		return true
	}

	return false
}

// SetFieldIndexForOwnedMateConnectors gets a reference to the given int32 and assigns it to the FieldIndexForOwnedMateConnectors field.
func (o *BTMGenerativeDesignFeature3195) SetFieldIndexForOwnedMateConnectors(v int32) {
	o.FieldIndexForOwnedMateConnectors = &v
}

// GetOccurrenceQueriesFromAllConfigurations returns the OccurrenceQueriesFromAllConfigurations field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetOccurrenceQueriesFromAllConfigurations() []BTMIndividualQueryWithOccurrenceBase904 {
	if o == nil || o.OccurrenceQueriesFromAllConfigurations == nil {
		var ret []BTMIndividualQueryWithOccurrenceBase904
		return ret
	}
	return o.OccurrenceQueriesFromAllConfigurations
}

// GetOccurrenceQueriesFromAllConfigurationsOk returns a tuple with the OccurrenceQueriesFromAllConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetOccurrenceQueriesFromAllConfigurationsOk() ([]BTMIndividualQueryWithOccurrenceBase904, bool) {
	if o == nil || o.OccurrenceQueriesFromAllConfigurations == nil {
		return nil, false
	}
	return o.OccurrenceQueriesFromAllConfigurations, true
}

// HasOccurrenceQueriesFromAllConfigurations returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasOccurrenceQueriesFromAllConfigurations() bool {
	if o != nil && o.OccurrenceQueriesFromAllConfigurations != nil {
		return true
	}

	return false
}

// SetOccurrenceQueriesFromAllConfigurations gets a reference to the given []BTMIndividualQueryWithOccurrenceBase904 and assigns it to the OccurrenceQueriesFromAllConfigurations field.
func (o *BTMGenerativeDesignFeature3195) SetOccurrenceQueriesFromAllConfigurations(v []BTMIndividualQueryWithOccurrenceBase904) {
	o.OccurrenceQueriesFromAllConfigurations = v
}

// GetParametricInstanceFeature returns the ParametricInstanceFeature field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetParametricInstanceFeature() bool {
	if o == nil || o.ParametricInstanceFeature == nil {
		var ret bool
		return ret
	}
	return *o.ParametricInstanceFeature
}

// GetParametricInstanceFeatureOk returns a tuple with the ParametricInstanceFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetParametricInstanceFeatureOk() (*bool, bool) {
	if o == nil || o.ParametricInstanceFeature == nil {
		return nil, false
	}
	return o.ParametricInstanceFeature, true
}

// HasParametricInstanceFeature returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasParametricInstanceFeature() bool {
	if o != nil && o.ParametricInstanceFeature != nil {
		return true
	}

	return false
}

// SetParametricInstanceFeature gets a reference to the given bool and assigns it to the ParametricInstanceFeature field.
func (o *BTMGenerativeDesignFeature3195) SetParametricInstanceFeature(v bool) {
	o.ParametricInstanceFeature = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *BTMGenerativeDesignFeature3195) SetVersion(v int32) {
	o.Version = &v
}

// GetAdditiveSupportsOccurrences returns the AdditiveSupportsOccurrences field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetAdditiveSupportsOccurrences() []BTOccurrence74 {
	if o == nil || o.AdditiveSupportsOccurrences == nil {
		var ret []BTOccurrence74
		return ret
	}
	return o.AdditiveSupportsOccurrences
}

// GetAdditiveSupportsOccurrencesOk returns a tuple with the AdditiveSupportsOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetAdditiveSupportsOccurrencesOk() ([]BTOccurrence74, bool) {
	if o == nil || o.AdditiveSupportsOccurrences == nil {
		return nil, false
	}
	return o.AdditiveSupportsOccurrences, true
}

// HasAdditiveSupportsOccurrences returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasAdditiveSupportsOccurrences() bool {
	if o != nil && o.AdditiveSupportsOccurrences != nil {
		return true
	}

	return false
}

// SetAdditiveSupportsOccurrences gets a reference to the given []BTOccurrence74 and assigns it to the AdditiveSupportsOccurrences field.
func (o *BTMGenerativeDesignFeature3195) SetAdditiveSupportsOccurrences(v []BTOccurrence74) {
	o.AdditiveSupportsOccurrences = v
}

// GetDraftOccurrences returns the DraftOccurrences field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetDraftOccurrences() []BTOccurrence74 {
	if o == nil || o.DraftOccurrences == nil {
		var ret []BTOccurrence74
		return ret
	}
	return o.DraftOccurrences
}

// GetDraftOccurrencesOk returns a tuple with the DraftOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetDraftOccurrencesOk() ([]BTOccurrence74, bool) {
	if o == nil || o.DraftOccurrences == nil {
		return nil, false
	}
	return o.DraftOccurrences, true
}

// HasDraftOccurrences returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasDraftOccurrences() bool {
	if o != nil && o.DraftOccurrences != nil {
		return true
	}

	return false
}

// SetDraftOccurrences gets a reference to the given []BTOccurrence74 and assigns it to the DraftOccurrences field.
func (o *BTMGenerativeDesignFeature3195) SetDraftOccurrences(v []BTOccurrence74) {
	o.DraftOccurrences = v
}

// GetExtrusionOccurrences returns the ExtrusionOccurrences field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetExtrusionOccurrences() []BTOccurrence74 {
	if o == nil || o.ExtrusionOccurrences == nil {
		var ret []BTOccurrence74
		return ret
	}
	return o.ExtrusionOccurrences
}

// GetExtrusionOccurrencesOk returns a tuple with the ExtrusionOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetExtrusionOccurrencesOk() ([]BTOccurrence74, bool) {
	if o == nil || o.ExtrusionOccurrences == nil {
		return nil, false
	}
	return o.ExtrusionOccurrences, true
}

// HasExtrusionOccurrences returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasExtrusionOccurrences() bool {
	if o != nil && o.ExtrusionOccurrences != nil {
		return true
	}

	return false
}

// SetExtrusionOccurrences gets a reference to the given []BTOccurrence74 and assigns it to the ExtrusionOccurrences field.
func (o *BTMGenerativeDesignFeature3195) SetExtrusionOccurrences(v []BTOccurrence74) {
	o.ExtrusionOccurrences = v
}

// GetJsonRecipe returns the JsonRecipe field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetJsonRecipe() BTMParameterReferenceJSON790 {
	if o == nil || o.JsonRecipe == nil {
		var ret BTMParameterReferenceJSON790
		return ret
	}
	return *o.JsonRecipe
}

// GetJsonRecipeOk returns a tuple with the JsonRecipe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetJsonRecipeOk() (*BTMParameterReferenceJSON790, bool) {
	if o == nil || o.JsonRecipe == nil {
		return nil, false
	}
	return o.JsonRecipe, true
}

// HasJsonRecipe returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasJsonRecipe() bool {
	if o != nil && o.JsonRecipe != nil {
		return true
	}

	return false
}

// SetJsonRecipe gets a reference to the given BTMParameterReferenceJSON790 and assigns it to the JsonRecipe field.
func (o *BTMGenerativeDesignFeature3195) SetJsonRecipe(v BTMParameterReferenceJSON790) {
	o.JsonRecipe = &v
}

// GetMillingOccurrences returns the MillingOccurrences field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetMillingOccurrences() []BTOccurrence74 {
	if o == nil || o.MillingOccurrences == nil {
		var ret []BTOccurrence74
		return ret
	}
	return o.MillingOccurrences
}

// GetMillingOccurrencesOk returns a tuple with the MillingOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetMillingOccurrencesOk() ([]BTOccurrence74, bool) {
	if o == nil || o.MillingOccurrences == nil {
		return nil, false
	}
	return o.MillingOccurrences, true
}

// HasMillingOccurrences returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasMillingOccurrences() bool {
	if o != nil && o.MillingOccurrences != nil {
		return true
	}

	return false
}

// SetMillingOccurrences gets a reference to the given []BTOccurrence74 and assigns it to the MillingOccurrences field.
func (o *BTMGenerativeDesignFeature3195) SetMillingOccurrences(v []BTOccurrence74) {
	o.MillingOccurrences = v
}

// GetOccurrences returns the Occurrences field value if set, zero value otherwise.
func (o *BTMGenerativeDesignFeature3195) GetOccurrences() []BTOccurrence74 {
	if o == nil || o.Occurrences == nil {
		var ret []BTOccurrence74
		return ret
	}
	return o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMGenerativeDesignFeature3195) GetOccurrencesOk() ([]BTOccurrence74, bool) {
	if o == nil || o.Occurrences == nil {
		return nil, false
	}
	return o.Occurrences, true
}

// HasOccurrences returns a boolean if a field has been set.
func (o *BTMGenerativeDesignFeature3195) HasOccurrences() bool {
	if o != nil && o.Occurrences != nil {
		return true
	}

	return false
}

// SetOccurrences gets a reference to the given []BTOccurrence74 and assigns it to the Occurrences field.
func (o *BTMGenerativeDesignFeature3195) SetOccurrences(v []BTOccurrence74) {
	o.Occurrences = v
}

func (o BTMGenerativeDesignFeature3195) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.FeatureType != nil {
		toSerialize["featureType"] = o.FeatureType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.ReturnAfterSubfeatures != nil {
		toSerialize["returnAfterSubfeatures"] = o.ReturnAfterSubfeatures
	}
	if o.SubFeatures != nil {
		toSerialize["subFeatures"] = o.SubFeatures
	}
	if o.Suppressed != nil {
		toSerialize["suppressed"] = o.Suppressed
	}
	if o.SuppressionConfigured != nil {
		toSerialize["suppressionConfigured"] = o.SuppressionConfigured
	}
	if o.VariableStudioReference != nil {
		toSerialize["variableStudioReference"] = o.VariableStudioReference
	}
	if o.AuxiliaryTreeFeature != nil {
		toSerialize["auxiliaryTreeFeature"] = o.AuxiliaryTreeFeature
	}
	if o.FeatureFolder != nil {
		toSerialize["featureFolder"] = o.FeatureFolder
	}
	if o.FeatureListFieldIndex != nil {
		toSerialize["featureListFieldIndex"] = o.FeatureListFieldIndex
	}
	if o.FieldIndexForOwnedMateConnectors != nil {
		toSerialize["fieldIndexForOwnedMateConnectors"] = o.FieldIndexForOwnedMateConnectors
	}
	if o.OccurrenceQueriesFromAllConfigurations != nil {
		toSerialize["occurrenceQueriesFromAllConfigurations"] = o.OccurrenceQueriesFromAllConfigurations
	}
	if o.ParametricInstanceFeature != nil {
		toSerialize["parametricInstanceFeature"] = o.ParametricInstanceFeature
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.AdditiveSupportsOccurrences != nil {
		toSerialize["additiveSupportsOccurrences"] = o.AdditiveSupportsOccurrences
	}
	if o.DraftOccurrences != nil {
		toSerialize["draftOccurrences"] = o.DraftOccurrences
	}
	if o.ExtrusionOccurrences != nil {
		toSerialize["extrusionOccurrences"] = o.ExtrusionOccurrences
	}
	if o.JsonRecipe != nil {
		toSerialize["jsonRecipe"] = o.JsonRecipe
	}
	if o.MillingOccurrences != nil {
		toSerialize["millingOccurrences"] = o.MillingOccurrences
	}
	if o.Occurrences != nil {
		toSerialize["occurrences"] = o.Occurrences
	}
	return json.Marshal(toSerialize)
}

type NullableBTMGenerativeDesignFeature3195 struct {
	value *BTMGenerativeDesignFeature3195
	isSet bool
}

func (v NullableBTMGenerativeDesignFeature3195) Get() *BTMGenerativeDesignFeature3195 {
	return v.value
}

func (v *NullableBTMGenerativeDesignFeature3195) Set(val *BTMGenerativeDesignFeature3195) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMGenerativeDesignFeature3195) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMGenerativeDesignFeature3195) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMGenerativeDesignFeature3195(val *BTMGenerativeDesignFeature3195) *NullableBTMGenerativeDesignFeature3195 {
	return &NullableBTMGenerativeDesignFeature3195{value: val, isSet: true}
}

func (v NullableBTMGenerativeDesignFeature3195) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMGenerativeDesignFeature3195) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
