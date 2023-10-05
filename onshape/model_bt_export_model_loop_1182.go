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

// BTExportModelLoop1182 struct for BTExportModelLoop1182
type BTExportModelLoop1182 struct {
	BtType  *string                   `json:"btType,omitempty"`
	Coedges []BTExportModelCoedge1342 `json:"coedges,omitempty"`
	IsInner *bool                     `json:"isInner,omitempty"`
	IsOuter *bool                     `json:"isOuter,omitempty"`
}

// NewBTExportModelLoop1182 instantiates a new BTExportModelLoop1182 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportModelLoop1182() *BTExportModelLoop1182 {
	this := BTExportModelLoop1182{}
	return &this
}

// NewBTExportModelLoop1182WithDefaults instantiates a new BTExportModelLoop1182 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportModelLoop1182WithDefaults() *BTExportModelLoop1182 {
	this := BTExportModelLoop1182{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExportModelLoop1182) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelLoop1182) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExportModelLoop1182) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExportModelLoop1182) SetBtType(v string) {
	o.BtType = &v
}

// GetCoedges returns the Coedges field value if set, zero value otherwise.
func (o *BTExportModelLoop1182) GetCoedges() []BTExportModelCoedge1342 {
	if o == nil || o.Coedges == nil {
		var ret []BTExportModelCoedge1342
		return ret
	}
	return o.Coedges
}

// GetCoedgesOk returns a tuple with the Coedges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelLoop1182) GetCoedgesOk() ([]BTExportModelCoedge1342, bool) {
	if o == nil || o.Coedges == nil {
		return nil, false
	}
	return o.Coedges, true
}

// HasCoedges returns a boolean if a field has been set.
func (o *BTExportModelLoop1182) HasCoedges() bool {
	if o != nil && o.Coedges != nil {
		return true
	}

	return false
}

// SetCoedges gets a reference to the given []BTExportModelCoedge1342 and assigns it to the Coedges field.
func (o *BTExportModelLoop1182) SetCoedges(v []BTExportModelCoedge1342) {
	o.Coedges = v
}

// GetIsInner returns the IsInner field value if set, zero value otherwise.
func (o *BTExportModelLoop1182) GetIsInner() bool {
	if o == nil || o.IsInner == nil {
		var ret bool
		return ret
	}
	return *o.IsInner
}

// GetIsInnerOk returns a tuple with the IsInner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelLoop1182) GetIsInnerOk() (*bool, bool) {
	if o == nil || o.IsInner == nil {
		return nil, false
	}
	return o.IsInner, true
}

// HasIsInner returns a boolean if a field has been set.
func (o *BTExportModelLoop1182) HasIsInner() bool {
	if o != nil && o.IsInner != nil {
		return true
	}

	return false
}

// SetIsInner gets a reference to the given bool and assigns it to the IsInner field.
func (o *BTExportModelLoop1182) SetIsInner(v bool) {
	o.IsInner = &v
}

// GetIsOuter returns the IsOuter field value if set, zero value otherwise.
func (o *BTExportModelLoop1182) GetIsOuter() bool {
	if o == nil || o.IsOuter == nil {
		var ret bool
		return ret
	}
	return *o.IsOuter
}

// GetIsOuterOk returns a tuple with the IsOuter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelLoop1182) GetIsOuterOk() (*bool, bool) {
	if o == nil || o.IsOuter == nil {
		return nil, false
	}
	return o.IsOuter, true
}

// HasIsOuter returns a boolean if a field has been set.
func (o *BTExportModelLoop1182) HasIsOuter() bool {
	if o != nil && o.IsOuter != nil {
		return true
	}

	return false
}

// SetIsOuter gets a reference to the given bool and assigns it to the IsOuter field.
func (o *BTExportModelLoop1182) SetIsOuter(v bool) {
	o.IsOuter = &v
}

func (o BTExportModelLoop1182) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Coedges != nil {
		toSerialize["coedges"] = o.Coedges
	}
	if o.IsInner != nil {
		toSerialize["isInner"] = o.IsInner
	}
	if o.IsOuter != nil {
		toSerialize["isOuter"] = o.IsOuter
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportModelLoop1182 struct {
	value *BTExportModelLoop1182
	isSet bool
}

func (v NullableBTExportModelLoop1182) Get() *BTExportModelLoop1182 {
	return v.value
}

func (v *NullableBTExportModelLoop1182) Set(val *BTExportModelLoop1182) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportModelLoop1182) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportModelLoop1182) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportModelLoop1182(val *BTExportModelLoop1182) *NullableBTExportModelLoop1182 {
	return &NullableBTExportModelLoop1182{value: val, isSet: true}
}

func (v NullableBTExportModelLoop1182) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportModelLoop1182) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
