/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTGTolConstraintType the model 'GBTGTolConstraintType'
type GBTGTolConstraintType string

// List of GBTGTolConstraintType
const (
	GBTGTolConstraintTypeTruePosition     GBTGTolConstraintType = "TRUE_POSITION"
	GBTGTolConstraintTypeParallelism      GBTGTolConstraintType = "PARALLELISM"
	GBTGTolConstraintTypePerpendicularity GBTGTolConstraintType = "PERPENDICULARITY"
	GBTGTolConstraintTypeProfileSurface   GBTGTolConstraintType = "PROFILE_SURFACE"
	GBTGTolConstraintTypeTotalRunout      GBTGTolConstraintType = "TOTAL_RUNOUT"
	GBTGTolConstraintTypeUnknown          GBTGTolConstraintType = "UNKNOWN"
)

// All allowed values of GBTGTolConstraintType enum
var AllowedGBTGTolConstraintTypeEnumValues = []GBTGTolConstraintType{
	"TRUE_POSITION",
	"PARALLELISM",
	"PERPENDICULARITY",
	"PROFILE_SURFACE",
	"TOTAL_RUNOUT",
	"UNKNOWN",
}

func (v *GBTGTolConstraintType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTGTolConstraintType(value)
	for _, existing := range AllowedGBTGTolConstraintTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTGTolConstraintType", value)
}

// NewGBTGTolConstraintTypeFromValue returns a pointer to a valid GBTGTolConstraintType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTGTolConstraintTypeFromValue(v string) (*GBTGTolConstraintType, error) {
	ev := GBTGTolConstraintType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTGTolConstraintType: valid values are %v", v, AllowedGBTGTolConstraintTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTGTolConstraintType) IsValid() bool {
	for _, existing := range AllowedGBTGTolConstraintTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTGTolConstraintType value
func (v GBTGTolConstraintType) Ptr() *GBTGTolConstraintType {
	return &v
}

type NullableGBTGTolConstraintType struct {
	value *GBTGTolConstraintType
	isSet bool
}

func (v NullableGBTGTolConstraintType) Get() *GBTGTolConstraintType {
	return v.value
}

func (v *NullableGBTGTolConstraintType) Set(val *GBTGTolConstraintType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTGTolConstraintType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTGTolConstraintType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTGTolConstraintType(val *GBTGTolConstraintType) *NullableGBTGTolConstraintType {
	return &NullableGBTGTolConstraintType{value: val, isSet: true}
}

func (v NullableGBTGTolConstraintType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTGTolConstraintType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
