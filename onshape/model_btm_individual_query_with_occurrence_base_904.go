/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.163.15296-122c93d7dbb6
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTMIndividualQueryWithOccurrenceBase904 - struct for BTMIndividualQueryWithOccurrenceBase904
type BTMIndividualQueryWithOccurrenceBase904 struct {
	implBTMIndividualQueryWithOccurrenceBase904 interface{}
}

// BTMIndividualQueryWithOccurrence811AsBTMIndividualQueryWithOccurrenceBase904 is a convenience function that returns BTMIndividualQueryWithOccurrence811 wrapped in BTMIndividualQueryWithOccurrenceBase904
func (o *BTMIndividualQueryWithOccurrence811) AsBTMIndividualQueryWithOccurrenceBase904() *BTMIndividualQueryWithOccurrenceBase904 {
	return &BTMIndividualQueryWithOccurrenceBase904{o}
}

// BTMIndividualOccurrenceQuery626AsBTMIndividualQueryWithOccurrenceBase904 is a convenience function that returns BTMIndividualOccurrenceQuery626 wrapped in BTMIndividualQueryWithOccurrenceBase904
func (o *BTMIndividualOccurrenceQuery626) AsBTMIndividualQueryWithOccurrenceBase904() *BTMIndividualQueryWithOccurrenceBase904 {
	return &BTMIndividualQueryWithOccurrenceBase904{o}
}

// BTMMeshPointQuery1183AsBTMIndividualQueryWithOccurrenceBase904 is a convenience function that returns BTMMeshPointQuery1183 wrapped in BTMIndividualQueryWithOccurrenceBase904
func (o *BTMMeshPointQuery1183) AsBTMIndividualQueryWithOccurrenceBase904() *BTMIndividualQueryWithOccurrenceBase904 {
	return &BTMIndividualQueryWithOccurrenceBase904{o}
}

// BTMFeatureQueryWithOccurrence157AsBTMIndividualQueryWithOccurrenceBase904 is a convenience function that returns BTMFeatureQueryWithOccurrence157 wrapped in BTMIndividualQueryWithOccurrenceBase904
func (o *BTMFeatureQueryWithOccurrence157) AsBTMIndividualQueryWithOccurrenceBase904() *BTMIndividualQueryWithOccurrenceBase904 {
	return &BTMIndividualQueryWithOccurrenceBase904{o}
}

// BTMPartStudioMateConnectorQuery1324AsBTMIndividualQueryWithOccurrenceBase904 is a convenience function that returns BTMPartStudioMateConnectorQuery1324 wrapped in BTMIndividualQueryWithOccurrenceBase904
func (o *BTMPartStudioMateConnectorQuery1324) AsBTMIndividualQueryWithOccurrenceBase904() *BTMIndividualQueryWithOccurrenceBase904 {
	return &BTMIndividualQueryWithOccurrenceBase904{o}
}

// BTMInferenceQueryWithOccurrence1083AsBTMIndividualQueryWithOccurrenceBase904 is a convenience function that returns BTMInferenceQueryWithOccurrence1083 wrapped in BTMIndividualQueryWithOccurrenceBase904
func (o *BTMInferenceQueryWithOccurrence1083) AsBTMIndividualQueryWithOccurrenceBase904() *BTMIndividualQueryWithOccurrenceBase904 {
	return &BTMIndividualQueryWithOccurrenceBase904{o}
}

// NewBTMIndividualQueryWithOccurrenceBase904 instantiates a new BTMIndividualQueryWithOccurrenceBase904 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMIndividualQueryWithOccurrenceBase904() *BTMIndividualQueryWithOccurrenceBase904 {
	this := BTMIndividualQueryWithOccurrenceBase904{Newbase_BTMIndividualQueryWithOccurrenceBase904()}
	return &this
}

