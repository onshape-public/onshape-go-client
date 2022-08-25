/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6210-7a182f03d281
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTExternalReference1936 struct for BTExternalReference1936
type BTExternalReference1936 struct {
	BtType                                  *string                               `json:"btType,omitempty"`
	Configured                              *bool                                 `json:"configured,omitempty"`
	ElementId                               *string                               `json:"elementId,omitempty"`
	ExternalDocumentWithVersion             *BTDocumentWithVersionId              `json:"externalDocumentWithVersion,omitempty"`
	ExternalDocumentWithVersionAndElementId *BTDocumentWithVersionAndElementId    `json:"externalDocumentWithVersionAndElementId,omitempty"`
	ExternalReference                       *bool                                 `json:"externalReference,omitempty"`
	FullElementId                           *BTFullElementId756                   `json:"fullElementId,omitempty"`
	MicroversionIdAndConfiguration          *BTMicroversionIdAndConfiguration2338 `json:"microversionIdAndConfiguration,omitempty"`
	NodeId                                  *string                               `json:"nodeId,omitempty"`
	DocumentVersionId                       *string                               `json:"documentVersionId,omitempty"`
}

// NewBTExternalReference1936 instantiates a new BTExternalReference1936 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExternalReference1936() *BTExternalReference1936 {
	this := BTExternalReference1936{}
	return &this
}

// NewBTExternalReference1936WithDefaults instantiates a new BTExternalReference1936 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExternalReference1936WithDefaults() *BTExternalReference1936 {
	this := BTExternalReference1936{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExternalReference1936) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalReference1936) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExternalReference1936) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExternalReference1936) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *BTExternalReference1936) GetConfigured() bool {
	if o == nil || o.Configured == nil {
		var ret bool
		return ret
	}
	return *o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalReference1936) GetConfiguredOk() (*bool, bool) {
	if o == nil || o.Configured == nil {
		return nil, false
	}
	return o.Configured, true
}

// HasConfigured returns a boolean if a field has been set.
func (o *BTExternalReference1936) HasConfigured() bool {
	if o != nil && o.Configured != nil {
		return true
	}

	return false
}

// SetConfigured gets a reference to the given bool and assigns it to the Configured field.
func (o *BTExternalReference1936) SetConfigured(v bool) {
	o.Configured = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTExternalReference1936) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalReference1936) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTExternalReference1936) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTExternalReference1936) SetElementId(v string) {
	o.ElementId = &v
}

// GetExternalDocumentWithVersion returns the ExternalDocumentWithVersion field value if set, zero value otherwise.
func (o *BTExternalReference1936) GetExternalDocumentWithVersion() BTDocumentWithVersionId {
	if o == nil || o.ExternalDocumentWithVersion == nil {
		var ret BTDocumentWithVersionId
		return ret
	}
	return *o.ExternalDocumentWithVersion
}

// GetExternalDocumentWithVersionOk returns a tuple with the ExternalDocumentWithVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalReference1936) GetExternalDocumentWithVersionOk() (*BTDocumentWithVersionId, bool) {
	if o == nil || o.ExternalDocumentWithVersion == nil {
		return nil, false
	}
	return o.ExternalDocumentWithVersion, true
}

// HasExternalDocumentWithVersion returns a boolean if a field has been set.
func (o *BTExternalReference1936) HasExternalDocumentWithVersion() bool {
	if o != nil && o.ExternalDocumentWithVersion != nil {
		return true
	}

	return false
}

// SetExternalDocumentWithVersion gets a reference to the given BTDocumentWithVersionId and assigns it to the ExternalDocumentWithVersion field.
func (o *BTExternalReference1936) SetExternalDocumentWithVersion(v BTDocumentWithVersionId) {
	o.ExternalDocumentWithVersion = &v
}

// GetExternalDocumentWithVersionAndElementId returns the ExternalDocumentWithVersionAndElementId field value if set, zero value otherwise.
func (o *BTExternalReference1936) GetExternalDocumentWithVersionAndElementId() BTDocumentWithVersionAndElementId {
	if o == nil || o.ExternalDocumentWithVersionAndElementId == nil {
		var ret BTDocumentWithVersionAndElementId
		return ret
	}
	return *o.ExternalDocumentWithVersionAndElementId
}

// GetExternalDocumentWithVersionAndElementIdOk returns a tuple with the ExternalDocumentWithVersionAndElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalReference1936) GetExternalDocumentWithVersionAndElementIdOk() (*BTDocumentWithVersionAndElementId, bool) {
	if o == nil || o.ExternalDocumentWithVersionAndElementId == nil {
		return nil, false
	}
	return o.ExternalDocumentWithVersionAndElementId, true
}

// HasExternalDocumentWithVersionAndElementId returns a boolean if a field has been set.
func (o *BTExternalReference1936) HasExternalDocumentWithVersionAndElementId() bool {
	if o != nil && o.ExternalDocumentWithVersionAndElementId != nil {
		return true
	}

	return false
}

// SetExternalDocumentWithVersionAndElementId gets a reference to the given BTDocumentWithVersionAndElementId and assigns it to the ExternalDocumentWithVersionAndElementId field.
func (o *BTExternalReference1936) SetExternalDocumentWithVersionAndElementId(v BTDocumentWithVersionAndElementId) {
	o.ExternalDocumentWithVersionAndElementId = &v
}

