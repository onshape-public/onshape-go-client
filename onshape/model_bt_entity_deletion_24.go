/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTEntityDeletion24 struct for BTEntityDeletion24
type BTEntityDeletion24 struct {
	BTBaseEntityData33
	BtType              *string              `json:"btType,omitempty"`
	ConstructionPlane   *bool                `json:"constructionPlane,omitempty"`
	CopyWithoutGeometry *BTBaseEntityData33  `json:"copyWithoutGeometry,omitempty"`
	Decompressed        *BTBaseEntityData33  `json:"decompressed,omitempty"`
	Deletion            *bool                `json:"deletion,omitempty"`
	FeatureIds          []string             `json:"featureIds,omitempty"`
	FromSketch          *bool                `json:"fromSketch,omitempty"`
	Geometries          []BTEntityGeometry35 `json:"geometries,omitempty"`
	Origin              *bool                `json:"origin,omitempty"`
	// Type of JSON object.
	BtType *string `json:"btType,omitempty"`
}

// NewBTEntityDeletion24 instantiates a new BTEntityDeletion24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEntityDeletion24() *BTEntityDeletion24 {
	this := BTEntityDeletion24{}
	return &this
}

// NewBTEntityDeletion24WithDefaults instantiates a new BTEntityDeletion24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEntityDeletion24WithDefaults() *BTEntityDeletion24 {
	this := BTEntityDeletion24{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTEntityDeletion24) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityDeletion24) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTEntityDeletion24) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTEntityDeletion24) SetBtType(v string) {
	o.BtType = &v
}

// GetConstructionPlane returns the ConstructionPlane field value if set, zero value otherwise.
func (o *BTEntityDeletion24) GetConstructionPlane() bool {
	if o == nil || o.ConstructionPlane == nil {
		var ret bool
		return ret
	}
	return *o.ConstructionPlane
}

// GetConstructionPlaneOk returns a tuple with the ConstructionPlane field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityDeletion24) GetConstructionPlaneOk() (*bool, bool) {
	if o == nil || o.ConstructionPlane == nil {
		return nil, false
	}
	return o.ConstructionPlane, true
}

// HasConstructionPlane returns a boolean if a field has been set.
func (o *BTEntityDeletion24) HasConstructionPlane() bool {
	if o != nil && o.ConstructionPlane != nil {
		return true
	}

	return false
}

// SetConstructionPlane gets a reference to the given bool and assigns it to the ConstructionPlane field.
func (o *BTEntityDeletion24) SetConstructionPlane(v bool) {
	o.ConstructionPlane = &v
}

// GetCopyWithoutGeometry returns the CopyWithoutGeometry field value if set, zero value otherwise.
func (o *BTEntityDeletion24) GetCopyWithoutGeometry() BTBaseEntityData33 {
	if o == nil || o.CopyWithoutGeometry == nil {
		var ret BTBaseEntityData33
		return ret
	}
	return *o.CopyWithoutGeometry
}

// GetCopyWithoutGeometryOk returns a tuple with the CopyWithoutGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityDeletion24) GetCopyWithoutGeometryOk() (*BTBaseEntityData33, bool) {
	if o == nil || o.CopyWithoutGeometry == nil {
		return nil, false
	}
	return o.CopyWithoutGeometry, true
}

// HasCopyWithoutGeometry returns a boolean if a field has been set.
func (o *BTEntityDeletion24) HasCopyWithoutGeometry() bool {
	if o != nil && o.CopyWithoutGeometry != nil {
		return true
	}

	return false
}

// SetCopyWithoutGeometry gets a reference to the given BTBaseEntityData33 and assigns it to the CopyWithoutGeometry field.
func (o *BTEntityDeletion24) SetCopyWithoutGeometry(v BTBaseEntityData33) {
	o.CopyWithoutGeometry = &v
}

// GetDecompressed returns the Decompressed field value if set, zero value otherwise.
func (o *BTEntityDeletion24) GetDecompressed() BTBaseEntityData33 {
	if o == nil || o.Decompressed == nil {
		var ret BTBaseEntityData33
		return ret
	}
	return *o.Decompressed
}

// GetDecompressedOk returns a tuple with the Decompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityDeletion24) GetDecompressedOk() (*BTBaseEntityData33, bool) {
	if o == nil || o.Decompressed == nil {
		return nil, false
	}
	return o.Decompressed, true
}

// HasDecompressed returns a boolean if a field has been set.
func (o *BTEntityDeletion24) HasDecompressed() bool {
	if o != nil && o.Decompressed != nil {
		return true
	}

	return false
}

// SetDecompressed gets a reference to the given BTBaseEntityData33 and assigns it to the Decompressed field.
func (o *BTEntityDeletion24) SetDecompressed(v BTBaseEntityData33) {
	o.Decompressed = &v
}

// GetDeletion returns the Deletion field value if set, zero value otherwise.
func (o *BTEntityDeletion24) GetDeletion() bool {
	if o == nil || o.Deletion == nil {
		var ret bool
		return ret
	}
	return *o.Deletion
}

