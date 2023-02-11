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

// BTClassroomInfo struct for BTClassroomInfo
type BTClassroomInfo struct {
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
	JsonType          string                  `json:"jsonType"`
	ModifiedAt        *JSONTime               `json:"modifiedAt,omitempty"`
	ModifiedBy        *BTUserBasicSummaryInfo `json:"modifiedBy,omitempty"`
	// Name of the resource.
	Name         *string      `json:"name,omitempty"`
	Owner        *BTOwnerInfo `json:"owner,omitempty"`
	ProjectId    *string      `json:"projectId,omitempty"`
	ResourceType *string      `json:"resourceType,omitempty"`
	TreeHref     *string      `json:"treeHref,omitempty"`
	UnparentHref *string      `json:"unparentHref,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef             *string `json:"viewRef,omitempty"`
	Educator            *bool   `json:"educator,omitempty"`
	NumberOfAssignments *int64  `json:"numberOfAssignments,omitempty"`
	NumberOfMembers     *int64  `json:"numberOfMembers,omitempty"`
	State               *int32  `json:"state,omitempty"`
	StateName           *string `json:"stateName,omitempty"`
	Title               *string `json:"title,omitempty"`
}

// NewBTClassroomInfo instantiates a new BTClassroomInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTClassroomInfo(jsonType string) *BTClassroomInfo {
	this := BTClassroomInfo{}
	this.JsonType = jsonType
	return &this
}

// NewBTClassroomInfoWithDefaults instantiates a new BTClassroomInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTClassroomInfoWithDefaults() *BTClassroomInfo {
	this := BTClassroomInfo{}
	return &this
}

// GetCanMove returns the CanMove field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetCanMove() bool {
	if o == nil || o.CanMove == nil {
		var ret bool
		return ret
	}
	return *o.CanMove
}

// GetCanMoveOk returns a tuple with the CanMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetCanMoveOk() (*bool, bool) {
	if o == nil || o.CanMove == nil {
		return nil, false
	}
	return o.CanMove, true
}

// HasCanMove returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasCanMove() bool {
	if o != nil && o.CanMove != nil {
		return true
	}

	return false
}

// SetCanMove gets a reference to the given bool and assigns it to the CanMove field.
func (o *BTClassroomInfo) SetCanMove(v bool) {
	o.CanMove = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTClassroomInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetCreatedBy() BTUserBasicSummaryInfo {
	if o == nil || o.CreatedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the CreatedBy field.
func (o *BTClassroomInfo) SetCreatedBy(v BTUserBasicSummaryInfo) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTClassroomInfo) SetDescription(v string) {
	o.Description = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTClassroomInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTClassroomInfo) SetId(v string) {
	o.Id = &v
}

// GetIsContainer returns the IsContainer field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetIsContainer() bool {
	if o == nil || o.IsContainer == nil {
		var ret bool
		return ret
	}
	return *o.IsContainer
}

// GetIsContainerOk returns a tuple with the IsContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetIsContainerOk() (*bool, bool) {
	if o == nil || o.IsContainer == nil {
		return nil, false
	}
	return o.IsContainer, true
}

// HasIsContainer returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasIsContainer() bool {
	if o != nil && o.IsContainer != nil {
		return true
	}

	return false
}

// SetIsContainer gets a reference to the given bool and assigns it to the IsContainer field.
func (o *BTClassroomInfo) SetIsContainer(v bool) {
	o.IsContainer = &v
}

// GetIsEnterpriseOwned returns the IsEnterpriseOwned field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetIsEnterpriseOwned() bool {
	if o == nil || o.IsEnterpriseOwned == nil {
		var ret bool
		return ret
	}
	return *o.IsEnterpriseOwned
}

// GetIsEnterpriseOwnedOk returns a tuple with the IsEnterpriseOwned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetIsEnterpriseOwnedOk() (*bool, bool) {
	if o == nil || o.IsEnterpriseOwned == nil {
		return nil, false
	}
	return o.IsEnterpriseOwned, true
}

// HasIsEnterpriseOwned returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasIsEnterpriseOwned() bool {
	if o != nil && o.IsEnterpriseOwned != nil {
		return true
	}

	return false
}

// SetIsEnterpriseOwned gets a reference to the given bool and assigns it to the IsEnterpriseOwned field.
func (o *BTClassroomInfo) SetIsEnterpriseOwned(v bool) {
	o.IsEnterpriseOwned = &v
}

// GetIsMutable returns the IsMutable field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetIsMutable() bool {
	if o == nil || o.IsMutable == nil {
		var ret bool
		return ret
	}
	return *o.IsMutable
}

// GetIsMutableOk returns a tuple with the IsMutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetIsMutableOk() (*bool, bool) {
	if o == nil || o.IsMutable == nil {
		return nil, false
	}
	return o.IsMutable, true
}

// HasIsMutable returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasIsMutable() bool {
	if o != nil && o.IsMutable != nil {
		return true
	}

	return false
}

// SetIsMutable gets a reference to the given bool and assigns it to the IsMutable field.
func (o *BTClassroomInfo) SetIsMutable(v bool) {
	o.IsMutable = &v
}

// GetJsonType returns the JsonType field value
func (o *BTClassroomInfo) GetJsonType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetJsonTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonType, true
}

// SetJsonType sets field value
func (o *BTClassroomInfo) SetJsonType(v string) {
	o.JsonType = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetModifiedAt() JSONTime {
	if o == nil || o.ModifiedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetModifiedAtOk() (*JSONTime, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given JSONTime and assigns it to the ModifiedAt field.
func (o *BTClassroomInfo) SetModifiedAt(v JSONTime) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetModifiedBy() BTUserBasicSummaryInfo {
	if o == nil || o.ModifiedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the ModifiedBy field.
func (o *BTClassroomInfo) SetModifiedBy(v BTUserBasicSummaryInfo) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTClassroomInfo) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetOwner() BTOwnerInfo {
	if o == nil || o.Owner == nil {
		var ret BTOwnerInfo
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetOwnerOk() (*BTOwnerInfo, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given BTOwnerInfo and assigns it to the Owner field.
func (o *BTClassroomInfo) SetOwner(v BTOwnerInfo) {
	o.Owner = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTClassroomInfo) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *BTClassroomInfo) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetTreeHref returns the TreeHref field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetTreeHref() string {
	if o == nil || o.TreeHref == nil {
		var ret string
		return ret
	}
	return *o.TreeHref
}

// GetTreeHrefOk returns a tuple with the TreeHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetTreeHrefOk() (*string, bool) {
	if o == nil || o.TreeHref == nil {
		return nil, false
	}
	return o.TreeHref, true
}

// HasTreeHref returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasTreeHref() bool {
	if o != nil && o.TreeHref != nil {
		return true
	}

	return false
}

// SetTreeHref gets a reference to the given string and assigns it to the TreeHref field.
func (o *BTClassroomInfo) SetTreeHref(v string) {
	o.TreeHref = &v
}

// GetUnparentHref returns the UnparentHref field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetUnparentHref() string {
	if o == nil || o.UnparentHref == nil {
		var ret string
		return ret
	}
	return *o.UnparentHref
}

// GetUnparentHrefOk returns a tuple with the UnparentHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetUnparentHrefOk() (*string, bool) {
	if o == nil || o.UnparentHref == nil {
		return nil, false
	}
	return o.UnparentHref, true
}

// HasUnparentHref returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasUnparentHref() bool {
	if o != nil && o.UnparentHref != nil {
		return true
	}

	return false
}

// SetUnparentHref gets a reference to the given string and assigns it to the UnparentHref field.
func (o *BTClassroomInfo) SetUnparentHref(v string) {
	o.UnparentHref = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTClassroomInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetEducator returns the Educator field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetEducator() bool {
	if o == nil || o.Educator == nil {
		var ret bool
		return ret
	}
	return *o.Educator
}

// GetEducatorOk returns a tuple with the Educator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetEducatorOk() (*bool, bool) {
	if o == nil || o.Educator == nil {
		return nil, false
	}
	return o.Educator, true
}

// HasEducator returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasEducator() bool {
	if o != nil && o.Educator != nil {
		return true
	}

	return false
}

// SetEducator gets a reference to the given bool and assigns it to the Educator field.
func (o *BTClassroomInfo) SetEducator(v bool) {
	o.Educator = &v
}

// GetNumberOfAssignments returns the NumberOfAssignments field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetNumberOfAssignments() int64 {
	if o == nil || o.NumberOfAssignments == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfAssignments
}

// GetNumberOfAssignmentsOk returns a tuple with the NumberOfAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetNumberOfAssignmentsOk() (*int64, bool) {
	if o == nil || o.NumberOfAssignments == nil {
		return nil, false
	}
	return o.NumberOfAssignments, true
}

// HasNumberOfAssignments returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasNumberOfAssignments() bool {
	if o != nil && o.NumberOfAssignments != nil {
		return true
	}

	return false
}

// SetNumberOfAssignments gets a reference to the given int64 and assigns it to the NumberOfAssignments field.
func (o *BTClassroomInfo) SetNumberOfAssignments(v int64) {
	o.NumberOfAssignments = &v
}

// GetNumberOfMembers returns the NumberOfMembers field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetNumberOfMembers() int64 {
	if o == nil || o.NumberOfMembers == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfMembers
}

// GetNumberOfMembersOk returns a tuple with the NumberOfMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetNumberOfMembersOk() (*int64, bool) {
	if o == nil || o.NumberOfMembers == nil {
		return nil, false
	}
	return o.NumberOfMembers, true
}

// HasNumberOfMembers returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasNumberOfMembers() bool {
	if o != nil && o.NumberOfMembers != nil {
		return true
	}

	return false
}

// SetNumberOfMembers gets a reference to the given int64 and assigns it to the NumberOfMembers field.
func (o *BTClassroomInfo) SetNumberOfMembers(v int64) {
	o.NumberOfMembers = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTClassroomInfo) SetState(v int32) {
	o.State = &v
}

// GetStateName returns the StateName field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetStateName() string {
	if o == nil || o.StateName == nil {
		var ret string
		return ret
	}
	return *o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetStateNameOk() (*string, bool) {
	if o == nil || o.StateName == nil {
		return nil, false
	}
	return o.StateName, true
}

// HasStateName returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasStateName() bool {
	if o != nil && o.StateName != nil {
		return true
	}

	return false
}

// SetStateName gets a reference to the given string and assigns it to the StateName field.
func (o *BTClassroomInfo) SetStateName(v string) {
	o.StateName = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BTClassroomInfo) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClassroomInfo) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BTClassroomInfo) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BTClassroomInfo) SetTitle(v string) {
	o.Title = &v
}

func (o BTClassroomInfo) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["jsonType"] = o.JsonType
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
	if o.Educator != nil {
		toSerialize["educator"] = o.Educator
	}
	if o.NumberOfAssignments != nil {
		toSerialize["numberOfAssignments"] = o.NumberOfAssignments
	}
	if o.NumberOfMembers != nil {
		toSerialize["numberOfMembers"] = o.NumberOfMembers
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.StateName != nil {
		toSerialize["stateName"] = o.StateName
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableBTClassroomInfo struct {
	value *BTClassroomInfo
	isSet bool
}

func (v NullableBTClassroomInfo) Get() *BTClassroomInfo {
	return v.value
}

func (v *NullableBTClassroomInfo) Set(val *BTClassroomInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTClassroomInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTClassroomInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTClassroomInfo(val *BTClassroomInfo) *NullableBTClassroomInfo {
	return &NullableBTClassroomInfo{value: val, isSet: true}
}

func (v NullableBTClassroomInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTClassroomInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
