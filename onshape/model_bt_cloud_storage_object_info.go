/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23196-ae5f57bec467
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCloudStorageObjectInfo struct for BTCloudStorageObjectInfo
type BTCloudStorageObjectInfo struct {
	CanMove               *bool                   `json:"canMove,omitempty"`
	CloudStorageAccountId *string                 `json:"cloudStorageAccountId,omitempty"`
	CloudStorageObjectId  *string                 `json:"cloudStorageObjectId,omitempty"`
	CloudStorageProvider  *int32                  `json:"cloudStorageProvider,omitempty"`
	CreatedAt             *JSONTime               `json:"createdAt,omitempty"`
	CreatedBy             *BTUserBasicSummaryInfo `json:"createdBy,omitempty"`
	CreatedById           *string                 `json:"createdById,omitempty"`
	Description           *string                 `json:"description,omitempty"`
	// URI to fetch complete information of the resource.
	Href     *string `json:"href,omitempty"`
	IconLink *string `json:"iconLink,omitempty"`
	// Id of the resource.
	Id                *string                 `json:"id,omitempty"`
	IsContainer       *bool                   `json:"isContainer,omitempty"`
	IsEnterpriseOwned *bool                   `json:"isEnterpriseOwned,omitempty"`
	IsMutable         *bool                   `json:"isMutable,omitempty"`
	MimeType          *string                 `json:"mimeType,omitempty"`
	ModifiedAt        *JSONTime               `json:"modifiedAt,omitempty"`
	ModifiedBy        *BTUserBasicSummaryInfo `json:"modifiedBy,omitempty"`
	ModifiedById      *string                 `json:"modifiedById,omitempty"`
	// Name of the resource.
	Name          *string          `json:"name,omitempty"`
	Owner         *BTOwnerInfo     `json:"owner,omitempty"`
	ParentId      *string          `json:"parentId,omitempty"`
	ProjectId     *string          `json:"projectId,omitempty"`
	ResourceType  *string          `json:"resourceType,omitempty"`
	SizeBytes     *int64           `json:"sizeBytes,omitempty"`
	ThumbnailInfo *BTThumbnailInfo `json:"thumbnailInfo,omitempty"`
	TreeHref      *string          `json:"treeHref,omitempty"`
	UnparentHref  *string          `json:"unparentHref,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef     *string `json:"viewRef,omitempty"`
	WebViewLink *string `json:"webViewLink,omitempty"`
}

// NewBTCloudStorageObjectInfo instantiates a new BTCloudStorageObjectInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCloudStorageObjectInfo() *BTCloudStorageObjectInfo {
	this := BTCloudStorageObjectInfo{}
	return &this
}

// NewBTCloudStorageObjectInfoWithDefaults instantiates a new BTCloudStorageObjectInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCloudStorageObjectInfoWithDefaults() *BTCloudStorageObjectInfo {
	this := BTCloudStorageObjectInfo{}
	return &this
}

// GetCanMove returns the CanMove field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetCanMove() bool {
	if o == nil || o.CanMove == nil {
		var ret bool
		return ret
	}
	return *o.CanMove
}

// GetCanMoveOk returns a tuple with the CanMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetCanMoveOk() (*bool, bool) {
	if o == nil || o.CanMove == nil {
		return nil, false
	}
	return o.CanMove, true
}

// HasCanMove returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasCanMove() bool {
	if o != nil && o.CanMove != nil {
		return true
	}

	return false
}

// SetCanMove gets a reference to the given bool and assigns it to the CanMove field.
func (o *BTCloudStorageObjectInfo) SetCanMove(v bool) {
	o.CanMove = &v
}

// GetCloudStorageAccountId returns the CloudStorageAccountId field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetCloudStorageAccountId() string {
	if o == nil || o.CloudStorageAccountId == nil {
		var ret string
		return ret
	}
	return *o.CloudStorageAccountId
}

// GetCloudStorageAccountIdOk returns a tuple with the CloudStorageAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetCloudStorageAccountIdOk() (*string, bool) {
	if o == nil || o.CloudStorageAccountId == nil {
		return nil, false
	}
	return o.CloudStorageAccountId, true
}

// HasCloudStorageAccountId returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasCloudStorageAccountId() bool {
	if o != nil && o.CloudStorageAccountId != nil {
		return true
	}

	return false
}

// SetCloudStorageAccountId gets a reference to the given string and assigns it to the CloudStorageAccountId field.
func (o *BTCloudStorageObjectInfo) SetCloudStorageAccountId(v string) {
	o.CloudStorageAccountId = &v
}

// GetCloudStorageObjectId returns the CloudStorageObjectId field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetCloudStorageObjectId() string {
	if o == nil || o.CloudStorageObjectId == nil {
		var ret string
		return ret
	}
	return *o.CloudStorageObjectId
}

// GetCloudStorageObjectIdOk returns a tuple with the CloudStorageObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetCloudStorageObjectIdOk() (*string, bool) {
	if o == nil || o.CloudStorageObjectId == nil {
		return nil, false
	}
	return o.CloudStorageObjectId, true
}

// HasCloudStorageObjectId returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasCloudStorageObjectId() bool {
	if o != nil && o.CloudStorageObjectId != nil {
		return true
	}

	return false
}

// SetCloudStorageObjectId gets a reference to the given string and assigns it to the CloudStorageObjectId field.
func (o *BTCloudStorageObjectInfo) SetCloudStorageObjectId(v string) {
	o.CloudStorageObjectId = &v
}

// GetCloudStorageProvider returns the CloudStorageProvider field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetCloudStorageProvider() int32 {
	if o == nil || o.CloudStorageProvider == nil {
		var ret int32
		return ret
	}
	return *o.CloudStorageProvider
}

// GetCloudStorageProviderOk returns a tuple with the CloudStorageProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetCloudStorageProviderOk() (*int32, bool) {
	if o == nil || o.CloudStorageProvider == nil {
		return nil, false
	}
	return o.CloudStorageProvider, true
}

// HasCloudStorageProvider returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasCloudStorageProvider() bool {
	if o != nil && o.CloudStorageProvider != nil {
		return true
	}

	return false
}

// SetCloudStorageProvider gets a reference to the given int32 and assigns it to the CloudStorageProvider field.
func (o *BTCloudStorageObjectInfo) SetCloudStorageProvider(v int32) {
	o.CloudStorageProvider = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTCloudStorageObjectInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetCreatedBy() BTUserBasicSummaryInfo {
	if o == nil || o.CreatedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the CreatedBy field.
func (o *BTCloudStorageObjectInfo) SetCreatedBy(v BTUserBasicSummaryInfo) {
	o.CreatedBy = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetCreatedById() string {
	if o == nil || o.CreatedById == nil {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetCreatedByIdOk() (*string, bool) {
	if o == nil || o.CreatedById == nil {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasCreatedById() bool {
	if o != nil && o.CreatedById != nil {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *BTCloudStorageObjectInfo) SetCreatedById(v string) {
	o.CreatedById = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTCloudStorageObjectInfo) SetDescription(v string) {
	o.Description = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTCloudStorageObjectInfo) SetHref(v string) {
	o.Href = &v
}

// GetIconLink returns the IconLink field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetIconLink() string {
	if o == nil || o.IconLink == nil {
		var ret string
		return ret
	}
	return *o.IconLink
}

// GetIconLinkOk returns a tuple with the IconLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetIconLinkOk() (*string, bool) {
	if o == nil || o.IconLink == nil {
		return nil, false
	}
	return o.IconLink, true
}

// HasIconLink returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasIconLink() bool {
	if o != nil && o.IconLink != nil {
		return true
	}

	return false
}

// SetIconLink gets a reference to the given string and assigns it to the IconLink field.
func (o *BTCloudStorageObjectInfo) SetIconLink(v string) {
	o.IconLink = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTCloudStorageObjectInfo) SetId(v string) {
	o.Id = &v
}

// GetIsContainer returns the IsContainer field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetIsContainer() bool {
	if o == nil || o.IsContainer == nil {
		var ret bool
		return ret
	}
	return *o.IsContainer
}

// GetIsContainerOk returns a tuple with the IsContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetIsContainerOk() (*bool, bool) {
	if o == nil || o.IsContainer == nil {
		return nil, false
	}
	return o.IsContainer, true
}

// HasIsContainer returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasIsContainer() bool {
	if o != nil && o.IsContainer != nil {
		return true
	}

	return false
}

// SetIsContainer gets a reference to the given bool and assigns it to the IsContainer field.
func (o *BTCloudStorageObjectInfo) SetIsContainer(v bool) {
	o.IsContainer = &v
}

// GetIsEnterpriseOwned returns the IsEnterpriseOwned field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetIsEnterpriseOwned() bool {
	if o == nil || o.IsEnterpriseOwned == nil {
		var ret bool
		return ret
	}
	return *o.IsEnterpriseOwned
}

// GetIsEnterpriseOwnedOk returns a tuple with the IsEnterpriseOwned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetIsEnterpriseOwnedOk() (*bool, bool) {
	if o == nil || o.IsEnterpriseOwned == nil {
		return nil, false
	}
	return o.IsEnterpriseOwned, true
}

// HasIsEnterpriseOwned returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasIsEnterpriseOwned() bool {
	if o != nil && o.IsEnterpriseOwned != nil {
		return true
	}

	return false
}

// SetIsEnterpriseOwned gets a reference to the given bool and assigns it to the IsEnterpriseOwned field.
func (o *BTCloudStorageObjectInfo) SetIsEnterpriseOwned(v bool) {
	o.IsEnterpriseOwned = &v
}

// GetIsMutable returns the IsMutable field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetIsMutable() bool {
	if o == nil || o.IsMutable == nil {
		var ret bool
		return ret
	}
	return *o.IsMutable
}

// GetIsMutableOk returns a tuple with the IsMutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetIsMutableOk() (*bool, bool) {
	if o == nil || o.IsMutable == nil {
		return nil, false
	}
	return o.IsMutable, true
}

// HasIsMutable returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasIsMutable() bool {
	if o != nil && o.IsMutable != nil {
		return true
	}

	return false
}

// SetIsMutable gets a reference to the given bool and assigns it to the IsMutable field.
func (o *BTCloudStorageObjectInfo) SetIsMutable(v bool) {
	o.IsMutable = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetMimeType() string {
	if o == nil || o.MimeType == nil {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetMimeTypeOk() (*string, bool) {
	if o == nil || o.MimeType == nil {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasMimeType() bool {
	if o != nil && o.MimeType != nil {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *BTCloudStorageObjectInfo) SetMimeType(v string) {
	o.MimeType = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetModifiedAt() JSONTime {
	if o == nil || o.ModifiedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetModifiedAtOk() (*JSONTime, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given JSONTime and assigns it to the ModifiedAt field.
func (o *BTCloudStorageObjectInfo) SetModifiedAt(v JSONTime) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetModifiedBy() BTUserBasicSummaryInfo {
	if o == nil || o.ModifiedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the ModifiedBy field.
func (o *BTCloudStorageObjectInfo) SetModifiedBy(v BTUserBasicSummaryInfo) {
	o.ModifiedBy = &v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetModifiedById() string {
	if o == nil || o.ModifiedById == nil {
		var ret string
		return ret
	}
	return *o.ModifiedById
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetModifiedByIdOk() (*string, bool) {
	if o == nil || o.ModifiedById == nil {
		return nil, false
	}
	return o.ModifiedById, true
}

// HasModifiedById returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasModifiedById() bool {
	if o != nil && o.ModifiedById != nil {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given string and assigns it to the ModifiedById field.
func (o *BTCloudStorageObjectInfo) SetModifiedById(v string) {
	o.ModifiedById = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTCloudStorageObjectInfo) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetOwner() BTOwnerInfo {
	if o == nil || o.Owner == nil {
		var ret BTOwnerInfo
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetOwnerOk() (*BTOwnerInfo, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given BTOwnerInfo and assigns it to the Owner field.
func (o *BTCloudStorageObjectInfo) SetOwner(v BTOwnerInfo) {
	o.Owner = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTCloudStorageObjectInfo) SetParentId(v string) {
	o.ParentId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTCloudStorageObjectInfo) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *BTCloudStorageObjectInfo) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetSizeBytes returns the SizeBytes field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetSizeBytes() int64 {
	if o == nil || o.SizeBytes == nil {
		var ret int64
		return ret
	}
	return *o.SizeBytes
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetSizeBytesOk() (*int64, bool) {
	if o == nil || o.SizeBytes == nil {
		return nil, false
	}
	return o.SizeBytes, true
}

// HasSizeBytes returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasSizeBytes() bool {
	if o != nil && o.SizeBytes != nil {
		return true
	}

	return false
}

// SetSizeBytes gets a reference to the given int64 and assigns it to the SizeBytes field.
func (o *BTCloudStorageObjectInfo) SetSizeBytes(v int64) {
	o.SizeBytes = &v
}

// GetThumbnailInfo returns the ThumbnailInfo field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetThumbnailInfo() BTThumbnailInfo {
	if o == nil || o.ThumbnailInfo == nil {
		var ret BTThumbnailInfo
		return ret
	}
	return *o.ThumbnailInfo
}

// GetThumbnailInfoOk returns a tuple with the ThumbnailInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetThumbnailInfoOk() (*BTThumbnailInfo, bool) {
	if o == nil || o.ThumbnailInfo == nil {
		return nil, false
	}
	return o.ThumbnailInfo, true
}

// HasThumbnailInfo returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasThumbnailInfo() bool {
	if o != nil && o.ThumbnailInfo != nil {
		return true
	}

	return false
}

// SetThumbnailInfo gets a reference to the given BTThumbnailInfo and assigns it to the ThumbnailInfo field.
func (o *BTCloudStorageObjectInfo) SetThumbnailInfo(v BTThumbnailInfo) {
	o.ThumbnailInfo = &v
}

// GetTreeHref returns the TreeHref field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetTreeHref() string {
	if o == nil || o.TreeHref == nil {
		var ret string
		return ret
	}
	return *o.TreeHref
}

// GetTreeHrefOk returns a tuple with the TreeHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetTreeHrefOk() (*string, bool) {
	if o == nil || o.TreeHref == nil {
		return nil, false
	}
	return o.TreeHref, true
}

// HasTreeHref returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasTreeHref() bool {
	if o != nil && o.TreeHref != nil {
		return true
	}

	return false
}

// SetTreeHref gets a reference to the given string and assigns it to the TreeHref field.
func (o *BTCloudStorageObjectInfo) SetTreeHref(v string) {
	o.TreeHref = &v
}

// GetUnparentHref returns the UnparentHref field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetUnparentHref() string {
	if o == nil || o.UnparentHref == nil {
		var ret string
		return ret
	}
	return *o.UnparentHref
}

// GetUnparentHrefOk returns a tuple with the UnparentHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetUnparentHrefOk() (*string, bool) {
	if o == nil || o.UnparentHref == nil {
		return nil, false
	}
	return o.UnparentHref, true
}

// HasUnparentHref returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasUnparentHref() bool {
	if o != nil && o.UnparentHref != nil {
		return true
	}

	return false
}

// SetUnparentHref gets a reference to the given string and assigns it to the UnparentHref field.
func (o *BTCloudStorageObjectInfo) SetUnparentHref(v string) {
	o.UnparentHref = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTCloudStorageObjectInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetWebViewLink returns the WebViewLink field value if set, zero value otherwise.
func (o *BTCloudStorageObjectInfo) GetWebViewLink() string {
	if o == nil || o.WebViewLink == nil {
		var ret string
		return ret
	}
	return *o.WebViewLink
}

// GetWebViewLinkOk returns a tuple with the WebViewLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageObjectInfo) GetWebViewLinkOk() (*string, bool) {
	if o == nil || o.WebViewLink == nil {
		return nil, false
	}
	return o.WebViewLink, true
}

// HasWebViewLink returns a boolean if a field has been set.
func (o *BTCloudStorageObjectInfo) HasWebViewLink() bool {
	if o != nil && o.WebViewLink != nil {
		return true
	}

	return false
}

// SetWebViewLink gets a reference to the given string and assigns it to the WebViewLink field.
func (o *BTCloudStorageObjectInfo) SetWebViewLink(v string) {
	o.WebViewLink = &v
}

func (o BTCloudStorageObjectInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanMove != nil {
		toSerialize["canMove"] = o.CanMove
	}
	if o.CloudStorageAccountId != nil {
		toSerialize["cloudStorageAccountId"] = o.CloudStorageAccountId
	}
	if o.CloudStorageObjectId != nil {
		toSerialize["cloudStorageObjectId"] = o.CloudStorageObjectId
	}
	if o.CloudStorageProvider != nil {
		toSerialize["cloudStorageProvider"] = o.CloudStorageProvider
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedById != nil {
		toSerialize["createdById"] = o.CreatedById
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.IconLink != nil {
		toSerialize["iconLink"] = o.IconLink
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
	if o.MimeType != nil {
		toSerialize["mimeType"] = o.MimeType
	}
	if o.ModifiedAt != nil {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if o.ModifiedBy != nil {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if o.ModifiedById != nil {
		toSerialize["modifiedById"] = o.ModifiedById
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.SizeBytes != nil {
		toSerialize["sizeBytes"] = o.SizeBytes
	}
	if o.ThumbnailInfo != nil {
		toSerialize["thumbnailInfo"] = o.ThumbnailInfo
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
	if o.WebViewLink != nil {
		toSerialize["webViewLink"] = o.WebViewLink
	}
	return json.Marshal(toSerialize)
}

type NullableBTCloudStorageObjectInfo struct {
	value *BTCloudStorageObjectInfo
	isSet bool
}

func (v NullableBTCloudStorageObjectInfo) Get() *BTCloudStorageObjectInfo {
	return v.value
}

func (v *NullableBTCloudStorageObjectInfo) Set(val *BTCloudStorageObjectInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCloudStorageObjectInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCloudStorageObjectInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCloudStorageObjectInfo(val *BTCloudStorageObjectInfo) *NullableBTCloudStorageObjectInfo {
	return &NullableBTCloudStorageObjectInfo{value: val, isSet: true}
}

func (v NullableBTCloudStorageObjectInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCloudStorageObjectInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
