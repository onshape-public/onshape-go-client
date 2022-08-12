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

// BTBillOfMaterialsRowInfo struct for BTBillOfMaterialsRowInfo
type BTBillOfMaterialsRowInfo struct {
	HeaderIdToValue    map[string]map[string]interface{} `json:"headerIdToValue,omitempty"`
	Href               *string                           `json:"href,omitempty"`
	Id                 *string                           `json:"id,omitempty"`
	IndentLevel        *int32                            `json:"indentLevel,omitempty"`
	ItemSource         *BTBillOfMaterialsItemSourceInfo  `json:"itemSource,omitempty"`
	Name               *string                           `json:"name,omitempty"`
	RelatedOccurrences []string                          `json:"relatedOccurrences,omitempty"`
	RowId              *string                           `json:"rowId,omitempty"`
	ViewRef            *string                           `json:"viewRef,omitempty"`
}

// NewBTBillOfMaterialsRowInfo instantiates a new BTBillOfMaterialsRowInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillOfMaterialsRowInfo() *BTBillOfMaterialsRowInfo {
	this := BTBillOfMaterialsRowInfo{}
	return &this
}

// NewBTBillOfMaterialsRowInfoWithDefaults instantiates a new BTBillOfMaterialsRowInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillOfMaterialsRowInfoWithDefaults() *BTBillOfMaterialsRowInfo {
	this := BTBillOfMaterialsRowInfo{}
	return &this
}

// GetHeaderIdToValue returns the HeaderIdToValue field value if set, zero value otherwise.
func (o *BTBillOfMaterialsRowInfo) GetHeaderIdToValue() map[string]map[string]interface{} {
	if o == nil || o.HeaderIdToValue == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.HeaderIdToValue
}

// GetHeaderIdToValueOk returns a tuple with the HeaderIdToValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsRowInfo) GetHeaderIdToValueOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.HeaderIdToValue == nil {
		return nil, false
	}
	return o.HeaderIdToValue, true
}

// HasHeaderIdToValue returns a boolean if a field has been set.
func (o *BTBillOfMaterialsRowInfo) HasHeaderIdToValue() bool {
	if o != nil && o.HeaderIdToValue != nil {
		return true
	}

	return false
}

// SetHeaderIdToValue gets a reference to the given map[string]map[string]interface{} and assigns it to the HeaderIdToValue field.
func (o *BTBillOfMaterialsRowInfo) SetHeaderIdToValue(v map[string]map[string]interface{}) {
	o.HeaderIdToValue = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTBillOfMaterialsRowInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsRowInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTBillOfMaterialsRowInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTBillOfMaterialsRowInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTBillOfMaterialsRowInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsRowInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsRowInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTBillOfMaterialsRowInfo) SetId(v string) {
	o.Id = &v
}

// GetIndentLevel returns the IndentLevel field value if set, zero value otherwise.
func (o *BTBillOfMaterialsRowInfo) GetIndentLevel() int32 {
	if o == nil || o.IndentLevel == nil {
		var ret int32
		return ret
	}
	return *o.IndentLevel
}

// GetIndentLevelOk returns a tuple with the IndentLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsRowInfo) GetIndentLevelOk() (*int32, bool) {
	if o == nil || o.IndentLevel == nil {
		return nil, false
	}
	return o.IndentLevel, true
}

// HasIndentLevel returns a boolean if a field has been set.
func (o *BTBillOfMaterialsRowInfo) HasIndentLevel() bool {
	if o != nil && o.IndentLevel != nil {
		return true
	}

	return false
}

// SetIndentLevel gets a reference to the given int32 and assigns it to the IndentLevel field.
func (o *BTBillOfMaterialsRowInfo) SetIndentLevel(v int32) {
	o.IndentLevel = &v
}

// GetItemSource returns the ItemSource field value if set, zero value otherwise.
func (o *BTBillOfMaterialsRowInfo) GetItemSource() BTBillOfMaterialsItemSourceInfo {
	if o == nil || o.ItemSource == nil {
		var ret BTBillOfMaterialsItemSourceInfo
		return ret
	}
	return *o.ItemSource
}

