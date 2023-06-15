/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.17497-411bb6b98e6b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDiscountInfo struct for BTDiscountInfo
type BTDiscountInfo struct {
	AccountBalance    *int32             `json:"accountBalance,omitempty"`
	AmountOff         *int32             `json:"amountOff,omitempty"`
	CouponType        *int32             `json:"couponType,omitempty"`
	CouponValidMonths *int32             `json:"couponValidMonths,omitempty"`
	CreatedAt         *JSONTime          `json:"createdAt,omitempty"`
	CreatedBy         *BTUserSummaryInfo `json:"createdBy,omitempty"`
	ExpiresAt         *JSONTime          `json:"expiresAt,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// Name of the resource.
	Name         *string   `json:"name,omitempty"`
	OwnerId      *string   `json:"ownerId,omitempty"`
	PercentOff   *int32    `json:"percentOff,omitempty"`
	PlanId       *string   `json:"planId,omitempty"`
	TrialEndDate *string   `json:"trialEndDate,omitempty"`
	UsedAt       *JSONTime `json:"usedAt,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTDiscountInfo instantiates a new BTDiscountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDiscountInfo() *BTDiscountInfo {
	this := BTDiscountInfo{}
	return &this
}

// NewBTDiscountInfoWithDefaults instantiates a new BTDiscountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDiscountInfoWithDefaults() *BTDiscountInfo {
	this := BTDiscountInfo{}
	return &this
}

// GetAccountBalance returns the AccountBalance field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetAccountBalance() int32 {
	if o == nil || o.AccountBalance == nil {
		var ret int32
		return ret
	}
	return *o.AccountBalance
}

// GetAccountBalanceOk returns a tuple with the AccountBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetAccountBalanceOk() (*int32, bool) {
	if o == nil || o.AccountBalance == nil {
		return nil, false
	}
	return o.AccountBalance, true
}

// HasAccountBalance returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasAccountBalance() bool {
	if o != nil && o.AccountBalance != nil {
		return true
	}

	return false
}

// SetAccountBalance gets a reference to the given int32 and assigns it to the AccountBalance field.
func (o *BTDiscountInfo) SetAccountBalance(v int32) {
	o.AccountBalance = &v
}

// GetAmountOff returns the AmountOff field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetAmountOff() int32 {
	if o == nil || o.AmountOff == nil {
		var ret int32
		return ret
	}
	return *o.AmountOff
}

// GetAmountOffOk returns a tuple with the AmountOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetAmountOffOk() (*int32, bool) {
	if o == nil || o.AmountOff == nil {
		return nil, false
	}
	return o.AmountOff, true
}

// HasAmountOff returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasAmountOff() bool {
	if o != nil && o.AmountOff != nil {
		return true
	}

	return false
}

// SetAmountOff gets a reference to the given int32 and assigns it to the AmountOff field.
func (o *BTDiscountInfo) SetAmountOff(v int32) {
	o.AmountOff = &v
}

// GetCouponType returns the CouponType field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetCouponType() int32 {
	if o == nil || o.CouponType == nil {
		var ret int32
		return ret
	}
	return *o.CouponType
}

// GetCouponTypeOk returns a tuple with the CouponType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetCouponTypeOk() (*int32, bool) {
	if o == nil || o.CouponType == nil {
		return nil, false
	}
	return o.CouponType, true
}

// HasCouponType returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasCouponType() bool {
	if o != nil && o.CouponType != nil {
		return true
	}

	return false
}

// SetCouponType gets a reference to the given int32 and assigns it to the CouponType field.
func (o *BTDiscountInfo) SetCouponType(v int32) {
	o.CouponType = &v
}

// GetCouponValidMonths returns the CouponValidMonths field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetCouponValidMonths() int32 {
	if o == nil || o.CouponValidMonths == nil {
		var ret int32
		return ret
	}
	return *o.CouponValidMonths
}

// GetCouponValidMonthsOk returns a tuple with the CouponValidMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetCouponValidMonthsOk() (*int32, bool) {
	if o == nil || o.CouponValidMonths == nil {
		return nil, false
	}
	return o.CouponValidMonths, true
}

// HasCouponValidMonths returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasCouponValidMonths() bool {
	if o != nil && o.CouponValidMonths != nil {
		return true
	}

	return false
}

