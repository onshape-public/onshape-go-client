/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23367-59c3c9f9feef
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTFeatureEntity34 - struct for BTFeatureEntity34
type BTFeatureEntity34 struct {
	implBTFeatureEntity34 interface{}
}

// BTMateConnectorEntity28AsBTFeatureEntity34 is a convenience function that returns BTMateConnectorEntity28 wrapped in BTFeatureEntity34
func (o *BTMateConnectorEntity28) AsBTFeatureEntity34() *BTFeatureEntity34 {
	return &BTFeatureEntity34{o}
}

// BTBodyEntity26AsBTFeatureEntity34 is a convenience function that returns BTBodyEntity26 wrapped in BTFeatureEntity34
func (o *BTBodyEntity26) AsBTFeatureEntity34() *BTFeatureEntity34 {
	return &BTFeatureEntity34{o}
}

// BTOriginEntity935AsBTFeatureEntity34 is a convenience function that returns BTOriginEntity935 wrapped in BTFeatureEntity34
func (o *BTOriginEntity935) AsBTFeatureEntity34() *BTFeatureEntity34 {
	return &BTFeatureEntity34{o}
}

// BTSketchEntity25AsBTFeatureEntity34 is a convenience function that returns BTSketchEntity25 wrapped in BTFeatureEntity34
func (o *BTSketchEntity25) AsBTFeatureEntity34() *BTFeatureEntity34 {
	return &BTFeatureEntity34{o}
}

// BTConstructionPlaneEntity27AsBTFeatureEntity34 is a convenience function that returns BTConstructionPlaneEntity27 wrapped in BTFeatureEntity34
func (o *BTConstructionPlaneEntity27) AsBTFeatureEntity34() *BTFeatureEntity34 {
	return &BTFeatureEntity34{o}
}

// BTPointEntity1439AsBTFeatureEntity34 is a convenience function that returns BTPointEntity1439 wrapped in BTFeatureEntity34
func (o *BTPointEntity1439) AsBTFeatureEntity34() *BTFeatureEntity34 {
	return &BTFeatureEntity34{o}
}

// NewBTFeatureEntity34 instantiates a new BTFeatureEntity34 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureEntity34() *BTFeatureEntity34 {
	this := BTFeatureEntity34{Newbase_BTFeatureEntity34()}
	return &this
}

// NewBTFeatureEntity34WithDefaults instantiates a new BTFeatureEntity34 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureEntity34WithDefaults() *BTFeatureEntity34 {
	this := BTFeatureEntity34{Newbase_BTFeatureEntity34WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFeatureEntity34) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureEntity34) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFeatureEntity34) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFeatureEntity34) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetCopyWithoutGeometry returns the CopyWithoutGeometry field value if set, zero value otherwise.
func (o *BTFeatureEntity34) GetCopyWithoutGeometry() BTBaseEntityData33 {
	type getResult interface {
		GetCopyWithoutGeometry() BTBaseEntityData33
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetCopyWithoutGeometry()
	} else {
		var de BTBaseEntityData33
		return de
	}
}

// GetCopyWithoutGeometryOk returns a tuple with the CopyWithoutGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureEntity34) GetCopyWithoutGeometryOk() (*BTBaseEntityData33, bool) {
	type getResult interface {
		GetCopyWithoutGeometryOk() (*BTBaseEntityData33, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetCopyWithoutGeometryOk()
	} else {
		return nil, false
	}
}

// HasCopyWithoutGeometry returns a boolean if a field has been set.
func (o *BTFeatureEntity34) HasCopyWithoutGeometry() bool {
	type getResult interface {
		HasCopyWithoutGeometry() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasCopyWithoutGeometry()
	} else {
		return false
	}
}

// SetCopyWithoutGeometry gets a reference to the given BTBaseEntityData33 and assigns it to the CopyWithoutGeometry field.
func (o *BTFeatureEntity34) SetCopyWithoutGeometry(v BTBaseEntityData33) {
	type getResult interface {
		SetCopyWithoutGeometry(v BTBaseEntityData33)
	}

	o.GetActualInstance().(getResult).SetCopyWithoutGeometry(v)
}

