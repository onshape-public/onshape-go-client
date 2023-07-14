/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.166.18819-fa27aca4c021
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDocumentWithVersionId struct for BTDocumentWithVersionId
type BTDocumentWithVersionId struct {
	DocumentId        *string `json:"documentId,omitempty"`
	DocumentVersionId *string `json:"documentVersionId,omitempty"`
}

// NewBTDocumentWithVersionId instantiates a new BTDocumentWithVersionId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentWithVersionId() *BTDocumentWithVersionId {
	this := BTDocumentWithVersionId{}
	return &this
}

// NewBTDocumentWithVersionIdWithDefaults instantiates a new BTDocumentWithVersionId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentWithVersionIdWithDefaults() *BTDocumentWithVersionId {
	this := BTDocumentWithVersionId{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTDocumentWithVersionId) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentWithVersionId) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTDocumentWithVersionId) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTDocumentWithVersionId) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentVersionId returns the DocumentVersionId field value if set, zero value otherwise.
func (o *BTDocumentWithVersionId) GetDocumentVersionId() string {
	if o == nil || o.DocumentVersionId == nil {
		var ret string
		return ret
	}
	return *o.DocumentVersionId
}

// GetDocumentVersionIdOk returns a tuple with the DocumentVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentWithVersionId) GetDocumentVersionIdOk() (*string, bool) {
	if o == nil || o.DocumentVersionId == nil {
		return nil, false
	}
	return o.DocumentVersionId, true
}

// HasDocumentVersionId returns a boolean if a field has been set.
func (o *BTDocumentWithVersionId) HasDocumentVersionId() bool {
	if o != nil && o.DocumentVersionId != nil {
		return true
	}

	return false
}

// SetDocumentVersionId gets a reference to the given string and assigns it to the DocumentVersionId field.
func (o *BTDocumentWithVersionId) SetDocumentVersionId(v string) {
	o.DocumentVersionId = &v
}

func (o BTDocumentWithVersionId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.DocumentVersionId != nil {
		toSerialize["documentVersionId"] = o.DocumentVersionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentWithVersionId struct {
	value *BTDocumentWithVersionId
	isSet bool
}

func (v NullableBTDocumentWithVersionId) Get() *BTDocumentWithVersionId {
	return v.value
}

func (v *NullableBTDocumentWithVersionId) Set(val *BTDocumentWithVersionId) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentWithVersionId) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentWithVersionId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentWithVersionId(val *BTDocumentWithVersionId) *NullableBTDocumentWithVersionId {
	return &NullableBTDocumentWithVersionId{value: val, isSet: true}
}

func (v NullableBTDocumentWithVersionId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentWithVersionId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
