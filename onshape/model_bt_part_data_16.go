/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13685-0fb99d06bde5
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPartData16 struct for BTPartData16
type BTPartData16 struct {
	BestAvailableTessellationSetting          *string                          `json:"bestAvailableTessellationSetting,omitempty"`
	BoundsDiameter                            *float64                         `json:"boundsDiameter,omitempty"`
	ClosedConstituentPartData                 *BTClosedConstituentPartData2911 `json:"closedConstituentPartData,omitempty"`
	CoarsePlanarFaceTriangleCount             *int32                           `json:"coarsePlanarFaceTriangleCount,omitempty"`
	CoarseTriangleCount                       *int32                           `json:"coarseTriangleCount,omitempty"`
	ConstituentBodyDeterministicIds           []string                         `json:"constituentBodyDeterministicIds,omitempty"`
	CopyWithoutEntities                       *BTPartData16                    `json:"copyWithoutEntities,omitempty"`
	EntityDIds                                []string                         `json:"entityDIds,omitempty"`
	EntityDeterministicIds                    []string                         `json:"entityDeterministicIds,omitempty"`
	FlattenedToUnflattenedMapping             *map[string]string               `json:"flattenedToUnflattenedMapping,omitempty"`
	HighBoxCorner                             *BTVector3d389                   `json:"highBoxCorner,omitempty"`
	IsACopyForTessellationSettings            *bool                            `json:"isACopyForTessellationSettings,omitempty"`
	IsAssociatedWithFlat                      *bool                            `json:"isAssociatedWithFlat,omitempty"`
	IsBendCenterLineBody                      *bool                            `json:"isBendCenterLineBody,omitempty"`
	IsClosedComposite                         *bool                            `json:"isClosedComposite,omitempty"`
	IsComposite                               *bool                            `json:"isComposite,omitempty"`
	IsDeletion                                *bool                            `json:"isDeletion,omitempty"`
	IsEntitylessPartData                      *bool                            `json:"isEntitylessPartData,omitempty"`
	IsFlattenedSheetMetalBody                 *bool                            `json:"isFlattenedSheetMetalBody,omitempty"`
	IsOpenComposite                           *bool                            `json:"isOpenComposite,omitempty"`
	LowBoxCorner                              *BTVector3d389                   `json:"lowBoxCorner,omitempty"`
	OwnerFlattenedBodyId                      *string                          `json:"ownerFlattenedBodyId,omitempty"`
	SheetMetalModelId                         *string                          `json:"sheetMetalModelId,omitempty"`
	SheetMetalOrderId                         *string                          `json:"sheetMetalOrderId,omitempty"`
	ShouldAlwaysUseHighestQualityTessellation *bool                            `json:"shouldAlwaysUseHighestQualityTessellation,omitempty"`
	TessellationSettings                      []int32                          `json:"tessellationSettings,omitempty"`
	TotalEntityCount                          *int32                           `json:"totalEntityCount,omitempty"`
	UserTessellationSetting                   *string                          `json:"userTessellationSetting,omitempty"`
}

// NewBTPartData16 instantiates a new BTPartData16 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPartData16() *BTPartData16 {
	this := BTPartData16{}
	return &this
}

// NewBTPartData16WithDefaults instantiates a new BTPartData16 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPartData16WithDefaults() *BTPartData16 {
	this := BTPartData16{}
	return &this
}

// GetBestAvailableTessellationSetting returns the BestAvailableTessellationSetting field value if set, zero value otherwise.
func (o *BTPartData16) GetBestAvailableTessellationSetting() string {
	if o == nil || o.BestAvailableTessellationSetting == nil {
		var ret string
		return ret
	}
	return *o.BestAvailableTessellationSetting
}

// GetBestAvailableTessellationSettingOk returns a tuple with the BestAvailableTessellationSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetBestAvailableTessellationSettingOk() (*string, bool) {
	if o == nil || o.BestAvailableTessellationSetting == nil {
		return nil, false
	}
	return o.BestAvailableTessellationSetting, true
}

// HasBestAvailableTessellationSetting returns a boolean if a field has been set.
func (o *BTPartData16) HasBestAvailableTessellationSetting() bool {
	if o != nil && o.BestAvailableTessellationSetting != nil {
		return true
	}

	return false
}

// SetBestAvailableTessellationSetting gets a reference to the given string and assigns it to the BestAvailableTessellationSetting field.
func (o *BTPartData16) SetBestAvailableTessellationSetting(v string) {
	o.BestAvailableTessellationSetting = &v
}

// GetBoundsDiameter returns the BoundsDiameter field value if set, zero value otherwise.
func (o *BTPartData16) GetBoundsDiameter() float64 {
	if o == nil || o.BoundsDiameter == nil {
		var ret float64
		return ret
	}
	return *o.BoundsDiameter
}

