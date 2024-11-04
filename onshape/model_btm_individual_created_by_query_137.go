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

// BTMIndividualCreatedByQuery137 struct for BTMIndividualCreatedByQuery137
type BTMIndividualCreatedByQuery137 struct {
	BTMIndividualQuery138
	BtType                     *string                    `json:"btType,omitempty"`
	DeterministicIdList        *BTMIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds           []string                   `json:"deterministicIds,omitempty"`
	GenerateSectionEntityQuery *bool                      `json:"generateSectionEntityQuery,omitempty"`
	GeneratedSectionQueryId    *string                    `json:"generatedSectionQueryId,omitempty"`
	// Microversion that resulted from the import.
	ImportMicroversion *string                    `json:"importMicroversion,omitempty"`
	NodeId             *string                    `json:"nodeId,omitempty"`
	PersistentQuery    *BTPStatement269           `json:"persistentQuery,omitempty"`
	Query              *BTMIndividualQueryBase139 `json:"query,omitempty"`
	QueryStatement     *BTPStatement269           `json:"queryStatement,omitempty"`
	QueryString        *string                    `json:"queryString,omitempty"`
	VariableName       *BTMIndividualQuery138     `json:"variableName,omitempty"`
	BodyType           *GBTBodyType               `json:"bodyType,omitempty"`
	EntityType         *GBTEntityType             `json:"entityType,omitempty"`
	FeatureId          *string                    `json:"featureId,omitempty"`
	FilterConstruction *bool                      `json:"filterConstruction,omitempty"`
	// Type of JSON object.
	BtType *string `json:"btType,omitempty"`
}

// NewBTMIndividualCreatedByQuery137 instantiates a new BTMIndividualCreatedByQuery137 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMIndividualCreatedByQuery137() *BTMIndividualCreatedByQuery137 {
	this := BTMIndividualCreatedByQuery137{}
	return &this
}

// NewBTMIndividualCreatedByQuery137WithDefaults instantiates a new BTMIndividualCreatedByQuery137 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMIndividualCreatedByQuery137WithDefaults() *BTMIndividualCreatedByQuery137 {
	this := BTMIndividualCreatedByQuery137{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMIndividualCreatedByQuery137) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetDeterministicIdList() BTMIndividualQueryBase139 {
	if o == nil || o.DeterministicIdList == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *BTMIndividualCreatedByQuery137) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	o.DeterministicIdList = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTMIndividualCreatedByQuery137) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetGenerateSectionEntityQuery returns the GenerateSectionEntityQuery field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetGenerateSectionEntityQuery() bool {
	if o == nil || o.GenerateSectionEntityQuery == nil {
		var ret bool
		return ret
	}
	return *o.GenerateSectionEntityQuery
}

// GetGenerateSectionEntityQueryOk returns a tuple with the GenerateSectionEntityQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetGenerateSectionEntityQueryOk() (*bool, bool) {
	if o == nil || o.GenerateSectionEntityQuery == nil {
		return nil, false
	}
	return o.GenerateSectionEntityQuery, true
}

// HasGenerateSectionEntityQuery returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasGenerateSectionEntityQuery() bool {
	if o != nil && o.GenerateSectionEntityQuery != nil {
		return true
	}

	return false
}

// SetGenerateSectionEntityQuery gets a reference to the given bool and assigns it to the GenerateSectionEntityQuery field.
func (o *BTMIndividualCreatedByQuery137) SetGenerateSectionEntityQuery(v bool) {
	o.GenerateSectionEntityQuery = &v
}

// GetGeneratedSectionQueryId returns the GeneratedSectionQueryId field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetGeneratedSectionQueryId() string {
	if o == nil || o.GeneratedSectionQueryId == nil {
		var ret string
		return ret
	}
	return *o.GeneratedSectionQueryId
}

// GetGeneratedSectionQueryIdOk returns a tuple with the GeneratedSectionQueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetGeneratedSectionQueryIdOk() (*string, bool) {
	if o == nil || o.GeneratedSectionQueryId == nil {
		return nil, false
	}
	return o.GeneratedSectionQueryId, true
}

