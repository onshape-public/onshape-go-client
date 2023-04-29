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
	"fmt"
)

// EditType the model 'EditType'
type EditType string

// List of EditType
const (
	NOTHING      EditType = "NOTHING"
	NEW_ROOT     EditType = "NEW_ROOT"
	MOVE         EditType = "MOVE"
	CHANGE       EditType = "CHANGE"
	CHANGE_FIELD EditType = "CHANGE_FIELD"
	INSERTION    EditType = "INSERTION"
	DELETION     EditType = "DELETION"
	LIST         EditType = "LIST"
)

// All allowed values of EditType enum
var AllowedEditTypeEnumValues = []EditType{
	"NOTHING",
	"NEW_ROOT",
	"MOVE",
	"CHANGE",
	"CHANGE_FIELD",
	"INSERTION",
	"DELETION",
	"LIST",
}

func (v *EditType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EditType(value)
	for _, existing := range AllowedEditTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EditType", value)
}

// NewEditTypeFromValue returns a pointer to a valid EditType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEditTypeFromValue(v string) (*EditType, error) {
	ev := EditType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EditType: valid values are %v", v, AllowedEditTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EditType) IsValid() bool {
	for _, existing := range AllowedEditTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EditType value
func (v EditType) Ptr() *EditType {
	return &v
}

type NullableEditType struct {
	value *EditType
	isSet bool
}

func (v NullableEditType) Get() *EditType {
	return v.value
}

func (v *NullableEditType) Set(val *EditType) {
	v.value = val
	v.isSet = true
}

func (v NullableEditType) IsSet() bool {
	return v.isSet
}

func (v *NullableEditType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditType(val *EditType) *NullableEditType {
	return &NullableEditType{value: val, isSet: true}
}

func (v NullableEditType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