// GetDecompressed returns the Decompressed field value if set, zero value otherwise.
func (o *BTFeatureEntity34) GetDecompressed() BTBaseEntityData33 {
	type getResult interface {
		GetDecompressed() BTBaseEntityData33
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDecompressed()
	} else {
		var de BTBaseEntityData33
		return de
	}
}

// GetDecompressedOk returns a tuple with the Decompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureEntity34) GetDecompressedOk() (*BTBaseEntityData33, bool) {
	type getResult interface {
		GetDecompressedOk() (*BTBaseEntityData33, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDecompressedOk()
	} else {
		return nil, false
	}
}

// HasDecompressed returns a boolean if a field has been set.
func (o *BTFeatureEntity34) HasDecompressed() bool {
	type getResult interface {
		HasDecompressed() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDecompressed()
	} else {
		return false
	}
}

// SetDecompressed gets a reference to the given BTBaseEntityData33 and assigns it to the Decompressed field.
func (o *BTFeatureEntity34) SetDecompressed(v BTBaseEntityData33) {
	type getResult interface {
		SetDecompressed(v BTBaseEntityData33)
	}

	o.GetActualInstance().(getResult).SetDecompressed(v)
}

// GetDeletion returns the Deletion field value if set, zero value otherwise.
func (o *BTFeatureEntity34) GetDeletion() bool {
	type getResult interface {
		GetDeletion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeletion()
	} else {
		var de bool
		return de
	}
}

// GetDeletionOk returns a tuple with the Deletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureEntity34) GetDeletionOk() (*bool, bool) {
	type getResult interface {
		GetDeletionOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeletionOk()
	} else {
		return nil, false
	}
}

// HasDeletion returns a boolean if a field has been set.
func (o *BTFeatureEntity34) HasDeletion() bool {
	type getResult interface {
		HasDeletion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDeletion()
	} else {
		return false
	}
}

// SetDeletion gets a reference to the given bool and assigns it to the Deletion field.
func (o *BTFeatureEntity34) SetDeletion(v bool) {
	type getResult interface {
		SetDeletion(v bool)
	}

	o.GetActualInstance().(getResult).SetDeletion(v)
}

// GetFeatureIds returns the FeatureIds field value if set, zero value otherwise.
func (o *BTFeatureEntity34) GetFeatureIds() []string {
	type getResult interface {
		GetFeatureIds() []string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFeatureIds()
	} else {
		var de []string
		return de
	}
}

// GetFeatureIdsOk returns a tuple with the FeatureIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureEntity34) GetFeatureIdsOk() ([]string, bool) {
	type getResult interface {
		GetFeatureIdsOk() ([]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFeatureIdsOk()
	} else {
		return nil, false
	}
}

// HasFeatureIds returns a boolean if a field has been set.
func (o *BTFeatureEntity34) HasFeatureIds() bool {
	type getResult interface {
		HasFeatureIds() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasFeatureIds()
	} else {
		return false
	}
}

// SetFeatureIds gets a reference to the given []string and assigns it to the FeatureIds field.
func (o *BTFeatureEntity34) SetFeatureIds(v []string) {
	type getResult interface {
		SetFeatureIds(v []string)
	}

	o.GetActualInstance().(getResult).SetFeatureIds(v)
}

// GetFromSketch returns the FromSketch field value if set, zero value otherwise.
func (o *BTFeatureEntity34) GetFromSketch() bool {
	type getResult interface {
		GetFromSketch() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFromSketch()
	} else {
		var de bool
		return de
	}
}

// GetFromSketchOk returns a tuple with the FromSketch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureEntity34) GetFromSketchOk() (*bool, bool) {
	type getResult interface {
		GetFromSketchOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFromSketchOk()
	} else {
		return nil, false
	}
}

// HasFromSketch returns a boolean if a field has been set.
func (o *BTFeatureEntity34) HasFromSketch() bool {
	type getResult interface {
		HasFromSketch() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasFromSketch()
	} else {
		return false
	}
}

// SetFromSketch gets a reference to the given bool and assigns it to the FromSketch field.
func (o *BTFeatureEntity34) SetFromSketch(v bool) {
	type getResult interface {
		SetFromSketch(v bool)
	}

	o.GetActualInstance().(getResult).SetFromSketch(v)
}

