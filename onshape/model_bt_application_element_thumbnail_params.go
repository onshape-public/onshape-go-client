/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.156.8014-206b7d55b208
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTApplicationElementThumbnailParams struct for BTApplicationElementThumbnailParams
type BTApplicationElementThumbnailParams struct {
	Base64EncodedImage *string `json:"base64EncodedImage,omitempty"`
	Description        *string `json:"description,omitempty"`
	ImageHeight        *int32  `json:"imageHeight,omitempty"`
	ImageWidth         *int32  `json:"imageWidth,omitempty"`
	IsPrimary          *bool   `json:"isPrimary,omitempty"`
	UniqueId           *string `json:"uniqueId,omitempty"`
}

// NewBTApplicationElementThumbnailParams instantiates a new BTApplicationElementThumbnailParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTApplicationElementThumbnailParams() *BTApplicationElementThumbnailParams {
	this := BTApplicationElementThumbnailParams{}
	return &this
}

// NewBTApplicationElementThumbnailParamsWithDefaults instantiates a new BTApplicationElementThumbnailParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTApplicationElementThumbnailParamsWithDefaults() *BTApplicationElementThumbnailParams {
	this := BTApplicationElementThumbnailParams{}
	return &this
}

// GetBase64EncodedImage returns the Base64EncodedImage field value if set, zero value otherwise.
func (o *BTApplicationElementThumbnailParams) GetBase64EncodedImage() string {
	if o == nil || o.Base64EncodedImage == nil {
		var ret string
		return ret
	}
	return *o.Base64EncodedImage
}

// GetBase64EncodedImageOk returns a tuple with the Base64EncodedImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApplicationElementThumbnailParams) GetBase64EncodedImageOk() (*string, bool) {
	if o == nil || o.Base64EncodedImage == nil {
		return nil, false
	}
	return o.Base64EncodedImage, true
}

// HasBase64EncodedImage returns a boolean if a field has been set.
func (o *BTApplicationElementThumbnailParams) HasBase64EncodedImage() bool {
	if o != nil && o.Base64EncodedImage != nil {
		return true
	}

	return false
}

// SetBase64EncodedImage gets a reference to the given string and assigns it to the Base64EncodedImage field.
func (o *BTApplicationElementThumbnailParams) SetBase64EncodedImage(v string) {
	o.Base64EncodedImage = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTApplicationElementThumbnailParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApplicationElementThumbnailParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTApplicationElementThumbnailParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTApplicationElementThumbnailParams) SetDescription(v string) {
	o.Description = &v
}

// GetImageHeight returns the ImageHeight field value if set, zero value otherwise.
func (o *BTApplicationElementThumbnailParams) GetImageHeight() int32 {
	if o == nil || o.ImageHeight == nil {
		var ret int32
		return ret
	}
	return *o.ImageHeight
}

// GetImageHeightOk returns a tuple with the ImageHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApplicationElementThumbnailParams) GetImageHeightOk() (*int32, bool) {
	if o == nil || o.ImageHeight == nil {
		return nil, false
	}
	return o.ImageHeight, true
}

// HasImageHeight returns a boolean if a field has been set.
func (o *BTApplicationElementThumbnailParams) HasImageHeight() bool {
	if o != nil && o.ImageHeight != nil {
		return true
	}

	return false
}

// SetImageHeight gets a reference to the given int32 and assigns it to the ImageHeight field.
func (o *BTApplicationElementThumbnailParams) SetImageHeight(v int32) {
	o.ImageHeight = &v
}

// GetImageWidth returns the ImageWidth field value if set, zero value otherwise.
func (o *BTApplicationElementThumbnailParams) GetImageWidth() int32 {
	if o == nil || o.ImageWidth == nil {
		var ret int32
		return ret
	}
	return *o.ImageWidth
}

// GetImageWidthOk returns a tuple with the ImageWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApplicationElementThumbnailParams) GetImageWidthOk() (*int32, bool) {
	if o == nil || o.ImageWidth == nil {
		return nil, false
	}
	return o.ImageWidth, true
}

// HasImageWidth returns a boolean if a field has been set.
func (o *BTApplicationElementThumbnailParams) HasImageWidth() bool {
	if o != nil && o.ImageWidth != nil {
		return true
	}

	return false
}

// SetImageWidth gets a reference to the given int32 and assigns it to the ImageWidth field.
func (o *BTApplicationElementThumbnailParams) SetImageWidth(v int32) {
	o.ImageWidth = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *BTApplicationElementThumbnailParams) GetIsPrimary() bool {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApplicationElementThumbnailParams) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || o.IsPrimary == nil {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *BTApplicationElementThumbnailParams) HasIsPrimary() bool {
	if o != nil && o.IsPrimary != nil {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *BTApplicationElementThumbnailParams) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise.
func (o *BTApplicationElementThumbnailParams) GetUniqueId() string {
	if o == nil || o.UniqueId == nil {
		var ret string
		return ret
	}
	return *o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApplicationElementThumbnailParams) GetUniqueIdOk() (*string, bool) {
	if o == nil || o.UniqueId == nil {
		return nil, false
	}
	return o.UniqueId, true
}

// HasUniqueId returns a boolean if a field has been set.
func (o *BTApplicationElementThumbnailParams) HasUniqueId() bool {
	if o != nil && o.UniqueId != nil {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given string and assigns it to the UniqueId field.
func (o *BTApplicationElementThumbnailParams) SetUniqueId(v string) {
	o.UniqueId = &v
}

func (o BTApplicationElementThumbnailParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Base64EncodedImage != nil {
		toSerialize["base64EncodedImage"] = o.Base64EncodedImage
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ImageHeight != nil {
		toSerialize["imageHeight"] = o.ImageHeight
	}
	if o.ImageWidth != nil {
		toSerialize["imageWidth"] = o.ImageWidth
	}
	if o.IsPrimary != nil {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	if o.UniqueId != nil {
		toSerialize["uniqueId"] = o.UniqueId
	}
	return json.Marshal(toSerialize)
}

type NullableBTApplicationElementThumbnailParams struct {
	value *BTApplicationElementThumbnailParams
	isSet bool
}

func (v NullableBTApplicationElementThumbnailParams) Get() *BTApplicationElementThumbnailParams {
	return v.value
}

func (v *NullableBTApplicationElementThumbnailParams) Set(val *BTApplicationElementThumbnailParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTApplicationElementThumbnailParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTApplicationElementThumbnailParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTApplicationElementThumbnailParams(val *BTApplicationElementThumbnailParams) *NullableBTApplicationElementThumbnailParams {
	return &NullableBTApplicationElementThumbnailParams{value: val, isSet: true}
}

func (v NullableBTApplicationElementThumbnailParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTApplicationElementThumbnailParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
