/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6438-5fd2d9755d52
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Card struct for Card
type Card struct {
	Account                *string            `json:"account,omitempty"`
	AddressCity            *string            `json:"addressCity,omitempty"`
	AddressCountry         *string            `json:"addressCountry,omitempty"`
	AddressLine1           *string            `json:"addressLine1,omitempty"`
	AddressLine1Check      *string            `json:"addressLine1Check,omitempty"`
	AddressLine2           *string            `json:"addressLine2,omitempty"`
	AddressState           *string            `json:"addressState,omitempty"`
	AddressZip             *string            `json:"addressZip,omitempty"`
	AddressZipCheck        *string            `json:"addressZipCheck,omitempty"`
	AvailablePayoutMethods []string           `json:"availablePayoutMethods,omitempty"`
	Brand                  *string            `json:"brand,omitempty"`
	Country                *string            `json:"country,omitempty"`
	Currency               *string            `json:"currency,omitempty"`
	Customer               *string            `json:"customer,omitempty"`
	CvcCheck               *string            `json:"cvcCheck,omitempty"`
	DefaultForCurrency     *bool              `json:"defaultForCurrency,omitempty"`
	Description            *string            `json:"description,omitempty"`
	DynamicLast4           *string            `json:"dynamicLast4,omitempty"`
	ExpMonth               *int32             `json:"expMonth,omitempty"`
	ExpYear                *int32             `json:"expYear,omitempty"`
	Fingerprint            *string            `json:"fingerprint,omitempty"`
	Funding                *string            `json:"funding,omitempty"`
	Id                     *string            `json:"id,omitempty"`
	Iin                    *string            `json:"iin,omitempty"`
	InstanceURL            *string            `json:"instanceURL,omitempty"`
	Issuer                 *string            `json:"issuer,omitempty"`
	Last4                  *string            `json:"last4,omitempty"`
	Metadata               *map[string]string `json:"metadata,omitempty"`
	Name                   *string            `json:"name,omitempty"`
	Object                 *string            `json:"object,omitempty"`
	Recipient              *string            `json:"recipient,omitempty"`
	Status                 *string            `json:"status,omitempty"`
	ThreeDSecure           *ThreeDSecure      `json:"threeDSecure,omitempty"`
	TokenizationMethod     *string            `json:"tokenizationMethod,omitempty"`
	Type                   *string            `json:"type,omitempty"`
}

// NewCard instantiates a new Card object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCard() *Card {
	this := Card{}
	return &this
}

// NewCardWithDefaults instantiates a new Card object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardWithDefaults() *Card {
	this := Card{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Card) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Card) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *Card) SetAccount(v string) {
	o.Account = &v
}

// GetAddressCity returns the AddressCity field value if set, zero value otherwise.
func (o *Card) GetAddressCity() string {
	if o == nil || o.AddressCity == nil {
		var ret string
		return ret
	}
	return *o.AddressCity
}

// GetAddressCityOk returns a tuple with the AddressCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetAddressCityOk() (*string, bool) {
	if o == nil || o.AddressCity == nil {
		return nil, false
	}
	return o.AddressCity, true
}

// HasAddressCity returns a boolean if a field has been set.
func (o *Card) HasAddressCity() bool {
	if o != nil && o.AddressCity != nil {
		return true
	}

	return false
}

// SetAddressCity gets a reference to the given string and assigns it to the AddressCity field.
func (o *Card) SetAddressCity(v string) {
	o.AddressCity = &v
}

// GetAddressCountry returns the AddressCountry field value if set, zero value otherwise.
func (o *Card) GetAddressCountry() string {
	if o == nil || o.AddressCountry == nil {
		var ret string
		return ret
	}
	return *o.AddressCountry
}

// GetAddressCountryOk returns a tuple with the AddressCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetAddressCountryOk() (*string, bool) {
	if o == nil || o.AddressCountry == nil {
		return nil, false
	}
	return o.AddressCountry, true
}

// HasAddressCountry returns a boolean if a field has been set.
func (o *Card) HasAddressCountry() bool {
	if o != nil && o.AddressCountry != nil {
		return true
	}

	return false
}

