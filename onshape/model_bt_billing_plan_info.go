/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5768-0013f50d23d2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBillingPlanInfo struct for BTBillingPlanInfo
type BTBillingPlanInfo struct {
	AmountCents        *int32          `json:"amountCents,omitempty"`
	ApplicationId      *string         `json:"applicationId,omitempty"`
	ClientId           *string         `json:"clientId,omitempty"`
	CompanyPlan        *bool           `json:"companyPlan,omitempty"`
	ConsumableQuantity *int32          `json:"consumableQuantity,omitempty"`
	Deprecated         *bool           `json:"deprecated,omitempty"`
	Description        *string         `json:"description,omitempty"`
	DiscountInfo       *BTDiscountInfo `json:"discountInfo,omitempty"`
	Group              *string         `json:"group,omitempty"`
	Hidden             *bool           `json:"hidden,omitempty"`
	Href               *string         `json:"href,omitempty"`
	Id                 *string         `json:"id,omitempty"`
	Interval           *string         `json:"interval,omitempty"`
	Name               *string         `json:"name,omitempty"`
	OnshapePlan        *bool           `json:"onshapePlan,omitempty"`
	PlanType           *int32          `json:"planType,omitempty"`
	TrialPeriodDays    *int32          `json:"trialPeriodDays,omitempty"`
	ViewRef            *string         `json:"viewRef,omitempty"`
}

// NewBTBillingPlanInfo instantiates a new BTBillingPlanInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillingPlanInfo() *BTBillingPlanInfo {
	this := BTBillingPlanInfo{}
	return &this
}

// NewBTBillingPlanInfoWithDefaults instantiates a new BTBillingPlanInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillingPlanInfoWithDefaults() *BTBillingPlanInfo {
	this := BTBillingPlanInfo{}
	return &this
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetAmountCents() int32 {
	if o == nil || o.AmountCents == nil {
		var ret int32
		return ret
	}
	return *o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetAmountCentsOk() (*int32, bool) {
	if o == nil || o.AmountCents == nil {
		return nil, false
	}
	return o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasAmountCents() bool {
	if o != nil && o.AmountCents != nil {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given int32 and assigns it to the AmountCents field.
func (o *BTBillingPlanInfo) SetAmountCents(v int32) {
	o.AmountCents = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *BTBillingPlanInfo) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BTBillingPlanInfo) SetClientId(v string) {
	o.ClientId = &v
}

// GetCompanyPlan returns the CompanyPlan field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetCompanyPlan() bool {
	if o == nil || o.CompanyPlan == nil {
		var ret bool
		return ret
	}
	return *o.CompanyPlan
}

// GetCompanyPlanOk returns a tuple with the CompanyPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetCompanyPlanOk() (*bool, bool) {
	if o == nil || o.CompanyPlan == nil {
		return nil, false
	}
	return o.CompanyPlan, true
}

// HasCompanyPlan returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasCompanyPlan() bool {
	if o != nil && o.CompanyPlan != nil {
		return true
	}

	return false
}

// SetCompanyPlan gets a reference to the given bool and assigns it to the CompanyPlan field.
func (o *BTBillingPlanInfo) SetCompanyPlan(v bool) {
	o.CompanyPlan = &v
}

// GetConsumableQuantity returns the ConsumableQuantity field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetConsumableQuantity() int32 {
	if o == nil || o.ConsumableQuantity == nil {
		var ret int32
		return ret
	}
	return *o.ConsumableQuantity
}

// GetConsumableQuantityOk returns a tuple with the ConsumableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetConsumableQuantityOk() (*int32, bool) {
	if o == nil || o.ConsumableQuantity == nil {
		return nil, false
	}
	return o.ConsumableQuantity, true
}

// HasConsumableQuantity returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasConsumableQuantity() bool {
	if o != nil && o.ConsumableQuantity != nil {
		return true
	}

	return false
}

