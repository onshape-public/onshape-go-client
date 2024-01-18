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

// Entry struct for Entry
type Entry struct {
	PermissionSet []string        `json:"permissionSet,omitempty"`
	Role          *BTRbacRoleInfo `json:"role,omitempty"`
}

// NewEntry instantiates a new Entry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntry() *Entry {
	this := Entry{}
	return &this
}

// NewEntryWithDefaults instantiates a new Entry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryWithDefaults() *Entry {
	this := Entry{}
	return &this
}

// GetPermissionSet returns the PermissionSet field value if set, zero value otherwise.
func (o *Entry) GetPermissionSet() []string {
	if o == nil || o.PermissionSet == nil {
		var ret []string
		return ret
	}
	return o.PermissionSet
}

// GetPermissionSetOk returns a tuple with the PermissionSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entry) GetPermissionSetOk() ([]string, bool) {
	if o == nil || o.PermissionSet == nil {
		return nil, false
	}
	return o.PermissionSet, true
}

// HasPermissionSet returns a boolean if a field has been set.
func (o *Entry) HasPermissionSet() bool {
	if o != nil && o.PermissionSet != nil {
		return true
	}

	return false
}

// SetPermissionSet gets a reference to the given []string and assigns it to the PermissionSet field.
func (o *Entry) SetPermissionSet(v []string) {
	o.PermissionSet = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *Entry) GetRole() BTRbacRoleInfo {
	if o == nil || o.Role == nil {
		var ret BTRbacRoleInfo
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entry) GetRoleOk() (*BTRbacRoleInfo, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *Entry) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given BTRbacRoleInfo and assigns it to the Role field.
func (o *Entry) SetRole(v BTRbacRoleInfo) {
	o.Role = &v
}

func (o Entry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PermissionSet != nil {
		toSerialize["permissionSet"] = o.PermissionSet
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableEntry struct {
	value *Entry
	isSet bool
}

func (v NullableEntry) Get() *Entry {
	return v.value
}

func (v *NullableEntry) Set(val *Entry) {
	v.value = val
	v.isSet = true
}

func (v NullableEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntry(val *Entry) *NullableEntry {
	return &NullableEntry{value: val, isSet: true}
}

func (v NullableEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
