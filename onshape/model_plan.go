/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Plan struct for Plan
type Plan struct {
	Amount               *int64             `json:"amount,omitempty"`
	Created              *int64             `json:"created,omitempty"`
	Currency             *string            `json:"currency,omitempty"`
	Id                   *string            `json:"id,omitempty"`
	Interval             *string            `json:"interval,omitempty"`
	IntervalCount        *int32             `json:"intervalCount,omitempty"`
	Livemode             *bool              `json:"livemode,omitempty"`
	Metadata             *map[string]string `json:"metadata,omitempty"`
	Name                 *string            `json:"name,omitempty"`
	Object               *string            `json:"object,omitempty"`
	StatementDescription *string            `json:"statementDescription,omitempty"`
	StatementDescriptor  *string            `json:"statementDescriptor,omitempty"`
	TrialPeriodDays      *int32             `json:"trialPeriodDays,omitempty"`
}

// NewPlan instantiates a new Plan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlan() *Plan {
	this := Plan{}
	return &this
}

// NewPlanWithDefaults instantiates a new Plan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanWithDefaults() *Plan {
	this := Plan{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Plan) GetAmount() int64 {
	if o == nil || o.Amount == nil {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetAmountOk() (*int64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Plan) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *Plan) SetAmount(v int64) {
	o.Amount = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Plan) GetCreated() int64 {
	if o == nil || o.Created == nil {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetCreatedOk() (*int64, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Plan) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *Plan) SetCreated(v int64) {
	o.Created = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Plan) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Plan) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Plan) SetCurrency(v string) {
	o.Currency = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Plan) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Plan) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Plan) SetId(v string) {
	o.Id = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *Plan) GetInterval() string {
	if o == nil || o.Interval == nil {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetIntervalOk() (*string, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *Plan) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *Plan) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalCount returns the IntervalCount field value if set, zero value otherwise.
func (o *Plan) GetIntervalCount() int32 {
	if o == nil || o.IntervalCount == nil {
		var ret int32
		return ret
	}
	return *o.IntervalCount
}

// GetIntervalCountOk returns a tuple with the IntervalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetIntervalCountOk() (*int32, bool) {
	if o == nil || o.IntervalCount == nil {
		return nil, false
	}
	return o.IntervalCount, true
}

// HasIntervalCount returns a boolean if a field has been set.
func (o *Plan) HasIntervalCount() bool {
	if o != nil && o.IntervalCount != nil {
		return true
	}

	return false
}

// SetIntervalCount gets a reference to the given int32 and assigns it to the IntervalCount field.
func (o *Plan) SetIntervalCount(v int32) {
	o.IntervalCount = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *Plan) GetLivemode() bool {
	if o == nil || o.Livemode == nil {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetLivemodeOk() (*bool, bool) {
	if o == nil || o.Livemode == nil {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *Plan) HasLivemode() bool {
	if o != nil && o.Livemode != nil {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *Plan) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Plan) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Plan) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *Plan) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Plan) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Plan) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Plan) SetName(v string) {
	o.Name = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *Plan) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *Plan) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *Plan) SetObject(v string) {
	o.Object = &v
}

// GetStatementDescription returns the StatementDescription field value if set, zero value otherwise.
func (o *Plan) GetStatementDescription() string {
	if o == nil || o.StatementDescription == nil {
		var ret string
		return ret
	}
	return *o.StatementDescription
}

// GetStatementDescriptionOk returns a tuple with the StatementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetStatementDescriptionOk() (*string, bool) {
	if o == nil || o.StatementDescription == nil {
		return nil, false
	}
	return o.StatementDescription, true
}

// HasStatementDescription returns a boolean if a field has been set.
func (o *Plan) HasStatementDescription() bool {
	if o != nil && o.StatementDescription != nil {
		return true
	}

	return false
}

// SetStatementDescription gets a reference to the given string and assigns it to the StatementDescription field.
func (o *Plan) SetStatementDescription(v string) {
	o.StatementDescription = &v
}

// GetStatementDescriptor returns the StatementDescriptor field value if set, zero value otherwise.
func (o *Plan) GetStatementDescriptor() string {
	if o == nil || o.StatementDescriptor == nil {
		var ret string
		return ret
	}
	return *o.StatementDescriptor
}

// GetStatementDescriptorOk returns a tuple with the StatementDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetStatementDescriptorOk() (*string, bool) {
	if o == nil || o.StatementDescriptor == nil {
		return nil, false
	}
	return o.StatementDescriptor, true
}

// HasStatementDescriptor returns a boolean if a field has been set.
func (o *Plan) HasStatementDescriptor() bool {
	if o != nil && o.StatementDescriptor != nil {
		return true
	}

	return false
}

// SetStatementDescriptor gets a reference to the given string and assigns it to the StatementDescriptor field.
func (o *Plan) SetStatementDescriptor(v string) {
	o.StatementDescriptor = &v
}

// GetTrialPeriodDays returns the TrialPeriodDays field value if set, zero value otherwise.
func (o *Plan) GetTrialPeriodDays() int32 {
	if o == nil || o.TrialPeriodDays == nil {
		var ret int32
		return ret
	}
	return *o.TrialPeriodDays
}

// GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetTrialPeriodDaysOk() (*int32, bool) {
	if o == nil || o.TrialPeriodDays == nil {
		return nil, false
	}
	return o.TrialPeriodDays, true
}

// HasTrialPeriodDays returns a boolean if a field has been set.
func (o *Plan) HasTrialPeriodDays() bool {
	if o != nil && o.TrialPeriodDays != nil {
		return true
	}

	return false
}

// SetTrialPeriodDays gets a reference to the given int32 and assigns it to the TrialPeriodDays field.
func (o *Plan) SetTrialPeriodDays(v int32) {
	o.TrialPeriodDays = &v
}

func (o Plan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.IntervalCount != nil {
		toSerialize["intervalCount"] = o.IntervalCount
	}
	if o.Livemode != nil {
		toSerialize["livemode"] = o.Livemode
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
	if o.StatementDescription != nil {
		toSerialize["statementDescription"] = o.StatementDescription
	}
	if o.StatementDescriptor != nil {
		toSerialize["statementDescriptor"] = o.StatementDescriptor
	}
	if o.TrialPeriodDays != nil {
		toSerialize["trialPeriodDays"] = o.TrialPeriodDays
	}
	return json.Marshal(toSerialize)
}

type NullablePlan struct {
	value *Plan
	isSet bool
}

func (v NullablePlan) Get() *Plan {
	return v.value
}

func (v *NullablePlan) Set(val *Plan) {
	v.value = val
	v.isSet = true
}

func (v NullablePlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlan(val *Plan) *NullablePlan {
	return &NullablePlan{value: val, isSet: true}
}

func (v NullablePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
