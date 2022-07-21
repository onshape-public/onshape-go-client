/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.150.5633-5ed6b38daa6b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBillOfMaterialsHeaderInfo struct for BTBillOfMaterialsHeaderInfo
type BTBillOfMaterialsHeaderInfo struct {
	Href         *string `json:"href,omitempty"`
	Id           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	PropertyName *string `json:"propertyName,omitempty"`
	ValueType    *string `json:"valueType,omitempty"`
	ViewRef      *string `json:"viewRef,omitempty"`
	Visible      *bool   `json:"visible,omitempty"`
}

// NewBTBillOfMaterialsHeaderInfo instantiates a new BTBillOfMaterialsHeaderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillOfMaterialsHeaderInfo() *BTBillOfMaterialsHeaderInfo {
	this := BTBillOfMaterialsHeaderInfo{}
	return &this
}

// NewBTBillOfMaterialsHeaderInfoWithDefaults instantiates a new BTBillOfMaterialsHeaderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillOfMaterialsHeaderInfoWithDefaults() *BTBillOfMaterialsHeaderInfo {
	this := BTBillOfMaterialsHeaderInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTBillOfMaterialsHeaderInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsHeaderInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTBillOfMaterialsHeaderInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTBillOfMaterialsHeaderInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTBillOfMaterialsHeaderInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsHeaderInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsHeaderInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTBillOfMaterialsHeaderInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTBillOfMaterialsHeaderInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsHeaderInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTBillOfMaterialsHeaderInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTBillOfMaterialsHeaderInfo) SetName(v string) {
	o.Name = &v
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *BTBillOfMaterialsHeaderInfo) GetPropertyName() string {
	if o == nil || o.PropertyName == nil {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsHeaderInfo) GetPropertyNameOk() (*string, bool) {
	if o == nil || o.PropertyName == nil {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *BTBillOfMaterialsHeaderInfo) HasPropertyName() bool {
	if o != nil && o.PropertyName != nil {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *BTBillOfMaterialsHeaderInfo) SetPropertyName(v string) {
	o.PropertyName = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *BTBillOfMaterialsHeaderInfo) GetValueType() string {
	if o == nil || o.ValueType == nil {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsHeaderInfo) GetValueTypeOk() (*string, bool) {
	if o == nil || o.ValueType == nil {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *BTBillOfMaterialsHeaderInfo) HasValueType() bool {
	if o != nil && o.ValueType != nil {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *BTBillOfMaterialsHeaderInfo) SetValueType(v string) {
	o.ValueType = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTBillOfMaterialsHeaderInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsHeaderInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTBillOfMaterialsHeaderInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTBillOfMaterialsHeaderInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *BTBillOfMaterialsHeaderInfo) GetVisible() bool {
	if o == nil || o.Visible == nil {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsHeaderInfo) GetVisibleOk() (*bool, bool) {
	if o == nil || o.Visible == nil {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *BTBillOfMaterialsHeaderInfo) HasVisible() bool {
	if o != nil && o.Visible != nil {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *BTBillOfMaterialsHeaderInfo) SetVisible(v bool) {
	o.Visible = &v
}

func (o BTBillOfMaterialsHeaderInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PropertyName != nil {
		toSerialize["propertyName"] = o.PropertyName
	}
	if o.ValueType != nil {
		toSerialize["valueType"] = o.ValueType
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.Visible != nil {
		toSerialize["visible"] = o.Visible
	}
	return json.Marshal(toSerialize)
}

type NullableBTBillOfMaterialsHeaderInfo struct {
	value *BTBillOfMaterialsHeaderInfo
	isSet bool
}

func (v NullableBTBillOfMaterialsHeaderInfo) Get() *BTBillOfMaterialsHeaderInfo {
	return v.value
}

func (v *NullableBTBillOfMaterialsHeaderInfo) Set(val *BTBillOfMaterialsHeaderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillOfMaterialsHeaderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillOfMaterialsHeaderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillOfMaterialsHeaderInfo(val *BTBillOfMaterialsHeaderInfo) *NullableBTBillOfMaterialsHeaderInfo {
	return &NullableBTBillOfMaterialsHeaderInfo{value: val, isSet: true}
}

func (v NullableBTBillOfMaterialsHeaderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillOfMaterialsHeaderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
