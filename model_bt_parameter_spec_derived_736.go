/*
 * Onshape REST API
 *
 * The Onshape REST API consumed by all clients.
 *
 * API version: 1.111
 * Contact: api-support@onshape.zendesk.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// BtParameterSpecDerived736 struct for BtParameterSpecDerived736
type BtParameterSpecDerived736 struct {
	AdditionalLocalizedStrings int32 `json:"additionalLocalizedStrings,omitempty"`
	BtType string `json:"btType,omitempty"`
	ColumnName string `json:"columnName,omitempty"`
	DefaultValue BtmParameter1 `json:"defaultValue,omitempty"`
	IconUri string `json:"iconUri,omitempty"`
	LocalizableName string `json:"localizableName,omitempty"`
	LocalizedName string `json:"localizedName,omitempty"`
	ParameterId string `json:"parameterId,omitempty"`
	ParameterName string `json:"parameterName,omitempty"`
	StringsToLocalize []string `json:"stringsToLocalize,omitempty"`
	UiHint string `json:"uiHint,omitempty"`
	UiHints []string `json:"uiHints,omitempty"`
	VisibilityCondition BtParameterVisibilityCondition177 `json:"visibilityCondition,omitempty"`
}
