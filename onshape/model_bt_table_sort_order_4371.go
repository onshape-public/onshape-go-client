/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.166.18819-fa27aca4c021
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTTableSortOrder4371 struct for BTTableSortOrder4371
type BTTableSortOrder4371 struct {
	BtType          *string `json:"btType,omitempty"`
	IsAscending     *bool   `json:"isAscending,omitempty"`
	NodeId          *string `json:"nodeId,omitempty"`
	SortingColumnId *string `json:"sortingColumnId,omitempty"`
}

// NewBTTableSortOrder4371 instantiates a new BTTableSortOrder4371 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableSortOrder4371() *BTTableSortOrder4371 {
	this := BTTableSortOrder4371{}
	return &this
}

// NewBTTableSortOrder4371WithDefaults instantiates a new BTTableSortOrder4371 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableSortOrder4371WithDefaults() *BTTableSortOrder4371 {
	this := BTTableSortOrder4371{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableSortOrder4371) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableSortOrder4371) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTableSortOrder4371) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTableSortOrder4371) SetBtType(v string) {
	o.BtType = &v
}

// GetIsAscending returns the IsAscending field value if set, zero value otherwise.
func (o *BTTableSortOrder4371) GetIsAscending() bool {
	if o == nil || o.IsAscending == nil {
		var ret bool
		return ret
	}
	return *o.IsAscending
}

// GetIsAscendingOk returns a tuple with the IsAscending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableSortOrder4371) GetIsAscendingOk() (*bool, bool) {
	if o == nil || o.IsAscending == nil {
		return nil, false
	}
	return o.IsAscending, true
}

// HasIsAscending returns a boolean if a field has been set.
func (o *BTTableSortOrder4371) HasIsAscending() bool {
	if o != nil && o.IsAscending != nil {
		return true
	}

	return false
}

// SetIsAscending gets a reference to the given bool and assigns it to the IsAscending field.
func (o *BTTableSortOrder4371) SetIsAscending(v bool) {
	o.IsAscending = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTTableSortOrder4371) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableSortOrder4371) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTTableSortOrder4371) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTTableSortOrder4371) SetNodeId(v string) {
	o.NodeId = &v
}

// GetSortingColumnId returns the SortingColumnId field value if set, zero value otherwise.
func (o *BTTableSortOrder4371) GetSortingColumnId() string {
	if o == nil || o.SortingColumnId == nil {
		var ret string
		return ret
	}
	return *o.SortingColumnId
}

// GetSortingColumnIdOk returns a tuple with the SortingColumnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableSortOrder4371) GetSortingColumnIdOk() (*string, bool) {
	if o == nil || o.SortingColumnId == nil {
		return nil, false
	}
	return o.SortingColumnId, true
}

// HasSortingColumnId returns a boolean if a field has been set.
func (o *BTTableSortOrder4371) HasSortingColumnId() bool {
	if o != nil && o.SortingColumnId != nil {
		return true
	}

	return false
}

// SetSortingColumnId gets a reference to the given string and assigns it to the SortingColumnId field.
func (o *BTTableSortOrder4371) SetSortingColumnId(v string) {
	o.SortingColumnId = &v
}

func (o BTTableSortOrder4371) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsAscending != nil {
		toSerialize["isAscending"] = o.IsAscending
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.SortingColumnId != nil {
		toSerialize["sortingColumnId"] = o.SortingColumnId
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableSortOrder4371 struct {
	value *BTTableSortOrder4371
	isSet bool
}

func (v NullableBTTableSortOrder4371) Get() *BTTableSortOrder4371 {
	return v.value
}

func (v *NullableBTTableSortOrder4371) Set(val *BTTableSortOrder4371) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableSortOrder4371) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableSortOrder4371) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableSortOrder4371(val *BTTableSortOrder4371) *NullableBTTableSortOrder4371 {
	return &NullableBTTableSortOrder4371{value: val, isSet: true}
}

func (v NullableBTTableSortOrder4371) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableSortOrder4371) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
