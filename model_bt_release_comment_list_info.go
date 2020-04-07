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
// BtReleaseCommentListInfo struct for BtReleaseCommentListInfo
type BtReleaseCommentListInfo struct {
	Comments []BtCommentInfo `json:"comments,omitempty"`
	RpId string `json:"rpId,omitempty"`
	RpName string `json:"rpName,omitempty"`
}
