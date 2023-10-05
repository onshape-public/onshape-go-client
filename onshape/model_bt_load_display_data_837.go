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

// BTLoadDisplayData837 struct for BTLoadDisplayData837
type BTLoadDisplayData837 struct {
	BtType                   *string                          `json:"btType,omitempty"`
	ComponentValues          *BTVector3d389                   `json:"componentValues,omitempty"`
	DirectionMateConnectorId *string                          `json:"directionMateConnectorId,omitempty"`
	FaceLoadDeterministicIds []string                         `json:"faceLoadDeterministicIds,omitempty"`
	Hidden                   *bool                            `json:"hidden,omitempty"`
	IsDerivedFeature         *bool                            `json:"isDerivedFeature,omitempty"`
	IsDirectionFlipped       *bool                            `json:"isDirectionFlipped,omitempty"`
	LoadType                 *GBTLoadType                     `json:"loadType,omitempty"`
	NodeId                   *string                          `json:"nodeId,omitempty"`
	Occurrence               *BTOccurrence74                  `json:"occurrence,omitempty"`
	OwnerOccurrence          *BTOccurrence74                  `json:"ownerOccurrence,omitempty"`
	Status                   *GBTAssemblyFeatureDisplayStatus `json:"status,omitempty"`
}

// NewBTLoadDisplayData837 instantiates a new BTLoadDisplayData837 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTLoadDisplayData837() *BTLoadDisplayData837 {
	this := BTLoadDisplayData837{}
	return &this
}

// NewBTLoadDisplayData837WithDefaults instantiates a new BTLoadDisplayData837 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTLoadDisplayData837WithDefaults() *BTLoadDisplayData837 {
	this := BTLoadDisplayData837{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTLoadDisplayData837) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoadDisplayData837) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTLoadDisplayData837) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTLoadDisplayData837) SetBtType(v string) {
	o.BtType = &v
}

// GetComponentValues returns the ComponentValues field value if set, zero value otherwise.
func (o *BTLoadDisplayData837) GetComponentValues() BTVector3d389 {
	if o == nil || o.ComponentValues == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.ComponentValues
}

// GetComponentValuesOk returns a tuple with the ComponentValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoadDisplayData837) GetComponentValuesOk() (*BTVector3d389, bool) {
	if o == nil || o.ComponentValues == nil {
		return nil, false
	}
	return o.ComponentValues, true
}

// HasComponentValues returns a boolean if a field has been set.
func (o *BTLoadDisplayData837) HasComponentValues() bool {
	if o != nil && o.ComponentValues != nil {
		return true
	}

	return false
}

// SetComponentValues gets a reference to the given BTVector3d389 and assigns it to the ComponentValues field.
func (o *BTLoadDisplayData837) SetComponentValues(v BTVector3d389) {
	o.ComponentValues = &v
}

// GetDirectionMateConnectorId returns the DirectionMateConnectorId field value if set, zero value otherwise.
func (o *BTLoadDisplayData837) GetDirectionMateConnectorId() string {
	if o == nil || o.DirectionMateConnectorId == nil {
		var ret string
		return ret
	}
	return *o.DirectionMateConnectorId
}

// GetDirectionMateConnectorIdOk returns a tuple with the DirectionMateConnectorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoadDisplayData837) GetDirectionMateConnectorIdOk() (*string, bool) {
	if o == nil || o.DirectionMateConnectorId == nil {
		return nil, false
	}
	return o.DirectionMateConnectorId, true
}

// HasDirectionMateConnectorId returns a boolean if a field has been set.
func (o *BTLoadDisplayData837) HasDirectionMateConnectorId() bool {
	if o != nil && o.DirectionMateConnectorId != nil {
		return true
	}

	return false
}

// SetDirectionMateConnectorId gets a reference to the given string and assigns it to the DirectionMateConnectorId field.
func (o *BTLoadDisplayData837) SetDirectionMateConnectorId(v string) {
	o.DirectionMateConnectorId = &v
}

// GetFaceLoadDeterministicIds returns the FaceLoadDeterministicIds field value if set, zero value otherwise.
func (o *BTLoadDisplayData837) GetFaceLoadDeterministicIds() []string {
	if o == nil || o.FaceLoadDeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.FaceLoadDeterministicIds
}

