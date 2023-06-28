/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.18120-f464f720d215
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSketchEntity25 struct for BTSketchEntity25
type BTSketchEntity25 struct {
	BtType                     *string                       `json:"btType,omitempty"`
	CopyWithoutGeometry        *BTBaseEntityData33           `json:"copyWithoutGeometry,omitempty"`
	Decompressed               *BTBaseEntityData33           `json:"decompressed,omitempty"`
	Deletion                   *bool                         `json:"deletion,omitempty"`
	FeatureIds                 []string                      `json:"featureIds,omitempty"`
	FromSketch                 *bool                         `json:"fromSketch,omitempty"`
	Geometries                 []BTEntityGeometry35          `json:"geometries,omitempty"`
	DomainSpecificMetadata     []BTDomainSpecificMetadata961 `json:"domainSpecificMetadata,omitempty"`
	FirstGeometry              *BTEntityGeometry35           `json:"firstGeometry,omitempty"`
	IsConstruction             *bool                         `json:"isConstruction,omitempty"`
	IsFromSplineControlPolygon *bool                         `json:"isFromSplineControlPolygon,omitempty"`
	IsFromSplineHandle         *bool                         `json:"isFromSplineHandle,omitempty"`
	IsTextStroke               *bool                         `json:"isTextStroke,omitempty"`
	IsUserPoint                *bool                         `json:"isUserPoint,omitempty"`
	SketchCurveType            *GBTSketchCurveType           `json:"sketchCurveType,omitempty"`
	SketchEntityId             *string                       `json:"sketchEntityId,omitempty"`
	SolveStatus                *int32                        `json:"solveStatus,omitempty"`
}

// NewBTSketchEntity25 instantiates a new BTSketchEntity25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSketchEntity25() *BTSketchEntity25 {
	this := BTSketchEntity25{}
	return &this
}

// NewBTSketchEntity25WithDefaults instantiates a new BTSketchEntity25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSketchEntity25WithDefaults() *BTSketchEntity25 {
	this := BTSketchEntity25{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSketchEntity25) SetBtType(v string) {
	o.BtType = &v
}

// GetCopyWithoutGeometry returns the CopyWithoutGeometry field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetCopyWithoutGeometry() BTBaseEntityData33 {
	if o == nil || o.CopyWithoutGeometry == nil {
		var ret BTBaseEntityData33
		return ret
	}
	return *o.CopyWithoutGeometry
}

// GetCopyWithoutGeometryOk returns a tuple with the CopyWithoutGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetCopyWithoutGeometryOk() (*BTBaseEntityData33, bool) {
	if o == nil || o.CopyWithoutGeometry == nil {
		return nil, false
	}
	return o.CopyWithoutGeometry, true
}

// HasCopyWithoutGeometry returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasCopyWithoutGeometry() bool {
	if o != nil && o.CopyWithoutGeometry != nil {
		return true
	}

	return false
}

// SetCopyWithoutGeometry gets a reference to the given BTBaseEntityData33 and assigns it to the CopyWithoutGeometry field.
func (o *BTSketchEntity25) SetCopyWithoutGeometry(v BTBaseEntityData33) {
	o.CopyWithoutGeometry = &v
}

// GetDecompressed returns the Decompressed field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetDecompressed() BTBaseEntityData33 {
	if o == nil || o.Decompressed == nil {
		var ret BTBaseEntityData33
		return ret
	}
	return *o.Decompressed
}

// GetDecompressedOk returns a tuple with the Decompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetDecompressedOk() (*BTBaseEntityData33, bool) {
	if o == nil || o.Decompressed == nil {
		return nil, false
	}
	return o.Decompressed, true
}

// HasDecompressed returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasDecompressed() bool {
	if o != nil && o.Decompressed != nil {
		return true
	}

	return false
}

// SetDecompressed gets a reference to the given BTBaseEntityData33 and assigns it to the Decompressed field.
func (o *BTSketchEntity25) SetDecompressed(v BTBaseEntityData33) {
	o.Decompressed = &v
}

// GetDeletion returns the Deletion field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetDeletion() bool {
	if o == nil || o.Deletion == nil {
		var ret bool
		return ret
	}
	return *o.Deletion
}

// GetDeletionOk returns a tuple with the Deletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetDeletionOk() (*bool, bool) {
	if o == nil || o.Deletion == nil {
		return nil, false
	}
	return o.Deletion, true
}

// HasDeletion returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasDeletion() bool {
	if o != nil && o.Deletion != nil {
		return true
	}

	return false
}

// SetDeletion gets a reference to the given bool and assigns it to the Deletion field.
func (o *BTSketchEntity25) SetDeletion(v bool) {
	o.Deletion = &v
}

