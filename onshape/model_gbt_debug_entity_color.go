/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.163.15457-d8ebaa9b9e42
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTDebugEntityColor the model 'GBTDebugEntityColor'
type GBTDebugEntityColor string

// List of GBTDebugEntityColor
const (
	RED                GBTDebugEntityColor = "RED"
	GREEN              GBTDebugEntityColor = "GREEN"
	BLUE               GBTDebugEntityColor = "BLUE"
	CYAN               GBTDebugEntityColor = "CYAN"
	YELLOW             GBTDebugEntityColor = "YELLOW"
	MAGENTA            GBTDebugEntityColor = "MAGENTA"
	BLACK              GBTDebugEntityColor = "BLACK"
	TRANSLUCENT_PURPLE GBTDebugEntityColor = "TRANSLUCENT_PURPLE"
	FEATURE_ERROR      GBTDebugEntityColor = "FEATURE_ERROR"
	FEATURE_DEBUG      GBTDebugEntityColor = "FEATURE_DEBUG"
	TRANSLUCENT_GREEN  GBTDebugEntityColor = "TRANSLUCENT_GREEN"
	TRANSLUCENT_BLUE   GBTDebugEntityColor = "TRANSLUCENT_BLUE"
	TRANSLUCENT_CYAN   GBTDebugEntityColor = "TRANSLUCENT_CYAN"
	TRANSLUCENT_YELLOW GBTDebugEntityColor = "TRANSLUCENT_YELLOW"
	TRANSLUCENT_BLACK  GBTDebugEntityColor = "TRANSLUCENT_BLACK"
	TRANSLUCENT_ORANGE GBTDebugEntityColor = "TRANSLUCENT_ORANGE"
	UNKNOWN            GBTDebugEntityColor = "UNKNOWN"
)

// All allowed values of GBTDebugEntityColor enum
var AllowedGBTDebugEntityColorEnumValues = []GBTDebugEntityColor{
	"RED",
	"GREEN",
	"BLUE",
	"CYAN",
	"YELLOW",
	"MAGENTA",
	"BLACK",
	"TRANSLUCENT_PURPLE",
	"FEATURE_ERROR",
	"FEATURE_DEBUG",
	"TRANSLUCENT_GREEN",
	"TRANSLUCENT_BLUE",
	"TRANSLUCENT_CYAN",
	"TRANSLUCENT_YELLOW",
	"TRANSLUCENT_BLACK",
	"TRANSLUCENT_ORANGE",
	"UNKNOWN",
}

func (v *GBTDebugEntityColor) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTDebugEntityColor(value)
	for _, existing := range AllowedGBTDebugEntityColorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTDebugEntityColor", value)
}

// NewGBTDebugEntityColorFromValue returns a pointer to a valid GBTDebugEntityColor
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTDebugEntityColorFromValue(v string) (*GBTDebugEntityColor, error) {
	ev := GBTDebugEntityColor(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTDebugEntityColor: valid values are %v", v, AllowedGBTDebugEntityColorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTDebugEntityColor) IsValid() bool {
	for _, existing := range AllowedGBTDebugEntityColorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTDebugEntityColor value
func (v GBTDebugEntityColor) Ptr() *GBTDebugEntityColor {
	return &v
}

type NullableGBTDebugEntityColor struct {
	value *GBTDebugEntityColor
	isSet bool
}

func (v NullableGBTDebugEntityColor) Get() *GBTDebugEntityColor {
	return v.value
}

func (v *NullableGBTDebugEntityColor) Set(val *GBTDebugEntityColor) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTDebugEntityColor) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTDebugEntityColor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTDebugEntityColor(val *GBTDebugEntityColor) *NullableGBTDebugEntityColor {
	return &NullableGBTDebugEntityColor{value: val, isSet: true}
}

func (v NullableGBTDebugEntityColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTDebugEntityColor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
