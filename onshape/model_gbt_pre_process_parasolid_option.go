/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25478-d4e5ab4765a4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTPreProcessParasolidOption the model 'GBTPreProcessParasolidOption'
type GBTPreProcessParasolidOption string

// List of GBTPreProcessParasolidOption
const (
	GBTPreProcessParasolidOptionNoPreProcessing                  GBTPreProcessParasolidOption = "NO_PRE_PROCESSING"
	GBTPreProcessParasolidOptionUseBodyshopPreProcessingAdvanced GBTPreProcessParasolidOption = "USE_BODYSHOP_PRE_PROCESSING_ADVANCED"
	GBTPreProcessParasolidOptionUseTranslatorPreProcessing       GBTPreProcessParasolidOption = "USE_TRANSLATOR_PRE_PROCESSING"
	GBTPreProcessParasolidOptionUseBodyshopPreProcessing         GBTPreProcessParasolidOption = "USE_BODYSHOP_PRE_PROCESSING"
	GBTPreProcessParasolidOptionUnknown                          GBTPreProcessParasolidOption = "UNKNOWN"
)

// All allowed values of GBTPreProcessParasolidOption enum
var AllowedGBTPreProcessParasolidOptionEnumValues = []GBTPreProcessParasolidOption{
	"NO_PRE_PROCESSING",
	"USE_BODYSHOP_PRE_PROCESSING_ADVANCED",
	"USE_TRANSLATOR_PRE_PROCESSING",
	"USE_BODYSHOP_PRE_PROCESSING",
	"UNKNOWN",
}

func (v *GBTPreProcessParasolidOption) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTPreProcessParasolidOption(value)
	for _, existing := range AllowedGBTPreProcessParasolidOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTPreProcessParasolidOption", value)
}

// NewGBTPreProcessParasolidOptionFromValue returns a pointer to a valid GBTPreProcessParasolidOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTPreProcessParasolidOptionFromValue(v string) (*GBTPreProcessParasolidOption, error) {
	ev := GBTPreProcessParasolidOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTPreProcessParasolidOption: valid values are %v", v, AllowedGBTPreProcessParasolidOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTPreProcessParasolidOption) IsValid() bool {
	for _, existing := range AllowedGBTPreProcessParasolidOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTPreProcessParasolidOption value
func (v GBTPreProcessParasolidOption) Ptr() *GBTPreProcessParasolidOption {
	return &v
}

type NullableGBTPreProcessParasolidOption struct {
	value *GBTPreProcessParasolidOption
	isSet bool
}

func (v NullableGBTPreProcessParasolidOption) Get() *GBTPreProcessParasolidOption {
	return v.value
}

func (v *NullableGBTPreProcessParasolidOption) Set(val *GBTPreProcessParasolidOption) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTPreProcessParasolidOption) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTPreProcessParasolidOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTPreProcessParasolidOption(val *GBTPreProcessParasolidOption) *NullableGBTPreProcessParasolidOption {
	return &NullableGBTPreProcessParasolidOption{value: val, isSet: true}
}

func (v NullableGBTPreProcessParasolidOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTPreProcessParasolidOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