// SetConsumableQuantity gets a reference to the given int32 and assigns it to the ConsumableQuantity field.
func (o *BTBillingPlanInfo) SetConsumableQuantity(v int32) {
	o.ConsumableQuantity = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetDeprecated() bool {
	if o == nil || o.Deprecated == nil {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetDeprecatedOk() (*bool, bool) {
	if o == nil || o.Deprecated == nil {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasDeprecated() bool {
	if o != nil && o.Deprecated != nil {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *BTBillingPlanInfo) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTBillingPlanInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDiscountInfo returns the DiscountInfo field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetDiscountInfo() BTDiscountInfo {
	if o == nil || o.DiscountInfo == nil {
		var ret BTDiscountInfo
		return ret
	}
	return *o.DiscountInfo
}

// GetDiscountInfoOk returns a tuple with the DiscountInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetDiscountInfoOk() (*BTDiscountInfo, bool) {
	if o == nil || o.DiscountInfo == nil {
		return nil, false
	}
	return o.DiscountInfo, true
}

// HasDiscountInfo returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasDiscountInfo() bool {
	if o != nil && o.DiscountInfo != nil {
		return true
	}

	return false
}

// SetDiscountInfo gets a reference to the given BTDiscountInfo and assigns it to the DiscountInfo field.
func (o *BTBillingPlanInfo) SetDiscountInfo(v BTDiscountInfo) {
	o.DiscountInfo = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *BTBillingPlanInfo) SetGroup(v string) {
	o.Group = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *BTBillingPlanInfo) SetHidden(v bool) {
	o.Hidden = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTBillingPlanInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTBillingPlanInfo) SetId(v string) {
	o.Id = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetInterval() string {
	if o == nil || o.Interval == nil {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetIntervalOk() (*string, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *BTBillingPlanInfo) SetInterval(v string) {
	o.Interval = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTBillingPlanInfo) SetName(v string) {
	o.Name = &v
}

// GetOnshapePlan returns the OnshapePlan field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetOnshapePlan() bool {
	if o == nil || o.OnshapePlan == nil {
		var ret bool
		return ret
	}
	return *o.OnshapePlan
}

// GetOnshapePlanOk returns a tuple with the OnshapePlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetOnshapePlanOk() (*bool, bool) {
	if o == nil || o.OnshapePlan == nil {
		return nil, false
	}
	return o.OnshapePlan, true
}

// HasOnshapePlan returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasOnshapePlan() bool {
	if o != nil && o.OnshapePlan != nil {
		return true
	}

	return false
}

// SetOnshapePlan gets a reference to the given bool and assigns it to the OnshapePlan field.
func (o *BTBillingPlanInfo) SetOnshapePlan(v bool) {
	o.OnshapePlan = &v
}

// GetPlanType returns the PlanType field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetPlanType() int32 {
	if o == nil || o.PlanType == nil {
		var ret int32
		return ret
	}
	return *o.PlanType
}

// GetPlanTypeOk returns a tuple with the PlanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetPlanTypeOk() (*int32, bool) {
	if o == nil || o.PlanType == nil {
		return nil, false
	}
	return o.PlanType, true
}

// HasPlanType returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasPlanType() bool {
	if o != nil && o.PlanType != nil {
		return true
	}

	return false
}

// SetPlanType gets a reference to the given int32 and assigns it to the PlanType field.
func (o *BTBillingPlanInfo) SetPlanType(v int32) {
	o.PlanType = &v
}

// GetTrialPeriodDays returns the TrialPeriodDays field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetTrialPeriodDays() int32 {
	if o == nil || o.TrialPeriodDays == nil {
		var ret int32
		return ret
	}
	return *o.TrialPeriodDays
}

// GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetTrialPeriodDaysOk() (*int32, bool) {
	if o == nil || o.TrialPeriodDays == nil {
		return nil, false
	}
	return o.TrialPeriodDays, true
}

// HasTrialPeriodDays returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasTrialPeriodDays() bool {
	if o != nil && o.TrialPeriodDays != nil {
		return true
	}

	return false
}

// SetTrialPeriodDays gets a reference to the given int32 and assigns it to the TrialPeriodDays field.
func (o *BTBillingPlanInfo) SetTrialPeriodDays(v int32) {
	o.TrialPeriodDays = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTBillingPlanInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillingPlanInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTBillingPlanInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTBillingPlanInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTBillingPlanInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AmountCents != nil {
		toSerialize["amountCents"] = o.AmountCents
	}
	if o.ApplicationId != nil {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.CompanyPlan != nil {
		toSerialize["companyPlan"] = o.CompanyPlan
	}
	if o.ConsumableQuantity != nil {
		toSerialize["consumableQuantity"] = o.ConsumableQuantity
	}
	if o.Deprecated != nil {
		toSerialize["deprecated"] = o.Deprecated
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DiscountInfo != nil {
		toSerialize["discountInfo"] = o.DiscountInfo
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OnshapePlan != nil {
		toSerialize["onshapePlan"] = o.OnshapePlan
	}
	if o.PlanType != nil {
		toSerialize["planType"] = o.PlanType
	}
	if o.TrialPeriodDays != nil {
		toSerialize["trialPeriodDays"] = o.TrialPeriodDays
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTBillingPlanInfo struct {
	value *BTBillingPlanInfo
	isSet bool
}

func (v NullableBTBillingPlanInfo) Get() *BTBillingPlanInfo {
	return v.value
}

func (v *NullableBTBillingPlanInfo) Set(val *BTBillingPlanInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillingPlanInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillingPlanInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillingPlanInfo(val *BTBillingPlanInfo) *NullableBTBillingPlanInfo {
	return &NullableBTBillingPlanInfo{value: val, isSet: true}
}

func (v NullableBTBillingPlanInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillingPlanInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