// SetAddressCountry gets a reference to the given string and assigns it to the AddressCountry field.
func (o *Card) SetAddressCountry(v string) {
	o.AddressCountry = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *Card) GetAddressLine1() string {
	if o == nil || o.AddressLine1 == nil {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetAddressLine1Ok() (*string, bool) {
	if o == nil || o.AddressLine1 == nil {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *Card) HasAddressLine1() bool {
	if o != nil && o.AddressLine1 != nil {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *Card) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine1Check returns the AddressLine1Check field value if set, zero value otherwise.
func (o *Card) GetAddressLine1Check() string {
	if o == nil || o.AddressLine1Check == nil {
		var ret string
		return ret
	}
	return *o.AddressLine1Check
}

// GetAddressLine1CheckOk returns a tuple with the AddressLine1Check field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetAddressLine1CheckOk() (*string, bool) {
	if o == nil || o.AddressLine1Check == nil {
		return nil, false
	}
	return o.AddressLine1Check, true
}

// HasAddressLine1Check returns a boolean if a field has been set.
func (o *Card) HasAddressLine1Check() bool {
	if o != nil && o.AddressLine1Check != nil {
		return true
	}

	return false
}

// SetAddressLine1Check gets a reference to the given string and assigns it to the AddressLine1Check field.
func (o *Card) SetAddressLine1Check(v string) {
	o.AddressLine1Check = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *Card) GetAddressLine2() string {
	if o == nil || o.AddressLine2 == nil {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetAddressLine2Ok() (*string, bool) {
	if o == nil || o.AddressLine2 == nil {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *Card) HasAddressLine2() bool {
	if o != nil && o.AddressLine2 != nil {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *Card) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressState returns the AddressState field value if set, zero value otherwise.
func (o *Card) GetAddressState() string {
	if o == nil || o.AddressState == nil {
		var ret string
		return ret
	}
	return *o.AddressState
}

// GetAddressStateOk returns a tuple with the AddressState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetAddressStateOk() (*string, bool) {
	if o == nil || o.AddressState == nil {
		return nil, false
	}
	return o.AddressState, true
}

// HasAddressState returns a boolean if a field has been set.
func (o *Card) HasAddressState() bool {
	if o != nil && o.AddressState != nil {
		return true
	}

	return false
}

// SetAddressState gets a reference to the given string and assigns it to the AddressState field.
func (o *Card) SetAddressState(v string) {
	o.AddressState = &v
}

// GetAddressZip returns the AddressZip field value if set, zero value otherwise.
func (o *Card) GetAddressZip() string {
	if o == nil || o.AddressZip == nil {
		var ret string
		return ret
	}
	return *o.AddressZip
}

// GetAddressZipOk returns a tuple with the AddressZip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetAddressZipOk() (*string, bool) {
	if o == nil || o.AddressZip == nil {
		return nil, false
	}
	return o.AddressZip, true
}

// HasAddressZip returns a boolean if a field has been set.
func (o *Card) HasAddressZip() bool {
	if o != nil && o.AddressZip != nil {
		return true
	}

	return false
}

// SetAddressZip gets a reference to the given string and assigns it to the AddressZip field.
func (o *Card) SetAddressZip(v string) {
	o.AddressZip = &v
}

// GetAddressZipCheck returns the AddressZipCheck field value if set, zero value otherwise.
func (o *Card) GetAddressZipCheck() string {
	if o == nil || o.AddressZipCheck == nil {
		var ret string
		return ret
	}
	return *o.AddressZipCheck
}

// GetAddressZipCheckOk returns a tuple with the AddressZipCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetAddressZipCheckOk() (*string, bool) {
	if o == nil || o.AddressZipCheck == nil {
		return nil, false
	}
	return o.AddressZipCheck, true
}

// HasAddressZipCheck returns a boolean if a field has been set.
func (o *Card) HasAddressZipCheck() bool {
	if o != nil && o.AddressZipCheck != nil {
		return true
	}

	return false
}

// SetAddressZipCheck gets a reference to the given string and assigns it to the AddressZipCheck field.
func (o *Card) SetAddressZipCheck(v string) {
	o.AddressZipCheck = &v
}

// GetAvailablePayoutMethods returns the AvailablePayoutMethods field value if set, zero value otherwise.
func (o *Card) GetAvailablePayoutMethods() []string {
	if o == nil || o.AvailablePayoutMethods == nil {
		var ret []string
		return ret
	}
	return o.AvailablePayoutMethods
}

// GetAvailablePayoutMethodsOk returns a tuple with the AvailablePayoutMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetAvailablePayoutMethodsOk() ([]string, bool) {
	if o == nil || o.AvailablePayoutMethods == nil {
		return nil, false
	}
	return o.AvailablePayoutMethods, true
}

// HasAvailablePayoutMethods returns a boolean if a field has been set.
func (o *Card) HasAvailablePayoutMethods() bool {
	if o != nil && o.AvailablePayoutMethods != nil {
		return true
	}

	return false
}

// SetAvailablePayoutMethods gets a reference to the given []string and assigns it to the AvailablePayoutMethods field.
func (o *Card) SetAvailablePayoutMethods(v []string) {
	o.AvailablePayoutMethods = v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *Card) GetBrand() string {
	if o == nil || o.Brand == nil {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetBrandOk() (*string, bool) {
	if o == nil || o.Brand == nil {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *Card) HasBrand() bool {
	if o != nil && o.Brand != nil {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *Card) SetBrand(v string) {
	o.Brand = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Card) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Card) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Card) SetCountry(v string) {
	o.Country = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Card) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Card) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Card) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *Card) GetCustomer() string {
	if o == nil || o.Customer == nil {
		var ret string
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetCustomerOk() (*string, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *Card) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given string and assigns it to the Customer field.
func (o *Card) SetCustomer(v string) {
	o.Customer = &v
}

// GetCvcCheck returns the CvcCheck field value if set, zero value otherwise.
func (o *Card) GetCvcCheck() string {
	if o == nil || o.CvcCheck == nil {
		var ret string
		return ret
	}
	return *o.CvcCheck
}

// GetCvcCheckOk returns a tuple with the CvcCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetCvcCheckOk() (*string, bool) {
	if o == nil || o.CvcCheck == nil {
		return nil, false
	}
	return o.CvcCheck, true
}

// HasCvcCheck returns a boolean if a field has been set.
func (o *Card) HasCvcCheck() bool {
	if o != nil && o.CvcCheck != nil {
		return true
	}

	return false
}

// SetCvcCheck gets a reference to the given string and assigns it to the CvcCheck field.
func (o *Card) SetCvcCheck(v string) {
	o.CvcCheck = &v
}

// GetDefaultForCurrency returns the DefaultForCurrency field value if set, zero value otherwise.
func (o *Card) GetDefaultForCurrency() bool {
	if o == nil || o.DefaultForCurrency == nil {
		var ret bool
		return ret
	}
	return *o.DefaultForCurrency
}

// GetDefaultForCurrencyOk returns a tuple with the DefaultForCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetDefaultForCurrencyOk() (*bool, bool) {
	if o == nil || o.DefaultForCurrency == nil {
		return nil, false
	}
	return o.DefaultForCurrency, true
}

// HasDefaultForCurrency returns a boolean if a field has been set.
func (o *Card) HasDefaultForCurrency() bool {
	if o != nil && o.DefaultForCurrency != nil {
		return true
	}

	return false
}

// SetDefaultForCurrency gets a reference to the given bool and assigns it to the DefaultForCurrency field.
func (o *Card) SetDefaultForCurrency(v bool) {
	o.DefaultForCurrency = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Card) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Card) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Card) SetDescription(v string) {
	o.Description = &v
}

