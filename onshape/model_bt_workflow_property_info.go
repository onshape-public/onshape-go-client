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

// BTWorkflowPropertyInfo struct for BTWorkflowPropertyInfo
type BTWorkflowPropertyInfo struct {
	ComputedAssemblyProperty *bool                            `json:"computedAssemblyProperty,omitempty"`
	ComputedProperty         *bool                            `json:"computedProperty,omitempty"`
	ComputedPropertyError    *string                          `json:"computedPropertyError,omitempty"`
	DefaultValue             *map[string]interface{}          `json:"defaultValue,omitempty"`
	Dirty                    *bool                            `json:"dirty,omitempty"`
	Editable                 *bool                            `json:"editable,omitempty"`
	EditableInUi             *bool                            `json:"editableInUi,omitempty"`
	EnumValues               []BTMetadataEnumValueInfo        `json:"enumValues,omitempty"`
	InitialValue             *map[string]interface{}          `json:"initialValue,omitempty"`
	IsApproverProperty       *bool                            `json:"isApproverProperty,omitempty"`
	IsNotifierProperty       *bool                            `json:"isNotifierProperty,omitempty"`
	Multivalued              *bool                            `json:"multivalued,omitempty"`
	Name                     *string                          `json:"name,omitempty"`
	Observers                []BTWorkflowableObjectObserver   `json:"observers,omitempty"`
	PropertyId               *string                          `json:"propertyId,omitempty"`
	PropertySource           *int32                           `json:"propertySource,omitempty"`
	Required                 *bool                            `json:"required,omitempty"`
	SchemaId                 *string                          `json:"schemaId,omitempty"`
	TeamsOnly                *bool                            `json:"teamsOnly,omitempty"`
	UiHints                  *BTMetadataPropertyUiHintsInfo   `json:"uiHints,omitempty"`
	UsersOnly                *bool                            `json:"usersOnly,omitempty"`
	Validator                *BTMetadataPropertyValidatorInfo `json:"validator,omitempty"`
	Value                    *map[string]interface{}          `json:"value,omitempty"`
	ValueType                *string                          `json:"valueType,omitempty"`
}

// NewBTWorkflowPropertyInfo instantiates a new BTWorkflowPropertyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWorkflowPropertyInfo() *BTWorkflowPropertyInfo {
	this := BTWorkflowPropertyInfo{}
	return &this
}

// NewBTWorkflowPropertyInfoWithDefaults instantiates a new BTWorkflowPropertyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWorkflowPropertyInfoWithDefaults() *BTWorkflowPropertyInfo {
	this := BTWorkflowPropertyInfo{}
	return &this
}

// GetComputedAssemblyProperty returns the ComputedAssemblyProperty field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetComputedAssemblyProperty() bool {
	if o == nil || o.ComputedAssemblyProperty == nil {
		var ret bool
		return ret
	}
	return *o.ComputedAssemblyProperty
}

// GetComputedAssemblyPropertyOk returns a tuple with the ComputedAssemblyProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetComputedAssemblyPropertyOk() (*bool, bool) {
	if o == nil || o.ComputedAssemblyProperty == nil {
		return nil, false
	}
	return o.ComputedAssemblyProperty, true
}

// HasComputedAssemblyProperty returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasComputedAssemblyProperty() bool {
	if o != nil && o.ComputedAssemblyProperty != nil {
		return true
	}

	return false
}

// SetComputedAssemblyProperty gets a reference to the given bool and assigns it to the ComputedAssemblyProperty field.
func (o *BTWorkflowPropertyInfo) SetComputedAssemblyProperty(v bool) {
	o.ComputedAssemblyProperty = &v
}

// GetComputedProperty returns the ComputedProperty field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetComputedProperty() bool {
	if o == nil || o.ComputedProperty == nil {
		var ret bool
		return ret
	}
	return *o.ComputedProperty
}

// GetComputedPropertyOk returns a tuple with the ComputedProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetComputedPropertyOk() (*bool, bool) {
	if o == nil || o.ComputedProperty == nil {
		return nil, false
	}
	return o.ComputedProperty, true
}

