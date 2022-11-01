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

// BTIdTranslationParams struct for BTIdTranslationParams
type BTIdTranslationParams struct {
	Ids                        []string `json:"ids,omitempty"`
	LinkDocumentId             *string  `json:"linkDocumentId,omitempty"`
	SourceConfiguration        *string  `json:"sourceConfiguration,omitempty"`
	SourceDocumentMicroversion *string  `json:"sourceDocumentMicroversion,omitempty"`
	TargetConfiguration        *string  `json:"targetConfiguration,omitempty"`
}

// NewBTIdTranslationParams instantiates a new BTIdTranslationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTIdTranslationParams() *BTIdTranslationParams {
	this := BTIdTranslationParams{}
	return &this
}

// NewBTIdTranslationParamsWithDefaults instantiates a new BTIdTranslationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTIdTranslationParamsWithDefaults() *BTIdTranslationParams {
	this := BTIdTranslationParams{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *BTIdTranslationParams) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdTranslationParams) GetIdsOk() ([]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *BTIdTranslationParams) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *BTIdTranslationParams) SetIds(v []string) {
	o.Ids = v
}

// GetLinkDocumentId returns the LinkDocumentId field value if set, zero value otherwise.
func (o *BTIdTranslationParams) GetLinkDocumentId() string {
	if o == nil || o.LinkDocumentId == nil {
		var ret string
		return ret
	}
	return *o.LinkDocumentId
}

// GetLinkDocumentIdOk returns a tuple with the LinkDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdTranslationParams) GetLinkDocumentIdOk() (*string, bool) {
	if o == nil || o.LinkDocumentId == nil {
		return nil, false
	}
	return o.LinkDocumentId, true
}

// HasLinkDocumentId returns a boolean if a field has been set.
func (o *BTIdTranslationParams) HasLinkDocumentId() bool {
	if o != nil && o.LinkDocumentId != nil {
		return true
	}

	return false
}

// SetLinkDocumentId gets a reference to the given string and assigns it to the LinkDocumentId field.
func (o *BTIdTranslationParams) SetLinkDocumentId(v string) {
	o.LinkDocumentId = &v
}

// GetSourceConfiguration returns the SourceConfiguration field value if set, zero value otherwise.
func (o *BTIdTranslationParams) GetSourceConfiguration() string {
	if o == nil || o.SourceConfiguration == nil {
		var ret string
		return ret
	}
	return *o.SourceConfiguration
}

// GetSourceConfigurationOk returns a tuple with the SourceConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdTranslationParams) GetSourceConfigurationOk() (*string, bool) {
	if o == nil || o.SourceConfiguration == nil {
		return nil, false
	}
	return o.SourceConfiguration, true
}

// HasSourceConfiguration returns a boolean if a field has been set.
func (o *BTIdTranslationParams) HasSourceConfiguration() bool {
	if o != nil && o.SourceConfiguration != nil {
		return true
	}

	return false
}

// SetSourceConfiguration gets a reference to the given string and assigns it to the SourceConfiguration field.
func (o *BTIdTranslationParams) SetSourceConfiguration(v string) {
	o.SourceConfiguration = &v
}

// GetSourceDocumentMicroversion returns the SourceDocumentMicroversion field value if set, zero value otherwise.
func (o *BTIdTranslationParams) GetSourceDocumentMicroversion() string {
	if o == nil || o.SourceDocumentMicroversion == nil {
		var ret string
		return ret
	}
	return *o.SourceDocumentMicroversion
}

// GetSourceDocumentMicroversionOk returns a tuple with the SourceDocumentMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdTranslationParams) GetSourceDocumentMicroversionOk() (*string, bool) {
	if o == nil || o.SourceDocumentMicroversion == nil {
		return nil, false
	}
	return o.SourceDocumentMicroversion, true
}

// HasSourceDocumentMicroversion returns a boolean if a field has been set.
func (o *BTIdTranslationParams) HasSourceDocumentMicroversion() bool {
	if o != nil && o.SourceDocumentMicroversion != nil {
		return true
	}

	return false
}

// SetSourceDocumentMicroversion gets a reference to the given string and assigns it to the SourceDocumentMicroversion field.
func (o *BTIdTranslationParams) SetSourceDocumentMicroversion(v string) {
	o.SourceDocumentMicroversion = &v
}

// GetTargetConfiguration returns the TargetConfiguration field value if set, zero value otherwise.
func (o *BTIdTranslationParams) GetTargetConfiguration() string {
	if o == nil || o.TargetConfiguration == nil {
		var ret string
		return ret
	}
	return *o.TargetConfiguration
}

// GetTargetConfigurationOk returns a tuple with the TargetConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdTranslationParams) GetTargetConfigurationOk() (*string, bool) {
	if o == nil || o.TargetConfiguration == nil {
		return nil, false
	}
	return o.TargetConfiguration, true
}

// HasTargetConfiguration returns a boolean if a field has been set.
func (o *BTIdTranslationParams) HasTargetConfiguration() bool {
	if o != nil && o.TargetConfiguration != nil {
		return true
	}

	return false
}

// SetTargetConfiguration gets a reference to the given string and assigns it to the TargetConfiguration field.
func (o *BTIdTranslationParams) SetTargetConfiguration(v string) {
	o.TargetConfiguration = &v
}

func (o BTIdTranslationParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.LinkDocumentId != nil {
		toSerialize["linkDocumentId"] = o.LinkDocumentId
	}
	if o.SourceConfiguration != nil {
		toSerialize["sourceConfiguration"] = o.SourceConfiguration
	}
	if o.SourceDocumentMicroversion != nil {
		toSerialize["sourceDocumentMicroversion"] = o.SourceDocumentMicroversion
	}
	if o.TargetConfiguration != nil {
		toSerialize["targetConfiguration"] = o.TargetConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableBTIdTranslationParams struct {
	value *BTIdTranslationParams
	isSet bool
}

func (v NullableBTIdTranslationParams) Get() *BTIdTranslationParams {
	return v.value
}

func (v *NullableBTIdTranslationParams) Set(val *BTIdTranslationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTIdTranslationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTIdTranslationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTIdTranslationParams(val *BTIdTranslationParams) *NullableBTIdTranslationParams {
	return &NullableBTIdTranslationParams{value: val, isSet: true}
}

func (v NullableBTIdTranslationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTIdTranslationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