// GetBoundsDiameterOk returns a tuple with the BoundsDiameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetBoundsDiameterOk() (*float64, bool) {
	if o == nil || o.BoundsDiameter == nil {
		return nil, false
	}
	return o.BoundsDiameter, true
}

// HasBoundsDiameter returns a boolean if a field has been set.
func (o *BTPartData16) HasBoundsDiameter() bool {
	if o != nil && o.BoundsDiameter != nil {
		return true
	}

	return false
}

// SetBoundsDiameter gets a reference to the given float64 and assigns it to the BoundsDiameter field.
func (o *BTPartData16) SetBoundsDiameter(v float64) {
	o.BoundsDiameter = &v
}

// GetClosedConstituentPartData returns the ClosedConstituentPartData field value if set, zero value otherwise.
func (o *BTPartData16) GetClosedConstituentPartData() BTClosedConstituentPartData2911 {
	if o == nil || o.ClosedConstituentPartData == nil {
		var ret BTClosedConstituentPartData2911
		return ret
	}
	return *o.ClosedConstituentPartData
}

// GetClosedConstituentPartDataOk returns a tuple with the ClosedConstituentPartData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetClosedConstituentPartDataOk() (*BTClosedConstituentPartData2911, bool) {
	if o == nil || o.ClosedConstituentPartData == nil {
		return nil, false
	}
	return o.ClosedConstituentPartData, true
}

// HasClosedConstituentPartData returns a boolean if a field has been set.
func (o *BTPartData16) HasClosedConstituentPartData() bool {
	if o != nil && o.ClosedConstituentPartData != nil {
		return true
	}

	return false
}

// SetClosedConstituentPartData gets a reference to the given BTClosedConstituentPartData2911 and assigns it to the ClosedConstituentPartData field.
func (o *BTPartData16) SetClosedConstituentPartData(v BTClosedConstituentPartData2911) {
	o.ClosedConstituentPartData = &v
}

// GetCoarsePlanarFaceTriangleCount returns the CoarsePlanarFaceTriangleCount field value if set, zero value otherwise.
func (o *BTPartData16) GetCoarsePlanarFaceTriangleCount() int32 {
	if o == nil || o.CoarsePlanarFaceTriangleCount == nil {
		var ret int32
		return ret
	}
	return *o.CoarsePlanarFaceTriangleCount
}

// GetCoarsePlanarFaceTriangleCountOk returns a tuple with the CoarsePlanarFaceTriangleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetCoarsePlanarFaceTriangleCountOk() (*int32, bool) {
	if o == nil || o.CoarsePlanarFaceTriangleCount == nil {
		return nil, false
	}
	return o.CoarsePlanarFaceTriangleCount, true
}

// HasCoarsePlanarFaceTriangleCount returns a boolean if a field has been set.
func (o *BTPartData16) HasCoarsePlanarFaceTriangleCount() bool {
	if o != nil && o.CoarsePlanarFaceTriangleCount != nil {
		return true
	}

	return false
}

// SetCoarsePlanarFaceTriangleCount gets a reference to the given int32 and assigns it to the CoarsePlanarFaceTriangleCount field.
func (o *BTPartData16) SetCoarsePlanarFaceTriangleCount(v int32) {
	o.CoarsePlanarFaceTriangleCount = &v
}

// GetCoarseTriangleCount returns the CoarseTriangleCount field value if set, zero value otherwise.
func (o *BTPartData16) GetCoarseTriangleCount() int32 {
	if o == nil || o.CoarseTriangleCount == nil {
		var ret int32
		return ret
	}
	return *o.CoarseTriangleCount
}

// GetCoarseTriangleCountOk returns a tuple with the CoarseTriangleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetCoarseTriangleCountOk() (*int32, bool) {
	if o == nil || o.CoarseTriangleCount == nil {
		return nil, false
	}
	return o.CoarseTriangleCount, true
}

// HasCoarseTriangleCount returns a boolean if a field has been set.
func (o *BTPartData16) HasCoarseTriangleCount() bool {
	if o != nil && o.CoarseTriangleCount != nil {
		return true
	}

	return false
}

// SetCoarseTriangleCount gets a reference to the given int32 and assigns it to the CoarseTriangleCount field.
func (o *BTPartData16) SetCoarseTriangleCount(v int32) {
	o.CoarseTriangleCount = &v
}

// GetConstituentBodyDeterministicIds returns the ConstituentBodyDeterministicIds field value if set, zero value otherwise.
func (o *BTPartData16) GetConstituentBodyDeterministicIds() []string {
	if o == nil || o.ConstituentBodyDeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.ConstituentBodyDeterministicIds
}

