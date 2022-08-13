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

// BTParameterSpecReferencePartStudio1256 struct for BTParameterSpecReferencePartStudio1256
type BTParameterSpecReferencePartStudio1256 struct {
	AdditionalLocalizedStrings  *int32                                 `json:"additionalLocalizedStrings,omitempty"`
	ColumnName                  *string                                `json:"columnName,omitempty"`
	DefaultValue                *BTMParameter1                         `json:"defaultValue,omitempty"`
	IconUri                     *string                                `json:"iconUri,omitempty"`
	LocalizableName             *string                                `json:"localizableName,omitempty"`
	LocalizedName               *string                                `json:"localizedName,omitempty"`
	ParameterDescription        *string                                `json:"parameterDescription,omitempty"`
	ParameterId                 *string                                `json:"parameterId,omitempty"`
	ParameterName               *string                                `json:"parameterName,omitempty"`
	StringsToLocalize           []string                               `json:"stringsToLocalize,omitempty"`
	UiHint                      *string                                `json:"uiHint,omitempty"`
	UiHints                     []string                               `json:"uiHints,omitempty"`
	VisibilityCondition         *BTParameterVisibilityCondition177     `json:"visibilityCondition,omitempty"`
	BtType                      *string                                `json:"btType,omitempty"`
	DefaultPurpose              *BTElementLibraryPurpose3353           `json:"defaultPurpose,omitempty"`
	AllowedInsertableTypes      []string                               `json:"allowedInsertableTypes,omitempty"`
	ComputedConfigurationInputs []BTComputedConfigurationInputSpec2525 `json:"computedConfigurationInputs,omitempty"`
	MaxNumberOfPicks            *int32                                 `json:"maxNumberOfPicks,omitempty"`
}

// NewBTParameterSpecReferencePartStudio1256 instantiates a new BTParameterSpecReferencePartStudio1256 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterSpecReferencePartStudio1256() *BTParameterSpecReferencePartStudio1256 {
	this := BTParameterSpecReferencePartStudio1256{}
	return &this
}

// NewBTParameterSpecReferencePartStudio1256WithDefaults instantiates a new BTParameterSpecReferencePartStudio1256 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterSpecReferencePartStudio1256WithDefaults() *BTParameterSpecReferencePartStudio1256 {
	this := BTParameterSpecReferencePartStudio1256{}
	return &this
}

// GetAdditionalLocalizedStrings returns the AdditionalLocalizedStrings field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetAdditionalLocalizedStrings() int32 {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		var ret int32
		return ret
	}
	return *o.AdditionalLocalizedStrings
}

// GetAdditionalLocalizedStringsOk returns a tuple with the AdditionalLocalizedStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetAdditionalLocalizedStringsOk() (*int32, bool) {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		return nil, false
	}
	return o.AdditionalLocalizedStrings, true
}

// HasAdditionalLocalizedStrings returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasAdditionalLocalizedStrings() bool {
	if o != nil && o.AdditionalLocalizedStrings != nil {
		return true
	}

	return false
}

// SetAdditionalLocalizedStrings gets a reference to the given int32 and assigns it to the AdditionalLocalizedStrings field.
func (o *BTParameterSpecReferencePartStudio1256) SetAdditionalLocalizedStrings(v int32) {
	o.AdditionalLocalizedStrings = &v
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasColumnName() bool {
	if o != nil && o.ColumnName != nil {
		return true
	}

	return false
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *BTParameterSpecReferencePartStudio1256) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetDefaultValue() BTMParameter1 {
	if o == nil || o.DefaultValue == nil {
		var ret BTMParameter1
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetDefaultValueOk() (*BTMParameter1, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given BTMParameter1 and assigns it to the DefaultValue field.
func (o *BTParameterSpecReferencePartStudio1256) SetDefaultValue(v BTMParameter1) {
	o.DefaultValue = &v
}

// GetIconUri returns the IconUri field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetIconUri() string {
	if o == nil || o.IconUri == nil {
		var ret string
		return ret
	}
	return *o.IconUri
}

// GetIconUriOk returns a tuple with the IconUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetIconUriOk() (*string, bool) {
	if o == nil || o.IconUri == nil {
		return nil, false
	}
	return o.IconUri, true
}

// HasIconUri returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasIconUri() bool {
	if o != nil && o.IconUri != nil {
		return true
	}

	return false
}

// SetIconUri gets a reference to the given string and assigns it to the IconUri field.
func (o *BTParameterSpecReferencePartStudio1256) SetIconUri(v string) {
	o.IconUri = &v
}

// GetLocalizableName returns the LocalizableName field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetLocalizableName() string {
	if o == nil || o.LocalizableName == nil {
		var ret string
		return ret
	}
	return *o.LocalizableName
}

