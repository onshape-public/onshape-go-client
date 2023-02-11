/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.159.11153-01265f451c50
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDocumentLabelInfo struct for BTDocumentLabelInfo
type BTDocumentLabelInfo struct {
	CanMove     *bool                   `json:"canMove,omitempty"`
	CreatedAt   *JSONTime               `json:"createdAt,omitempty"`
	CreatedBy   *BTUserBasicSummaryInfo `json:"createdBy,omitempty"`
	Description *string                 `json:"description,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id                *string                 `json:"id,omitempty"`
	IsContainer       *bool                   `json:"isContainer,omitempty"`
	IsEnterpriseOwned *bool                   `json:"isEnterpriseOwned,omitempty"`
	IsMutable         *bool                   `json:"isMutable,omitempty"`
	ModifiedAt        *JSONTime               `json:"modifiedAt,omitempty"`
	ModifiedBy        *BTUserBasicSummaryInfo `json:"modifiedBy,omitempty"`
	// Name of the resource.
	Name          *string      `json:"name,omitempty"`
	Owner         *BTOwnerInfo `json:"owner,omitempty"`
	ParentLabelId *string      `json:"parentLabelId,omitempty"`
	Path          *string      `json:"path,omitempty"`
	ProjectId     *string      `json:"projectId,omitempty"`
	ResourceType  *string      `json:"resourceType,omitempty"`
	TreeHref      *string      `json:"treeHref,omitempty"`
	UnparentHref  *string      `json:"unparentHref,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTDocumentLabelInfo instantiates a new BTDocumentLabelInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentLabelInfo() *BTDocumentLabelInfo {
	this := BTDocumentLabelInfo{}
	return &this
}

// NewBTDocumentLabelInfoWithDefaults instantiates a new BTDocumentLabelInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentLabelInfoWithDefaults() *BTDocumentLabelInfo {
	this := BTDocumentLabelInfo{}
	return &this
}

// GetCanMove returns the CanMove field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetCanMove() bool {
	if o == nil || o.CanMove == nil {
		var ret bool
		return ret
	}
	return *o.CanMove
}

// GetCanMoveOk returns a tuple with the CanMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetCanMoveOk() (*bool, bool) {
	if o == nil || o.CanMove == nil {
		return nil, false
	}
	return o.CanMove, true
}

// HasCanMove returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasCanMove() bool {
	if o != nil && o.CanMove != nil {
		return true
	}

	return false
}

// SetCanMove gets a reference to the given bool and assigns it to the CanMove field.
func (o *BTDocumentLabelInfo) SetCanMove(v bool) {
	o.CanMove = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTDocumentLabelInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetCreatedBy() BTUserBasicSummaryInfo {
	if o == nil || o.CreatedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the CreatedBy field.
func (o *BTDocumentLabelInfo) SetCreatedBy(v BTUserBasicSummaryInfo) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTDocumentLabelInfo) SetDescription(v string) {
	o.Description = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTDocumentLabelInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTDocumentLabelInfo) SetId(v string) {
	o.Id = &v
}

// GetIsContainer returns the IsContainer field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetIsContainer() bool {
	if o == nil || o.IsContainer == nil {
		var ret bool
		return ret
	}
	return *o.IsContainer
}

// GetIsContainerOk returns a tuple with the IsContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetIsContainerOk() (*bool, bool) {
	if o == nil || o.IsContainer == nil {
		return nil, false
	}
	return o.IsContainer, true
}

// HasIsContainer returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasIsContainer() bool {
	if o != nil && o.IsContainer != nil {
		return true
	}

	return false
}

// SetIsContainer gets a reference to the given bool and assigns it to the IsContainer field.
func (o *BTDocumentLabelInfo) SetIsContainer(v bool) {
	o.IsContainer = &v
}

// GetIsEnterpriseOwned returns the IsEnterpriseOwned field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetIsEnterpriseOwned() bool {
	if o == nil || o.IsEnterpriseOwned == nil {
		var ret bool
		return ret
	}
	return *o.IsEnterpriseOwned
}

// GetIsEnterpriseOwnedOk returns a tuple with the IsEnterpriseOwned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetIsEnterpriseOwnedOk() (*bool, bool) {
	if o == nil || o.IsEnterpriseOwned == nil {
		return nil, false
	}
	return o.IsEnterpriseOwned, true
}

// HasIsEnterpriseOwned returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasIsEnterpriseOwned() bool {
	if o != nil && o.IsEnterpriseOwned != nil {
		return true
	}

	return false
}

// SetIsEnterpriseOwned gets a reference to the given bool and assigns it to the IsEnterpriseOwned field.
func (o *BTDocumentLabelInfo) SetIsEnterpriseOwned(v bool) {
	o.IsEnterpriseOwned = &v
}

// GetIsMutable returns the IsMutable field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetIsMutable() bool {
	if o == nil || o.IsMutable == nil {
		var ret bool
		return ret
	}
	return *o.IsMutable
}

// GetIsMutableOk returns a tuple with the IsMutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetIsMutableOk() (*bool, bool) {
	if o == nil || o.IsMutable == nil {
		return nil, false
	}
	return o.IsMutable, true
}

