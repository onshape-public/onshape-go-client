/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23196-ae5f57bec467
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTAppElementErrorCode the model 'BTAppElementErrorCode'
type BTAppElementErrorCode string

// List of BTAppElementErrorCode
const (
	BTAppElementErrorCodeOk                  BTAppElementErrorCode = "OK"
	BTAppElementErrorCodeTransactionConflict BTAppElementErrorCode = "TRANSACTION_CONFLICT"
	BTAppElementErrorCodeNotFound            BTAppElementErrorCode = "NOT_FOUND"
	BTAppElementErrorCodeInconsistentChanges BTAppElementErrorCode = "INCONSISTENT_CHANGES"
)

// All allowed values of BTAppElementErrorCode enum
var AllowedBTAppElementErrorCodeEnumValues = []BTAppElementErrorCode{
	"OK",
	"TRANSACTION_CONFLICT",
	"NOT_FOUND",
	"INCONSISTENT_CHANGES",
}

func (v *BTAppElementErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTAppElementErrorCode(value)
	for _, existing := range AllowedBTAppElementErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTAppElementErrorCode", value)
}

// NewBTAppElementErrorCodeFromValue returns a pointer to a valid BTAppElementErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTAppElementErrorCodeFromValue(v string) (*BTAppElementErrorCode, error) {
	ev := BTAppElementErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTAppElementErrorCode: valid values are %v", v, AllowedBTAppElementErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTAppElementErrorCode) IsValid() bool {
	for _, existing := range AllowedBTAppElementErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTAppElementErrorCode value
func (v BTAppElementErrorCode) Ptr() *BTAppElementErrorCode {
	return &v
}

type NullableBTAppElementErrorCode struct {
	value *BTAppElementErrorCode
	isSet bool
}

func (v NullableBTAppElementErrorCode) Get() *BTAppElementErrorCode {
	return v.value
}

func (v *NullableBTAppElementErrorCode) Set(val *BTAppElementErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementErrorCode(val *BTAppElementErrorCode) *NullableBTAppElementErrorCode {
	return &NullableBTAppElementErrorCode{value: val, isSet: true}
}

func (v NullableBTAppElementErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
