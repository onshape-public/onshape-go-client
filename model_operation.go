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
// Operation struct for Operation
type Operation struct {
	Callbacks map[string]Callback `json:"callbacks,omitempty"`
	Deprecated bool `json:"deprecated,omitempty"`
	Description string `json:"description,omitempty"`
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
	OperationId string `json:"operationId,omitempty"`
	Parameters []Parameter `json:"parameters,omitempty"`
	RequestBody RequestBody `json:"requestBody,omitempty"`
	Responses map[string]ApiResponse `json:"responses,omitempty"`
	Security []SecurityRequirement `json:"security,omitempty"`
	Servers []Server `json:"servers,omitempty"`
	Summary string `json:"summary,omitempty"`
	Tags []string `json:"tags,omitempty"`
}
