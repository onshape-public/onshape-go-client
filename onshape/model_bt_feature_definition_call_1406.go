/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.160.12117-533984354e9a
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTFeatureDefinitionCall1406 struct for BTFeatureDefinitionCall1406
type BTFeatureDefinitionCall1406 struct {
	Feature                *BTMFeature134 `json:"feature,omitempty"`
	LibraryVersion         *int32         `json:"libraryVersion,omitempty"`
	MicroversionSkew       *bool          `json:"microversionSkew,omitempty"`
	RejectMicroversionSkew *bool          `json:"rejectMicroversionSkew,omitempty"`
	SerializationVersion   *string        `json:"serializationVersion,omitempty"`
	SourceMicroversion     *string        `json:"sourceMicroversion,omitempty"`
}

// NewBTFeatureDefinitionCall1406 instantiates a new BTFeatureDefinitionCall1406 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureDefinitionCall1406() *BTFeatureDefinitionCall1406 {
	this := BTFeatureDefinitionCall1406{}
	return &this
}

// NewBTFeatureDefinitionCall1406WithDefaults instantiates a new BTFeatureDefinitionCall1406 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureDefinitionCall1406WithDefaults() *BTFeatureDefinitionCall1406 {
	this := BTFeatureDefinitionCall1406{}
	return &this
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *BTFeatureDefinitionCall1406) GetFeature() BTMFeature134 {
	if o == nil || o.Feature == nil {
		var ret BTMFeature134
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureDefinitionCall1406) GetFeatureOk() (*BTMFeature134, bool) {
	if o == nil || o.Feature == nil {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *BTFeatureDefinitionCall1406) HasFeature() bool {
	if o != nil && o.Feature != nil {
		return true
	}

	return false
}

// SetFeature gets a reference to the given BTMFeature134 and assigns it to the Feature field.
func (o *BTFeatureDefinitionCall1406) SetFeature(v BTMFeature134) {
	o.Feature = &v
}

// GetLibraryVersion returns the LibraryVersion field value if set, zero value otherwise.
func (o *BTFeatureDefinitionCall1406) GetLibraryVersion() int32 {
	if o == nil || o.LibraryVersion == nil {
		var ret int32
		return ret
	}
	return *o.LibraryVersion
}

// GetLibraryVersionOk returns a tuple with the LibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureDefinitionCall1406) GetLibraryVersionOk() (*int32, bool) {
	if o == nil || o.LibraryVersion == nil {
		return nil, false
	}
	return o.LibraryVersion, true
}

// HasLibraryVersion returns a boolean if a field has been set.
func (o *BTFeatureDefinitionCall1406) HasLibraryVersion() bool {
	if o != nil && o.LibraryVersion != nil {
		return true
	}

	return false
}

// SetLibraryVersion gets a reference to the given int32 and assigns it to the LibraryVersion field.
func (o *BTFeatureDefinitionCall1406) SetLibraryVersion(v int32) {
	o.LibraryVersion = &v
}

// GetMicroversionSkew returns the MicroversionSkew field value if set, zero value otherwise.
func (o *BTFeatureDefinitionCall1406) GetMicroversionSkew() bool {
	if o == nil || o.MicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.MicroversionSkew
}

// GetMicroversionSkewOk returns a tuple with the MicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureDefinitionCall1406) GetMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.MicroversionSkew == nil {
		return nil, false
	}
	return o.MicroversionSkew, true
}

// HasMicroversionSkew returns a boolean if a field has been set.
func (o *BTFeatureDefinitionCall1406) HasMicroversionSkew() bool {
	if o != nil && o.MicroversionSkew != nil {
		return true
	}

	return false
}

// SetMicroversionSkew gets a reference to the given bool and assigns it to the MicroversionSkew field.
func (o *BTFeatureDefinitionCall1406) SetMicroversionSkew(v bool) {
	o.MicroversionSkew = &v
}

// GetRejectMicroversionSkew returns the RejectMicroversionSkew field value if set, zero value otherwise.
func (o *BTFeatureDefinitionCall1406) GetRejectMicroversionSkew() bool {
	if o == nil || o.RejectMicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.RejectMicroversionSkew
}

// GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureDefinitionCall1406) GetRejectMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.RejectMicroversionSkew == nil {
		return nil, false
	}
	return o.RejectMicroversionSkew, true
}

// HasRejectMicroversionSkew returns a boolean if a field has been set.
func (o *BTFeatureDefinitionCall1406) HasRejectMicroversionSkew() bool {
	if o != nil && o.RejectMicroversionSkew != nil {
		return true
	}

	return false
}

// SetRejectMicroversionSkew gets a reference to the given bool and assigns it to the RejectMicroversionSkew field.
func (o *BTFeatureDefinitionCall1406) SetRejectMicroversionSkew(v bool) {
	o.RejectMicroversionSkew = &v
}

// GetSerializationVersion returns the SerializationVersion field value if set, zero value otherwise.
func (o *BTFeatureDefinitionCall1406) GetSerializationVersion() string {
	if o == nil || o.SerializationVersion == nil {
		var ret string
		return ret
	}
	return *o.SerializationVersion
}

// GetSerializationVersionOk returns a tuple with the SerializationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureDefinitionCall1406) GetSerializationVersionOk() (*string, bool) {
	if o == nil || o.SerializationVersion == nil {
		return nil, false
	}
	return o.SerializationVersion, true
}

// HasSerializationVersion returns a boolean if a field has been set.
func (o *BTFeatureDefinitionCall1406) HasSerializationVersion() bool {
	if o != nil && o.SerializationVersion != nil {
		return true
	}

	return false
}

// SetSerializationVersion gets a reference to the given string and assigns it to the SerializationVersion field.
func (o *BTFeatureDefinitionCall1406) SetSerializationVersion(v string) {
	o.SerializationVersion = &v
}

// GetSourceMicroversion returns the SourceMicroversion field value if set, zero value otherwise.
func (o *BTFeatureDefinitionCall1406) GetSourceMicroversion() string {
	if o == nil || o.SourceMicroversion == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversion
}

// GetSourceMicroversionOk returns a tuple with the SourceMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureDefinitionCall1406) GetSourceMicroversionOk() (*string, bool) {
	if o == nil || o.SourceMicroversion == nil {
		return nil, false
	}
	return o.SourceMicroversion, true
}

// HasSourceMicroversion returns a boolean if a field has been set.
func (o *BTFeatureDefinitionCall1406) HasSourceMicroversion() bool {
	if o != nil && o.SourceMicroversion != nil {
		return true
	}

	return false
}

// SetSourceMicroversion gets a reference to the given string and assigns it to the SourceMicroversion field.
func (o *BTFeatureDefinitionCall1406) SetSourceMicroversion(v string) {
	o.SourceMicroversion = &v
}

func (o BTFeatureDefinitionCall1406) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Feature != nil {
		toSerialize["feature"] = o.Feature
	}
	if o.LibraryVersion != nil {
		toSerialize["libraryVersion"] = o.LibraryVersion
	}
	if o.MicroversionSkew != nil {
		toSerialize["microversionSkew"] = o.MicroversionSkew
	}
	if o.RejectMicroversionSkew != nil {
		toSerialize["rejectMicroversionSkew"] = o.RejectMicroversionSkew
	}
	if o.SerializationVersion != nil {
		toSerialize["serializationVersion"] = o.SerializationVersion
	}
	if o.SourceMicroversion != nil {
		toSerialize["sourceMicroversion"] = o.SourceMicroversion
	}
	return json.Marshal(toSerialize)
}

type NullableBTFeatureDefinitionCall1406 struct {
	value *BTFeatureDefinitionCall1406
	isSet bool
}

func (v NullableBTFeatureDefinitionCall1406) Get() *BTFeatureDefinitionCall1406 {
	return v.value
}

func (v *NullableBTFeatureDefinitionCall1406) Set(val *BTFeatureDefinitionCall1406) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureDefinitionCall1406) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureDefinitionCall1406) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureDefinitionCall1406(val *BTFeatureDefinitionCall1406) *NullableBTFeatureDefinitionCall1406 {
	return &NullableBTFeatureDefinitionCall1406{value: val, isSet: true}
}

func (v NullableBTFeatureDefinitionCall1406) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureDefinitionCall1406) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
