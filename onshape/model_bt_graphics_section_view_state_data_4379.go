/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6946-bb9367a9d0cc
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTGraphicsSectionViewStateData4379 struct for BTGraphicsSectionViewStateData4379
type BTGraphicsSectionViewStateData4379 struct {
	ElementId           *string                          `json:"elementId,omitempty"`
	IsExcluding         *bool                            `json:"isExcluding,omitempty"`
	SectionPlanes       []BTGraphicsSectionPlaneData1429 `json:"sectionPlanes,omitempty"`
	SelectionsToExclude []BTUiSelection1185              `json:"selectionsToExclude,omitempty"`
	SelectionsToInclude []BTUiSelection1185              `json:"selectionsToInclude,omitempty"`
}

// NewBTGraphicsSectionViewStateData4379 instantiates a new BTGraphicsSectionViewStateData4379 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTGraphicsSectionViewStateData4379() *BTGraphicsSectionViewStateData4379 {
	this := BTGraphicsSectionViewStateData4379{}
	return &this
}

// NewBTGraphicsSectionViewStateData4379WithDefaults instantiates a new BTGraphicsSectionViewStateData4379 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTGraphicsSectionViewStateData4379WithDefaults() *BTGraphicsSectionViewStateData4379 {
	this := BTGraphicsSectionViewStateData4379{}
	return &this
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTGraphicsSectionViewStateData4379) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGraphicsSectionViewStateData4379) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTGraphicsSectionViewStateData4379) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTGraphicsSectionViewStateData4379) SetElementId(v string) {
	o.ElementId = &v
}

// GetIsExcluding returns the IsExcluding field value if set, zero value otherwise.
func (o *BTGraphicsSectionViewStateData4379) GetIsExcluding() bool {
	if o == nil || o.IsExcluding == nil {
		var ret bool
		return ret
	}
	return *o.IsExcluding
}

// GetIsExcludingOk returns a tuple with the IsExcluding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGraphicsSectionViewStateData4379) GetIsExcludingOk() (*bool, bool) {
	if o == nil || o.IsExcluding == nil {
		return nil, false
	}
	return o.IsExcluding, true
}

// HasIsExcluding returns a boolean if a field has been set.
func (o *BTGraphicsSectionViewStateData4379) HasIsExcluding() bool {
	if o != nil && o.IsExcluding != nil {
		return true
	}

	return false
}

// SetIsExcluding gets a reference to the given bool and assigns it to the IsExcluding field.
func (o *BTGraphicsSectionViewStateData4379) SetIsExcluding(v bool) {
	o.IsExcluding = &v
}

// GetSectionPlanes returns the SectionPlanes field value if set, zero value otherwise.
func (o *BTGraphicsSectionViewStateData4379) GetSectionPlanes() []BTGraphicsSectionPlaneData1429 {
	if o == nil || o.SectionPlanes == nil {
		var ret []BTGraphicsSectionPlaneData1429
		return ret
	}
	return o.SectionPlanes
}

// GetSectionPlanesOk returns a tuple with the SectionPlanes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGraphicsSectionViewStateData4379) GetSectionPlanesOk() ([]BTGraphicsSectionPlaneData1429, bool) {
	if o == nil || o.SectionPlanes == nil {
		return nil, false
	}
	return o.SectionPlanes, true
}

// HasSectionPlanes returns a boolean if a field has been set.
func (o *BTGraphicsSectionViewStateData4379) HasSectionPlanes() bool {
	if o != nil && o.SectionPlanes != nil {
		return true
	}

	return false
}

// SetSectionPlanes gets a reference to the given []BTGraphicsSectionPlaneData1429 and assigns it to the SectionPlanes field.
func (o *BTGraphicsSectionViewStateData4379) SetSectionPlanes(v []BTGraphicsSectionPlaneData1429) {
	o.SectionPlanes = v
}

