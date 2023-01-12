/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.158.9878-ff51e1211d95
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTFeatureApiBase1430 - struct for BTFeatureApiBase1430
type BTFeatureApiBase1430 struct {
	implBTFeatureApiBase1430 interface{}
}

// BTUpdateFeaturesCall1748AsBTFeatureApiBase1430 is a convenience function that returns BTUpdateFeaturesCall1748 wrapped in BTFeatureApiBase1430
func (o *BTUpdateFeaturesCall1748) AsBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	return &BTFeatureApiBase1430{o}
}

// BTUpdateFeaturesResponse1333AsBTFeatureApiBase1430 is a convenience function that returns BTUpdateFeaturesResponse1333 wrapped in BTFeatureApiBase1430
func (o *BTUpdateFeaturesResponse1333) AsBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	return &BTFeatureApiBase1430{o}
}

// BTSetFeatureRollbackCall1899AsBTFeatureApiBase1430 is a convenience function that returns BTSetFeatureRollbackCall1899 wrapped in BTFeatureApiBase1430
func (o *BTSetFeatureRollbackCall1899) AsBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	return &BTFeatureApiBase1430{o}
}

// BTFeatureSpecsResponse664AsBTFeatureApiBase1430 is a convenience function that returns BTFeatureSpecsResponse664 wrapped in BTFeatureApiBase1430
func (o *BTFeatureSpecsResponse664) AsBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	return &BTFeatureApiBase1430{o}
}

// BTFeatureDefinitionCall1406AsBTFeatureApiBase1430 is a convenience function that returns BTFeatureDefinitionCall1406 wrapped in BTFeatureApiBase1430
func (o *BTFeatureDefinitionCall1406) AsBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	return &BTFeatureApiBase1430{o}
}

// BTAssemblyFeatureListResponse1174AsBTFeatureApiBase1430 is a convenience function that returns BTAssemblyFeatureListResponse1174 wrapped in BTFeatureApiBase1430
func (o *BTAssemblyFeatureListResponse1174) AsBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	return &BTFeatureApiBase1430{o}
}

// BTFeatureDefinitionResponse1617AsBTFeatureApiBase1430 is a convenience function that returns BTFeatureDefinitionResponse1617 wrapped in BTFeatureApiBase1430
func (o *BTFeatureDefinitionResponse1617) AsBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	return &BTFeatureApiBase1430{o}
}

// BTConfigurationResponse2019AsBTFeatureApiBase1430 is a convenience function that returns BTConfigurationResponse2019 wrapped in BTFeatureApiBase1430
func (o *BTConfigurationResponse2019) AsBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	return &BTFeatureApiBase1430{o}
}

// BTFeatureListResponse2457AsBTFeatureApiBase1430 is a convenience function that returns BTFeatureListResponse2457 wrapped in BTFeatureApiBase1430
func (o *BTFeatureListResponse2457) AsBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	return &BTFeatureApiBase1430{o}
}

// BTConfigurationUpdateCall2933AsBTFeatureApiBase1430 is a convenience function that returns BTConfigurationUpdateCall2933 wrapped in BTFeatureApiBase1430
func (o *BTConfigurationUpdateCall2933) AsBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	return &BTFeatureApiBase1430{o}
}

// BTFeatureScriptEvalResponse1859AsBTFeatureApiBase1430 is a convenience function that returns BTFeatureScriptEvalResponse1859 wrapped in BTFeatureApiBase1430
func (o *BTFeatureScriptEvalResponse1859) AsBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	return &BTFeatureApiBase1430{o}
}

// BTFeatureScriptEvalCall2377AsBTFeatureApiBase1430 is a convenience function that returns BTFeatureScriptEvalCall2377 wrapped in BTFeatureApiBase1430
func (o *BTFeatureScriptEvalCall2377) AsBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	return &BTFeatureApiBase1430{o}
}

