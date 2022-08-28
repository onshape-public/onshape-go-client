/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6251-33aac52a3d4e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTTeamMemberInfo struct for BTTeamMemberInfo
type BTTeamMemberInfo struct {
	Admin   *bool              `json:"admin,omitempty"`
	Href    *string            `json:"href,omitempty"`
	Id      *string            `json:"id,omitempty"`
	Member  *BTUserSummaryInfo `json:"member,omitempty"`
	Name    *string            `json:"name,omitempty"`
	Team    *BTTeamSummaryInfo `json:"team,omitempty"`
	ViewRef *string            `json:"viewRef,omitempty"`
}

// NewBTTeamMemberInfo instantiates a new BTTeamMemberInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTeamMemberInfo() *BTTeamMemberInfo {
	this := BTTeamMemberInfo{}
	return &this
}

// NewBTTeamMemberInfoWithDefaults instantiates a new BTTeamMemberInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTeamMemberInfoWithDefaults() *BTTeamMemberInfo {
	this := BTTeamMemberInfo{}
	return &this
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *BTTeamMemberInfo) GetAdmin() bool {
	if o == nil || o.Admin == nil {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTeamMemberInfo) GetAdminOk() (*bool, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *BTTeamMemberInfo) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *BTTeamMemberInfo) SetAdmin(v bool) {
	o.Admin = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTTeamMemberInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTeamMemberInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTTeamMemberInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTTeamMemberInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTTeamMemberInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTeamMemberInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTTeamMemberInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTTeamMemberInfo) SetId(v string) {
	o.Id = &v
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *BTTeamMemberInfo) GetMember() BTUserSummaryInfo {
	if o == nil || o.Member == nil {
		var ret BTUserSummaryInfo
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTeamMemberInfo) GetMemberOk() (*BTUserSummaryInfo, bool) {
	if o == nil || o.Member == nil {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *BTTeamMemberInfo) HasMember() bool {
	if o != nil && o.Member != nil {
		return true
	}

	return false
}

// SetMember gets a reference to the given BTUserSummaryInfo and assigns it to the Member field.
func (o *BTTeamMemberInfo) SetMember(v BTUserSummaryInfo) {
	o.Member = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTTeamMemberInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTeamMemberInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTTeamMemberInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTTeamMemberInfo) SetName(v string) {
	o.Name = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *BTTeamMemberInfo) GetTeam() BTTeamSummaryInfo {
	if o == nil || o.Team == nil {
		var ret BTTeamSummaryInfo
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTeamMemberInfo) GetTeamOk() (*BTTeamSummaryInfo, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *BTTeamMemberInfo) HasTeam() bool {
	if o != nil && o.Team != nil {
		return true
	}

	return false
}

// SetTeam gets a reference to the given BTTeamSummaryInfo and assigns it to the Team field.
func (o *BTTeamMemberInfo) SetTeam(v BTTeamSummaryInfo) {
	o.Team = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTTeamMemberInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTeamMemberInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTTeamMemberInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTTeamMemberInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTTeamMemberInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Member != nil {
		toSerialize["member"] = o.Member
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Team != nil {
		toSerialize["team"] = o.Team
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTTeamMemberInfo struct {
	value *BTTeamMemberInfo
	isSet bool
}

func (v NullableBTTeamMemberInfo) Get() *BTTeamMemberInfo {
	return v.value
}

func (v *NullableBTTeamMemberInfo) Set(val *BTTeamMemberInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTeamMemberInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTeamMemberInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTeamMemberInfo(val *BTTeamMemberInfo) *NullableBTTeamMemberInfo {
	return &NullableBTTeamMemberInfo{value: val, isSet: true}
}

func (v NullableBTTeamMemberInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTeamMemberInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
