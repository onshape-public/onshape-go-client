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

// BTConfiguredFeatureColumnInfo1014 struct for BTConfiguredFeatureColumnInfo1014
type BTConfiguredFeatureColumnInfo1014 struct {
	BtType        *string                  `json:"btType,omitempty"`
	Id            *string                  `json:"id,omitempty"`
	NodeId        *string                  `json:"nodeId,omitempty"`
	Specification *BTTableColumnSpec1967   `json:"specification,omitempty"`
	ParentId      *string                  `json:"parentId,omitempty"`
	ParentName    *string                  `json:"parentName,omitempty"`
	ParentType    *GBTConfiguredParentType `json:"parentType,omitempty"`
}

// NewBTConfiguredFeatureColumnInfo1014 instantiates a new BTConfiguredFeatureColumnInfo1014 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfiguredFeatureColumnInfo1014() *BTConfiguredFeatureColumnInfo1014 {
	this := BTConfiguredFeatureColumnInfo1014{}
	return &this
}

// NewBTConfiguredFeatureColumnInfo1014WithDefaults instantiates a new BTConfiguredFeatureColumnInfo1014 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfiguredFeatureColumnInfo1014WithDefaults() *BTConfiguredFeatureColumnInfo1014 {
	this := BTConfiguredFeatureColumnInfo1014{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConfiguredFeatureColumnInfo1014) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredFeatureColumnInfo1014) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConfiguredFeatureColumnInfo1014) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConfiguredFeatureColumnInfo1014) SetBtType(v string) {
	o.BtType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTConfiguredFeatureColumnInfo1014) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredFeatureColumnInfo1014) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTConfiguredFeatureColumnInfo1014) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTConfiguredFeatureColumnInfo1014) SetId(v string) {
	o.Id = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTConfiguredFeatureColumnInfo1014) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredFeatureColumnInfo1014) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTConfiguredFeatureColumnInfo1014) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTConfiguredFeatureColumnInfo1014) SetNodeId(v string) {
	o.NodeId = &v
}

// GetSpecification returns the Specification field value if set, zero value otherwise.
func (o *BTConfiguredFeatureColumnInfo1014) GetSpecification() BTTableColumnSpec1967 {
	if o == nil || o.Specification == nil {
		var ret BTTableColumnSpec1967
		return ret
	}
	return *o.Specification
}

// GetSpecificationOk returns a tuple with the Specification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredFeatureColumnInfo1014) GetSpecificationOk() (*BTTableColumnSpec1967, bool) {
	if o == nil || o.Specification == nil {
		return nil, false
	}
	return o.Specification, true
}

// HasSpecification returns a boolean if a field has been set.
func (o *BTConfiguredFeatureColumnInfo1014) HasSpecification() bool {
	if o != nil && o.Specification != nil {
		return true
	}

	return false
}

// SetSpecification gets a reference to the given BTTableColumnSpec1967 and assigns it to the Specification field.
func (o *BTConfiguredFeatureColumnInfo1014) SetSpecification(v BTTableColumnSpec1967) {
	o.Specification = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTConfiguredFeatureColumnInfo1014) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredFeatureColumnInfo1014) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BTConfiguredFeatureColumnInfo1014) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTConfiguredFeatureColumnInfo1014) SetParentId(v string) {
	o.ParentId = &v
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *BTConfiguredFeatureColumnInfo1014) GetParentName() string {
	if o == nil || o.ParentName == nil {
		var ret string
		return ret
	}
	return *o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredFeatureColumnInfo1014) GetParentNameOk() (*string, bool) {
	if o == nil || o.ParentName == nil {
		return nil, false
	}
	return o.ParentName, true
}

// HasParentName returns a boolean if a field has been set.
func (o *BTConfiguredFeatureColumnInfo1014) HasParentName() bool {
	if o != nil && o.ParentName != nil {
		return true
	}

	return false
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *BTConfiguredFeatureColumnInfo1014) SetParentName(v string) {
	o.ParentName = &v
}

// GetParentType returns the ParentType field value if set, zero value otherwise.
func (o *BTConfiguredFeatureColumnInfo1014) GetParentType() GBTConfiguredParentType {
	if o == nil || o.ParentType == nil {
		var ret GBTConfiguredParentType
		return ret
	}
	return *o.ParentType
}

// GetParentTypeOk returns a tuple with the ParentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredFeatureColumnInfo1014) GetParentTypeOk() (*GBTConfiguredParentType, bool) {
	if o == nil || o.ParentType == nil {
		return nil, false
	}
	return o.ParentType, true
}

// HasParentType returns a boolean if a field has been set.
func (o *BTConfiguredFeatureColumnInfo1014) HasParentType() bool {
	if o != nil && o.ParentType != nil {
		return true
	}

	return false
}

// SetParentType gets a reference to the given GBTConfiguredParentType and assigns it to the ParentType field.
func (o *BTConfiguredFeatureColumnInfo1014) SetParentType(v GBTConfiguredParentType) {
	o.ParentType = &v
}

func (o BTConfiguredFeatureColumnInfo1014) MarshalJSON() ([]byte, error) {
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
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.ParentName != nil {
		toSerialize["parentName"] = o.ParentName
	}
	if o.ParentType != nil {
		toSerialize["parentType"] = o.ParentType
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfiguredFeatureColumnInfo1014 struct {
	value *BTConfiguredFeatureColumnInfo1014
	isSet bool
}

func (v NullableBTConfiguredFeatureColumnInfo1014) Get() *BTConfiguredFeatureColumnInfo1014 {
	return v.value
}

func (v *NullableBTConfiguredFeatureColumnInfo1014) Set(val *BTConfiguredFeatureColumnInfo1014) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfiguredFeatureColumnInfo1014) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfiguredFeatureColumnInfo1014) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfiguredFeatureColumnInfo1014(val *BTConfiguredFeatureColumnInfo1014) *NullableBTConfiguredFeatureColumnInfo1014 {
	return &NullableBTConfiguredFeatureColumnInfo1014{value: val, isSet: true}
}

func (v NullableBTConfiguredFeatureColumnInfo1014) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfiguredFeatureColumnInfo1014) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