// GetConstituentBodyDeterministicIdsOk returns a tuple with the ConstituentBodyDeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetConstituentBodyDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.ConstituentBodyDeterministicIds == nil {
		return nil, false
	}
	return o.ConstituentBodyDeterministicIds, true
}

// HasConstituentBodyDeterministicIds returns a boolean if a field has been set.
func (o *BTPartData16) HasConstituentBodyDeterministicIds() bool {
	if o != nil && o.ConstituentBodyDeterministicIds != nil {
		return true
	}

	return false
}

// SetConstituentBodyDeterministicIds gets a reference to the given []string and assigns it to the ConstituentBodyDeterministicIds field.
func (o *BTPartData16) SetConstituentBodyDeterministicIds(v []string) {
	o.ConstituentBodyDeterministicIds = v
}

// GetCopyWithoutEntities returns the CopyWithoutEntities field value if set, zero value otherwise.
func (o *BTPartData16) GetCopyWithoutEntities() BTPartData16 {
	if o == nil || o.CopyWithoutEntities == nil {
		var ret BTPartData16
		return ret
	}
	return *o.CopyWithoutEntities
}

// GetCopyWithoutEntitiesOk returns a tuple with the CopyWithoutEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetCopyWithoutEntitiesOk() (*BTPartData16, bool) {
	if o == nil || o.CopyWithoutEntities == nil {
		return nil, false
	}
	return o.CopyWithoutEntities, true
}

// HasCopyWithoutEntities returns a boolean if a field has been set.
func (o *BTPartData16) HasCopyWithoutEntities() bool {
	if o != nil && o.CopyWithoutEntities != nil {
		return true
	}

	return false
}

// SetCopyWithoutEntities gets a reference to the given BTPartData16 and assigns it to the CopyWithoutEntities field.
func (o *BTPartData16) SetCopyWithoutEntities(v BTPartData16) {
	o.CopyWithoutEntities = &v
}

// GetEntityDIds returns the EntityDIds field value if set, zero value otherwise.
func (o *BTPartData16) GetEntityDIds() []string {
	if o == nil || o.EntityDIds == nil {
		var ret []string
		return ret
	}
	return o.EntityDIds
}

// GetEntityDIdsOk returns a tuple with the EntityDIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetEntityDIdsOk() ([]string, bool) {
	if o == nil || o.EntityDIds == nil {
		return nil, false
	}
	return o.EntityDIds, true
}

// HasEntityDIds returns a boolean if a field has been set.
func (o *BTPartData16) HasEntityDIds() bool {
	if o != nil && o.EntityDIds != nil {
		return true
	}

	return false
}

// SetEntityDIds gets a reference to the given []string and assigns it to the EntityDIds field.
func (o *BTPartData16) SetEntityDIds(v []string) {
	o.EntityDIds = v
}

// GetEntityDeterministicIds returns the EntityDeterministicIds field value if set, zero value otherwise.
func (o *BTPartData16) GetEntityDeterministicIds() []string {
	if o == nil || o.EntityDeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.EntityDeterministicIds
}

// GetEntityDeterministicIdsOk returns a tuple with the EntityDeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetEntityDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.EntityDeterministicIds == nil {
		return nil, false
	}
	return o.EntityDeterministicIds, true
}

// HasEntityDeterministicIds returns a boolean if a field has been set.
func (o *BTPartData16) HasEntityDeterministicIds() bool {
	if o != nil && o.EntityDeterministicIds != nil {
		return true
	}

	return false
}

// SetEntityDeterministicIds gets a reference to the given []string and assigns it to the EntityDeterministicIds field.
func (o *BTPartData16) SetEntityDeterministicIds(v []string) {
	o.EntityDeterministicIds = v
}

// GetFlattenedToUnflattenedMapping returns the FlattenedToUnflattenedMapping field value if set, zero value otherwise.
func (o *BTPartData16) GetFlattenedToUnflattenedMapping() map[string]string {
	if o == nil || o.FlattenedToUnflattenedMapping == nil {
		var ret map[string]string
		return ret
	}
	return *o.FlattenedToUnflattenedMapping
}

// GetFlattenedToUnflattenedMappingOk returns a tuple with the FlattenedToUnflattenedMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetFlattenedToUnflattenedMappingOk() (*map[string]string, bool) {
	if o == nil || o.FlattenedToUnflattenedMapping == nil {
		return nil, false
	}
	return o.FlattenedToUnflattenedMapping, true
}

// HasFlattenedToUnflattenedMapping returns a boolean if a field has been set.
func (o *BTPartData16) HasFlattenedToUnflattenedMapping() bool {
	if o != nil && o.FlattenedToUnflattenedMapping != nil {
		return true
	}

	return false
}

