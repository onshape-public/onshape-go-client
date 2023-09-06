/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.169.21759-9ddbba9fdfb8
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyPsFeatureInfo struct for BTAssemblyPsFeatureInfo
type BTAssemblyPsFeatureInfo struct {
	Configuration        *string `json:"configuration,omitempty"`
	DocumentId           *string `json:"documentId,omitempty"`
	DocumentMicroversion *string `json:"documentMicroversion,omitempty"`
	DocumentVersion      *string `json:"documentVersion,omitempty"`
	ElementId            *string `json:"elementId,omitempty"`
	FeatureId            *string `json:"featureId,omitempty"`
	FeatureType          *string `json:"featureType,omitempty"`
	FullConfiguration    *string `json:"fullConfiguration,omitempty"`
	PartNumber           *string `json:"partNumber,omitempty"`
	Revision             *string `json:"revision,omitempty"`
}

// NewBTAssemblyPsFeatureInfo instantiates a new BTAssemblyPsFeatureInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyPsFeatureInfo() *BTAssemblyPsFeatureInfo {
	this := BTAssemblyPsFeatureInfo{}
	return &this
}

// NewBTAssemblyPsFeatureInfoWithDefaults instantiates a new BTAssemblyPsFeatureInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyPsFeatureInfoWithDefaults() *BTAssemblyPsFeatureInfo {
	this := BTAssemblyPsFeatureInfo{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTAssemblyPsFeatureInfo) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPsFeatureInfo) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTAssemblyPsFeatureInfo) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTAssemblyPsFeatureInfo) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTAssemblyPsFeatureInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPsFeatureInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTAssemblyPsFeatureInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTAssemblyPsFeatureInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentMicroversion returns the DocumentMicroversion field value if set, zero value otherwise.
func (o *BTAssemblyPsFeatureInfo) GetDocumentMicroversion() string {
	if o == nil || o.DocumentMicroversion == nil {
		var ret string
		return ret
	}
	return *o.DocumentMicroversion
}

// GetDocumentMicroversionOk returns a tuple with the DocumentMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPsFeatureInfo) GetDocumentMicroversionOk() (*string, bool) {
	if o == nil || o.DocumentMicroversion == nil {
		return nil, false
	}
	return o.DocumentMicroversion, true
}

// HasDocumentMicroversion returns a boolean if a field has been set.
func (o *BTAssemblyPsFeatureInfo) HasDocumentMicroversion() bool {
	if o != nil && o.DocumentMicroversion != nil {
		return true
	}

	return false
}

// SetDocumentMicroversion gets a reference to the given string and assigns it to the DocumentMicroversion field.
func (o *BTAssemblyPsFeatureInfo) SetDocumentMicroversion(v string) {
	o.DocumentMicroversion = &v
}

// GetDocumentVersion returns the DocumentVersion field value if set, zero value otherwise.
func (o *BTAssemblyPsFeatureInfo) GetDocumentVersion() string {
	if o == nil || o.DocumentVersion == nil {
		var ret string
		return ret
	}
	return *o.DocumentVersion
}

// GetDocumentVersionOk returns a tuple with the DocumentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPsFeatureInfo) GetDocumentVersionOk() (*string, bool) {
	if o == nil || o.DocumentVersion == nil {
		return nil, false
	}
	return o.DocumentVersion, true
}

// HasDocumentVersion returns a boolean if a field has been set.
func (o *BTAssemblyPsFeatureInfo) HasDocumentVersion() bool {
	if o != nil && o.DocumentVersion != nil {
		return true
	}

	return false
}

// SetDocumentVersion gets a reference to the given string and assigns it to the DocumentVersion field.
func (o *BTAssemblyPsFeatureInfo) SetDocumentVersion(v string) {
	o.DocumentVersion = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTAssemblyPsFeatureInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPsFeatureInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTAssemblyPsFeatureInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTAssemblyPsFeatureInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTAssemblyPsFeatureInfo) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPsFeatureInfo) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTAssemblyPsFeatureInfo) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTAssemblyPsFeatureInfo) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *BTAssemblyPsFeatureInfo) GetFeatureType() string {
	if o == nil || o.FeatureType == nil {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPsFeatureInfo) GetFeatureTypeOk() (*string, bool) {
	if o == nil || o.FeatureType == nil {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *BTAssemblyPsFeatureInfo) HasFeatureType() bool {
	if o != nil && o.FeatureType != nil {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *BTAssemblyPsFeatureInfo) SetFeatureType(v string) {
	o.FeatureType = &v
}

// GetFullConfiguration returns the FullConfiguration field value if set, zero value otherwise.
func (o *BTAssemblyPsFeatureInfo) GetFullConfiguration() string {
	if o == nil || o.FullConfiguration == nil {
		var ret string
		return ret
	}
	return *o.FullConfiguration
}

// GetFullConfigurationOk returns a tuple with the FullConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPsFeatureInfo) GetFullConfigurationOk() (*string, bool) {
	if o == nil || o.FullConfiguration == nil {
		return nil, false
	}
	return o.FullConfiguration, true
}

// HasFullConfiguration returns a boolean if a field has been set.
func (o *BTAssemblyPsFeatureInfo) HasFullConfiguration() bool {
	if o != nil && o.FullConfiguration != nil {
		return true
	}

	return false
}

// SetFullConfiguration gets a reference to the given string and assigns it to the FullConfiguration field.
func (o *BTAssemblyPsFeatureInfo) SetFullConfiguration(v string) {
	o.FullConfiguration = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTAssemblyPsFeatureInfo) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPsFeatureInfo) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTAssemblyPsFeatureInfo) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTAssemblyPsFeatureInfo) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BTAssemblyPsFeatureInfo) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPsFeatureInfo) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BTAssemblyPsFeatureInfo) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BTAssemblyPsFeatureInfo) SetRevision(v string) {
	o.Revision = &v
}

func (o BTAssemblyPsFeatureInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.DocumentMicroversion != nil {
		toSerialize["documentMicroversion"] = o.DocumentMicroversion
	}
	if o.DocumentVersion != nil {
		toSerialize["documentVersion"] = o.DocumentVersion
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.FeatureType != nil {
		toSerialize["featureType"] = o.FeatureType
	}
	if o.FullConfiguration != nil {
		toSerialize["fullConfiguration"] = o.FullConfiguration
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyPsFeatureInfo struct {
	value *BTAssemblyPsFeatureInfo
	isSet bool
}

func (v NullableBTAssemblyPsFeatureInfo) Get() *BTAssemblyPsFeatureInfo {
	return v.value
}

func (v *NullableBTAssemblyPsFeatureInfo) Set(val *BTAssemblyPsFeatureInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyPsFeatureInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyPsFeatureInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyPsFeatureInfo(val *BTAssemblyPsFeatureInfo) *NullableBTAssemblyPsFeatureInfo {
	return &NullableBTAssemblyPsFeatureInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyPsFeatureInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyPsFeatureInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
