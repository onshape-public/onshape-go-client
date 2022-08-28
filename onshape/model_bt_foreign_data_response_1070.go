/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6251-33aac52a3d4e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTForeignDataResponse1070 struct for BTForeignDataResponse1070
type BTForeignDataResponse1070 struct {
	BtType          *string  `json:"btType,omitempty"`
	BucketName      *string  `json:"bucketName,omitempty"`
	BucketPath      *string  `json:"bucketPath,omitempty"`
	CacheChunkList  []string `json:"cacheChunkList,omitempty"`
	DataId          *string  `json:"dataId,omitempty"`
	Format          *string  `json:"format,omitempty"`
	Name            *string  `json:"name,omitempty"`
	Region          *string  `json:"region,omitempty"`
	Size            *int32   `json:"size,omitempty"`
	StorageType     *string  `json:"storageType,omitempty"`
	UseLocalStorage *bool    `json:"useLocalStorage,omitempty"`
}

// NewBTForeignDataResponse1070 instantiates a new BTForeignDataResponse1070 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTForeignDataResponse1070() *BTForeignDataResponse1070 {
	this := BTForeignDataResponse1070{}
	return &this
}

// NewBTForeignDataResponse1070WithDefaults instantiates a new BTForeignDataResponse1070 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTForeignDataResponse1070WithDefaults() *BTForeignDataResponse1070 {
	this := BTForeignDataResponse1070{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTForeignDataResponse1070) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTForeignDataResponse1070) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTForeignDataResponse1070) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTForeignDataResponse1070) SetBtType(v string) {
	o.BtType = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *BTForeignDataResponse1070) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTForeignDataResponse1070) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *BTForeignDataResponse1070) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *BTForeignDataResponse1070) SetBucketName(v string) {
	o.BucketName = &v
}

// GetBucketPath returns the BucketPath field value if set, zero value otherwise.
func (o *BTForeignDataResponse1070) GetBucketPath() string {
	if o == nil || o.BucketPath == nil {
		var ret string
		return ret
	}
	return *o.BucketPath
}

// GetBucketPathOk returns a tuple with the BucketPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTForeignDataResponse1070) GetBucketPathOk() (*string, bool) {
	if o == nil || o.BucketPath == nil {
		return nil, false
	}
	return o.BucketPath, true
}

// HasBucketPath returns a boolean if a field has been set.
func (o *BTForeignDataResponse1070) HasBucketPath() bool {
	if o != nil && o.BucketPath != nil {
		return true
	}

	return false
}

// SetBucketPath gets a reference to the given string and assigns it to the BucketPath field.
func (o *BTForeignDataResponse1070) SetBucketPath(v string) {
	o.BucketPath = &v
}

// GetCacheChunkList returns the CacheChunkList field value if set, zero value otherwise.
func (o *BTForeignDataResponse1070) GetCacheChunkList() []string {
	if o == nil || o.CacheChunkList == nil {
		var ret []string
		return ret
	}
	return o.CacheChunkList
}

// GetCacheChunkListOk returns a tuple with the CacheChunkList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTForeignDataResponse1070) GetCacheChunkListOk() ([]string, bool) {
	if o == nil || o.CacheChunkList == nil {
		return nil, false
	}
	return o.CacheChunkList, true
}

// HasCacheChunkList returns a boolean if a field has been set.
func (o *BTForeignDataResponse1070) HasCacheChunkList() bool {
	if o != nil && o.CacheChunkList != nil {
		return true
	}

	return false
}

// SetCacheChunkList gets a reference to the given []string and assigns it to the CacheChunkList field.
func (o *BTForeignDataResponse1070) SetCacheChunkList(v []string) {
	o.CacheChunkList = v
}

// GetDataId returns the DataId field value if set, zero value otherwise.
func (o *BTForeignDataResponse1070) GetDataId() string {
	if o == nil || o.DataId == nil {
		var ret string
		return ret
	}
	return *o.DataId
}

// GetDataIdOk returns a tuple with the DataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTForeignDataResponse1070) GetDataIdOk() (*string, bool) {
	if o == nil || o.DataId == nil {
		return nil, false
	}
	return o.DataId, true
}

// HasDataId returns a boolean if a field has been set.
func (o *BTForeignDataResponse1070) HasDataId() bool {
	if o != nil && o.DataId != nil {
		return true
	}

	return false
}

