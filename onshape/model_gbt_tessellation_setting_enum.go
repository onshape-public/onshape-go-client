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

// GBTTessellationSettingEnum the model 'GBTTessellationSettingEnum'
type GBTTessellationSettingEnum string

// List of GBTTessellationSettingEnum
const (
	GBTTessellationSettingEnumAuto                   GBTTessellationSettingEnum = "AUTO"
	GBTTessellationSettingEnumCoarse                 GBTTessellationSettingEnum = "COARSE"
	GBTTessellationSettingEnumMedium                 GBTTessellationSettingEnum = "MEDIUM"
	GBTTessellationSettingEnumFine                   GBTTessellationSettingEnum = "FINE"
	GBTTessellationSettingEnumVeryFine               GBTTessellationSettingEnum = "VERY_FINE"
	GBTTessellationSettingEnumCurvatureVisualization GBTTessellationSettingEnum = "CURVATURE_VISUALIZATION"
	GBTTessellationSettingEnumUnknown                GBTTessellationSettingEnum = "UNKNOWN"
)

// All allowed values of GBTTessellationSettingEnum enum
var AllowedGBTTessellationSettingEnumEnumValues = []GBTTessellationSettingEnum{
	"AUTO",
	"COARSE",
	"MEDIUM",
	"FINE",
	"VERY_FINE",
	"CURVATURE_VISUALIZATION",
	"UNKNOWN",
}

func (v *GBTTessellationSettingEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTTessellationSettingEnum(value)
	for _, existing := range AllowedGBTTessellationSettingEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTTessellationSettingEnum", value)
}

// NewGBTTessellationSettingEnumFromValue returns a pointer to a valid GBTTessellationSettingEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTTessellationSettingEnumFromValue(v string) (*GBTTessellationSettingEnum, error) {
	ev := GBTTessellationSettingEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTTessellationSettingEnum: valid values are %v", v, AllowedGBTTessellationSettingEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTTessellationSettingEnum) IsValid() bool {
	for _, existing := range AllowedGBTTessellationSettingEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTTessellationSettingEnum value
func (v GBTTessellationSettingEnum) Ptr() *GBTTessellationSettingEnum {
	return &v
}

type NullableGBTTessellationSettingEnum struct {
	value *GBTTessellationSettingEnum
	isSet bool
}

func (v NullableGBTTessellationSettingEnum) Get() *GBTTessellationSettingEnum {
	return v.value
}

func (v *NullableGBTTessellationSettingEnum) Set(val *GBTTessellationSettingEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTTessellationSettingEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTTessellationSettingEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTTessellationSettingEnum(val *GBTTessellationSettingEnum) *NullableGBTTessellationSettingEnum {
	return &NullableGBTTessellationSettingEnum{value: val, isSet: true}
}

func (v NullableGBTTessellationSettingEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTTessellationSettingEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