// GetDeletionOk returns a tuple with the Deletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityDeletion24) GetDeletionOk() (*bool, bool) {
	if o == nil || o.Deletion == nil {
		return nil, false
	}
	return o.Deletion, true
}

// HasDeletion returns a boolean if a field has been set.
func (o *BTEntityDeletion24) HasDeletion() bool {
	if o != nil && o.Deletion != nil {
		return true
	}

	return false
}

// SetDeletion gets a reference to the given bool and assigns it to the Deletion field.
func (o *BTEntityDeletion24) SetDeletion(v bool) {
	o.Deletion = &v
}

// GetFeatureIds returns the FeatureIds field value if set, zero value otherwise.
func (o *BTEntityDeletion24) GetFeatureIds() []string {
	if o == nil || o.FeatureIds == nil {
		var ret []string
		return ret
	}
	return o.FeatureIds
}

// GetFeatureIdsOk returns a tuple with the FeatureIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityDeletion24) GetFeatureIdsOk() ([]string, bool) {
	if o == nil || o.FeatureIds == nil {
		return nil, false
	}
	return o.FeatureIds, true
}

// HasFeatureIds returns a boolean if a field has been set.
func (o *BTEntityDeletion24) HasFeatureIds() bool {
	if o != nil && o.FeatureIds != nil {
		return true
	}

	return false
}

// SetFeatureIds gets a reference to the given []string and assigns it to the FeatureIds field.
func (o *BTEntityDeletion24) SetFeatureIds(v []string) {
	o.FeatureIds = v
}

// GetFromSketch returns the FromSketch field value if set, zero value otherwise.
func (o *BTEntityDeletion24) GetFromSketch() bool {
	if o == nil || o.FromSketch == nil {
		var ret bool
		return ret
	}
	return *o.FromSketch
}

// GetFromSketchOk returns a tuple with the FromSketch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityDeletion24) GetFromSketchOk() (*bool, bool) {
	if o == nil || o.FromSketch == nil {
		return nil, false
	}
	return o.FromSketch, true
}

// HasFromSketch returns a boolean if a field has been set.
func (o *BTEntityDeletion24) HasFromSketch() bool {
	if o != nil && o.FromSketch != nil {
		return true
	}

	return false
}

// SetFromSketch gets a reference to the given bool and assigns it to the FromSketch field.
func (o *BTEntityDeletion24) SetFromSketch(v bool) {
	o.FromSketch = &v
}

// GetGeometries returns the Geometries field value if set, zero value otherwise.
func (o *BTEntityDeletion24) GetGeometries() []BTEntityGeometry35 {
	if o == nil || o.Geometries == nil {
		var ret []BTEntityGeometry35
		return ret
	}
	return o.Geometries
}

// GetGeometriesOk returns a tuple with the Geometries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityDeletion24) GetGeometriesOk() ([]BTEntityGeometry35, bool) {
	if o == nil || o.Geometries == nil {
		return nil, false
	}
	return o.Geometries, true
}

// HasGeometries returns a boolean if a field has been set.
func (o *BTEntityDeletion24) HasGeometries() bool {
	if o != nil && o.Geometries != nil {
		return true
	}

	return false
}

// SetGeometries gets a reference to the given []BTEntityGeometry35 and assigns it to the Geometries field.
func (o *BTEntityDeletion24) SetGeometries(v []BTEntityGeometry35) {
	o.Geometries = v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTEntityDeletion24) GetOrigin() bool {
	if o == nil || o.Origin == nil {
		var ret bool
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityDeletion24) GetOriginOk() (*bool, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTEntityDeletion24) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given bool and assigns it to the Origin field.
func (o *BTEntityDeletion24) SetOrigin(v bool) {
	o.Origin = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTEntityDeletion24) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityDeletion24) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTEntityDeletion24) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTEntityDeletion24) SetBtType(v string) {
	o.BtType = &v
}

func (o BTEntityDeletion24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTBaseEntityData33, errBTBaseEntityData33 := json.Marshal(o.BTBaseEntityData33)
	if errBTBaseEntityData33 != nil {
		return []byte{}, errBTBaseEntityData33
	}
	errBTBaseEntityData33 = json.Unmarshal([]byte(serializedBTBaseEntityData33), &toSerialize)
	if errBTBaseEntityData33 != nil {
		return []byte{}, errBTBaseEntityData33
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ConstructionPlane != nil {
		toSerialize["constructionPlane"] = o.ConstructionPlane
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
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTEntityDeletion24 struct {
	value *BTEntityDeletion24
	isSet bool
}

func (v NullableBTEntityDeletion24) Get() *BTEntityDeletion24 {
	return v.value
}

func (v *NullableBTEntityDeletion24) Set(val *BTEntityDeletion24) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEntityDeletion24) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEntityDeletion24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEntityDeletion24(val *BTEntityDeletion24) *NullableBTEntityDeletion24 {
	return &NullableBTEntityDeletion24{value: val, isSet: true}
}

func (v NullableBTEntityDeletion24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEntityDeletion24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