// GetGeometries returns the Geometries field value if set, zero value otherwise.
func (o *BTFeatureEntity34) GetGeometries() []BTEntityGeometry35 {
	type getResult interface {
		GetGeometries() []BTEntityGeometry35
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetGeometries()
	} else {
		var de []BTEntityGeometry35
		return de
	}
}

// GetGeometriesOk returns a tuple with the Geometries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureEntity34) GetGeometriesOk() ([]BTEntityGeometry35, bool) {
	type getResult interface {
		GetGeometriesOk() ([]BTEntityGeometry35, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetGeometriesOk()
	} else {
		return nil, false
	}
}

// HasGeometries returns a boolean if a field has been set.
func (o *BTFeatureEntity34) HasGeometries() bool {
	type getResult interface {
		HasGeometries() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasGeometries()
	} else {
		return false
	}
}

// SetGeometries gets a reference to the given []BTEntityGeometry35 and assigns it to the Geometries field.
func (o *BTFeatureEntity34) SetGeometries(v []BTEntityGeometry35) {
	type getResult interface {
		SetGeometries(v []BTEntityGeometry35)
	}

	o.GetActualInstance().(getResult).SetGeometries(v)
}

// GetDomainSpecificMetadata returns the DomainSpecificMetadata field value if set, zero value otherwise.
func (o *BTFeatureEntity34) GetDomainSpecificMetadata() []BTDomainSpecificMetadata961 {
	type getResult interface {
		GetDomainSpecificMetadata() []BTDomainSpecificMetadata961
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDomainSpecificMetadata()
	} else {
		var de []BTDomainSpecificMetadata961
		return de
	}
}

// GetDomainSpecificMetadataOk returns a tuple with the DomainSpecificMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureEntity34) GetDomainSpecificMetadataOk() ([]BTDomainSpecificMetadata961, bool) {
	type getResult interface {
		GetDomainSpecificMetadataOk() ([]BTDomainSpecificMetadata961, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDomainSpecificMetadataOk()
	} else {
		return nil, false
	}
}

// HasDomainSpecificMetadata returns a boolean if a field has been set.
func (o *BTFeatureEntity34) HasDomainSpecificMetadata() bool {
	type getResult interface {
		HasDomainSpecificMetadata() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDomainSpecificMetadata()
	} else {
		return false
	}
}

// SetDomainSpecificMetadata gets a reference to the given []BTDomainSpecificMetadata961 and assigns it to the DomainSpecificMetadata field.
func (o *BTFeatureEntity34) SetDomainSpecificMetadata(v []BTDomainSpecificMetadata961) {
	type getResult interface {
		SetDomainSpecificMetadata(v []BTDomainSpecificMetadata961)
	}

	o.GetActualInstance().(getResult).SetDomainSpecificMetadata(v)
}

// GetFirstGeometry returns the FirstGeometry field value if set, zero value otherwise.
func (o *BTFeatureEntity34) GetFirstGeometry() BTEntityGeometry35 {
	type getResult interface {
		GetFirstGeometry() BTEntityGeometry35
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFirstGeometry()
	} else {
		var de BTEntityGeometry35
		return de
	}
}

// GetFirstGeometryOk returns a tuple with the FirstGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureEntity34) GetFirstGeometryOk() (*BTEntityGeometry35, bool) {
	type getResult interface {
		GetFirstGeometryOk() (*BTEntityGeometry35, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFirstGeometryOk()
	} else {
		return nil, false
	}
}

// HasFirstGeometry returns a boolean if a field has been set.
func (o *BTFeatureEntity34) HasFirstGeometry() bool {
	type getResult interface {
		HasFirstGeometry() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasFirstGeometry()
	} else {
		return false
	}
}

