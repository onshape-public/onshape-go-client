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
// BtmParameterMaterial1388 struct for BtmParameterMaterial1388
type BtmParameterMaterial1388 struct {
	BtType string `json:"btType,omitempty"`
	ImportMicroversion string `json:"importMicroversion,omitempty"`
	NodeId string `json:"nodeId,omitempty"`
	ParameterId string `json:"parameterId,omitempty"`
	Material BtPartMaterial1445 `json:"material,omitempty"`
}
