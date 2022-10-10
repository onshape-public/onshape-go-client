/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6832-bc6eeadb5087
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTSurfaceDescription1564 - struct for BTSurfaceDescription1564
type BTSurfaceDescription1564 struct {
	implBTSurfaceDescription1564 interface{}
}

// BTSpunDescription657AsBTSurfaceDescription1564 is a convenience function that returns BTSpunDescription657 wrapped in BTSurfaceDescription1564
func (o *BTSpunDescription657) AsBTSurfaceDescription1564() *BTSurfaceDescription1564 {
	return &BTSurfaceDescription1564{o}
}

// BTSphereDescription1263AsBTSurfaceDescription1564 is a convenience function that returns BTSphereDescription1263 wrapped in BTSurfaceDescription1564
func (o *BTSphereDescription1263) AsBTSurfaceDescription1564() *BTSurfaceDescription1564 {
	return &BTSurfaceDescription1564{o}
}

// BTTorusDescription1834AsBTSurfaceDescription1564 is a convenience function that returns BTTorusDescription1834 wrapped in BTSurfaceDescription1564
func (o *BTTorusDescription1834) AsBTSurfaceDescription1564() *BTSurfaceDescription1564 {
	return &BTSurfaceDescription1564{o}
}

// BTConeDescription860AsBTSurfaceDescription1564 is a convenience function that returns BTConeDescription860 wrapped in BTSurfaceDescription1564
func (o *BTConeDescription860) AsBTSurfaceDescription1564() *BTSurfaceDescription1564 {
	return &BTSurfaceDescription1564{o}
}

// BTPlaneDescription692AsBTSurfaceDescription1564 is a convenience function that returns BTPlaneDescription692 wrapped in BTSurfaceDescription1564
func (o *BTPlaneDescription692) AsBTSurfaceDescription1564() *BTSurfaceDescription1564 {
	return &BTSurfaceDescription1564{o}
}

// BTSweepDescription1473AsBTSurfaceDescription1564 is a convenience function that returns BTSweepDescription1473 wrapped in BTSurfaceDescription1564
func (o *BTSweepDescription1473) AsBTSurfaceDescription1564() *BTSurfaceDescription1564 {
	return &BTSurfaceDescription1564{o}
}

// BTCylinderDescription686AsBTSurfaceDescription1564 is a convenience function that returns BTCylinderDescription686 wrapped in BTSurfaceDescription1564
func (o *BTCylinderDescription686) AsBTSurfaceDescription1564() *BTSurfaceDescription1564 {
	return &BTSurfaceDescription1564{o}
}

// NewBTSurfaceDescription1564 instantiates a new BTSurfaceDescription1564 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSurfaceDescription1564() *BTSurfaceDescription1564 {
	this := BTSurfaceDescription1564{Newbase_BTSurfaceDescription1564()}
	return &this
}

