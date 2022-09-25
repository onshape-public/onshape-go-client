/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6617-832d47e03347
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPSpace10 struct for BTPSpace10
type BTPSpace10 struct {
	BtType *string  `json:"btType,omitempty"`
	Lines  []string `json:"lines,omitempty"`
	NodeId *string  `json:"nodeId,omitempty"`
	Text   *string  `json:"text,omitempty"`
}

// NewBTPSpace10 instantiates a new BTPSpace10 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPSpace10() *BTPSpace10 {
	this := BTPSpace10{}
	return &this
}

// NewBTPSpace10WithDefaults instantiates a new BTPSpace10 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPSpace10WithDefaults() *BTPSpace10 {
	this := BTPSpace10{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPSpace10) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPSpace10) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPSpace10) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPSpace10) SetBtType(v string) {
	o.BtType = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *BTPSpace10) GetLines() []string {
	if o == nil || o.Lines == nil {
		var ret []string
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPSpace10) GetLinesOk() ([]string, bool) {
	if o == nil || o.Lines == nil {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *BTPSpace10) HasLines() bool {
	if o != nil && o.Lines != nil {
		return true
	}

	return false
}

// SetLines gets a reference to the given []string and assigns it to the Lines field.
func (o *BTPSpace10) SetLines(v []string) {
	o.Lines = v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPSpace10) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPSpace10) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPSpace10) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPSpace10) SetNodeId(v string) {
	o.NodeId = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *BTPSpace10) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPSpace10) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *BTPSpace10) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *BTPSpace10) SetText(v string) {
	o.Text = &v
}

func (o BTPSpace10) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Lines != nil {
		toSerialize["lines"] = o.Lines
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableBTPSpace10 struct {
	value *BTPSpace10
	isSet bool
}

func (v NullableBTPSpace10) Get() *BTPSpace10 {
	return v.value
}

func (v *NullableBTPSpace10) Set(val *BTPSpace10) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPSpace10) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPSpace10) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPSpace10(val *BTPSpace10) *NullableBTPSpace10 {
	return &NullableBTPSpace10{value: val, isSet: true}
}

func (v NullableBTPSpace10) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPSpace10) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
