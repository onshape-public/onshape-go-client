/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6946-bb9367a9d0cc
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTWebRendererPerformanceMeasurementParams struct for BTWebRendererPerformanceMeasurementParams
type BTWebRendererPerformanceMeasurementParams struct {
	LinesPerSecond     *float32 `json:"linesPerSecond,omitempty"`
	Renderer           *string  `json:"renderer,omitempty"`
	TrianglesPerSecond *float32 `json:"trianglesPerSecond,omitempty"`
	Vendor             *string  `json:"vendor,omitempty"`
}

// NewBTWebRendererPerformanceMeasurementParams instantiates a new BTWebRendererPerformanceMeasurementParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWebRendererPerformanceMeasurementParams() *BTWebRendererPerformanceMeasurementParams {
	this := BTWebRendererPerformanceMeasurementParams{}
	return &this
}

// NewBTWebRendererPerformanceMeasurementParamsWithDefaults instantiates a new BTWebRendererPerformanceMeasurementParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWebRendererPerformanceMeasurementParamsWithDefaults() *BTWebRendererPerformanceMeasurementParams {
	this := BTWebRendererPerformanceMeasurementParams{}
	return &this
}

// GetLinesPerSecond returns the LinesPerSecond field value if set, zero value otherwise.
func (o *BTWebRendererPerformanceMeasurementParams) GetLinesPerSecond() float32 {
	if o == nil || o.LinesPerSecond == nil {
		var ret float32
		return ret
	}
	return *o.LinesPerSecond
}

// GetLinesPerSecondOk returns a tuple with the LinesPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebRendererPerformanceMeasurementParams) GetLinesPerSecondOk() (*float32, bool) {
	if o == nil || o.LinesPerSecond == nil {
		return nil, false
	}
	return o.LinesPerSecond, true
}

// HasLinesPerSecond returns a boolean if a field has been set.
func (o *BTWebRendererPerformanceMeasurementParams) HasLinesPerSecond() bool {
	if o != nil && o.LinesPerSecond != nil {
		return true
	}

	return false
}

// SetLinesPerSecond gets a reference to the given float32 and assigns it to the LinesPerSecond field.
func (o *BTWebRendererPerformanceMeasurementParams) SetLinesPerSecond(v float32) {
	o.LinesPerSecond = &v
}

// GetRenderer returns the Renderer field value if set, zero value otherwise.
func (o *BTWebRendererPerformanceMeasurementParams) GetRenderer() string {
	if o == nil || o.Renderer == nil {
		var ret string
		return ret
	}
	return *o.Renderer
}

// GetRendererOk returns a tuple with the Renderer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebRendererPerformanceMeasurementParams) GetRendererOk() (*string, bool) {
	if o == nil || o.Renderer == nil {
		return nil, false
	}
	return o.Renderer, true
}

// HasRenderer returns a boolean if a field has been set.
func (o *BTWebRendererPerformanceMeasurementParams) HasRenderer() bool {
	if o != nil && o.Renderer != nil {
		return true
	}

	return false
}

// SetRenderer gets a reference to the given string and assigns it to the Renderer field.
func (o *BTWebRendererPerformanceMeasurementParams) SetRenderer(v string) {
	o.Renderer = &v
}

// GetTrianglesPerSecond returns the TrianglesPerSecond field value if set, zero value otherwise.
func (o *BTWebRendererPerformanceMeasurementParams) GetTrianglesPerSecond() float32 {
	if o == nil || o.TrianglesPerSecond == nil {
		var ret float32
		return ret
	}
	return *o.TrianglesPerSecond
}

// GetTrianglesPerSecondOk returns a tuple with the TrianglesPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebRendererPerformanceMeasurementParams) GetTrianglesPerSecondOk() (*float32, bool) {
	if o == nil || o.TrianglesPerSecond == nil {
		return nil, false
	}
	return o.TrianglesPerSecond, true
}

// HasTrianglesPerSecond returns a boolean if a field has been set.
func (o *BTWebRendererPerformanceMeasurementParams) HasTrianglesPerSecond() bool {
	if o != nil && o.TrianglesPerSecond != nil {
		return true
	}

	return false
}

// SetTrianglesPerSecond gets a reference to the given float32 and assigns it to the TrianglesPerSecond field.
func (o *BTWebRendererPerformanceMeasurementParams) SetTrianglesPerSecond(v float32) {
	o.TrianglesPerSecond = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *BTWebRendererPerformanceMeasurementParams) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebRendererPerformanceMeasurementParams) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *BTWebRendererPerformanceMeasurementParams) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *BTWebRendererPerformanceMeasurementParams) SetVendor(v string) {
	o.Vendor = &v
}

func (o BTWebRendererPerformanceMeasurementParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LinesPerSecond != nil {
		toSerialize["linesPerSecond"] = o.LinesPerSecond
	}
	if o.Renderer != nil {
		toSerialize["renderer"] = o.Renderer
	}
	if o.TrianglesPerSecond != nil {
		toSerialize["trianglesPerSecond"] = o.TrianglesPerSecond
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableBTWebRendererPerformanceMeasurementParams struct {
	value *BTWebRendererPerformanceMeasurementParams
	isSet bool
}

func (v NullableBTWebRendererPerformanceMeasurementParams) Get() *BTWebRendererPerformanceMeasurementParams {
	return v.value
}

func (v *NullableBTWebRendererPerformanceMeasurementParams) Set(val *BTWebRendererPerformanceMeasurementParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWebRendererPerformanceMeasurementParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWebRendererPerformanceMeasurementParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWebRendererPerformanceMeasurementParams(val *BTWebRendererPerformanceMeasurementParams) *NullableBTWebRendererPerformanceMeasurementParams {
	return &NullableBTWebRendererPerformanceMeasurementParams{value: val, isSet: true}
}

func (v NullableBTWebRendererPerformanceMeasurementParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWebRendererPerformanceMeasurementParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
