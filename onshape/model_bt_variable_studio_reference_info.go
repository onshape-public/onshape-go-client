/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.163.15457-d8ebaa9b9e42
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTVariableStudioReferenceInfo List of variable studio references
type BTVariableStudioReferenceInfo struct {
	// Whether all variables in the referenced variable studio are included
	EntireVariableStudio *bool `json:"entireVariableStudio,omitempty"`
	// DocumentId of referenced variable studio, blank for intra-workspace references
	ReferenceDocumentId *string `json:"referenceDocumentId,omitempty"`
	// ElementId of referenced variable studio
	ReferenceElementId string `json:"referenceElementId"`
	// VersionId of referenced variable studio, blank for intra-workspace references
	ReferenceVersionId *string `json:"referenceVersionId,omitempty"`
	// Optional list of selected variables
	VariableNames []string `json:"variableNames,omitempty"`
}

// NewBTVariableStudioReferenceInfo instantiates a new BTVariableStudioReferenceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTVariableStudioReferenceInfo(referenceElementId string) *BTVariableStudioReferenceInfo {
	this := BTVariableStudioReferenceInfo{}
	this.ReferenceElementId = referenceElementId
	return &this
}

// NewBTVariableStudioReferenceInfoWithDefaults instantiates a new BTVariableStudioReferenceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTVariableStudioReferenceInfoWithDefaults() *BTVariableStudioReferenceInfo {
	this := BTVariableStudioReferenceInfo{}
	return &this
}

// GetEntireVariableStudio returns the EntireVariableStudio field value if set, zero value otherwise.
func (o *BTVariableStudioReferenceInfo) GetEntireVariableStudio() bool {
	if o == nil || o.EntireVariableStudio == nil {
		var ret bool
		return ret
	}
	return *o.EntireVariableStudio
}

// GetEntireVariableStudioOk returns a tuple with the EntireVariableStudio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableStudioReferenceInfo) GetEntireVariableStudioOk() (*bool, bool) {
	if o == nil || o.EntireVariableStudio == nil {
		return nil, false
	}
	return o.EntireVariableStudio, true
}

// HasEntireVariableStudio returns a boolean if a field has been set.
func (o *BTVariableStudioReferenceInfo) HasEntireVariableStudio() bool {
	if o != nil && o.EntireVariableStudio != nil {
		return true
	}

	return false
}

// SetEntireVariableStudio gets a reference to the given bool and assigns it to the EntireVariableStudio field.
func (o *BTVariableStudioReferenceInfo) SetEntireVariableStudio(v bool) {
	o.EntireVariableStudio = &v
}

// GetReferenceDocumentId returns the ReferenceDocumentId field value if set, zero value otherwise.
func (o *BTVariableStudioReferenceInfo) GetReferenceDocumentId() string {
	if o == nil || o.ReferenceDocumentId == nil {
		var ret string
		return ret
	}
	return *o.ReferenceDocumentId
}

// GetReferenceDocumentIdOk returns a tuple with the ReferenceDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableStudioReferenceInfo) GetReferenceDocumentIdOk() (*string, bool) {
	if o == nil || o.ReferenceDocumentId == nil {
		return nil, false
	}
	return o.ReferenceDocumentId, true
}

// HasReferenceDocumentId returns a boolean if a field has been set.
func (o *BTVariableStudioReferenceInfo) HasReferenceDocumentId() bool {
	if o != nil && o.ReferenceDocumentId != nil {
		return true
	}

	return false
}

// SetReferenceDocumentId gets a reference to the given string and assigns it to the ReferenceDocumentId field.
func (o *BTVariableStudioReferenceInfo) SetReferenceDocumentId(v string) {
	o.ReferenceDocumentId = &v
}

// GetReferenceElementId returns the ReferenceElementId field value
func (o *BTVariableStudioReferenceInfo) GetReferenceElementId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceElementId
}

// GetReferenceElementIdOk returns a tuple with the ReferenceElementId field value
// and a boolean to check if the value has been set.
func (o *BTVariableStudioReferenceInfo) GetReferenceElementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceElementId, true
}

// SetReferenceElementId sets field value
func (o *BTVariableStudioReferenceInfo) SetReferenceElementId(v string) {
	o.ReferenceElementId = v
}

// GetReferenceVersionId returns the ReferenceVersionId field value if set, zero value otherwise.
func (o *BTVariableStudioReferenceInfo) GetReferenceVersionId() string {
	if o == nil || o.ReferenceVersionId == nil {
		var ret string
		return ret
	}
	return *o.ReferenceVersionId
}

// GetReferenceVersionIdOk returns a tuple with the ReferenceVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableStudioReferenceInfo) GetReferenceVersionIdOk() (*string, bool) {
	if o == nil || o.ReferenceVersionId == nil {
		return nil, false
	}
	return o.ReferenceVersionId, true
}

// HasReferenceVersionId returns a boolean if a field has been set.
func (o *BTVariableStudioReferenceInfo) HasReferenceVersionId() bool {
	if o != nil && o.ReferenceVersionId != nil {
		return true
	}

	return false
}

// SetReferenceVersionId gets a reference to the given string and assigns it to the ReferenceVersionId field.
func (o *BTVariableStudioReferenceInfo) SetReferenceVersionId(v string) {
	o.ReferenceVersionId = &v
}

// GetVariableNames returns the VariableNames field value if set, zero value otherwise.
func (o *BTVariableStudioReferenceInfo) GetVariableNames() []string {
	if o == nil || o.VariableNames == nil {
		var ret []string
		return ret
	}
	return o.VariableNames
}

// GetVariableNamesOk returns a tuple with the VariableNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableStudioReferenceInfo) GetVariableNamesOk() ([]string, bool) {
	if o == nil || o.VariableNames == nil {
		return nil, false
	}
	return o.VariableNames, true
}

// HasVariableNames returns a boolean if a field has been set.
func (o *BTVariableStudioReferenceInfo) HasVariableNames() bool {
	if o != nil && o.VariableNames != nil {
		return true
	}

	return false
}

// SetVariableNames gets a reference to the given []string and assigns it to the VariableNames field.
func (o *BTVariableStudioReferenceInfo) SetVariableNames(v []string) {
	o.VariableNames = v
}

func (o BTVariableStudioReferenceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntireVariableStudio != nil {
		toSerialize["entireVariableStudio"] = o.EntireVariableStudio
	}
	if o.ReferenceDocumentId != nil {
		toSerialize["referenceDocumentId"] = o.ReferenceDocumentId
	}
	if true {
		toSerialize["referenceElementId"] = o.ReferenceElementId
	}
	if o.ReferenceVersionId != nil {
		toSerialize["referenceVersionId"] = o.ReferenceVersionId
	}
	if o.VariableNames != nil {
		toSerialize["variableNames"] = o.VariableNames
	}
	return json.Marshal(toSerialize)
}

type NullableBTVariableStudioReferenceInfo struct {
	value *BTVariableStudioReferenceInfo
	isSet bool
}

func (v NullableBTVariableStudioReferenceInfo) Get() *BTVariableStudioReferenceInfo {
	return v.value
}

func (v *NullableBTVariableStudioReferenceInfo) Set(val *BTVariableStudioReferenceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTVariableStudioReferenceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTVariableStudioReferenceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTVariableStudioReferenceInfo(val *BTVariableStudioReferenceInfo) *NullableBTVariableStudioReferenceInfo {
	return &NullableBTVariableStudioReferenceInfo{value: val, isSet: true}
}

func (v NullableBTVariableStudioReferenceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTVariableStudioReferenceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