// GetDynamicLast4 returns the DynamicLast4 field value if set, zero value otherwise.
func (o *Card) GetDynamicLast4() string {
	if o == nil || o.DynamicLast4 == nil {
		var ret string
		return ret
	}
	return *o.DynamicLast4
}

// GetDynamicLast4Ok returns a tuple with the DynamicLast4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetDynamicLast4Ok() (*string, bool) {
	if o == nil || o.DynamicLast4 == nil {
		return nil, false
	}
	return o.DynamicLast4, true
}

// HasDynamicLast4 returns a boolean if a field has been set.
func (o *Card) HasDynamicLast4() bool {
	if o != nil && o.DynamicLast4 != nil {
		return true
	}

	return false
}

// SetDynamicLast4 gets a reference to the given string and assigns it to the DynamicLast4 field.
func (o *Card) SetDynamicLast4(v string) {
	o.DynamicLast4 = &v
}

// GetExpMonth returns the ExpMonth field value if set, zero value otherwise.
func (o *Card) GetExpMonth() int32 {
	if o == nil || o.ExpMonth == nil {
		var ret int32
		return ret
	}
	return *o.ExpMonth
}

// GetExpMonthOk returns a tuple with the ExpMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetExpMonthOk() (*int32, bool) {
	if o == nil || o.ExpMonth == nil {
		return nil, false
	}
	return o.ExpMonth, true
}

// HasExpMonth returns a boolean if a field has been set.
func (o *Card) HasExpMonth() bool {
	if o != nil && o.ExpMonth != nil {
		return true
	}

	return false
}

