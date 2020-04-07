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
// BtExternalElementReferenceInfo struct for BtExternalElementReferenceInfo
type BtExternalElementReferenceInfo struct {
	DocumentId string `json:"documentId,omitempty"`
	ElementId string `json:"elementId,omitempty"`
	ElementMicroversionId string `json:"elementMicroversionId,omitempty"`
	VersionId string `json:"versionId,omitempty"`
}
