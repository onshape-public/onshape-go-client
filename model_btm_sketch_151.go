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
// BtmSketch151 struct for BtmSketch151
type BtmSketch151 struct {
	BtType string `json:"btType,omitempty"`
	FeatureId string `json:"featureId,omitempty"`
	FeatureType string `json:"featureType,omitempty"`
	ImportMicroversion string `json:"importMicroversion,omitempty"`
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	NodeId string `json:"nodeId,omitempty"`
	Parameters []BtmParameter1 `json:"parameters,omitempty"`
	ReturnAfterSubfeatures bool `json:"returnAfterSubfeatures,omitempty"`
	SubFeatures []BtmFeature134 `json:"subFeatures,omitempty"`
	Suppressed bool `json:"suppressed,omitempty"`
	Entities []BtmSketchGeomEntity5 `json:"entities,omitempty"`
	Constraints []BtmSketchConstraint2 `json:"constraints,omitempty"`
}