// BTFeatureStudioContents2239AsBTFeatureApiBase1430 is a convenience function that returns BTFeatureStudioContents2239 wrapped in BTFeatureApiBase1430
func (o *BTFeatureStudioContents2239) AsBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	return &BTFeatureApiBase1430{o}
}

// BTSetFeatureRollbackResponse1042AsBTFeatureApiBase1430 is a convenience function that returns BTSetFeatureRollbackResponse1042 wrapped in BTFeatureApiBase1430
func (o *BTSetFeatureRollbackResponse1042) AsBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	return &BTFeatureApiBase1430{o}
}

// NewBTFeatureApiBase1430 instantiates a new BTFeatureApiBase1430 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureApiBase1430() *BTFeatureApiBase1430 {
	this := BTFeatureApiBase1430{Newbase_BTFeatureApiBase1430()}
	return &this
}

// NewBTFeatureApiBase1430WithDefaults instantiates a new BTFeatureApiBase1430 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureApiBase1430WithDefaults() *BTFeatureApiBase1430 {
	this := BTFeatureApiBase1430{Newbase_BTFeatureApiBase1430WithDefaults()}
	return &this
}

// GetLibraryVersion returns the LibraryVersion field value if set, zero value otherwise.
func (o *BTFeatureApiBase1430) GetLibraryVersion() int32 {
	type getResult interface {
		GetLibraryVersion() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLibraryVersion()
	} else {
		var de int32
		return de
	}
}

// GetLibraryVersionOk returns a tuple with the LibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureApiBase1430) GetLibraryVersionOk() (*int32, bool) {
	type getResult interface {
		GetLibraryVersionOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLibraryVersionOk()
	} else {
		return nil, false
	}
}

// HasLibraryVersion returns a boolean if a field has been set.
func (o *BTFeatureApiBase1430) HasLibraryVersion() bool {
	type getResult interface {
		HasLibraryVersion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasLibraryVersion()
	} else {
		return false
	}
}

// SetLibraryVersion gets a reference to the given int32 and assigns it to the LibraryVersion field.
func (o *BTFeatureApiBase1430) SetLibraryVersion(v int32) {
	type getResult interface {
		SetLibraryVersion(v int32)
	}

	o.GetActualInstance().(getResult).SetLibraryVersion(v)
}

// GetMicroversionSkew returns the MicroversionSkew field value if set, zero value otherwise.
func (o *BTFeatureApiBase1430) GetMicroversionSkew() bool {
	type getResult interface {
		GetMicroversionSkew() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionSkew()
	} else {
		var de bool
		return de
	}
}

// GetMicroversionSkewOk returns a tuple with the MicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureApiBase1430) GetMicroversionSkewOk() (*bool, bool) {
	type getResult interface {
		GetMicroversionSkewOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionSkewOk()
	} else {
		return nil, false
	}
}

// HasMicroversionSkew returns a boolean if a field has been set.
func (o *BTFeatureApiBase1430) HasMicroversionSkew() bool {
	type getResult interface {
		HasMicroversionSkew() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMicroversionSkew()
	} else {
		return false
	}
}

// SetMicroversionSkew gets a reference to the given bool and assigns it to the MicroversionSkew field.
func (o *BTFeatureApiBase1430) SetMicroversionSkew(v bool) {
	type getResult interface {
		SetMicroversionSkew(v bool)
	}

	o.GetActualInstance().(getResult).SetMicroversionSkew(v)
}

// GetRejectMicroversionSkew returns the RejectMicroversionSkew field value if set, zero value otherwise.
func (o *BTFeatureApiBase1430) GetRejectMicroversionSkew() bool {
	type getResult interface {
		GetRejectMicroversionSkew() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRejectMicroversionSkew()
	} else {
		var de bool
		return de
	}
}

// GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureApiBase1430) GetRejectMicroversionSkewOk() (*bool, bool) {
	type getResult interface {
		GetRejectMicroversionSkewOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRejectMicroversionSkewOk()
	} else {
		return nil, false
	}
}

