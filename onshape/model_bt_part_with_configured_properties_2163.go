/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6017-d06351faf6f2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPartWithConfiguredProperties2163 struct for BTPartWithConfiguredProperties2163
type BTPartWithConfiguredProperties2163 struct {
	BtType                  *string                                `json:"btType,omitempty"`
	ConfigurationProperties []BTOneConfigurationPartProperties1661 `json:"configurationProperties,omitempty"`
	ForSubPartProperties    *bool                                  `json:"forSubPartProperties,omitempty"`
	NodeId                  *string                                `json:"nodeId,omitempty"`
	ParsedQuery             *BTPFunctionDeclaration246             `json:"parsedQuery,omitempty"`
	PropertyNodeId          *string                                `json:"propertyNodeId,omitempty"`
	Query                   *string                                `json:"query,omitempty"`
}

// NewBTPartWithConfiguredProperties2163 instantiates a new BTPartWithConfiguredProperties2163 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPartWithConfiguredProperties2163() *BTPartWithConfiguredProperties2163 {
	this := BTPartWithConfiguredProperties2163{}
	return &this
}

// NewBTPartWithConfiguredProperties2163WithDefaults instantiates a new BTPartWithConfiguredProperties2163 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPartWithConfiguredProperties2163WithDefaults() *BTPartWithConfiguredProperties2163 {
	this := BTPartWithConfiguredProperties2163{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPartWithConfiguredProperties2163) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartWithConfiguredProperties2163) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPartWithConfiguredProperties2163) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPartWithConfiguredProperties2163) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigurationProperties returns the ConfigurationProperties field value if set, zero value otherwise.
func (o *BTPartWithConfiguredProperties2163) GetConfigurationProperties() []BTOneConfigurationPartProperties1661 {
	if o == nil || o.ConfigurationProperties == nil {
		var ret []BTOneConfigurationPartProperties1661
		return ret
	}
	return o.ConfigurationProperties
}

// GetConfigurationPropertiesOk returns a tuple with the ConfigurationProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartWithConfiguredProperties2163) GetConfigurationPropertiesOk() ([]BTOneConfigurationPartProperties1661, bool) {
	if o == nil || o.ConfigurationProperties == nil {
		return nil, false
	}
	return o.ConfigurationProperties, true
}

// HasConfigurationProperties returns a boolean if a field has been set.
func (o *BTPartWithConfiguredProperties2163) HasConfigurationProperties() bool {
	if o != nil && o.ConfigurationProperties != nil {
		return true
	}

	return false
}

// SetConfigurationProperties gets a reference to the given []BTOneConfigurationPartProperties1661 and assigns it to the ConfigurationProperties field.
func (o *BTPartWithConfiguredProperties2163) SetConfigurationProperties(v []BTOneConfigurationPartProperties1661) {
	o.ConfigurationProperties = v
}

// GetForSubPartProperties returns the ForSubPartProperties field value if set, zero value otherwise.
func (o *BTPartWithConfiguredProperties2163) GetForSubPartProperties() bool {
	if o == nil || o.ForSubPartProperties == nil {
		var ret bool
		return ret
	}
	return *o.ForSubPartProperties
}

// GetForSubPartPropertiesOk returns a tuple with the ForSubPartProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartWithConfiguredProperties2163) GetForSubPartPropertiesOk() (*bool, bool) {
	if o == nil || o.ForSubPartProperties == nil {
		return nil, false
	}
	return o.ForSubPartProperties, true
}

// HasForSubPartProperties returns a boolean if a field has been set.
func (o *BTPartWithConfiguredProperties2163) HasForSubPartProperties() bool {
	if o != nil && o.ForSubPartProperties != nil {
		return true
	}

	return false
}