// SetFirstGeometry gets a reference to the given BTEntityGeometry35 and assigns it to the FirstGeometry field.
func (o *BTFeatureEntity34) SetFirstGeometry(v BTEntityGeometry35) {
	type getResult interface {
		SetFirstGeometry(v BTEntityGeometry35)
	}

	o.GetActualInstance().(getResult).SetFirstGeometry(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTFeatureEntity34) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTBodyEntity-26'
	if jsonDict["btType"] == "BTBodyEntity-26" {
		// try to unmarshal JSON data into BTBodyEntity26
		var qr *BTBodyEntity26
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureEntity34 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureEntity34 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureEntity34 as BTBodyEntity26: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTConstructionPlaneEntity-27'
	if jsonDict["btType"] == "BTConstructionPlaneEntity-27" {
		// try to unmarshal JSON data into BTConstructionPlaneEntity27
		var qr *BTConstructionPlaneEntity27
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureEntity34 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureEntity34 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureEntity34 as BTConstructionPlaneEntity27: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMateConnectorEntity-28'
	if jsonDict["btType"] == "BTMateConnectorEntity-28" {
		// try to unmarshal JSON data into BTMateConnectorEntity28
		var qr *BTMateConnectorEntity28
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureEntity34 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureEntity34 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureEntity34 as BTMateConnectorEntity28: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTOriginEntity-935'
	if jsonDict["btType"] == "BTOriginEntity-935" {
		// try to unmarshal JSON data into BTOriginEntity935
		var qr *BTOriginEntity935
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureEntity34 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureEntity34 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureEntity34 as BTOriginEntity935: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTPointEntity-1439'
	if jsonDict["btType"] == "BTPointEntity-1439" {
		// try to unmarshal JSON data into BTPointEntity1439
		var qr *BTPointEntity1439
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureEntity34 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureEntity34 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureEntity34 as BTPointEntity1439: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSketchEntity-25'
	if jsonDict["btType"] == "BTSketchEntity-25" {
		// try to unmarshal JSON data into BTSketchEntity25
		var qr *BTSketchEntity25
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureEntity34 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureEntity34 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureEntity34 as BTSketchEntity25: %s", err.Error())
		}
	}

	var qtx *base_BTFeatureEntity34
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTFeatureEntity34 = qtx
		return nil // data stored in dst.base_BTFeatureEntity34, return on the first match
	} else {
		dst.implBTFeatureEntity34 = nil
		return fmt.Errorf("Failed to unmarshal BTFeatureEntity34 as base_BTFeatureEntity34: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTFeatureEntity34) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTFeatureEntity34) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTFeatureEntity34
}

type NullableBTFeatureEntity34 struct {
	value *BTFeatureEntity34
	isSet bool
}

func (v NullableBTFeatureEntity34) Get() *BTFeatureEntity34 {
	return v.value
}

func (v *NullableBTFeatureEntity34) Set(val *BTFeatureEntity34) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureEntity34) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureEntity34) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureEntity34(val *BTFeatureEntity34) *NullableBTFeatureEntity34 {
	return &NullableBTFeatureEntity34{value: val, isSet: true}
}

func (v NullableBTFeatureEntity34) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureEntity34) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTFeatureEntity34 struct {
	BtType                 *string                       `json:"btType,omitempty"`
	CopyWithoutGeometry    *BTBaseEntityData33           `json:"copyWithoutGeometry,omitempty"`
	Decompressed           *BTBaseEntityData33           `json:"decompressed,omitempty"`
	Deletion               *bool                         `json:"deletion,omitempty"`
	FeatureIds             []string                      `json:"featureIds,omitempty"`
	FromSketch             *bool                         `json:"fromSketch,omitempty"`
	Geometries             []BTEntityGeometry35          `json:"geometries,omitempty"`
	DomainSpecificMetadata []BTDomainSpecificMetadata961 `json:"domainSpecificMetadata,omitempty"`
	FirstGeometry          *BTEntityGeometry35           `json:"firstGeometry,omitempty"`
}

// Newbase_BTFeatureEntity34 instantiates a new base_BTFeatureEntity34 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTFeatureEntity34() *base_BTFeatureEntity34 {
	this := base_BTFeatureEntity34{}
	return &this
}

// Newbase_BTFeatureEntity34WithDefaults instantiates a new base_BTFeatureEntity34 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTFeatureEntity34WithDefaults() *base_BTFeatureEntity34 {
	this := base_BTFeatureEntity34{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTFeatureEntity34) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureEntity34) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTFeatureEntity34) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTFeatureEntity34) SetBtType(v string) {
	o.BtType = &v
}

