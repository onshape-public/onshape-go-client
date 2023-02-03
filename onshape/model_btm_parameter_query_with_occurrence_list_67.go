/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.158.10882-bf44380b65e2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMParameterQueryWithOccurrenceList67 struct for BTMParameterQueryWithOccurrenceList67
type BTMParameterQueryWithOccurrenceList67 struct {
	BtType             *string                                   `json:"btType,omitempty"`
	ImportMicroversion *string                                   `json:"importMicroversion,omitempty"`
	NodeId             *string                                   `json:"nodeId,omitempty"`
	Occurrences        []BTOccurrence74                          `json:"occurrences,omitempty"`
	ParameterId        *string                                   `json:"parameterId,omitempty"`
	Queries            []BTMIndividualQueryWithOccurrenceBase904 `json:"queries,omitempty"`
}

// NewBTMParameterQueryWithOccurrenceList67 instantiates a new BTMParameterQueryWithOccurrenceList67 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterQueryWithOccurrenceList67() *BTMParameterQueryWithOccurrenceList67 {
	this := BTMParameterQueryWithOccurrenceList67{}
	return &this
}

// NewBTMParameterQueryWithOccurrenceList67WithDefaults instantiates a new BTMParameterQueryWithOccurrenceList67 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterQueryWithOccurrenceList67WithDefaults() *BTMParameterQueryWithOccurrenceList67 {
	this := BTMParameterQueryWithOccurrenceList67{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterQueryWithOccurrenceList67) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQueryWithOccurrenceList67) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterQueryWithOccurrenceList67) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterQueryWithOccurrenceList67) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterQueryWithOccurrenceList67) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQueryWithOccurrenceList67) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterQueryWithOccurrenceList67) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterQueryWithOccurrenceList67) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterQueryWithOccurrenceList67) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQueryWithOccurrenceList67) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameterQueryWithOccurrenceList67) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameterQueryWithOccurrenceList67) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOccurrences returns the Occurrences field value if set, zero value otherwise.
func (o *BTMParameterQueryWithOccurrenceList67) GetOccurrences() []BTOccurrence74 {
	if o == nil || o.Occurrences == nil {
		var ret []BTOccurrence74
		return ret
	}
	return o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQueryWithOccurrenceList67) GetOccurrencesOk() ([]BTOccurrence74, bool) {
	if o == nil || o.Occurrences == nil {
		return nil, false
	}
	return o.Occurrences, true
}

// HasOccurrences returns a boolean if a field has been set.
func (o *BTMParameterQueryWithOccurrenceList67) HasOccurrences() bool {
	if o != nil && o.Occurrences != nil {
		return true
	}

	return false
}

// SetOccurrences gets a reference to the given []BTOccurrence74 and assigns it to the Occurrences field.
func (o *BTMParameterQueryWithOccurrenceList67) SetOccurrences(v []BTOccurrence74) {
	o.Occurrences = v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterQueryWithOccurrenceList67) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQueryWithOccurrenceList67) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterQueryWithOccurrenceList67) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterQueryWithOccurrenceList67) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *BTMParameterQueryWithOccurrenceList67) GetQueries() []BTMIndividualQueryWithOccurrenceBase904 {
	if o == nil || o.Queries == nil {
		var ret []BTMIndividualQueryWithOccurrenceBase904
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQueryWithOccurrenceList67) GetQueriesOk() ([]BTMIndividualQueryWithOccurrenceBase904, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *BTMParameterQueryWithOccurrenceList67) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []BTMIndividualQueryWithOccurrenceBase904 and assigns it to the Queries field.
func (o *BTMParameterQueryWithOccurrenceList67) SetQueries(v []BTMIndividualQueryWithOccurrenceBase904) {
	o.Queries = v
}

func (o BTMParameterQueryWithOccurrenceList67) MarshalJSON() ([]byte, error) {
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
	if o.Occurrences != nil {
		toSerialize["occurrences"] = o.Occurrences
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterQueryWithOccurrenceList67 struct {
	value *BTMParameterQueryWithOccurrenceList67
	isSet bool
}

func (v NullableBTMParameterQueryWithOccurrenceList67) Get() *BTMParameterQueryWithOccurrenceList67 {
	return v.value
}

func (v *NullableBTMParameterQueryWithOccurrenceList67) Set(val *BTMParameterQueryWithOccurrenceList67) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterQueryWithOccurrenceList67) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterQueryWithOccurrenceList67) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterQueryWithOccurrenceList67(val *BTMParameterQueryWithOccurrenceList67) *NullableBTMParameterQueryWithOccurrenceList67 {
	return &NullableBTMParameterQueryWithOccurrenceList67{value: val, isSet: true}
}

func (v NullableBTMParameterQueryWithOccurrenceList67) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterQueryWithOccurrenceList67) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
