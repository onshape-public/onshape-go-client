/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.175.29152-5591e93bd1c1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTNextPartNumberParam struct for BTNextPartNumberParam
type BTNextPartNumberParam struct {
	Categories                 []BTCategoryParam `json:"categories,omitempty"`
	Configuration              *string           `json:"configuration,omitempty"`
	DocumentId                 *string           `json:"documentId,omitempty"`
	ElementId                  *string           `json:"elementId,omitempty"`
	ElementType                *int32            `json:"elementType,omitempty"`
	Id                         *string           `json:"id,omitempty"`
	MimeType                   *string           `json:"mimeType,omitempty"`
	NumberSchemeResourceTypeId *string           `json:"numberSchemeResourceTypeId,omitempty"`
	PartId                     *string           `json:"partId,omitempty"`
	PartNumber                 *string           `json:"partNumber,omitempty"`
	VersionId                  *string           `json:"versionId,omitempty"`
	WorkspaceId                *string           `json:"workspaceId,omitempty"`
}

// NewBTNextPartNumberParam instantiates a new BTNextPartNumberParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTNextPartNumberParam() *BTNextPartNumberParam {
	this := BTNextPartNumberParam{}
	return &this
}

// NewBTNextPartNumberParamWithDefaults instantiates a new BTNextPartNumberParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTNextPartNumberParamWithDefaults() *BTNextPartNumberParam {
	this := BTNextPartNumberParam{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *BTNextPartNumberParam) GetCategories() []BTCategoryParam {
	if o == nil || o.Categories == nil {
		var ret []BTCategoryParam
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNextPartNumberParam) GetCategoriesOk() ([]BTCategoryParam, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *BTNextPartNumberParam) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []BTCategoryParam and assigns it to the Categories field.
func (o *BTNextPartNumberParam) SetCategories(v []BTCategoryParam) {
	o.Categories = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTNextPartNumberParam) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNextPartNumberParam) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTNextPartNumberParam) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTNextPartNumberParam) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTNextPartNumberParam) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNextPartNumberParam) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTNextPartNumberParam) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTNextPartNumberParam) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTNextPartNumberParam) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNextPartNumberParam) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTNextPartNumberParam) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTNextPartNumberParam) SetElementId(v string) {
	o.ElementId = &v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *BTNextPartNumberParam) GetElementType() int32 {
	if o == nil || o.ElementType == nil {
		var ret int32
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNextPartNumberParam) GetElementTypeOk() (*int32, bool) {
	if o == nil || o.ElementType == nil {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *BTNextPartNumberParam) HasElementType() bool {
	if o != nil && o.ElementType != nil {
		return true
	}

	return false
}

// SetElementType gets a reference to the given int32 and assigns it to the ElementType field.
func (o *BTNextPartNumberParam) SetElementType(v int32) {
	o.ElementType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTNextPartNumberParam) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNextPartNumberParam) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTNextPartNumberParam) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTNextPartNumberParam) SetId(v string) {
	o.Id = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *BTNextPartNumberParam) GetMimeType() string {
	if o == nil || o.MimeType == nil {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNextPartNumberParam) GetMimeTypeOk() (*string, bool) {
	if o == nil || o.MimeType == nil {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *BTNextPartNumberParam) HasMimeType() bool {
	if o != nil && o.MimeType != nil {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *BTNextPartNumberParam) SetMimeType(v string) {
	o.MimeType = &v
}

// GetNumberSchemeResourceTypeId returns the NumberSchemeResourceTypeId field value if set, zero value otherwise.
func (o *BTNextPartNumberParam) GetNumberSchemeResourceTypeId() string {
	if o == nil || o.NumberSchemeResourceTypeId == nil {
		var ret string
		return ret
	}
	return *o.NumberSchemeResourceTypeId
}

// GetNumberSchemeResourceTypeIdOk returns a tuple with the NumberSchemeResourceTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNextPartNumberParam) GetNumberSchemeResourceTypeIdOk() (*string, bool) {
	if o == nil || o.NumberSchemeResourceTypeId == nil {
		return nil, false
	}
	return o.NumberSchemeResourceTypeId, true
}

// HasNumberSchemeResourceTypeId returns a boolean if a field has been set.
func (o *BTNextPartNumberParam) HasNumberSchemeResourceTypeId() bool {
	if o != nil && o.NumberSchemeResourceTypeId != nil {
		return true
	}

	return false
}

// SetNumberSchemeResourceTypeId gets a reference to the given string and assigns it to the NumberSchemeResourceTypeId field.
func (o *BTNextPartNumberParam) SetNumberSchemeResourceTypeId(v string) {
	o.NumberSchemeResourceTypeId = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTNextPartNumberParam) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNextPartNumberParam) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTNextPartNumberParam) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTNextPartNumberParam) SetPartId(v string) {
	o.PartId = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTNextPartNumberParam) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNextPartNumberParam) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTNextPartNumberParam) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTNextPartNumberParam) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTNextPartNumberParam) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNextPartNumberParam) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTNextPartNumberParam) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTNextPartNumberParam) SetVersionId(v string) {
	o.VersionId = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTNextPartNumberParam) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNextPartNumberParam) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTNextPartNumberParam) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTNextPartNumberParam) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTNextPartNumberParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ElementType != nil {
		toSerialize["elementType"] = o.ElementType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MimeType != nil {
		toSerialize["mimeType"] = o.MimeType
	}
	if o.NumberSchemeResourceTypeId != nil {
		toSerialize["numberSchemeResourceTypeId"] = o.NumberSchemeResourceTypeId
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTNextPartNumberParam struct {
	value *BTNextPartNumberParam
	isSet bool
}

func (v NullableBTNextPartNumberParam) Get() *BTNextPartNumberParam {
	return v.value
}

func (v *NullableBTNextPartNumberParam) Set(val *BTNextPartNumberParam) {
	v.value = val
	v.isSet = true
}

func (v NullableBTNextPartNumberParam) IsSet() bool {
	return v.isSet
}

func (v *NullableBTNextPartNumberParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTNextPartNumberParam(val *BTNextPartNumberParam) *NullableBTNextPartNumberParam {
	return &NullableBTNextPartNumberParam{value: val, isSet: true}
}

func (v NullableBTNextPartNumberParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTNextPartNumberParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
