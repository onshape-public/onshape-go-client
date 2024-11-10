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

// BTMMeshPointQuery1183 struct for BTMMeshPointQuery1183
type BTMMeshPointQuery1183 struct {
	BTMIndividualQueryWithOccurrenceBase904
	BtType *string `json:"btType,omitempty"`
	// Microversion that resulted from the import.
	ImportMicroversion         *string                    `json:"importMicroversion,omitempty"`
	NodeId                     *string                    `json:"nodeId,omitempty"`
	DeterministicIdList        *BTMIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds           []string                   `json:"deterministicIds,omitempty"`
	GenerateSectionEntityQuery *bool                      `json:"generateSectionEntityQuery,omitempty"`
	GeneratedSectionQueryId    *string                    `json:"generatedSectionQueryId,omitempty"`
	Query                      *BTMIndividualQueryBase139 `json:"query,omitempty"`
	QueryString                *string                    `json:"queryString,omitempty"`
	FullPathAsString           *string                    `json:"fullPathAsString,omitempty"`
	Occurrence                 *BTOccurrence74            `json:"occurrence,omitempty"`
	Path                       []string                   `json:"path,omitempty"`
	Origin                     *BTVector3d389             `json:"origin,omitempty"`
	// Type of JSON object.
	BtType *string `json:"btType,omitempty"`
}

// NewBTMMeshPointQuery1183 instantiates a new BTMMeshPointQuery1183 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMMeshPointQuery1183() *BTMMeshPointQuery1183 {
	this := BTMMeshPointQuery1183{}
	return &this
}

// NewBTMMeshPointQuery1183WithDefaults instantiates a new BTMMeshPointQuery1183 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMMeshPointQuery1183WithDefaults() *BTMMeshPointQuery1183 {
	this := BTMMeshPointQuery1183{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMMeshPointQuery1183) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMMeshPointQuery1183) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMMeshPointQuery1183) SetNodeId(v string) {
	o.NodeId = &v
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetDeterministicIdList() BTMIndividualQueryBase139 {
	if o == nil || o.DeterministicIdList == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *BTMMeshPointQuery1183) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	o.DeterministicIdList = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTMMeshPointQuery1183) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetGenerateSectionEntityQuery returns the GenerateSectionEntityQuery field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetGenerateSectionEntityQuery() bool {
	if o == nil || o.GenerateSectionEntityQuery == nil {
		var ret bool
		return ret
	}
	return *o.GenerateSectionEntityQuery
}

// GetGenerateSectionEntityQueryOk returns a tuple with the GenerateSectionEntityQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetGenerateSectionEntityQueryOk() (*bool, bool) {
	if o == nil || o.GenerateSectionEntityQuery == nil {
		return nil, false
	}
	return o.GenerateSectionEntityQuery, true
}

// HasGenerateSectionEntityQuery returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasGenerateSectionEntityQuery() bool {
	if o != nil && o.GenerateSectionEntityQuery != nil {
		return true
	}

	return false
}

// SetGenerateSectionEntityQuery gets a reference to the given bool and assigns it to the GenerateSectionEntityQuery field.
func (o *BTMMeshPointQuery1183) SetGenerateSectionEntityQuery(v bool) {
	o.GenerateSectionEntityQuery = &v
}

// GetGeneratedSectionQueryId returns the GeneratedSectionQueryId field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetGeneratedSectionQueryId() string {
	if o == nil || o.GeneratedSectionQueryId == nil {
		var ret string
		return ret
	}
	return *o.GeneratedSectionQueryId
}

// GetGeneratedSectionQueryIdOk returns a tuple with the GeneratedSectionQueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetGeneratedSectionQueryIdOk() (*string, bool) {
	if o == nil || o.GeneratedSectionQueryId == nil {
		return nil, false
	}
	return o.GeneratedSectionQueryId, true
}

// HasGeneratedSectionQueryId returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasGeneratedSectionQueryId() bool {
	if o != nil && o.GeneratedSectionQueryId != nil {
		return true
	}

	return false
}

