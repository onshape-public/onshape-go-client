/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.169.21759-9ddbba9fdfb8
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTBillOfMaterialsExpansionStatus the model 'GBTBillOfMaterialsExpansionStatus'
type GBTBillOfMaterialsExpansionStatus string

// List of GBTBillOfMaterialsExpansionStatus
const (
	GBTBillOfMaterialsExpansionStatusNotExpandable GBTBillOfMaterialsExpansionStatus = "NOT_EXPANDABLE"
	GBTBillOfMaterialsExpansionStatusExpanded      GBTBillOfMaterialsExpansionStatus = "EXPANDED"
	GBTBillOfMaterialsExpansionStatusCollapsed     GBTBillOfMaterialsExpansionStatus = "COLLAPSED"
	GBTBillOfMaterialsExpansionStatusUnknown       GBTBillOfMaterialsExpansionStatus = "UNKNOWN"
)

// All allowed values of GBTBillOfMaterialsExpansionStatus enum
var AllowedGBTBillOfMaterialsExpansionStatusEnumValues = []GBTBillOfMaterialsExpansionStatus{
	"NOT_EXPANDABLE",
	"EXPANDED",
	"COLLAPSED",
	"UNKNOWN",
}

func (v *GBTBillOfMaterialsExpansionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTBillOfMaterialsExpansionStatus(value)
	for _, existing := range AllowedGBTBillOfMaterialsExpansionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTBillOfMaterialsExpansionStatus", value)
}

// NewGBTBillOfMaterialsExpansionStatusFromValue returns a pointer to a valid GBTBillOfMaterialsExpansionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTBillOfMaterialsExpansionStatusFromValue(v string) (*GBTBillOfMaterialsExpansionStatus, error) {
	ev := GBTBillOfMaterialsExpansionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTBillOfMaterialsExpansionStatus: valid values are %v", v, AllowedGBTBillOfMaterialsExpansionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTBillOfMaterialsExpansionStatus) IsValid() bool {
	for _, existing := range AllowedGBTBillOfMaterialsExpansionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTBillOfMaterialsExpansionStatus value
func (v GBTBillOfMaterialsExpansionStatus) Ptr() *GBTBillOfMaterialsExpansionStatus {
	return &v
}

type NullableGBTBillOfMaterialsExpansionStatus struct {
	value *GBTBillOfMaterialsExpansionStatus
	isSet bool
}

func (v NullableGBTBillOfMaterialsExpansionStatus) Get() *GBTBillOfMaterialsExpansionStatus {
	return v.value
}

func (v *NullableGBTBillOfMaterialsExpansionStatus) Set(val *GBTBillOfMaterialsExpansionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTBillOfMaterialsExpansionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTBillOfMaterialsExpansionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTBillOfMaterialsExpansionStatus(val *GBTBillOfMaterialsExpansionStatus) *NullableGBTBillOfMaterialsExpansionStatus {
	return &NullableGBTBillOfMaterialsExpansionStatus{value: val, isSet: true}
}

func (v NullableGBTBillOfMaterialsExpansionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTBillOfMaterialsExpansionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
