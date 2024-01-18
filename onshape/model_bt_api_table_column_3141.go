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

// BTApiTableColumn3141 struct for BTApiTableColumn3141
type BTApiTableColumn3141 struct {
	BtType        *string                `json:"btType,omitempty"`
	Header        *string                `json:"header,omitempty"`
	Id            *string                `json:"id,omitempty"`
	TextAlignment *GBTTableTextAlignment `json:"textAlignment,omitempty"`
}

// NewBTApiTableColumn3141 instantiates a new BTApiTableColumn3141 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTApiTableColumn3141() *BTApiTableColumn3141 {
	this := BTApiTableColumn3141{}
	return &this
}

// NewBTApiTableColumn3141WithDefaults instantiates a new BTApiTableColumn3141 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTApiTableColumn3141WithDefaults() *BTApiTableColumn3141 {
	this := BTApiTableColumn3141{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTApiTableColumn3141) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTableColumn3141) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTApiTableColumn3141) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTApiTableColumn3141) SetBtType(v string) {
	o.BtType = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *BTApiTableColumn3141) GetHeader() string {
	if o == nil || o.Header == nil {
		var ret string
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTableColumn3141) GetHeaderOk() (*string, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *BTApiTableColumn3141) HasHeader() bool {
	if o != nil && o.Header != nil {
		return true
	}

	return false
}

// SetHeader gets a reference to the given string and assigns it to the Header field.
func (o *BTApiTableColumn3141) SetHeader(v string) {
	o.Header = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTApiTableColumn3141) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTableColumn3141) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTApiTableColumn3141) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTApiTableColumn3141) SetId(v string) {
	o.Id = &v
}

// GetTextAlignment returns the TextAlignment field value if set, zero value otherwise.
func (o *BTApiTableColumn3141) GetTextAlignment() GBTTableTextAlignment {
	if o == nil || o.TextAlignment == nil {
		var ret GBTTableTextAlignment
		return ret
	}
	return *o.TextAlignment
}

// GetTextAlignmentOk returns a tuple with the TextAlignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTableColumn3141) GetTextAlignmentOk() (*GBTTableTextAlignment, bool) {
	if o == nil || o.TextAlignment == nil {
		return nil, false
	}
	return o.TextAlignment, true
}

// HasTextAlignment returns a boolean if a field has been set.
func (o *BTApiTableColumn3141) HasTextAlignment() bool {
	if o != nil && o.TextAlignment != nil {
		return true
	}

	return false
}

// SetTextAlignment gets a reference to the given GBTTableTextAlignment and assigns it to the TextAlignment field.
func (o *BTApiTableColumn3141) SetTextAlignment(v GBTTableTextAlignment) {
	o.TextAlignment = &v
}

func (o BTApiTableColumn3141) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Header != nil {
		toSerialize["header"] = o.Header
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TextAlignment != nil {
		toSerialize["textAlignment"] = o.TextAlignment
	}
	return json.Marshal(toSerialize)
}

type NullableBTApiTableColumn3141 struct {
	value *BTApiTableColumn3141
	isSet bool
}

func (v NullableBTApiTableColumn3141) Get() *BTApiTableColumn3141 {
	return v.value
}

func (v *NullableBTApiTableColumn3141) Set(val *BTApiTableColumn3141) {
	v.value = val
	v.isSet = true
}

func (v NullableBTApiTableColumn3141) IsSet() bool {
	return v.isSet
}

func (v *NullableBTApiTableColumn3141) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTApiTableColumn3141(val *BTApiTableColumn3141) *NullableBTApiTableColumn3141 {
	return &NullableBTApiTableColumn3141{value: val, isSet: true}
}

func (v NullableBTApiTableColumn3141) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTApiTableColumn3141) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
