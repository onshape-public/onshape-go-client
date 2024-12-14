/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAnnotationGTolDisplayData4887 struct for BTAnnotationGTolDisplayData4887
type BTAnnotationGTolDisplayData4887 struct {
	BTAnnotationDisplayData3225
	AnnotationPlane         []float32                            `json:"annotationPlane,omitempty"`
	BaseNormal              []float32                            `json:"baseNormal,omitempty"`
	BasePoint               []float32                            `json:"basePoint,omitempty"`
	BtType                  *string                              `json:"btType,omitempty"`
	DeterministicId         *string                              `json:"deterministicId,omitempty"`
	DxdySegments            []float32                            `json:"dxdySegments,omitempty"`
	NumberOfLeaderSegements *int32                               `json:"numberOfLeaderSegements,omitempty"`
	Rows                    []BTAnnotationGTolRowDisplayData4397 `json:"rows,omitempty"`
}

// NewBTAnnotationGTolDisplayData4887 instantiates a new BTAnnotationGTolDisplayData4887 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAnnotationGTolDisplayData4887() *BTAnnotationGTolDisplayData4887 {
	this := BTAnnotationGTolDisplayData4887{}
	return &this
}

// NewBTAnnotationGTolDisplayData4887WithDefaults instantiates a new BTAnnotationGTolDisplayData4887 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAnnotationGTolDisplayData4887WithDefaults() *BTAnnotationGTolDisplayData4887 {
	this := BTAnnotationGTolDisplayData4887{}
	return &this
}

// GetAnnotationPlane returns the AnnotationPlane field value if set, zero value otherwise.
func (o *BTAnnotationGTolDisplayData4887) GetAnnotationPlane() []float32 {
	if o == nil || o.AnnotationPlane == nil {
		var ret []float32
		return ret
	}
	return o.AnnotationPlane
}

// GetAnnotationPlaneOk returns a tuple with the AnnotationPlane field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAnnotationGTolDisplayData4887) GetAnnotationPlaneOk() ([]float32, bool) {
	if o == nil || o.AnnotationPlane == nil {
		return nil, false
	}
	return o.AnnotationPlane, true
}

// HasAnnotationPlane returns a boolean if a field has been set.
func (o *BTAnnotationGTolDisplayData4887) HasAnnotationPlane() bool {
	if o != nil && o.AnnotationPlane != nil {
		return true
	}

	return false
}

// SetAnnotationPlane gets a reference to the given []float32 and assigns it to the AnnotationPlane field.
func (o *BTAnnotationGTolDisplayData4887) SetAnnotationPlane(v []float32) {
	o.AnnotationPlane = v
}

// GetBaseNormal returns the BaseNormal field value if set, zero value otherwise.
func (o *BTAnnotationGTolDisplayData4887) GetBaseNormal() []float32 {
	if o == nil || o.BaseNormal == nil {
		var ret []float32
		return ret
	}
	return o.BaseNormal
}

// GetBaseNormalOk returns a tuple with the BaseNormal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAnnotationGTolDisplayData4887) GetBaseNormalOk() ([]float32, bool) {
	if o == nil || o.BaseNormal == nil {
		return nil, false
	}
	return o.BaseNormal, true
}

// HasBaseNormal returns a boolean if a field has been set.
func (o *BTAnnotationGTolDisplayData4887) HasBaseNormal() bool {
	if o != nil && o.BaseNormal != nil {
		return true
	}

	return false
}

// SetBaseNormal gets a reference to the given []float32 and assigns it to the BaseNormal field.
func (o *BTAnnotationGTolDisplayData4887) SetBaseNormal(v []float32) {
	o.BaseNormal = v
}

// GetBasePoint returns the BasePoint field value if set, zero value otherwise.
func (o *BTAnnotationGTolDisplayData4887) GetBasePoint() []float32 {
	if o == nil || o.BasePoint == nil {
		var ret []float32
		return ret
	}
	return o.BasePoint
}

// GetBasePointOk returns a tuple with the BasePoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAnnotationGTolDisplayData4887) GetBasePointOk() ([]float32, bool) {
	if o == nil || o.BasePoint == nil {
		return nil, false
	}
	return o.BasePoint, true
}

// HasBasePoint returns a boolean if a field has been set.
func (o *BTAnnotationGTolDisplayData4887) HasBasePoint() bool {
	if o != nil && o.BasePoint != nil {
		return true
	}

	return false
}

// SetBasePoint gets a reference to the given []float32 and assigns it to the BasePoint field.
func (o *BTAnnotationGTolDisplayData4887) SetBasePoint(v []float32) {
	o.BasePoint = v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTAnnotationGTolDisplayData4887) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAnnotationGTolDisplayData4887) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTAnnotationGTolDisplayData4887) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTAnnotationGTolDisplayData4887) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicId returns the DeterministicId field value if set, zero value otherwise.
func (o *BTAnnotationGTolDisplayData4887) GetDeterministicId() string {
	if o == nil || o.DeterministicId == nil {
		var ret string
		return ret
	}
	return *o.DeterministicId
}

// GetDeterministicIdOk returns a tuple with the DeterministicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAnnotationGTolDisplayData4887) GetDeterministicIdOk() (*string, bool) {
	if o == nil || o.DeterministicId == nil {
		return nil, false
	}
	return o.DeterministicId, true
}

// HasDeterministicId returns a boolean if a field has been set.
func (o *BTAnnotationGTolDisplayData4887) HasDeterministicId() bool {
	if o != nil && o.DeterministicId != nil {
		return true
	}

	return false
}