// HasIsMutable returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasIsMutable() bool {
	if o != nil && o.IsMutable != nil {
		return true
	}

	return false
}

// SetIsMutable gets a reference to the given bool and assigns it to the IsMutable field.
func (o *BTDocumentLabelInfo) SetIsMutable(v bool) {
	o.IsMutable = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetModifiedAt() JSONTime {
	if o == nil || o.ModifiedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetModifiedAtOk() (*JSONTime, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given JSONTime and assigns it to the ModifiedAt field.
func (o *BTDocumentLabelInfo) SetModifiedAt(v JSONTime) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetModifiedBy() BTUserBasicSummaryInfo {
	if o == nil || o.ModifiedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the ModifiedBy field.
func (o *BTDocumentLabelInfo) SetModifiedBy(v BTUserBasicSummaryInfo) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTDocumentLabelInfo) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetOwner() BTOwnerInfo {
	if o == nil || o.Owner == nil {
		var ret BTOwnerInfo
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetOwnerOk() (*BTOwnerInfo, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given BTOwnerInfo and assigns it to the Owner field.
func (o *BTDocumentLabelInfo) SetOwner(v BTOwnerInfo) {
	o.Owner = &v
}

// GetParentLabelId returns the ParentLabelId field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetParentLabelId() string {
	if o == nil || o.ParentLabelId == nil {
		var ret string
		return ret
	}
	return *o.ParentLabelId
}

// GetParentLabelIdOk returns a tuple with the ParentLabelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetParentLabelIdOk() (*string, bool) {
	if o == nil || o.ParentLabelId == nil {
		return nil, false
	}
	return o.ParentLabelId, true
}

// HasParentLabelId returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasParentLabelId() bool {
	if o != nil && o.ParentLabelId != nil {
		return true
	}

	return false
}

// SetParentLabelId gets a reference to the given string and assigns it to the ParentLabelId field.
func (o *BTDocumentLabelInfo) SetParentLabelId(v string) {
	o.ParentLabelId = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *BTDocumentLabelInfo) SetPath(v string) {
	o.Path = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTDocumentLabelInfo) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *BTDocumentLabelInfo) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetTreeHref returns the TreeHref field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetTreeHref() string {
	if o == nil || o.TreeHref == nil {
		var ret string
		return ret
	}
	return *o.TreeHref
}

// GetTreeHrefOk returns a tuple with the TreeHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetTreeHrefOk() (*string, bool) {
	if o == nil || o.TreeHref == nil {
		return nil, false
	}
	return o.TreeHref, true
}

// HasTreeHref returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasTreeHref() bool {
	if o != nil && o.TreeHref != nil {
		return true
	}

	return false
}

// SetTreeHref gets a reference to the given string and assigns it to the TreeHref field.
func (o *BTDocumentLabelInfo) SetTreeHref(v string) {
	o.TreeHref = &v
}

// GetUnparentHref returns the UnparentHref field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetUnparentHref() string {
	if o == nil || o.UnparentHref == nil {
		var ret string
		return ret
	}
	return *o.UnparentHref
}

// GetUnparentHrefOk returns a tuple with the UnparentHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetUnparentHrefOk() (*string, bool) {
	if o == nil || o.UnparentHref == nil {
		return nil, false
	}
	return o.UnparentHref, true
}

// HasUnparentHref returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasUnparentHref() bool {
	if o != nil && o.UnparentHref != nil {
		return true
	}

	return false
}

// SetUnparentHref gets a reference to the given string and assigns it to the UnparentHref field.
func (o *BTDocumentLabelInfo) SetUnparentHref(v string) {
	o.UnparentHref = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTDocumentLabelInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentLabelInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTDocumentLabelInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTDocumentLabelInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTDocumentLabelInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanMove != nil {
		toSerialize["canMove"] = o.CanMove
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsContainer != nil {
		toSerialize["isContainer"] = o.IsContainer
	}
	if o.IsEnterpriseOwned != nil {
		toSerialize["isEnterpriseOwned"] = o.IsEnterpriseOwned
	}
	if o.IsMutable != nil {
		toSerialize["isMutable"] = o.IsMutable
	}
	if o.ModifiedAt != nil {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if o.ModifiedBy != nil {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.ParentLabelId != nil {
		toSerialize["parentLabelId"] = o.ParentLabelId
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.TreeHref != nil {
		toSerialize["treeHref"] = o.TreeHref
	}
	if o.UnparentHref != nil {
		toSerialize["unparentHref"] = o.UnparentHref
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentLabelInfo struct {
	value *BTDocumentLabelInfo
	isSet bool
}

func (v NullableBTDocumentLabelInfo) Get() *BTDocumentLabelInfo {
	return v.value
}

func (v *NullableBTDocumentLabelInfo) Set(val *BTDocumentLabelInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentLabelInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentLabelInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentLabelInfo(val *BTDocumentLabelInfo) *NullableBTDocumentLabelInfo {
	return &NullableBTDocumentLabelInfo{value: val, isSet: true}
}

func (v NullableBTDocumentLabelInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentLabelInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