// HasComputedProperty returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasComputedProperty() bool {
	if o != nil && o.ComputedProperty != nil {
		return true
	}

	return false
}

// SetComputedProperty gets a reference to the given bool and assigns it to the ComputedProperty field.
func (o *BTWorkflowPropertyInfo) SetComputedProperty(v bool) {
	o.ComputedProperty = &v
}

// GetComputedPropertyError returns the ComputedPropertyError field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetComputedPropertyError() string {
	if o == nil || o.ComputedPropertyError == nil {
		var ret string
		return ret
	}
	return *o.ComputedPropertyError
}

// GetComputedPropertyErrorOk returns a tuple with the ComputedPropertyError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetComputedPropertyErrorOk() (*string, bool) {
	if o == nil || o.ComputedPropertyError == nil {
		return nil, false
	}
	return o.ComputedPropertyError, true
}

// HasComputedPropertyError returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasComputedPropertyError() bool {
	if o != nil && o.ComputedPropertyError != nil {
		return true
	}

	return false
}

// SetComputedPropertyError gets a reference to the given string and assigns it to the ComputedPropertyError field.
func (o *BTWorkflowPropertyInfo) SetComputedPropertyError(v string) {
	o.ComputedPropertyError = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetDefaultValue() map[string]interface{} {
	if o == nil || o.DefaultValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetDefaultValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given map[string]interface{} and assigns it to the DefaultValue field.
func (o *BTWorkflowPropertyInfo) SetDefaultValue(v map[string]interface{}) {
	o.DefaultValue = &v
}

// GetDirty returns the Dirty field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetDirty() bool {
	if o == nil || o.Dirty == nil {
		var ret bool
		return ret
	}
	return *o.Dirty
}

// GetDirtyOk returns a tuple with the Dirty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetDirtyOk() (*bool, bool) {
	if o == nil || o.Dirty == nil {
		return nil, false
	}
	return o.Dirty, true
}

// HasDirty returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasDirty() bool {
	if o != nil && o.Dirty != nil {
		return true
	}

	return false
}

// SetDirty gets a reference to the given bool and assigns it to the Dirty field.
func (o *BTWorkflowPropertyInfo) SetDirty(v bool) {
	o.Dirty = &v
}

// GetEditable returns the Editable field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetEditable() bool {
	if o == nil || o.Editable == nil {
		var ret bool
		return ret
	}
	return *o.Editable
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetEditableOk() (*bool, bool) {
	if o == nil || o.Editable == nil {
		return nil, false
	}
	return o.Editable, true
}

// HasEditable returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasEditable() bool {
	if o != nil && o.Editable != nil {
		return true
	}

	return false
}

// SetEditable gets a reference to the given bool and assigns it to the Editable field.
func (o *BTWorkflowPropertyInfo) SetEditable(v bool) {
	o.Editable = &v
}

// GetEditableInUi returns the EditableInUi field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetEditableInUi() bool {
	if o == nil || o.EditableInUi == nil {
		var ret bool
		return ret
	}
	return *o.EditableInUi
}

// GetEditableInUiOk returns a tuple with the EditableInUi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetEditableInUiOk() (*bool, bool) {
	if o == nil || o.EditableInUi == nil {
		return nil, false
	}
	return o.EditableInUi, true
}

// HasEditableInUi returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasEditableInUi() bool {
	if o != nil && o.EditableInUi != nil {
		return true
	}

	return false
}

// SetEditableInUi gets a reference to the given bool and assigns it to the EditableInUi field.
func (o *BTWorkflowPropertyInfo) SetEditableInUi(v bool) {
	o.EditableInUi = &v
}

// GetEnumValues returns the EnumValues field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetEnumValues() []BTMetadataEnumValueInfo {
	if o == nil || o.EnumValues == nil {
		var ret []BTMetadataEnumValueInfo
		return ret
	}
	return o.EnumValues
}

// GetEnumValuesOk returns a tuple with the EnumValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetEnumValuesOk() ([]BTMetadataEnumValueInfo, bool) {
	if o == nil || o.EnumValues == nil {
		return nil, false
	}
	return o.EnumValues, true
}

