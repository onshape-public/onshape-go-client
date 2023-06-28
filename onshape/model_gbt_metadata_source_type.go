/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.18120-f464f720d215
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTMetadataSourceType the model 'GBTMetadataSourceType'
type GBTMetadataSourceType string

// List of GBTMetadataSourceType
const (
	GBTMetadataSourceTypeAutomatic          GBTMetadataSourceType = "AUTOMATIC"
	GBTMetadataSourceTypeMerged             GBTMetadataSourceType = "MERGED"
	GBTMetadataSourceTypeFeature            GBTMetadataSourceType = "FEATURE"
	GBTMetadataSourceTypeUnconfigured       GBTMetadataSourceType = "UNCONFIGURED"
	GBTMetadataSourceTypeConfigured         GBTMetadataSourceType = "CONFIGURED"
	GBTMetadataSourceTypeStandardContent    GBTMetadataSourceType = "STANDARD_CONTENT"
	GBTMetadataSourceTypeDefault            GBTMetadataSourceType = "DEFAULT"
	GBTMetadataSourceTypeComputed           GBTMetadataSourceType = "COMPUTED"
	GBTMetadataSourceTypeComputedConfigured GBTMetadataSourceType = "COMPUTED_CONFIGURED"
	GBTMetadataSourceTypeUnknown            GBTMetadataSourceType = "UNKNOWN"
)

// All allowed values of GBTMetadataSourceType enum
var AllowedGBTMetadataSourceTypeEnumValues = []GBTMetadataSourceType{
	"AUTOMATIC",
	"MERGED",
	"FEATURE",
	"UNCONFIGURED",
	"CONFIGURED",
	"STANDARD_CONTENT",
	"DEFAULT",
	"COMPUTED",
	"COMPUTED_CONFIGURED",
	"UNKNOWN",
}

func (v *GBTMetadataSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTMetadataSourceType(value)
	for _, existing := range AllowedGBTMetadataSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTMetadataSourceType", value)
}

// NewGBTMetadataSourceTypeFromValue returns a pointer to a valid GBTMetadataSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTMetadataSourceTypeFromValue(v string) (*GBTMetadataSourceType, error) {
	ev := GBTMetadataSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTMetadataSourceType: valid values are %v", v, AllowedGBTMetadataSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTMetadataSourceType) IsValid() bool {
	for _, existing := range AllowedGBTMetadataSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTMetadataSourceType value
func (v GBTMetadataSourceType) Ptr() *GBTMetadataSourceType {
	return &v
}

type NullableGBTMetadataSourceType struct {
	value *GBTMetadataSourceType
	isSet bool
}

func (v NullableGBTMetadataSourceType) Get() *GBTMetadataSourceType {
	return v.value
}

func (v *NullableGBTMetadataSourceType) Set(val *GBTMetadataSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTMetadataSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTMetadataSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTMetadataSourceType(val *GBTMetadataSourceType) *NullableGBTMetadataSourceType {
	return &NullableGBTMetadataSourceType{value: val, isSet: true}
}

func (v NullableGBTMetadataSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTMetadataSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
