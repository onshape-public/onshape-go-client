/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.159.11153-01265f451c50
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTLocationInfo226 struct for BTLocationInfo226
type BTLocationInfo226 struct {
	BtType              *string                          `json:"btType,omitempty"`
	Character           *int32                           `json:"character,omitempty"`
	Column              *int32                           `json:"column,omitempty"`
	Document            *string                          `json:"document,omitempty"`
	ElementMicroversion *string                          `json:"elementMicroversion,omitempty"`
	EndCharacter        *int32                           `json:"endCharacter,omitempty"`
	EndColumn           *int32                           `json:"endColumn,omitempty"`
	EndLine             *int32                           `json:"endLine,omitempty"`
	FromNode            *BTPNode7                        `json:"fromNode,omitempty"`
	FromTemplate        *BTLocationInfo226               `json:"fromTemplate,omitempty"`
	LanguageVersion     *int32                           `json:"languageVersion,omitempty"`
	Line                *int32                           `json:"line,omitempty"`
	ModuleIds           *BTDocumentVersionElementIds1897 `json:"moduleIds,omitempty"`
	NodeId              *string                          `json:"nodeId,omitempty"`
	ParseNodeId         *string                          `json:"parseNodeId,omitempty"`
	ParseNodeIdRaw      *BTObjectId                      `json:"parseNodeIdRaw,omitempty"`
	TopLevel            *string                          `json:"topLevel,omitempty"`
	Version             *string                          `json:"version,omitempty"`
}

// NewBTLocationInfo226 instantiates a new BTLocationInfo226 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTLocationInfo226() *BTLocationInfo226 {
	this := BTLocationInfo226{}
	return &this
}

// NewBTLocationInfo226WithDefaults instantiates a new BTLocationInfo226 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTLocationInfo226WithDefaults() *BTLocationInfo226 {
	this := BTLocationInfo226{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTLocationInfo226) SetBtType(v string) {
	o.BtType = &v
}

// GetCharacter returns the Character field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetCharacter() int32 {
	if o == nil || o.Character == nil {
		var ret int32
		return ret
	}
	return *o.Character
}

// GetCharacterOk returns a tuple with the Character field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetCharacterOk() (*int32, bool) {
	if o == nil || o.Character == nil {
		return nil, false
	}
	return o.Character, true
}

// HasCharacter returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasCharacter() bool {
	if o != nil && o.Character != nil {
		return true
	}

	return false
}

// SetCharacter gets a reference to the given int32 and assigns it to the Character field.
func (o *BTLocationInfo226) SetCharacter(v int32) {
	o.Character = &v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetColumn() int32 {
	if o == nil || o.Column == nil {
		var ret int32
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetColumnOk() (*int32, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasColumn() bool {
	if o != nil && o.Column != nil {
		return true
	}

	return false
}

// SetColumn gets a reference to the given int32 and assigns it to the Column field.
func (o *BTLocationInfo226) SetColumn(v int32) {
	o.Column = &v
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetDocument() string {
	if o == nil || o.Document == nil {
		var ret string
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetDocumentOk() (*string, bool) {
	if o == nil || o.Document == nil {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasDocument() bool {
	if o != nil && o.Document != nil {
		return true
	}

	return false
}

// SetDocument gets a reference to the given string and assigns it to the Document field.
func (o *BTLocationInfo226) SetDocument(v string) {
	o.Document = &v
}

// GetElementMicroversion returns the ElementMicroversion field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetElementMicroversion() string {
	if o == nil || o.ElementMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ElementMicroversion
}

// GetElementMicroversionOk returns a tuple with the ElementMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetElementMicroversionOk() (*string, bool) {
	if o == nil || o.ElementMicroversion == nil {
		return nil, false
	}
	return o.ElementMicroversion, true
}

// HasElementMicroversion returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasElementMicroversion() bool {
	if o != nil && o.ElementMicroversion != nil {
		return true
	}

	return false
}

// SetElementMicroversion gets a reference to the given string and assigns it to the ElementMicroversion field.
func (o *BTLocationInfo226) SetElementMicroversion(v string) {
	o.ElementMicroversion = &v
}

// GetEndCharacter returns the EndCharacter field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetEndCharacter() int32 {
	if o == nil || o.EndCharacter == nil {
		var ret int32
		return ret
	}
	return *o.EndCharacter
}

// GetEndCharacterOk returns a tuple with the EndCharacter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetEndCharacterOk() (*int32, bool) {
	if o == nil || o.EndCharacter == nil {
		return nil, false
	}
	return o.EndCharacter, true
}

// HasEndCharacter returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasEndCharacter() bool {
	if o != nil && o.EndCharacter != nil {
		return true
	}

	return false
}

// SetEndCharacter gets a reference to the given int32 and assigns it to the EndCharacter field.
func (o *BTLocationInfo226) SetEndCharacter(v int32) {
	o.EndCharacter = &v
}

// GetEndColumn returns the EndColumn field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetEndColumn() int32 {
	if o == nil || o.EndColumn == nil {
		var ret int32
		return ret
	}
	return *o.EndColumn
}

// GetEndColumnOk returns a tuple with the EndColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetEndColumnOk() (*int32, bool) {
	if o == nil || o.EndColumn == nil {
		return nil, false
	}
	return o.EndColumn, true
}

// HasEndColumn returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasEndColumn() bool {
	if o != nil && o.EndColumn != nil {
		return true
	}

	return false
}