// GetFaceLoadDeterministicIdsOk returns a tuple with the FaceLoadDeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoadDisplayData837) GetFaceLoadDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.FaceLoadDeterministicIds == nil {
		return nil, false
	}
	return o.FaceLoadDeterministicIds, true
}

// HasFaceLoadDeterministicIds returns a boolean if a field has been set.
func (o *BTLoadDisplayData837) HasFaceLoadDeterministicIds() bool {
	if o != nil && o.FaceLoadDeterministicIds != nil {
		return true
	}

	return false
}

// SetFaceLoadDeterministicIds gets a reference to the given []string and assigns it to the FaceLoadDeterministicIds field.
func (o *BTLoadDisplayData837) SetFaceLoadDeterministicIds(v []string) {
	o.FaceLoadDeterministicIds = v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *BTLoadDisplayData837) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoadDisplayData837) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *BTLoadDisplayData837) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *BTLoadDisplayData837) SetHidden(v bool) {
	o.Hidden = &v
}

// GetIsDerivedFeature returns the IsDerivedFeature field value if set, zero value otherwise.
func (o *BTLoadDisplayData837) GetIsDerivedFeature() bool {
	if o == nil || o.IsDerivedFeature == nil {
		var ret bool
		return ret
	}
	return *o.IsDerivedFeature
}

// GetIsDerivedFeatureOk returns a tuple with the IsDerivedFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoadDisplayData837) GetIsDerivedFeatureOk() (*bool, bool) {
	if o == nil || o.IsDerivedFeature == nil {
		return nil, false
	}
	return o.IsDerivedFeature, true
}

// HasIsDerivedFeature returns a boolean if a field has been set.
func (o *BTLoadDisplayData837) HasIsDerivedFeature() bool {
	if o != nil && o.IsDerivedFeature != nil {
		return true
	}

	return false
}

// SetIsDerivedFeature gets a reference to the given bool and assigns it to the IsDerivedFeature field.
func (o *BTLoadDisplayData837) SetIsDerivedFeature(v bool) {
	o.IsDerivedFeature = &v
}

// GetIsDirectionFlipped returns the IsDirectionFlipped field value if set, zero value otherwise.
func (o *BTLoadDisplayData837) GetIsDirectionFlipped() bool {
	if o == nil || o.IsDirectionFlipped == nil {
		var ret bool
		return ret
	}
	return *o.IsDirectionFlipped
}

// GetIsDirectionFlippedOk returns a tuple with the IsDirectionFlipped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoadDisplayData837) GetIsDirectionFlippedOk() (*bool, bool) {
	if o == nil || o.IsDirectionFlipped == nil {
		return nil, false
	}
	return o.IsDirectionFlipped, true
}

// HasIsDirectionFlipped returns a boolean if a field has been set.
func (o *BTLoadDisplayData837) HasIsDirectionFlipped() bool {
	if o != nil && o.IsDirectionFlipped != nil {
		return true
	}

	return false
}

// SetIsDirectionFlipped gets a reference to the given bool and assigns it to the IsDirectionFlipped field.
func (o *BTLoadDisplayData837) SetIsDirectionFlipped(v bool) {
	o.IsDirectionFlipped = &v
}

// GetLoadType returns the LoadType field value if set, zero value otherwise.
func (o *BTLoadDisplayData837) GetLoadType() GBTLoadType {
	if o == nil || o.LoadType == nil {
		var ret GBTLoadType
		return ret
	}
	return *o.LoadType
}

// GetLoadTypeOk returns a tuple with the LoadType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoadDisplayData837) GetLoadTypeOk() (*GBTLoadType, bool) {
	if o == nil || o.LoadType == nil {
		return nil, false
	}
	return o.LoadType, true
}

// HasLoadType returns a boolean if a field has been set.
func (o *BTLoadDisplayData837) HasLoadType() bool {
	if o != nil && o.LoadType != nil {
		return true
	}

	return false
}

// SetLoadType gets a reference to the given GBTLoadType and assigns it to the LoadType field.
func (o *BTLoadDisplayData837) SetLoadType(v GBTLoadType) {
	o.LoadType = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTLoadDisplayData837) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoadDisplayData837) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTLoadDisplayData837) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTLoadDisplayData837) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOccurrence returns the Occurrence field value if set, zero value otherwise.
