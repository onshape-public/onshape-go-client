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

// BTDiffJsonResponse2725Patch struct for BTDiffJsonResponse2725Patch
type BTDiffJsonResponse2725Patch struct {
	BtType *string `json:"btType,omitempty"`
}

// NewBTDiffJsonResponse2725Patch instantiates a new BTDiffJsonResponse2725Patch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDiffJsonResponse2725Patch() *BTDiffJsonResponse2725Patch {
	this := BTDiffJsonResponse2725Patch{}
	return &this
}

// NewBTDiffJsonResponse2725PatchWithDefaults instantiates a new BTDiffJsonResponse2725Patch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDiffJsonResponse2725PatchWithDefaults() *BTDiffJsonResponse2725Patch {
	this := BTDiffJsonResponse2725Patch{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTDiffJsonResponse2725Patch) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffJsonResponse2725Patch) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTDiffJsonResponse2725Patch) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTDiffJsonResponse2725Patch) SetBtType(v string) {
	o.BtType = &v
}

func (o BTDiffJsonResponse2725Patch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTDiffJsonResponse2725Patch struct {
	value *BTDiffJsonResponse2725Patch
	isSet bool
}

func (v NullableBTDiffJsonResponse2725Patch) Get() *BTDiffJsonResponse2725Patch {
	return v.value
}

func (v *NullableBTDiffJsonResponse2725Patch) Set(val *BTDiffJsonResponse2725Patch) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDiffJsonResponse2725Patch) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDiffJsonResponse2725Patch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDiffJsonResponse2725Patch(val *BTDiffJsonResponse2725Patch) *NullableBTDiffJsonResponse2725Patch {
	return &NullableBTDiffJsonResponse2725Patch{value: val, isSet: true}
}

func (v NullableBTDiffJsonResponse2725Patch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDiffJsonResponse2725Patch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
