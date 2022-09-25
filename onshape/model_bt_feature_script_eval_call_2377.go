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

// BTFeatureScriptEvalCall2377 struct for BTFeatureScriptEvalCall2377
type BTFeatureScriptEvalCall2377 struct {
	LibraryVersion         *int32               `json:"libraryVersion,omitempty"`
	MicroversionSkew       *bool                `json:"microversionSkew,omitempty"`
	Queries                *map[string][]string `json:"queries,omitempty"`
	RejectMicroversionSkew *bool                `json:"rejectMicroversionSkew,omitempty"`
	Script                 *string              `json:"script,omitempty"`
	SerializationVersion   *string              `json:"serializationVersion,omitempty"`
	SourceMicroversion     *string              `json:"sourceMicroversion,omitempty"`
}

// NewBTFeatureScriptEvalCall2377 instantiates a new BTFeatureScriptEvalCall2377 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureScriptEvalCall2377() *BTFeatureScriptEvalCall2377 {
	this := BTFeatureScriptEvalCall2377{}
	return &this
}

// NewBTFeatureScriptEvalCall2377WithDefaults instantiates a new BTFeatureScriptEvalCall2377 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureScriptEvalCall2377WithDefaults() *BTFeatureScriptEvalCall2377 {
	this := BTFeatureScriptEvalCall2377{}
	return &this
}

// GetLibraryVersion returns the LibraryVersion field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalCall2377) GetLibraryVersion() int32 {
	if o == nil || o.LibraryVersion == nil {
		var ret int32
		return ret
	}
	return *o.LibraryVersion
}

// GetLibraryVersionOk returns a tuple with the LibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalCall2377) GetLibraryVersionOk() (*int32, bool) {
	if o == nil || o.LibraryVersion == nil {
		return nil, false
	}
	return o.LibraryVersion, true
}

// HasLibraryVersion returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalCall2377) HasLibraryVersion() bool {
	if o != nil && o.LibraryVersion != nil {
		return true
	}

	return false
}

// SetLibraryVersion gets a reference to the given int32 and assigns it to the LibraryVersion field.
func (o *BTFeatureScriptEvalCall2377) SetLibraryVersion(v int32) {
	o.LibraryVersion = &v
}

// GetMicroversionSkew returns the MicroversionSkew field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalCall2377) GetMicroversionSkew() bool {
	if o == nil || o.MicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.MicroversionSkew
}

// GetMicroversionSkewOk returns a tuple with the MicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalCall2377) GetMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.MicroversionSkew == nil {
		return nil, false
	}
	return o.MicroversionSkew, true
}

// HasMicroversionSkew returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalCall2377) HasMicroversionSkew() bool {
	if o != nil && o.MicroversionSkew != nil {
		return true
	}

	return false
}

// SetMicroversionSkew gets a reference to the given bool and assigns it to the MicroversionSkew field.
func (o *BTFeatureScriptEvalCall2377) SetMicroversionSkew(v bool) {
	o.MicroversionSkew = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalCall2377) GetQueries() map[string][]string {
	if o == nil || o.Queries == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalCall2377) GetQueriesOk() (*map[string][]string, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalCall2377) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given map[string][]string and assigns it to the Queries field.
func (o *BTFeatureScriptEvalCall2377) SetQueries(v map[string][]string) {
	o.Queries = &v
}

// GetRejectMicroversionSkew returns the RejectMicroversionSkew field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalCall2377) GetRejectMicroversionSkew() bool {
	if o == nil || o.RejectMicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.RejectMicroversionSkew
}

// GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalCall2377) GetRejectMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.RejectMicroversionSkew == nil {
		return nil, false
	}
	return o.RejectMicroversionSkew, true
}

// HasRejectMicroversionSkew returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalCall2377) HasRejectMicroversionSkew() bool {
	if o != nil && o.RejectMicroversionSkew != nil {
		return true
	}

	return false
}