// GetCopyWithoutGeometry returns the CopyWithoutGeometry field value if set, zero value otherwise.
func (o *base_BTFeatureEntity34) GetCopyWithoutGeometry() BTBaseEntityData33 {
	if o == nil || o.CopyWithoutGeometry == nil {
		var ret BTBaseEntityData33
		return ret
	}
	return *o.CopyWithoutGeometry
}

// GetCopyWithoutGeometryOk returns a tuple with the CopyWithoutGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureEntity34) GetCopyWithoutGeometryOk() (*BTBaseEntityData33, bool) {
	if o == nil || o.CopyWithoutGeometry == nil {
		return nil, false
	}
	return o.CopyWithoutGeometry, true
}

// HasCopyWithoutGeometry returns a boolean if a field has been set.
func (o *base_BTFeatureEntity34) HasCopyWithoutGeometry() bool {
	if o != nil && o.CopyWithoutGeometry != nil {
		return true
	}

	return false
}

// SetCopyWithoutGeometry gets a reference to the given BTBaseEntityData33 and assigns it to the CopyWithoutGeometry field.
func (o *base_BTFeatureEntity34) SetCopyWithoutGeometry(v BTBaseEntityData33) {
	o.CopyWithoutGeometry = &v
}

// GetDecompressed returns the Decompressed field value if set, zero value otherwise.
func (o *base_BTFeatureEntity34) GetDecompressed() BTBaseEntityData33 {
	if o == nil || o.Decompressed == nil {
		var ret BTBaseEntityData33
		return ret
	}
	return *o.Decompressed
}

// GetDecompressedOk returns a tuple with the Decompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureEntity34) GetDecompressedOk() (*BTBaseEntityData33, bool) {
	if o == nil || o.Decompressed == nil {
		return nil, false
	}
	return o.Decompressed, true
}

// HasDecompressed returns a boolean if a field has been set.
func (o *base_BTFeatureEntity34) HasDecompressed() bool {
	if o != nil && o.Decompressed != nil {
		return true
	}

	return false
}

// SetDecompressed gets a reference to the given BTBaseEntityData33 and assigns it to the Decompressed field.
func (o *base_BTFeatureEntity34) SetDecompressed(v BTBaseEntityData33) {
	o.Decompressed = &v
}

// GetDeletion returns the Deletion field value if set, zero value otherwise.
func (o *base_BTFeatureEntity34) GetDeletion() bool {
	if o == nil || o.Deletion == nil {
		var ret bool
		return ret
	}
	return *o.Deletion
}

// GetDeletionOk returns a tuple with the Deletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureEntity34) GetDeletionOk() (*bool, bool) {
	if o == nil || o.Deletion == nil {
		return nil, false
	}
	return o.Deletion, true
}

// HasDeletion returns a boolean if a field has been set.
func (o *base_BTFeatureEntity34) HasDeletion() bool {
	if o != nil && o.Deletion != nil {
		return true
	}

	return false
}

// SetDeletion gets a reference to the given bool and assigns it to the Deletion field.
func (o *base_BTFeatureEntity34) SetDeletion(v bool) {
	o.Deletion = &v
}

// GetFeatureIds returns the FeatureIds field value if set, zero value otherwise.
func (o *base_BTFeatureEntity34) GetFeatureIds() []string {
	if o == nil || o.FeatureIds == nil {
		var ret []string
		return ret
	}
	return o.FeatureIds
}

// GetFeatureIdsOk returns a tuple with the FeatureIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureEntity34) GetFeatureIdsOk() ([]string, bool) {
	if o == nil || o.FeatureIds == nil {
		return nil, false
	}
	return o.FeatureIds, true
}

// HasFeatureIds returns a boolean if a field has been set.
func (o *base_BTFeatureEntity34) HasFeatureIds() bool {
	if o != nil && o.FeatureIds != nil {
		return true
	}

	return false
}

// SetFeatureIds gets a reference to the given []string and assigns it to the FeatureIds field.
func (o *base_BTFeatureEntity34) SetFeatureIds(v []string) {
	o.FeatureIds = v
}

// GetFromSketch returns the FromSketch field value if set, zero value otherwise.
func (o *base_BTFeatureEntity34) GetFromSketch() bool {
	if o == nil || o.FromSketch == nil {
		var ret bool
		return ret
	}
	return *o.FromSketch
}

