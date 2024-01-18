/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.175.29152-5591e93bd1c1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTBillOfMaterialsExclusionStatus the model 'GBTBillOfMaterialsExclusionStatus'
type GBTBillOfMaterialsExclusionStatus string

// List of GBTBillOfMaterialsExclusionStatus
const (
	GBTBillOfMaterialsExclusionStatusNotExcluded    GBTBillOfMaterialsExclusionStatus = "NOT_EXCLUDED"
	GBTBillOfMaterialsExclusionStatusParentExcluded GBTBillOfMaterialsExclusionStatus = "PARENT_EXCLUDED"
	GBTBillOfMaterialsExclusionStatusExcluded       GBTBillOfMaterialsExclusionStatus = "EXCLUDED"
	GBTBillOfMaterialsExclusionStatusUnknown        GBTBillOfMaterialsExclusionStatus = "UNKNOWN"
)

// All allowed values of GBTBillOfMaterialsExclusionStatus enum
var AllowedGBTBillOfMaterialsExclusionStatusEnumValues = []GBTBillOfMaterialsExclusionStatus{
	"NOT_EXCLUDED",
	"PARENT_EXCLUDED",
	"EXCLUDED",
	"UNKNOWN",
}

func (v *GBTBillOfMaterialsExclusionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTBillOfMaterialsExclusionStatus(value)
	for _, existing := range AllowedGBTBillOfMaterialsExclusionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTBillOfMaterialsExclusionStatus", value)
}

// NewGBTBillOfMaterialsExclusionStatusFromValue returns a pointer to a valid GBTBillOfMaterialsExclusionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTBillOfMaterialsExclusionStatusFromValue(v string) (*GBTBillOfMaterialsExclusionStatus, error) {
	ev := GBTBillOfMaterialsExclusionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTBillOfMaterialsExclusionStatus: valid values are %v", v, AllowedGBTBillOfMaterialsExclusionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTBillOfMaterialsExclusionStatus) IsValid() bool {
	for _, existing := range AllowedGBTBillOfMaterialsExclusionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTBillOfMaterialsExclusionStatus value
func (v GBTBillOfMaterialsExclusionStatus) Ptr() *GBTBillOfMaterialsExclusionStatus {
	return &v
}

type NullableGBTBillOfMaterialsExclusionStatus struct {
	value *GBTBillOfMaterialsExclusionStatus
	isSet bool
}

func (v NullableGBTBillOfMaterialsExclusionStatus) Get() *GBTBillOfMaterialsExclusionStatus {
	return v.value
}

func (v *NullableGBTBillOfMaterialsExclusionStatus) Set(val *GBTBillOfMaterialsExclusionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTBillOfMaterialsExclusionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTBillOfMaterialsExclusionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTBillOfMaterialsExclusionStatus(val *GBTBillOfMaterialsExclusionStatus) *NullableGBTBillOfMaterialsExclusionStatus {
	return &NullableGBTBillOfMaterialsExclusionStatus{value: val, isSet: true}
}

func (v NullableGBTBillOfMaterialsExclusionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTBillOfMaterialsExclusionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
