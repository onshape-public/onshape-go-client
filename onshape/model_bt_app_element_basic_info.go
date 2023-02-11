/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.159.11153-01265f451c50
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppElementBasicInfo struct for BTAppElementBasicInfo
type BTAppElementBasicInfo struct {
	ChangeId *string `json:"changeId,omitempty"`
	// The numeric code identifying the error that occurred, if one occurred.
	ErrorCode *int32 `json:"errorCode,omitempty"`
	// A human-readable value for the error that occurred, if one occurred.
	ErrorDescription *string `json:"errorDescription,omitempty"`
	ErrorValue       *string `json:"errorValue,omitempty"`
	ParentChangeId   *string `json:"parentChangeId,omitempty"`
}

// NewBTAppElementBasicInfo instantiates a new BTAppElementBasicInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementBasicInfo() *BTAppElementBasicInfo {
	this := BTAppElementBasicInfo{}
	return &this
}

// NewBTAppElementBasicInfoWithDefaults instantiates a new BTAppElementBasicInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementBasicInfoWithDefaults() *BTAppElementBasicInfo {
	this := BTAppElementBasicInfo{}
	return &this
}

// GetChangeId returns the ChangeId field value if set, zero value otherwise.
func (o *BTAppElementBasicInfo) GetChangeId() string {
	if o == nil || o.ChangeId == nil {
		var ret string
		return ret
	}
	return *o.ChangeId
}

// GetChangeIdOk returns a tuple with the ChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementBasicInfo) GetChangeIdOk() (*string, bool) {
	if o == nil || o.ChangeId == nil {
		return nil, false
	}
	return o.ChangeId, true
}

// HasChangeId returns a boolean if a field has been set.
func (o *BTAppElementBasicInfo) HasChangeId() bool {
	if o != nil && o.ChangeId != nil {
		return true
	}

	return false
}

// SetChangeId gets a reference to the given string and assigns it to the ChangeId field.
func (o *BTAppElementBasicInfo) SetChangeId(v string) {
	o.ChangeId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *BTAppElementBasicInfo) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementBasicInfo) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *BTAppElementBasicInfo) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *BTAppElementBasicInfo) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *BTAppElementBasicInfo) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementBasicInfo) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *BTAppElementBasicInfo) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *BTAppElementBasicInfo) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorValue returns the ErrorValue field value if set, zero value otherwise.
func (o *BTAppElementBasicInfo) GetErrorValue() string {
	if o == nil || o.ErrorValue == nil {
		var ret string
		return ret
	}
	return *o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementBasicInfo) GetErrorValueOk() (*string, bool) {
	if o == nil || o.ErrorValue == nil {
		return nil, false
	}
	return o.ErrorValue, true
}

// HasErrorValue returns a boolean if a field has been set.
func (o *BTAppElementBasicInfo) HasErrorValue() bool {
	if o != nil && o.ErrorValue != nil {
		return true
	}

	return false
}

// SetErrorValue gets a reference to the given string and assigns it to the ErrorValue field.
func (o *BTAppElementBasicInfo) SetErrorValue(v string) {
	o.ErrorValue = &v
}

// GetParentChangeId returns the ParentChangeId field value if set, zero value otherwise.
func (o *BTAppElementBasicInfo) GetParentChangeId() string {
	if o == nil || o.ParentChangeId == nil {
		var ret string
		return ret
	}
	return *o.ParentChangeId
}

// GetParentChangeIdOk returns a tuple with the ParentChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementBasicInfo) GetParentChangeIdOk() (*string, bool) {
	if o == nil || o.ParentChangeId == nil {
		return nil, false
	}
	return o.ParentChangeId, true
}

// HasParentChangeId returns a boolean if a field has been set.
func (o *BTAppElementBasicInfo) HasParentChangeId() bool {
	if o != nil && o.ParentChangeId != nil {
		return true
	}

	return false
}

// SetParentChangeId gets a reference to the given string and assigns it to the ParentChangeId field.
func (o *BTAppElementBasicInfo) SetParentChangeId(v string) {
	o.ParentChangeId = &v
}

func (o BTAppElementBasicInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeId != nil {
		toSerialize["changeId"] = o.ChangeId
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorDescription != nil {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if o.ErrorValue != nil {
		toSerialize["errorValue"] = o.ErrorValue
	}
	if o.ParentChangeId != nil {
		toSerialize["parentChangeId"] = o.ParentChangeId
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementBasicInfo struct {
	value *BTAppElementBasicInfo
	isSet bool
}

func (v NullableBTAppElementBasicInfo) Get() *BTAppElementBasicInfo {
	return v.value
}

func (v *NullableBTAppElementBasicInfo) Set(val *BTAppElementBasicInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementBasicInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementBasicInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementBasicInfo(val *BTAppElementBasicInfo) *NullableBTAppElementBasicInfo {
	return &NullableBTAppElementBasicInfo{value: val, isSet: true}
}

func (v NullableBTAppElementBasicInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementBasicInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