// HasRejectMicroversionSkew returns a boolean if a field has been set.
func (o *BTFeatureApiBase1430) HasRejectMicroversionSkew() bool {
	type getResult interface {
		HasRejectMicroversionSkew() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasRejectMicroversionSkew()
	} else {
		return false
	}
}

// SetRejectMicroversionSkew gets a reference to the given bool and assigns it to the RejectMicroversionSkew field.
func (o *BTFeatureApiBase1430) SetRejectMicroversionSkew(v bool) {
	type getResult interface {
		SetRejectMicroversionSkew(v bool)
	}

	o.GetActualInstance().(getResult).SetRejectMicroversionSkew(v)
}

// GetSerializationVersion returns the SerializationVersion field value if set, zero value otherwise.
func (o *BTFeatureApiBase1430) GetSerializationVersion() string {
	type getResult interface {
		GetSerializationVersion() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSerializationVersion()
	} else {
		var de string
		return de
	}
}

// GetSerializationVersionOk returns a tuple with the SerializationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureApiBase1430) GetSerializationVersionOk() (*string, bool) {
	type getResult interface {
		GetSerializationVersionOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSerializationVersionOk()
	} else {
		return nil, false
	}
}

// HasSerializationVersion returns a boolean if a field has been set.
func (o *BTFeatureApiBase1430) HasSerializationVersion() bool {
	type getResult interface {
		HasSerializationVersion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasSerializationVersion()
	} else {
		return false
	}
}

// SetSerializationVersion gets a reference to the given string and assigns it to the SerializationVersion field.
func (o *BTFeatureApiBase1430) SetSerializationVersion(v string) {
	type getResult interface {
		SetSerializationVersion(v string)
	}

	o.GetActualInstance().(getResult).SetSerializationVersion(v)
}

// GetSourceMicroversion returns the SourceMicroversion field value if set, zero value otherwise.
func (o *BTFeatureApiBase1430) GetSourceMicroversion() string {
	type getResult interface {
		GetSourceMicroversion() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSourceMicroversion()
	} else {
		var de string
		return de
	}
}

// GetSourceMicroversionOk returns a tuple with the SourceMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureApiBase1430) GetSourceMicroversionOk() (*string, bool) {
	type getResult interface {
		GetSourceMicroversionOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSourceMicroversionOk()
	} else {
		return nil, false
	}
}

// HasSourceMicroversion returns a boolean if a field has been set.
func (o *BTFeatureApiBase1430) HasSourceMicroversion() bool {
	type getResult interface {
		HasSourceMicroversion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasSourceMicroversion()
	} else {
		return false
	}
}

