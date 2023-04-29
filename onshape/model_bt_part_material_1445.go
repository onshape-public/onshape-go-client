/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.163.15296-122c93d7dbb6
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPartMaterial1445 struct for BTPartMaterial1445
type BTPartMaterial1445 struct {
	BtType           *string                                `json:"btType,omitempty"`
	Id               *string                                `json:"id,omitempty"`
	LibraryName      *string                                `json:"libraryName,omitempty"`
	LibraryReference *BTElementReference725                 `json:"libraryReference,omitempty"`
	Name             *string                                `json:"name,omitempty"`
	Properties       *map[string]BTPartMaterialProperty1453 `json:"properties,omitempty"`
	Version          *int32                                 `json:"version,omitempty"`
}

// NewBTPartMaterial1445 instantiates a new BTPartMaterial1445 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPartMaterial1445() *BTPartMaterial1445 {
	this := BTPartMaterial1445{}
	return &this
}

// NewBTPartMaterial1445WithDefaults instantiates a new BTPartMaterial1445 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPartMaterial1445WithDefaults() *BTPartMaterial1445 {
	this := BTPartMaterial1445{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPartMaterial1445) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterial1445) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPartMaterial1445) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPartMaterial1445) SetBtType(v string) {
	o.BtType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTPartMaterial1445) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterial1445) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTPartMaterial1445) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTPartMaterial1445) SetId(v string) {
	o.Id = &v
}

// GetLibraryName returns the LibraryName field value if set, zero value otherwise.
func (o *BTPartMaterial1445) GetLibraryName() string {
	if o == nil || o.LibraryName == nil {
		var ret string
		return ret
	}
	return *o.LibraryName
}

// GetLibraryNameOk returns a tuple with the LibraryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterial1445) GetLibraryNameOk() (*string, bool) {
	if o == nil || o.LibraryName == nil {
		return nil, false
	}
	return o.LibraryName, true
}

// HasLibraryName returns a boolean if a field has been set.
func (o *BTPartMaterial1445) HasLibraryName() bool {
	if o != nil && o.LibraryName != nil {
		return true
	}

	return false
}

// SetLibraryName gets a reference to the given string and assigns it to the LibraryName field.
func (o *BTPartMaterial1445) SetLibraryName(v string) {
	o.LibraryName = &v
}

// GetLibraryReference returns the LibraryReference field value if set, zero value otherwise.
func (o *BTPartMaterial1445) GetLibraryReference() BTElementReference725 {
	if o == nil || o.LibraryReference == nil {
		var ret BTElementReference725
		return ret
	}
	return *o.LibraryReference
}

// GetLibraryReferenceOk returns a tuple with the LibraryReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterial1445) GetLibraryReferenceOk() (*BTElementReference725, bool) {
	if o == nil || o.LibraryReference == nil {
		return nil, false
	}
	return o.LibraryReference, true
}

// HasLibraryReference returns a boolean if a field has been set.
func (o *BTPartMaterial1445) HasLibraryReference() bool {
	if o != nil && o.LibraryReference != nil {
		return true
	}

	return false
}

// SetLibraryReference gets a reference to the given BTElementReference725 and assigns it to the LibraryReference field.
func (o *BTPartMaterial1445) SetLibraryReference(v BTElementReference725) {
	o.LibraryReference = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPartMaterial1445) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterial1445) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPartMaterial1445) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTPartMaterial1445) SetName(v string) {
	o.Name = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTPartMaterial1445) GetProperties() map[string]BTPartMaterialProperty1453 {
	if o == nil || o.Properties == nil {
		var ret map[string]BTPartMaterialProperty1453
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterial1445) GetPropertiesOk() (*map[string]BTPartMaterialProperty1453, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTPartMaterial1445) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]BTPartMaterialProperty1453 and assigns it to the Properties field.
func (o *BTPartMaterial1445) SetProperties(v map[string]BTPartMaterialProperty1453) {
	o.Properties = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTPartMaterial1445) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterial1445) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTPartMaterial1445) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *BTPartMaterial1445) SetVersion(v int32) {
	o.Version = &v
}

func (o BTPartMaterial1445) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LibraryName != nil {
		toSerialize["libraryName"] = o.LibraryName
	}
	if o.LibraryReference != nil {
		toSerialize["libraryReference"] = o.LibraryReference
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableBTPartMaterial1445 struct {
	value *BTPartMaterial1445
	isSet bool
}

func (v NullableBTPartMaterial1445) Get() *BTPartMaterial1445 {
	return v.value
}

func (v *NullableBTPartMaterial1445) Set(val *BTPartMaterial1445) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPartMaterial1445) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPartMaterial1445) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPartMaterial1445(val *BTPartMaterial1445) *NullableBTPartMaterial1445 {
	return &NullableBTPartMaterial1445{value: val, isSet: true}
}

func (v NullableBTPartMaterial1445) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPartMaterial1445) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