// SetEndColumn gets a reference to the given int32 and assigns it to the EndColumn field.
func (o *BTLocationInfo226) SetEndColumn(v int32) {
	o.EndColumn = &v
}

// GetEndLine returns the EndLine field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetEndLine() int32 {
	if o == nil || o.EndLine == nil {
		var ret int32
		return ret
	}
	return *o.EndLine
}

// GetEndLineOk returns a tuple with the EndLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetEndLineOk() (*int32, bool) {
	if o == nil || o.EndLine == nil {
		return nil, false
	}
	return o.EndLine, true
}

// HasEndLine returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasEndLine() bool {
	if o != nil && o.EndLine != nil {
		return true
	}

	return false
}

// SetEndLine gets a reference to the given int32 and assigns it to the EndLine field.
func (o *BTLocationInfo226) SetEndLine(v int32) {
	o.EndLine = &v
}

// GetFromNode returns the FromNode field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetFromNode() BTPNode7 {
	if o == nil || o.FromNode == nil {
		var ret BTPNode7
		return ret
	}
	return *o.FromNode
}

// GetFromNodeOk returns a tuple with the FromNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetFromNodeOk() (*BTPNode7, bool) {
	if o == nil || o.FromNode == nil {
		return nil, false
	}
	return o.FromNode, true
}

// HasFromNode returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasFromNode() bool {
	if o != nil && o.FromNode != nil {
		return true
	}

	return false
}

// SetFromNode gets a reference to the given BTPNode7 and assigns it to the FromNode field.
func (o *BTLocationInfo226) SetFromNode(v BTPNode7) {
	o.FromNode = &v
}

// GetFromTemplate returns the FromTemplate field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetFromTemplate() BTLocationInfo226 {
	if o == nil || o.FromTemplate == nil {
		var ret BTLocationInfo226
		return ret
	}
	return *o.FromTemplate
}

// GetFromTemplateOk returns a tuple with the FromTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetFromTemplateOk() (*BTLocationInfo226, bool) {
	if o == nil || o.FromTemplate == nil {
		return nil, false
	}
	return o.FromTemplate, true
}

// HasFromTemplate returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasFromTemplate() bool {
	if o != nil && o.FromTemplate != nil {
		return true
	}

	return false
}

// SetFromTemplate gets a reference to the given BTLocationInfo226 and assigns it to the FromTemplate field.
func (o *BTLocationInfo226) SetFromTemplate(v BTLocationInfo226) {
	o.FromTemplate = &v
}

// GetLanguageVersion returns the LanguageVersion field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetLanguageVersion() int32 {
	if o == nil || o.LanguageVersion == nil {
		var ret int32
		return ret
	}
	return *o.LanguageVersion
}

// GetLanguageVersionOk returns a tuple with the LanguageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetLanguageVersionOk() (*int32, bool) {
	if o == nil || o.LanguageVersion == nil {
		return nil, false
	}
	return o.LanguageVersion, true
}

// HasLanguageVersion returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasLanguageVersion() bool {
	if o != nil && o.LanguageVersion != nil {
		return true
	}

	return false
}

// SetLanguageVersion gets a reference to the given int32 and assigns it to the LanguageVersion field.
func (o *BTLocationInfo226) SetLanguageVersion(v int32) {
	o.LanguageVersion = &v
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetLine() int32 {
	if o == nil || o.Line == nil {
		var ret int32
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetLineOk() (*int32, bool) {
	if o == nil || o.Line == nil {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given int32 and assigns it to the Line field.
func (o *BTLocationInfo226) SetLine(v int32) {
	o.Line = &v
}

// GetModuleIds returns the ModuleIds field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetModuleIds() BTDocumentVersionElementIds1897 {
	if o == nil || o.ModuleIds == nil {
		var ret BTDocumentVersionElementIds1897
		return ret
	}
	return *o.ModuleIds
}

// GetModuleIdsOk returns a tuple with the ModuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetModuleIdsOk() (*BTDocumentVersionElementIds1897, bool) {
	if o == nil || o.ModuleIds == nil {
		return nil, false
	}
	return o.ModuleIds, true
}

// HasModuleIds returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasModuleIds() bool {
	if o != nil && o.ModuleIds != nil {
		return true
	}

	return false
}

// SetModuleIds gets a reference to the given BTDocumentVersionElementIds1897 and assigns it to the ModuleIds field.
func (o *BTLocationInfo226) SetModuleIds(v BTDocumentVersionElementIds1897) {
	o.ModuleIds = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTLocationInfo226) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParseNodeId returns the ParseNodeId field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetParseNodeId() string {
	if o == nil || o.ParseNodeId == nil {
		var ret string
		return ret
	}
	return *o.ParseNodeId
}

// GetParseNodeIdOk returns a tuple with the ParseNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetParseNodeIdOk() (*string, bool) {
	if o == nil || o.ParseNodeId == nil {
		return nil, false
	}
	return o.ParseNodeId, true
}

// HasParseNodeId returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasParseNodeId() bool {
	if o != nil && o.ParseNodeId != nil {
		return true
	}

	return false
}

