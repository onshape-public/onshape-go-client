package onshape_test

import (
	"encoding/base64"
	"io"
	"os"
	"testing"

	"github.com/onshape-public/go-client/onshape"
	"github.com/onshape-public/go-client/onshape_test/testhelper"
	"github.com/stretchr/testify/require"
)

func TestAppElementAPI(t *testing.T) {
	InitializeTester[*onshape.AppElementApiService](t)

	fileName := "./test_data/hf.txt"
	osFile, err := os.Open(fileName)
	require.NoError(Tester(), err)
	file := onshape.NewHttpFileFromOsFile(osFile)

	SetContext(TestingContext{
		"wvm":                      "w",
		"bTAppElementParams":       GetDefaultAppElementParams(),
		"bTAppElementUpdateParams": GetDefaultAppUpdateParams(),
		"bTAppElementBulkCreateParams": &onshape.BTAppElementBulkCreateParams{
			Names:    []string{"Q.PRT", "R.PRT", "S.PRT", "T.PRT", "U.PRT", "V.PRT", "W.PRT", "X.PRT", "Y.PRT", "Z.PRT"},
			FormatId: "CollabBulkCreate",
		},
		"bid":      "blob1",
		"fileName": fileName,
		"file":     &file,
	}.InheritDefaults(Context()))

	OpenAPITest{
		Call: onshape.ApiCreateDocumentRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		}.BTDocumentParams(onshape.BTDocumentParams{
			Name:        &testhelper.DocumentName,
			Description: &testhelper.DocumentDescription,
			IsPublic:    Ptr(false),
		}),
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiCreateElementRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiBulkCreateElementRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTAppElementBulkCreateInfo) {
			bcp := *Context()["bTAppElementBulkCreateParams"].(*onshape.BTAppElementBulkCreateParams)
			require.EqualValues(Tester(), len(r.ElementIds), len(bcp.Names))
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetSubElementContentRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTAppElementContentInfo) {
			require.Zero(Tester(), len(r.GetData()))
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiUploadBlobSubelementRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiDownloadBlobSubelementWorkspaceRequest{},
		Expect: NoAPIErrorAnd(func(res *onshape.HttpFile) {
			fileBytes, err := io.ReadAll(res.Data)
			require.NoError(Tester(), err)
			buf, err := os.ReadFile(Context()["fileName"].(string))
			require.NoError(Tester(), err)
			require.EqualValues(Tester(), buf, fileBytes)
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetJsonRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTGetJsonResponse2137) {
			require.EqualValues(Tester(), r.GetTree(), Context()["bTAppElementParams"].(*onshape.BTAppElementParams).JsonTree)
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiUpdateAppElementRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetJsonRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTGetJsonResponse2137) {
			nm := make(map[string]interface{})
			for k, v := range Context()["bTAppElementParams"].(*onshape.BTAppElementParams).JsonTree {
				nm[k] = v
			}

			insert := Context()["bTAppElementUpdateParams"].(*onshape.BTAppElementUpdateParams).JsonTreeEdit.GetActualInstance().(*onshape.BTJEditInsert2523)
			nm[insert.Path.GetPath()[0].GetActualInstance().(*onshape.BTJPathKey3221).GetKey()] = insert.GetValue()
			require.EqualValues(t, r.GetTree(), nm)
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiDeleteDocumentRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		},
		Expect: NoAPIError(),
	}.Execute()
}

