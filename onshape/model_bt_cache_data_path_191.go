/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.171.24257-687de06de652
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCacheDataPath191 struct for BTCacheDataPath191
type BTCacheDataPath191 struct {
	BtType                           *string `json:"btType,omitempty"`
	DocumentId                       *string `json:"documentId,omitempty"`
	ElementId                        *string `json:"elementId,omitempty"`
	ImmutablePathContentsShouldExist *bool   `json:"immutablePathContentsShouldExist,omitempty"`
	IsImmutableContextPath           *bool   `json:"isImmutableContextPath,omitempty"`
	Key                              *string `json:"key,omitempty"`
	KeyContainsConfiguration         *bool   `json:"keyContainsConfiguration,omitempty"`
	UseLocalFileCache                *bool   `json:"useLocalFileCache,omitempty"`
}

// NewBTCacheDataPath191 instantiates a new BTCacheDataPath191 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCacheDataPath191() *BTCacheDataPath191 {
	this := BTCacheDataPath191{}
	return &this
}

// NewBTCacheDataPath191WithDefaults instantiates a new BTCacheDataPath191 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCacheDataPath191WithDefaults() *BTCacheDataPath191 {
	this := BTCacheDataPath191{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTCacheDataPath191) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTCacheDataPath191) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTCacheDataPath191) SetElementId(v string) {
	o.ElementId = &v
}

// GetImmutablePathContentsShouldExist returns the ImmutablePathContentsShouldExist field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetImmutablePathContentsShouldExist() bool {
	if o == nil || o.ImmutablePathContentsShouldExist == nil {
		var ret bool
		return ret
	}
	return *o.ImmutablePathContentsShouldExist
}

// GetImmutablePathContentsShouldExistOk returns a tuple with the ImmutablePathContentsShouldExist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetImmutablePathContentsShouldExistOk() (*bool, bool) {
	if o == nil || o.ImmutablePathContentsShouldExist == nil {
		return nil, false
	}
	return o.ImmutablePathContentsShouldExist, true
}

// HasImmutablePathContentsShouldExist returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasImmutablePathContentsShouldExist() bool {
	if o != nil && o.ImmutablePathContentsShouldExist != nil {
		return true
	}

	return false
}

// SetImmutablePathContentsShouldExist gets a reference to the given bool and assigns it to the ImmutablePathContentsShouldExist field.
func (o *BTCacheDataPath191) SetImmutablePathContentsShouldExist(v bool) {
	o.ImmutablePathContentsShouldExist = &v
}

// GetIsImmutableContextPath returns the IsImmutableContextPath field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetIsImmutableContextPath() bool {
	if o == nil || o.IsImmutableContextPath == nil {
		var ret bool
		return ret
	}
	return *o.IsImmutableContextPath
}

// GetIsImmutableContextPathOk returns a tuple with the IsImmutableContextPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetIsImmutableContextPathOk() (*bool, bool) {
	if o == nil || o.IsImmutableContextPath == nil {
		return nil, false
	}
	return o.IsImmutableContextPath, true
}

// HasIsImmutableContextPath returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasIsImmutableContextPath() bool {
	if o != nil && o.IsImmutableContextPath != nil {
		return true
	}

	return false
}

// SetIsImmutableContextPath gets a reference to the given bool and assigns it to the IsImmutableContextPath field.
func (o *BTCacheDataPath191) SetIsImmutableContextPath(v bool) {
	o.IsImmutableContextPath = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *BTCacheDataPath191) SetKey(v string) {
	o.Key = &v
}

// GetKeyContainsConfiguration returns the KeyContainsConfiguration field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetKeyContainsConfiguration() bool {
	if o == nil || o.KeyContainsConfiguration == nil {
		var ret bool
		return ret
	}
	return *o.KeyContainsConfiguration
}

// GetKeyContainsConfigurationOk returns a tuple with the KeyContainsConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetKeyContainsConfigurationOk() (*bool, bool) {
	if o == nil || o.KeyContainsConfiguration == nil {
		return nil, false
	}
	return o.KeyContainsConfiguration, true
}

// HasKeyContainsConfiguration returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasKeyContainsConfiguration() bool {
	if o != nil && o.KeyContainsConfiguration != nil {
		return true
	}

	return false
}

// SetKeyContainsConfiguration gets a reference to the given bool and assigns it to the KeyContainsConfiguration field.
func (o *BTCacheDataPath191) SetKeyContainsConfiguration(v bool) {
	o.KeyContainsConfiguration = &v
}

// GetUseLocalFileCache returns the UseLocalFileCache field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetUseLocalFileCache() bool {
	if o == nil || o.UseLocalFileCache == nil {
		var ret bool
		return ret
	}
	return *o.UseLocalFileCache
}

// GetUseLocalFileCacheOk returns a tuple with the UseLocalFileCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetUseLocalFileCacheOk() (*bool, bool) {
	if o == nil || o.UseLocalFileCache == nil {
		return nil, false
	}
	return o.UseLocalFileCache, true
}

// HasUseLocalFileCache returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasUseLocalFileCache() bool {
	if o != nil && o.UseLocalFileCache != nil {
		return true
	}

	return false
}

// SetUseLocalFileCache gets a reference to the given bool and assigns it to the UseLocalFileCache field.
func (o *BTCacheDataPath191) SetUseLocalFileCache(v bool) {
	o.UseLocalFileCache = &v
}

func (o BTCacheDataPath191) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ImmutablePathContentsShouldExist != nil {
		toSerialize["immutablePathContentsShouldExist"] = o.ImmutablePathContentsShouldExist
	}
	if o.IsImmutableContextPath != nil {
		toSerialize["isImmutableContextPath"] = o.IsImmutableContextPath
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.KeyContainsConfiguration != nil {
		toSerialize["keyContainsConfiguration"] = o.KeyContainsConfiguration
	}
	if o.UseLocalFileCache != nil {
		toSerialize["useLocalFileCache"] = o.UseLocalFileCache
	}
	return json.Marshal(toSerialize)
}

type NullableBTCacheDataPath191 struct {
	value *BTCacheDataPath191
	isSet bool
}

func (v NullableBTCacheDataPath191) Get() *BTCacheDataPath191 {
	return v.value
}

func (v *NullableBTCacheDataPath191) Set(val *BTCacheDataPath191) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCacheDataPath191) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCacheDataPath191) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCacheDataPath191(val *BTCacheDataPath191) *NullableBTCacheDataPath191 {
	return &NullableBTCacheDataPath191{value: val, isSet: true}
}

func (v NullableBTCacheDataPath191) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCacheDataPath191) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
