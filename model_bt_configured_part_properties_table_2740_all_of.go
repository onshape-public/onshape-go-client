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
// BtConfiguredPartPropertiesTable2740AllOf struct for BtConfiguredPartPropertiesTable2740AllOf
type BtConfiguredPartPropertiesTable2740AllOf struct {
	PartDeterministicId string `json:"partDeterministicId,omitempty"`
	PropertyNodeId string `json:"propertyNodeId,omitempty"`
	PartDeterministicIds []string `json:"partDeterministicIds,omitempty"`
	BtType string `json:"btType,omitempty"`
}