// GetFromSketchOk returns a tuple with the FromSketch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureEntity34) GetFromSketchOk() (*bool, bool) {
	if o == nil || o.FromSketch == nil {
		return nil, false
	}
	return o.FromSketch, true
}

// HasFromSketch returns a boolean if a field has been set.
func (o *base_BTFeatureEntity34) HasFromSketch() bool {
	if o != nil && o.FromSketch != nil {
		return true
	}

	return false
}

// SetFromSketch gets a reference to the given bool and assigns it to the FromSketch field.
func (o *base_BTFeatureEntity34) SetFromSketch(v bool) {
	o.FromSketch = &v
}

// GetGeometries returns the Geometries field value if set, zero value otherwise.
func (o *base_BTFeatureEntity34) GetGeometries() []BTEntityGeometry35 {
	if o == nil || o.Geometries == nil {
		var ret []BTEntityGeometry35
		return ret
	}
	return o.Geometries
}

// GetGeometriesOk returns a tuple with the Geometries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureEntity34) GetGeometriesOk() ([]BTEntityGeometry35, bool) {
	if o == nil || o.Geometries == nil {
		return nil, false
	}
	return o.Geometries, true
}

// HasGeometries returns a boolean if a field has been set.
func (o *base_BTFeatureEntity34) HasGeometries() bool {
	if o != nil && o.Geometries != nil {
		return true
	}

	return false
}

// SetGeometries gets a reference to the given []BTEntityGeometry35 and assigns it to the Geometries field.
func (o *base_BTFeatureEntity34) SetGeometries(v []BTEntityGeometry35) {
	o.Geometries = v
}

// GetDomainSpecificMetadata returns the DomainSpecificMetadata field value if set, zero value otherwise.
func (o *base_BTFeatureEntity34) GetDomainSpecificMetadata() []BTDomainSpecificMetadata961 {
	if o == nil || o.DomainSpecificMetadata == nil {
		var ret []BTDomainSpecificMetadata961
		return ret
	}
	return o.DomainSpecificMetadata
}

// GetDomainSpecificMetadataOk returns a tuple with the DomainSpecificMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureEntity34) GetDomainSpecificMetadataOk() ([]BTDomainSpecificMetadata961, bool) {
	if o == nil || o.DomainSpecificMetadata == nil {
		return nil, false
	}
	return o.DomainSpecificMetadata, true
}

// HasDomainSpecificMetadata returns a boolean if a field has been set.
func (o *base_BTFeatureEntity34) HasDomainSpecificMetadata() bool {
	if o != nil && o.DomainSpecificMetadata != nil {
		return true
	}

	return false
}

// SetDomainSpecificMetadata gets a reference to the given []BTDomainSpecificMetadata961 and assigns it to the DomainSpecificMetadata field.
func (o *base_BTFeatureEntity34) SetDomainSpecificMetadata(v []BTDomainSpecificMetadata961) {
	o.DomainSpecificMetadata = v
}

// GetFirstGeometry returns the FirstGeometry field value if set, zero value otherwise.
func (o *base_BTFeatureEntity34) GetFirstGeometry() BTEntityGeometry35 {
	if o == nil || o.FirstGeometry == nil {
		var ret BTEntityGeometry35
		return ret
	}
	return *o.FirstGeometry
}

// GetFirstGeometryOk returns a tuple with the FirstGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureEntity34) GetFirstGeometryOk() (*BTEntityGeometry35, bool) {
	if o == nil || o.FirstGeometry == nil {
		return nil, false
	}
	return o.FirstGeometry, true
}

// HasFirstGeometry returns a boolean if a field has been set.
func (o *base_BTFeatureEntity34) HasFirstGeometry() bool {
	if o != nil && o.FirstGeometry != nil {
		return true
	}

	return false
}

// SetFirstGeometry gets a reference to the given BTEntityGeometry35 and assigns it to the FirstGeometry field.
func (o *base_BTFeatureEntity34) SetFirstGeometry(v BTEntityGeometry35) {
	o.FirstGeometry = &v
}

func (o base_BTFeatureEntity34) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}
