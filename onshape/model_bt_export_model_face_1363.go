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

// BTExportModelFace1363 struct for BTExportModelFace1363
type BTExportModelFace1363 struct {
	// Identifies the application of the appearance. Faces that share a value were assigned an appearance together.
	AppearancePropertyNodeId *string                      `json:"appearancePropertyNodeId,omitempty"`
	Area                     *float64                     `json:"area,omitempty"`
	Box                      *BTBoundingBox1052           `json:"box,omitempty"`
	FaceProperties           *BTExportModelProperties3216 `json:"faceProperties,omitempty"`
	Id                       *string                      `json:"id,omitempty"`
	Loops                    []BTExportModelLoop1182      `json:"loops,omitempty"`
	Orientation              *bool                        `json:"orientation,omitempty"`
	Surface                  *BTSurfaceDescription1564    `json:"surface,omitempty"`
}

// NewBTExportModelFace1363 instantiates a new BTExportModelFace1363 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportModelFace1363() *BTExportModelFace1363 {
	this := BTExportModelFace1363{}
	return &this
}

// NewBTExportModelFace1363WithDefaults instantiates a new BTExportModelFace1363 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportModelFace1363WithDefaults() *BTExportModelFace1363 {
	this := BTExportModelFace1363{}
	return &this
}

// GetAppearancePropertyNodeId returns the AppearancePropertyNodeId field value if set, zero value otherwise.
func (o *BTExportModelFace1363) GetAppearancePropertyNodeId() string {
	if o == nil || o.AppearancePropertyNodeId == nil {
		var ret string
		return ret
	}
	return *o.AppearancePropertyNodeId
}

// GetAppearancePropertyNodeIdOk returns a tuple with the AppearancePropertyNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelFace1363) GetAppearancePropertyNodeIdOk() (*string, bool) {
	if o == nil || o.AppearancePropertyNodeId == nil {
		return nil, false
	}
	return o.AppearancePropertyNodeId, true
}

// HasAppearancePropertyNodeId returns a boolean if a field has been set.
func (o *BTExportModelFace1363) HasAppearancePropertyNodeId() bool {
	if o != nil && o.AppearancePropertyNodeId != nil {
		return true
	}

	return false
}

// SetAppearancePropertyNodeId gets a reference to the given string and assigns it to the AppearancePropertyNodeId field.
func (o *BTExportModelFace1363) SetAppearancePropertyNodeId(v string) {
	o.AppearancePropertyNodeId = &v
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *BTExportModelFace1363) GetArea() float64 {
	if o == nil || o.Area == nil {
		var ret float64
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelFace1363) GetAreaOk() (*float64, bool) {
	if o == nil || o.Area == nil {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *BTExportModelFace1363) HasArea() bool {
	if o != nil && o.Area != nil {
		return true
	}

	return false
}

// SetArea gets a reference to the given float64 and assigns it to the Area field.
func (o *BTExportModelFace1363) SetArea(v float64) {
	o.Area = &v
}

// GetBox returns the Box field value if set, zero value otherwise.
func (o *BTExportModelFace1363) GetBox() BTBoundingBox1052 {
	if o == nil || o.Box == nil {
		var ret BTBoundingBox1052
		return ret
	}
	return *o.Box
}

// GetBoxOk returns a tuple with the Box field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelFace1363) GetBoxOk() (*BTBoundingBox1052, bool) {
	if o == nil || o.Box == nil {
		return nil, false
	}
	return o.Box, true
}

// HasBox returns a boolean if a field has been set.
func (o *BTExportModelFace1363) HasBox() bool {
	if o != nil && o.Box != nil {
		return true
	}

	return false
}

// SetBox gets a reference to the given BTBoundingBox1052 and assigns it to the Box field.
func (o *BTExportModelFace1363) SetBox(v BTBoundingBox1052) {
	o.Box = &v
}

// GetFaceProperties returns the FaceProperties field value if set, zero value otherwise.
func (o *BTExportModelFace1363) GetFaceProperties() BTExportModelProperties3216 {
	if o == nil || o.FaceProperties == nil {
		var ret BTExportModelProperties3216
		return ret
	}
	return *o.FaceProperties
}

// GetFacePropertiesOk returns a tuple with the FaceProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelFace1363) GetFacePropertiesOk() (*BTExportModelProperties3216, bool) {
	if o == nil || o.FaceProperties == nil {
		return nil, false
	}
	return o.FaceProperties, true
}