// GetSelectionsToExclude returns the SelectionsToExclude field value if set, zero value otherwise.
func (o *BTGraphicsSectionViewStateData4379) GetSelectionsToExclude() []BTUiSelection1185 {
	if o == nil || o.SelectionsToExclude == nil {
		var ret []BTUiSelection1185
		return ret
	}
	return o.SelectionsToExclude
}

// GetSelectionsToExcludeOk returns a tuple with the SelectionsToExclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGraphicsSectionViewStateData4379) GetSelectionsToExcludeOk() ([]BTUiSelection1185, bool) {
	if o == nil || o.SelectionsToExclude == nil {
		return nil, false
	}
	return o.SelectionsToExclude, true
}

// HasSelectionsToExclude returns a boolean if a field has been set.
func (o *BTGraphicsSectionViewStateData4379) HasSelectionsToExclude() bool {
	if o != nil && o.SelectionsToExclude != nil {
		return true
	}

	return false
}

// SetSelectionsToExclude gets a reference to the given []BTUiSelection1185 and assigns it to the SelectionsToExclude field.
func (o *BTGraphicsSectionViewStateData4379) SetSelectionsToExclude(v []BTUiSelection1185) {
	o.SelectionsToExclude = v
}

// GetSelectionsToInclude returns the SelectionsToInclude field value if set, zero value otherwise.
func (o *BTGraphicsSectionViewStateData4379) GetSelectionsToInclude() []BTUiSelection1185 {
	if o == nil || o.SelectionsToInclude == nil {
		var ret []BTUiSelection1185
		return ret
	}
	return o.SelectionsToInclude
}

// GetSelectionsToIncludeOk returns a tuple with the SelectionsToInclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGraphicsSectionViewStateData4379) GetSelectionsToIncludeOk() ([]BTUiSelection1185, bool) {
	if o == nil || o.SelectionsToInclude == nil {
		return nil, false
	}
	return o.SelectionsToInclude, true
}

// HasSelectionsToInclude returns a boolean if a field has been set.
func (o *BTGraphicsSectionViewStateData4379) HasSelectionsToInclude() bool {
	if o != nil && o.SelectionsToInclude != nil {
		return true
	}

	return false
}

// SetSelectionsToInclude gets a reference to the given []BTUiSelection1185 and assigns it to the SelectionsToInclude field.
func (o *BTGraphicsSectionViewStateData4379) SetSelectionsToInclude(v []BTUiSelection1185) {
	o.SelectionsToInclude = v
}

func (o BTGraphicsSectionViewStateData4379) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.IsExcluding != nil {
		toSerialize["isExcluding"] = o.IsExcluding
	}
	if o.SectionPlanes != nil {
		toSerialize["sectionPlanes"] = o.SectionPlanes
	}
	if o.SelectionsToExclude != nil {
		toSerialize["selectionsToExclude"] = o.SelectionsToExclude
	}
	if o.SelectionsToInclude != nil {
		toSerialize["selectionsToInclude"] = o.SelectionsToInclude
	}
	return json.Marshal(toSerialize)
}

type NullableBTGraphicsSectionViewStateData4379 struct {
	value *BTGraphicsSectionViewStateData4379
	isSet bool
}

func (v NullableBTGraphicsSectionViewStateData4379) Get() *BTGraphicsSectionViewStateData4379 {
	return v.value
}

func (v *NullableBTGraphicsSectionViewStateData4379) Set(val *BTGraphicsSectionViewStateData4379) {
	v.value = val
	v.isSet = true
}

func (v NullableBTGraphicsSectionViewStateData4379) IsSet() bool {
	return v.isSet
}

func (v *NullableBTGraphicsSectionViewStateData4379) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTGraphicsSectionViewStateData4379(val *BTGraphicsSectionViewStateData4379) *NullableBTGraphicsSectionViewStateData4379 {
	return &NullableBTGraphicsSectionViewStateData4379{value: val, isSet: true}
}

func (v NullableBTGraphicsSectionViewStateData4379) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTGraphicsSectionViewStateData4379) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