// HasGeneratedSectionQueryId returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasGeneratedSectionQueryId() bool {
	if o != nil && o.GeneratedSectionQueryId != nil {
		return true
	}

	return false
}

// SetGeneratedSectionQueryId gets a reference to the given string and assigns it to the GeneratedSectionQueryId field.
func (o *BTMIndividualCreatedByQuery137) SetGeneratedSectionQueryId(v string) {
	o.GeneratedSectionQueryId = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMIndividualCreatedByQuery137) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMIndividualCreatedByQuery137) SetNodeId(v string) {
	o.NodeId = &v
}

// GetPersistentQuery returns the PersistentQuery field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetPersistentQuery() BTPStatement269 {
	if o == nil || o.PersistentQuery == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.PersistentQuery
}

// GetPersistentQueryOk returns a tuple with the PersistentQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetPersistentQueryOk() (*BTPStatement269, bool) {
	if o == nil || o.PersistentQuery == nil {
		return nil, false
	}
	return o.PersistentQuery, true
}

// HasPersistentQuery returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasPersistentQuery() bool {
	if o != nil && o.PersistentQuery != nil {
		return true
	}

	return false
}

// SetPersistentQuery gets a reference to the given BTPStatement269 and assigns it to the PersistentQuery field.
func (o *BTMIndividualCreatedByQuery137) SetPersistentQuery(v BTPStatement269) {
	o.PersistentQuery = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetQuery() BTMIndividualQueryBase139 {
	if o == nil || o.Query == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *BTMIndividualCreatedByQuery137) SetQuery(v BTMIndividualQueryBase139) {
	o.Query = &v
}

// GetQueryStatement returns the QueryStatement field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetQueryStatement() BTPStatement269 {
	if o == nil || o.QueryStatement == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.QueryStatement
}

// GetQueryStatementOk returns a tuple with the QueryStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetQueryStatementOk() (*BTPStatement269, bool) {
	if o == nil || o.QueryStatement == nil {
		return nil, false
	}
	return o.QueryStatement, true
}

// HasQueryStatement returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasQueryStatement() bool {
	if o != nil && o.QueryStatement != nil {
		return true
	}

	return false
}

// SetQueryStatement gets a reference to the given BTPStatement269 and assigns it to the QueryStatement field.
func (o *BTMIndividualCreatedByQuery137) SetQueryStatement(v BTPStatement269) {
	o.QueryStatement = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *BTMIndividualCreatedByQuery137) SetQueryString(v string) {
	o.QueryString = &v
}

// GetVariableName returns the VariableName field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetVariableName() BTMIndividualQuery138 {
	if o == nil || o.VariableName == nil {
		var ret BTMIndividualQuery138
		return ret
	}
	return *o.VariableName
}

// GetVariableNameOk returns a tuple with the VariableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetVariableNameOk() (*BTMIndividualQuery138, bool) {
	if o == nil || o.VariableName == nil {
		return nil, false
	}
	return o.VariableName, true
}

// HasVariableName returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasVariableName() bool {
	if o != nil && o.VariableName != nil {
		return true
	}

	return false
}

// SetVariableName gets a reference to the given BTMIndividualQuery138 and assigns it to the VariableName field.
func (o *BTMIndividualCreatedByQuery137) SetVariableName(v BTMIndividualQuery138) {
	o.VariableName = &v
}

// GetBodyType returns the BodyType field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetBodyType() GBTBodyType {
	if o == nil || o.BodyType == nil {
		var ret GBTBodyType
		return ret
	}
	return *o.BodyType
}

// GetBodyTypeOk returns a tuple with the BodyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetBodyTypeOk() (*GBTBodyType, bool) {
	if o == nil || o.BodyType == nil {
		return nil, false
	}
	return o.BodyType, true
}

// HasBodyType returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasBodyType() bool {
	if o != nil && o.BodyType != nil {
		return true
	}

	return false
}