// SetRejectMicroversionSkew gets a reference to the given bool and assigns it to the RejectMicroversionSkew field.
func (o *BTFeatureScriptEvalCall2377) SetRejectMicroversionSkew(v bool) {
	o.RejectMicroversionSkew = &v
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalCall2377) GetScript() string {
	if o == nil || o.Script == nil {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalCall2377) GetScriptOk() (*string, bool) {
	if o == nil || o.Script == nil {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalCall2377) HasScript() bool {
	if o != nil && o.Script != nil {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *BTFeatureScriptEvalCall2377) SetScript(v string) {
	o.Script = &v
}

// GetSerializationVersion returns the SerializationVersion field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalCall2377) GetSerializationVersion() string {
	if o == nil || o.SerializationVersion == nil {
		var ret string
		return ret
	}
	return *o.SerializationVersion
}

// GetSerializationVersionOk returns a tuple with the SerializationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalCall2377) GetSerializationVersionOk() (*string, bool) {
	if o == nil || o.SerializationVersion == nil {
		return nil, false
	}
	return o.SerializationVersion, true
}

// HasSerializationVersion returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalCall2377) HasSerializationVersion() bool {
	if o != nil && o.SerializationVersion != nil {
		return true
	}

	return false
}

// SetSerializationVersion gets a reference to the given string and assigns it to the SerializationVersion field.
func (o *BTFeatureScriptEvalCall2377) SetSerializationVersion(v string) {
	o.SerializationVersion = &v
}

// GetSourceMicroversion returns the SourceMicroversion field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalCall2377) GetSourceMicroversion() string {
	if o == nil || o.SourceMicroversion == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversion
}

// GetSourceMicroversionOk returns a tuple with the SourceMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalCall2377) GetSourceMicroversionOk() (*string, bool) {
	if o == nil || o.SourceMicroversion == nil {
		return nil, false
	}
	return o.SourceMicroversion, true
}

// HasSourceMicroversion returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalCall2377) HasSourceMicroversion() bool {
	if o != nil && o.SourceMicroversion != nil {
		return true
	}

	return false
}

// SetSourceMicroversion gets a reference to the given string and assigns it to the SourceMicroversion field.
func (o *BTFeatureScriptEvalCall2377) SetSourceMicroversion(v string) {
	o.SourceMicroversion = &v
}

func (o BTFeatureScriptEvalCall2377) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LibraryVersion != nil {
		toSerialize["libraryVersion"] = o.LibraryVersion
	}
	if o.MicroversionSkew != nil {
		toSerialize["microversionSkew"] = o.MicroversionSkew
	}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.RejectMicroversionSkew != nil {
		toSerialize["rejectMicroversionSkew"] = o.RejectMicroversionSkew
	}
	if o.Script != nil {
		toSerialize["script"] = o.Script
	}
	if o.SerializationVersion != nil {
		toSerialize["serializationVersion"] = o.SerializationVersion
	}
	if o.SourceMicroversion != nil {
		toSerialize["sourceMicroversion"] = o.SourceMicroversion
	}
	return json.Marshal(toSerialize)
}

type NullableBTFeatureScriptEvalCall2377 struct {
	value *BTFeatureScriptEvalCall2377
	isSet bool
}

func (v NullableBTFeatureScriptEvalCall2377) Get() *BTFeatureScriptEvalCall2377 {
	return v.value
}

func (v *NullableBTFeatureScriptEvalCall2377) Set(val *BTFeatureScriptEvalCall2377) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureScriptEvalCall2377) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureScriptEvalCall2377) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureScriptEvalCall2377(val *BTFeatureScriptEvalCall2377) *NullableBTFeatureScriptEvalCall2377 {
	return &NullableBTFeatureScriptEvalCall2377{value: val, isSet: true}
}

func (v NullableBTFeatureScriptEvalCall2377) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureScriptEvalCall2377) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