// SetSourceMicroversion gets a reference to the given string and assigns it to the SourceMicroversion field.
func (o *BTFeatureApiBase1430) SetSourceMicroversion(v string) {
	type getResult interface {
		SetSourceMicroversion(v string)
	}

	o.GetActualInstance().(getResult).SetSourceMicroversion(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTFeatureApiBase1430) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTAssemblyFeatureListResponse-1174'
	if jsonDict["btType"] == "BTAssemblyFeatureListResponse-1174" {
		// try to unmarshal JSON data into BTAssemblyFeatureListResponse1174
		var qr *BTAssemblyFeatureListResponse1174
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureApiBase1430 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureApiBase1430 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as BTAssemblyFeatureListResponse1174: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTConfigurationResponse-2019'
	if jsonDict["btType"] == "BTConfigurationResponse-2019" {
		// try to unmarshal JSON data into BTConfigurationResponse2019
		var qr *BTConfigurationResponse2019
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureApiBase1430 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureApiBase1430 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as BTConfigurationResponse2019: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTConfigurationUpdateCall-2933'
	if jsonDict["btType"] == "BTConfigurationUpdateCall-2933" {
		// try to unmarshal JSON data into BTConfigurationUpdateCall2933
		var qr *BTConfigurationUpdateCall2933
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureApiBase1430 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureApiBase1430 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as BTConfigurationUpdateCall2933: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFeatureDefinitionCall-1406'
	if jsonDict["btType"] == "BTFeatureDefinitionCall-1406" {
		// try to unmarshal JSON data into BTFeatureDefinitionCall1406
		var qr *BTFeatureDefinitionCall1406
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureApiBase1430 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureApiBase1430 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as BTFeatureDefinitionCall1406: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFeatureDefinitionResponse-1617'
	if jsonDict["btType"] == "BTFeatureDefinitionResponse-1617" {
		// try to unmarshal JSON data into BTFeatureDefinitionResponse1617
		var qr *BTFeatureDefinitionResponse1617
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureApiBase1430 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureApiBase1430 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as BTFeatureDefinitionResponse1617: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFeatureListResponse-2457'
	if jsonDict["btType"] == "BTFeatureListResponse-2457" {
		// try to unmarshal JSON data into BTFeatureListResponse2457
		var qr *BTFeatureListResponse2457
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureApiBase1430 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureApiBase1430 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as BTFeatureListResponse2457: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFeatureScriptEvalCall-2377'
	if jsonDict["btType"] == "BTFeatureScriptEvalCall-2377" {
		// try to unmarshal JSON data into BTFeatureScriptEvalCall2377
		var qr *BTFeatureScriptEvalCall2377
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureApiBase1430 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureApiBase1430 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as BTFeatureScriptEvalCall2377: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFeatureScriptEvalResponse-1859'
	if jsonDict["btType"] == "BTFeatureScriptEvalResponse-1859" {
		// try to unmarshal JSON data into BTFeatureScriptEvalResponse1859
		var qr *BTFeatureScriptEvalResponse1859
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureApiBase1430 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureApiBase1430 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as BTFeatureScriptEvalResponse1859: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFeatureSpecsResponse-664'
	if jsonDict["btType"] == "BTFeatureSpecsResponse-664" {
		// try to unmarshal JSON data into BTFeatureSpecsResponse664
		var qr *BTFeatureSpecsResponse664
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureApiBase1430 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureApiBase1430 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as BTFeatureSpecsResponse664: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFeatureStudioContents-2239'
	if jsonDict["btType"] == "BTFeatureStudioContents-2239" {
		// try to unmarshal JSON data into BTFeatureStudioContents2239
		var qr *BTFeatureStudioContents2239
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureApiBase1430 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureApiBase1430 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as BTFeatureStudioContents2239: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSetFeatureRollbackCall-1899'
	if jsonDict["btType"] == "BTSetFeatureRollbackCall-1899" {
		// try to unmarshal JSON data into BTSetFeatureRollbackCall1899
		var qr *BTSetFeatureRollbackCall1899
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureApiBase1430 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureApiBase1430 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as BTSetFeatureRollbackCall1899: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSetFeatureRollbackResponse-1042'
	if jsonDict["btType"] == "BTSetFeatureRollbackResponse-1042" {
		// try to unmarshal JSON data into BTSetFeatureRollbackResponse1042
		var qr *BTSetFeatureRollbackResponse1042
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureApiBase1430 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureApiBase1430 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as BTSetFeatureRollbackResponse1042: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTUpdateFeaturesCall-1748'
	if jsonDict["btType"] == "BTUpdateFeaturesCall-1748" {
		// try to unmarshal JSON data into BTUpdateFeaturesCall1748
		var qr *BTUpdateFeaturesCall1748
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureApiBase1430 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureApiBase1430 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as BTUpdateFeaturesCall1748: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTUpdateFeaturesResponse-1333'
	if jsonDict["btType"] == "BTUpdateFeaturesResponse-1333" {
		// try to unmarshal JSON data into BTUpdateFeaturesResponse1333
		var qr *BTUpdateFeaturesResponse1333
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFeatureApiBase1430 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFeatureApiBase1430 = nil
			return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as BTUpdateFeaturesResponse1333: %s", err.Error())
		}
	}

	var qtx *base_BTFeatureApiBase1430
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTFeatureApiBase1430 = qtx
		return nil // data stored in dst.base_BTFeatureApiBase1430, return on the first match
	} else {
		dst.implBTFeatureApiBase1430 = nil
		return fmt.Errorf("Failed to unmarshal BTFeatureApiBase1430 as base_BTFeatureApiBase1430: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTFeatureApiBase1430) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTFeatureApiBase1430) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTFeatureApiBase1430
}

type NullableBTFeatureApiBase1430 struct {
	value *BTFeatureApiBase1430
	isSet bool
}

func (v NullableBTFeatureApiBase1430) Get() *BTFeatureApiBase1430 {
	return v.value
}

func (v *NullableBTFeatureApiBase1430) Set(val *BTFeatureApiBase1430) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureApiBase1430) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureApiBase1430) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureApiBase1430(val *BTFeatureApiBase1430) *NullableBTFeatureApiBase1430 {
	return &NullableBTFeatureApiBase1430{value: val, isSet: true}
}

func (v NullableBTFeatureApiBase1430) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureApiBase1430) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTFeatureApiBase1430 struct {
	LibraryVersion         *int32  `json:"libraryVersion,omitempty"`
	MicroversionSkew       *bool   `json:"microversionSkew,omitempty"`
	RejectMicroversionSkew *bool   `json:"rejectMicroversionSkew,omitempty"`
	SerializationVersion   *string `json:"serializationVersion,omitempty"`
	SourceMicroversion     *string `json:"sourceMicroversion,omitempty"`
}

// Newbase_BTFeatureApiBase1430 instantiates a new base_BTFeatureApiBase1430 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTFeatureApiBase1430() *base_BTFeatureApiBase1430 {
	this := base_BTFeatureApiBase1430{}
	return &this
}

// Newbase_BTFeatureApiBase1430WithDefaults instantiates a new base_BTFeatureApiBase1430 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTFeatureApiBase1430WithDefaults() *base_BTFeatureApiBase1430 {
	this := base_BTFeatureApiBase1430{}
	return &this
}

// GetLibraryVersion returns the LibraryVersion field value if set, zero value otherwise.
func (o *base_BTFeatureApiBase1430) GetLibraryVersion() int32 {
	if o == nil || o.LibraryVersion == nil {
		var ret int32
		return ret
	}
	return *o.LibraryVersion
}

// GetLibraryVersionOk returns a tuple with the LibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureApiBase1430) GetLibraryVersionOk() (*int32, bool) {
	if o == nil || o.LibraryVersion == nil {
		return nil, false
	}
	return o.LibraryVersion, true
}

// HasLibraryVersion returns a boolean if a field has been set.
func (o *base_BTFeatureApiBase1430) HasLibraryVersion() bool {
	if o != nil && o.LibraryVersion != nil {
		return true
	}

	return false
}

// SetLibraryVersion gets a reference to the given int32 and assigns it to the LibraryVersion field.
func (o *base_BTFeatureApiBase1430) SetLibraryVersion(v int32) {
	o.LibraryVersion = &v
}

// GetMicroversionSkew returns the MicroversionSkew field value if set, zero value otherwise.
func (o *base_BTFeatureApiBase1430) GetMicroversionSkew() bool {
	if o == nil || o.MicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.MicroversionSkew
}

// GetMicroversionSkewOk returns a tuple with the MicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureApiBase1430) GetMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.MicroversionSkew == nil {
		return nil, false
	}
	return o.MicroversionSkew, true
}

// HasMicroversionSkew returns a boolean if a field has been set.
func (o *base_BTFeatureApiBase1430) HasMicroversionSkew() bool {
	if o != nil && o.MicroversionSkew != nil {
		return true
	}

	return false
}

// SetMicroversionSkew gets a reference to the given bool and assigns it to the MicroversionSkew field.
func (o *base_BTFeatureApiBase1430) SetMicroversionSkew(v bool) {
	o.MicroversionSkew = &v
}

// GetRejectMicroversionSkew returns the RejectMicroversionSkew field value if set, zero value otherwise.
func (o *base_BTFeatureApiBase1430) GetRejectMicroversionSkew() bool {
	if o == nil || o.RejectMicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.RejectMicroversionSkew
}

// GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureApiBase1430) GetRejectMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.RejectMicroversionSkew == nil {
		return nil, false
	}
	return o.RejectMicroversionSkew, true
}

// HasRejectMicroversionSkew returns a boolean if a field has been set.
func (o *base_BTFeatureApiBase1430) HasRejectMicroversionSkew() bool {
	if o != nil && o.RejectMicroversionSkew != nil {
		return true
	}

	return false
}

// SetRejectMicroversionSkew gets a reference to the given bool and assigns it to the RejectMicroversionSkew field.
func (o *base_BTFeatureApiBase1430) SetRejectMicroversionSkew(v bool) {
	o.RejectMicroversionSkew = &v
}

// GetSerializationVersion returns the SerializationVersion field value if set, zero value otherwise.
func (o *base_BTFeatureApiBase1430) GetSerializationVersion() string {
	if o == nil || o.SerializationVersion == nil {
		var ret string
		return ret
	}
	return *o.SerializationVersion
}

// GetSerializationVersionOk returns a tuple with the SerializationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureApiBase1430) GetSerializationVersionOk() (*string, bool) {
	if o == nil || o.SerializationVersion == nil {
		return nil, false
	}
	return o.SerializationVersion, true
}

// HasSerializationVersion returns a boolean if a field has been set.
func (o *base_BTFeatureApiBase1430) HasSerializationVersion() bool {
	if o != nil && o.SerializationVersion != nil {
		return true
	}

	return false
}

// SetSerializationVersion gets a reference to the given string and assigns it to the SerializationVersion field.
func (o *base_BTFeatureApiBase1430) SetSerializationVersion(v string) {
	o.SerializationVersion = &v
}

// GetSourceMicroversion returns the SourceMicroversion field value if set, zero value otherwise.
func (o *base_BTFeatureApiBase1430) GetSourceMicroversion() string {
	if o == nil || o.SourceMicroversion == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversion
}

// GetSourceMicroversionOk returns a tuple with the SourceMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFeatureApiBase1430) GetSourceMicroversionOk() (*string, bool) {
	if o == nil || o.SourceMicroversion == nil {
		return nil, false
	}
	return o.SourceMicroversion, true
}

// HasSourceMicroversion returns a boolean if a field has been set.
func (o *base_BTFeatureApiBase1430) HasSourceMicroversion() bool {
	if o != nil && o.SourceMicroversion != nil {
		return true
	}

	return false
}

// SetSourceMicroversion gets a reference to the given string and assigns it to the SourceMicroversion field.
func (o *base_BTFeatureApiBase1430) SetSourceMicroversion(v string) {
	o.SourceMicroversion = &v
}

func (o base_BTFeatureApiBase1430) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LibraryVersion != nil {
		toSerialize["libraryVersion"] = o.LibraryVersion
	}
	if o.MicroversionSkew != nil {
		toSerialize["microversionSkew"] = o.MicroversionSkew
	}
	if o.RejectMicroversionSkew != nil {
		toSerialize["rejectMicroversionSkew"] = o.RejectMicroversionSkew
	}
	if o.SerializationVersion != nil {
		toSerialize["serializationVersion"] = o.SerializationVersion
	}
	if o.SourceMicroversion != nil {
		toSerialize["sourceMicroversion"] = o.SourceMicroversion
	}
	return json.Marshal(toSerialize)
}
