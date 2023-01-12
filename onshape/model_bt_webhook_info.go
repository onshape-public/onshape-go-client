/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.158.9878-ff51e1211d95
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTWebhookInfo struct for BTWebhookInfo
type BTWebhookInfo struct {
	CompanyId         *string            `json:"companyId,omitempty"`
	CreatedBy         *BTUserSummaryInfo `json:"createdBy,omitempty"`
	Data              *string            `json:"data,omitempty"`
	Description       *string            `json:"description,omitempty"`
	DroppedEventCount *int32             `json:"droppedEventCount,omitempty"`
	Events            []string           `json:"events,omitempty"`
	Filter            *string            `json:"filter,omitempty"`
	FolderId          *string            `json:"folderId,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// Name of the resource.
	Name      *string           `json:"name,omitempty"`
	Options   *BTWebhookOptions `json:"options,omitempty"`
	ProjectId *string           `json:"projectId,omitempty"`
	Url       *string           `json:"url,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTWebhookInfo instantiates a new BTWebhookInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWebhookInfo() *BTWebhookInfo {
	this := BTWebhookInfo{}
	return &this
}

// NewBTWebhookInfoWithDefaults instantiates a new BTWebhookInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWebhookInfoWithDefaults() *BTWebhookInfo {
	this := BTWebhookInfo{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTWebhookInfo) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetCreatedBy() BTUserSummaryInfo {
	if o == nil || o.CreatedBy == nil {
		var ret BTUserSummaryInfo
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetCreatedByOk() (*BTUserSummaryInfo, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given BTUserSummaryInfo and assigns it to the CreatedBy field.
func (o *BTWebhookInfo) SetCreatedBy(v BTUserSummaryInfo) {
	o.CreatedBy = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *BTWebhookInfo) SetData(v string) {
	o.Data = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTWebhookInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDroppedEventCount returns the DroppedEventCount field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetDroppedEventCount() int32 {
	if o == nil || o.DroppedEventCount == nil {
		var ret int32
		return ret
	}
	return *o.DroppedEventCount
}

// GetDroppedEventCountOk returns a tuple with the DroppedEventCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetDroppedEventCountOk() (*int32, bool) {
	if o == nil || o.DroppedEventCount == nil {
		return nil, false
	}
	return o.DroppedEventCount, true
}

// HasDroppedEventCount returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasDroppedEventCount() bool {
	if o != nil && o.DroppedEventCount != nil {
		return true
	}

	return false
}

// SetDroppedEventCount gets a reference to the given int32 and assigns it to the DroppedEventCount field.
func (o *BTWebhookInfo) SetDroppedEventCount(v int32) {
	o.DroppedEventCount = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetEvents() []string {
	if o == nil || o.Events == nil {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetEventsOk() ([]string, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *BTWebhookInfo) SetEvents(v []string) {
	o.Events = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *BTWebhookInfo) SetFilter(v string) {
	o.Filter = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetFolderId() string {
	if o == nil || o.FolderId == nil {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetFolderIdOk() (*string, bool) {
	if o == nil || o.FolderId == nil {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasFolderId() bool {
	if o != nil && o.FolderId != nil {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *BTWebhookInfo) SetFolderId(v string) {
	o.FolderId = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTWebhookInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTWebhookInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTWebhookInfo) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetOptions() BTWebhookOptions {
	if o == nil || o.Options == nil {
		var ret BTWebhookOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetOptionsOk() (*BTWebhookOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given BTWebhookOptions and assigns it to the Options field.
func (o *BTWebhookInfo) SetOptions(v BTWebhookOptions) {
	o.Options = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTWebhookInfo) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *BTWebhookInfo) SetUrl(v string) {
	o.Url = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTWebhookInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTWebhookInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTWebhookInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTWebhookInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DroppedEventCount != nil {
		toSerialize["droppedEventCount"] = o.DroppedEventCount
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.FolderId != nil {
		toSerialize["folderId"] = o.FolderId
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTWebhookInfo struct {
	value *BTWebhookInfo
	isSet bool
}

func (v NullableBTWebhookInfo) Get() *BTWebhookInfo {
	return v.value
}

func (v *NullableBTWebhookInfo) Set(val *BTWebhookInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWebhookInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWebhookInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWebhookInfo(val *BTWebhookInfo) *NullableBTWebhookInfo {
	return &NullableBTWebhookInfo{value: val, isSet: true}
}

func (v NullableBTWebhookInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWebhookInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
