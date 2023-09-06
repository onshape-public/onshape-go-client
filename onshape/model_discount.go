/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.169.21759-9ddbba9fdfb8
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Discount struct for Discount
type Discount struct {
	Coupon       *Coupon `json:"coupon,omitempty"`
	Customer     *string `json:"customer,omitempty"`
	End          *int64  `json:"end,omitempty"`
	Id           *string `json:"id,omitempty"`
	Object       *string `json:"object,omitempty"`
	Start        *int64  `json:"start,omitempty"`
	Subscription *string `json:"subscription,omitempty"`
}

// NewDiscount instantiates a new Discount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscount() *Discount {
	this := Discount{}
	return &this
}

// NewDiscountWithDefaults instantiates a new Discount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountWithDefaults() *Discount {
	this := Discount{}
	return &this
}

// GetCoupon returns the Coupon field value if set, zero value otherwise.
func (o *Discount) GetCoupon() Coupon {
	if o == nil || o.Coupon == nil {
		var ret Coupon
		return ret
	}
	return *o.Coupon
}

// GetCouponOk returns a tuple with the Coupon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetCouponOk() (*Coupon, bool) {
	if o == nil || o.Coupon == nil {
		return nil, false
	}
	return o.Coupon, true
}

// HasCoupon returns a boolean if a field has been set.
func (o *Discount) HasCoupon() bool {
	if o != nil && o.Coupon != nil {
		return true
	}

	return false
}

// SetCoupon gets a reference to the given Coupon and assigns it to the Coupon field.
func (o *Discount) SetCoupon(v Coupon) {
	o.Coupon = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *Discount) GetCustomer() string {
	if o == nil || o.Customer == nil {
		var ret string
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetCustomerOk() (*string, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *Discount) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given string and assigns it to the Customer field.
func (o *Discount) SetCustomer(v string) {
	o.Customer = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Discount) GetEnd() int64 {
	if o == nil || o.End == nil {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetEndOk() (*int64, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Discount) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *Discount) SetEnd(v int64) {
	o.End = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Discount) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Discount) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Discount) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *Discount) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *Discount) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *Discount) SetObject(v string) {
	o.Object = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *Discount) GetStart() int64 {
	if o == nil || o.Start == nil {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetStartOk() (*int64, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Discount) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *Discount) SetStart(v int64) {
	o.Start = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *Discount) GetSubscription() string {
	if o == nil || o.Subscription == nil {
		var ret string
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetSubscriptionOk() (*string, bool) {
	if o == nil || o.Subscription == nil {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *Discount) HasSubscription() bool {
	if o != nil && o.Subscription != nil {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given string and assigns it to the Subscription field.
func (o *Discount) SetSubscription(v string) {
	o.Subscription = &v
}

func (o Discount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Coupon != nil {
		toSerialize["coupon"] = o.Coupon
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.Subscription != nil {
		toSerialize["subscription"] = o.Subscription
	}
	return json.Marshal(toSerialize)
}

type NullableDiscount struct {
	value *Discount
	isSet bool
}

func (v NullableDiscount) Get() *Discount {
	return v.value
}

func (v *NullableDiscount) Set(val *Discount) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscount(val *Discount) *NullableDiscount {
	return &NullableDiscount{value: val, isSet: true}
}

func (v NullableDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
