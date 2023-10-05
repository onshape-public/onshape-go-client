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
)

// BTExportBodyProperties3559 struct for BTExportBodyProperties3559
type BTExportBodyProperties3559 struct {
	Appearance *BTGraphicsAppearance1152 `json:"appearance,omitempty"`
	BtType     *string                   `json:"btType,omitempty"`
	Name       *string                   `json:"name,omitempty"`
	Material   *BTPartMaterial1445       `json:"material,omitempty"`
	Visibility *GBTPartVisibility        `json:"visibility,omitempty"`
}

// NewBTExportBodyProperties3559 instantiates a new BTExportBodyProperties3559 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportBodyProperties3559() *BTExportBodyProperties3559 {
	this := BTExportBodyProperties3559{}
	return &this
}

// NewBTExportBodyProperties3559WithDefaults instantiates a new BTExportBodyProperties3559 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportBodyProperties3559WithDefaults() *BTExportBodyProperties3559 {
	this := BTExportBodyProperties3559{}
	return &this
}

// GetAppearance returns the Appearance field value if set, zero value otherwise.
func (o *BTExportBodyProperties3559) GetAppearance() BTGraphicsAppearance1152 {
	if o == nil || o.Appearance == nil {
		var ret BTGraphicsAppearance1152
		return ret
	}
	return *o.Appearance
}

// GetAppearanceOk returns a tuple with the Appearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportBodyProperties3559) GetAppearanceOk() (*BTGraphicsAppearance1152, bool) {
	if o == nil || o.Appearance == nil {
		return nil, false
	}
	return o.Appearance, true
}

// HasAppearance returns a boolean if a field has been set.
func (o *BTExportBodyProperties3559) HasAppearance() bool {
	if o != nil && o.Appearance != nil {
		return true
	}

	return false
}

// SetAppearance gets a reference to the given BTGraphicsAppearance1152 and assigns it to the Appearance field.
func (o *BTExportBodyProperties3559) SetAppearance(v BTGraphicsAppearance1152) {
	o.Appearance = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExportBodyProperties3559) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportBodyProperties3559) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExportBodyProperties3559) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExportBodyProperties3559) SetBtType(v string) {
	o.BtType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTExportBodyProperties3559) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportBodyProperties3559) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTExportBodyProperties3559) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTExportBodyProperties3559) SetName(v string) {
	o.Name = &v
}

// GetMaterial returns the Material field value if set, zero value otherwise.
func (o *BTExportBodyProperties3559) GetMaterial() BTPartMaterial1445 {
	if o == nil || o.Material == nil {
		var ret BTPartMaterial1445
		return ret
	}
	return *o.Material
}

// GetMaterialOk returns a tuple with the Material field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportBodyProperties3559) GetMaterialOk() (*BTPartMaterial1445, bool) {
	if o == nil || o.Material == nil {
		return nil, false
	}
	return o.Material, true
}

// HasMaterial returns a boolean if a field has been set.
func (o *BTExportBodyProperties3559) HasMaterial() bool {
	if o != nil && o.Material != nil {
		return true
	}

	return false
}

// SetMaterial gets a reference to the given BTPartMaterial1445 and assigns it to the Material field.
func (o *BTExportBodyProperties3559) SetMaterial(v BTPartMaterial1445) {
	o.Material = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *BTExportBodyProperties3559) GetVisibility() GBTPartVisibility {
	if o == nil || o.Visibility == nil {
		var ret GBTPartVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportBodyProperties3559) GetVisibilityOk() (*GBTPartVisibility, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *BTExportBodyProperties3559) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given GBTPartVisibility and assigns it to the Visibility field.
func (o *BTExportBodyProperties3559) SetVisibility(v GBTPartVisibility) {
	o.Visibility = &v
}

func (o BTExportBodyProperties3559) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Appearance != nil {
		toSerialize["appearance"] = o.Appearance
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Material != nil {
		toSerialize["material"] = o.Material
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportBodyProperties3559 struct {
	value *BTExportBodyProperties3559
	isSet bool
}

func (v NullableBTExportBodyProperties3559) Get() *BTExportBodyProperties3559 {
	return v.value
}

func (v *NullableBTExportBodyProperties3559) Set(val *BTExportBodyProperties3559) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportBodyProperties3559) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportBodyProperties3559) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportBodyProperties3559(val *BTExportBodyProperties3559) *NullableBTExportBodyProperties3559 {
	return &NullableBTExportBodyProperties3559{value: val, isSet: true}
}

func (v NullableBTExportBodyProperties3559) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportBodyProperties3559) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