// SetCouponValidMonths gets a reference to the given int32 and assigns it to the CouponValidMonths field.
func (o *BTDiscountInfo) SetCouponValidMonths(v int32) {
	o.CouponValidMonths = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTDiscountInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetCreatedBy() BTUserSummaryInfo {
	if o == nil || o.CreatedBy == nil {
		var ret BTUserSummaryInfo
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetCreatedByOk() (*BTUserSummaryInfo, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given BTUserSummaryInfo and assigns it to the CreatedBy field.
func (o *BTDiscountInfo) SetCreatedBy(v BTUserSummaryInfo) {
	o.CreatedBy = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetExpiresAt() JSONTime {
	if o == nil || o.ExpiresAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetExpiresAtOk() (*JSONTime, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given JSONTime and assigns it to the ExpiresAt field.
func (o *BTDiscountInfo) SetExpiresAt(v JSONTime) {
	o.ExpiresAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTDiscountInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTDiscountInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTDiscountInfo) SetName(v string) {
	o.Name = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *BTDiscountInfo) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetPercentOff returns the PercentOff field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetPercentOff() int32 {
	if o == nil || o.PercentOff == nil {
		var ret int32
		return ret
	}
	return *o.PercentOff
}

// GetPercentOffOk returns a tuple with the PercentOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetPercentOffOk() (*int32, bool) {
	if o == nil || o.PercentOff == nil {
		return nil, false
	}
	return o.PercentOff, true
}

// HasPercentOff returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasPercentOff() bool {
	if o != nil && o.PercentOff != nil {
		return true
	}

	return false
}

// SetPercentOff gets a reference to the given int32 and assigns it to the PercentOff field.
func (o *BTDiscountInfo) SetPercentOff(v int32) {
	o.PercentOff = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetPlanId() string {
	if o == nil || o.PlanId == nil {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetPlanIdOk() (*string, bool) {
	if o == nil || o.PlanId == nil {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasPlanId() bool {
	if o != nil && o.PlanId != nil {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *BTDiscountInfo) SetPlanId(v string) {
	o.PlanId = &v
}

// GetTrialEndDate returns the TrialEndDate field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetTrialEndDate() string {
	if o == nil || o.TrialEndDate == nil {
		var ret string
		return ret
	}
	return *o.TrialEndDate
}

// GetTrialEndDateOk returns a tuple with the TrialEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetTrialEndDateOk() (*string, bool) {
	if o == nil || o.TrialEndDate == nil {
		return nil, false
	}
	return o.TrialEndDate, true
}

// HasTrialEndDate returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasTrialEndDate() bool {
	if o != nil && o.TrialEndDate != nil {
		return true
	}

	return false
}

// SetTrialEndDate gets a reference to the given string and assigns it to the TrialEndDate field.
func (o *BTDiscountInfo) SetTrialEndDate(v string) {
	o.TrialEndDate = &v
}

// GetUsedAt returns the UsedAt field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetUsedAt() JSONTime {
	if o == nil || o.UsedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.UsedAt
}

// GetUsedAtOk returns a tuple with the UsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetUsedAtOk() (*JSONTime, bool) {
	if o == nil || o.UsedAt == nil {
		return nil, false
	}
	return o.UsedAt, true
}

// HasUsedAt returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasUsedAt() bool {
	if o != nil && o.UsedAt != nil {
		return true
	}

	return false
}

// SetUsedAt gets a reference to the given JSONTime and assigns it to the UsedAt field.
func (o *BTDiscountInfo) SetUsedAt(v JSONTime) {
	o.UsedAt = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTDiscountInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTDiscountInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTDiscountInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTDiscountInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountBalance != nil {
		toSerialize["accountBalance"] = o.AccountBalance
	}
	if o.AmountOff != nil {
		toSerialize["amountOff"] = o.AmountOff
	}
	if o.CouponType != nil {
		toSerialize["couponType"] = o.CouponType
	}
	if o.CouponValidMonths != nil {
		toSerialize["couponValidMonths"] = o.CouponValidMonths
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.PercentOff != nil {
		toSerialize["percentOff"] = o.PercentOff
	}
	if o.PlanId != nil {
		toSerialize["planId"] = o.PlanId
	}
	if o.TrialEndDate != nil {
		toSerialize["trialEndDate"] = o.TrialEndDate
	}
	if o.UsedAt != nil {
		toSerialize["usedAt"] = o.UsedAt
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTDiscountInfo struct {
	value *BTDiscountInfo
	isSet bool
}

func (v NullableBTDiscountInfo) Get() *BTDiscountInfo {
	return v.value
}

func (v *NullableBTDiscountInfo) Set(val *BTDiscountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDiscountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDiscountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDiscountInfo(val *BTDiscountInfo) *NullableBTDiscountInfo {
	return &NullableBTDiscountInfo{value: val, isSet: true}
}

func (v NullableBTDiscountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDiscountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