func (o *BTLoadDisplayData837) GetOccurrence() BTOccurrence74 {
	if o == nil || o.Occurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.Occurrence
}

// GetOccurrenceOk returns a tuple with the Occurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoadDisplayData837) GetOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.Occurrence == nil {
		return nil, false
	}
	return o.Occurrence, true
}

// HasOccurrence returns a boolean if a field has been set.
func (o *BTLoadDisplayData837) HasOccurrence() bool {
	if o != nil && o.Occurrence != nil {
		return true
	}

	return false
}

// SetOccurrence gets a reference to the given BTOccurrence74 and assigns it to the Occurrence field.
func (o *BTLoadDisplayData837) SetOccurrence(v BTOccurrence74) {
	o.Occurrence = &v
}

// GetOwnerOccurrence returns the OwnerOccurrence field value if set, zero value otherwise.
func (o *BTLoadDisplayData837) GetOwnerOccurrence() BTOccurrence74 {
	if o == nil || o.OwnerOccurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.OwnerOccurrence
}

// GetOwnerOccurrenceOk returns a tuple with the OwnerOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoadDisplayData837) GetOwnerOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.OwnerOccurrence == nil {
		return nil, false
	}
	return o.OwnerOccurrence, true
}

// HasOwnerOccurrence returns a boolean if a field has been set.
func (o *BTLoadDisplayData837) HasOwnerOccurrence() bool {
	if o != nil && o.OwnerOccurrence != nil {
		return true
	}

	return false
}

// SetOwnerOccurrence gets a reference to the given BTOccurrence74 and assigns it to the OwnerOccurrence field.
func (o *BTLoadDisplayData837) SetOwnerOccurrence(v BTOccurrence74) {
	o.OwnerOccurrence = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BTLoadDisplayData837) GetStatus() GBTAssemblyFeatureDisplayStatus {
	if o == nil || o.Status == nil {
		var ret GBTAssemblyFeatureDisplayStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoadDisplayData837) GetStatusOk() (*GBTAssemblyFeatureDisplayStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BTLoadDisplayData837) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GBTAssemblyFeatureDisplayStatus and assigns it to the Status field.
func (o *BTLoadDisplayData837) SetStatus(v GBTAssemblyFeatureDisplayStatus) {
	o.Status = &v
}

func (o BTLoadDisplayData837) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ComponentValues != nil {
		toSerialize["componentValues"] = o.ComponentValues
	}
	if o.DirectionMateConnectorId != nil {
		toSerialize["directionMateConnectorId"] = o.DirectionMateConnectorId
	}
	if o.FaceLoadDeterministicIds != nil {
		toSerialize["faceLoadDeterministicIds"] = o.FaceLoadDeterministicIds
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.IsDerivedFeature != nil {
		toSerialize["isDerivedFeature"] = o.IsDerivedFeature
	}
	if o.IsDirectionFlipped != nil {
		toSerialize["isDirectionFlipped"] = o.IsDirectionFlipped
	}
	if o.LoadType != nil {
		toSerialize["loadType"] = o.LoadType
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Occurrence != nil {
		toSerialize["occurrence"] = o.Occurrence
	}
	if o.OwnerOccurrence != nil {
		toSerialize["ownerOccurrence"] = o.OwnerOccurrence
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableBTLoadDisplayData837 struct {
	value *BTLoadDisplayData837
	isSet bool
}

func (v NullableBTLoadDisplayData837) Get() *BTLoadDisplayData837 {
	return v.value
}

func (v *NullableBTLoadDisplayData837) Set(val *BTLoadDisplayData837) {
	v.value = val
	v.isSet = true
}

func (v NullableBTLoadDisplayData837) IsSet() bool {
	return v.isSet
}

func (v *NullableBTLoadDisplayData837) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTLoadDisplayData837(val *BTLoadDisplayData837) *NullableBTLoadDisplayData837 {
	return &NullableBTLoadDisplayData837{value: val, isSet: true}
}

func (v NullableBTLoadDisplayData837) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTLoadDisplayData837) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