// SetExpMonth gets a reference to the given int32 and assigns it to the ExpMonth field.
func (o *Card) SetExpMonth(v int32) {
	o.ExpMonth = &v
}

// GetExpYear returns the ExpYear field value if set, zero value otherwise.
func (o *Card) GetExpYear() int32 {
	if o == nil || o.ExpYear == nil {
		var ret int32
		return ret
	}
	return *o.ExpYear
}

// GetExpYearOk returns a tuple with the ExpYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetExpYearOk() (*int32, bool) {
	if o == nil || o.ExpYear == nil {
		return nil, false
	}
	return o.ExpYear, true
}

// HasExpYear returns a boolean if a field has been set.
func (o *Card) HasExpYear() bool {
	if o != nil && o.ExpYear != nil {
		return true
	}

	return false
}

// SetExpYear gets a reference to the given int32 and assigns it to the ExpYear field.
func (o *Card) SetExpYear(v int32) {
	o.ExpYear = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *Card) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *Card) HasFingerprint() bool {
	if o != nil && o.Fingerprint != nil {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *Card) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetFunding returns the Funding field value if set, zero value otherwise.
func (o *Card) GetFunding() string {
	if o == nil || o.Funding == nil {
		var ret string
		return ret
	}
	return *o.Funding
}

// GetFundingOk returns a tuple with the Funding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetFundingOk() (*string, bool) {
	if o == nil || o.Funding == nil {
		return nil, false
	}
	return o.Funding, true
}

// HasFunding returns a boolean if a field has been set.
func (o *Card) HasFunding() bool {
	if o != nil && o.Funding != nil {
		return true
	}

	return false
}

// SetFunding gets a reference to the given string and assigns it to the Funding field.
func (o *Card) SetFunding(v string) {
	o.Funding = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Card) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Card) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Card) SetId(v string) {
	o.Id = &v
}

// GetIin returns the Iin field value if set, zero value otherwise.
func (o *Card) GetIin() string {
	if o == nil || o.Iin == nil {
		var ret string
		return ret
	}
	return *o.Iin
}

// GetIinOk returns a tuple with the Iin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetIinOk() (*string, bool) {
	if o == nil || o.Iin == nil {
		return nil, false
	}
	return o.Iin, true
}

// HasIin returns a boolean if a field has been set.
func (o *Card) HasIin() bool {
	if o != nil && o.Iin != nil {
		return true
	}

	return false
}

// SetIin gets a reference to the given string and assigns it to the Iin field.
func (o *Card) SetIin(v string) {
	o.Iin = &v
}

// GetInstanceURL returns the InstanceURL field value if set, zero value otherwise.
func (o *Card) GetInstanceURL() string {
	if o == nil || o.InstanceURL == nil {
		var ret string
		return ret
	}
	return *o.InstanceURL
}

// GetInstanceURLOk returns a tuple with the InstanceURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetInstanceURLOk() (*string, bool) {
	if o == nil || o.InstanceURL == nil {
		return nil, false
	}
	return o.InstanceURL, true
}

// HasInstanceURL returns a boolean if a field has been set.
func (o *Card) HasInstanceURL() bool {
	if o != nil && o.InstanceURL != nil {
		return true
	}

	return false
}