// GetFeatureIds returns the FeatureIds field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetFeatureIds() []string {
	if o == nil || o.FeatureIds == nil {
		var ret []string
		return ret
	}
	return o.FeatureIds
}

// GetFeatureIdsOk returns a tuple with the FeatureIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetFeatureIdsOk() ([]string, bool) {
	if o == nil || o.FeatureIds == nil {
		return nil, false
	}
	return o.FeatureIds, true
}

// HasFeatureIds returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasFeatureIds() bool {
	if o != nil && o.FeatureIds != nil {
		return true
	}

	return false
}

// SetFeatureIds gets a reference to the given []string and assigns it to the FeatureIds field.
func (o *BTSketchEntity25) SetFeatureIds(v []string) {
	o.FeatureIds = v
}

// GetFromSketch returns the FromSketch field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetFromSketch() bool {
	if o == nil || o.FromSketch == nil {
		var ret bool
		return ret
	}
	return *o.FromSketch
}

// GetFromSketchOk returns a tuple with the FromSketch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetFromSketchOk() (*bool, bool) {
	if o == nil || o.FromSketch == nil {
		return nil, false
	}
	return o.FromSketch, true
}

// HasFromSketch returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasFromSketch() bool {
	if o != nil && o.FromSketch != nil {
		return true
	}

	return false
}

// SetFromSketch gets a reference to the given bool and assigns it to the FromSketch field.
func (o *BTSketchEntity25) SetFromSketch(v bool) {
	o.FromSketch = &v
}

// GetGeometries returns the Geometries field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetGeometries() []BTEntityGeometry35 {
	if o == nil || o.Geometries == nil {
		var ret []BTEntityGeometry35
		return ret
	}
	return o.Geometries
}

// GetGeometriesOk returns a tuple with the Geometries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetGeometriesOk() ([]BTEntityGeometry35, bool) {
	if o == nil || o.Geometries == nil {
		return nil, false
	}
	return o.Geometries, true
}

// HasGeometries returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasGeometries() bool {
	if o != nil && o.Geometries != nil {
		return true
	}

	return false
}

// SetGeometries gets a reference to the given []BTEntityGeometry35 and assigns it to the Geometries field.
func (o *BTSketchEntity25) SetGeometries(v []BTEntityGeometry35) {
	o.Geometries = v
}

// GetDomainSpecificMetadata returns the DomainSpecificMetadata field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetDomainSpecificMetadata() []BTDomainSpecificMetadata961 {
	if o == nil || o.DomainSpecificMetadata == nil {
		var ret []BTDomainSpecificMetadata961
		return ret
	}
	return o.DomainSpecificMetadata
}

// GetDomainSpecificMetadataOk returns a tuple with the DomainSpecificMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetDomainSpecificMetadataOk() ([]BTDomainSpecificMetadata961, bool) {
	if o == nil || o.DomainSpecificMetadata == nil {
		return nil, false
	}
	return o.DomainSpecificMetadata, true
}

// HasDomainSpecificMetadata returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasDomainSpecificMetadata() bool {
	if o != nil && o.DomainSpecificMetadata != nil {
		return true
	}

	return false
}

// SetDomainSpecificMetadata gets a reference to the given []BTDomainSpecificMetadata961 and assigns it to the DomainSpecificMetadata field.
func (o *BTSketchEntity25) SetDomainSpecificMetadata(v []BTDomainSpecificMetadata961) {
	o.DomainSpecificMetadata = v
}

// GetFirstGeometry returns the FirstGeometry field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetFirstGeometry() BTEntityGeometry35 {
	if o == nil || o.FirstGeometry == nil {
		var ret BTEntityGeometry35
		return ret
	}
	return *o.FirstGeometry
}

// GetFirstGeometryOk returns a tuple with the FirstGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetFirstGeometryOk() (*BTEntityGeometry35, bool) {
	if o == nil || o.FirstGeometry == nil {
		return nil, false
	}
	return o.FirstGeometry, true
}

// HasFirstGeometry returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasFirstGeometry() bool {
	if o != nil && o.FirstGeometry != nil {
		return true
	}

	return false
}

// SetFirstGeometry gets a reference to the given BTEntityGeometry35 and assigns it to the FirstGeometry field.
func (o *BTSketchEntity25) SetFirstGeometry(v BTEntityGeometry35) {
	o.FirstGeometry = &v
}

// GetIsConstruction returns the IsConstruction field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetIsConstruction() bool {
	if o == nil || o.IsConstruction == nil {
		var ret bool
		return ret
	}
	return *o.IsConstruction
}