// SetDeterministicId gets a reference to the given string and assigns it to the DeterministicId field.
func (o *BTAnnotationGTolDisplayData4887) SetDeterministicId(v string) {
	o.DeterministicId = &v
}

// GetDxdySegments returns the DxdySegments field value if set, zero value otherwise.
func (o *BTAnnotationGTolDisplayData4887) GetDxdySegments() []float32 {
	if o == nil || o.DxdySegments == nil {
		var ret []float32
		return ret
	}
	return o.DxdySegments
}

// GetDxdySegmentsOk returns a tuple with the DxdySegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAnnotationGTolDisplayData4887) GetDxdySegmentsOk() ([]float32, bool) {
	if o == nil || o.DxdySegments == nil {
		return nil, false
	}
	return o.DxdySegments, true
}

// HasDxdySegments returns a boolean if a field has been set.
func (o *BTAnnotationGTolDisplayData4887) HasDxdySegments() bool {
	if o != nil && o.DxdySegments != nil {
		return true
	}

	return false
}

// SetDxdySegments gets a reference to the given []float32 and assigns it to the DxdySegments field.
func (o *BTAnnotationGTolDisplayData4887) SetDxdySegments(v []float32) {
	o.DxdySegments = v
}

// GetNumberOfLeaderSegements returns the NumberOfLeaderSegements field value if set, zero value otherwise.
func (o *BTAnnotationGTolDisplayData4887) GetNumberOfLeaderSegements() int32 {
	if o == nil || o.NumberOfLeaderSegements == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfLeaderSegements
}

// GetNumberOfLeaderSegementsOk returns a tuple with the NumberOfLeaderSegements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAnnotationGTolDisplayData4887) GetNumberOfLeaderSegementsOk() (*int32, bool) {
	if o == nil || o.NumberOfLeaderSegements == nil {
		return nil, false
	}
	return o.NumberOfLeaderSegements, true
}

// HasNumberOfLeaderSegements returns a boolean if a field has been set.
func (o *BTAnnotationGTolDisplayData4887) HasNumberOfLeaderSegements() bool {
	if o != nil && o.NumberOfLeaderSegements != nil {
		return true
	}

	return false
}

// SetNumberOfLeaderSegements gets a reference to the given int32 and assigns it to the NumberOfLeaderSegements field.
func (o *BTAnnotationGTolDisplayData4887) SetNumberOfLeaderSegements(v int32) {
	o.NumberOfLeaderSegements = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *BTAnnotationGTolDisplayData4887) GetRows() []BTAnnotationGTolRowDisplayData4397 {
	if o == nil || o.Rows == nil {
		var ret []BTAnnotationGTolRowDisplayData4397
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAnnotationGTolDisplayData4887) GetRowsOk() ([]BTAnnotationGTolRowDisplayData4397, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *BTAnnotationGTolDisplayData4887) HasRows() bool {
	if o != nil && o.Rows != nil {
		return true
	}

	return false
}

// SetRows gets a reference to the given []BTAnnotationGTolRowDisplayData4397 and assigns it to the Rows field.
func (o *BTAnnotationGTolDisplayData4887) SetRows(v []BTAnnotationGTolRowDisplayData4397) {
	o.Rows = v
}

func (o BTAnnotationGTolDisplayData4887) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTAnnotationDisplayData3225, errBTAnnotationDisplayData3225 := json.Marshal(o.BTAnnotationDisplayData3225)
	if errBTAnnotationDisplayData3225 != nil {
		return []byte{}, errBTAnnotationDisplayData3225
	}
	errBTAnnotationDisplayData3225 = json.Unmarshal([]byte(serializedBTAnnotationDisplayData3225), &toSerialize)
	if errBTAnnotationDisplayData3225 != nil {
		return []byte{}, errBTAnnotationDisplayData3225
	}
	if o.AnnotationPlane != nil {
		toSerialize["annotationPlane"] = o.AnnotationPlane
	}
	if o.BaseNormal != nil {
		toSerialize["baseNormal"] = o.BaseNormal
	}
	if o.BasePoint != nil {
		toSerialize["basePoint"] = o.BasePoint
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DeterministicId != nil {
		toSerialize["deterministicId"] = o.DeterministicId
	}
	if o.DxdySegments != nil {
		toSerialize["dxdySegments"] = o.DxdySegments
	}
	if o.NumberOfLeaderSegements != nil {
		toSerialize["numberOfLeaderSegements"] = o.NumberOfLeaderSegements
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	return json.Marshal(toSerialize)
}

type NullableBTAnnotationGTolDisplayData4887 struct {
	value *BTAnnotationGTolDisplayData4887
	isSet bool
}

func (v NullableBTAnnotationGTolDisplayData4887) Get() *BTAnnotationGTolDisplayData4887 {
	return v.value
}

func (v *NullableBTAnnotationGTolDisplayData4887) Set(val *BTAnnotationGTolDisplayData4887) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAnnotationGTolDisplayData4887) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAnnotationGTolDisplayData4887) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAnnotationGTolDisplayData4887(val *BTAnnotationGTolDisplayData4887) *NullableBTAnnotationGTolDisplayData4887 {
	return &NullableBTAnnotationGTolDisplayData4887{value: val, isSet: true}
}

func (v NullableBTAnnotationGTolDisplayData4887) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAnnotationGTolDisplayData4887) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
