/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6438-5fd2d9755d52
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAPIApplicationExtensionInfo struct for BTAPIApplicationExtensionInfo
type BTAPIApplicationExtensionInfo struct {
	ActionBody             *string `json:"actionBody,omitempty"`
	ActionType             *int32  `json:"actionType,omitempty"`
	ActionUrl              *string `json:"actionUrl,omitempty"`
	ApplicationId          *string `json:"applicationId,omitempty"`
	ClientId               *string `json:"clientId,omitempty"`
	Description            *string `json:"description,omitempty"`
	ExtensionContext       *int64  `json:"extensionContext,omitempty"`
	ExtensionLocation      *int64  `json:"extensionLocation,omitempty"`
	HasIcon                *bool   `json:"hasIcon,omitempty"`
	HasPendingIcon         *bool   `json:"hasPendingIcon,omitempty"`
	Href                   *string `json:"href,omitempty"`
	IconUrl                *string `json:"iconUrl,omitempty"`
	Id                     *string `json:"id,omitempty"`
	Name                   *string `json:"name,omitempty"`
	ParentAppPrimaryFormat *string `json:"parentAppPrimaryFormat,omitempty"`
	ShowResponse           *bool   `json:"showResponse,omitempty"`
	ViewRef                *string `json:"viewRef,omitempty"`
}

// NewBTAPIApplicationExtensionInfo instantiates a new BTAPIApplicationExtensionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAPIApplicationExtensionInfo() *BTAPIApplicationExtensionInfo {
	this := BTAPIApplicationExtensionInfo{}
	return &this
}

// NewBTAPIApplicationExtensionInfoWithDefaults instantiates a new BTAPIApplicationExtensionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAPIApplicationExtensionInfoWithDefaults() *BTAPIApplicationExtensionInfo {
	this := BTAPIApplicationExtensionInfo{}
	return &this
}

// GetActionBody returns the ActionBody field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetActionBody() string {
	if o == nil || o.ActionBody == nil {
		var ret string
		return ret
	}
	return *o.ActionBody
}

// GetActionBodyOk returns a tuple with the ActionBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetActionBodyOk() (*string, bool) {
	if o == nil || o.ActionBody == nil {
		return nil, false
	}
	return o.ActionBody, true
}

// HasActionBody returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasActionBody() bool {
	if o != nil && o.ActionBody != nil {
		return true
	}

	return false
}

// SetActionBody gets a reference to the given string and assigns it to the ActionBody field.
func (o *BTAPIApplicationExtensionInfo) SetActionBody(v string) {
	o.ActionBody = &v
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetActionType() int32 {
	if o == nil || o.ActionType == nil {
		var ret int32
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetActionTypeOk() (*int32, bool) {
	if o == nil || o.ActionType == nil {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasActionType() bool {
	if o != nil && o.ActionType != nil {
		return true
	}

	return false
}

// SetActionType gets a reference to the given int32 and assigns it to the ActionType field.
func (o *BTAPIApplicationExtensionInfo) SetActionType(v int32) {
	o.ActionType = &v
}

// GetActionUrl returns the ActionUrl field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetActionUrl() string {
	if o == nil || o.ActionUrl == nil {
		var ret string
		return ret
	}
	return *o.ActionUrl
}

// GetActionUrlOk returns a tuple with the ActionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetActionUrlOk() (*string, bool) {
	if o == nil || o.ActionUrl == nil {
		return nil, false
	}
	return o.ActionUrl, true
}

// HasActionUrl returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasActionUrl() bool {
	if o != nil && o.ActionUrl != nil {
		return true
	}

	return false
}

// SetActionUrl gets a reference to the given string and assigns it to the ActionUrl field.
func (o *BTAPIApplicationExtensionInfo) SetActionUrl(v string) {
	o.ActionUrl = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *BTAPIApplicationExtensionInfo) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BTAPIApplicationExtensionInfo) SetClientId(v string) {
	o.ClientId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTAPIApplicationExtensionInfo) SetDescription(v string) {
	o.Description = &v
}

// GetExtensionContext returns the ExtensionContext field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetExtensionContext() int64 {
	if o == nil || o.ExtensionContext == nil {
		var ret int64
		return ret
	}
	return *o.ExtensionContext
}

// GetExtensionContextOk returns a tuple with the ExtensionContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetExtensionContextOk() (*int64, bool) {
	if o == nil || o.ExtensionContext == nil {
		return nil, false
	}
	return o.ExtensionContext, true
}

// HasExtensionContext returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasExtensionContext() bool {
	if o != nil && o.ExtensionContext != nil {
		return true
	}

	return false
}

// SetExtensionContext gets a reference to the given int64 and assigns it to the ExtensionContext field.
func (o *BTAPIApplicationExtensionInfo) SetExtensionContext(v int64) {
	o.ExtensionContext = &v
}

// GetExtensionLocation returns the ExtensionLocation field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetExtensionLocation() int64 {
	if o == nil || o.ExtensionLocation == nil {
		var ret int64
		return ret
	}
	return *o.ExtensionLocation
}

// GetExtensionLocationOk returns a tuple with the ExtensionLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetExtensionLocationOk() (*int64, bool) {
	if o == nil || o.ExtensionLocation == nil {
		return nil, false
	}
	return o.ExtensionLocation, true
}

// HasExtensionLocation returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasExtensionLocation() bool {
	if o != nil && o.ExtensionLocation != nil {
		return true
	}

	return false
}

// SetExtensionLocation gets a reference to the given int64 and assigns it to the ExtensionLocation field.
func (o *BTAPIApplicationExtensionInfo) SetExtensionLocation(v int64) {
	o.ExtensionLocation = &v
}

// GetHasIcon returns the HasIcon field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetHasIcon() bool {
	if o == nil || o.HasIcon == nil {
		var ret bool
		return ret
	}
	return *o.HasIcon
}

