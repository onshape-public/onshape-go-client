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
// BtfsTable953 struct for BtfsTable953
type BtfsTable953 struct {
	AllRowValues [][]string `json:"allRowValues,omitempty"`
	BtType string `json:"btType,omitempty"`
	ColumnCount int32 `json:"columnCount,omitempty"`
	FrozenColumns int32 `json:"frozenColumns,omitempty"`
	NodeId string `json:"nodeId,omitempty"`
	ReadOnly bool `json:"readOnly,omitempty"`
	RowCount int32 `json:"rowCount,omitempty"`
	TableColumns []BtTableColumnInfo1222 `json:"tableColumns,omitempty"`
	TableId string `json:"tableId,omitempty"`
	TableRows []BtTableRow1054 `json:"tableRows,omitempty"`
	Title string `json:"title,omitempty"`
	CrossHighlightData BtTableBaseCrossHighlightData2609 `json:"crossHighlightData,omitempty"`
}
