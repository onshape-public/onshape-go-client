/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.158.9878-ff51e1211d95
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMConfigurationParameterString872 struct for BTMConfigurationParameterString872
type BTMConfigurationParameterString872 struct {
	GeneratedParameterId *BTTreeNode20 `json:"generatedParameterId,omitempty"`
	ImportMicroversion   *string       `json:"importMicroversion,omitempty"`
	NodeId               *string       `json:"nodeId,omitempty"`
	ParameterId          *string       `json:"parameterId,omitempty"`
	ParameterName        *string       `json:"parameterName,omitempty"`
	ParameterType        *string       `json:"parameterType,omitempty"`
	Valid                *bool         `json:"valid,omitempty"`
	BtType               *string       `json:"btType,omitempty"`
	DefaultValue         *string       `json:"defaultValue,omitempty"`
}

// NewBTMConfigurationParameterString872 instantiates a new BTMConfigurationParameterString872 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMConfigurationParameterString872() *BTMConfigurationParameterString872 {
	this := BTMConfigurationParameterString872{}
	return &this
}

// NewBTMConfigurationParameterString872WithDefaults instantiates a new BTMConfigurationParameterString872 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMConfigurationParameterString872WithDefaults() *BTMConfigurationParameterString872 {
	this := BTMConfigurationParameterString872{}
	return &this
}

// GetGeneratedParameterId returns the GeneratedParameterId field value if set, zero value otherwise.
func (o *BTMConfigurationParameterString872) GetGeneratedParameterId() BTTreeNode20 {
	if o == nil || o.GeneratedParameterId == nil {
		var ret BTTreeNode20
		return ret
	}
	return *o.GeneratedParameterId
}

// GetGeneratedParameterIdOk returns a tuple with the GeneratedParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterString872) GetGeneratedParameterIdOk() (*BTTreeNode20, bool) {
	if o == nil || o.GeneratedParameterId == nil {
		return nil, false
	}
	return o.GeneratedParameterId, true
}

// HasGeneratedParameterId returns a boolean if a field has been set.
func (o *BTMConfigurationParameterString872) HasGeneratedParameterId() bool {
	if o != nil && o.GeneratedParameterId != nil {
		return true
	}

	return false
}

// SetGeneratedParameterId gets a reference to the given BTTreeNode20 and assigns it to the GeneratedParameterId field.
func (o *BTMConfigurationParameterString872) SetGeneratedParameterId(v BTTreeNode20) {
	o.GeneratedParameterId = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMConfigurationParameterString872) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterString872) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMConfigurationParameterString872) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMConfigurationParameterString872) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMConfigurationParameterString872) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterString872) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMConfigurationParameterString872) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMConfigurationParameterString872) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMConfigurationParameterString872) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterString872) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMConfigurationParameterString872) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMConfigurationParameterString872) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *BTMConfigurationParameterString872) GetParameterName() string {
	if o == nil || o.ParameterName == nil {
		var ret string
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterString872) GetParameterNameOk() (*string, bool) {
	if o == nil || o.ParameterName == nil {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *BTMConfigurationParameterString872) HasParameterName() bool {
	if o != nil && o.ParameterName != nil {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *BTMConfigurationParameterString872) SetParameterName(v string) {
	o.ParameterName = &v
}

// GetParameterType returns the ParameterType field value if set, zero value otherwise.
func (o *BTMConfigurationParameterString872) GetParameterType() string {
	if o == nil || o.ParameterType == nil {
		var ret string
		return ret
	}
	return *o.ParameterType
}

// GetParameterTypeOk returns a tuple with the ParameterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterString872) GetParameterTypeOk() (*string, bool) {
	if o == nil || o.ParameterType == nil {
		return nil, false
	}
	return o.ParameterType, true
}

// HasParameterType returns a boolean if a field has been set.
func (o *BTMConfigurationParameterString872) HasParameterType() bool {
	if o != nil && o.ParameterType != nil {
		return true
	}

	return false
}

// SetParameterType gets a reference to the given string and assigns it to the ParameterType field.
func (o *BTMConfigurationParameterString872) SetParameterType(v string) {
	o.ParameterType = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *BTMConfigurationParameterString872) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterString872) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *BTMConfigurationParameterString872) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *BTMConfigurationParameterString872) SetValid(v bool) {
	o.Valid = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMConfigurationParameterString872) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterString872) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMConfigurationParameterString872) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMConfigurationParameterString872) SetBtType(v string) {
	o.BtType = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *BTMConfigurationParameterString872) GetDefaultValue() string {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterString872) GetDefaultValueOk() (*string, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BTMConfigurationParameterString872) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *BTMConfigurationParameterString872) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

func (o BTMConfigurationParameterString872) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GeneratedParameterId != nil {
		toSerialize["generatedParameterId"] = o.GeneratedParameterId
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.ParameterName != nil {
		toSerialize["parameterName"] = o.ParameterName
	}
	if o.ParameterType != nil {
		toSerialize["parameterType"] = o.ParameterType
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	return json.Marshal(toSerialize)
}

type NullableBTMConfigurationParameterString872 struct {
	value *BTMConfigurationParameterString872
	isSet bool
}

func (v NullableBTMConfigurationParameterString872) Get() *BTMConfigurationParameterString872 {
	return v.value
}

func (v *NullableBTMConfigurationParameterString872) Set(val *BTMConfigurationParameterString872) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMConfigurationParameterString872) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMConfigurationParameterString872) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMConfigurationParameterString872(val *BTMConfigurationParameterString872) *NullableBTMConfigurationParameterString872 {
	return &NullableBTMConfigurationParameterString872{value: val, isSet: true}
}

func (v NullableBTMConfigurationParameterString872) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMConfigurationParameterString872) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
