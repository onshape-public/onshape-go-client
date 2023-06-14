/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.17369-82f2ed5d514e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTBodyType the model 'GBTBodyType'
type GBTBodyType string

// List of GBTBodyType
const (
	GBTBodyTypeSolid         GBTBodyType = "SOLID"
	GBTBodyTypeSheet         GBTBodyType = "SHEET"
	GBTBodyTypeWire          GBTBodyType = "WIRE"
	GBTBodyTypePoint         GBTBodyType = "POINT"
	GBTBodyTypeMateConnector GBTBodyType = "MATE_CONNECTOR"
	GBTBodyTypeComposite     GBTBodyType = "COMPOSITE"
	GBTBodyTypeUnknown       GBTBodyType = "UNKNOWN"
)

// All allowed values of GBTBodyType enum
var AllowedGBTBodyTypeEnumValues = []GBTBodyType{
	"SOLID",
	"SHEET",
	"WIRE",
	"POINT",
	"MATE_CONNECTOR",
	"COMPOSITE",
	"UNKNOWN",
}

func (v *GBTBodyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTBodyType(value)
	for _, existing := range AllowedGBTBodyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTBodyType", value)
}

// NewGBTBodyTypeFromValue returns a pointer to a valid GBTBodyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTBodyTypeFromValue(v string) (*GBTBodyType, error) {
	ev := GBTBodyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTBodyType: valid values are %v", v, AllowedGBTBodyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTBodyType) IsValid() bool {
	for _, existing := range AllowedGBTBodyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTBodyType value
func (v GBTBodyType) Ptr() *GBTBodyType {
	return &v
}

type NullableGBTBodyType struct {
	value *GBTBodyType
	isSet bool
}

func (v NullableGBTBodyType) Get() *GBTBodyType {
	return v.value
}

func (v *NullableGBTBodyType) Set(val *GBTBodyType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTBodyType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTBodyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTBodyType(val *GBTBodyType) *NullableGBTBodyType {
	return &NullableGBTBodyType{value: val, isSet: true}
}

func (v NullableGBTBodyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTBodyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
