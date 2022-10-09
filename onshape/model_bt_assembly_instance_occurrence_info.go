/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6826-13f16e212af4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyInstanceOccurrenceInfo struct for BTAssemblyInstanceOccurrenceInfo
type BTAssemblyInstanceOccurrenceInfo struct {
	Occurrences []BTAssemblyOccurrenceInfo `json:"occurrences,omitempty"`
}

// NewBTAssemblyInstanceOccurrenceInfo instantiates a new BTAssemblyInstanceOccurrenceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyInstanceOccurrenceInfo() *BTAssemblyInstanceOccurrenceInfo {
	this := BTAssemblyInstanceOccurrenceInfo{}
	return &this
}

// NewBTAssemblyInstanceOccurrenceInfoWithDefaults instantiates a new BTAssemblyInstanceOccurrenceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyInstanceOccurrenceInfoWithDefaults() *BTAssemblyInstanceOccurrenceInfo {
	this := BTAssemblyInstanceOccurrenceInfo{}
	return &this
}

// GetOccurrences returns the Occurrences field value if set, zero value otherwise.
func (o *BTAssemblyInstanceOccurrenceInfo) GetOccurrences() []BTAssemblyOccurrenceInfo {
	if o == nil || o.Occurrences == nil {
		var ret []BTAssemblyOccurrenceInfo
		return ret
	}
	return o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceOccurrenceInfo) GetOccurrencesOk() ([]BTAssemblyOccurrenceInfo, bool) {
	if o == nil || o.Occurrences == nil {
		return nil, false
	}
	return o.Occurrences, true
}

// HasOccurrences returns a boolean if a field has been set.
func (o *BTAssemblyInstanceOccurrenceInfo) HasOccurrences() bool {
	if o != nil && o.Occurrences != nil {
		return true
	}

	return false
}

// SetOccurrences gets a reference to the given []BTAssemblyOccurrenceInfo and assigns it to the Occurrences field.
func (o *BTAssemblyInstanceOccurrenceInfo) SetOccurrences(v []BTAssemblyOccurrenceInfo) {
	o.Occurrences = v
}

func (o BTAssemblyInstanceOccurrenceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Occurrences != nil {
		toSerialize["occurrences"] = o.Occurrences
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyInstanceOccurrenceInfo struct {
	value *BTAssemblyInstanceOccurrenceInfo
	isSet bool
}

func (v NullableBTAssemblyInstanceOccurrenceInfo) Get() *BTAssemblyInstanceOccurrenceInfo {
	return v.value
}

func (v *NullableBTAssemblyInstanceOccurrenceInfo) Set(val *BTAssemblyInstanceOccurrenceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyInstanceOccurrenceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyInstanceOccurrenceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyInstanceOccurrenceInfo(val *BTAssemblyInstanceOccurrenceInfo) *NullableBTAssemblyInstanceOccurrenceInfo {
	return &NullableBTAssemblyInstanceOccurrenceInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyInstanceOccurrenceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyInstanceOccurrenceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
