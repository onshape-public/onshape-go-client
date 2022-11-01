/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7120-bb0c8e12993e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// ConfigurationEntry struct for ConfigurationEntry
type ConfigurationEntry struct {
	ParameterId    *string `json:"parameterId,omitempty"`
	ParameterValue *string `json:"parameterValue,omitempty"`
}

// NewConfigurationEntry instantiates a new ConfigurationEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationEntry() *ConfigurationEntry {
	this := ConfigurationEntry{}
	return &this
}

// NewConfigurationEntryWithDefaults instantiates a new ConfigurationEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationEntryWithDefaults() *ConfigurationEntry {
	this := ConfigurationEntry{}
	return &this
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *ConfigurationEntry) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationEntry) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *ConfigurationEntry) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *ConfigurationEntry) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetParameterValue returns the ParameterValue field value if set, zero value otherwise.
func (o *ConfigurationEntry) GetParameterValue() string {
	if o == nil || o.ParameterValue == nil {
		var ret string
		return ret
	}
	return *o.ParameterValue
}

// GetParameterValueOk returns a tuple with the ParameterValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationEntry) GetParameterValueOk() (*string, bool) {
	if o == nil || o.ParameterValue == nil {
		return nil, false
	}
	return o.ParameterValue, true
}

// HasParameterValue returns a boolean if a field has been set.
func (o *ConfigurationEntry) HasParameterValue() bool {
	if o != nil && o.ParameterValue != nil {
		return true
	}

	return false
}

// SetParameterValue gets a reference to the given string and assigns it to the ParameterValue field.
func (o *ConfigurationEntry) SetParameterValue(v string) {
	o.ParameterValue = &v
}

func (o ConfigurationEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.ParameterValue != nil {
		toSerialize["parameterValue"] = o.ParameterValue
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurationEntry struct {
	value *ConfigurationEntry
	isSet bool
}

func (v NullableConfigurationEntry) Get() *ConfigurationEntry {
	return v.value
}

func (v *NullableConfigurationEntry) Set(val *ConfigurationEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationEntry(val *ConfigurationEntry) *NullableConfigurationEntry {
	return &NullableConfigurationEntry{value: val, isSet: true}
}

func (v NullableConfigurationEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