// HasEnumValues returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasEnumValues() bool {
	if o != nil && o.EnumValues != nil {
		return true
	}

	return false
}

// SetEnumValues gets a reference to the given []BTMetadataEnumValueInfo and assigns it to the EnumValues field.
func (o *BTWorkflowPropertyInfo) SetEnumValues(v []BTMetadataEnumValueInfo) {
	o.EnumValues = v
}

// GetInitialValue returns the InitialValue field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetInitialValue() map[string]interface{} {
	if o == nil || o.InitialValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.InitialValue
}

// GetInitialValueOk returns a tuple with the InitialValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetInitialValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.InitialValue == nil {
		return nil, false
	}
	return o.InitialValue, true
}

// HasInitialValue returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasInitialValue() bool {
	if o != nil && o.InitialValue != nil {
		return true
	}

	return false
}

// SetInitialValue gets a reference to the given map[string]interface{} and assigns it to the InitialValue field.
func (o *BTWorkflowPropertyInfo) SetInitialValue(v map[string]interface{}) {
	o.InitialValue = &v
}

// GetIsApproverProperty returns the IsApproverProperty field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetIsApproverProperty() bool {
	if o == nil || o.IsApproverProperty == nil {
		var ret bool
		return ret
	}
	return *o.IsApproverProperty
}

// GetIsApproverPropertyOk returns a tuple with the IsApproverProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetIsApproverPropertyOk() (*bool, bool) {
	if o == nil || o.IsApproverProperty == nil {
		return nil, false
	}
	return o.IsApproverProperty, true
}

// HasIsApproverProperty returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasIsApproverProperty() bool {
	if o != nil && o.IsApproverProperty != nil {
		return true
	}

	return false
}

// SetIsApproverProperty gets a reference to the given bool and assigns it to the IsApproverProperty field.
func (o *BTWorkflowPropertyInfo) SetIsApproverProperty(v bool) {
	o.IsApproverProperty = &v
}

// GetIsNotifierProperty returns the IsNotifierProperty field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetIsNotifierProperty() bool {
	if o == nil || o.IsNotifierProperty == nil {
		var ret bool
		return ret
	}
	return *o.IsNotifierProperty
}

// GetIsNotifierPropertyOk returns a tuple with the IsNotifierProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetIsNotifierPropertyOk() (*bool, bool) {
	if o == nil || o.IsNotifierProperty == nil {
		return nil, false
	}
	return o.IsNotifierProperty, true
}

// HasIsNotifierProperty returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasIsNotifierProperty() bool {
	if o != nil && o.IsNotifierProperty != nil {
		return true
	}

	return false
}

// SetIsNotifierProperty gets a reference to the given bool and assigns it to the IsNotifierProperty field.
func (o *BTWorkflowPropertyInfo) SetIsNotifierProperty(v bool) {
	o.IsNotifierProperty = &v
}

// GetMultivalued returns the Multivalued field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetMultivalued() bool {
	if o == nil || o.Multivalued == nil {
		var ret bool
		return ret
	}
	return *o.Multivalued
}

// GetMultivaluedOk returns a tuple with the Multivalued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetMultivaluedOk() (*bool, bool) {
	if o == nil || o.Multivalued == nil {
		return nil, false
	}
	return o.Multivalued, true
}

// HasMultivalued returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasMultivalued() bool {
	if o != nil && o.Multivalued != nil {
		return true
	}

	return false
}

// SetMultivalued gets a reference to the given bool and assigns it to the Multivalued field.
func (o *BTWorkflowPropertyInfo) SetMultivalued(v bool) {
	o.Multivalued = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTWorkflowPropertyInfo) SetName(v string) {
	o.Name = &v
}

// GetObservers returns the Observers field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetObservers() []BTWorkflowableObjectObserver {
	if o == nil || o.Observers == nil {
		var ret []BTWorkflowableObjectObserver
		return ret
	}
	return o.Observers
}

// GetObserversOk returns a tuple with the Observers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetObserversOk() ([]BTWorkflowableObjectObserver, bool) {
	if o == nil || o.Observers == nil {
		return nil, false
	}
	return o.Observers, true
}

