/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5768-0013f50d23d2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTParameterGroupSpec3469 struct for BTParameterGroupSpec3469
type BTParameterGroupSpec3469 struct {
	AdditionalLocalizedStrings *int32   `json:"additionalLocalizedStrings,omitempty"`
	CollapsedByDefault         *bool    `json:"collapsedByDefault,omitempty"`
	DrivingParameterId         *string  `json:"drivingParameterId,omitempty"`
	GroupId                    *string  `json:"groupId,omitempty"`
	GroupName                  *string  `json:"groupName,omitempty"`
	GroupOrParameterIds        []string `json:"groupOrParameterIds,omitempty"`
	LocalizableName            *string  `json:"localizableName,omitempty"`
	LocalizedName              *string  `json:"localizedName,omitempty"`
	StringsToLocalize          []string `json:"stringsToLocalize,omitempty"`
}

// NewBTParameterGroupSpec3469 instantiates a new BTParameterGroupSpec3469 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterGroupSpec3469() *BTParameterGroupSpec3469 {
	this := BTParameterGroupSpec3469{}
	return &this
}

// NewBTParameterGroupSpec3469WithDefaults instantiates a new BTParameterGroupSpec3469 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterGroupSpec3469WithDefaults() *BTParameterGroupSpec3469 {
	this := BTParameterGroupSpec3469{}
	return &this
}

// GetAdditionalLocalizedStrings returns the AdditionalLocalizedStrings field value if set, zero value otherwise.
func (o *BTParameterGroupSpec3469) GetAdditionalLocalizedStrings() int32 {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		var ret int32
		return ret
	}
	return *o.AdditionalLocalizedStrings
}

// GetAdditionalLocalizedStringsOk returns a tuple with the AdditionalLocalizedStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterGroupSpec3469) GetAdditionalLocalizedStringsOk() (*int32, bool) {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		return nil, false
	}
	return o.AdditionalLocalizedStrings, true
}

// HasAdditionalLocalizedStrings returns a boolean if a field has been set.
func (o *BTParameterGroupSpec3469) HasAdditionalLocalizedStrings() bool {
	if o != nil && o.AdditionalLocalizedStrings != nil {
		return true
	}

	return false
}

// SetAdditionalLocalizedStrings gets a reference to the given int32 and assigns it to the AdditionalLocalizedStrings field.
func (o *BTParameterGroupSpec3469) SetAdditionalLocalizedStrings(v int32) {
	o.AdditionalLocalizedStrings = &v
}

// GetCollapsedByDefault returns the CollapsedByDefault field value if set, zero value otherwise.
func (o *BTParameterGroupSpec3469) GetCollapsedByDefault() bool {
	if o == nil || o.CollapsedByDefault == nil {
		var ret bool
		return ret
	}
	return *o.CollapsedByDefault
}

// GetCollapsedByDefaultOk returns a tuple with the CollapsedByDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterGroupSpec3469) GetCollapsedByDefaultOk() (*bool, bool) {
	if o == nil || o.CollapsedByDefault == nil {
		return nil, false
	}
	return o.CollapsedByDefault, true
}

// HasCollapsedByDefault returns a boolean if a field has been set.
func (o *BTParameterGroupSpec3469) HasCollapsedByDefault() bool {
	if o != nil && o.CollapsedByDefault != nil {
		return true
	}

	return false
}

// SetCollapsedByDefault gets a reference to the given bool and assigns it to the CollapsedByDefault field.
func (o *BTParameterGroupSpec3469) SetCollapsedByDefault(v bool) {
	o.CollapsedByDefault = &v
}

// GetDrivingParameterId returns the DrivingParameterId field value if set, zero value otherwise.
func (o *BTParameterGroupSpec3469) GetDrivingParameterId() string {
	if o == nil || o.DrivingParameterId == nil {
		var ret string
		return ret
	}
	return *o.DrivingParameterId
}

// GetDrivingParameterIdOk returns a tuple with the DrivingParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterGroupSpec3469) GetDrivingParameterIdOk() (*string, bool) {
	if o == nil || o.DrivingParameterId == nil {
		return nil, false
	}
	return o.DrivingParameterId, true
}

// HasDrivingParameterId returns a boolean if a field has been set.
func (o *BTParameterGroupSpec3469) HasDrivingParameterId() bool {
	if o != nil && o.DrivingParameterId != nil {
		return true
	}

	return false
}

// SetDrivingParameterId gets a reference to the given string and assigns it to the DrivingParameterId field.
func (o *BTParameterGroupSpec3469) SetDrivingParameterId(v string) {
	o.DrivingParameterId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *BTParameterGroupSpec3469) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterGroupSpec3469) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *BTParameterGroupSpec3469) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *BTParameterGroupSpec3469) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *BTParameterGroupSpec3469) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterGroupSpec3469) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *BTParameterGroupSpec3469) HasGroupName() bool {
	if o != nil && o.GroupName != nil {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *BTParameterGroupSpec3469) SetGroupName(v string) {
	o.GroupName = &v
}

// GetGroupOrParameterIds returns the GroupOrParameterIds field value if set, zero value otherwise.
func (o *BTParameterGroupSpec3469) GetGroupOrParameterIds() []string {
	if o == nil || o.GroupOrParameterIds == nil {
		var ret []string
		return ret
	}
	return o.GroupOrParameterIds
}

// GetGroupOrParameterIdsOk returns a tuple with the GroupOrParameterIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterGroupSpec3469) GetGroupOrParameterIdsOk() ([]string, bool) {
	if o == nil || o.GroupOrParameterIds == nil {
		return nil, false
	}
	return o.GroupOrParameterIds, true
}