// SetFlattenedToUnflattenedMapping gets a reference to the given map[string]string and assigns it to the FlattenedToUnflattenedMapping field.
func (o *BTPartData16) SetFlattenedToUnflattenedMapping(v map[string]string) {
	o.FlattenedToUnflattenedMapping = &v
}

// GetHighBoxCorner returns the HighBoxCorner field value if set, zero value otherwise.
func (o *BTPartData16) GetHighBoxCorner() BTVector3d389 {
	if o == nil || o.HighBoxCorner == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.HighBoxCorner
}

// GetHighBoxCornerOk returns a tuple with the HighBoxCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetHighBoxCornerOk() (*BTVector3d389, bool) {
	if o == nil || o.HighBoxCorner == nil {
		return nil, false
	}
	return o.HighBoxCorner, true
}

// HasHighBoxCorner returns a boolean if a field has been set.
func (o *BTPartData16) HasHighBoxCorner() bool {
	if o != nil && o.HighBoxCorner != nil {
		return true
	}

	return false
}

// SetHighBoxCorner gets a reference to the given BTVector3d389 and assigns it to the HighBoxCorner field.
func (o *BTPartData16) SetHighBoxCorner(v BTVector3d389) {
	o.HighBoxCorner = &v
}

// GetIsACopyForTessellationSettings returns the IsACopyForTessellationSettings field value if set, zero value otherwise.
func (o *BTPartData16) GetIsACopyForTessellationSettings() bool {
	if o == nil || o.IsACopyForTessellationSettings == nil {
		var ret bool
		return ret
	}
	return *o.IsACopyForTessellationSettings
}

// GetIsACopyForTessellationSettingsOk returns a tuple with the IsACopyForTessellationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetIsACopyForTessellationSettingsOk() (*bool, bool) {
	if o == nil || o.IsACopyForTessellationSettings == nil {
		return nil, false
	}
	return o.IsACopyForTessellationSettings, true
}

// HasIsACopyForTessellationSettings returns a boolean if a field has been set.
func (o *BTPartData16) HasIsACopyForTessellationSettings() bool {
	if o != nil && o.IsACopyForTessellationSettings != nil {
		return true
	}

	return false
}

// SetIsACopyForTessellationSettings gets a reference to the given bool and assigns it to the IsACopyForTessellationSettings field.
func (o *BTPartData16) SetIsACopyForTessellationSettings(v bool) {
	o.IsACopyForTessellationSettings = &v
}

// GetIsAssociatedWithFlat returns the IsAssociatedWithFlat field value if set, zero value otherwise.
func (o *BTPartData16) GetIsAssociatedWithFlat() bool {
	if o == nil || o.IsAssociatedWithFlat == nil {
		var ret bool
		return ret
	}
	return *o.IsAssociatedWithFlat
}

// GetIsAssociatedWithFlatOk returns a tuple with the IsAssociatedWithFlat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetIsAssociatedWithFlatOk() (*bool, bool) {
	if o == nil || o.IsAssociatedWithFlat == nil {
		return nil, false
	}
	return o.IsAssociatedWithFlat, true
}

// HasIsAssociatedWithFlat returns a boolean if a field has been set.
func (o *BTPartData16) HasIsAssociatedWithFlat() bool {
	if o != nil && o.IsAssociatedWithFlat != nil {
		return true
	}

	return false
}

// SetIsAssociatedWithFlat gets a reference to the given bool and assigns it to the IsAssociatedWithFlat field.
func (o *BTPartData16) SetIsAssociatedWithFlat(v bool) {
	o.IsAssociatedWithFlat = &v
}

// GetIsBendCenterLineBody returns the IsBendCenterLineBody field value if set, zero value otherwise.
func (o *BTPartData16) GetIsBendCenterLineBody() bool {
	if o == nil || o.IsBendCenterLineBody == nil {
		var ret bool
		return ret
	}
	return *o.IsBendCenterLineBody
}

// GetIsBendCenterLineBodyOk returns a tuple with the IsBendCenterLineBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetIsBendCenterLineBodyOk() (*bool, bool) {
	if o == nil || o.IsBendCenterLineBody == nil {
		return nil, false
	}
	return o.IsBendCenterLineBody, true
}

// HasIsBendCenterLineBody returns a boolean if a field has been set.
func (o *BTPartData16) HasIsBendCenterLineBody() bool {
	if o != nil && o.IsBendCenterLineBody != nil {
		return true
	}

	return false
}

// SetIsBendCenterLineBody gets a reference to the given bool and assigns it to the IsBendCenterLineBody field.
func (o *BTPartData16) SetIsBendCenterLineBody(v bool) {
	o.IsBendCenterLineBody = &v
}

