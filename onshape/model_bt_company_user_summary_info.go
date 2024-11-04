/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCompanyUserSummaryInfo struct for BTCompanyUserSummaryInfo
type BTCompanyUserSummaryInfo struct {
	JsonType string `json:"jsonType"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// Name of the resource.
	Name *string `json:"name,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef                   *string               `json:"viewRef,omitempty"`
	Image                     *string               `json:"image,omitempty"`
	IsOnshapeSupport          *bool                 `json:"isOnshapeSupport,omitempty"`
	State                     *int32                `json:"state,omitempty"`
	DocumentationName         *string               `json:"documentationName,omitempty"`
	Email                     *string               `json:"email,omitempty"`
	FirstName                 *string               `json:"firstName,omitempty"`
	LastName                  *string               `json:"lastName,omitempty"`
	Company                   *BTCompanySummaryInfo `json:"company,omitempty"`
	DocumentationNameOverride *string               `json:"documentationNameOverride,omitempty"`
	GlobalPermissions         *GlobalPermissionInfo `json:"globalPermissions,omitempty"`
	InvitationState           *int32                `json:"invitationState,omitempty"`
	IsGuest                   *bool                 `json:"isGuest,omitempty"`
	IsLight                   *bool                 `json:"isLight,omitempty"`
	LastLoginTime             *JSONTime             `json:"lastLoginTime,omitempty"`
	PersonalMessageAllowed    *bool                 `json:"personalMessageAllowed,omitempty"`
	Source                    *int32                `json:"source,omitempty"`
	Admin                     *bool                 `json:"admin,omitempty"`
	Cls                       *string               `json:"cls,omitempty"`
}

// NewBTCompanyUserSummaryInfo instantiates a new BTCompanyUserSummaryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCompanyUserSummaryInfo(jsonType string) *BTCompanyUserSummaryInfo {
	this := BTCompanyUserSummaryInfo{}
	this.JsonType = jsonType
	return &this
}

// NewBTCompanyUserSummaryInfoWithDefaults instantiates a new BTCompanyUserSummaryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCompanyUserSummaryInfoWithDefaults() *BTCompanyUserSummaryInfo {
	this := BTCompanyUserSummaryInfo{}
	return &this
}

// GetJsonType returns the JsonType field value
func (o *BTCompanyUserSummaryInfo) GetJsonType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetJsonTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonType, true
}

// SetJsonType sets field value
func (o *BTCompanyUserSummaryInfo) SetJsonType(v string) {
	o.JsonType = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTCompanyUserSummaryInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTCompanyUserSummaryInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTCompanyUserSummaryInfo) SetName(v string) {
	o.Name = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTCompanyUserSummaryInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BTCompanyUserSummaryInfo) SetImage(v string) {
	o.Image = &v
}

// GetIsOnshapeSupport returns the IsOnshapeSupport field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetIsOnshapeSupport() bool {
	if o == nil || o.IsOnshapeSupport == nil {
		var ret bool
		return ret
	}
	return *o.IsOnshapeSupport
}

// GetIsOnshapeSupportOk returns a tuple with the IsOnshapeSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetIsOnshapeSupportOk() (*bool, bool) {
	if o == nil || o.IsOnshapeSupport == nil {
		return nil, false
	}
	return o.IsOnshapeSupport, true
}

// HasIsOnshapeSupport returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasIsOnshapeSupport() bool {
	if o != nil && o.IsOnshapeSupport != nil {
		return true
	}

	return false
}

// SetIsOnshapeSupport gets a reference to the given bool and assigns it to the IsOnshapeSupport field.
func (o *BTCompanyUserSummaryInfo) SetIsOnshapeSupport(v bool) {
	o.IsOnshapeSupport = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTCompanyUserSummaryInfo) SetState(v int32) {
	o.State = &v
}

// GetDocumentationName returns the DocumentationName field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetDocumentationName() string {
	if o == nil || o.DocumentationName == nil {
		var ret string
		return ret
	}
	return *o.DocumentationName
}

// GetDocumentationNameOk returns a tuple with the DocumentationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetDocumentationNameOk() (*string, bool) {
	if o == nil || o.DocumentationName == nil {
		return nil, false
	}
	return o.DocumentationName, true
}

// HasDocumentationName returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasDocumentationName() bool {
	if o != nil && o.DocumentationName != nil {
		return true
	}

	return false
}