// SetGeneratedSectionQueryId gets a reference to the given string and assigns it to the GeneratedSectionQueryId field.
func (o *BTMMeshPointQuery1183) SetGeneratedSectionQueryId(v string) {
	o.GeneratedSectionQueryId = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetQuery() BTMIndividualQueryBase139 {
	if o == nil || o.Query == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *BTMMeshPointQuery1183) SetQuery(v BTMIndividualQueryBase139) {
	o.Query = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *BTMMeshPointQuery1183) SetQueryString(v string) {
	o.QueryString = &v
}

// GetFullPathAsString returns the FullPathAsString field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetFullPathAsString() string {
	if o == nil || o.FullPathAsString == nil {
		var ret string
		return ret
	}
	return *o.FullPathAsString
}

// GetFullPathAsStringOk returns a tuple with the FullPathAsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetFullPathAsStringOk() (*string, bool) {
	if o == nil || o.FullPathAsString == nil {
		return nil, false
	}
	return o.FullPathAsString, true
}

// HasFullPathAsString returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasFullPathAsString() bool {
	if o != nil && o.FullPathAsString != nil {
		return true
	}

	return false
}

// SetFullPathAsString gets a reference to the given string and assigns it to the FullPathAsString field.
func (o *BTMMeshPointQuery1183) SetFullPathAsString(v string) {
	o.FullPathAsString = &v
}

// GetOccurrence returns the Occurrence field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetOccurrence() BTOccurrence74 {
	if o == nil || o.Occurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.Occurrence
}

// GetOccurrenceOk returns a tuple with the Occurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.Occurrence == nil {
		return nil, false
	}
	return o.Occurrence, true
}

// HasOccurrence returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasOccurrence() bool {
	if o != nil && o.Occurrence != nil {
		return true
	}

	return false
}

// SetOccurrence gets a reference to the given BTOccurrence74 and assigns it to the Occurrence field.
func (o *BTMMeshPointQuery1183) SetOccurrence(v BTOccurrence74) {
	o.Occurrence = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetPath() []string {
	if o == nil || o.Path == nil {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetPathOk() ([]string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *BTMMeshPointQuery1183) SetPath(v []string) {
	o.Path = v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTMMeshPointQuery1183) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMMeshPointQuery1183) SetBtType(v string) {
	o.BtType = &v
}

func (o BTMMeshPointQuery1183) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMIndividualQueryWithOccurrenceBase904, errBTMIndividualQueryWithOccurrenceBase904 := json.Marshal(o.BTMIndividualQueryWithOccurrenceBase904)
	if errBTMIndividualQueryWithOccurrenceBase904 != nil {
		return []byte{}, errBTMIndividualQueryWithOccurrenceBase904
	}
	errBTMIndividualQueryWithOccurrenceBase904 = json.Unmarshal([]byte(serializedBTMIndividualQueryWithOccurrenceBase904), &toSerialize)
	if errBTMIndividualQueryWithOccurrenceBase904 != nil {
		return []byte{}, errBTMIndividualQueryWithOccurrenceBase904
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
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
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryString != nil {
		toSerialize["queryString"] = o.QueryString
	}
	if o.FullPathAsString != nil {
		toSerialize["fullPathAsString"] = o.FullPathAsString
	}
	if o.Occurrence != nil {
		toSerialize["occurrence"] = o.Occurrence
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTMMeshPointQuery1183 struct {
	value *BTMMeshPointQuery1183
	isSet bool
}

func (v NullableBTMMeshPointQuery1183) Get() *BTMMeshPointQuery1183 {
	return v.value
}

func (v *NullableBTMMeshPointQuery1183) Set(val *BTMMeshPointQuery1183) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMMeshPointQuery1183) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMMeshPointQuery1183) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMMeshPointQuery1183(val *BTMMeshPointQuery1183) *NullableBTMMeshPointQuery1183 {
	return &NullableBTMMeshPointQuery1183{value: val, isSet: true}
}

func (v NullableBTMMeshPointQuery1183) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMMeshPointQuery1183) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
