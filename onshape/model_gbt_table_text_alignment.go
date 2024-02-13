/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTTableTextAlignment the model 'GBTTableTextAlignment'
type GBTTableTextAlignment string

// List of GBTTableTextAlignment
const (
	GBTTableTextAlignmentLeft    GBTTableTextAlignment = "LEFT"
	GBTTableTextAlignmentCenter  GBTTableTextAlignment = "CENTER"
	GBTTableTextAlignmentRight   GBTTableTextAlignment = "RIGHT"
	GBTTableTextAlignmentUnknown GBTTableTextAlignment = "UNKNOWN"
)

// All allowed values of GBTTableTextAlignment enum
var AllowedGBTTableTextAlignmentEnumValues = []GBTTableTextAlignment{
	"LEFT",
	"CENTER",
	"RIGHT",
	"UNKNOWN",
}

func (v *GBTTableTextAlignment) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTTableTextAlignment(value)
	for _, existing := range AllowedGBTTableTextAlignmentEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTTableTextAlignment", value)
}

// NewGBTTableTextAlignmentFromValue returns a pointer to a valid GBTTableTextAlignment
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTTableTextAlignmentFromValue(v string) (*GBTTableTextAlignment, error) {
	ev := GBTTableTextAlignment(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTTableTextAlignment: valid values are %v", v, AllowedGBTTableTextAlignmentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTTableTextAlignment) IsValid() bool {
	for _, existing := range AllowedGBTTableTextAlignmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTTableTextAlignment value
func (v GBTTableTextAlignment) Ptr() *GBTTableTextAlignment {
	return &v
}

type NullableGBTTableTextAlignment struct {
	value *GBTTableTextAlignment
	isSet bool
}

func (v NullableGBTTableTextAlignment) Get() *GBTTableTextAlignment {
	return v.value
}

func (v *NullableGBTTableTextAlignment) Set(val *GBTTableTextAlignment) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTTableTextAlignment) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTTableTextAlignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTTableTextAlignment(val *GBTTableTextAlignment) *NullableGBTTableTextAlignment {
	return &NullableGBTTableTextAlignment{value: val, isSet: true}
}

func (v NullableGBTTableTextAlignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTTableTextAlignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
