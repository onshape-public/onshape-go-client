/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6471-f221dba4a96a
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTTableCrossHighlightData1753 struct for BTTableCrossHighlightData1753
type BTTableCrossHighlightData1753 struct {
	BtType              *string  `json:"btType,omitempty"`
	DeterministicIdList []string `json:"deterministicIdList,omitempty"`
	FeatureIdList       []string `json:"featureIdList,omitempty"`
}

// NewBTTableCrossHighlightData1753 instantiates a new BTTableCrossHighlightData1753 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableCrossHighlightData1753() *BTTableCrossHighlightData1753 {
	this := BTTableCrossHighlightData1753{}
	return &this
}

// NewBTTableCrossHighlightData1753WithDefaults instantiates a new BTTableCrossHighlightData1753 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableCrossHighlightData1753WithDefaults() *BTTableCrossHighlightData1753 {
	this := BTTableCrossHighlightData1753{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableCrossHighlightData1753) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCrossHighlightData1753) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTableCrossHighlightData1753) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTableCrossHighlightData1753) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTTableCrossHighlightData1753) GetDeterministicIdList() []string {
	if o == nil || o.DeterministicIdList == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCrossHighlightData1753) GetDeterministicIdListOk() ([]string, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTTableCrossHighlightData1753) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given []string and assigns it to the DeterministicIdList field.
func (o *BTTableCrossHighlightData1753) SetDeterministicIdList(v []string) {
	o.DeterministicIdList = v
}

// GetFeatureIdList returns the FeatureIdList field value if set, zero value otherwise.
func (o *BTTableCrossHighlightData1753) GetFeatureIdList() []string {
	if o == nil || o.FeatureIdList == nil {
		var ret []string
		return ret
	}
	return o.FeatureIdList
}

// GetFeatureIdListOk returns a tuple with the FeatureIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCrossHighlightData1753) GetFeatureIdListOk() ([]string, bool) {
	if o == nil || o.FeatureIdList == nil {
		return nil, false
	}
	return o.FeatureIdList, true
}

// HasFeatureIdList returns a boolean if a field has been set.
func (o *BTTableCrossHighlightData1753) HasFeatureIdList() bool {
	if o != nil && o.FeatureIdList != nil {
		return true
	}

	return false
}

// SetFeatureIdList gets a reference to the given []string and assigns it to the FeatureIdList field.
func (o *BTTableCrossHighlightData1753) SetFeatureIdList(v []string) {
	o.FeatureIdList = v
}

func (o BTTableCrossHighlightData1753) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DeterministicIdList != nil {
		toSerialize["deterministicIdList"] = o.DeterministicIdList
	}
	if o.FeatureIdList != nil {
		toSerialize["featureIdList"] = o.FeatureIdList
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableCrossHighlightData1753 struct {
	value *BTTableCrossHighlightData1753
	isSet bool
}

func (v NullableBTTableCrossHighlightData1753) Get() *BTTableCrossHighlightData1753 {
	return v.value
}

func (v *NullableBTTableCrossHighlightData1753) Set(val *BTTableCrossHighlightData1753) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableCrossHighlightData1753) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableCrossHighlightData1753) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableCrossHighlightData1753(val *BTTableCrossHighlightData1753) *NullableBTTableCrossHighlightData1753 {
	return &NullableBTTableCrossHighlightData1753{value: val, isSet: true}
}

func (v NullableBTTableCrossHighlightData1753) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableCrossHighlightData1753) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