// SetDocumentationName gets a reference to the given string and assigns it to the DocumentationName field.
func (o *BTCompanyUserSummaryInfo) SetDocumentationName(v string) {
	o.DocumentationName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BTCompanyUserSummaryInfo) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BTCompanyUserSummaryInfo) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BTCompanyUserSummaryInfo) SetLastName(v string) {
	o.LastName = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetCompany() BTCompanySummaryInfo {
	if o == nil || o.Company == nil {
		var ret BTCompanySummaryInfo
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetCompanyOk() (*BTCompanySummaryInfo, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given BTCompanySummaryInfo and assigns it to the Company field.
func (o *BTCompanyUserSummaryInfo) SetCompany(v BTCompanySummaryInfo) {
	o.Company = &v
}

// GetDocumentationNameOverride returns the DocumentationNameOverride field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetDocumentationNameOverride() string {
	if o == nil || o.DocumentationNameOverride == nil {
		var ret string
		return ret
	}
	return *o.DocumentationNameOverride
}

// GetDocumentationNameOverrideOk returns a tuple with the DocumentationNameOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetDocumentationNameOverrideOk() (*string, bool) {
	if o == nil || o.DocumentationNameOverride == nil {
		return nil, false
	}
	return o.DocumentationNameOverride, true
}

// HasDocumentationNameOverride returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasDocumentationNameOverride() bool {
	if o != nil && o.DocumentationNameOverride != nil {
		return true
	}

	return false
}

// SetDocumentationNameOverride gets a reference to the given string and assigns it to the DocumentationNameOverride field.
func (o *BTCompanyUserSummaryInfo) SetDocumentationNameOverride(v string) {
	o.DocumentationNameOverride = &v
}

// GetGlobalPermissions returns the GlobalPermissions field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetGlobalPermissions() GlobalPermissionInfo {
	if o == nil || o.GlobalPermissions == nil {
		var ret GlobalPermissionInfo
		return ret
	}
	return *o.GlobalPermissions
}

// GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetGlobalPermissionsOk() (*GlobalPermissionInfo, bool) {
	if o == nil || o.GlobalPermissions == nil {
		return nil, false
	}
	return o.GlobalPermissions, true
}

// HasGlobalPermissions returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasGlobalPermissions() bool {
	if o != nil && o.GlobalPermissions != nil {
		return true
	}

	return false
}

// SetGlobalPermissions gets a reference to the given GlobalPermissionInfo and assigns it to the GlobalPermissions field.
func (o *BTCompanyUserSummaryInfo) SetGlobalPermissions(v GlobalPermissionInfo) {
	o.GlobalPermissions = &v
}

// GetInvitationState returns the InvitationState field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetInvitationState() int32 {
	if o == nil || o.InvitationState == nil {
		var ret int32
		return ret
	}
	return *o.InvitationState
}

// GetInvitationStateOk returns a tuple with the InvitationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetInvitationStateOk() (*int32, bool) {
	if o == nil || o.InvitationState == nil {
		return nil, false
	}
	return o.InvitationState, true
}

// HasInvitationState returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasInvitationState() bool {
	if o != nil && o.InvitationState != nil {
		return true
	}

	return false
}

// SetInvitationState gets a reference to the given int32 and assigns it to the InvitationState field.
func (o *BTCompanyUserSummaryInfo) SetInvitationState(v int32) {
	o.InvitationState = &v
}

// GetIsGuest returns the IsGuest field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetIsGuest() bool {
	if o == nil || o.IsGuest == nil {
		var ret bool
		return ret
	}
	return *o.IsGuest
}

// GetIsGuestOk returns a tuple with the IsGuest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetIsGuestOk() (*bool, bool) {
	if o == nil || o.IsGuest == nil {
		return nil, false
	}
	return o.IsGuest, true
}

// HasIsGuest returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasIsGuest() bool {
	if o != nil && o.IsGuest != nil {
		return true
	}

	return false
}

// SetIsGuest gets a reference to the given bool and assigns it to the IsGuest field.
func (o *BTCompanyUserSummaryInfo) SetIsGuest(v bool) {
	o.IsGuest = &v
}

// GetIsLight returns the IsLight field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetIsLight() bool {
	if o == nil || o.IsLight == nil {
		var ret bool
		return ret
	}
	return *o.IsLight
}

// GetIsLightOk returns a tuple with the IsLight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetIsLightOk() (*bool, bool) {
	if o == nil || o.IsLight == nil {
		return nil, false
	}
	return o.IsLight, true
}

// HasIsLight returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasIsLight() bool {
	if o != nil && o.IsLight != nil {
		return true
	}

	return false
}

// SetIsLight gets a reference to the given bool and assigns it to the IsLight field.
func (o *BTCompanyUserSummaryInfo) SetIsLight(v bool) {
	o.IsLight = &v
}

// GetLastLoginTime returns the LastLoginTime field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetLastLoginTime() JSONTime {
	if o == nil || o.LastLoginTime == nil {
		var ret JSONTime
		return ret
	}
	return *o.LastLoginTime
}