// SetBodyType gets a reference to the given GBTBodyType and assigns it to the BodyType field.
func (o *BTMIndividualCreatedByQuery137) SetBodyType(v GBTBodyType) {
	o.BodyType = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetEntityType() GBTEntityType {
	if o == nil || o.EntityType == nil {
		var ret GBTEntityType
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetEntityTypeOk() (*GBTEntityType, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given GBTEntityType and assigns it to the EntityType field.
func (o *BTMIndividualCreatedByQuery137) SetEntityType(v GBTEntityType) {
	o.EntityType = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTMIndividualCreatedByQuery137) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFilterConstruction returns the FilterConstruction field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetFilterConstruction() bool {
	if o == nil || o.FilterConstruction == nil {
		var ret bool
		return ret
	}
	return *o.FilterConstruction
}

// GetFilterConstructionOk returns a tuple with the FilterConstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetFilterConstructionOk() (*bool, bool) {
	if o == nil || o.FilterConstruction == nil {
		return nil, false
	}
	return o.FilterConstruction, true
}

// HasFilterConstruction returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasFilterConstruction() bool {
	if o != nil && o.FilterConstruction != nil {
		return true
	}

	return false
}

// SetFilterConstruction gets a reference to the given bool and assigns it to the FilterConstruction field.
func (o *BTMIndividualCreatedByQuery137) SetFilterConstruction(v bool) {
	o.FilterConstruction = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMIndividualCreatedByQuery137) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCreatedByQuery137) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMIndividualCreatedByQuery137) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMIndividualCreatedByQuery137) SetBtType(v string) {
	o.BtType = &v
}

func (o BTMIndividualCreatedByQuery137) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMIndividualQuery138, errBTMIndividualQuery138 := json.Marshal(o.BTMIndividualQuery138)
	if errBTMIndividualQuery138 != nil {
		return []byte{}, errBTMIndividualQuery138
	}
	errBTMIndividualQuery138 = json.Unmarshal([]byte(serializedBTMIndividualQuery138), &toSerialize)
	if errBTMIndividualQuery138 != nil {
		return []byte{}, errBTMIndividualQuery138
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DeterministicIdList != nil {
		toSerialize["deterministicIdList"] = o.DeterministicIdList
	}
	if o.DeterministicIds != nil {
		toSerialize["deterministicIds"] = o.DeterministicIds
	}
	if o.GenerateSectionEntityQuery != nil {
		toSerialize["generateSectionEntityQuery"] = o.GenerateSectionEntityQuery
	}
	if o.GeneratedSectionQueryId != nil {
		toSerialize["generatedSectionQueryId"] = o.GeneratedSectionQueryId
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.PersistentQuery != nil {
		toSerialize["persistentQuery"] = o.PersistentQuery
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryStatement != nil {
		toSerialize["queryStatement"] = o.QueryStatement
	}
	if o.QueryString != nil {
		toSerialize["queryString"] = o.QueryString
	}
	if o.VariableName != nil {
		toSerialize["variableName"] = o.VariableName
	}
	if o.BodyType != nil {
		toSerialize["bodyType"] = o.BodyType
	}
	if o.EntityType != nil {
		toSerialize["entityType"] = o.EntityType
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.FilterConstruction != nil {
		toSerialize["filterConstruction"] = o.FilterConstruction
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTMIndividualCreatedByQuery137 struct {
	value *BTMIndividualCreatedByQuery137
	isSet bool
}

func (v NullableBTMIndividualCreatedByQuery137) Get() *BTMIndividualCreatedByQuery137 {
	return v.value
}

func (v *NullableBTMIndividualCreatedByQuery137) Set(val *BTMIndividualCreatedByQuery137) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMIndividualCreatedByQuery137) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMIndividualCreatedByQuery137) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMIndividualCreatedByQuery137(val *BTMIndividualCreatedByQuery137) *NullableBTMIndividualCreatedByQuery137 {
	return &NullableBTMIndividualCreatedByQuery137{value: val, isSet: true}
}

func (v NullableBTMIndividualCreatedByQuery137) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMIndividualCreatedByQuery137) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
