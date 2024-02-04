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

// BTSearchEntityType the model 'BTSearchEntityType'
type BTSearchEntityType string

// List of BTSearchEntityType
const (
	BTSearchEntityTypeUnknown     BTSearchEntityType = "unknown"
	BTSearchEntityTypeCapability  BTSearchEntityType = "capability"
	BTSearchEntityTypeCompany     BTSearchEntityType = "company"
	BTSearchEntityTypeDocument    BTSearchEntityType = "document"
	BTSearchEntityTypeElement     BTSearchEntityType = "element"
	BTSearchEntityTypeFriend      BTSearchEntityType = "friend"
	BTSearchEntityTypeItem        BTSearchEntityType = "item"
	BTSearchEntityTypePart        BTSearchEntityType = "part"
	BTSearchEntityTypeTeam        BTSearchEntityType = "team"
	BTSearchEntityTypeUser        BTSearchEntityType = "user"
	BTSearchEntityTypeVersion     BTSearchEntityType = "version"
	BTSearchEntityTypeWorkspace   BTSearchEntityType = "workspace"
	BTSearchEntityTypeProject     BTSearchEntityType = "project"
	BTSearchEntityTypePublication BTSearchEntityType = "publication"
)

// All allowed values of BTSearchEntityType enum
var AllowedBTSearchEntityTypeEnumValues = []BTSearchEntityType{
	"unknown",
	"capability",
	"company",
	"document",
	"element",
	"friend",
	"item",
	"part",
	"team",
	"user",
	"version",
	"workspace",
	"project",
	"publication",
}

func (v *BTSearchEntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTSearchEntityType(value)
	for _, existing := range AllowedBTSearchEntityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTSearchEntityType", value)
}

// NewBTSearchEntityTypeFromValue returns a pointer to a valid BTSearchEntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTSearchEntityTypeFromValue(v string) (*BTSearchEntityType, error) {
	ev := BTSearchEntityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTSearchEntityType: valid values are %v", v, AllowedBTSearchEntityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTSearchEntityType) IsValid() bool {
	for _, existing := range AllowedBTSearchEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTSearchEntityType value
func (v BTSearchEntityType) Ptr() *BTSearchEntityType {
	return &v
}

type NullableBTSearchEntityType struct {
	value *BTSearchEntityType
	isSet bool
}

func (v NullableBTSearchEntityType) Get() *BTSearchEntityType {
	return v.value
}

func (v *NullableBTSearchEntityType) Set(val *BTSearchEntityType) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSearchEntityType) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSearchEntityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSearchEntityType(val *BTSearchEntityType) *NullableBTSearchEntityType {
	return &NullableBTSearchEntityType{value: val, isSet: true}
}

func (v NullableBTSearchEntityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSearchEntityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
