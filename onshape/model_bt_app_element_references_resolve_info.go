/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13883-0fccf6f231ad
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppElementReferencesResolveInfo struct for BTAppElementReferencesResolveInfo
type BTAppElementReferencesResolveInfo struct {
	ResolvedReferences     []BTAppElementReferenceResolveInfo `json:"resolvedReferences,omitempty"`
	UnresolvedReferenceIds []string                           `json:"unresolvedReferenceIds,omitempty"`
}

// NewBTAppElementReferencesResolveInfo instantiates a new BTAppElementReferencesResolveInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementReferencesResolveInfo() *BTAppElementReferencesResolveInfo {
	this := BTAppElementReferencesResolveInfo{}
	return &this
}

// NewBTAppElementReferencesResolveInfoWithDefaults instantiates a new BTAppElementReferencesResolveInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementReferencesResolveInfoWithDefaults() *BTAppElementReferencesResolveInfo {
	this := BTAppElementReferencesResolveInfo{}
	return &this
}

// GetResolvedReferences returns the ResolvedReferences field value if set, zero value otherwise.
func (o *BTAppElementReferencesResolveInfo) GetResolvedReferences() []BTAppElementReferenceResolveInfo {
	if o == nil || o.ResolvedReferences == nil {
		var ret []BTAppElementReferenceResolveInfo
		return ret
	}
	return o.ResolvedReferences
}

// GetResolvedReferencesOk returns a tuple with the ResolvedReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferencesResolveInfo) GetResolvedReferencesOk() ([]BTAppElementReferenceResolveInfo, bool) {
	if o == nil || o.ResolvedReferences == nil {
		return nil, false
	}
	return o.ResolvedReferences, true
}

// HasResolvedReferences returns a boolean if a field has been set.
func (o *BTAppElementReferencesResolveInfo) HasResolvedReferences() bool {
	if o != nil && o.ResolvedReferences != nil {
		return true
	}

	return false
}

// SetResolvedReferences gets a reference to the given []BTAppElementReferenceResolveInfo and assigns it to the ResolvedReferences field.
func (o *BTAppElementReferencesResolveInfo) SetResolvedReferences(v []BTAppElementReferenceResolveInfo) {
	o.ResolvedReferences = v
}

// GetUnresolvedReferenceIds returns the UnresolvedReferenceIds field value if set, zero value otherwise.
func (o *BTAppElementReferencesResolveInfo) GetUnresolvedReferenceIds() []string {
	if o == nil || o.UnresolvedReferenceIds == nil {
		var ret []string
		return ret
	}
	return o.UnresolvedReferenceIds
}

// GetUnresolvedReferenceIdsOk returns a tuple with the UnresolvedReferenceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferencesResolveInfo) GetUnresolvedReferenceIdsOk() ([]string, bool) {
	if o == nil || o.UnresolvedReferenceIds == nil {
		return nil, false
	}
	return o.UnresolvedReferenceIds, true
}

// HasUnresolvedReferenceIds returns a boolean if a field has been set.
func (o *BTAppElementReferencesResolveInfo) HasUnresolvedReferenceIds() bool {
	if o != nil && o.UnresolvedReferenceIds != nil {
		return true
	}

	return false
}

// SetUnresolvedReferenceIds gets a reference to the given []string and assigns it to the UnresolvedReferenceIds field.
func (o *BTAppElementReferencesResolveInfo) SetUnresolvedReferenceIds(v []string) {
	o.UnresolvedReferenceIds = v
}

func (o BTAppElementReferencesResolveInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResolvedReferences != nil {
		toSerialize["resolvedReferences"] = o.ResolvedReferences
	}
	if o.UnresolvedReferenceIds != nil {
		toSerialize["unresolvedReferenceIds"] = o.UnresolvedReferenceIds
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementReferencesResolveInfo struct {
	value *BTAppElementReferencesResolveInfo
	isSet bool
}

func (v NullableBTAppElementReferencesResolveInfo) Get() *BTAppElementReferencesResolveInfo {
	return v.value
}

func (v *NullableBTAppElementReferencesResolveInfo) Set(val *BTAppElementReferencesResolveInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementReferencesResolveInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementReferencesResolveInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementReferencesResolveInfo(val *BTAppElementReferencesResolveInfo) *NullableBTAppElementReferencesResolveInfo {
	return &NullableBTAppElementReferencesResolveInfo{value: val, isSet: true}
}

func (v NullableBTAppElementReferencesResolveInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementReferencesResolveInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
