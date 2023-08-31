/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.168.21279-402b6292597b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCommentAttachmentInfo struct for BTCommentAttachmentInfo
type BTCommentAttachmentInfo struct {
	FileName *string `json:"fileName,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id       *string `json:"id,omitempty"`
	MimeType *string `json:"mimeType,omitempty"`
	// Name of the resource.
	Name         *string `json:"name,omitempty"`
	ThumbnailFor *string `json:"thumbnailFor,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTCommentAttachmentInfo instantiates a new BTCommentAttachmentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCommentAttachmentInfo() *BTCommentAttachmentInfo {
	this := BTCommentAttachmentInfo{}
	return &this
}

// NewBTCommentAttachmentInfoWithDefaults instantiates a new BTCommentAttachmentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCommentAttachmentInfoWithDefaults() *BTCommentAttachmentInfo {
	this := BTCommentAttachmentInfo{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *BTCommentAttachmentInfo) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentAttachmentInfo) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *BTCommentAttachmentInfo) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *BTCommentAttachmentInfo) SetFileName(v string) {
	o.FileName = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTCommentAttachmentInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentAttachmentInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTCommentAttachmentInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTCommentAttachmentInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTCommentAttachmentInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentAttachmentInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTCommentAttachmentInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTCommentAttachmentInfo) SetId(v string) {
	o.Id = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *BTCommentAttachmentInfo) GetMimeType() string {
	if o == nil || o.MimeType == nil {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentAttachmentInfo) GetMimeTypeOk() (*string, bool) {
	if o == nil || o.MimeType == nil {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *BTCommentAttachmentInfo) HasMimeType() bool {
	if o != nil && o.MimeType != nil {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *BTCommentAttachmentInfo) SetMimeType(v string) {
	o.MimeType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTCommentAttachmentInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentAttachmentInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTCommentAttachmentInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTCommentAttachmentInfo) SetName(v string) {
	o.Name = &v
}

// GetThumbnailFor returns the ThumbnailFor field value if set, zero value otherwise.
func (o *BTCommentAttachmentInfo) GetThumbnailFor() string {
	if o == nil || o.ThumbnailFor == nil {
		var ret string
		return ret
	}
	return *o.ThumbnailFor
}

// GetThumbnailForOk returns a tuple with the ThumbnailFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentAttachmentInfo) GetThumbnailForOk() (*string, bool) {
	if o == nil || o.ThumbnailFor == nil {
		return nil, false
	}
	return o.ThumbnailFor, true
}

// HasThumbnailFor returns a boolean if a field has been set.
func (o *BTCommentAttachmentInfo) HasThumbnailFor() bool {
	if o != nil && o.ThumbnailFor != nil {
		return true
	}

	return false
}

// SetThumbnailFor gets a reference to the given string and assigns it to the ThumbnailFor field.
func (o *BTCommentAttachmentInfo) SetThumbnailFor(v string) {
	o.ThumbnailFor = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTCommentAttachmentInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentAttachmentInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTCommentAttachmentInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTCommentAttachmentInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTCommentAttachmentInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FileName != nil {
		toSerialize["fileName"] = o.FileName
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MimeType != nil {
		toSerialize["mimeType"] = o.MimeType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ThumbnailFor != nil {
		toSerialize["thumbnailFor"] = o.ThumbnailFor
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTCommentAttachmentInfo struct {
	value *BTCommentAttachmentInfo
	isSet bool
}

func (v NullableBTCommentAttachmentInfo) Get() *BTCommentAttachmentInfo {
	return v.value
}

func (v *NullableBTCommentAttachmentInfo) Set(val *BTCommentAttachmentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCommentAttachmentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCommentAttachmentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCommentAttachmentInfo(val *BTCommentAttachmentInfo) *NullableBTCommentAttachmentInfo {
	return &NullableBTCommentAttachmentInfo{value: val, isSet: true}
}

func (v NullableBTCommentAttachmentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCommentAttachmentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
