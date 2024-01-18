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

// BTMergeStrategy the model 'BTMergeStrategy'
type BTMergeStrategy string

// List of BTMergeStrategy
const (
	BTMergeStrategyMerge   BTMergeStrategy = "MERGE"
	BTMergeStrategyKeep    BTMergeStrategy = "KEEP"
	BTMergeStrategyReplace BTMergeStrategy = "REPLACE"
)

// All allowed values of BTMergeStrategy enum
var AllowedBTMergeStrategyEnumValues = []BTMergeStrategy{
	"MERGE",
	"KEEP",
	"REPLACE",
}

func (v *BTMergeStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTMergeStrategy(value)
	for _, existing := range AllowedBTMergeStrategyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTMergeStrategy", value)
}

// NewBTMergeStrategyFromValue returns a pointer to a valid BTMergeStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTMergeStrategyFromValue(v string) (*BTMergeStrategy, error) {
	ev := BTMergeStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTMergeStrategy: valid values are %v", v, AllowedBTMergeStrategyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTMergeStrategy) IsValid() bool {
	for _, existing := range AllowedBTMergeStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTMergeStrategy value
func (v BTMergeStrategy) Ptr() *BTMergeStrategy {
	return &v
}

type NullableBTMergeStrategy struct {
	value *BTMergeStrategy
	isSet bool
}

func (v NullableBTMergeStrategy) Get() *BTMergeStrategy {
	return v.value
}

func (v *NullableBTMergeStrategy) Set(val *BTMergeStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMergeStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMergeStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMergeStrategy(val *BTMergeStrategy) *NullableBTMergeStrategy {
	return &NullableBTMergeStrategy{value: val, isSet: true}
}

func (v NullableBTMergeStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMergeStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
