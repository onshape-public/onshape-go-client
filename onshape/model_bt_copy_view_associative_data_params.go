/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.171.24440-f37a82fd6450
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCopyViewAssociativeDataParams struct for BTCopyViewAssociativeDataParams
type BTCopyViewAssociativeDataParams struct {
	AssociativeDataIds []string `json:"associativeDataIds,omitempty"`
	DestinationViewId  *string  `json:"destinationViewId,omitempty"`
	SourceElementId    *string  `json:"sourceElementId,omitempty"`
	SourceViewId       *string  `json:"sourceViewId,omitempty"`
}

// NewBTCopyViewAssociativeDataParams instantiates a new BTCopyViewAssociativeDataParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCopyViewAssociativeDataParams() *BTCopyViewAssociativeDataParams {
	this := BTCopyViewAssociativeDataParams{}
	return &this
}

// NewBTCopyViewAssociativeDataParamsWithDefaults instantiates a new BTCopyViewAssociativeDataParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCopyViewAssociativeDataParamsWithDefaults() *BTCopyViewAssociativeDataParams {
	this := BTCopyViewAssociativeDataParams{}
	return &this
}

// GetAssociativeDataIds returns the AssociativeDataIds field value if set, zero value otherwise.
func (o *BTCopyViewAssociativeDataParams) GetAssociativeDataIds() []string {
	if o == nil || o.AssociativeDataIds == nil {
		var ret []string
		return ret
	}
	return o.AssociativeDataIds
}

// GetAssociativeDataIdsOk returns a tuple with the AssociativeDataIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyViewAssociativeDataParams) GetAssociativeDataIdsOk() ([]string, bool) {
	if o == nil || o.AssociativeDataIds == nil {
		return nil, false
	}
	return o.AssociativeDataIds, true
}

// HasAssociativeDataIds returns a boolean if a field has been set.
func (o *BTCopyViewAssociativeDataParams) HasAssociativeDataIds() bool {
	if o != nil && o.AssociativeDataIds != nil {
		return true
	}

	return false
}

// SetAssociativeDataIds gets a reference to the given []string and assigns it to the AssociativeDataIds field.
func (o *BTCopyViewAssociativeDataParams) SetAssociativeDataIds(v []string) {
	o.AssociativeDataIds = v
}

// GetDestinationViewId returns the DestinationViewId field value if set, zero value otherwise.
func (o *BTCopyViewAssociativeDataParams) GetDestinationViewId() string {
	if o == nil || o.DestinationViewId == nil {
		var ret string
		return ret
	}
	return *o.DestinationViewId
}

// GetDestinationViewIdOk returns a tuple with the DestinationViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyViewAssociativeDataParams) GetDestinationViewIdOk() (*string, bool) {
	if o == nil || o.DestinationViewId == nil {
		return nil, false
	}
	return o.DestinationViewId, true
}

// HasDestinationViewId returns a boolean if a field has been set.
func (o *BTCopyViewAssociativeDataParams) HasDestinationViewId() bool {
	if o != nil && o.DestinationViewId != nil {
		return true
	}

	return false
}

// SetDestinationViewId gets a reference to the given string and assigns it to the DestinationViewId field.
func (o *BTCopyViewAssociativeDataParams) SetDestinationViewId(v string) {
	o.DestinationViewId = &v
}

// GetSourceElementId returns the SourceElementId field value if set, zero value otherwise.
func (o *BTCopyViewAssociativeDataParams) GetSourceElementId() string {
	if o == nil || o.SourceElementId == nil {
		var ret string
		return ret
	}
	return *o.SourceElementId
}

// GetSourceElementIdOk returns a tuple with the SourceElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyViewAssociativeDataParams) GetSourceElementIdOk() (*string, bool) {
	if o == nil || o.SourceElementId == nil {
		return nil, false
	}
	return o.SourceElementId, true
}

// HasSourceElementId returns a boolean if a field has been set.
func (o *BTCopyViewAssociativeDataParams) HasSourceElementId() bool {
	if o != nil && o.SourceElementId != nil {
		return true
	}

	return false
}

// SetSourceElementId gets a reference to the given string and assigns it to the SourceElementId field.
func (o *BTCopyViewAssociativeDataParams) SetSourceElementId(v string) {
	o.SourceElementId = &v
}

// GetSourceViewId returns the SourceViewId field value if set, zero value otherwise.
func (o *BTCopyViewAssociativeDataParams) GetSourceViewId() string {
	if o == nil || o.SourceViewId == nil {
		var ret string
		return ret
	}
	return *o.SourceViewId
}

// GetSourceViewIdOk returns a tuple with the SourceViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyViewAssociativeDataParams) GetSourceViewIdOk() (*string, bool) {
	if o == nil || o.SourceViewId == nil {
		return nil, false
	}
	return o.SourceViewId, true
}

// HasSourceViewId returns a boolean if a field has been set.
func (o *BTCopyViewAssociativeDataParams) HasSourceViewId() bool {
	if o != nil && o.SourceViewId != nil {
		return true
	}

	return false
}

// SetSourceViewId gets a reference to the given string and assigns it to the SourceViewId field.
func (o *BTCopyViewAssociativeDataParams) SetSourceViewId(v string) {
	o.SourceViewId = &v
}

func (o BTCopyViewAssociativeDataParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssociativeDataIds != nil {
		toSerialize["associativeDataIds"] = o.AssociativeDataIds
	}
	if o.DestinationViewId != nil {
		toSerialize["destinationViewId"] = o.DestinationViewId
	}
	if o.SourceElementId != nil {
		toSerialize["sourceElementId"] = o.SourceElementId
	}
	if o.SourceViewId != nil {
		toSerialize["sourceViewId"] = o.SourceViewId
	}
	return json.Marshal(toSerialize)
}

type NullableBTCopyViewAssociativeDataParams struct {
	value *BTCopyViewAssociativeDataParams
	isSet bool
}

func (v NullableBTCopyViewAssociativeDataParams) Get() *BTCopyViewAssociativeDataParams {
	return v.value
}

func (v *NullableBTCopyViewAssociativeDataParams) Set(val *BTCopyViewAssociativeDataParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCopyViewAssociativeDataParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCopyViewAssociativeDataParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCopyViewAssociativeDataParams(val *BTCopyViewAssociativeDataParams) *NullableBTCopyViewAssociativeDataParams {
	return &NullableBTCopyViewAssociativeDataParams{value: val, isSet: true}
}

func (v NullableBTCopyViewAssociativeDataParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCopyViewAssociativeDataParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