// NewBTSurfaceDescription1564WithDefaults instantiates a new BTSurfaceDescription1564 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSurfaceDescription1564WithDefaults() *BTSurfaceDescription1564 {
	this := BTSurfaceDescription1564{Newbase_BTSurfaceDescription1564WithDefaults()}
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *BTSurfaceDescription1564) GetDirection() BTVector3d389 {
	type getResult interface {
		GetDirection() BTVector3d389
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDirection()
	} else {
		var de BTVector3d389
		return de
	}
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSurfaceDescription1564) GetDirectionOk() (*BTVector3d389, bool) {
	type getResult interface {
		GetDirectionOk() (*BTVector3d389, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDirectionOk()
	} else {
		return nil, false
	}
}

// HasDirection returns a boolean if a field has been set.
func (o *BTSurfaceDescription1564) HasDirection() bool {
	type getResult interface {
		HasDirection() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDirection()
	} else {
		return false
	}
}

// SetDirection gets a reference to the given BTVector3d389 and assigns it to the Direction field.
func (o *BTSurfaceDescription1564) SetDirection(v BTVector3d389) {
	type getResult interface {
		SetDirection(v BTVector3d389)
	}

	o.GetActualInstance().(getResult).SetDirection(v)
}

// GetDirectionOrientedWithFace returns the DirectionOrientedWithFace field value if set, zero value otherwise.
func (o *BTSurfaceDescription1564) GetDirectionOrientedWithFace() BTVector3d389 {
	type getResult interface {
		GetDirectionOrientedWithFace() BTVector3d389
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDirectionOrientedWithFace()
	} else {
		var de BTVector3d389
		return de
	}
}

// GetDirectionOrientedWithFaceOk returns a tuple with the DirectionOrientedWithFace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSurfaceDescription1564) GetDirectionOrientedWithFaceOk() (*BTVector3d389, bool) {
	type getResult interface {
		GetDirectionOrientedWithFaceOk() (*BTVector3d389, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDirectionOrientedWithFaceOk()
	} else {
		return nil, false
	}
}

// HasDirectionOrientedWithFace returns a boolean if a field has been set.
func (o *BTSurfaceDescription1564) HasDirectionOrientedWithFace() bool {
	type getResult interface {
		HasDirectionOrientedWithFace() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDirectionOrientedWithFace()
	} else {
		return false
	}
}

// SetDirectionOrientedWithFace gets a reference to the given BTVector3d389 and assigns it to the DirectionOrientedWithFace field.
func (o *BTSurfaceDescription1564) SetDirectionOrientedWithFace(v BTVector3d389) {
	type getResult interface {
		SetDirectionOrientedWithFace(v BTVector3d389)
	}

	o.GetActualInstance().(getResult).SetDirectionOrientedWithFace(v)
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTSurfaceDescription1564) GetOrigin() BTVector3d389 {
	type getResult interface {
		GetOrigin() BTVector3d389
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOrigin()
	} else {
		var de BTVector3d389
		return de
	}
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSurfaceDescription1564) GetOriginOk() (*BTVector3d389, bool) {
	type getResult interface {
		GetOriginOk() (*BTVector3d389, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOriginOk()
	} else {
		return nil, false
	}
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTSurfaceDescription1564) HasOrigin() bool {
	type getResult interface {
		HasOrigin() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasOrigin()
	} else {
		return false
	}
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTSurfaceDescription1564) SetOrigin(v BTVector3d389) {
	type getResult interface {
		SetOrigin(v BTVector3d389)
	}

	o.GetActualInstance().(getResult).SetOrigin(v)
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTSurfaceDescription1564) GetType() string {
	type getResult interface {
		GetType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetType()
	} else {
		var de string
		return de
	}
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSurfaceDescription1564) GetTypeOk() (*string, bool) {
	type getResult interface {
		GetTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTypeOk()
	} else {
		return nil, false
	}
}

// HasType returns a boolean if a field has been set.
func (o *BTSurfaceDescription1564) HasType() bool {
	type getResult interface {
		HasType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasType()
	} else {
		return false
	}
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTSurfaceDescription1564) SetType(v string) {
	type getResult interface {
		SetType(v string)
	}

	o.GetActualInstance().(getResult).SetType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTSurfaceDescription1564) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTConeDescription-860'
	if jsonDict["btType"] == "BTConeDescription-860" {
		// try to unmarshal JSON data into BTConeDescription860
		var qr *BTConeDescription860
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSurfaceDescription1564 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSurfaceDescription1564 = nil
			return fmt.Errorf("Failed to unmarshal BTSurfaceDescription1564 as BTConeDescription860: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTCylinderDescription-686'
	if jsonDict["btType"] == "BTCylinderDescription-686" {
		// try to unmarshal JSON data into BTCylinderDescription686
		var qr *BTCylinderDescription686
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSurfaceDescription1564 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSurfaceDescription1564 = nil
			return fmt.Errorf("Failed to unmarshal BTSurfaceDescription1564 as BTCylinderDescription686: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTPlaneDescription-692'
	if jsonDict["btType"] == "BTPlaneDescription-692" {
		// try to unmarshal JSON data into BTPlaneDescription692
		var qr *BTPlaneDescription692
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSurfaceDescription1564 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSurfaceDescription1564 = nil
			return fmt.Errorf("Failed to unmarshal BTSurfaceDescription1564 as BTPlaneDescription692: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSphereDescription-1263'
	if jsonDict["btType"] == "BTSphereDescription-1263" {
		// try to unmarshal JSON data into BTSphereDescription1263
		var qr *BTSphereDescription1263
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSurfaceDescription1564 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSurfaceDescription1564 = nil
			return fmt.Errorf("Failed to unmarshal BTSurfaceDescription1564 as BTSphereDescription1263: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSpunDescription-657'
	if jsonDict["btType"] == "BTSpunDescription-657" {
		// try to unmarshal JSON data into BTSpunDescription657
		var qr *BTSpunDescription657
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSurfaceDescription1564 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSurfaceDescription1564 = nil
			return fmt.Errorf("Failed to unmarshal BTSurfaceDescription1564 as BTSpunDescription657: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSweepDescription-1473'
	if jsonDict["btType"] == "BTSweepDescription-1473" {
		// try to unmarshal JSON data into BTSweepDescription1473
		var qr *BTSweepDescription1473
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSurfaceDescription1564 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSurfaceDescription1564 = nil
			return fmt.Errorf("Failed to unmarshal BTSurfaceDescription1564 as BTSweepDescription1473: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTTorusDescription-1834'
	if jsonDict["btType"] == "BTTorusDescription-1834" {
		// try to unmarshal JSON data into BTTorusDescription1834
		var qr *BTTorusDescription1834
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSurfaceDescription1564 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSurfaceDescription1564 = nil
			return fmt.Errorf("Failed to unmarshal BTSurfaceDescription1564 as BTTorusDescription1834: %s", err.Error())
		}
	}

	var qtx *base_BTSurfaceDescription1564
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTSurfaceDescription1564 = qtx
		return nil // data stored in dst.base_BTSurfaceDescription1564, return on the first match
	} else {
		dst.implBTSurfaceDescription1564 = nil
		return fmt.Errorf("Failed to unmarshal BTSurfaceDescription1564 as base_BTSurfaceDescription1564: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTSurfaceDescription1564) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTSurfaceDescription1564) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTSurfaceDescription1564
}

type NullableBTSurfaceDescription1564 struct {
	value *BTSurfaceDescription1564
	isSet bool
}

func (v NullableBTSurfaceDescription1564) Get() *BTSurfaceDescription1564 {
	return v.value
}

func (v *NullableBTSurfaceDescription1564) Set(val *BTSurfaceDescription1564) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSurfaceDescription1564) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSurfaceDescription1564) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSurfaceDescription1564(val *BTSurfaceDescription1564) *NullableBTSurfaceDescription1564 {
	return &NullableBTSurfaceDescription1564{value: val, isSet: true}
}

func (v NullableBTSurfaceDescription1564) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSurfaceDescription1564) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTSurfaceDescription1564 struct {
	Direction                 *BTVector3d389 `json:"direction,omitempty"`
	DirectionOrientedWithFace *BTVector3d389 `json:"directionOrientedWithFace,omitempty"`
	Origin                    *BTVector3d389 `json:"origin,omitempty"`
	Type                      *string        `json:"type,omitempty"`
}

// Newbase_BTSurfaceDescription1564 instantiates a new base_BTSurfaceDescription1564 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTSurfaceDescription1564() *base_BTSurfaceDescription1564 {
	this := base_BTSurfaceDescription1564{}
	return &this
}

// Newbase_BTSurfaceDescription1564WithDefaults instantiates a new base_BTSurfaceDescription1564 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTSurfaceDescription1564WithDefaults() *base_BTSurfaceDescription1564 {
	this := base_BTSurfaceDescription1564{}
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *base_BTSurfaceDescription1564) GetDirection() BTVector3d389 {
	if o == nil || o.Direction == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSurfaceDescription1564) GetDirectionOk() (*BTVector3d389, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *base_BTSurfaceDescription1564) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given BTVector3d389 and assigns it to the Direction field.
func (o *base_BTSurfaceDescription1564) SetDirection(v BTVector3d389) {
	o.Direction = &v
}

// GetDirectionOrientedWithFace returns the DirectionOrientedWithFace field value if set, zero value otherwise.
func (o *base_BTSurfaceDescription1564) GetDirectionOrientedWithFace() BTVector3d389 {
	if o == nil || o.DirectionOrientedWithFace == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.DirectionOrientedWithFace
}

// GetDirectionOrientedWithFaceOk returns a tuple with the DirectionOrientedWithFace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSurfaceDescription1564) GetDirectionOrientedWithFaceOk() (*BTVector3d389, bool) {
	if o == nil || o.DirectionOrientedWithFace == nil {
		return nil, false
	}
	return o.DirectionOrientedWithFace, true
}

// HasDirectionOrientedWithFace returns a boolean if a field has been set.
func (o *base_BTSurfaceDescription1564) HasDirectionOrientedWithFace() bool {
	if o != nil && o.DirectionOrientedWithFace != nil {
		return true
	}

	return false
}

// SetDirectionOrientedWithFace gets a reference to the given BTVector3d389 and assigns it to the DirectionOrientedWithFace field.
func (o *base_BTSurfaceDescription1564) SetDirectionOrientedWithFace(v BTVector3d389) {
	o.DirectionOrientedWithFace = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *base_BTSurfaceDescription1564) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSurfaceDescription1564) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *base_BTSurfaceDescription1564) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *base_BTSurfaceDescription1564) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *base_BTSurfaceDescription1564) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSurfaceDescription1564) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *base_BTSurfaceDescription1564) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *base_BTSurfaceDescription1564) SetType(v string) {
	o.Type = &v
}

func (o base_BTSurfaceDescription1564) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.DirectionOrientedWithFace != nil {
		toSerialize["directionOrientedWithFace"] = o.DirectionOrientedWithFace
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}