// GetLastLoginTimeOk returns a tuple with the LastLoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetLastLoginTimeOk() (*JSONTime, bool) {
	if o == nil || o.LastLoginTime == nil {
		return nil, false
	}
	return o.LastLoginTime, true
}

// HasLastLoginTime returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasLastLoginTime() bool {
	if o != nil && o.LastLoginTime != nil {
		return true
	}

	return false
}

// SetLastLoginTime gets a reference to the given JSONTime and assigns it to the LastLoginTime field.
func (o *BTCompanyUserSummaryInfo) SetLastLoginTime(v JSONTime) {
	o.LastLoginTime = &v
}

// GetPersonalMessageAllowed returns the PersonalMessageAllowed field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetPersonalMessageAllowed() bool {
	if o == nil || o.PersonalMessageAllowed == nil {
		var ret bool
		return ret
	}
	return *o.PersonalMessageAllowed
}

// GetPersonalMessageAllowedOk returns a tuple with the PersonalMessageAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetPersonalMessageAllowedOk() (*bool, bool) {
	if o == nil || o.PersonalMessageAllowed == nil {
		return nil, false
	}
	return o.PersonalMessageAllowed, true
}

// HasPersonalMessageAllowed returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasPersonalMessageAllowed() bool {
	if o != nil && o.PersonalMessageAllowed != nil {
		return true
	}

	return false
}

// SetPersonalMessageAllowed gets a reference to the given bool and assigns it to the PersonalMessageAllowed field.
func (o *BTCompanyUserSummaryInfo) SetPersonalMessageAllowed(v bool) {
	o.PersonalMessageAllowed = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetSource() int32 {
	if o == nil || o.Source == nil {
		var ret int32
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetSourceOk() (*int32, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given int32 and assigns it to the Source field.
func (o *BTCompanyUserSummaryInfo) SetSource(v int32) {
	o.Source = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetAdmin() bool {
	if o == nil || o.Admin == nil {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetAdminOk() (*bool, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *BTCompanyUserSummaryInfo) SetAdmin(v bool) {
	o.Admin = &v
}

// GetCls returns the Cls field value if set, zero value otherwise.
func (o *BTCompanyUserSummaryInfo) GetCls() string {
	if o == nil || o.Cls == nil {
		var ret string
		return ret
	}
	return *o.Cls
}

// GetClsOk returns a tuple with the Cls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserSummaryInfo) GetClsOk() (*string, bool) {
	if o == nil || o.Cls == nil {
		return nil, false
	}
	return o.Cls, true
}

// HasCls returns a boolean if a field has been set.
func (o *BTCompanyUserSummaryInfo) HasCls() bool {
	if o != nil && o.Cls != nil {
		return true
	}

	return false
}

// SetCls gets a reference to the given string and assigns it to the Cls field.
func (o *BTCompanyUserSummaryInfo) SetCls(v string) {
	o.Cls = &v
}

func (o BTCompanyUserSummaryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["jsonType"] = o.JsonType
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
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.IsOnshapeSupport != nil {
		toSerialize["isOnshapeSupport"] = o.IsOnshapeSupport
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.DocumentationName != nil {
		toSerialize["documentationName"] = o.DocumentationName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.DocumentationNameOverride != nil {
		toSerialize["documentationNameOverride"] = o.DocumentationNameOverride
	}
	if o.GlobalPermissions != nil {
		toSerialize["globalPermissions"] = o.GlobalPermissions
	}
	if o.InvitationState != nil {
		toSerialize["invitationState"] = o.InvitationState
	}
	if o.IsGuest != nil {
		toSerialize["isGuest"] = o.IsGuest
	}
	if o.IsLight != nil {
		toSerialize["isLight"] = o.IsLight
	}
	if o.LastLoginTime != nil {
		toSerialize["lastLoginTime"] = o.LastLoginTime
	}
	if o.PersonalMessageAllowed != nil {
		toSerialize["personalMessageAllowed"] = o.PersonalMessageAllowed
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	if o.Cls != nil {
		toSerialize["cls"] = o.Cls
	}
	return json.Marshal(toSerialize)
}

type NullableBTCompanyUserSummaryInfo struct {
	value *BTCompanyUserSummaryInfo
	isSet bool
}

func (v NullableBTCompanyUserSummaryInfo) Get() *BTCompanyUserSummaryInfo {
	return v.value
}

func (v *NullableBTCompanyUserSummaryInfo) Set(val *BTCompanyUserSummaryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCompanyUserSummaryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCompanyUserSummaryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCompanyUserSummaryInfo(val *BTCompanyUserSummaryInfo) *NullableBTCompanyUserSummaryInfo {
	return &NullableBTCompanyUserSummaryInfo{value: val, isSet: true}
}

func (v NullableBTCompanyUserSummaryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCompanyUserSummaryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
