/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6017-d06351faf6f2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAllowedMateTypeFilter1511 struct for BTAllowedMateTypeFilter1511
type BTAllowedMateTypeFilter1511 struct {
	BtType               *string  `json:"btType,omitempty"`
	RequireMateQueryData *bool    `json:"requireMateQueryData,omitempty"`
	TopLevelMateOnly     *bool    `json:"topLevelMateOnly,omitempty"`
	AllowedMateTypes     []string `json:"allowedMateTypes,omitempty"`
}

// NewBTAllowedMateTypeFilter1511 instantiates a new BTAllowedMateTypeFilter1511 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAllowedMateTypeFilter1511() *BTAllowedMateTypeFilter1511 {
	this := BTAllowedMateTypeFilter1511{}
	return &this
}

// NewBTAllowedMateTypeFilter1511WithDefaults instantiates a new BTAllowedMateTypeFilter1511 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAllowedMateTypeFilter1511WithDefaults() *BTAllowedMateTypeFilter1511 {
	this := BTAllowedMateTypeFilter1511{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTAllowedMateTypeFilter1511) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowedMateTypeFilter1511) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTAllowedMateTypeFilter1511) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTAllowedMateTypeFilter1511) SetBtType(v string) {
	o.BtType = &v
}

// GetRequireMateQueryData returns the RequireMateQueryData field value if set, zero value otherwise.
func (o *BTAllowedMateTypeFilter1511) GetRequireMateQueryData() bool {
	if o == nil || o.RequireMateQueryData == nil {
		var ret bool
		return ret
	}
	return *o.RequireMateQueryData
}

// GetRequireMateQueryDataOk returns a tuple with the RequireMateQueryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowedMateTypeFilter1511) GetRequireMateQueryDataOk() (*bool, bool) {
	if o == nil || o.RequireMateQueryData == nil {
		return nil, false
	}
	return o.RequireMateQueryData, true
}

// HasRequireMateQueryData returns a boolean if a field has been set.
func (o *BTAllowedMateTypeFilter1511) HasRequireMateQueryData() bool {
	if o != nil && o.RequireMateQueryData != nil {
		return true
	}

	return false
}

// SetRequireMateQueryData gets a reference to the given bool and assigns it to the RequireMateQueryData field.
func (o *BTAllowedMateTypeFilter1511) SetRequireMateQueryData(v bool) {
	o.RequireMateQueryData = &v
}

// GetTopLevelMateOnly returns the TopLevelMateOnly field value if set, zero value otherwise.
func (o *BTAllowedMateTypeFilter1511) GetTopLevelMateOnly() bool {
	if o == nil || o.TopLevelMateOnly == nil {
		var ret bool
		return ret
	}
	return *o.TopLevelMateOnly
}

// GetTopLevelMateOnlyOk returns a tuple with the TopLevelMateOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowedMateTypeFilter1511) GetTopLevelMateOnlyOk() (*bool, bool) {
	if o == nil || o.TopLevelMateOnly == nil {
		return nil, false
	}
	return o.TopLevelMateOnly, true
}

// HasTopLevelMateOnly returns a boolean if a field has been set.
func (o *BTAllowedMateTypeFilter1511) HasTopLevelMateOnly() bool {
	if o != nil && o.TopLevelMateOnly != nil {
		return true
	}

	return false
}

// SetTopLevelMateOnly gets a reference to the given bool and assigns it to the TopLevelMateOnly field.
func (o *BTAllowedMateTypeFilter1511) SetTopLevelMateOnly(v bool) {
	o.TopLevelMateOnly = &v
}

// GetAllowedMateTypes returns the AllowedMateTypes field value if set, zero value otherwise.
func (o *BTAllowedMateTypeFilter1511) GetAllowedMateTypes() []string {
	if o == nil || o.AllowedMateTypes == nil {
		var ret []string
		return ret
	}
	return o.AllowedMateTypes
}

// GetAllowedMateTypesOk returns a tuple with the AllowedMateTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowedMateTypeFilter1511) GetAllowedMateTypesOk() ([]string, bool) {
	if o == nil || o.AllowedMateTypes == nil {
		return nil, false
	}
	return o.AllowedMateTypes, true
}

// HasAllowedMateTypes returns a boolean if a field has been set.
func (o *BTAllowedMateTypeFilter1511) HasAllowedMateTypes() bool {
	if o != nil && o.AllowedMateTypes != nil {
		return true
	}

	return false
}

// SetAllowedMateTypes gets a reference to the given []string and assigns it to the AllowedMateTypes field.
func (o *BTAllowedMateTypeFilter1511) SetAllowedMateTypes(v []string) {
	o.AllowedMateTypes = v
}

func (o BTAllowedMateTypeFilter1511) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.RequireMateQueryData != nil {
		toSerialize["requireMateQueryData"] = o.RequireMateQueryData
	}
	if o.TopLevelMateOnly != nil {
		toSerialize["topLevelMateOnly"] = o.TopLevelMateOnly
	}
	if o.AllowedMateTypes != nil {
		toSerialize["allowedMateTypes"] = o.AllowedMateTypes
	}
	return json.Marshal(toSerialize)
}

type NullableBTAllowedMateTypeFilter1511 struct {
	value *BTAllowedMateTypeFilter1511
	isSet bool
}

func (v NullableBTAllowedMateTypeFilter1511) Get() *BTAllowedMateTypeFilter1511 {
	return v.value
}

func (v *NullableBTAllowedMateTypeFilter1511) Set(val *BTAllowedMateTypeFilter1511) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAllowedMateTypeFilter1511) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAllowedMateTypeFilter1511) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAllowedMateTypeFilter1511(val *BTAllowedMateTypeFilter1511) *NullableBTAllowedMateTypeFilter1511 {
	return &NullableBTAllowedMateTypeFilter1511{value: val, isSet: true}
}

func (v NullableBTAllowedMateTypeFilter1511) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAllowedMateTypeFilter1511) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
