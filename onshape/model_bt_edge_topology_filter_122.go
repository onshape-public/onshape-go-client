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

// BTEdgeTopologyFilter122 struct for BTEdgeTopologyFilter122
type BTEdgeTopologyFilter122 struct {
	BTQueryFilter183
	BtType         *string          `json:"btType,omitempty"`
	EdgeTopology   *GBTEdgeTopology `json:"edgeTopology,omitempty"`
	IsInternalEdge *bool            `json:"isInternalEdge,omitempty"`
	// Type of JSON object.
	BtType *string `json:"btType,omitempty"`
}

// NewBTEdgeTopologyFilter122 instantiates a new BTEdgeTopologyFilter122 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEdgeTopologyFilter122() *BTEdgeTopologyFilter122 {
	this := BTEdgeTopologyFilter122{}
	return &this
}

// NewBTEdgeTopologyFilter122WithDefaults instantiates a new BTEdgeTopologyFilter122 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEdgeTopologyFilter122WithDefaults() *BTEdgeTopologyFilter122 {
	this := BTEdgeTopologyFilter122{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTEdgeTopologyFilter122) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEdgeTopologyFilter122) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTEdgeTopologyFilter122) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTEdgeTopologyFilter122) SetBtType(v string) {
	o.BtType = &v
}

// GetEdgeTopology returns the EdgeTopology field value if set, zero value otherwise.
func (o *BTEdgeTopologyFilter122) GetEdgeTopology() GBTEdgeTopology {
	if o == nil || o.EdgeTopology == nil {
		var ret GBTEdgeTopology
		return ret
	}
	return *o.EdgeTopology
}

// GetEdgeTopologyOk returns a tuple with the EdgeTopology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEdgeTopologyFilter122) GetEdgeTopologyOk() (*GBTEdgeTopology, bool) {
	if o == nil || o.EdgeTopology == nil {
		return nil, false
	}
	return o.EdgeTopology, true
}

// HasEdgeTopology returns a boolean if a field has been set.
func (o *BTEdgeTopologyFilter122) HasEdgeTopology() bool {
	if o != nil && o.EdgeTopology != nil {
		return true
	}

	return false
}

// SetEdgeTopology gets a reference to the given GBTEdgeTopology and assigns it to the EdgeTopology field.
func (o *BTEdgeTopologyFilter122) SetEdgeTopology(v GBTEdgeTopology) {
	o.EdgeTopology = &v
}

// GetIsInternalEdge returns the IsInternalEdge field value if set, zero value otherwise.
func (o *BTEdgeTopologyFilter122) GetIsInternalEdge() bool {
	if o == nil || o.IsInternalEdge == nil {
		var ret bool
		return ret
	}
	return *o.IsInternalEdge
}

// GetIsInternalEdgeOk returns a tuple with the IsInternalEdge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEdgeTopologyFilter122) GetIsInternalEdgeOk() (*bool, bool) {
	if o == nil || o.IsInternalEdge == nil {
		return nil, false
	}
	return o.IsInternalEdge, true
}

// HasIsInternalEdge returns a boolean if a field has been set.
func (o *BTEdgeTopologyFilter122) HasIsInternalEdge() bool {
	if o != nil && o.IsInternalEdge != nil {
		return true
	}

	return false
}

// SetIsInternalEdge gets a reference to the given bool and assigns it to the IsInternalEdge field.
func (o *BTEdgeTopologyFilter122) SetIsInternalEdge(v bool) {
	o.IsInternalEdge = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTEdgeTopologyFilter122) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEdgeTopologyFilter122) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTEdgeTopologyFilter122) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTEdgeTopologyFilter122) SetBtType(v string) {
	o.BtType = &v
}

func (o BTEdgeTopologyFilter122) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTQueryFilter183, errBTQueryFilter183 := json.Marshal(o.BTQueryFilter183)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	errBTQueryFilter183 = json.Unmarshal([]byte(serializedBTQueryFilter183), &toSerialize)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.EdgeTopology != nil {
		toSerialize["edgeTopology"] = o.EdgeTopology
	}
	if o.IsInternalEdge != nil {
		toSerialize["isInternalEdge"] = o.IsInternalEdge
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTEdgeTopologyFilter122 struct {
	value *BTEdgeTopologyFilter122
	isSet bool
}

func (v NullableBTEdgeTopologyFilter122) Get() *BTEdgeTopologyFilter122 {
	return v.value
}

func (v *NullableBTEdgeTopologyFilter122) Set(val *BTEdgeTopologyFilter122) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEdgeTopologyFilter122) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEdgeTopologyFilter122) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEdgeTopologyFilter122(val *BTEdgeTopologyFilter122) *NullableBTEdgeTopologyFilter122 {
	return &NullableBTEdgeTopologyFilter122{value: val, isSet: true}
}

func (v NullableBTEdgeTopologyFilter122) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEdgeTopologyFilter122) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
