/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.159.11153-01265f451c50
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAPIApplicationSummaryInfo struct for BTAPIApplicationSummaryInfo
type BTAPIApplicationSummaryInfo struct {
	ClientId    *string `json:"clientId,omitempty"`
	Description *string `json:"description,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// Name of the resource.
	Name  *string `json:"name,omitempty"`
	State *int32  `json:"state,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTAPIApplicationSummaryInfo instantiates a new BTAPIApplicationSummaryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAPIApplicationSummaryInfo() *BTAPIApplicationSummaryInfo {
	this := BTAPIApplicationSummaryInfo{}
	return &this
}

// NewBTAPIApplicationSummaryInfoWithDefaults instantiates a new BTAPIApplicationSummaryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAPIApplicationSummaryInfoWithDefaults() *BTAPIApplicationSummaryInfo {
	this := BTAPIApplicationSummaryInfo{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BTAPIApplicationSummaryInfo) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationSummaryInfo) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BTAPIApplicationSummaryInfo) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BTAPIApplicationSummaryInfo) SetClientId(v string) {
	o.ClientId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTAPIApplicationSummaryInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationSummaryInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTAPIApplicationSummaryInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTAPIApplicationSummaryInfo) SetDescription(v string) {
	o.Description = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTAPIApplicationSummaryInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationSummaryInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTAPIApplicationSummaryInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTAPIApplicationSummaryInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTAPIApplicationSummaryInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationSummaryInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTAPIApplicationSummaryInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTAPIApplicationSummaryInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTAPIApplicationSummaryInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationSummaryInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTAPIApplicationSummaryInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTAPIApplicationSummaryInfo) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTAPIApplicationSummaryInfo) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationSummaryInfo) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTAPIApplicationSummaryInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTAPIApplicationSummaryInfo) SetState(v int32) {
	o.State = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTAPIApplicationSummaryInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAPIApplicationSummaryInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTAPIApplicationSummaryInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTAPIApplicationSummaryInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTAPIApplicationSummaryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTAPIApplicationSummaryInfo struct {
	value *BTAPIApplicationSummaryInfo
	isSet bool
}

func (v NullableBTAPIApplicationSummaryInfo) Get() *BTAPIApplicationSummaryInfo {
	return v.value
}

func (v *NullableBTAPIApplicationSummaryInfo) Set(val *BTAPIApplicationSummaryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAPIApplicationSummaryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAPIApplicationSummaryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAPIApplicationSummaryInfo(val *BTAPIApplicationSummaryInfo) *NullableBTAPIApplicationSummaryInfo {
	return &NullableBTAPIApplicationSummaryInfo{value: val, isSet: true}
}

func (v NullableBTAPIApplicationSummaryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAPIApplicationSummaryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