// GetIsClosedComposite returns the IsClosedComposite field value if set, zero value otherwise.
func (o *BTPartData16) GetIsClosedComposite() bool {
	if o == nil || o.IsClosedComposite == nil {
		var ret bool
		return ret
	}
	return *o.IsClosedComposite
}

// GetIsClosedCompositeOk returns a tuple with the IsClosedComposite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetIsClosedCompositeOk() (*bool, bool) {
	if o == nil || o.IsClosedComposite == nil {
		return nil, false
	}
	return o.IsClosedComposite, true
}

// HasIsClosedComposite returns a boolean if a field has been set.
func (o *BTPartData16) HasIsClosedComposite() bool {
	if o != nil && o.IsClosedComposite != nil {
		return true
	}

	return false
}

// SetIsClosedComposite gets a reference to the given bool and assigns it to the IsClosedComposite field.
func (o *BTPartData16) SetIsClosedComposite(v bool) {
	o.IsClosedComposite = &v
}

// GetIsComposite returns the IsComposite field value if set, zero value otherwise.
func (o *BTPartData16) GetIsComposite() bool {
	if o == nil || o.IsComposite == nil {
		var ret bool
		return ret
	}
	return *o.IsComposite
}

// GetIsCompositeOk returns a tuple with the IsComposite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetIsCompositeOk() (*bool, bool) {
	if o == nil || o.IsComposite == nil {
		return nil, false
	}
	return o.IsComposite, true
}

// HasIsComposite returns a boolean if a field has been set.
func (o *BTPartData16) HasIsComposite() bool {
	if o != nil && o.IsComposite != nil {
		return true
	}

	return false
}

// SetIsComposite gets a reference to the given bool and assigns it to the IsComposite field.
func (o *BTPartData16) SetIsComposite(v bool) {
	o.IsComposite = &v
}

// GetIsDeletion returns the IsDeletion field value if set, zero value otherwise.
func (o *BTPartData16) GetIsDeletion() bool {
	if o == nil || o.IsDeletion == nil {
		var ret bool
		return ret
	}
	return *o.IsDeletion
}

// GetIsDeletionOk returns a tuple with the IsDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetIsDeletionOk() (*bool, bool) {
	if o == nil || o.IsDeletion == nil {
		return nil, false
	}
	return o.IsDeletion, true
}

// HasIsDeletion returns a boolean if a field has been set.
func (o *BTPartData16) HasIsDeletion() bool {
	if o != nil && o.IsDeletion != nil {
		return true
	}

	return false
}

// SetIsDeletion gets a reference to the given bool and assigns it to the IsDeletion field.
func (o *BTPartData16) SetIsDeletion(v bool) {
	o.IsDeletion = &v
}

// GetIsEntitylessPartData returns the IsEntitylessPartData field value if set, zero value otherwise.
func (o *BTPartData16) GetIsEntitylessPartData() bool {
	if o == nil || o.IsEntitylessPartData == nil {
		var ret bool
		return ret
	}
	return *o.IsEntitylessPartData
}

// GetIsEntitylessPartDataOk returns a tuple with the IsEntitylessPartData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetIsEntitylessPartDataOk() (*bool, bool) {
	if o == nil || o.IsEntitylessPartData == nil {
		return nil, false
	}
	return o.IsEntitylessPartData, true
}

// HasIsEntitylessPartData returns a boolean if a field has been set.
func (o *BTPartData16) HasIsEntitylessPartData() bool {
	if o != nil && o.IsEntitylessPartData != nil {
		return true
	}

	return false
}

// SetIsEntitylessPartData gets a reference to the given bool and assigns it to the IsEntitylessPartData field.
func (o *BTPartData16) SetIsEntitylessPartData(v bool) {
	o.IsEntitylessPartData = &v
}

// GetIsFlattenedSheetMetalBody returns the IsFlattenedSheetMetalBody field value if set, zero value otherwise.
func (o *BTPartData16) GetIsFlattenedSheetMetalBody() bool {
	if o == nil || o.IsFlattenedSheetMetalBody == nil {
		var ret bool
		return ret
	}
	return *o.IsFlattenedSheetMetalBody
}

// GetIsFlattenedSheetMetalBodyOk returns a tuple with the IsFlattenedSheetMetalBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetIsFlattenedSheetMetalBodyOk() (*bool, bool) {
	if o == nil || o.IsFlattenedSheetMetalBody == nil {
		return nil, false
	}
	return o.IsFlattenedSheetMetalBody, true
}

// HasIsFlattenedSheetMetalBody returns a boolean if a field has been set.
func (o *BTPartData16) HasIsFlattenedSheetMetalBody() bool {
	if o != nil && o.IsFlattenedSheetMetalBody != nil {
		return true
	}

	return false
}