// GetIsConstructionOk returns a tuple with the IsConstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetIsConstructionOk() (*bool, bool) {
	if o == nil || o.IsConstruction == nil {
		return nil, false
	}
	return o.IsConstruction, true
}

// HasIsConstruction returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasIsConstruction() bool {
	if o != nil && o.IsConstruction != nil {
		return true
	}

	return false
}

// SetIsConstruction gets a reference to the given bool and assigns it to the IsConstruction field.
func (o *BTSketchEntity25) SetIsConstruction(v bool) {
	o.IsConstruction = &v
}

// GetIsFromSplineControlPolygon returns the IsFromSplineControlPolygon field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetIsFromSplineControlPolygon() bool {
	if o == nil || o.IsFromSplineControlPolygon == nil {
		var ret bool
		return ret
	}
	return *o.IsFromSplineControlPolygon
}

// GetIsFromSplineControlPolygonOk returns a tuple with the IsFromSplineControlPolygon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetIsFromSplineControlPolygonOk() (*bool, bool) {
	if o == nil || o.IsFromSplineControlPolygon == nil {
		return nil, false
	}
	return o.IsFromSplineControlPolygon, true
}

// HasIsFromSplineControlPolygon returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasIsFromSplineControlPolygon() bool {
	if o != nil && o.IsFromSplineControlPolygon != nil {
		return true
	}

	return false
}

// SetIsFromSplineControlPolygon gets a reference to the given bool and assigns it to the IsFromSplineControlPolygon field.
func (o *BTSketchEntity25) SetIsFromSplineControlPolygon(v bool) {
	o.IsFromSplineControlPolygon = &v
}

// GetIsFromSplineHandle returns the IsFromSplineHandle field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetIsFromSplineHandle() bool {
	if o == nil || o.IsFromSplineHandle == nil {
		var ret bool
		return ret
	}
	return *o.IsFromSplineHandle
}

// GetIsFromSplineHandleOk returns a tuple with the IsFromSplineHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetIsFromSplineHandleOk() (*bool, bool) {
	if o == nil || o.IsFromSplineHandle == nil {
		return nil, false
	}
	return o.IsFromSplineHandle, true
}

// HasIsFromSplineHandle returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasIsFromSplineHandle() bool {
	if o != nil && o.IsFromSplineHandle != nil {
		return true
	}

	return false
}

// SetIsFromSplineHandle gets a reference to the given bool and assigns it to the IsFromSplineHandle field.
func (o *BTSketchEntity25) SetIsFromSplineHandle(v bool) {
	o.IsFromSplineHandle = &v
}

// GetIsTextStroke returns the IsTextStroke field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetIsTextStroke() bool {
	if o == nil || o.IsTextStroke == nil {
		var ret bool
		return ret
	}
	return *o.IsTextStroke
}

// GetIsTextStrokeOk returns a tuple with the IsTextStroke field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetIsTextStrokeOk() (*bool, bool) {
	if o == nil || o.IsTextStroke == nil {
		return nil, false
	}
	return o.IsTextStroke, true
}

// HasIsTextStroke returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasIsTextStroke() bool {
	if o != nil && o.IsTextStroke != nil {
		return true
	}

	return false
}

// SetIsTextStroke gets a reference to the given bool and assigns it to the IsTextStroke field.
func (o *BTSketchEntity25) SetIsTextStroke(v bool) {
	o.IsTextStroke = &v
}

// GetIsUserPoint returns the IsUserPoint field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetIsUserPoint() bool {
	if o == nil || o.IsUserPoint == nil {
		var ret bool
		return ret
	}
	return *o.IsUserPoint
}

// GetIsUserPointOk returns a tuple with the IsUserPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetIsUserPointOk() (*bool, bool) {
	if o == nil || o.IsUserPoint == nil {
		return nil, false
	}
	return o.IsUserPoint, true
}

// HasIsUserPoint returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasIsUserPoint() bool {
	if o != nil && o.IsUserPoint != nil {
		return true
	}

	return false
}

// SetIsUserPoint gets a reference to the given bool and assigns it to the IsUserPoint field.
func (o *BTSketchEntity25) SetIsUserPoint(v bool) {
	o.IsUserPoint = &v
}

// GetSketchCurveType returns the SketchCurveType field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetSketchCurveType() GBTSketchCurveType {
	if o == nil || o.SketchCurveType == nil {
		var ret GBTSketchCurveType
		return ret
	}
	return *o.SketchCurveType
}

// GetSketchCurveTypeOk returns a tuple with the SketchCurveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetSketchCurveTypeOk() (*GBTSketchCurveType, bool) {
	if o == nil || o.SketchCurveType == nil {
		return nil, false
	}
	return o.SketchCurveType, true
}