// HasObservers returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasObservers() bool {
	if o != nil && o.Observers != nil {
		return true
	}

	return false
}

// SetObservers gets a reference to the given []BTWorkflowableObjectObserver and assigns it to the Observers field.
func (o *BTWorkflowPropertyInfo) SetObservers(v []BTWorkflowableObjectObserver) {
	o.Observers = v
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetPropertyId() string {
	if o == nil || o.PropertyId == nil {
		var ret string
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetPropertyIdOk() (*string, bool) {
	if o == nil || o.PropertyId == nil {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasPropertyId() bool {
	if o != nil && o.PropertyId != nil {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given string and assigns it to the PropertyId field.
func (o *BTWorkflowPropertyInfo) SetPropertyId(v string) {
	o.PropertyId = &v
}

// GetPropertySource returns the PropertySource field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetPropertySource() int32 {
	if o == nil || o.PropertySource == nil {
		var ret int32
		return ret
	}
	return *o.PropertySource
}

// GetPropertySourceOk returns a tuple with the PropertySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetPropertySourceOk() (*int32, bool) {
	if o == nil || o.PropertySource == nil {
		return nil, false
	}
	return o.PropertySource, true
}

// HasPropertySource returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasPropertySource() bool {
	if o != nil && o.PropertySource != nil {
		return true
	}

	return false
}

// SetPropertySource gets a reference to the given int32 and assigns it to the PropertySource field.
func (o *BTWorkflowPropertyInfo) SetPropertySource(v int32) {
	o.PropertySource = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *BTWorkflowPropertyInfo) SetRequired(v bool) {
	o.Required = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *BTWorkflowPropertyInfo) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetTeamsOnly returns the TeamsOnly field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetTeamsOnly() bool {
	if o == nil || o.TeamsOnly == nil {
		var ret bool
		return ret
	}
	return *o.TeamsOnly
}

// GetTeamsOnlyOk returns a tuple with the TeamsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetTeamsOnlyOk() (*bool, bool) {
	if o == nil || o.TeamsOnly == nil {
		return nil, false
	}
	return o.TeamsOnly, true
}

// HasTeamsOnly returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasTeamsOnly() bool {
	if o != nil && o.TeamsOnly != nil {
		return true
	}

	return false
}

// SetTeamsOnly gets a reference to the given bool and assigns it to the TeamsOnly field.
func (o *BTWorkflowPropertyInfo) SetTeamsOnly(v bool) {
	o.TeamsOnly = &v
}

// GetUiHints returns the UiHints field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetUiHints() BTMetadataPropertyUiHintsInfo {
	if o == nil || o.UiHints == nil {
		var ret BTMetadataPropertyUiHintsInfo
		return ret
	}
	return *o.UiHints
}

// GetUiHintsOk returns a tuple with the UiHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetUiHintsOk() (*BTMetadataPropertyUiHintsInfo, bool) {
	if o == nil || o.UiHints == nil {
		return nil, false
	}
	return o.UiHints, true
}

// HasUiHints returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasUiHints() bool {
	if o != nil && o.UiHints != nil {
		return true
	}

	return false
}

// SetUiHints gets a reference to the given BTMetadataPropertyUiHintsInfo and assigns it to the UiHints field.
func (o *BTWorkflowPropertyInfo) SetUiHints(v BTMetadataPropertyUiHintsInfo) {
	o.UiHints = &v
}

// GetUsersOnly returns the UsersOnly field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetUsersOnly() bool {
	if o == nil || o.UsersOnly == nil {
		var ret bool
		return ret
	}
	return *o.UsersOnly
}

// GetUsersOnlyOk returns a tuple with the UsersOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetUsersOnlyOk() (*bool, bool) {
	if o == nil || o.UsersOnly == nil {
		return nil, false
	}
	return o.UsersOnly, true
}

// HasUsersOnly returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasUsersOnly() bool {
	if o != nil && o.UsersOnly != nil {
		return true
	}

	return false
}

// SetUsersOnly gets a reference to the given bool and assigns it to the UsersOnly field.
func (o *BTWorkflowPropertyInfo) SetUsersOnly(v bool) {
	o.UsersOnly = &v
}