// GetItemSourceOk returns a tuple with the ItemSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsRowInfo) GetItemSourceOk() (*BTBillOfMaterialsItemSourceInfo, bool) {
	if o == nil || o.ItemSource == nil {
		return nil, false
	}
	return o.ItemSource, true
}

// HasItemSource returns a boolean if a field has been set.
func (o *BTBillOfMaterialsRowInfo) HasItemSource() bool {
	if o != nil && o.ItemSource != nil {
		return true
	}

	return false
}

// SetItemSource gets a reference to the given BTBillOfMaterialsItemSourceInfo and assigns it to the ItemSource field.
func (o *BTBillOfMaterialsRowInfo) SetItemSource(v BTBillOfMaterialsItemSourceInfo) {
	o.ItemSource = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTBillOfMaterialsRowInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsRowInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTBillOfMaterialsRowInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTBillOfMaterialsRowInfo) SetName(v string) {
	o.Name = &v
}

// GetRelatedOccurrences returns the RelatedOccurrences field value if set, zero value otherwise.
func (o *BTBillOfMaterialsRowInfo) GetRelatedOccurrences() []string {
	if o == nil || o.RelatedOccurrences == nil {
		var ret []string
		return ret
	}
	return o.RelatedOccurrences
}

// GetRelatedOccurrencesOk returns a tuple with the RelatedOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsRowInfo) GetRelatedOccurrencesOk() ([]string, bool) {
	if o == nil || o.RelatedOccurrences == nil {
		return nil, false
	}
	return o.RelatedOccurrences, true
}

// HasRelatedOccurrences returns a boolean if a field has been set.
func (o *BTBillOfMaterialsRowInfo) HasRelatedOccurrences() bool {
	if o != nil && o.RelatedOccurrences != nil {
		return true
	}

	return false
}

// SetRelatedOccurrences gets a reference to the given []string and assigns it to the RelatedOccurrences field.
func (o *BTBillOfMaterialsRowInfo) SetRelatedOccurrences(v []string) {
	o.RelatedOccurrences = v
}

// GetRowId returns the RowId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsRowInfo) GetRowId() string {
	if o == nil || o.RowId == nil {
		var ret string
		return ret
	}
	return *o.RowId
}

// GetRowIdOk returns a tuple with the RowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsRowInfo) GetRowIdOk() (*string, bool) {
	if o == nil || o.RowId == nil {
		return nil, false
	}
	return o.RowId, true
}

// HasRowId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsRowInfo) HasRowId() bool {
	if o != nil && o.RowId != nil {
		return true
	}

	return false
}

// SetRowId gets a reference to the given string and assigns it to the RowId field.
func (o *BTBillOfMaterialsRowInfo) SetRowId(v string) {
	o.RowId = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTBillOfMaterialsRowInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsRowInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTBillOfMaterialsRowInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTBillOfMaterialsRowInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTBillOfMaterialsRowInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HeaderIdToValue != nil {
		toSerialize["headerIdToValue"] = o.HeaderIdToValue
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IndentLevel != nil {
		toSerialize["indentLevel"] = o.IndentLevel
	}
	if o.ItemSource != nil {
		toSerialize["itemSource"] = o.ItemSource
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RelatedOccurrences != nil {
		toSerialize["relatedOccurrences"] = o.RelatedOccurrences
	}
	if o.RowId != nil {
		toSerialize["rowId"] = o.RowId
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTBillOfMaterialsRowInfo struct {
	value *BTBillOfMaterialsRowInfo
	isSet bool
}

func (v NullableBTBillOfMaterialsRowInfo) Get() *BTBillOfMaterialsRowInfo {
	return v.value
}

func (v *NullableBTBillOfMaterialsRowInfo) Set(val *BTBillOfMaterialsRowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillOfMaterialsRowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillOfMaterialsRowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillOfMaterialsRowInfo(val *BTBillOfMaterialsRowInfo) *NullableBTBillOfMaterialsRowInfo {
	return &NullableBTBillOfMaterialsRowInfo{value: val, isSet: true}
}

func (v NullableBTBillOfMaterialsRowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillOfMaterialsRowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