// SetInstanceURL gets a reference to the given string and assigns it to the InstanceURL field.
func (o *Card) SetInstanceURL(v string) {
	o.InstanceURL = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *Card) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *Card) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *Card) SetIssuer(v string) {
	o.Issuer = &v
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *Card) GetLast4() string {
	if o == nil || o.Last4 == nil {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetLast4Ok() (*string, bool) {
	if o == nil || o.Last4 == nil {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *Card) HasLast4() bool {
	if o != nil && o.Last4 != nil {
		return true
	}

	return false
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *Card) SetLast4(v string) {
	o.Last4 = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Card) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Card) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *Card) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Card) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Card) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Card) SetName(v string) {
	o.Name = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *Card) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *Card) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *Card) SetObject(v string) {
	o.Object = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *Card) GetRecipient() string {
	if o == nil || o.Recipient == nil {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetRecipientOk() (*string, bool) {
	if o == nil || o.Recipient == nil {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *Card) HasRecipient() bool {
	if o != nil && o.Recipient != nil {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *Card) SetRecipient(v string) {
	o.Recipient = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Card) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Card) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Card) SetStatus(v string) {
	o.Status = &v
}

// GetThreeDSecure returns the ThreeDSecure field value if set, zero value otherwise.
func (o *Card) GetThreeDSecure() ThreeDSecure {
	if o == nil || o.ThreeDSecure == nil {
		var ret ThreeDSecure
		return ret
	}
	return *o.ThreeDSecure
}

// GetThreeDSecureOk returns a tuple with the ThreeDSecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetThreeDSecureOk() (*ThreeDSecure, bool) {
	if o == nil || o.ThreeDSecure == nil {
		return nil, false
	}
	return o.ThreeDSecure, true
}

// HasThreeDSecure returns a boolean if a field has been set.
func (o *Card) HasThreeDSecure() bool {
	if o != nil && o.ThreeDSecure != nil {
		return true
	}

	return false
}

// SetThreeDSecure gets a reference to the given ThreeDSecure and assigns it to the ThreeDSecure field.
func (o *Card) SetThreeDSecure(v ThreeDSecure) {
	o.ThreeDSecure = &v
}

// GetTokenizationMethod returns the TokenizationMethod field value if set, zero value otherwise.
func (o *Card) GetTokenizationMethod() string {
	if o == nil || o.TokenizationMethod == nil {
		var ret string
		return ret
	}
	return *o.TokenizationMethod
}

// GetTokenizationMethodOk returns a tuple with the TokenizationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetTokenizationMethodOk() (*string, bool) {
	if o == nil || o.TokenizationMethod == nil {
		return nil, false
	}
	return o.TokenizationMethod, true
}

// HasTokenizationMethod returns a boolean if a field has been set.
func (o *Card) HasTokenizationMethod() bool {
	if o != nil && o.TokenizationMethod != nil {
		return true
	}

	return false
}

// SetTokenizationMethod gets a reference to the given string and assigns it to the TokenizationMethod field.
func (o *Card) SetTokenizationMethod(v string) {
	o.TokenizationMethod = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Card) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Card) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Card) SetType(v string) {
	o.Type = &v
}

func (o Card) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.AddressCity != nil {
		toSerialize["addressCity"] = o.AddressCity
	}
	if o.AddressCountry != nil {
		toSerialize["addressCountry"] = o.AddressCountry
	}
	if o.AddressLine1 != nil {
		toSerialize["addressLine1"] = o.AddressLine1
	}
	if o.AddressLine1Check != nil {
		toSerialize["addressLine1Check"] = o.AddressLine1Check
	}
	if o.AddressLine2 != nil {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if o.AddressState != nil {
		toSerialize["addressState"] = o.AddressState
	}
	if o.AddressZip != nil {
		toSerialize["addressZip"] = o.AddressZip
	}
	if o.AddressZipCheck != nil {
		toSerialize["addressZipCheck"] = o.AddressZipCheck
	}
	if o.AvailablePayoutMethods != nil {
		toSerialize["availablePayoutMethods"] = o.AvailablePayoutMethods
	}
	if o.Brand != nil {
		toSerialize["brand"] = o.Brand
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.CvcCheck != nil {
		toSerialize["cvcCheck"] = o.CvcCheck
	}
	if o.DefaultForCurrency != nil {
		toSerialize["defaultForCurrency"] = o.DefaultForCurrency
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DynamicLast4 != nil {
		toSerialize["dynamicLast4"] = o.DynamicLast4
	}
	if o.ExpMonth != nil {
		toSerialize["expMonth"] = o.ExpMonth
	}
	if o.ExpYear != nil {
		toSerialize["expYear"] = o.ExpYear
	}
	if o.Fingerprint != nil {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if o.Funding != nil {
		toSerialize["funding"] = o.Funding
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Iin != nil {
		toSerialize["iin"] = o.Iin
	}
	if o.InstanceURL != nil {
		toSerialize["instanceURL"] = o.InstanceURL
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.Last4 != nil {
		toSerialize["last4"] = o.Last4
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Recipient != nil {
		toSerialize["recipient"] = o.Recipient
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ThreeDSecure != nil {
		toSerialize["threeDSecure"] = o.ThreeDSecure
	}
	if o.TokenizationMethod != nil {
		toSerialize["tokenizationMethod"] = o.TokenizationMethod
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCard struct {
	value *Card
	isSet bool
}

func (v NullableCard) Get() *Card {
	return v.value
}

func (v *NullableCard) Set(val *Card) {
	v.value = val
	v.isSet = true
}

func (v NullableCard) IsSet() bool {
	return v.isSet
}

func (v *NullableCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCard(val *Card) *NullableCard {
	return &NullableCard{value: val, isSet: true}
}

func (v NullableCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
