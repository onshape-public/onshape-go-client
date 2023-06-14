/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.17369-82f2ed5d514e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMUserCode161 struct for BTMUserCode161
type BTMUserCode161 struct {
	BtType             *string          `json:"btType,omitempty"`
	ImportMicroversion *string          `json:"importMicroversion,omitempty"`
	NodeId             *string          `json:"nodeId,omitempty"`
	Parsed             *BTPStatement269 `json:"parsed,omitempty"`
	Statement          *string          `json:"statement,omitempty"`
}

// NewBTMUserCode161 instantiates a new BTMUserCode161 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMUserCode161() *BTMUserCode161 {
	this := BTMUserCode161{}
	return &this
}

// NewBTMUserCode161WithDefaults instantiates a new BTMUserCode161 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMUserCode161WithDefaults() *BTMUserCode161 {
	this := BTMUserCode161{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMUserCode161) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMUserCode161) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMUserCode161) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMUserCode161) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMUserCode161) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMUserCode161) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMUserCode161) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMUserCode161) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMUserCode161) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMUserCode161) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMUserCode161) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMUserCode161) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParsed returns the Parsed field value if set, zero value otherwise.
func (o *BTMUserCode161) GetParsed() BTPStatement269 {
	if o == nil || o.Parsed == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.Parsed
}

// GetParsedOk returns a tuple with the Parsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMUserCode161) GetParsedOk() (*BTPStatement269, bool) {
	if o == nil || o.Parsed == nil {
		return nil, false
	}
	return o.Parsed, true
}

// HasParsed returns a boolean if a field has been set.
func (o *BTMUserCode161) HasParsed() bool {
	if o != nil && o.Parsed != nil {
		return true
	}

	return false
}

// SetParsed gets a reference to the given BTPStatement269 and assigns it to the Parsed field.
func (o *BTMUserCode161) SetParsed(v BTPStatement269) {
	o.Parsed = &v
}

// GetStatement returns the Statement field value if set, zero value otherwise.
func (o *BTMUserCode161) GetStatement() string {
	if o == nil || o.Statement == nil {
		var ret string
		return ret
	}
	return *o.Statement
}

// GetStatementOk returns a tuple with the Statement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMUserCode161) GetStatementOk() (*string, bool) {
	if o == nil || o.Statement == nil {
		return nil, false
	}
	return o.Statement, true
}

// HasStatement returns a boolean if a field has been set.
func (o *BTMUserCode161) HasStatement() bool {
	if o != nil && o.Statement != nil {
		return true
	}

	return false
}

// SetStatement gets a reference to the given string and assigns it to the Statement field.
func (o *BTMUserCode161) SetStatement(v string) {
	o.Statement = &v
}

func (o BTMUserCode161) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Parsed != nil {
		toSerialize["parsed"] = o.Parsed
	}
	if o.Statement != nil {
		toSerialize["statement"] = o.Statement
	}
	return json.Marshal(toSerialize)
}

type NullableBTMUserCode161 struct {
	value *BTMUserCode161
	isSet bool
}

func (v NullableBTMUserCode161) Get() *BTMUserCode161 {
	return v.value
}

func (v *NullableBTMUserCode161) Set(val *BTMUserCode161) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMUserCode161) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMUserCode161) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMUserCode161(val *BTMUserCode161) *NullableBTMUserCode161 {
	return &NullableBTMUserCode161{value: val, isSet: true}
}

func (v NullableBTMUserCode161) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMUserCode161) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