// SetIsFlattenedSheetMetalBody gets a reference to the given bool and assigns it to the IsFlattenedSheetMetalBody field.
func (o *BTPartData16) SetIsFlattenedSheetMetalBody(v bool) {
	o.IsFlattenedSheetMetalBody = &v
}

// GetIsOpenComposite returns the IsOpenComposite field value if set, zero value otherwise.
func (o *BTPartData16) GetIsOpenComposite() bool {
	if o == nil || o.IsOpenComposite == nil {
		var ret bool
		return ret
	}
	return *o.IsOpenComposite
}

// GetIsOpenCompositeOk returns a tuple with the IsOpenComposite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetIsOpenCompositeOk() (*bool, bool) {
	if o == nil || o.IsOpenComposite == nil {
		return nil, false
	}
	return o.IsOpenComposite, true
}

// HasIsOpenComposite returns a boolean if a field has been set.
func (o *BTPartData16) HasIsOpenComposite() bool {
	if o != nil && o.IsOpenComposite != nil {
		return true
	}

	return false
}

// SetIsOpenComposite gets a reference to the given bool and assigns it to the IsOpenComposite field.
func (o *BTPartData16) SetIsOpenComposite(v bool) {
	o.IsOpenComposite = &v
}

// GetLowBoxCorner returns the LowBoxCorner field value if set, zero value otherwise.
func (o *BTPartData16) GetLowBoxCorner() BTVector3d389 {
	if o == nil || o.LowBoxCorner == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.LowBoxCorner
}

// GetLowBoxCornerOk returns a tuple with the LowBoxCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetLowBoxCornerOk() (*BTVector3d389, bool) {
	if o == nil || o.LowBoxCorner == nil {
		return nil, false
	}
	return o.LowBoxCorner, true
}

// HasLowBoxCorner returns a boolean if a field has been set.
func (o *BTPartData16) HasLowBoxCorner() bool {
	if o != nil && o.LowBoxCorner != nil {
		return true
	}

	return false
}

// SetLowBoxCorner gets a reference to the given BTVector3d389 and assigns it to the LowBoxCorner field.
func (o *BTPartData16) SetLowBoxCorner(v BTVector3d389) {
	o.LowBoxCorner = &v
}

// GetOwnerFlattenedBodyId returns the OwnerFlattenedBodyId field value if set, zero value otherwise.
func (o *BTPartData16) GetOwnerFlattenedBodyId() string {
	if o == nil || o.OwnerFlattenedBodyId == nil {
		var ret string
		return ret
	}
	return *o.OwnerFlattenedBodyId
}

// GetOwnerFlattenedBodyIdOk returns a tuple with the OwnerFlattenedBodyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetOwnerFlattenedBodyIdOk() (*string, bool) {
	if o == nil || o.OwnerFlattenedBodyId == nil {
		return nil, false
	}
	return o.OwnerFlattenedBodyId, true
}

// HasOwnerFlattenedBodyId returns a boolean if a field has been set.
func (o *BTPartData16) HasOwnerFlattenedBodyId() bool {
	if o != nil && o.OwnerFlattenedBodyId != nil {
		return true
	}

	return false
}

// SetOwnerFlattenedBodyId gets a reference to the given string and assigns it to the OwnerFlattenedBodyId field.
func (o *BTPartData16) SetOwnerFlattenedBodyId(v string) {
	o.OwnerFlattenedBodyId = &v
}

// GetSheetMetalModelId returns the SheetMetalModelId field value if set, zero value otherwise.
func (o *BTPartData16) GetSheetMetalModelId() string {
	if o == nil || o.SheetMetalModelId == nil {
		var ret string
		return ret
	}
	return *o.SheetMetalModelId
}

// GetSheetMetalModelIdOk returns a tuple with the SheetMetalModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetSheetMetalModelIdOk() (*string, bool) {
	if o == nil || o.SheetMetalModelId == nil {
		return nil, false
	}
	return o.SheetMetalModelId, true
}

// HasSheetMetalModelId returns a boolean if a field has been set.
func (o *BTPartData16) HasSheetMetalModelId() bool {
	if o != nil && o.SheetMetalModelId != nil {
		return true
	}

	return false
}

// SetSheetMetalModelId gets a reference to the given string and assigns it to the SheetMetalModelId field.
func (o *BTPartData16) SetSheetMetalModelId(v string) {
	o.SheetMetalModelId = &v
}

// GetSheetMetalOrderId returns the SheetMetalOrderId field value if set, zero value otherwise.
func (o *BTPartData16) GetSheetMetalOrderId() string {
	if o == nil || o.SheetMetalOrderId == nil {
		var ret string
		return ret
	}
	return *o.SheetMetalOrderId
}

