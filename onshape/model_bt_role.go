/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.22862-4427d042758b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTRole the model 'BTRole'
type BTRole string

// List of BTRole
const (
	BTRoleAnonymous          BTRole = "ANONYMOUS"
	BTRoleTotppendinguser    BTRole = "TOTPPENDINGUSER"
	BTRoleUser               BTRole = "USER"
	BTRoleDeveloper          BTRole = "DEVELOPER"
	BTRolePartner            BTRole = "PARTNER"
	BTRoleOnshapecompanyuser BTRole = "ONSHAPECOMPANYUSER"
	BTRoleAdmin              BTRole = "ADMIN"
	BTRoleService            BTRole = "SERVICE"
)

// All allowed values of BTRole enum
var AllowedBTRoleEnumValues = []BTRole{
	"ANONYMOUS",
	"TOTPPENDINGUSER",
	"USER",
	"DEVELOPER",
	"PARTNER",
	"ONSHAPECOMPANYUSER",
	"ADMIN",
	"SERVICE",
}

func (v *BTRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTRole(value)
	for _, existing := range AllowedBTRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTRole", value)
}

// NewBTRoleFromValue returns a pointer to a valid BTRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTRoleFromValue(v string) (*BTRole, error) {
	ev := BTRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTRole: valid values are %v", v, AllowedBTRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTRole) IsValid() bool {
	for _, existing := range AllowedBTRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTRole value
func (v BTRole) Ptr() *BTRole {
	return &v
}

type NullableBTRole struct {
	value *BTRole
	isSet bool
}

func (v NullableBTRole) Get() *BTRole {
	return v.value
}

func (v *NullableBTRole) Set(val *BTRole) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRole) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRole(val *BTRole) *NullableBTRole {
	return &NullableBTRole{value: val, isSet: true}
}

func (v NullableBTRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