// GetLocalizableNameOk returns a tuple with the LocalizableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetLocalizableNameOk() (*string, bool) {
	if o == nil || o.LocalizableName == nil {
		return nil, false
	}
	return o.LocalizableName, true
}

// HasLocalizableName returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasLocalizableName() bool {
	if o != nil && o.LocalizableName != nil {
		return true
	}

	return false
}

// SetLocalizableName gets a reference to the given string and assigns it to the LocalizableName field.
func (o *BTParameterSpecReferencePartStudio1256) SetLocalizableName(v string) {
	o.LocalizableName = &v
}

// GetLocalizedName returns the LocalizedName field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetLocalizedName() string {
	if o == nil || o.LocalizedName == nil {
		var ret string
		return ret
	}
	return *o.LocalizedName
}

// GetLocalizedNameOk returns a tuple with the LocalizedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetLocalizedNameOk() (*string, bool) {
	if o == nil || o.LocalizedName == nil {
		return nil, false
	}
	return o.LocalizedName, true
}

// HasLocalizedName returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasLocalizedName() bool {
	if o != nil && o.LocalizedName != nil {
		return true
	}

	return false
}

// SetLocalizedName gets a reference to the given string and assigns it to the LocalizedName field.
func (o *BTParameterSpecReferencePartStudio1256) SetLocalizedName(v string) {
	o.LocalizedName = &v
}

// GetParameterDescription returns the ParameterDescription field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetParameterDescription() string {
	if o == nil || o.ParameterDescription == nil {
		var ret string
		return ret
	}
	return *o.ParameterDescription
}

// GetParameterDescriptionOk returns a tuple with the ParameterDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetParameterDescriptionOk() (*string, bool) {
	if o == nil || o.ParameterDescription == nil {
		return nil, false
	}
	return o.ParameterDescription, true
}

// HasParameterDescription returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasParameterDescription() bool {
	if o != nil && o.ParameterDescription != nil {
		return true
	}

	return false
}

// SetParameterDescription gets a reference to the given string and assigns it to the ParameterDescription field.
func (o *BTParameterSpecReferencePartStudio1256) SetParameterDescription(v string) {
	o.ParameterDescription = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTParameterSpecReferencePartStudio1256) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetParameterName() string {
	if o == nil || o.ParameterName == nil {
		var ret string
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetParameterNameOk() (*string, bool) {
	if o == nil || o.ParameterName == nil {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasParameterName() bool {
	if o != nil && o.ParameterName != nil {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *BTParameterSpecReferencePartStudio1256) SetParameterName(v string) {
	o.ParameterName = &v
}

// GetStringsToLocalize returns the StringsToLocalize field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetStringsToLocalize() []string {
	if o == nil || o.StringsToLocalize == nil {
		var ret []string
		return ret
	}
	return o.StringsToLocalize
}

// GetStringsToLocalizeOk returns a tuple with the StringsToLocalize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetStringsToLocalizeOk() ([]string, bool) {
	if o == nil || o.StringsToLocalize == nil {
		return nil, false
	}
	return o.StringsToLocalize, true
}

// HasStringsToLocalize returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasStringsToLocalize() bool {
	if o != nil && o.StringsToLocalize != nil {
		return true
	}

	return false
}

// SetStringsToLocalize gets a reference to the given []string and assigns it to the StringsToLocalize field.
func (o *BTParameterSpecReferencePartStudio1256) SetStringsToLocalize(v []string) {
	o.StringsToLocalize = v
}

// GetUiHint returns the UiHint field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetUiHint() string {
	if o == nil || o.UiHint == nil {
		var ret string
		return ret
	}
	return *o.UiHint
}

// GetUiHintOk returns a tuple with the UiHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetUiHintOk() (*string, bool) {
	if o == nil || o.UiHint == nil {
		return nil, false
	}
	return o.UiHint, true
}

// HasUiHint returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasUiHint() bool {
	if o != nil && o.UiHint != nil {
		return true
	}

	return false
}