// NewBTMIndividualQueryWithOccurrenceBase904WithDefaults instantiates a new BTMIndividualQueryWithOccurrenceBase904 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMIndividualQueryWithOccurrenceBase904WithDefaults() *BTMIndividualQueryWithOccurrenceBase904 {
	this := BTMIndividualQueryWithOccurrenceBase904{Newbase_BTMIndividualQueryWithOccurrenceBase904WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetBtType() string {
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
func (o *BTMIndividualQueryWithOccurrenceBase904) GetBtTypeOk() (*string, bool) {
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
func (o *BTMIndividualQueryWithOccurrenceBase904) HasBtType() bool {
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
func (o *BTMIndividualQueryWithOccurrenceBase904) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetDeterministicIdList() BTMIndividualQueryBase139 {
	type getResult interface {
		GetDeterministicIdList() BTMIndividualQueryBase139
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIdList()
	} else {
		var de BTMIndividualQueryBase139
		return de
	}
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	type getResult interface {
		GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIdListOk()
	} else {
		return nil, false
	}
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrenceBase904) HasDeterministicIdList() bool {
	type getResult interface {
		HasDeterministicIdList() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDeterministicIdList()
	} else {
		return false
	}
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *BTMIndividualQueryWithOccurrenceBase904) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	type getResult interface {
		SetDeterministicIdList(v BTMIndividualQueryBase139)
	}

	o.GetActualInstance().(getResult).SetDeterministicIdList(v)
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetDeterministicIds() []string {
	type getResult interface {
		GetDeterministicIds() []string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIds()
	} else {
		var de []string
		return de
	}
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetDeterministicIdsOk() ([]string, bool) {
	type getResult interface {
		GetDeterministicIdsOk() ([]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIdsOk()
	} else {
		return nil, false
	}
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrenceBase904) HasDeterministicIds() bool {
	type getResult interface {
		HasDeterministicIds() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDeterministicIds()
	} else {
		return false
	}
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTMIndividualQueryWithOccurrenceBase904) SetDeterministicIds(v []string) {
	type getResult interface {
		SetDeterministicIds(v []string)
	}

	o.GetActualInstance().(getResult).SetDeterministicIds(v)
}

// GetFullPathAsString returns the FullPathAsString field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetFullPathAsString() string {
	type getResult interface {
		GetFullPathAsString() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFullPathAsString()
	} else {
		var de string
		return de
	}
}

// GetFullPathAsStringOk returns a tuple with the FullPathAsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetFullPathAsStringOk() (*string, bool) {
	type getResult interface {
		GetFullPathAsStringOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFullPathAsStringOk()
	} else {
		return nil, false
	}
}

// HasFullPathAsString returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrenceBase904) HasFullPathAsString() bool {
	type getResult interface {
		HasFullPathAsString() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasFullPathAsString()
	} else {
		return false
	}
}

// SetFullPathAsString gets a reference to the given string and assigns it to the FullPathAsString field.
func (o *BTMIndividualQueryWithOccurrenceBase904) SetFullPathAsString(v string) {
	type getResult interface {
		SetFullPathAsString(v string)
	}

	o.GetActualInstance().(getResult).SetFullPathAsString(v)
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetImportMicroversion() string {
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
func (o *BTMIndividualQueryWithOccurrenceBase904) GetImportMicroversionOk() (*string, bool) {
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
func (o *BTMIndividualQueryWithOccurrenceBase904) HasImportMicroversion() bool {
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
func (o *BTMIndividualQueryWithOccurrenceBase904) SetImportMicroversion(v string) {
	type getResult interface {
		SetImportMicroversion(v string)
	}

	o.GetActualInstance().(getResult).SetImportMicroversion(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetNodeId() string {
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
func (o *BTMIndividualQueryWithOccurrenceBase904) GetNodeIdOk() (*string, bool) {
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
func (o *BTMIndividualQueryWithOccurrenceBase904) HasNodeId() bool {
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
func (o *BTMIndividualQueryWithOccurrenceBase904) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetOccurrence returns the Occurrence field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetOccurrence() BTOccurrence74 {
	type getResult interface {
		GetOccurrence() BTOccurrence74
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOccurrence()
	} else {
		var de BTOccurrence74
		return de
	}
}

// GetOccurrenceOk returns a tuple with the Occurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetOccurrenceOk() (*BTOccurrence74, bool) {
	type getResult interface {
		GetOccurrenceOk() (*BTOccurrence74, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOccurrenceOk()
	} else {
		return nil, false
	}
}

// HasOccurrence returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrenceBase904) HasOccurrence() bool {
	type getResult interface {
		HasOccurrence() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasOccurrence()
	} else {
		return false
	}
}

// SetOccurrence gets a reference to the given BTOccurrence74 and assigns it to the Occurrence field.
func (o *BTMIndividualQueryWithOccurrenceBase904) SetOccurrence(v BTOccurrence74) {
	type getResult interface {
		SetOccurrence(v BTOccurrence74)
	}

	o.GetActualInstance().(getResult).SetOccurrence(v)
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetPath() []string {
	type getResult interface {
		GetPath() []string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPath()
	} else {
		var de []string
		return de
	}
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetPathOk() ([]string, bool) {
	type getResult interface {
		GetPathOk() ([]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPathOk()
	} else {
		return nil, false
	}
}

// HasPath returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrenceBase904) HasPath() bool {
	type getResult interface {
		HasPath() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasPath()
	} else {
		return false
	}
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *BTMIndividualQueryWithOccurrenceBase904) SetPath(v []string) {
	type getResult interface {
		SetPath(v []string)
	}

	o.GetActualInstance().(getResult).SetPath(v)
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetQuery() BTMIndividualQueryBase139 {
	type getResult interface {
		GetQuery() BTMIndividualQueryBase139
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQuery()
	} else {
		var de BTMIndividualQueryBase139
		return de
	}
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	type getResult interface {
		GetQueryOk() (*BTMIndividualQueryBase139, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQueryOk()
	} else {
		return nil, false
	}
}

// HasQuery returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrenceBase904) HasQuery() bool {
	type getResult interface {
		HasQuery() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasQuery()
	} else {
		return false
	}
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *BTMIndividualQueryWithOccurrenceBase904) SetQuery(v BTMIndividualQueryBase139) {
	type getResult interface {
		SetQuery(v BTMIndividualQueryBase139)
	}

	o.GetActualInstance().(getResult).SetQuery(v)
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetQueryString() string {
	type getResult interface {
		GetQueryString() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQueryString()
	} else {
		var de string
		return de
	}
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrenceBase904) GetQueryStringOk() (*string, bool) {
	type getResult interface {
		GetQueryStringOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQueryStringOk()
	} else {
		return nil, false
	}
}

// HasQueryString returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrenceBase904) HasQueryString() bool {
	type getResult interface {
		HasQueryString() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasQueryString()
	} else {
		return false
	}
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *BTMIndividualQueryWithOccurrenceBase904) SetQueryString(v string) {
	type getResult interface {
		SetQueryString(v string)
	}

	o.GetActualInstance().(getResult).SetQueryString(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMIndividualQueryWithOccurrenceBase904) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTMFeatureQueryWithOccurrence-157'
	if jsonDict["btType"] == "BTMFeatureQueryWithOccurrence-157" {
		// try to unmarshal JSON data into BTMFeatureQueryWithOccurrence157
		var qr *BTMFeatureQueryWithOccurrence157
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQueryWithOccurrenceBase904 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQueryWithOccurrenceBase904 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQueryWithOccurrenceBase904 as BTMFeatureQueryWithOccurrence157: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMIndividualOccurrenceQuery-626'
	if jsonDict["btType"] == "BTMIndividualOccurrenceQuery-626" {
		// try to unmarshal JSON data into BTMIndividualOccurrenceQuery626
		var qr *BTMIndividualOccurrenceQuery626
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQueryWithOccurrenceBase904 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQueryWithOccurrenceBase904 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQueryWithOccurrenceBase904 as BTMIndividualOccurrenceQuery626: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMIndividualQueryWithOccurrence-811'
	if jsonDict["btType"] == "BTMIndividualQueryWithOccurrence-811" {
		// try to unmarshal JSON data into BTMIndividualQueryWithOccurrence811
		var qr *BTMIndividualQueryWithOccurrence811
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQueryWithOccurrenceBase904 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQueryWithOccurrenceBase904 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQueryWithOccurrenceBase904 as BTMIndividualQueryWithOccurrence811: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMInferenceQueryWithOccurrence-1083'
	if jsonDict["btType"] == "BTMInferenceQueryWithOccurrence-1083" {
		// try to unmarshal JSON data into BTMInferenceQueryWithOccurrence1083
		var qr *BTMInferenceQueryWithOccurrence1083
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQueryWithOccurrenceBase904 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQueryWithOccurrenceBase904 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQueryWithOccurrenceBase904 as BTMInferenceQueryWithOccurrence1083: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMMeshPointQuery-1183'
	if jsonDict["btType"] == "BTMMeshPointQuery-1183" {
		// try to unmarshal JSON data into BTMMeshPointQuery1183
		var qr *BTMMeshPointQuery1183
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQueryWithOccurrenceBase904 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQueryWithOccurrenceBase904 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQueryWithOccurrenceBase904 as BTMMeshPointQuery1183: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMPartStudioMateConnectorQuery-1324'
	if jsonDict["btType"] == "BTMPartStudioMateConnectorQuery-1324" {
		// try to unmarshal JSON data into BTMPartStudioMateConnectorQuery1324
		var qr *BTMPartStudioMateConnectorQuery1324
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQueryWithOccurrenceBase904 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQueryWithOccurrenceBase904 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQueryWithOccurrenceBase904 as BTMPartStudioMateConnectorQuery1324: %s", err.Error())
		}
	}

	var qtx *base_BTMIndividualQueryWithOccurrenceBase904
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMIndividualQueryWithOccurrenceBase904 = qtx
		return nil // data stored in dst.base_BTMIndividualQueryWithOccurrenceBase904, return on the first match
	} else {
		dst.implBTMIndividualQueryWithOccurrenceBase904 = nil
		return fmt.Errorf("Failed to unmarshal BTMIndividualQueryWithOccurrenceBase904 as base_BTMIndividualQueryWithOccurrenceBase904: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMIndividualQueryWithOccurrenceBase904) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMIndividualQueryWithOccurrenceBase904) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMIndividualQueryWithOccurrenceBase904
}

type NullableBTMIndividualQueryWithOccurrenceBase904 struct {
	value *BTMIndividualQueryWithOccurrenceBase904
	isSet bool
}

func (v NullableBTMIndividualQueryWithOccurrenceBase904) Get() *BTMIndividualQueryWithOccurrenceBase904 {
	return v.value
}

func (v *NullableBTMIndividualQueryWithOccurrenceBase904) Set(val *BTMIndividualQueryWithOccurrenceBase904) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMIndividualQueryWithOccurrenceBase904) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMIndividualQueryWithOccurrenceBase904) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMIndividualQueryWithOccurrenceBase904(val *BTMIndividualQueryWithOccurrenceBase904) *NullableBTMIndividualQueryWithOccurrenceBase904 {
	return &NullableBTMIndividualQueryWithOccurrenceBase904{value: val, isSet: true}
}

func (v NullableBTMIndividualQueryWithOccurrenceBase904) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMIndividualQueryWithOccurrenceBase904) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMIndividualQueryWithOccurrenceBase904 struct {
	BtType              *string                    `json:"btType,omitempty"`
	DeterministicIdList *BTMIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds    []string                   `json:"deterministicIds,omitempty"`
	FullPathAsString    *string                    `json:"fullPathAsString,omitempty"`
	ImportMicroversion  *string                    `json:"importMicroversion,omitempty"`
	NodeId              *string                    `json:"nodeId,omitempty"`
	Occurrence          *BTOccurrence74            `json:"occurrence,omitempty"`
	Path                []string                   `json:"path,omitempty"`
	Query               *BTMIndividualQueryBase139 `json:"query,omitempty"`
	QueryString         *string                    `json:"queryString,omitempty"`
}

// Newbase_BTMIndividualQueryWithOccurrenceBase904 instantiates a new base_BTMIndividualQueryWithOccurrenceBase904 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMIndividualQueryWithOccurrenceBase904() *base_BTMIndividualQueryWithOccurrenceBase904 {
	this := base_BTMIndividualQueryWithOccurrenceBase904{}
	return &this
}

// Newbase_BTMIndividualQueryWithOccurrenceBase904WithDefaults instantiates a new base_BTMIndividualQueryWithOccurrenceBase904 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMIndividualQueryWithOccurrenceBase904WithDefaults() *base_BTMIndividualQueryWithOccurrenceBase904 {
	this := base_BTMIndividualQueryWithOccurrenceBase904{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetDeterministicIdList() BTMIndividualQueryBase139 {
	if o == nil || o.DeterministicIdList == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	o.DeterministicIdList = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetFullPathAsString returns the FullPathAsString field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetFullPathAsString() string {
	if o == nil || o.FullPathAsString == nil {
		var ret string
		return ret
	}
	return *o.FullPathAsString
}

// GetFullPathAsStringOk returns a tuple with the FullPathAsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetFullPathAsStringOk() (*string, bool) {
	if o == nil || o.FullPathAsString == nil {
		return nil, false
	}
	return o.FullPathAsString, true
}

// HasFullPathAsString returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) HasFullPathAsString() bool {
	if o != nil && o.FullPathAsString != nil {
		return true
	}

	return false
}

// SetFullPathAsString gets a reference to the given string and assigns it to the FullPathAsString field.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) SetFullPathAsString(v string) {
	o.FullPathAsString = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOccurrence returns the Occurrence field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetOccurrence() BTOccurrence74 {
	if o == nil || o.Occurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.Occurrence
}

// GetOccurrenceOk returns a tuple with the Occurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.Occurrence == nil {
		return nil, false
	}
	return o.Occurrence, true
}

// HasOccurrence returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) HasOccurrence() bool {
	if o != nil && o.Occurrence != nil {
		return true
	}

	return false
}

// SetOccurrence gets a reference to the given BTOccurrence74 and assigns it to the Occurrence field.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) SetOccurrence(v BTOccurrence74) {
	o.Occurrence = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetPath() []string {
	if o == nil || o.Path == nil {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetPathOk() ([]string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) SetPath(v []string) {
	o.Path = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetQuery() BTMIndividualQueryBase139 {
	if o == nil || o.Query == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) SetQuery(v BTMIndividualQueryBase139) {
	o.Query = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *base_BTMIndividualQueryWithOccurrenceBase904) SetQueryString(v string) {
	o.QueryString = &v
}

func (o base_BTMIndividualQueryWithOccurrenceBase904) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DeterministicIdList != nil {
		toSerialize["deterministicIdList"] = o.DeterministicIdList
	}
	if o.DeterministicIds != nil {
		toSerialize["deterministicIds"] = o.DeterministicIds
	}
	if o.FullPathAsString != nil {
		toSerialize["fullPathAsString"] = o.FullPathAsString
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Occurrence != nil {
		toSerialize["occurrence"] = o.Occurrence
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryString != nil {
		toSerialize["queryString"] = o.QueryString
	}
	return json.Marshal(toSerialize)
}
