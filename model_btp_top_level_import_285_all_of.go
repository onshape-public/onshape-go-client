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
// BtpTopLevelImport285AllOf struct for BtpTopLevelImport285AllOf
type BtpTopLevelImport285AllOf struct {
	ImportMicroversion string `json:"importMicroversion,omitempty"`
	NamespaceString string `json:"namespaceString,omitempty"`
	CombinedNamespacePathAndVersion string `json:"combinedNamespacePathAndVersion,omitempty"`
	ModuleId BtpModuleId235 `json:"moduleId,omitempty"`
	SpaceBeforeImport BtpSpace10 `json:"spaceBeforeImport,omitempty"`
	Namespace []BtpIdentifier8 `json:"namespace,omitempty"`
	BtType string `json:"btType,omitempty"`
}
