/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.5998-d3227e94fd7e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAllowMeshGeometryFilter1026 struct for BTAllowMeshGeometryFilter1026
type BTAllowMeshGeometryFilter1026 struct {
	BtType             *string `json:"btType,omitempty"`
	AllowsMeshGeometry *bool   `json:"allowsMeshGeometry,omitempty"`
}

// NewBTAllowMeshGeometryFilter1026 instantiates a new BTAllowMeshGeometryFilter1026 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAllowMeshGeometryFilter1026() *BTAllowMeshGeometryFilter1026 {
	this := BTAllowMeshGeometryFilter1026{}
	return &this
}

// NewBTAllowMeshGeometryFilter1026WithDefaults instantiates a new BTAllowMeshGeometryFilter1026 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAllowMeshGeometryFilter1026WithDefaults() *BTAllowMeshGeometryFilter1026 {
	this := BTAllowMeshGeometryFilter1026{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTAllowMeshGeometryFilter1026) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowMeshGeometryFilter1026) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTAllowMeshGeometryFilter1026) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTAllowMeshGeometryFilter1026) SetBtType(v string) {
	o.BtType = &v
}

// GetAllowsMeshGeometry returns the AllowsMeshGeometry field value if set, zero value otherwise.
func (o *BTAllowMeshGeometryFilter1026) GetAllowsMeshGeometry() bool {
	if o == nil || o.AllowsMeshGeometry == nil {
		var ret bool
		return ret
	}
	return *o.AllowsMeshGeometry
}

// GetAllowsMeshGeometryOk returns a tuple with the AllowsMeshGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowMeshGeometryFilter1026) GetAllowsMeshGeometryOk() (*bool, bool) {
	if o == nil || o.AllowsMeshGeometry == nil {
		return nil, false
	}
	return o.AllowsMeshGeometry, true
}

// HasAllowsMeshGeometry returns a boolean if a field has been set.
func (o *BTAllowMeshGeometryFilter1026) HasAllowsMeshGeometry() bool {
	if o != nil && o.AllowsMeshGeometry != nil {
		return true
	}

	return false
}

// SetAllowsMeshGeometry gets a reference to the given bool and assigns it to the AllowsMeshGeometry field.
func (o *BTAllowMeshGeometryFilter1026) SetAllowsMeshGeometry(v bool) {
	o.AllowsMeshGeometry = &v
}

func (o BTAllowMeshGeometryFilter1026) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.AllowsMeshGeometry != nil {
		toSerialize["allowsMeshGeometry"] = o.AllowsMeshGeometry
	}
	return json.Marshal(toSerialize)
}

type NullableBTAllowMeshGeometryFilter1026 struct {
	value *BTAllowMeshGeometryFilter1026
	isSet bool
}

func (v NullableBTAllowMeshGeometryFilter1026) Get() *BTAllowMeshGeometryFilter1026 {
	return v.value
}

func (v *NullableBTAllowMeshGeometryFilter1026) Set(val *BTAllowMeshGeometryFilter1026) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAllowMeshGeometryFilter1026) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAllowMeshGeometryFilter1026) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAllowMeshGeometryFilter1026(val *BTAllowMeshGeometryFilter1026) *NullableBTAllowMeshGeometryFilter1026 {
	return &NullableBTAllowMeshGeometryFilter1026{value: val, isSet: true}
}

func (v NullableBTAllowMeshGeometryFilter1026) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAllowMeshGeometryFilter1026) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
