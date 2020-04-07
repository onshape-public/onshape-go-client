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
// BtUpdateFeaturesResponse1333AllOf struct for BtUpdateFeaturesResponse1333AllOf
type BtUpdateFeaturesResponse1333AllOf struct {
	Features []BtmFeature134 `json:"features,omitempty"`
	FeatureStates map[string]BtFeatureState1688 `json:"featureStates,omitempty"`
	BtType string `json:"btType,omitempty"`
}
