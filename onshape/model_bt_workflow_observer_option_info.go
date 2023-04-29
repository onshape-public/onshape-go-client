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

// BTWorkflowObserverOptionInfo Array of items in the current page.
type BTWorkflowObserverOptionInfo struct {
	Alias *BTAliasInfo `json:"alias,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id           *string            `json:"id,omitempty"`
	IdentityType *int32             `json:"identityType,omitempty"`
	Role         *BTRbacRoleInfo    `json:"role,omitempty"`
	Team         *BTTeamSummaryInfo `json:"team,omitempty"`
	User         *BTUserSummaryInfo `json:"user,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTWorkflowObserverOptionInfo instantiates a new BTWorkflowObserverOptionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWorkflowObserverOptionInfo() *BTWorkflowObserverOptionInfo {
	this := BTWorkflowObserverOptionInfo{}
	return &this
}

// NewBTWorkflowObserverOptionInfoWithDefaults instantiates a new BTWorkflowObserverOptionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWorkflowObserverOptionInfoWithDefaults() *BTWorkflowObserverOptionInfo {
	this := BTWorkflowObserverOptionInfo{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *BTWorkflowObserverOptionInfo) GetAlias() BTAliasInfo {
	if o == nil || o.Alias == nil {
		var ret BTAliasInfo
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowObserverOptionInfo) GetAliasOk() (*BTAliasInfo, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *BTWorkflowObserverOptionInfo) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given BTAliasInfo and assigns it to the Alias field.
func (o *BTWorkflowObserverOptionInfo) SetAlias(v BTAliasInfo) {
	o.Alias = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTWorkflowObserverOptionInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowObserverOptionInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTWorkflowObserverOptionInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTWorkflowObserverOptionInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTWorkflowObserverOptionInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowObserverOptionInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTWorkflowObserverOptionInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTWorkflowObserverOptionInfo) SetId(v string) {
	o.Id = &v
}

// GetIdentityType returns the IdentityType field value if set, zero value otherwise.
func (o *BTWorkflowObserverOptionInfo) GetIdentityType() int32 {
	if o == nil || o.IdentityType == nil {
		var ret int32
		return ret
	}
	return *o.IdentityType
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowObserverOptionInfo) GetIdentityTypeOk() (*int32, bool) {
	if o == nil || o.IdentityType == nil {
		return nil, false
	}
	return o.IdentityType, true
}

// HasIdentityType returns a boolean if a field has been set.
func (o *BTWorkflowObserverOptionInfo) HasIdentityType() bool {
	if o != nil && o.IdentityType != nil {
		return true
	}

	return false
}

// SetIdentityType gets a reference to the given int32 and assigns it to the IdentityType field.
func (o *BTWorkflowObserverOptionInfo) SetIdentityType(v int32) {
	o.IdentityType = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *BTWorkflowObserverOptionInfo) GetRole() BTRbacRoleInfo {
	if o == nil || o.Role == nil {
		var ret BTRbacRoleInfo
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowObserverOptionInfo) GetRoleOk() (*BTRbacRoleInfo, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *BTWorkflowObserverOptionInfo) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given BTRbacRoleInfo and assigns it to the Role field.
func (o *BTWorkflowObserverOptionInfo) SetRole(v BTRbacRoleInfo) {
	o.Role = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *BTWorkflowObserverOptionInfo) GetTeam() BTTeamSummaryInfo {
	if o == nil || o.Team == nil {
		var ret BTTeamSummaryInfo
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowObserverOptionInfo) GetTeamOk() (*BTTeamSummaryInfo, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *BTWorkflowObserverOptionInfo) HasTeam() bool {
	if o != nil && o.Team != nil {
		return true
	}

	return false
}

// SetTeam gets a reference to the given BTTeamSummaryInfo and assigns it to the Team field.
func (o *BTWorkflowObserverOptionInfo) SetTeam(v BTTeamSummaryInfo) {
	o.Team = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *BTWorkflowObserverOptionInfo) GetUser() BTUserSummaryInfo {
	if o == nil || o.User == nil {
		var ret BTUserSummaryInfo
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowObserverOptionInfo) GetUserOk() (*BTUserSummaryInfo, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *BTWorkflowObserverOptionInfo) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given BTUserSummaryInfo and assigns it to the User field.
func (o *BTWorkflowObserverOptionInfo) SetUser(v BTUserSummaryInfo) {
	o.User = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTWorkflowObserverOptionInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowObserverOptionInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTWorkflowObserverOptionInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTWorkflowObserverOptionInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTWorkflowObserverOptionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdentityType != nil {
		toSerialize["identityType"] = o.IdentityType
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Team != nil {
		toSerialize["team"] = o.Team
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTWorkflowObserverOptionInfo struct {
	value *BTWorkflowObserverOptionInfo
	isSet bool
}

func (v NullableBTWorkflowObserverOptionInfo) Get() *BTWorkflowObserverOptionInfo {
	return v.value
}

func (v *NullableBTWorkflowObserverOptionInfo) Set(val *BTWorkflowObserverOptionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkflowObserverOptionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkflowObserverOptionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkflowObserverOptionInfo(val *BTWorkflowObserverOptionInfo) *NullableBTWorkflowObserverOptionInfo {
	return &NullableBTWorkflowObserverOptionInfo{value: val, isSet: true}
}

func (v NullableBTWorkflowObserverOptionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkflowObserverOptionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
