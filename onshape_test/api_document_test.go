package onshape_test

import (
	"context"
	"log"
	"slices"
	"testing"

	"github.com/onshape-public/go-client/onshape"
	"github.com/onshape-public/go-client/onshape_test/testhelper"
	"github.com/stretchr/testify/require"
)

const (
	ShareTestGroupId         = "62203d19a8823131729808eb-nonexistant"
	ShareTestGroupType int32 = 2
)

func TestDocumentAPI(t *testing.T) {
	InitializeTester[*onshape.DocumentApiService](t)

	SetContext(TestingContext{
		"label":     &testhelper.DocumentName,
		"docPublic": false,
		"bTDocumentParams": &onshape.BTDocumentParams{
			Name:        &testhelper.DocumentName,
			Description: &testhelper.DocumentDescription,
			IsPublic:    Ptr(false),
		},
		"wm":  "w",
		"wv":  "w",
		"wvm": "w",
		"vm":  "v",
	}.InheritDefaults(Context()))
	client := Context()["client"].(*onshape.APIClient)
	ctx := Context()["ctx"].(context.Context)
	did, wid, _ := testhelper.SetupDocument(ctx, client, "TestDocumentApi")
	targetDocumentName := testhelper.DocumentName+"_target_doc"
	elemInfo, rawResp, err := client.DocumentApi.GetElementsInDocument(ctx, did, "w", wid).Execute()
	if err != nil || (rawResp != nil && rawResp.StatusCode >= 300) {
		log.Fatal("err: ", err, " -- Response status: ", rawResp)
	}
	elementsToMove := []string{elemInfo[0].GetId()}
	OpenAPITest{
		Call: Context()["ApiService"].(*onshape.DocumentApiService).MoveElementsToDocument(ctx,  did, wid),
		Context: TestingContext{"bTMoveElementParams": &onshape.BTMoveElementParams{Name: &targetDocumentName, SourceDocumentId: &did, SourceWorkspaceId: &wid, Elements: elementsToMove}},
		Expect: NoAPIError(),
	}.Execute()


	OpenAPITest{
		Call:   onshape.ApiCreateDocumentRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentsRequest{},
		Expect: NoAPIErrorAnd(func(result *onshape.BTGlobalTreeNodeListResponse) {
			require.NotZero(Tester(), len(result.Items))
			require.GreaterOrEqual(Tester(), slices.IndexFunc(result.Items, func(q onshape.BTGlobalTreeNodeInfo) bool {
				btd := q.GetActualInstance().(*onshape.BTDocumentSummaryInfo)
				return btd.GetId() == Context()["did"] && Ptr(btd.GetCreatedBy()).GetId() == *Context()["creator"].(*string)
			}), 0)
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiUpdateDocumentAttributesRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTDocumentInfo) {
			require.Equal(Tester(), r.GetId(), Context()["did"])
			require.Equal(Tester(), r.GetName(), *Context()["label"].(*string))
			require.Equal(Tester(), r.GetPublic(), Context()["docPublic"])
			require.Equal(Tester(), r.GetDescription(), *Context()["bTDocumentParams"].(*onshape.BTDocumentParams).Description)
			require.True(Tester(), r.HasDefaultWorkspace())
		}),
	}.Execute()

	OpenAPITest{
		Name:    "Test ApiGetDocumentRequest nonexistant",
		Call:    onshape.ApiGetDocumentRequest{},
		Context: TestingContext{"did": "this-doesnt-exit"},
		Expect:  APIError(),
		Pass:    Nothing(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentAclRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTAclInfo) {
			require.Equal(Tester(), r.Owner.GetId(), *Context()["owner"].(*string))
			require.Equal(Tester(), r.GetPublic(), Context()["docPublic"])
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentPermissionSetRequest{},
		Expect: NoAPIErrorAnd(func(r []string) {
			val := []string{"READ", "COPY", "WRITE", "RESHARE", "EXPORT", "DELETE", "COMMENT", "LINK"}
			for _, c := range val {
				require.True(Tester(), slices.Contains(r, c))
			}
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiCreateVersionRequest{}.BTVersionOrWorkspaceParams(onshape.BTVersionOrWorkspaceParams{
			Name:        Ptr("Test version"),
			Description: Ptr("bindings test"),
			DocumentId:  Ptr(Context()["did"].(string)),
			FromHistory: Ptr(true),
		}),
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentVersionsRequest{},
		Expect: NoAPIErrorAnd(func(r []onshape.BTVersionInfo) {
			require.Equal(Tester(), 2, len(r))
			Context()["vmid"] = r[0].GetId()
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiCreateWorkspaceRequest{}.BTVersionOrWorkspaceParams(onshape.BTVersionOrWorkspaceParams{Name: Ptr("branch")}),
		Expect: NoAPIErrorAnd(func(r *onshape.BTWorkspaceInfo) {
			Context()["swid"] = r.Id
		}),
		Pass: Nothing(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentWorkspacesRequest{},
		Expect: NoAPIErrorAnd(func(r []onshape.BTWorkspaceInfo) {
			require.Len(Tester(), r, 2)
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetCurrentMicroversionRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTMicroversionInfo) {
			require.Equal(Tester(), Context()["mv"], r.GetMicroversion())
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiMergeIntoWorkspaceRequest{}.
			BTVersionOrWorkspaceMergeInfo(onshape.BTVersionOrWorkspaceMergeInfo{
				Id:                            Context()["swid"].(*string),
				DefaultMergeStrategy:          onshape.BTMergeStrategyKeep.Ptr(),
				MergeStrategyElementOverrides: Ptr(map[string]onshape.BTMergeStrategy{}),
				Type:                          Ptr("workspace")}),
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetElementsInDocumentRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetCurrentMicroversionRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTMicroversionInfo) {
			// When there are no changes, merging doesn't create a new mv anymore
			require.Equal(Tester(), Context()["mv"], r.GetMicroversion())
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentHistoryRequest{},
		Expect: NoAPIErrorAnd(func(r []onshape.BTDocumentHistoryInfo) {
			// Merge does not create a mv anymore
			require.Len(Tester(), r, 1)
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiRestoreFromHistoryRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentHistoryRequest{},
		Expect: NoAPIErrorAnd(func(r []onshape.BTDocumentHistoryInfo) {
			// Restore to the same mv does not create a new mv anymore
			require.Len(Tester(), r, 1)
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiShareDocumentRequest{}.BTShareParams(onshape.BTShareParams{
			DocumentId:    Ptr(Context()["did"].(string)),
			WorkspaceId:   Ptr(Context()["wid"].(string)),
			PermissionSet: []string{"COPY", "EXPORT", "COMMENT", "WRITE"},
			Entries: []onshape.BTShareEntryParams{
				{
					TeamId:    Ptr(ShareTestGroupId),
					EntryType: Ptr(ShareTestGroupType),
				},
			},
		}),
		Expect: APIError(),
	}.Execute()

	OpenAPITest{
		Call:    onshape.ApiUnShareDocumentRequest{},
		Context: TestingContext{"entryType": Ptr(ShareTestGroupType), "eid": ShareTestGroupId},
		Expect:  APIError(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiShareWithSupportRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiUnshareFromSupportRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call:    onshape.ApiMergePreviewRequest{},
		Context: TestingContext{"sourceType": Ptr("w"), "sourceId": Context()["swid"].(*string)},
		Expect: NoAPIErrorAnd(func(r *onshape.BTMergePreviewInfo) {
			require.NotNil(Tester(), r)
			require.Equalf(Tester(), Context()["mv"], r.GetTargetMicroversionId(), "Target Microversion Id should be the same as the workspace id")
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiDeleteDocumentRequest{},
		Expect: NoAPIError(),
	}.Execute()
}

/*** ADDITIONAL TESTS

OpenAPITest{
    Call: onshape.ApiDownloadExternalDataRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiGetVersionRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiUpdateExternalReferencesToLatestDocumentsRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiMoveElementsToDocumentRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiRevertUnchangedToRevisionsRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiSyncApplicationElementsRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiDeleteWorkspaceRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiGetUnitInfoRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiExport2JsonRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiGetInsertablesRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiSearchRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiCopyWorkspaceRequest{},
    Expect: Todo(),
}.Execute()


***/
