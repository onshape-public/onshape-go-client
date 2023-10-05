/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23367-59c3c9f9feef
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTSearchEntityType the model 'BTSearchEntityType'
type BTSearchEntityType string

// List of BTSearchEntityType
const (
	BTSearchEntityTypeUnknown     BTSearchEntityType = "unknown"
	BTSearchEntityTypeCapability  BTSearchEntityType = "capability"
	BTSearchEntityTypeCompany     BTSearchEntityType = "company"
	BTSearchEntityTypeDocument    BTSearchEntityType = "document"
	BTSearchEntityTypeElement     BTSearchEntityType = "element"
	BTSearchEntityTypeFriend      BTSearchEntityType = "friend"
	BTSearchEntityTypeItem        BTSearchEntityType = "item"
	BTSearchEntityTypePart        BTSearchEntityType = "part"
	BTSearchEntityTypeTeam        BTSearchEntityType = "team"
	BTSearchEntityTypeUser        BTSearchEntityType = "user"
	BTSearchEntityTypeVersion     BTSearchEntityType = "version"
	BTSearchEntityTypeWorkspace   BTSearchEntityType = "workspace"
	BTSearchEntityTypeProject     BTSearchEntityType = "project"
	BTSearchEntityTypePublication BTSearchEntityType = "publication"
)

// All allowed values of BTSearchEntityType enum
var AllowedBTSearchEntityTypeEnumValues = []BTSearchEntityType{
	"unknown",
	"capability",
	"company",
	"document",
	"element",
	"friend",
	"item",
	"part",
	"team",
	"user",
	"version",
	"workspace",
	"project",
	"publication",
}

func (v *BTSearchEntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTSearchEntityType(value)
	for _, existing := range AllowedBTSearchEntityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTSearchEntityType", value)
}

// NewBTSearchEntityTypeFromValue returns a pointer to a valid BTSearchEntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTSearchEntityTypeFromValue(v string) (*BTSearchEntityType, error) {
	ev := BTSearchEntityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTSearchEntityType: valid values are %v", v, AllowedBTSearchEntityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTSearchEntityType) IsValid() bool {
	for _, existing := range AllowedBTSearchEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTSearchEntityType value
func (v BTSearchEntityType) Ptr() *BTSearchEntityType {
	return &v
}

type NullableBTSearchEntityType struct {
	value *BTSearchEntityType
	isSet bool
}

func (v NullableBTSearchEntityType) Get() *BTSearchEntityType {
	return v.value
}

func (v *NullableBTSearchEntityType) Set(val *BTSearchEntityType) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSearchEntityType) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSearchEntityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSearchEntityType(val *BTSearchEntityType) *NullableBTSearchEntityType {
	return &NullableBTSearchEntityType{value: val, isSet: true}
}

func (v NullableBTSearchEntityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSearchEntityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