// SetUiHint gets a reference to the given string and assigns it to the UiHint field.
func (o *BTParameterSpecReferencePartStudio1256) SetUiHint(v string) {
	o.UiHint = &v
}

// GetUiHints returns the UiHints field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetUiHints() []string {
	if o == nil || o.UiHints == nil {
		var ret []string
		return ret
	}
	return o.UiHints
}

// GetUiHintsOk returns a tuple with the UiHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetUiHintsOk() ([]string, bool) {
	if o == nil || o.UiHints == nil {
		return nil, false
	}
	return o.UiHints, true
}

// HasUiHints returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasUiHints() bool {
	if o != nil && o.UiHints != nil {
		return true
	}

	return false
}

// SetUiHints gets a reference to the given []string and assigns it to the UiHints field.
func (o *BTParameterSpecReferencePartStudio1256) SetUiHints(v []string) {
	o.UiHints = v
}

// GetVisibilityCondition returns the VisibilityCondition field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetVisibilityCondition() BTParameterVisibilityCondition177 {
	if o == nil || o.VisibilityCondition == nil {
		var ret BTParameterVisibilityCondition177
		return ret
	}
	return *o.VisibilityCondition
}

// GetVisibilityConditionOk returns a tuple with the VisibilityCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetVisibilityConditionOk() (*BTParameterVisibilityCondition177, bool) {
	if o == nil || o.VisibilityCondition == nil {
		return nil, false
	}
	return o.VisibilityCondition, true
}

// HasVisibilityCondition returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasVisibilityCondition() bool {
	if o != nil && o.VisibilityCondition != nil {
		return true
	}

	return false
}

// SetVisibilityCondition gets a reference to the given BTParameterVisibilityCondition177 and assigns it to the VisibilityCondition field.
func (o *BTParameterSpecReferencePartStudio1256) SetVisibilityCondition(v BTParameterVisibilityCondition177) {
	o.VisibilityCondition = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterSpecReferencePartStudio1256) SetBtType(v string) {
	o.BtType = &v
}

// GetDefaultPurpose returns the DefaultPurpose field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetDefaultPurpose() BTElementLibraryPurpose3353 {
	if o == nil || o.DefaultPurpose == nil {
		var ret BTElementLibraryPurpose3353
		return ret
	}
	return *o.DefaultPurpose
}

// GetDefaultPurposeOk returns a tuple with the DefaultPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetDefaultPurposeOk() (*BTElementLibraryPurpose3353, bool) {
	if o == nil || o.DefaultPurpose == nil {
		return nil, false
	}
	return o.DefaultPurpose, true
}

// HasDefaultPurpose returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasDefaultPurpose() bool {
	if o != nil && o.DefaultPurpose != nil {
		return true
	}

	return false
}

// SetDefaultPurpose gets a reference to the given BTElementLibraryPurpose3353 and assigns it to the DefaultPurpose field.
func (o *BTParameterSpecReferencePartStudio1256) SetDefaultPurpose(v BTElementLibraryPurpose3353) {
	o.DefaultPurpose = &v
}

// GetAllowedInsertableTypes returns the AllowedInsertableTypes field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetAllowedInsertableTypes() []string {
	if o == nil || o.AllowedInsertableTypes == nil {
		var ret []string
		return ret
	}
	return o.AllowedInsertableTypes
}

// GetAllowedInsertableTypesOk returns a tuple with the AllowedInsertableTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetAllowedInsertableTypesOk() ([]string, bool) {
	if o == nil || o.AllowedInsertableTypes == nil {
		return nil, false
	}
	return o.AllowedInsertableTypes, true
}

// HasAllowedInsertableTypes returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasAllowedInsertableTypes() bool {
	if o != nil && o.AllowedInsertableTypes != nil {
		return true
	}

	return false
}

// SetAllowedInsertableTypes gets a reference to the given []string and assigns it to the AllowedInsertableTypes field.
func (o *BTParameterSpecReferencePartStudio1256) SetAllowedInsertableTypes(v []string) {
	o.AllowedInsertableTypes = v
}

// GetComputedConfigurationInputs returns the ComputedConfigurationInputs field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetComputedConfigurationInputs() []BTComputedConfigurationInputSpec2525 {
	if o == nil || o.ComputedConfigurationInputs == nil {
		var ret []BTComputedConfigurationInputSpec2525
		return ret
	}
	return o.ComputedConfigurationInputs
}