// HasGroupOrParameterIds returns a boolean if a field has been set.
func (o *BTParameterGroupSpec3469) HasGroupOrParameterIds() bool {
	if o != nil && o.GroupOrParameterIds != nil {
		return true
	}

	return false
}

// SetGroupOrParameterIds gets a reference to the given []string and assigns it to the GroupOrParameterIds field.
func (o *BTParameterGroupSpec3469) SetGroupOrParameterIds(v []string) {
	o.GroupOrParameterIds = v
}

// GetLocalizableName returns the LocalizableName field value if set, zero value otherwise.
func (o *BTParameterGroupSpec3469) GetLocalizableName() string {
	if o == nil || o.LocalizableName == nil {
		var ret string
		return ret
	}
	return *o.LocalizableName
}

// GetLocalizableNameOk returns a tuple with the LocalizableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterGroupSpec3469) GetLocalizableNameOk() (*string, bool) {
	if o == nil || o.LocalizableName == nil {
		return nil, false
	}
	return o.LocalizableName, true
}

// HasLocalizableName returns a boolean if a field has been set.
func (o *BTParameterGroupSpec3469) HasLocalizableName() bool {
	if o != nil && o.LocalizableName != nil {
		return true
	}

	return false
}

// SetLocalizableName gets a reference to the given string and assigns it to the LocalizableName field.
func (o *BTParameterGroupSpec3469) SetLocalizableName(v string) {
	o.LocalizableName = &v
}

// GetLocalizedName returns the LocalizedName field value if set, zero value otherwise.
func (o *BTParameterGroupSpec3469) GetLocalizedName() string {
	if o == nil || o.LocalizedName == nil {
		var ret string
		return ret
	}
	return *o.LocalizedName
}

// GetLocalizedNameOk returns a tuple with the LocalizedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterGroupSpec3469) GetLocalizedNameOk() (*string, bool) {
	if o == nil || o.LocalizedName == nil {
		return nil, false
	}
	return o.LocalizedName, true
}

// HasLocalizedName returns a boolean if a field has been set.
func (o *BTParameterGroupSpec3469) HasLocalizedName() bool {
	if o != nil && o.LocalizedName != nil {
		return true
	}

	return false
}

// SetLocalizedName gets a reference to the given string and assigns it to the LocalizedName field.
func (o *BTParameterGroupSpec3469) SetLocalizedName(v string) {
	o.LocalizedName = &v
}

// GetStringsToLocalize returns the StringsToLocalize field value if set, zero value otherwise.
func (o *BTParameterGroupSpec3469) GetStringsToLocalize() []string {
	if o == nil || o.StringsToLocalize == nil {
		var ret []string
		return ret
	}
	return o.StringsToLocalize
}

// GetStringsToLocalizeOk returns a tuple with the StringsToLocalize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterGroupSpec3469) GetStringsToLocalizeOk() ([]string, bool) {
	if o == nil || o.StringsToLocalize == nil {
		return nil, false
	}
	return o.StringsToLocalize, true
}

// HasStringsToLocalize returns a boolean if a field has been set.
func (o *BTParameterGroupSpec3469) HasStringsToLocalize() bool {
	if o != nil && o.StringsToLocalize != nil {
		return true
	}

	return false
}

// SetStringsToLocalize gets a reference to the given []string and assigns it to the StringsToLocalize field.
func (o *BTParameterGroupSpec3469) SetStringsToLocalize(v []string) {
	o.StringsToLocalize = v
}

func (o BTParameterGroupSpec3469) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalLocalizedStrings != nil {
		toSerialize["additionalLocalizedStrings"] = o.AdditionalLocalizedStrings
	}
	if o.CollapsedByDefault != nil {
		toSerialize["collapsedByDefault"] = o.CollapsedByDefault
	}
	if o.DrivingParameterId != nil {
		toSerialize["drivingParameterId"] = o.DrivingParameterId
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.GroupName != nil {
		toSerialize["groupName"] = o.GroupName
	}
	if o.GroupOrParameterIds != nil {
		toSerialize["groupOrParameterIds"] = o.GroupOrParameterIds
	}
	if o.LocalizableName != nil {
		toSerialize["localizableName"] = o.LocalizableName
	}
	if o.LocalizedName != nil {
		toSerialize["localizedName"] = o.LocalizedName
	}
	if o.StringsToLocalize != nil {
		toSerialize["stringsToLocalize"] = o.StringsToLocalize
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterGroupSpec3469 struct {
	value *BTParameterGroupSpec3469
	isSet bool
}

func (v NullableBTParameterGroupSpec3469) Get() *BTParameterGroupSpec3469 {
	return v.value
}

func (v *NullableBTParameterGroupSpec3469) Set(val *BTParameterGroupSpec3469) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterGroupSpec3469) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterGroupSpec3469) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterGroupSpec3469(val *BTParameterGroupSpec3469) *NullableBTParameterGroupSpec3469 {
	return &NullableBTParameterGroupSpec3469{value: val, isSet: true}
}

func (v NullableBTParameterGroupSpec3469) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterGroupSpec3469) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
