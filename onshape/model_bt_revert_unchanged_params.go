/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6946-bb9367a9d0cc
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTRevertUnchangedParams struct for BTRevertUnchangedParams
type BTRevertUnchangedParams struct {
	CompanyId    *string                          `json:"companyId,omitempty"`
	ConnectionId *string                          `json:"connectionId,omitempty"`
	DoUpdate     *bool                            `json:"doUpdate,omitempty"`
	Elements     []BTRevertUnchangedElementParams `json:"elements,omitempty"`
}

// NewBTRevertUnchangedParams instantiates a new BTRevertUnchangedParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRevertUnchangedParams() *BTRevertUnchangedParams {
	this := BTRevertUnchangedParams{}
	return &this
}

// NewBTRevertUnchangedParamsWithDefaults instantiates a new BTRevertUnchangedParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRevertUnchangedParamsWithDefaults() *BTRevertUnchangedParams {
	this := BTRevertUnchangedParams{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTRevertUnchangedParams) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevertUnchangedParams) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTRevertUnchangedParams) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTRevertUnchangedParams) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *BTRevertUnchangedParams) GetConnectionId() string {
	if o == nil || o.ConnectionId == nil {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevertUnchangedParams) GetConnectionIdOk() (*string, bool) {
	if o == nil || o.ConnectionId == nil {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *BTRevertUnchangedParams) HasConnectionId() bool {
	if o != nil && o.ConnectionId != nil {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *BTRevertUnchangedParams) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetDoUpdate returns the DoUpdate field value if set, zero value otherwise.
func (o *BTRevertUnchangedParams) GetDoUpdate() bool {
	if o == nil || o.DoUpdate == nil {
		var ret bool
		return ret
	}
	return *o.DoUpdate
}

// GetDoUpdateOk returns a tuple with the DoUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevertUnchangedParams) GetDoUpdateOk() (*bool, bool) {
	if o == nil || o.DoUpdate == nil {
		return nil, false
	}
	return o.DoUpdate, true
}

// HasDoUpdate returns a boolean if a field has been set.
func (o *BTRevertUnchangedParams) HasDoUpdate() bool {
	if o != nil && o.DoUpdate != nil {
		return true
	}

	return false
}

// SetDoUpdate gets a reference to the given bool and assigns it to the DoUpdate field.
func (o *BTRevertUnchangedParams) SetDoUpdate(v bool) {
	o.DoUpdate = &v
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *BTRevertUnchangedParams) GetElements() []BTRevertUnchangedElementParams {
	if o == nil || o.Elements == nil {
		var ret []BTRevertUnchangedElementParams
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevertUnchangedParams) GetElementsOk() ([]BTRevertUnchangedElementParams, bool) {
	if o == nil || o.Elements == nil {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *BTRevertUnchangedParams) HasElements() bool {
	if o != nil && o.Elements != nil {
		return true
	}

	return false
}

// SetElements gets a reference to the given []BTRevertUnchangedElementParams and assigns it to the Elements field.
func (o *BTRevertUnchangedParams) SetElements(v []BTRevertUnchangedElementParams) {
	o.Elements = v
}

func (o BTRevertUnchangedParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.ConnectionId != nil {
		toSerialize["connectionId"] = o.ConnectionId
	}
	if o.DoUpdate != nil {
		toSerialize["doUpdate"] = o.DoUpdate
	}
	if o.Elements != nil {
		toSerialize["elements"] = o.Elements
	}
	return json.Marshal(toSerialize)
}

type NullableBTRevertUnchangedParams struct {
	value *BTRevertUnchangedParams
	isSet bool
}

func (v NullableBTRevertUnchangedParams) Get() *BTRevertUnchangedParams {
	return v.value
}

func (v *NullableBTRevertUnchangedParams) Set(val *BTRevertUnchangedParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRevertUnchangedParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRevertUnchangedParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRevertUnchangedParams(val *BTRevertUnchangedParams) *NullableBTRevertUnchangedParams {
	return &NullableBTRevertUnchangedParams{value: val, isSet: true}
}

func (v NullableBTRevertUnchangedParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRevertUnchangedParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
