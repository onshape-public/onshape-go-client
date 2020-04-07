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
// BtpTopLevelEnumDeclaration284 struct for BtpTopLevelEnumDeclaration284
type BtpTopLevelEnumDeclaration284 struct {
	Atomic bool `json:"atomic,omitempty"`
	BtType string `json:"btType,omitempty"`
	DocumentationType string `json:"documentationType,omitempty"`
	EndSourceLocation int32 `json:"endSourceLocation,omitempty"`
	NodeId string `json:"nodeId,omitempty"`
	ShortDescriptor string `json:"shortDescriptor,omitempty"`
	SpaceAfter BtpSpace10 `json:"spaceAfter,omitempty"`
	SpaceBefore BtpSpace10 `json:"spaceBefore,omitempty"`
	SpaceDefault bool `json:"spaceDefault,omitempty"`
	StartSourceLocation int32 `json:"startSourceLocation,omitempty"`
	Deprecated bool `json:"deprecated,omitempty"`
	SymbolName BtpIdentifier8 `json:"symbolName,omitempty"`
	ArgumentsToDocument []BtpArgumentDeclaration232 `json:"argumentsToDocument,omitempty"`
	DeprecatedExplanation string `json:"deprecatedExplanation,omitempty"`
	ForExport bool `json:"forExport,omitempty"`
	SpaceAfterExport BtpSpace10 `json:"spaceAfterExport,omitempty"`
	Annotation BtpAnnotation231 `json:"annotation,omitempty"`
	SpaceAfterVersion BtpSpace10 `json:"spaceAfterVersion,omitempty"`
	Version BtpLiteralNumber258 `json:"version,omitempty"`
	Name BtpIdentifier8 `json:"name,omitempty"`
	TrailingComma bool `json:"trailingComma,omitempty"`
	Annotations []BtpAnnotation231 `json:"annotations,omitempty"`
	SpaceInEmptyList BtpSpace10 `json:"spaceInEmptyList,omitempty"`
	Values []BtpIdentifier8 `json:"values,omitempty"`
}