// SetParseNodeId gets a reference to the given string and assigns it to the ParseNodeId field.
func (o *BTLocationInfo226) SetParseNodeId(v string) {
	o.ParseNodeId = &v
}

// GetParseNodeIdRaw returns the ParseNodeIdRaw field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetParseNodeIdRaw() BTObjectId {
	if o == nil || o.ParseNodeIdRaw == nil {
		var ret BTObjectId
		return ret
	}
	return *o.ParseNodeIdRaw
}

// GetParseNodeIdRawOk returns a tuple with the ParseNodeIdRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetParseNodeIdRawOk() (*BTObjectId, bool) {
	if o == nil || o.ParseNodeIdRaw == nil {
		return nil, false
	}
	return o.ParseNodeIdRaw, true
}

// HasParseNodeIdRaw returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasParseNodeIdRaw() bool {
	if o != nil && o.ParseNodeIdRaw != nil {
		return true
	}

	return false
}

// SetParseNodeIdRaw gets a reference to the given BTObjectId and assigns it to the ParseNodeIdRaw field.
func (o *BTLocationInfo226) SetParseNodeIdRaw(v BTObjectId) {
	o.ParseNodeIdRaw = &v
}

// GetTopLevel returns the TopLevel field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetTopLevel() string {
	if o == nil || o.TopLevel == nil {
		var ret string
		return ret
	}
	return *o.TopLevel
}

// GetTopLevelOk returns a tuple with the TopLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetTopLevelOk() (*string, bool) {
	if o == nil || o.TopLevel == nil {
		return nil, false
	}
	return o.TopLevel, true
}

// HasTopLevel returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasTopLevel() bool {
	if o != nil && o.TopLevel != nil {
		return true
	}

	return false
}

// SetTopLevel gets a reference to the given string and assigns it to the TopLevel field.
func (o *BTLocationInfo226) SetTopLevel(v string) {
	o.TopLevel = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTLocationInfo226) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLocationInfo226) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTLocationInfo226) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *BTLocationInfo226) SetVersion(v string) {
	o.Version = &v
}

func (o BTLocationInfo226) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Character != nil {
		toSerialize["character"] = o.Character
	}
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	if o.Document != nil {
		toSerialize["document"] = o.Document
	}
	if o.ElementMicroversion != nil {
		toSerialize["elementMicroversion"] = o.ElementMicroversion
	}
	if o.EndCharacter != nil {
		toSerialize["endCharacter"] = o.EndCharacter
	}
	if o.EndColumn != nil {
		toSerialize["endColumn"] = o.EndColumn
	}
	if o.EndLine != nil {
		toSerialize["endLine"] = o.EndLine
	}
	if o.FromNode != nil {
		toSerialize["fromNode"] = o.FromNode
	}
	if o.FromTemplate != nil {
		toSerialize["fromTemplate"] = o.FromTemplate
	}
	if o.LanguageVersion != nil {
		toSerialize["languageVersion"] = o.LanguageVersion
	}
	if o.Line != nil {
		toSerialize["line"] = o.Line
	}
	if o.ModuleIds != nil {
		toSerialize["moduleIds"] = o.ModuleIds
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParseNodeId != nil {
		toSerialize["parseNodeId"] = o.ParseNodeId
	}
	if o.ParseNodeIdRaw != nil {
		toSerialize["parseNodeIdRaw"] = o.ParseNodeIdRaw
	}
	if o.TopLevel != nil {
		toSerialize["topLevel"] = o.TopLevel
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableBTLocationInfo226 struct {
	value *BTLocationInfo226
	isSet bool
}

func (v NullableBTLocationInfo226) Get() *BTLocationInfo226 {
	return v.value
}

func (v *NullableBTLocationInfo226) Set(val *BTLocationInfo226) {
	v.value = val
	v.isSet = true
}

func (v NullableBTLocationInfo226) IsSet() bool {
	return v.isSet
}

func (v *NullableBTLocationInfo226) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTLocationInfo226(val *BTLocationInfo226) *NullableBTLocationInfo226 {
	return &NullableBTLocationInfo226{value: val, isSet: true}
}

func (v NullableBTLocationInfo226) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTLocationInfo226) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