// GetComputedConfigurationInputsOk returns a tuple with the ComputedConfigurationInputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetComputedConfigurationInputsOk() ([]BTComputedConfigurationInputSpec2525, bool) {
	if o == nil || o.ComputedConfigurationInputs == nil {
		return nil, false
	}
	return o.ComputedConfigurationInputs, true
}

// HasComputedConfigurationInputs returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasComputedConfigurationInputs() bool {
	if o != nil && o.ComputedConfigurationInputs != nil {
		return true
	}

	return false
}

// SetComputedConfigurationInputs gets a reference to the given []BTComputedConfigurationInputSpec2525 and assigns it to the ComputedConfigurationInputs field.
func (o *BTParameterSpecReferencePartStudio1256) SetComputedConfigurationInputs(v []BTComputedConfigurationInputSpec2525) {
	o.ComputedConfigurationInputs = v
}

// GetMaxNumberOfPicks returns the MaxNumberOfPicks field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetMaxNumberOfPicks() int32 {
	if o == nil || o.MaxNumberOfPicks == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumberOfPicks
}

// GetMaxNumberOfPicksOk returns a tuple with the MaxNumberOfPicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetMaxNumberOfPicksOk() (*int32, bool) {
	if o == nil || o.MaxNumberOfPicks == nil {
		return nil, false
	}
	return o.MaxNumberOfPicks, true
}

// HasMaxNumberOfPicks returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasMaxNumberOfPicks() bool {
	if o != nil && o.MaxNumberOfPicks != nil {
		return true
	}

	return false
}

// SetMaxNumberOfPicks gets a reference to the given int32 and assigns it to the MaxNumberOfPicks field.
func (o *BTParameterSpecReferencePartStudio1256) SetMaxNumberOfPicks(v int32) {
	o.MaxNumberOfPicks = &v
}

func (o BTParameterSpecReferencePartStudio1256) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalLocalizedStrings != nil {
		toSerialize["additionalLocalizedStrings"] = o.AdditionalLocalizedStrings
	}
	if o.ColumnName != nil {
		toSerialize["columnName"] = o.ColumnName
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.IconUri != nil {
		toSerialize["iconUri"] = o.IconUri
	}
	if o.LocalizableName != nil {
		toSerialize["localizableName"] = o.LocalizableName
	}
	if o.LocalizedName != nil {
		toSerialize["localizedName"] = o.LocalizedName
	}
	if o.ParameterDescription != nil {
		toSerialize["parameterDescription"] = o.ParameterDescription
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.ParameterName != nil {
		toSerialize["parameterName"] = o.ParameterName
	}
	if o.StringsToLocalize != nil {
		toSerialize["stringsToLocalize"] = o.StringsToLocalize
	}
	if o.UiHint != nil {
		toSerialize["uiHint"] = o.UiHint
	}
	if o.UiHints != nil {
		toSerialize["uiHints"] = o.UiHints
	}
	if o.VisibilityCondition != nil {
		toSerialize["visibilityCondition"] = o.VisibilityCondition
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DefaultPurpose != nil {
		toSerialize["defaultPurpose"] = o.DefaultPurpose
	}
	if o.AllowedInsertableTypes != nil {
		toSerialize["allowedInsertableTypes"] = o.AllowedInsertableTypes
	}
	if o.ComputedConfigurationInputs != nil {
		toSerialize["computedConfigurationInputs"] = o.ComputedConfigurationInputs
	}
	if o.MaxNumberOfPicks != nil {
		toSerialize["maxNumberOfPicks"] = o.MaxNumberOfPicks
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterSpecReferencePartStudio1256 struct {
	value *BTParameterSpecReferencePartStudio1256
	isSet bool
}

func (v NullableBTParameterSpecReferencePartStudio1256) Get() *BTParameterSpecReferencePartStudio1256 {
	return v.value
}

func (v *NullableBTParameterSpecReferencePartStudio1256) Set(val *BTParameterSpecReferencePartStudio1256) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterSpecReferencePartStudio1256) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterSpecReferencePartStudio1256) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterSpecReferencePartStudio1256(val *BTParameterSpecReferencePartStudio1256) *NullableBTParameterSpecReferencePartStudio1256 {
	return &NullableBTParameterSpecReferencePartStudio1256{value: val, isSet: true}
}

func (v NullableBTParameterSpecReferencePartStudio1256) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterSpecReferencePartStudio1256) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
