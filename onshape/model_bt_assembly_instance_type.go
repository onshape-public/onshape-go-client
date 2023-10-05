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

// BTAssemblyInstanceType the model 'BTAssemblyInstanceType'
type BTAssemblyInstanceType string

// List of BTAssemblyInstanceType
const (
	BTAssemblyInstanceTypeAssembly BTAssemblyInstanceType = "Assembly"
	BTAssemblyInstanceTypeFeature  BTAssemblyInstanceType = "Feature"
	BTAssemblyInstanceTypePart     BTAssemblyInstanceType = "Part"
	BTAssemblyInstanceTypeUnknown  BTAssemblyInstanceType = "Unknown"
)

// All allowed values of BTAssemblyInstanceType enum
var AllowedBTAssemblyInstanceTypeEnumValues = []BTAssemblyInstanceType{
	"Assembly",
	"Feature",
	"Part",
	"Unknown",
}

func (v *BTAssemblyInstanceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTAssemblyInstanceType(value)
	for _, existing := range AllowedBTAssemblyInstanceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTAssemblyInstanceType", value)
}

// NewBTAssemblyInstanceTypeFromValue returns a pointer to a valid BTAssemblyInstanceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTAssemblyInstanceTypeFromValue(v string) (*BTAssemblyInstanceType, error) {
	ev := BTAssemblyInstanceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTAssemblyInstanceType: valid values are %v", v, AllowedBTAssemblyInstanceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTAssemblyInstanceType) IsValid() bool {
	for _, existing := range AllowedBTAssemblyInstanceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTAssemblyInstanceType value
func (v BTAssemblyInstanceType) Ptr() *BTAssemblyInstanceType {
	return &v
}

type NullableBTAssemblyInstanceType struct {
	value *BTAssemblyInstanceType
	isSet bool
}

func (v NullableBTAssemblyInstanceType) Get() *BTAssemblyInstanceType {
	return v.value
}

func (v *NullableBTAssemblyInstanceType) Set(val *BTAssemblyInstanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyInstanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyInstanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyInstanceType(val *BTAssemblyInstanceType) *NullableBTAssemblyInstanceType {
	return &NullableBTAssemblyInstanceType{value: val, isSet: true}
}

func (v NullableBTAssemblyInstanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyInstanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
