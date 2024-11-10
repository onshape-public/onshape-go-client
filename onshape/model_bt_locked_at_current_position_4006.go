/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTLockedAtCurrentPosition4006 struct for BTLockedAtCurrentPosition4006
type BTLockedAtCurrentPosition4006 struct {
	BTLockedSubAssembly4590
	BtType                      *string                                   `json:"btType,omitempty"`
	LockType                    *GBTSubAssemblyLockType                   `json:"lockType,omitempty"`
	LockedSubAssemblyOutputInfo *BTRigidOrLockedSubAssemblyOutputInfo3860 `json:"lockedSubAssemblyOutputInfo,omitempty"`
	// Type of JSON object.
	BtType *string `json:"btType,omitempty"`
}

// NewBTLockedAtCurrentPosition4006 instantiates a new BTLockedAtCurrentPosition4006 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTLockedAtCurrentPosition4006() *BTLockedAtCurrentPosition4006 {
	this := BTLockedAtCurrentPosition4006{}
	return &this
}

// NewBTLockedAtCurrentPosition4006WithDefaults instantiates a new BTLockedAtCurrentPosition4006 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTLockedAtCurrentPosition4006WithDefaults() *BTLockedAtCurrentPosition4006 {
	this := BTLockedAtCurrentPosition4006{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTLockedAtCurrentPosition4006) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedAtCurrentPosition4006) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTLockedAtCurrentPosition4006) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTLockedAtCurrentPosition4006) SetBtType(v string) {
	o.BtType = &v
}

// GetLockType returns the LockType field value if set, zero value otherwise.
func (o *BTLockedAtCurrentPosition4006) GetLockType() GBTSubAssemblyLockType {
	if o == nil || o.LockType == nil {
		var ret GBTSubAssemblyLockType
		return ret
	}
	return *o.LockType
}

// GetLockTypeOk returns a tuple with the LockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedAtCurrentPosition4006) GetLockTypeOk() (*GBTSubAssemblyLockType, bool) {
	if o == nil || o.LockType == nil {
		return nil, false
	}
	return o.LockType, true
}

// HasLockType returns a boolean if a field has been set.
func (o *BTLockedAtCurrentPosition4006) HasLockType() bool {
	if o != nil && o.LockType != nil {
		return true
	}

	return false
}

// SetLockType gets a reference to the given GBTSubAssemblyLockType and assigns it to the LockType field.
func (o *BTLockedAtCurrentPosition4006) SetLockType(v GBTSubAssemblyLockType) {
	o.LockType = &v
}

// GetLockedSubAssemblyOutputInfo returns the LockedSubAssemblyOutputInfo field value if set, zero value otherwise.
func (o *BTLockedAtCurrentPosition4006) GetLockedSubAssemblyOutputInfo() BTRigidOrLockedSubAssemblyOutputInfo3860 {
	if o == nil || o.LockedSubAssemblyOutputInfo == nil {
		var ret BTRigidOrLockedSubAssemblyOutputInfo3860
		return ret
	}
	return *o.LockedSubAssemblyOutputInfo
}

// GetLockedSubAssemblyOutputInfoOk returns a tuple with the LockedSubAssemblyOutputInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedAtCurrentPosition4006) GetLockedSubAssemblyOutputInfoOk() (*BTRigidOrLockedSubAssemblyOutputInfo3860, bool) {
	if o == nil || o.LockedSubAssemblyOutputInfo == nil {
		return nil, false
	}
	return o.LockedSubAssemblyOutputInfo, true
}

// HasLockedSubAssemblyOutputInfo returns a boolean if a field has been set.
func (o *BTLockedAtCurrentPosition4006) HasLockedSubAssemblyOutputInfo() bool {
	if o != nil && o.LockedSubAssemblyOutputInfo != nil {
		return true
	}

	return false
}

// SetLockedSubAssemblyOutputInfo gets a reference to the given BTRigidOrLockedSubAssemblyOutputInfo3860 and assigns it to the LockedSubAssemblyOutputInfo field.
func (o *BTLockedAtCurrentPosition4006) SetLockedSubAssemblyOutputInfo(v BTRigidOrLockedSubAssemblyOutputInfo3860) {
	o.LockedSubAssemblyOutputInfo = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTLockedAtCurrentPosition4006) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedAtCurrentPosition4006) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTLockedAtCurrentPosition4006) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTLockedAtCurrentPosition4006) SetBtType(v string) {
	o.BtType = &v
}

func (o BTLockedAtCurrentPosition4006) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTLockedSubAssembly4590, errBTLockedSubAssembly4590 := json.Marshal(o.BTLockedSubAssembly4590)
	if errBTLockedSubAssembly4590 != nil {
		return []byte{}, errBTLockedSubAssembly4590
	}
	errBTLockedSubAssembly4590 = json.Unmarshal([]byte(serializedBTLockedSubAssembly4590), &toSerialize)
	if errBTLockedSubAssembly4590 != nil {
		return []byte{}, errBTLockedSubAssembly4590
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.LockType != nil {
		toSerialize["lockType"] = o.LockType
	}
	if o.LockedSubAssemblyOutputInfo != nil {
		toSerialize["lockedSubAssemblyOutputInfo"] = o.LockedSubAssemblyOutputInfo
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTLockedAtCurrentPosition4006 struct {
	value *BTLockedAtCurrentPosition4006
	isSet bool
}

func (v NullableBTLockedAtCurrentPosition4006) Get() *BTLockedAtCurrentPosition4006 {
	return v.value
}

func (v *NullableBTLockedAtCurrentPosition4006) Set(val *BTLockedAtCurrentPosition4006) {
	v.value = val
	v.isSet = true
}

func (v NullableBTLockedAtCurrentPosition4006) IsSet() bool {
	return v.isSet
}

func (v *NullableBTLockedAtCurrentPosition4006) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTLockedAtCurrentPosition4006(val *BTLockedAtCurrentPosition4006) *NullableBTLockedAtCurrentPosition4006 {
	return &NullableBTLockedAtCurrentPosition4006{value: val, isSet: true}
}

func (v NullableBTLockedAtCurrentPosition4006) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTLockedAtCurrentPosition4006) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
