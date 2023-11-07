/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25478-d4e5ab4765a4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTTessellationProperties927 struct for BTTessellationProperties927
type BTTessellationProperties927 struct {
	AngularTolerance   *float64 `json:"angularTolerance,omitempty"`
	BtType             *string  `json:"btType,omitempty"`
	ChordalTolerance   *float64 `json:"chordalTolerance,omitempty"`
	NodeId             *string  `json:"nodeId,omitempty"`
	TessellationBudget *int32   `json:"tessellationBudget,omitempty"`
}

// NewBTTessellationProperties927 instantiates a new BTTessellationProperties927 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTessellationProperties927() *BTTessellationProperties927 {
	this := BTTessellationProperties927{}
	return &this
}

// NewBTTessellationProperties927WithDefaults instantiates a new BTTessellationProperties927 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTessellationProperties927WithDefaults() *BTTessellationProperties927 {
	this := BTTessellationProperties927{}
	return &this
}

// GetAngularTolerance returns the AngularTolerance field value if set, zero value otherwise.
func (o *BTTessellationProperties927) GetAngularTolerance() float64 {
	if o == nil || o.AngularTolerance == nil {
		var ret float64
		return ret
	}
	return *o.AngularTolerance
}

// GetAngularToleranceOk returns a tuple with the AngularTolerance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTessellationProperties927) GetAngularToleranceOk() (*float64, bool) {
	if o == nil || o.AngularTolerance == nil {
		return nil, false
	}
	return o.AngularTolerance, true
}

// HasAngularTolerance returns a boolean if a field has been set.
func (o *BTTessellationProperties927) HasAngularTolerance() bool {
	if o != nil && o.AngularTolerance != nil {
		return true
	}

	return false
}

// SetAngularTolerance gets a reference to the given float64 and assigns it to the AngularTolerance field.
func (o *BTTessellationProperties927) SetAngularTolerance(v float64) {
	o.AngularTolerance = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTessellationProperties927) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTessellationProperties927) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTessellationProperties927) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTessellationProperties927) SetBtType(v string) {
	o.BtType = &v
}

// GetChordalTolerance returns the ChordalTolerance field value if set, zero value otherwise.
func (o *BTTessellationProperties927) GetChordalTolerance() float64 {
	if o == nil || o.ChordalTolerance == nil {
		var ret float64
		return ret
	}
	return *o.ChordalTolerance
}

// GetChordalToleranceOk returns a tuple with the ChordalTolerance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTessellationProperties927) GetChordalToleranceOk() (*float64, bool) {
	if o == nil || o.ChordalTolerance == nil {
		return nil, false
	}
	return o.ChordalTolerance, true
}

// HasChordalTolerance returns a boolean if a field has been set.
func (o *BTTessellationProperties927) HasChordalTolerance() bool {
	if o != nil && o.ChordalTolerance != nil {
		return true
	}

	return false
}

// SetChordalTolerance gets a reference to the given float64 and assigns it to the ChordalTolerance field.
func (o *BTTessellationProperties927) SetChordalTolerance(v float64) {
	o.ChordalTolerance = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTTessellationProperties927) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTessellationProperties927) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTTessellationProperties927) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTTessellationProperties927) SetNodeId(v string) {
	o.NodeId = &v
}

// GetTessellationBudget returns the TessellationBudget field value if set, zero value otherwise.
func (o *BTTessellationProperties927) GetTessellationBudget() int32 {
	if o == nil || o.TessellationBudget == nil {
		var ret int32
		return ret
	}
	return *o.TessellationBudget
}

// GetTessellationBudgetOk returns a tuple with the TessellationBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTessellationProperties927) GetTessellationBudgetOk() (*int32, bool) {
	if o == nil || o.TessellationBudget == nil {
		return nil, false
	}
	return o.TessellationBudget, true
}

// HasTessellationBudget returns a boolean if a field has been set.
func (o *BTTessellationProperties927) HasTessellationBudget() bool {
	if o != nil && o.TessellationBudget != nil {
		return true
	}

	return false
}

// SetTessellationBudget gets a reference to the given int32 and assigns it to the TessellationBudget field.
func (o *BTTessellationProperties927) SetTessellationBudget(v int32) {
	o.TessellationBudget = &v
}

func (o BTTessellationProperties927) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AngularTolerance != nil {
		toSerialize["angularTolerance"] = o.AngularTolerance
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ChordalTolerance != nil {
		toSerialize["chordalTolerance"] = o.ChordalTolerance
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.TessellationBudget != nil {
		toSerialize["tessellationBudget"] = o.TessellationBudget
	}
	return json.Marshal(toSerialize)
}

type NullableBTTessellationProperties927 struct {
	value *BTTessellationProperties927
	isSet bool
}

func (v NullableBTTessellationProperties927) Get() *BTTessellationProperties927 {
	return v.value
}

func (v *NullableBTTessellationProperties927) Set(val *BTTessellationProperties927) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTessellationProperties927) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTessellationProperties927) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTessellationProperties927(val *BTTessellationProperties927) *NullableBTTessellationProperties927 {
	return &NullableBTTessellationProperties927{value: val, isSet: true}
}

func (v NullableBTTessellationProperties927) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTessellationProperties927) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
