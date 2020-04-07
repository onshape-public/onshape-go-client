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
// BtpExpressionFunction1325AllOf struct for BtpExpressionFunction1325AllOf
type BtpExpressionFunction1325AllOf struct {
	Body BtpStatementBlock271 `json:"body,omitempty"`
	Arguments []BtpArgumentDeclaration232 `json:"arguments,omitempty"`
	SpaceAfterFunction BtpSpace10 `json:"spaceAfterFunction,omitempty"`
	Precondition BtpStatement269 `json:"precondition,omitempty"`
	SpaceAfterArglist BtpSpace10 `json:"spaceAfterArglist,omitempty"`
	SpaceInEmptyList BtpSpace10 `json:"spaceInEmptyList,omitempty"`
	ReturnType BtpTypeName290 `json:"returnType,omitempty"`
	BtType string `json:"btType,omitempty"`
}