// GetHasIconOk returns a tuple with the HasIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetHasIconOk() (*bool, bool) {
	if o == nil || o.HasIcon == nil {
		return nil, false
	}
	return o.HasIcon, true
}

// HasHasIcon returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasHasIcon() bool {
	if o != nil && o.HasIcon != nil {
		return true
	}

	return false
}

// SetHasIcon gets a reference to the given bool and assigns it to the HasIcon field.
func (o *BTAPIApplicationExtensionInfo) SetHasIcon(v bool) {
	o.HasIcon = &v
}

// GetHasPendingIcon returns the HasPendingIcon field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetHasPendingIcon() bool {
	if o == nil || o.HasPendingIcon == nil {
		var ret bool
		return ret
	}
	return *o.HasPendingIcon
}

// GetHasPendingIconOk returns a tuple with the HasPendingIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetHasPendingIconOk() (*bool, bool) {
	if o == nil || o.HasPendingIcon == nil {
		return nil, false
	}
	return o.HasPendingIcon, true
}

// HasHasPendingIcon returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasHasPendingIcon() bool {
	if o != nil && o.HasPendingIcon != nil {
		return true
	}

	return false
}

// SetHasPendingIcon gets a reference to the given bool and assigns it to the HasPendingIcon field.
func (o *BTAPIApplicationExtensionInfo) SetHasPendingIcon(v bool) {
	o.HasPendingIcon = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTAPIApplicationExtensionInfo) SetHref(v string) {
	o.Href = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetIconUrl() string {
	if o == nil || o.IconUrl == nil {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetIconUrlOk() (*string, bool) {
	if o == nil || o.IconUrl == nil {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasIconUrl() bool {
	if o != nil && o.IconUrl != nil {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *BTAPIApplicationExtensionInfo) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTAPIApplicationExtensionInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTAPIApplicationExtensionInfo) SetName(v string) {
	o.Name = &v
}

// GetParentAppPrimaryFormat returns the ParentAppPrimaryFormat field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetParentAppPrimaryFormat() string {
	if o == nil || o.ParentAppPrimaryFormat == nil {
		var ret string
		return ret
	}
	return *o.ParentAppPrimaryFormat
}

// GetParentAppPrimaryFormatOk returns a tuple with the ParentAppPrimaryFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetParentAppPrimaryFormatOk() (*string, bool) {
	if o == nil || o.ParentAppPrimaryFormat == nil {
		return nil, false
	}
	return o.ParentAppPrimaryFormat, true
}

// HasParentAppPrimaryFormat returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasParentAppPrimaryFormat() bool {
	if o != nil && o.ParentAppPrimaryFormat != nil {
		return true
	}

	return false
}

// SetParentAppPrimaryFormat gets a reference to the given string and assigns it to the ParentAppPrimaryFormat field.
func (o *BTAPIApplicationExtensionInfo) SetParentAppPrimaryFormat(v string) {
	o.ParentAppPrimaryFormat = &v
}

// GetShowResponse returns the ShowResponse field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetShowResponse() bool {
	if o == nil || o.ShowResponse == nil {
		var ret bool
		return ret
	}
	return *o.ShowResponse
}

// GetShowResponseOk returns a tuple with the ShowResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetShowResponseOk() (*bool, bool) {
	if o == nil || o.ShowResponse == nil {
		return nil, false
	}
	return o.ShowResponse, true
}

// HasShowResponse returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasShowResponse() bool {
	if o != nil && o.ShowResponse != nil {
		return true
	}

	return false
}

// SetShowResponse gets a reference to the given bool and assigns it to the ShowResponse field.
func (o *BTAPIApplicationExtensionInfo) SetShowResponse(v bool) {
	o.ShowResponse = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTAPIApplicationExtensionInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationExtensionInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTAPIApplicationExtensionInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTAPIApplicationExtensionInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTAPIApplicationExtensionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActionBody != nil {
		toSerialize["actionBody"] = o.ActionBody
	}
	if o.ActionType != nil {
		toSerialize["actionType"] = o.ActionType
	}
	if o.ActionUrl != nil {
		toSerialize["actionUrl"] = o.ActionUrl
	}
	if o.ApplicationId != nil {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExtensionContext != nil {
		toSerialize["extensionContext"] = o.ExtensionContext
	}
	if o.ExtensionLocation != nil {
		toSerialize["extensionLocation"] = o.ExtensionLocation
	}
	if o.HasIcon != nil {
		toSerialize["hasIcon"] = o.HasIcon
	}
	if o.HasPendingIcon != nil {
		toSerialize["hasPendingIcon"] = o.HasPendingIcon
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.IconUrl != nil {
		toSerialize["iconUrl"] = o.IconUrl
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ParentAppPrimaryFormat != nil {
		toSerialize["parentAppPrimaryFormat"] = o.ParentAppPrimaryFormat
	}
	if o.ShowResponse != nil {
		toSerialize["showResponse"] = o.ShowResponse
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTAPIApplicationExtensionInfo struct {
	value *BTAPIApplicationExtensionInfo
	isSet bool
}

func (v NullableBTAPIApplicationExtensionInfo) Get() *BTAPIApplicationExtensionInfo {
	return v.value
}

func (v *NullableBTAPIApplicationExtensionInfo) Set(val *BTAPIApplicationExtensionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAPIApplicationExtensionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAPIApplicationExtensionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAPIApplicationExtensionInfo(val *BTAPIApplicationExtensionInfo) *NullableBTAPIApplicationExtensionInfo {
	return &NullableBTAPIApplicationExtensionInfo{value: val, isSet: true}
}

func (v NullableBTAPIApplicationExtensionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAPIApplicationExtensionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