func TestTransactionAppElementAPI(t *testing.T) {
	InitializeTester[*onshape.AppElementApiService](t)

	SetContext(TestingContext{
		"bTVersionOrWorkspaceParams": &onshape.BTVersionOrWorkspaceParams{
			Name: Ptr("branch"),
		},
		"bTAppElementParams": &onshape.BTAppElementParams{
			FormatId: "String",
			Name:     Ptr("test-element"),
		},
		"wvm": "w",
		"wm":  "w",
	}.InheritDefaults(Context()))

	OpenAPITest{
		Call: onshape.ApiCreateDocumentRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		}.BTDocumentParams(onshape.BTDocumentParams{
			Name:        &testhelper.DocumentName,
			Description: &testhelper.DocumentDescription,
			IsPublic:    Ptr(false),
		}),
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiCreateWorkspaceRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiCreateElementRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTAppElementModifyInfo) {
			Context()["bTAppElementStartTransactionParams"] = &onshape.BTAppElementStartTransactionParams{
				Description:    Ptr("Transaction to populate Assembly w/Data"),
				ParentChangeId: Ptr(r.GetChangeId()),
			}
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiStartTransactionRequest{},
		Expect: NoAPIError(),
	}.Execute()

	for _, eup := range []*onshape.BTAppElementUpdateParams{
		CreateAppElementChangeUpdate("Chapter 1"), CreateAppElementChangeUpdate("Chapter 2"), CreateAppElementChangeUpdate("Chapter 3"),
	} {
		eup.SetTransactionId(Context()["tid"].(string))

		OpenAPITest{
			Call: onshape.ApiUpdateAppElementRequest{}.BTAppElementUpdateParams(*eup),
			Expect: NoAPIErrorAnd(func(r *onshape.BTAppElementModifyInfo) {
				require.Equal(Tester(), Context()["tid"].(string), r.GetTransactionId())
			}),
		}.Execute()
	}

	OpenAPITest{
		Call: onshape.ApiCommitTransactionsRequest{}.BTAppElementCommitTransactionParams(onshape.BTAppElementCommitTransactionParams{
			Description:    Ptr("Done Updating the Assembly"),
			TransactionIds: []string{Context()["tid"].(string)},
		}),
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentHistoryRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		},
		Expect: NoAPIErrorAnd(func(r []onshape.BTDocumentHistoryInfo) {
			require.LessOrEqual(Tester(), 3, len(r))
			Context()["wvmid"] = r[0].GetMicroversionId()
		}),
	}.Execute()

	OpenAPITest{
		Call:    onshape.ApiGetSubelementIdsRequest{},
		Context: TestingContext{"wvm": "m"},
		Expect: NoAPIErrorAnd(func(r *onshape.BTAppElementIdsInfo) {
			require.Equal(Tester(), 3, len(r.GetSubelementIds()))
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiStartTransactionRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiAbortTransactionRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiDeleteDocumentRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		},
		Expect: NoAPIError(),
	}.Execute()
}

func GetDefaultAppElementParams() *onshape.BTAppElementParams {
	return &onshape.BTAppElementParams{
		Name:     Ptr("Test-element-1"),
		FormatId: "String",
		JsonTree: map[string]interface{}{
			"name":    "p.prt",
			"_nodeId": "master",
			"masterProperties": map[string]interface{}{
				"attr1":   "val1",
				"attr2":   "val2",
				"_nodeId": "masterProperties",
			},
			"versionProperties": map[string]interface{}{
				"attr1":   "v11",
				"attr2":   "v22",
				"attr3":   "v33",
				"_nodeId": "versionProperties",
			},
		},
	}
}

func GetDefaultAppUpdateParams() *onshape.BTAppElementUpdateParams {
	bTAppElementUpdateParams := onshape.NewBTAppElementUpdateParams()
	//Top Level Edit Element
	btjEditInsert := onshape.NewBTJEditInsert2523()
	btjEditInsert.SetBtType("BTJEditInsert-2523")
	btjEditInsert.BTJEdit3734 = *onshape.NewBTJEdit3734()
	//Path Element of the Edit
	path := onshape.NewBTJPath3073("master")
	pathType := string("BTJPath-3073")
	path.BtType = &pathType
	// path.SetStartNode("master") -- not needed any more, passing it in the constructor
	//Path.path element
	pathKey := onshape.NewBTJPathKey3221()
	pathKey.SetBtType("BTJPathKey-3221")
	pathKey.SetKey("chapterProperties")
	pathKey.BTJPathElement2297 = *onshape.NewBTJPathElement2297()
	path.SetPath([]onshape.BTJPathElement2297{*pathKey.AsBTJPathElement2297()})
	btjEditInsert.SetPath(*path)
	value := map[string]interface{}{
		"_nodeId": "chapterProperties",
		"chp1":    "v1",
		"chp2":    "v2",
	}
	btjEditInsert.SetValue(value)
	bTAppElementUpdateParams.SetJsonTreeEdit(*btjEditInsert.AsBTJEdit3734())

	return bTAppElementUpdateParams
}

func CreateChapterTestSubelement(name string) onshape.BTAppElementChangeParams {
	return onshape.BTAppElementChangeParams{SubelementId: name, BaseContent: Ptr(base64.StdEncoding.EncodeToString([]byte(name)))}
}

func CreateAppElementChangeUpdate(name string) *onshape.BTAppElementUpdateParams {
	return &onshape.BTAppElementUpdateParams{
		Description: Ptr("Update " + name),
		Changes:     []onshape.BTAppElementChangeParams{CreateChapterTestSubelement(name)},
	}
}

/*** ADDITIONAL TESTS

OpenAPITest{
    Call: onshape.ApiDeleteBlobSubelementRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiGetElementTransactionsRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiDownloadBlobSubelementRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiGetBlobSubelementIdsRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiCompareAppElementJsonRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiGetAppElementHistoryRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiGetJsonPathsRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiDeleteAppElementContentRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiCreateReferenceRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiResolveReferenceRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiUpdateReferenceRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiDeleteReferenceRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiResolveReferencesRequest{},
    Expect: Todo(),
}.Execute()


***/

/*** ADDITIONAL TESTS

OpenAPITest{
    Call: onshape.ApiDeleteAppElementContentBatchRequest{},
    Expect: Todo(),
}.Execute()


***/