// GetExternalReference returns the ExternalReference field value if set, zero value otherwise.
func (o *BTExternalReference1936) GetExternalReference() bool {
	if o == nil || o.ExternalReference == nil {
		var ret bool
		return ret
	}
	return *o.ExternalReference
}

// GetExternalReferenceOk returns a tuple with the ExternalReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalReference1936) GetExternalReferenceOk() (*bool, bool) {
	if o == nil || o.ExternalReference == nil {
		return nil, false
	}
	return o.ExternalReference, true
}

// HasExternalReference returns a boolean if a field has been set.
func (o *BTExternalReference1936) HasExternalReference() bool {
	if o != nil && o.ExternalReference != nil {
		return true
	}

	return false
}

// SetExternalReference gets a reference to the given bool and assigns it to the ExternalReference field.
func (o *BTExternalReference1936) SetExternalReference(v bool) {
	o.ExternalReference = &v
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *BTExternalReference1936) GetFullElementId() BTFullElementId756 {
	if o == nil || o.FullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FullElementId
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalReference1936) GetFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FullElementId == nil {
		return nil, false
	}
	return o.FullElementId, true
}

// HasFullElementId returns a boolean if a field has been set.
func (o *BTExternalReference1936) HasFullElementId() bool {
	if o != nil && o.FullElementId != nil {
		return true
	}

	return false
}

// SetFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FullElementId field.
func (o *BTExternalReference1936) SetFullElementId(v BTFullElementId756) {
	o.FullElementId = &v
}

// GetMicroversionIdAndConfiguration returns the MicroversionIdAndConfiguration field value if set, zero value otherwise.
func (o *BTExternalReference1936) GetMicroversionIdAndConfiguration() BTMicroversionIdAndConfiguration2338 {
	if o == nil || o.MicroversionIdAndConfiguration == nil {
		var ret BTMicroversionIdAndConfiguration2338
		return ret
	}
	return *o.MicroversionIdAndConfiguration
}

// GetMicroversionIdAndConfigurationOk returns a tuple with the MicroversionIdAndConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalReference1936) GetMicroversionIdAndConfigurationOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	if o == nil || o.MicroversionIdAndConfiguration == nil {
		return nil, false
	}
	return o.MicroversionIdAndConfiguration, true
}

// HasMicroversionIdAndConfiguration returns a boolean if a field has been set.
func (o *BTExternalReference1936) HasMicroversionIdAndConfiguration() bool {
	if o != nil && o.MicroversionIdAndConfiguration != nil {
		return true
	}

	return false
}

// SetMicroversionIdAndConfiguration gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the MicroversionIdAndConfiguration field.
func (o *BTExternalReference1936) SetMicroversionIdAndConfiguration(v BTMicroversionIdAndConfiguration2338) {
	o.MicroversionIdAndConfiguration = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTExternalReference1936) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalReference1936) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTExternalReference1936) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTExternalReference1936) SetNodeId(v string) {
	o.NodeId = &v
}

// GetDocumentVersionId returns the DocumentVersionId field value if set, zero value otherwise.
func (o *BTExternalReference1936) GetDocumentVersionId() string {
	if o == nil || o.DocumentVersionId == nil {
		var ret string
		return ret
	}
	return *o.DocumentVersionId
}

// GetDocumentVersionIdOk returns a tuple with the DocumentVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalReference1936) GetDocumentVersionIdOk() (*string, bool) {
	if o == nil || o.DocumentVersionId == nil {
		return nil, false
	}
	return o.DocumentVersionId, true
}

// HasDocumentVersionId returns a boolean if a field has been set.
func (o *BTExternalReference1936) HasDocumentVersionId() bool {
	if o != nil && o.DocumentVersionId != nil {
		return true
	}

	return false
}

// SetDocumentVersionId gets a reference to the given string and assigns it to the DocumentVersionId field.
func (o *BTExternalReference1936) SetDocumentVersionId(v string) {
	o.DocumentVersionId = &v
}

func (o BTExternalReference1936) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Configured != nil {
		toSerialize["configured"] = o.Configured
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ExternalDocumentWithVersion != nil {
		toSerialize["externalDocumentWithVersion"] = o.ExternalDocumentWithVersion
	}
	if o.ExternalDocumentWithVersionAndElementId != nil {
		toSerialize["externalDocumentWithVersionAndElementId"] = o.ExternalDocumentWithVersionAndElementId
	}
	if o.ExternalReference != nil {
		toSerialize["externalReference"] = o.ExternalReference
	}
	if o.FullElementId != nil {
		toSerialize["fullElementId"] = o.FullElementId
	}
	if o.MicroversionIdAndConfiguration != nil {
		toSerialize["microversionIdAndConfiguration"] = o.MicroversionIdAndConfiguration
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.DocumentVersionId != nil {
		toSerialize["documentVersionId"] = o.DocumentVersionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTExternalReference1936 struct {
	value *BTExternalReference1936
	isSet bool
}

func (v NullableBTExternalReference1936) Get() *BTExternalReference1936 {
	return v.value
}

func (v *NullableBTExternalReference1936) Set(val *BTExternalReference1936) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExternalReference1936) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExternalReference1936) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExternalReference1936(val *BTExternalReference1936) *NullableBTExternalReference1936 {
	return &NullableBTExternalReference1936{value: val, isSet: true}
}

func (v NullableBTExternalReference1936) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExternalReference1936) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
