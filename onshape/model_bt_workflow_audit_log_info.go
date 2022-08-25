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

// BTWorkflowAuditLogInfo struct for BTWorkflowAuditLogInfo
type BTWorkflowAuditLogInfo struct {
	CompanyId           *string                       `json:"companyId,omitempty"`
	DebugMicroversionId *string                       `json:"debugMicroversionId,omitempty"`
	Entries             []BTWorkflowAuditLogEntryInfo `json:"entries,omitempty"`
	ObjectId            *string                       `json:"objectId,omitempty"`
	ObjectType          *int32                        `json:"objectType,omitempty"`
	PublishedWorkflowId *BTPublishedWorkflowId        `json:"publishedWorkflowId,omitempty"`
}

// NewBTWorkflowAuditLogInfo instantiates a new BTWorkflowAuditLogInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWorkflowAuditLogInfo() *BTWorkflowAuditLogInfo {
	this := BTWorkflowAuditLogInfo{}
	return &this
}

// NewBTWorkflowAuditLogInfoWithDefaults instantiates a new BTWorkflowAuditLogInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWorkflowAuditLogInfoWithDefaults() *BTWorkflowAuditLogInfo {
	this := BTWorkflowAuditLogInfo{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogInfo) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogInfo) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogInfo) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTWorkflowAuditLogInfo) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetDebugMicroversionId returns the DebugMicroversionId field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogInfo) GetDebugMicroversionId() string {
	if o == nil || o.DebugMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.DebugMicroversionId
}

// GetDebugMicroversionIdOk returns a tuple with the DebugMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogInfo) GetDebugMicroversionIdOk() (*string, bool) {
	if o == nil || o.DebugMicroversionId == nil {
		return nil, false
	}
	return o.DebugMicroversionId, true
}

// HasDebugMicroversionId returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogInfo) HasDebugMicroversionId() bool {
	if o != nil && o.DebugMicroversionId != nil {
		return true
	}

	return false
}

// SetDebugMicroversionId gets a reference to the given string and assigns it to the DebugMicroversionId field.
func (o *BTWorkflowAuditLogInfo) SetDebugMicroversionId(v string) {
	o.DebugMicroversionId = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogInfo) GetEntries() []BTWorkflowAuditLogEntryInfo {
	if o == nil || o.Entries == nil {
		var ret []BTWorkflowAuditLogEntryInfo
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogInfo) GetEntriesOk() ([]BTWorkflowAuditLogEntryInfo, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogInfo) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []BTWorkflowAuditLogEntryInfo and assigns it to the Entries field.
func (o *BTWorkflowAuditLogInfo) SetEntries(v []BTWorkflowAuditLogEntryInfo) {
	o.Entries = v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogInfo) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogInfo) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogInfo) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *BTWorkflowAuditLogInfo) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogInfo) GetObjectType() int32 {
	if o == nil || o.ObjectType == nil {
		var ret int32
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogInfo) GetObjectTypeOk() (*int32, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogInfo) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given int32 and assigns it to the ObjectType field.
func (o *BTWorkflowAuditLogInfo) SetObjectType(v int32) {
	o.ObjectType = &v
}

// GetPublishedWorkflowId returns the PublishedWorkflowId field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogInfo) GetPublishedWorkflowId() BTPublishedWorkflowId {
	if o == nil || o.PublishedWorkflowId == nil {
		var ret BTPublishedWorkflowId
		return ret
	}
	return *o.PublishedWorkflowId
}

// GetPublishedWorkflowIdOk returns a tuple with the PublishedWorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogInfo) GetPublishedWorkflowIdOk() (*BTPublishedWorkflowId, bool) {
	if o == nil || o.PublishedWorkflowId == nil {
		return nil, false
	}
	return o.PublishedWorkflowId, true
}

// HasPublishedWorkflowId returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogInfo) HasPublishedWorkflowId() bool {
	if o != nil && o.PublishedWorkflowId != nil {
		return true
	}

	return false
}

// SetPublishedWorkflowId gets a reference to the given BTPublishedWorkflowId and assigns it to the PublishedWorkflowId field.
func (o *BTWorkflowAuditLogInfo) SetPublishedWorkflowId(v BTPublishedWorkflowId) {
	o.PublishedWorkflowId = &v
}

func (o BTWorkflowAuditLogInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.DebugMicroversionId != nil {
		toSerialize["debugMicroversionId"] = o.DebugMicroversionId
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.PublishedWorkflowId != nil {
		toSerialize["publishedWorkflowId"] = o.PublishedWorkflowId
	}
	return json.Marshal(toSerialize)
}

type NullableBTWorkflowAuditLogInfo struct {
	value *BTWorkflowAuditLogInfo
	isSet bool
}

func (v NullableBTWorkflowAuditLogInfo) Get() *BTWorkflowAuditLogInfo {
	return v.value
}

func (v *NullableBTWorkflowAuditLogInfo) Set(val *BTWorkflowAuditLogInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkflowAuditLogInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkflowAuditLogInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkflowAuditLogInfo(val *BTWorkflowAuditLogInfo) *NullableBTWorkflowAuditLogInfo {
	return &NullableBTWorkflowAuditLogInfo{value: val, isSet: true}
}

func (v NullableBTWorkflowAuditLogInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkflowAuditLogInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
