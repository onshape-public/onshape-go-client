/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25478-d4e5ab4765a4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// OAuthFlows struct for OAuthFlows
type OAuthFlows struct {
	AuthorizationCode *OAuthFlow                        `json:"authorizationCode,omitempty"`
	ClientCredentials *OAuthFlow                        `json:"clientCredentials,omitempty"`
	Extensions        map[string]map[string]interface{} `json:"extensions,omitempty"`
	Implicit          *OAuthFlow                        `json:"implicit,omitempty"`
	Password          *OAuthFlow                        `json:"password,omitempty"`
}

// NewOAuthFlows instantiates a new OAuthFlows object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthFlows() *OAuthFlows {
	this := OAuthFlows{}
	return &this
}

// NewOAuthFlowsWithDefaults instantiates a new OAuthFlows object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthFlowsWithDefaults() *OAuthFlows {
	this := OAuthFlows{}
	return &this
}

// GetAuthorizationCode returns the AuthorizationCode field value if set, zero value otherwise.
func (o *OAuthFlows) GetAuthorizationCode() OAuthFlow {
	if o == nil || o.AuthorizationCode == nil {
		var ret OAuthFlow
		return ret
	}
	return *o.AuthorizationCode
}

// GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlows) GetAuthorizationCodeOk() (*OAuthFlow, bool) {
	if o == nil || o.AuthorizationCode == nil {
		return nil, false
	}
	return o.AuthorizationCode, true
}

// HasAuthorizationCode returns a boolean if a field has been set.
func (o *OAuthFlows) HasAuthorizationCode() bool {
	if o != nil && o.AuthorizationCode != nil {
		return true
	}

	return false
}

// SetAuthorizationCode gets a reference to the given OAuthFlow and assigns it to the AuthorizationCode field.
func (o *OAuthFlows) SetAuthorizationCode(v OAuthFlow) {
	o.AuthorizationCode = &v
}

// GetClientCredentials returns the ClientCredentials field value if set, zero value otherwise.
func (o *OAuthFlows) GetClientCredentials() OAuthFlow {
	if o == nil || o.ClientCredentials == nil {
		var ret OAuthFlow
		return ret
	}
	return *o.ClientCredentials
}

// GetClientCredentialsOk returns a tuple with the ClientCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlows) GetClientCredentialsOk() (*OAuthFlow, bool) {
	if o == nil || o.ClientCredentials == nil {
		return nil, false
	}
	return o.ClientCredentials, true
}

// HasClientCredentials returns a boolean if a field has been set.
func (o *OAuthFlows) HasClientCredentials() bool {
	if o != nil && o.ClientCredentials != nil {
		return true
	}

	return false
}

// SetClientCredentials gets a reference to the given OAuthFlow and assigns it to the ClientCredentials field.
func (o *OAuthFlows) SetClientCredentials(v OAuthFlow) {
	o.ClientCredentials = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *OAuthFlows) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlows) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *OAuthFlows) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *OAuthFlows) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetImplicit returns the Implicit field value if set, zero value otherwise.
func (o *OAuthFlows) GetImplicit() OAuthFlow {
	if o == nil || o.Implicit == nil {
		var ret OAuthFlow
		return ret
	}
	return *o.Implicit
}

// GetImplicitOk returns a tuple with the Implicit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlows) GetImplicitOk() (*OAuthFlow, bool) {
	if o == nil || o.Implicit == nil {
		return nil, false
	}
	return o.Implicit, true
}

// HasImplicit returns a boolean if a field has been set.
func (o *OAuthFlows) HasImplicit() bool {
	if o != nil && o.Implicit != nil {
		return true
	}

	return false
}

// SetImplicit gets a reference to the given OAuthFlow and assigns it to the Implicit field.
func (o *OAuthFlows) SetImplicit(v OAuthFlow) {
	o.Implicit = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *OAuthFlows) GetPassword() OAuthFlow {
	if o == nil || o.Password == nil {
		var ret OAuthFlow
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlows) GetPasswordOk() (*OAuthFlow, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *OAuthFlows) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given OAuthFlow and assigns it to the Password field.
func (o *OAuthFlows) SetPassword(v OAuthFlow) {
	o.Password = &v
}

func (o OAuthFlows) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationCode != nil {
		toSerialize["authorizationCode"] = o.AuthorizationCode
	}
	if o.ClientCredentials != nil {
		toSerialize["clientCredentials"] = o.ClientCredentials
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Implicit != nil {
		toSerialize["implicit"] = o.Implicit
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableOAuthFlows struct {
	value *OAuthFlows
	isSet bool
}

func (v NullableOAuthFlows) Get() *OAuthFlows {
	return v.value
}

func (v *NullableOAuthFlows) Set(val *OAuthFlows) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthFlows) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthFlows) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthFlows(val *OAuthFlows) *NullableOAuthFlows {
	return &NullableOAuthFlows{value: val, isSet: true}
}

func (v NullableOAuthFlows) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthFlows) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
