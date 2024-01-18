/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.175.29152-5591e93bd1c1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTGlobalTreeNodeListResponseBTTeamInfo struct for BTGlobalTreeNodeListResponseBTTeamInfo
type BTGlobalTreeNodeListResponseBTTeamInfo struct {
	// Requested Document URL
	Href *string `json:"href,omitempty"`
	// Document Items array. Array entries are the same as that returned from \"/api/documents/{did}\".
	Items []BTTeamInfo `json:"items,omitempty"`
	// The URL for the next page of items. Responses are limited to 20 items per page.
	Next *string `json:"next,omitempty"`
	// The URL for the previous page of items. Responses are limited to 20 items per page.
	Previous *string `json:"previous,omitempty"`
}

// NewBTGlobalTreeNodeListResponseBTTeamInfo instantiates a new BTGlobalTreeNodeListResponseBTTeamInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTGlobalTreeNodeListResponseBTTeamInfo() *BTGlobalTreeNodeListResponseBTTeamInfo {
	this := BTGlobalTreeNodeListResponseBTTeamInfo{}
	return &this
}

// NewBTGlobalTreeNodeListResponseBTTeamInfoWithDefaults instantiates a new BTGlobalTreeNodeListResponseBTTeamInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTGlobalTreeNodeListResponseBTTeamInfoWithDefaults() *BTGlobalTreeNodeListResponseBTTeamInfo {
	this := BTGlobalTreeNodeListResponseBTTeamInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) SetHref(v string) {
	o.Href = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) GetItems() []BTTeamInfo {
	if o == nil || o.Items == nil {
		var ret []BTTeamInfo
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) GetItemsOk() ([]BTTeamInfo, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BTTeamInfo and assigns it to the Items field.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) SetItems(v []BTTeamInfo) {
	o.Items = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *BTGlobalTreeNodeListResponseBTTeamInfo) SetPrevious(v string) {
	o.Previous = &v
}

func (o BTGlobalTreeNodeListResponseBTTeamInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	return json.Marshal(toSerialize)
}

type NullableBTGlobalTreeNodeListResponseBTTeamInfo struct {
	value *BTGlobalTreeNodeListResponseBTTeamInfo
	isSet bool
}

func (v NullableBTGlobalTreeNodeListResponseBTTeamInfo) Get() *BTGlobalTreeNodeListResponseBTTeamInfo {
	return v.value
}

func (v *NullableBTGlobalTreeNodeListResponseBTTeamInfo) Set(val *BTGlobalTreeNodeListResponseBTTeamInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTGlobalTreeNodeListResponseBTTeamInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTGlobalTreeNodeListResponseBTTeamInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTGlobalTreeNodeListResponseBTTeamInfo(val *BTGlobalTreeNodeListResponseBTTeamInfo) *NullableBTGlobalTreeNodeListResponseBTTeamInfo {
	return &NullableBTGlobalTreeNodeListResponseBTTeamInfo{value: val, isSet: true}
}

func (v NullableBTGlobalTreeNodeListResponseBTTeamInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTGlobalTreeNodeListResponseBTTeamInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
