/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23196-ae5f57bec467
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTJEditList2707 A list of edits that will be applied in order.
type BTJEditList2707 struct {
	BtType *string       `json:"btType,omitempty"`
	Edits  []BTJEdit3734 `json:"edits,omitempty"`
}

// NewBTJEditList2707 instantiates a new BTJEditList2707 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTJEditList2707() *BTJEditList2707 {
	this := BTJEditList2707{}
	return &this
}

// NewBTJEditList2707WithDefaults instantiates a new BTJEditList2707 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTJEditList2707WithDefaults() *BTJEditList2707 {
	this := BTJEditList2707{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTJEditList2707) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJEditList2707) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTJEditList2707) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTJEditList2707) SetBtType(v string) {
	o.BtType = &v
}

// GetEdits returns the Edits field value if set, zero value otherwise.
func (o *BTJEditList2707) GetEdits() []BTJEdit3734 {
	if o == nil || o.Edits == nil {
		var ret []BTJEdit3734
		return ret
	}
	return o.Edits
}

// GetEditsOk returns a tuple with the Edits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJEditList2707) GetEditsOk() ([]BTJEdit3734, bool) {
	if o == nil || o.Edits == nil {
		return nil, false
	}
	return o.Edits, true
}

// HasEdits returns a boolean if a field has been set.
func (o *BTJEditList2707) HasEdits() bool {
	if o != nil && o.Edits != nil {
		return true
	}

	return false
}

// SetEdits gets a reference to the given []BTJEdit3734 and assigns it to the Edits field.
func (o *BTJEditList2707) SetEdits(v []BTJEdit3734) {
	o.Edits = v
}

func (o BTJEditList2707) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Edits != nil {
		toSerialize["edits"] = o.Edits
	}
	return json.Marshal(toSerialize)
}

type NullableBTJEditList2707 struct {
	value *BTJEditList2707
	isSet bool
}

func (v NullableBTJEditList2707) Get() *BTJEditList2707 {
	return v.value
}

func (v *NullableBTJEditList2707) Set(val *BTJEditList2707) {
	v.value = val
	v.isSet = true
}

func (v NullableBTJEditList2707) IsSet() bool {
	return v.isSet
}

func (v *NullableBTJEditList2707) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTJEditList2707(val *BTJEditList2707) *NullableBTJEditList2707 {
	return &NullableBTJEditList2707{value: val, isSet: true}
}

func (v NullableBTJEditList2707) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTJEditList2707) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