// GetValidator returns the Validator field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetValidator() BTMetadataPropertyValidatorInfo {
	if o == nil || o.Validator == nil {
		var ret BTMetadataPropertyValidatorInfo
		return ret
	}
	return *o.Validator
}

// GetValidatorOk returns a tuple with the Validator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetValidatorOk() (*BTMetadataPropertyValidatorInfo, bool) {
	if o == nil || o.Validator == nil {
		return nil, false
	}
	return o.Validator, true
}

// HasValidator returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasValidator() bool {
	if o != nil && o.Validator != nil {
		return true
	}

	return false
}

// SetValidator gets a reference to the given BTMetadataPropertyValidatorInfo and assigns it to the Validator field.
func (o *BTWorkflowPropertyInfo) SetValidator(v BTMetadataPropertyValidatorInfo) {
	o.Validator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *BTWorkflowPropertyInfo) SetValue(v map[string]interface{}) {
	o.Value = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *BTWorkflowPropertyInfo) GetValueType() string {
	if o == nil || o.ValueType == nil {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowPropertyInfo) GetValueTypeOk() (*string, bool) {
	if o == nil || o.ValueType == nil {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *BTWorkflowPropertyInfo) HasValueType() bool {
	if o != nil && o.ValueType != nil {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *BTWorkflowPropertyInfo) SetValueType(v string) {
	o.ValueType = &v
}

func (o BTWorkflowPropertyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ComputedAssemblyProperty != nil {
		toSerialize["computedAssemblyProperty"] = o.ComputedAssemblyProperty
	}
	if o.ComputedProperty != nil {
		toSerialize["computedProperty"] = o.ComputedProperty
	}
	if o.ComputedPropertyError != nil {
		toSerialize["computedPropertyError"] = o.ComputedPropertyError
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.Dirty != nil {
		toSerialize["dirty"] = o.Dirty
	}
	if o.Editable != nil {
		toSerialize["editable"] = o.Editable
	}
	if o.EditableInUi != nil {
		toSerialize["editableInUi"] = o.EditableInUi
	}
	if o.EnumValues != nil {
		toSerialize["enumValues"] = o.EnumValues
	}
	if o.InitialValue != nil {
		toSerialize["initialValue"] = o.InitialValue
	}
	if o.IsApproverProperty != nil {
		toSerialize["isApproverProperty"] = o.IsApproverProperty
	}
	if o.IsNotifierProperty != nil {
		toSerialize["isNotifierProperty"] = o.IsNotifierProperty
	}
	if o.Multivalued != nil {
		toSerialize["multivalued"] = o.Multivalued
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Observers != nil {
		toSerialize["observers"] = o.Observers
	}
	if o.PropertyId != nil {
		toSerialize["propertyId"] = o.PropertyId
	}
	if o.PropertySource != nil {
		toSerialize["propertySource"] = o.PropertySource
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.SchemaId != nil {
		toSerialize["schemaId"] = o.SchemaId
	}
	if o.TeamsOnly != nil {
		toSerialize["teamsOnly"] = o.TeamsOnly
	}
	if o.UiHints != nil {
		toSerialize["uiHints"] = o.UiHints
	}
	if o.UsersOnly != nil {
		toSerialize["usersOnly"] = o.UsersOnly
	}
	if o.Validator != nil {
		toSerialize["validator"] = o.Validator
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ValueType != nil {
		toSerialize["valueType"] = o.ValueType
	}
	return json.Marshal(toSerialize)
}

type NullableBTWorkflowPropertyInfo struct {
	value *BTWorkflowPropertyInfo
	isSet bool
}

func (v NullableBTWorkflowPropertyInfo) Get() *BTWorkflowPropertyInfo {
	return v.value
}

func (v *NullableBTWorkflowPropertyInfo) Set(val *BTWorkflowPropertyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkflowPropertyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkflowPropertyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkflowPropertyInfo(val *BTWorkflowPropertyInfo) *NullableBTWorkflowPropertyInfo {
	return &NullableBTWorkflowPropertyInfo{value: val, isSet: true}
}

func (v NullableBTWorkflowPropertyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkflowPropertyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