// HasSketchCurveType returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasSketchCurveType() bool {
	if o != nil && o.SketchCurveType != nil {
		return true
	}

	return false
}

// SetSketchCurveType gets a reference to the given GBTSketchCurveType and assigns it to the SketchCurveType field.
func (o *BTSketchEntity25) SetSketchCurveType(v GBTSketchCurveType) {
	o.SketchCurveType = &v
}

// GetSketchEntityId returns the SketchEntityId field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetSketchEntityId() string {
	if o == nil || o.SketchEntityId == nil {
		var ret string
		return ret
	}
	return *o.SketchEntityId
}

// GetSketchEntityIdOk returns a tuple with the SketchEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetSketchEntityIdOk() (*string, bool) {
	if o == nil || o.SketchEntityId == nil {
		return nil, false
	}
	return o.SketchEntityId, true
}

// HasSketchEntityId returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasSketchEntityId() bool {
	if o != nil && o.SketchEntityId != nil {
		return true
	}

	return false
}

// SetSketchEntityId gets a reference to the given string and assigns it to the SketchEntityId field.
func (o *BTSketchEntity25) SetSketchEntityId(v string) {
	o.SketchEntityId = &v
}

// GetSolveStatus returns the SolveStatus field value if set, zero value otherwise.
func (o *BTSketchEntity25) GetSolveStatus() int32 {
	if o == nil || o.SolveStatus == nil {
		var ret int32
		return ret
	}
	return *o.SolveStatus
}

// GetSolveStatusOk returns a tuple with the SolveStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchEntity25) GetSolveStatusOk() (*int32, bool) {
	if o == nil || o.SolveStatus == nil {
		return nil, false
	}
	return o.SolveStatus, true
}

// HasSolveStatus returns a boolean if a field has been set.
func (o *BTSketchEntity25) HasSolveStatus() bool {
	if o != nil && o.SolveStatus != nil {
		return true
	}

	return false
}

// SetSolveStatus gets a reference to the given int32 and assigns it to the SolveStatus field.
func (o *BTSketchEntity25) SetSolveStatus(v int32) {
	o.SolveStatus = &v
}

func (o BTSketchEntity25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CopyWithoutGeometry != nil {
		toSerialize["copyWithoutGeometry"] = o.CopyWithoutGeometry
	}
	if o.Decompressed != nil {
		toSerialize["decompressed"] = o.Decompressed
	}
	if o.Deletion != nil {
		toSerialize["deletion"] = o.Deletion
	}
	if o.FeatureIds != nil {
		toSerialize["featureIds"] = o.FeatureIds
	}
	if o.FromSketch != nil {
		toSerialize["fromSketch"] = o.FromSketch
	}
	if o.Geometries != nil {
		toSerialize["geometries"] = o.Geometries
	}
	if o.DomainSpecificMetadata != nil {
		toSerialize["domainSpecificMetadata"] = o.DomainSpecificMetadata
	}
	if o.FirstGeometry != nil {
		toSerialize["firstGeometry"] = o.FirstGeometry
	}
	if o.IsConstruction != nil {
		toSerialize["isConstruction"] = o.IsConstruction
	}
	if o.IsFromSplineControlPolygon != nil {
		toSerialize["isFromSplineControlPolygon"] = o.IsFromSplineControlPolygon
	}
	if o.IsFromSplineHandle != nil {
		toSerialize["isFromSplineHandle"] = o.IsFromSplineHandle
	}
	if o.IsTextStroke != nil {
		toSerialize["isTextStroke"] = o.IsTextStroke
	}
	if o.IsUserPoint != nil {
		toSerialize["isUserPoint"] = o.IsUserPoint
	}
	if o.SketchCurveType != nil {
		toSerialize["sketchCurveType"] = o.SketchCurveType
	}
	if o.SketchEntityId != nil {
		toSerialize["sketchEntityId"] = o.SketchEntityId
	}
	if o.SolveStatus != nil {
		toSerialize["solveStatus"] = o.SolveStatus
	}
	return json.Marshal(toSerialize)
}

type NullableBTSketchEntity25 struct {
	value *BTSketchEntity25
	isSet bool
}

func (v NullableBTSketchEntity25) Get() *BTSketchEntity25 {
	return v.value
}

func (v *NullableBTSketchEntity25) Set(val *BTSketchEntity25) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSketchEntity25) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSketchEntity25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSketchEntity25(val *BTSketchEntity25) *NullableBTSketchEntity25 {
	return &NullableBTSketchEntity25{value: val, isSet: true}
}

func (v NullableBTSketchEntity25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSketchEntity25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
