/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.150.5633-5ed6b38daa6b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTStringNodeWrapper4224 struct for BTStringNodeWrapper4224
type BTStringNodeWrapper4224 struct {
	BtType *string `json:"btType,omitempty"`
	NodeId *string `json:"nodeId,omitempty"`
	String *string `json:"string,omitempty"`
}

// NewBTStringNodeWrapper4224 instantiates a new BTStringNodeWrapper4224 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTStringNodeWrapper4224() *BTStringNodeWrapper4224 {
	this := BTStringNodeWrapper4224{}
	return &this
}

// NewBTStringNodeWrapper4224WithDefaults instantiates a new BTStringNodeWrapper4224 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTStringNodeWrapper4224WithDefaults() *BTStringNodeWrapper4224 {
	this := BTStringNodeWrapper4224{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTStringNodeWrapper4224) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringNodeWrapper4224) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTStringNodeWrapper4224) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTStringNodeWrapper4224) SetBtType(v string) {
	o.BtType = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTStringNodeWrapper4224) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringNodeWrapper4224) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTStringNodeWrapper4224) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTStringNodeWrapper4224) SetNodeId(v string) {
	o.NodeId = &v
}

// GetString returns the String field value if set, zero value otherwise.
func (o *BTStringNodeWrapper4224) GetString() string {
	if o == nil || o.String == nil {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringNodeWrapper4224) GetStringOk() (*string, bool) {
	if o == nil || o.String == nil {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *BTStringNodeWrapper4224) HasString() bool {
	if o != nil && o.String != nil {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *BTStringNodeWrapper4224) SetString(v string) {
	o.String = &v
}

func (o BTStringNodeWrapper4224) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.String != nil {
		toSerialize["string"] = o.String
	}
	return json.Marshal(toSerialize)
}

type NullableBTStringNodeWrapper4224 struct {
	value *BTStringNodeWrapper4224
	isSet bool
}

func (v NullableBTStringNodeWrapper4224) Get() *BTStringNodeWrapper4224 {
	return v.value
}

func (v *NullableBTStringNodeWrapper4224) Set(val *BTStringNodeWrapper4224) {
	v.value = val
	v.isSet = true
}

func (v NullableBTStringNodeWrapper4224) IsSet() bool {
	return v.isSet
}

func (v *NullableBTStringNodeWrapper4224) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTStringNodeWrapper4224(val *BTStringNodeWrapper4224) *NullableBTStringNodeWrapper4224 {
	return &NullableBTStringNodeWrapper4224{value: val, isSet: true}
}

func (v NullableBTStringNodeWrapper4224) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTStringNodeWrapper4224) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