// GetSheetMetalOrderIdOk returns a tuple with the SheetMetalOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetSheetMetalOrderIdOk() (*string, bool) {
	if o == nil || o.SheetMetalOrderId == nil {
		return nil, false
	}
	return o.SheetMetalOrderId, true
}

// HasSheetMetalOrderId returns a boolean if a field has been set.
func (o *BTPartData16) HasSheetMetalOrderId() bool {
	if o != nil && o.SheetMetalOrderId != nil {
		return true
	}

	return false
}

// SetSheetMetalOrderId gets a reference to the given string and assigns it to the SheetMetalOrderId field.
func (o *BTPartData16) SetSheetMetalOrderId(v string) {
	o.SheetMetalOrderId = &v
}

// GetShouldAlwaysUseHighestQualityTessellation returns the ShouldAlwaysUseHighestQualityTessellation field value if set, zero value otherwise.
func (o *BTPartData16) GetShouldAlwaysUseHighestQualityTessellation() bool {
	if o == nil || o.ShouldAlwaysUseHighestQualityTessellation == nil {
		var ret bool
		return ret
	}
	return *o.ShouldAlwaysUseHighestQualityTessellation
}

// GetShouldAlwaysUseHighestQualityTessellationOk returns a tuple with the ShouldAlwaysUseHighestQualityTessellation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetShouldAlwaysUseHighestQualityTessellationOk() (*bool, bool) {
	if o == nil || o.ShouldAlwaysUseHighestQualityTessellation == nil {
		return nil, false
	}
	return o.ShouldAlwaysUseHighestQualityTessellation, true
}

// HasShouldAlwaysUseHighestQualityTessellation returns a boolean if a field has been set.
func (o *BTPartData16) HasShouldAlwaysUseHighestQualityTessellation() bool {
	if o != nil && o.ShouldAlwaysUseHighestQualityTessellation != nil {
		return true
	}

	return false
}

// SetShouldAlwaysUseHighestQualityTessellation gets a reference to the given bool and assigns it to the ShouldAlwaysUseHighestQualityTessellation field.
func (o *BTPartData16) SetShouldAlwaysUseHighestQualityTessellation(v bool) {
	o.ShouldAlwaysUseHighestQualityTessellation = &v
}

// GetTessellationSettings returns the TessellationSettings field value if set, zero value otherwise.
func (o *BTPartData16) GetTessellationSettings() []int32 {
	if o == nil || o.TessellationSettings == nil {
		var ret []int32
		return ret
	}
	return o.TessellationSettings
}

// GetTessellationSettingsOk returns a tuple with the TessellationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetTessellationSettingsOk() ([]int32, bool) {
	if o == nil || o.TessellationSettings == nil {
		return nil, false
	}
	return o.TessellationSettings, true
}

// HasTessellationSettings returns a boolean if a field has been set.
func (o *BTPartData16) HasTessellationSettings() bool {
	if o != nil && o.TessellationSettings != nil {
		return true
	}

	return false
}

// SetTessellationSettings gets a reference to the given []int32 and assigns it to the TessellationSettings field.
func (o *BTPartData16) SetTessellationSettings(v []int32) {
	o.TessellationSettings = v
}

// GetTotalEntityCount returns the TotalEntityCount field value if set, zero value otherwise.
func (o *BTPartData16) GetTotalEntityCount() int32 {
	if o == nil || o.TotalEntityCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalEntityCount
}

// GetTotalEntityCountOk returns a tuple with the TotalEntityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetTotalEntityCountOk() (*int32, bool) {
	if o == nil || o.TotalEntityCount == nil {
		return nil, false
	}
	return o.TotalEntityCount, true
}

// HasTotalEntityCount returns a boolean if a field has been set.
func (o *BTPartData16) HasTotalEntityCount() bool {
	if o != nil && o.TotalEntityCount != nil {
		return true
	}

	return false
}

// SetTotalEntityCount gets a reference to the given int32 and assigns it to the TotalEntityCount field.
func (o *BTPartData16) SetTotalEntityCount(v int32) {
	o.TotalEntityCount = &v
}

// GetUserTessellationSetting returns the UserTessellationSetting field value if set, zero value otherwise.
func (o *BTPartData16) GetUserTessellationSetting() string {
	if o == nil || o.UserTessellationSetting == nil {
		var ret string
		return ret
	}
	return *o.UserTessellationSetting
}

// GetUserTessellationSettingOk returns a tuple with the UserTessellationSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartData16) GetUserTessellationSettingOk() (*string, bool) {
	if o == nil || o.UserTessellationSetting == nil {
		return nil, false
	}
	return o.UserTessellationSetting, true
}

// HasUserTessellationSetting returns a boolean if a field has been set.
func (o *BTPartData16) HasUserTessellationSetting() bool {
	if o != nil && o.UserTessellationSetting != nil {
		return true
	}

	return false
}