// HasFaceProperties returns a boolean if a field has been set.
func (o *BTExportModelFace1363) HasFaceProperties() bool {
	if o != nil && o.FaceProperties != nil {
		return true
	}

	return false
}

// SetFaceProperties gets a reference to the given BTExportModelProperties3216 and assigns it to the FaceProperties field.
func (o *BTExportModelFace1363) SetFaceProperties(v BTExportModelProperties3216) {
	o.FaceProperties = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTExportModelFace1363) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelFace1363) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTExportModelFace1363) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTExportModelFace1363) SetId(v string) {
	o.Id = &v
}

// GetLoops returns the Loops field value if set, zero value otherwise.
func (o *BTExportModelFace1363) GetLoops() []BTExportModelLoop1182 {
	if o == nil || o.Loops == nil {
		var ret []BTExportModelLoop1182
		return ret
	}
	return o.Loops
}

// GetLoopsOk returns a tuple with the Loops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelFace1363) GetLoopsOk() ([]BTExportModelLoop1182, bool) {
	if o == nil || o.Loops == nil {
		return nil, false
	}
	return o.Loops, true
}

// HasLoops returns a boolean if a field has been set.
func (o *BTExportModelFace1363) HasLoops() bool {
	if o != nil && o.Loops != nil {
		return true
	}

	return false
}

// SetLoops gets a reference to the given []BTExportModelLoop1182 and assigns it to the Loops field.
func (o *BTExportModelFace1363) SetLoops(v []BTExportModelLoop1182) {
	o.Loops = v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *BTExportModelFace1363) GetOrientation() bool {
	if o == nil || o.Orientation == nil {
		var ret bool
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelFace1363) GetOrientationOk() (*bool, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *BTExportModelFace1363) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given bool and assigns it to the Orientation field.
func (o *BTExportModelFace1363) SetOrientation(v bool) {
	o.Orientation = &v
}

// GetSurface returns the Surface field value if set, zero value otherwise.
func (o *BTExportModelFace1363) GetSurface() BTSurfaceDescription1564 {
	if o == nil || o.Surface == nil {
		var ret BTSurfaceDescription1564
		return ret
	}
	return *o.Surface
}

// GetSurfaceOk returns a tuple with the Surface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelFace1363) GetSurfaceOk() (*BTSurfaceDescription1564, bool) {
	if o == nil || o.Surface == nil {
		return nil, false
	}
	return o.Surface, true
}

// HasSurface returns a boolean if a field has been set.
func (o *BTExportModelFace1363) HasSurface() bool {
	if o != nil && o.Surface != nil {
		return true
	}

	return false
}

// SetSurface gets a reference to the given BTSurfaceDescription1564 and assigns it to the Surface field.
func (o *BTExportModelFace1363) SetSurface(v BTSurfaceDescription1564) {
	o.Surface = &v
}

func (o BTExportModelFace1363) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppearancePropertyNodeId != nil {
		toSerialize["appearancePropertyNodeId"] = o.AppearancePropertyNodeId
	}
	if o.Area != nil {
		toSerialize["area"] = o.Area
	}
	if o.Box != nil {
		toSerialize["box"] = o.Box
	}
	if o.FaceProperties != nil {
		toSerialize["faceProperties"] = o.FaceProperties
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Loops != nil {
		toSerialize["loops"] = o.Loops
	}
	if o.Orientation != nil {
		toSerialize["orientation"] = o.Orientation
	}
	if o.Surface != nil {
		toSerialize["surface"] = o.Surface
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportModelFace1363 struct {
	value *BTExportModelFace1363
	isSet bool
}

func (v NullableBTExportModelFace1363) Get() *BTExportModelFace1363 {
	return v.value
}

func (v *NullableBTExportModelFace1363) Set(val *BTExportModelFace1363) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportModelFace1363) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportModelFace1363) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportModelFace1363(val *BTExportModelFace1363) *NullableBTExportModelFace1363 {
	return &NullableBTExportModelFace1363{value: val, isSet: true}
}

func (v NullableBTExportModelFace1363) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportModelFace1363) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
