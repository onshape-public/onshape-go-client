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
// Xml struct for Xml
type Xml struct {
	Attribute bool `json:"attribute,omitempty"`
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Prefix string `json:"prefix,omitempty"`
	Wrapped bool `json:"wrapped,omitempty"`
}