// SetForSubPartProperties gets a reference to the given bool and assigns it to the ForSubPartProperties field.
func (o *BTPartWithConfiguredProperties2163) SetForSubPartProperties(v bool) {
	o.ForSubPartProperties = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPartWithConfiguredProperties2163) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartWithConfiguredProperties2163) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPartWithConfiguredProperties2163) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPartWithConfiguredProperties2163) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParsedQuery returns the ParsedQuery field value if set, zero value otherwise.
func (o *BTPartWithConfiguredProperties2163) GetParsedQuery() BTPFunctionDeclaration246 {
	if o == nil || o.ParsedQuery == nil {
		var ret BTPFunctionDeclaration246
		return ret
	}
	return *o.ParsedQuery
}

// GetParsedQueryOk returns a tuple with the ParsedQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartWithConfiguredProperties2163) GetParsedQueryOk() (*BTPFunctionDeclaration246, bool) {
	if o == nil || o.ParsedQuery == nil {
		return nil, false
	}
	return o.ParsedQuery, true
}

// HasParsedQuery returns a boolean if a field has been set.
func (o *BTPartWithConfiguredProperties2163) HasParsedQuery() bool {
	if o != nil && o.ParsedQuery != nil {
		return true
	}

	return false
}

// SetParsedQuery gets a reference to the given BTPFunctionDeclaration246 and assigns it to the ParsedQuery field.
func (o *BTPartWithConfiguredProperties2163) SetParsedQuery(v BTPFunctionDeclaration246) {
	o.ParsedQuery = &v
}

// GetPropertyNodeId returns the PropertyNodeId field value if set, zero value otherwise.
func (o *BTPartWithConfiguredProperties2163) GetPropertyNodeId() string {
	if o == nil || o.PropertyNodeId == nil {
		var ret string
		return ret
	}
	return *o.PropertyNodeId
}

// GetPropertyNodeIdOk returns a tuple with the PropertyNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartWithConfiguredProperties2163) GetPropertyNodeIdOk() (*string, bool) {
	if o == nil || o.PropertyNodeId == nil {
		return nil, false
	}
	return o.PropertyNodeId, true
}

// HasPropertyNodeId returns a boolean if a field has been set.
func (o *BTPartWithConfiguredProperties2163) HasPropertyNodeId() bool {
	if o != nil && o.PropertyNodeId != nil {
		return true
	}

	return false
}

// SetPropertyNodeId gets a reference to the given string and assigns it to the PropertyNodeId field.
func (o *BTPartWithConfiguredProperties2163) SetPropertyNodeId(v string) {
	o.PropertyNodeId = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTPartWithConfiguredProperties2163) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartWithConfiguredProperties2163) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *BTPartWithConfiguredProperties2163) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *BTPartWithConfiguredProperties2163) SetQuery(v string) {
	o.Query = &v
}

func (o BTPartWithConfiguredProperties2163) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ConfigurationProperties != nil {
		toSerialize["configurationProperties"] = o.ConfigurationProperties
	}
	if o.ForSubPartProperties != nil {
		toSerialize["forSubPartProperties"] = o.ForSubPartProperties
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParsedQuery != nil {
		toSerialize["parsedQuery"] = o.ParsedQuery
	}
	if o.PropertyNodeId != nil {
		toSerialize["propertyNodeId"] = o.PropertyNodeId
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableBTPartWithConfiguredProperties2163 struct {
	value *BTPartWithConfiguredProperties2163
	isSet bool
}

func (v NullableBTPartWithConfiguredProperties2163) Get() *BTPartWithConfiguredProperties2163 {
	return v.value
}

func (v *NullableBTPartWithConfiguredProperties2163) Set(val *BTPartWithConfiguredProperties2163) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPartWithConfiguredProperties2163) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPartWithConfiguredProperties2163) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPartWithConfiguredProperties2163(val *BTPartWithConfiguredProperties2163) *NullableBTPartWithConfiguredProperties2163 {
	return &NullableBTPartWithConfiguredProperties2163{value: val, isSet: true}
}

func (v NullableBTPartWithConfiguredProperties2163) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPartWithConfiguredProperties2163) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