// SetDataId gets a reference to the given string and assigns it to the DataId field.
func (o *BTForeignDataResponse1070) SetDataId(v string) {
	o.DataId = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *BTForeignDataResponse1070) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTForeignDataResponse1070) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *BTForeignDataResponse1070) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *BTForeignDataResponse1070) SetFormat(v string) {
	o.Format = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTForeignDataResponse1070) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTForeignDataResponse1070) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTForeignDataResponse1070) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTForeignDataResponse1070) SetName(v string) {
	o.Name = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *BTForeignDataResponse1070) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTForeignDataResponse1070) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *BTForeignDataResponse1070) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *BTForeignDataResponse1070) SetRegion(v string) {
	o.Region = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BTForeignDataResponse1070) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTForeignDataResponse1070) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BTForeignDataResponse1070) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *BTForeignDataResponse1070) SetSize(v int32) {
	o.Size = &v
}

// GetStorageType returns the StorageType field value if set, zero value otherwise.
func (o *BTForeignDataResponse1070) GetStorageType() string {
	if o == nil || o.StorageType == nil {
		var ret string
		return ret
	}
	return *o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTForeignDataResponse1070) GetStorageTypeOk() (*string, bool) {
	if o == nil || o.StorageType == nil {
		return nil, false
	}
	return o.StorageType, true
}

// HasStorageType returns a boolean if a field has been set.
func (o *BTForeignDataResponse1070) HasStorageType() bool {
	if o != nil && o.StorageType != nil {
		return true
	}

	return false
}

// SetStorageType gets a reference to the given string and assigns it to the StorageType field.
func (o *BTForeignDataResponse1070) SetStorageType(v string) {
	o.StorageType = &v
}

// GetUseLocalStorage returns the UseLocalStorage field value if set, zero value otherwise.
func (o *BTForeignDataResponse1070) GetUseLocalStorage() bool {
	if o == nil || o.UseLocalStorage == nil {
		var ret bool
		return ret
	}
	return *o.UseLocalStorage
}

// GetUseLocalStorageOk returns a tuple with the UseLocalStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTForeignDataResponse1070) GetUseLocalStorageOk() (*bool, bool) {
	if o == nil || o.UseLocalStorage == nil {
		return nil, false
	}
	return o.UseLocalStorage, true
}

// HasUseLocalStorage returns a boolean if a field has been set.
func (o *BTForeignDataResponse1070) HasUseLocalStorage() bool {
	if o != nil && o.UseLocalStorage != nil {
		return true
	}

	return false
}

// SetUseLocalStorage gets a reference to the given bool and assigns it to the UseLocalStorage field.
func (o *BTForeignDataResponse1070) SetUseLocalStorage(v bool) {
	o.UseLocalStorage = &v
}

func (o BTForeignDataResponse1070) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.BucketName != nil {
		toSerialize["bucketName"] = o.BucketName
	}
	if o.BucketPath != nil {
		toSerialize["bucketPath"] = o.BucketPath
	}
	if o.CacheChunkList != nil {
		toSerialize["cacheChunkList"] = o.CacheChunkList
	}
	if o.DataId != nil {
		toSerialize["dataId"] = o.DataId
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.StorageType != nil {
		toSerialize["storageType"] = o.StorageType
	}
	if o.UseLocalStorage != nil {
		toSerialize["useLocalStorage"] = o.UseLocalStorage
	}
	return json.Marshal(toSerialize)
}

type NullableBTForeignDataResponse1070 struct {
	value *BTForeignDataResponse1070
	isSet bool
}

func (v NullableBTForeignDataResponse1070) Get() *BTForeignDataResponse1070 {
	return v.value
}

func (v *NullableBTForeignDataResponse1070) Set(val *BTForeignDataResponse1070) {
	v.value = val
	v.isSet = true
}

func (v NullableBTForeignDataResponse1070) IsSet() bool {
	return v.isSet
}

func (v *NullableBTForeignDataResponse1070) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTForeignDataResponse1070(val *BTForeignDataResponse1070) *NullableBTForeignDataResponse1070 {
	return &NullableBTForeignDataResponse1070{value: val, isSet: true}
}

func (v NullableBTForeignDataResponse1070) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTForeignDataResponse1070) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
