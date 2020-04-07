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
// BtpStatementLoopForIn279AllOf struct for BtpStatementLoopForIn279AllOf
type BtpStatementLoopForIn279AllOf struct {
	StandardType string `json:"standardType,omitempty"`
	Name BtpIdentifier8 `json:"name,omitempty"`
	TypeName string `json:"typeName,omitempty"`
	Container BtpExpression9 `json:"container,omitempty"`
	Var BtpIdentifier8 `json:"var,omitempty"`
	IsVarDeclaredHere bool `json:"isVarDeclaredHere,omitempty"`
	SpaceBeforeVar BtpSpace10 `json:"spaceBeforeVar,omitempty"`
	BtType string `json:"btType,omitempty"`
}
