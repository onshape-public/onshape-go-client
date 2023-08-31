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

// BTFSTableColumnInfo623 struct for BTFSTableColumnInfo623
type BTFSTableColumnInfo623 struct {
	BtType             *string                            `json:"btType,omitempty"`
	Id                 *string                            `json:"id,omitempty"`
	NodeId             *string                            `json:"nodeId,omitempty"`
	Specification      *BTTableColumnSpec1967             `json:"specification,omitempty"`
	CrossHighlightData *BTTableBaseCrossHighlightData2609 `json:"crossHighlightData,omitempty"`
}

// NewBTFSTableColumnInfo623 instantiates a new BTFSTableColumnInfo623 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSTableColumnInfo623() *BTFSTableColumnInfo623 {
	this := BTFSTableColumnInfo623{}
	return &this
}

// NewBTFSTableColumnInfo623WithDefaults instantiates a new BTFSTableColumnInfo623 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSTableColumnInfo623WithDefaults() *BTFSTableColumnInfo623 {
	this := BTFSTableColumnInfo623{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFSTableColumnInfo623) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSTableColumnInfo623) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFSTableColumnInfo623) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFSTableColumnInfo623) SetBtType(v string) {
	o.BtType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTFSTableColumnInfo623) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSTableColumnInfo623) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTFSTableColumnInfo623) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTFSTableColumnInfo623) SetId(v string) {
	o.Id = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTFSTableColumnInfo623) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSTableColumnInfo623) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTFSTableColumnInfo623) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTFSTableColumnInfo623) SetNodeId(v string) {
	o.NodeId = &v
}

// GetSpecification returns the Specification field value if set, zero value otherwise.
func (o *BTFSTableColumnInfo623) GetSpecification() BTTableColumnSpec1967 {
	if o == nil || o.Specification == nil {
		var ret BTTableColumnSpec1967
		return ret
	}
	return *o.Specification
}

// GetSpecificationOk returns a tuple with the Specification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSTableColumnInfo623) GetSpecificationOk() (*BTTableColumnSpec1967, bool) {
	if o == nil || o.Specification == nil {
		return nil, false
	}
	return o.Specification, true
}

// HasSpecification returns a boolean if a field has been set.
func (o *BTFSTableColumnInfo623) HasSpecification() bool {
	if o != nil && o.Specification != nil {
		return true
	}

	return false
}

// SetSpecification gets a reference to the given BTTableColumnSpec1967 and assigns it to the Specification field.
func (o *BTFSTableColumnInfo623) SetSpecification(v BTTableColumnSpec1967) {
	o.Specification = &v
}

// GetCrossHighlightData returns the CrossHighlightData field value if set, zero value otherwise.
func (o *BTFSTableColumnInfo623) GetCrossHighlightData() BTTableBaseCrossHighlightData2609 {
	if o == nil || o.CrossHighlightData == nil {
		var ret BTTableBaseCrossHighlightData2609
		return ret
	}
	return *o.CrossHighlightData
}

// GetCrossHighlightDataOk returns a tuple with the CrossHighlightData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSTableColumnInfo623) GetCrossHighlightDataOk() (*BTTableBaseCrossHighlightData2609, bool) {
	if o == nil || o.CrossHighlightData == nil {
		return nil, false
	}
	return o.CrossHighlightData, true
}

// HasCrossHighlightData returns a boolean if a field has been set.
func (o *BTFSTableColumnInfo623) HasCrossHighlightData() bool {
	if o != nil && o.CrossHighlightData != nil {
		return true
	}

	return false
}

// SetCrossHighlightData gets a reference to the given BTTableBaseCrossHighlightData2609 and assigns it to the CrossHighlightData field.
func (o *BTFSTableColumnInfo623) SetCrossHighlightData(v BTTableBaseCrossHighlightData2609) {
	o.CrossHighlightData = &v
}

func (o BTFSTableColumnInfo623) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Specification != nil {
		toSerialize["specification"] = o.Specification
	}
	if o.CrossHighlightData != nil {
		toSerialize["crossHighlightData"] = o.CrossHighlightData
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSTableColumnInfo623 struct {
	value *BTFSTableColumnInfo623
	isSet bool
}

func (v NullableBTFSTableColumnInfo623) Get() *BTFSTableColumnInfo623 {
	return v.value
}

func (v *NullableBTFSTableColumnInfo623) Set(val *BTFSTableColumnInfo623) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSTableColumnInfo623) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSTableColumnInfo623) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSTableColumnInfo623(val *BTFSTableColumnInfo623) *NullableBTFSTableColumnInfo623 {
	return &NullableBTFSTableColumnInfo623{value: val, isSet: true}
}

func (v NullableBTFSTableColumnInfo623) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSTableColumnInfo623) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
