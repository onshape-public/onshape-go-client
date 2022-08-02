/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5843-26f51e563fa4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTRevisionListResponse struct for BTRevisionListResponse
type BTRevisionListResponse struct {
	Href       *string          `json:"href,omitempty"`
	Items      []BTRevisionInfo `json:"items,omitempty"`
	Next       *string          `json:"next,omitempty"`
	PartNumber *string          `json:"partNumber,omitempty"`
	Previous   *string          `json:"previous,omitempty"`
}

// NewBTRevisionListResponse instantiates a new BTRevisionListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRevisionListResponse() *BTRevisionListResponse {
	this := BTRevisionListResponse{}
	return &this
}

// NewBTRevisionListResponseWithDefaults instantiates a new BTRevisionListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRevisionListResponseWithDefaults() *BTRevisionListResponse {
	this := BTRevisionListResponse{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTRevisionListResponse) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionListResponse) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTRevisionListResponse) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTRevisionListResponse) SetHref(v string) {
	o.Href = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BTRevisionListResponse) GetItems() []BTRevisionInfo {
	if o == nil || o.Items == nil {
		var ret []BTRevisionInfo
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionListResponse) GetItemsOk() ([]BTRevisionInfo, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BTRevisionListResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BTRevisionInfo and assigns it to the Items field.
func (o *BTRevisionListResponse) SetItems(v []BTRevisionInfo) {
	o.Items = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *BTRevisionListResponse) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionListResponse) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *BTRevisionListResponse) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *BTRevisionListResponse) SetNext(v string) {
	o.Next = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTRevisionListResponse) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionListResponse) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTRevisionListResponse) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTRevisionListResponse) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *BTRevisionListResponse) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionListResponse) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *BTRevisionListResponse) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *BTRevisionListResponse) SetPrevious(v string) {
	o.Previous = &v
}

func (o BTRevisionListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	return json.Marshal(toSerialize)
}

type NullableBTRevisionListResponse struct {
	value *BTRevisionListResponse
	isSet bool
}

func (v NullableBTRevisionListResponse) Get() *BTRevisionListResponse {
	return v.value
}

func (v *NullableBTRevisionListResponse) Set(val *BTRevisionListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRevisionListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRevisionListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRevisionListResponse(val *BTRevisionListResponse) *NullableBTRevisionListResponse {
	return &NullableBTRevisionListResponse{value: val, isSet: true}
}

func (v NullableBTRevisionListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRevisionListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
