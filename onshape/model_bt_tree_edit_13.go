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

// BTTreeEdit13 struct for BTTreeEdit13
type BTTreeEdit13 struct {
	BtType   *string        `json:"btType,omitempty"`
	EditType *EditType      `json:"editType,omitempty"`
	NewNodes []BTTreeNode20 `json:"newNodes,omitempty"`
	Nothing  *bool          `json:"nothing,omitempty"`
}

// NewBTTreeEdit13 instantiates a new BTTreeEdit13 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTreeEdit13() *BTTreeEdit13 {
	this := BTTreeEdit13{}
	return &this
}

// NewBTTreeEdit13WithDefaults instantiates a new BTTreeEdit13 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTreeEdit13WithDefaults() *BTTreeEdit13 {
	this := BTTreeEdit13{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTreeEdit13) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTreeEdit13) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTreeEdit13) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTreeEdit13) SetBtType(v string) {
	o.BtType = &v
}

// GetEditType returns the EditType field value if set, zero value otherwise.
func (o *BTTreeEdit13) GetEditType() EditType {
	if o == nil || o.EditType == nil {
		var ret EditType
		return ret
	}
	return *o.EditType
}

// GetEditTypeOk returns a tuple with the EditType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTreeEdit13) GetEditTypeOk() (*EditType, bool) {
	if o == nil || o.EditType == nil {
		return nil, false
	}
	return o.EditType, true
}

// HasEditType returns a boolean if a field has been set.
func (o *BTTreeEdit13) HasEditType() bool {
	if o != nil && o.EditType != nil {
		return true
	}

	return false
}

// SetEditType gets a reference to the given EditType and assigns it to the EditType field.
func (o *BTTreeEdit13) SetEditType(v EditType) {
	o.EditType = &v
}

// GetNewNodes returns the NewNodes field value if set, zero value otherwise.
func (o *BTTreeEdit13) GetNewNodes() []BTTreeNode20 {
	if o == nil || o.NewNodes == nil {
		var ret []BTTreeNode20
		return ret
	}
	return o.NewNodes
}

// GetNewNodesOk returns a tuple with the NewNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTreeEdit13) GetNewNodesOk() ([]BTTreeNode20, bool) {
	if o == nil || o.NewNodes == nil {
		return nil, false
	}
	return o.NewNodes, true
}

// HasNewNodes returns a boolean if a field has been set.
func (o *BTTreeEdit13) HasNewNodes() bool {
	if o != nil && o.NewNodes != nil {
		return true
	}

	return false
}

// SetNewNodes gets a reference to the given []BTTreeNode20 and assigns it to the NewNodes field.
func (o *BTTreeEdit13) SetNewNodes(v []BTTreeNode20) {
	o.NewNodes = v
}

// GetNothing returns the Nothing field value if set, zero value otherwise.
func (o *BTTreeEdit13) GetNothing() bool {
	if o == nil || o.Nothing == nil {
		var ret bool
		return ret
	}
	return *o.Nothing
}

// GetNothingOk returns a tuple with the Nothing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTreeEdit13) GetNothingOk() (*bool, bool) {
	if o == nil || o.Nothing == nil {
		return nil, false
	}
	return o.Nothing, true
}

// HasNothing returns a boolean if a field has been set.
func (o *BTTreeEdit13) HasNothing() bool {
	if o != nil && o.Nothing != nil {
		return true
	}

	return false
}

// SetNothing gets a reference to the given bool and assigns it to the Nothing field.
func (o *BTTreeEdit13) SetNothing(v bool) {
	o.Nothing = &v
}

func (o BTTreeEdit13) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.EditType != nil {
		toSerialize["editType"] = o.EditType
	}
	if o.NewNodes != nil {
		toSerialize["newNodes"] = o.NewNodes
	}
	if o.Nothing != nil {
		toSerialize["nothing"] = o.Nothing
	}
	return json.Marshal(toSerialize)
}

type NullableBTTreeEdit13 struct {
	value *BTTreeEdit13
	isSet bool
}

func (v NullableBTTreeEdit13) Get() *BTTreeEdit13 {
	return v.value
}

func (v *NullableBTTreeEdit13) Set(val *BTTreeEdit13) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTreeEdit13) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTreeEdit13) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTreeEdit13(val *BTTreeEdit13) *NullableBTTreeEdit13 {
	return &NullableBTTreeEdit13{value: val, isSet: true}
}

func (v NullableBTTreeEdit13) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTreeEdit13) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
