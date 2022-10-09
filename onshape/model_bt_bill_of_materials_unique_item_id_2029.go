/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6826-13f16e212af4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBillOfMaterialsUniqueItemId2029 struct for BTBillOfMaterialsUniqueItemId2029
type BTBillOfMaterialsUniqueItemId2029 struct {
	ApiConfiguration  *string                `json:"apiConfiguration,omitempty"`
	BtType            *string                `json:"btType,omitempty"`
	IsStandardContent *bool                  `json:"isStandardContent,omitempty"`
	ItemDefinitionId  *string                `json:"itemDefinitionId,omitempty"`
	PartId            *string                `json:"partId,omitempty"`
	SourceElement     *BTElementReference725 `json:"sourceElement,omitempty"`
}

// NewBTBillOfMaterialsUniqueItemId2029 instantiates a new BTBillOfMaterialsUniqueItemId2029 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillOfMaterialsUniqueItemId2029() *BTBillOfMaterialsUniqueItemId2029 {
	this := BTBillOfMaterialsUniqueItemId2029{}
	return &this
}

// NewBTBillOfMaterialsUniqueItemId2029WithDefaults instantiates a new BTBillOfMaterialsUniqueItemId2029 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillOfMaterialsUniqueItemId2029WithDefaults() *BTBillOfMaterialsUniqueItemId2029 {
	this := BTBillOfMaterialsUniqueItemId2029{}
	return &this
}

// GetApiConfiguration returns the ApiConfiguration field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetApiConfiguration() string {
	if o == nil || o.ApiConfiguration == nil {
		var ret string
		return ret
	}
	return *o.ApiConfiguration
}

// GetApiConfigurationOk returns a tuple with the ApiConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetApiConfigurationOk() (*string, bool) {
	if o == nil || o.ApiConfiguration == nil {
		return nil, false
	}
	return o.ApiConfiguration, true
}

// HasApiConfiguration returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasApiConfiguration() bool {
	if o != nil && o.ApiConfiguration != nil {
		return true
	}

	return false
}

// SetApiConfiguration gets a reference to the given string and assigns it to the ApiConfiguration field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetApiConfiguration(v string) {
	o.ApiConfiguration = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetBtType(v string) {
	o.BtType = &v
}

// GetIsStandardContent returns the IsStandardContent field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetIsStandardContent() bool {
	if o == nil || o.IsStandardContent == nil {
		var ret bool
		return ret
	}
	return *o.IsStandardContent
}

// GetIsStandardContentOk returns a tuple with the IsStandardContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetIsStandardContentOk() (*bool, bool) {
	if o == nil || o.IsStandardContent == nil {
		return nil, false
	}
	return o.IsStandardContent, true
}

// HasIsStandardContent returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasIsStandardContent() bool {
	if o != nil && o.IsStandardContent != nil {
		return true
	}

	return false
}

// SetIsStandardContent gets a reference to the given bool and assigns it to the IsStandardContent field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetIsStandardContent(v bool) {
	o.IsStandardContent = &v
}

// GetItemDefinitionId returns the ItemDefinitionId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetItemDefinitionId() string {
	if o == nil || o.ItemDefinitionId == nil {
		var ret string
		return ret
	}
	return *o.ItemDefinitionId
}

// GetItemDefinitionIdOk returns a tuple with the ItemDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetItemDefinitionIdOk() (*string, bool) {
	if o == nil || o.ItemDefinitionId == nil {
		return nil, false
	}
	return o.ItemDefinitionId, true
}

// HasItemDefinitionId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasItemDefinitionId() bool {
	if o != nil && o.ItemDefinitionId != nil {
		return true
	}

	return false
}

// SetItemDefinitionId gets a reference to the given string and assigns it to the ItemDefinitionId field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetItemDefinitionId(v string) {
	o.ItemDefinitionId = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetPartId(v string) {
	o.PartId = &v
}

// GetSourceElement returns the SourceElement field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetSourceElement() BTElementReference725 {
	if o == nil || o.SourceElement == nil {
		var ret BTElementReference725
		return ret
	}
	return *o.SourceElement
}

// GetSourceElementOk returns a tuple with the SourceElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetSourceElementOk() (*BTElementReference725, bool) {
	if o == nil || o.SourceElement == nil {
		return nil, false
	}
	return o.SourceElement, true
}

// HasSourceElement returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasSourceElement() bool {
	if o != nil && o.SourceElement != nil {
		return true
	}

	return false
}

// SetSourceElement gets a reference to the given BTElementReference725 and assigns it to the SourceElement field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetSourceElement(v BTElementReference725) {
	o.SourceElement = &v
}

func (o BTBillOfMaterialsUniqueItemId2029) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiConfiguration != nil {
		toSerialize["apiConfiguration"] = o.ApiConfiguration
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsStandardContent != nil {
		toSerialize["isStandardContent"] = o.IsStandardContent
	}
	if o.ItemDefinitionId != nil {
		toSerialize["itemDefinitionId"] = o.ItemDefinitionId
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.SourceElement != nil {
		toSerialize["sourceElement"] = o.SourceElement
	}
	return json.Marshal(toSerialize)
}

type NullableBTBillOfMaterialsUniqueItemId2029 struct {
	value *BTBillOfMaterialsUniqueItemId2029
	isSet bool
}

func (v NullableBTBillOfMaterialsUniqueItemId2029) Get() *BTBillOfMaterialsUniqueItemId2029 {
	return v.value
}

func (v *NullableBTBillOfMaterialsUniqueItemId2029) Set(val *BTBillOfMaterialsUniqueItemId2029) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillOfMaterialsUniqueItemId2029) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillOfMaterialsUniqueItemId2029) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillOfMaterialsUniqueItemId2029(val *BTBillOfMaterialsUniqueItemId2029) *NullableBTBillOfMaterialsUniqueItemId2029 {
	return &NullableBTBillOfMaterialsUniqueItemId2029{value: val, isSet: true}
}

func (v NullableBTBillOfMaterialsUniqueItemId2029) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillOfMaterialsUniqueItemId2029) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