// SetUserTessellationSetting gets a reference to the given string and assigns it to the UserTessellationSetting field.
func (o *BTPartData16) SetUserTessellationSetting(v string) {
	o.UserTessellationSetting = &v
}

func (o BTPartData16) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BestAvailableTessellationSetting != nil {
		toSerialize["bestAvailableTessellationSetting"] = o.BestAvailableTessellationSetting
	}
	if o.BoundsDiameter != nil {
		toSerialize["boundsDiameter"] = o.BoundsDiameter
	}
	if o.ClosedConstituentPartData != nil {
		toSerialize["closedConstituentPartData"] = o.ClosedConstituentPartData
	}
	if o.CoarsePlanarFaceTriangleCount != nil {
		toSerialize["coarsePlanarFaceTriangleCount"] = o.CoarsePlanarFaceTriangleCount
	}
	if o.CoarseTriangleCount != nil {
		toSerialize["coarseTriangleCount"] = o.CoarseTriangleCount
	}
	if o.ConstituentBodyDeterministicIds != nil {
		toSerialize["constituentBodyDeterministicIds"] = o.ConstituentBodyDeterministicIds
	}
	if o.CopyWithoutEntities != nil {
		toSerialize["copyWithoutEntities"] = o.CopyWithoutEntities
	}
	if o.EntityDIds != nil {
		toSerialize["entityDIds"] = o.EntityDIds
	}
	if o.EntityDeterministicIds != nil {
		toSerialize["entityDeterministicIds"] = o.EntityDeterministicIds
	}
	if o.FlattenedToUnflattenedMapping != nil {
		toSerialize["flattenedToUnflattenedMapping"] = o.FlattenedToUnflattenedMapping
	}
	if o.HighBoxCorner != nil {
		toSerialize["highBoxCorner"] = o.HighBoxCorner
	}
	if o.IsACopyForTessellationSettings != nil {
		toSerialize["isACopyForTessellationSettings"] = o.IsACopyForTessellationSettings
	}
	if o.IsAssociatedWithFlat != nil {
		toSerialize["isAssociatedWithFlat"] = o.IsAssociatedWithFlat
	}
	if o.IsBendCenterLineBody != nil {
		toSerialize["isBendCenterLineBody"] = o.IsBendCenterLineBody
	}
	if o.IsClosedComposite != nil {
		toSerialize["isClosedComposite"] = o.IsClosedComposite
	}
	if o.IsComposite != nil {
		toSerialize["isComposite"] = o.IsComposite
	}
	if o.IsDeletion != nil {
		toSerialize["isDeletion"] = o.IsDeletion
	}
	if o.IsEntitylessPartData != nil {
		toSerialize["isEntitylessPartData"] = o.IsEntitylessPartData
	}
	if o.IsFlattenedSheetMetalBody != nil {
		toSerialize["isFlattenedSheetMetalBody"] = o.IsFlattenedSheetMetalBody
	}
	if o.IsOpenComposite != nil {
		toSerialize["isOpenComposite"] = o.IsOpenComposite
	}
	if o.LowBoxCorner != nil {
		toSerialize["lowBoxCorner"] = o.LowBoxCorner
	}
	if o.OwnerFlattenedBodyId != nil {
		toSerialize["ownerFlattenedBodyId"] = o.OwnerFlattenedBodyId
	}
	if o.SheetMetalModelId != nil {
		toSerialize["sheetMetalModelId"] = o.SheetMetalModelId
	}
	if o.SheetMetalOrderId != nil {
		toSerialize["sheetMetalOrderId"] = o.SheetMetalOrderId
	}
	if o.ShouldAlwaysUseHighestQualityTessellation != nil {
		toSerialize["shouldAlwaysUseHighestQualityTessellation"] = o.ShouldAlwaysUseHighestQualityTessellation
	}
	if o.TessellationSettings != nil {
		toSerialize["tessellationSettings"] = o.TessellationSettings
	}
	if o.TotalEntityCount != nil {
		toSerialize["totalEntityCount"] = o.TotalEntityCount
	}
	if o.UserTessellationSetting != nil {
		toSerialize["userTessellationSetting"] = o.UserTessellationSetting
	}
	return json.Marshal(toSerialize)
}

type NullableBTPartData16 struct {
	value *BTPartData16
	isSet bool
}

func (v NullableBTPartData16) Get() *BTPartData16 {
	return v.value
}

func (v *NullableBTPartData16) Set(val *BTPartData16) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPartData16) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPartData16) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPartData16(val *BTPartData16) *NullableBTPartData16 {
	return &NullableBTPartData16{value: val, isSet: true}
}

func (v NullableBTPartData16) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPartData16) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
