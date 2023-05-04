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

// GBTEdgeTopology the model 'GBTEdgeTopology'
type GBTEdgeTopology string

// List of GBTEdgeTopology
const (
	WIRE      GBTEdgeTopology = "WIRE"
	ONE_SIDED GBTEdgeTopology = "ONE_SIDED"
	TWO_SIDED GBTEdgeTopology = "TWO_SIDED"
	LAMINAR   GBTEdgeTopology = "LAMINAR"
	UNKNOWN   GBTEdgeTopology = "UNKNOWN"
)

// All allowed values of GBTEdgeTopology enum
var AllowedGBTEdgeTopologyEnumValues = []GBTEdgeTopology{
	"WIRE",
	"ONE_SIDED",
	"TWO_SIDED",
	"LAMINAR",
	"UNKNOWN",
}

func (v *GBTEdgeTopology) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTEdgeTopology(value)
	for _, existing := range AllowedGBTEdgeTopologyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTEdgeTopology", value)
}

// NewGBTEdgeTopologyFromValue returns a pointer to a valid GBTEdgeTopology
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTEdgeTopologyFromValue(v string) (*GBTEdgeTopology, error) {
	ev := GBTEdgeTopology(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTEdgeTopology: valid values are %v", v, AllowedGBTEdgeTopologyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTEdgeTopology) IsValid() bool {
	for _, existing := range AllowedGBTEdgeTopologyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTEdgeTopology value
func (v GBTEdgeTopology) Ptr() *GBTEdgeTopology {
	return &v
}

type NullableGBTEdgeTopology struct {
	value *GBTEdgeTopology
	isSet bool
}

func (v NullableGBTEdgeTopology) Get() *GBTEdgeTopology {
	return v.value
}

func (v *NullableGBTEdgeTopology) Set(val *GBTEdgeTopology) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTEdgeTopology) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTEdgeTopology) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTEdgeTopology(val *GBTEdgeTopology) *NullableGBTEdgeTopology {
	return &NullableGBTEdgeTopology{value: val, isSet: true}
}

func (v NullableGBTEdgeTopology) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTEdgeTopology) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
