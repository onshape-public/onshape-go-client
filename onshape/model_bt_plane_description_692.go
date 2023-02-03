/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.158.10882-bf44380b65e2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPlaneDescription692 struct for BTPlaneDescription692
type BTPlaneDescription692 struct {
	Direction                 *BTVector3d389 `json:"direction,omitempty"`
	DirectionOrientedWithFace *BTVector3d389 `json:"directionOrientedWithFace,omitempty"`
	Origin                    *BTVector3d389 `json:"origin,omitempty"`
	Type                      *string        `json:"type,omitempty"`
	BtType                    *string        `json:"btType,omitempty"`
	IsOrientedWithFace        *bool          `json:"isOrientedWithFace,omitempty"`
	Normal                    *BTVector3d389 `json:"normal,omitempty"`
}

// NewBTPlaneDescription692 instantiates a new BTPlaneDescription692 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPlaneDescription692() *BTPlaneDescription692 {
	this := BTPlaneDescription692{}
	return &this
}

// NewBTPlaneDescription692WithDefaults instantiates a new BTPlaneDescription692 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPlaneDescription692WithDefaults() *BTPlaneDescription692 {
	this := BTPlaneDescription692{}
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *BTPlaneDescription692) GetDirection() BTVector3d389 {
	if o == nil || o.Direction == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlaneDescription692) GetDirectionOk() (*BTVector3d389, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *BTPlaneDescription692) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given BTVector3d389 and assigns it to the Direction field.
func (o *BTPlaneDescription692) SetDirection(v BTVector3d389) {
	o.Direction = &v
}

// GetDirectionOrientedWithFace returns the DirectionOrientedWithFace field value if set, zero value otherwise.
func (o *BTPlaneDescription692) GetDirectionOrientedWithFace() BTVector3d389 {
	if o == nil || o.DirectionOrientedWithFace == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.DirectionOrientedWithFace
}

// GetDirectionOrientedWithFaceOk returns a tuple with the DirectionOrientedWithFace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlaneDescription692) GetDirectionOrientedWithFaceOk() (*BTVector3d389, bool) {
	if o == nil || o.DirectionOrientedWithFace == nil {
		return nil, false
	}
	return o.DirectionOrientedWithFace, true
}

// HasDirectionOrientedWithFace returns a boolean if a field has been set.
func (o *BTPlaneDescription692) HasDirectionOrientedWithFace() bool {
	if o != nil && o.DirectionOrientedWithFace != nil {
		return true
	}

	return false
}

// SetDirectionOrientedWithFace gets a reference to the given BTVector3d389 and assigns it to the DirectionOrientedWithFace field.
func (o *BTPlaneDescription692) SetDirectionOrientedWithFace(v BTVector3d389) {
	o.DirectionOrientedWithFace = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTPlaneDescription692) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlaneDescription692) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTPlaneDescription692) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTPlaneDescription692) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTPlaneDescription692) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlaneDescription692) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTPlaneDescription692) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTPlaneDescription692) SetType(v string) {
	o.Type = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPlaneDescription692) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlaneDescription692) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPlaneDescription692) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPlaneDescription692) SetBtType(v string) {
	o.BtType = &v
}

// GetIsOrientedWithFace returns the IsOrientedWithFace field value if set, zero value otherwise.
func (o *BTPlaneDescription692) GetIsOrientedWithFace() bool {
	if o == nil || o.IsOrientedWithFace == nil {
		var ret bool
		return ret
	}
	return *o.IsOrientedWithFace
}

// GetIsOrientedWithFaceOk returns a tuple with the IsOrientedWithFace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlaneDescription692) GetIsOrientedWithFaceOk() (*bool, bool) {
	if o == nil || o.IsOrientedWithFace == nil {
		return nil, false
	}
	return o.IsOrientedWithFace, true
}

// HasIsOrientedWithFace returns a boolean if a field has been set.
func (o *BTPlaneDescription692) HasIsOrientedWithFace() bool {
	if o != nil && o.IsOrientedWithFace != nil {
		return true
	}

	return false
}

// SetIsOrientedWithFace gets a reference to the given bool and assigns it to the IsOrientedWithFace field.
func (o *BTPlaneDescription692) SetIsOrientedWithFace(v bool) {
	o.IsOrientedWithFace = &v
}

// GetNormal returns the Normal field value if set, zero value otherwise.
func (o *BTPlaneDescription692) GetNormal() BTVector3d389 {
	if o == nil || o.Normal == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Normal
}

// GetNormalOk returns a tuple with the Normal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlaneDescription692) GetNormalOk() (*BTVector3d389, bool) {
	if o == nil || o.Normal == nil {
		return nil, false
	}
	return o.Normal, true
}

// HasNormal returns a boolean if a field has been set.
func (o *BTPlaneDescription692) HasNormal() bool {
	if o != nil && o.Normal != nil {
		return true
	}

	return false
}

// SetNormal gets a reference to the given BTVector3d389 and assigns it to the Normal field.
func (o *BTPlaneDescription692) SetNormal(v BTVector3d389) {
	o.Normal = &v
}

func (o BTPlaneDescription692) MarshalJSON() ([]byte, error) {
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
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsOrientedWithFace != nil {
		toSerialize["isOrientedWithFace"] = o.IsOrientedWithFace
	}
	if o.Normal != nil {
		toSerialize["normal"] = o.Normal
	}
	return json.Marshal(toSerialize)
}

type NullableBTPlaneDescription692 struct {
	value *BTPlaneDescription692
	isSet bool
}

func (v NullableBTPlaneDescription692) Get() *BTPlaneDescription692 {
	return v.value
}

func (v *NullableBTPlaneDescription692) Set(val *BTPlaneDescription692) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPlaneDescription692) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPlaneDescription692) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPlaneDescription692(val *BTPlaneDescription692) *NullableBTPlaneDescription692 {
	return &NullableBTPlaneDescription692{value: val, isSet: true}
}

func (v NullableBTPlaneDescription692) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPlaneDescription692) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
