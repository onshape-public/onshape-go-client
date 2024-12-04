package onshape_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/onshape-public/go-client/onshape"
	"github.com/onshape-public/go-client/onshape_test/testhelper"
	"github.com/stretchr/testify/require"
)

func TestCommentAPI(t *testing.T) {
	InitializeTester[*onshape.CommentApiService](t)
	message := "test comment"
	updatedMessage := "updated test comment"
	fileName := "spider-man-1.jpg"
	fileSize := 9084
	osFile, err := os.Open("./test_data/" + fileName)
	require.NoError(Tester(), err)
	file := onshape.NewHttpFileFromOsFile(osFile)

	onshapeAPIClient := Context()["client"].(*onshape.APIClient)
	commentApi := onshapeAPIClient.CommentApi

	OpenAPITest{
		Call: onshape.ApiCreateDocumentRequest{
			ApiService: onshapeAPIClient.DocumentApi,
		}.BTDocumentParams(onshape.BTDocumentParams{
			Name:        &testhelper.DocumentName,
			Description: &testhelper.DocumentDescription,
			IsPublic:    Ptr(false),
		}),
		Expect: NoAPIError(),
	}.Execute()

	did := Context()["did"].(string)
	wid := Context()["wid"].(string)

	OpenAPITest{
		Call: onshape.ApiCreateCommentRequest{
			ApiService: commentApi,
		}.BTCommentParams(onshape.BTCommentParams{
			DocumentId:  Ptr(did),
			ObjectId:    Ptr(did),
			WorkspaceId: Ptr(wid),
			Message:     Ptr(message),
		}),
		Expect: NoAPIErrorAnd(func(r *onshape.BTCommentInfo) {
			require.NotNil(Tester(), r)
			require.NotEmpty(Tester(), r.GetId())
			require.Equal(Tester(), message, r.GetMessage())
		}),
	}.Execute()

	cid := Context()["cid"].(string)
	OpenAPITest{
		Call: onshape.ApiGetCommentsRequest{ApiService: commentApi}.Did(did),
		Expect: NoAPIErrorAnd(func(r *onshape.BTListResponseBTCommentInfo) {
			require.Equal(Tester(), 1, len(r.GetItems()))
		}),
	}.Execute()

	ctx := Context()["ctx"].(context.Context)
	OpenAPITest{
		Call: commentApi.GetComment(ctx, cid),
		Expect: NoAPIErrorAnd(func(r *onshape.BTCommentInfo) {
			require.NotNil(Tester(), r)
			require.NotEmpty(Tester(), r.GetId())
			require.Equal(Tester(), message, r.GetMessage())
		}),
	}.Execute()

	OpenAPITest{
		Call: commentApi.UpdateComment(ctx, cid).BTCommentParams(onshape.BTCommentParams{
			Id:      Ptr(cid),
			Message: Ptr(updatedMessage),
		}),
		Expect: NoAPIErrorAnd(func(r *onshape.BTCommentInfo) {
			require.NotNil(Tester(), r)
			require.NotEmpty(Tester(), r.GetId())
			require.Equal(Tester(), updatedMessage, r.GetMessage())
		}),
	}.Execute()

	OpenAPITest{
		Call: commentApi.Resolve(ctx, cid),
		Expect: NoAPIErrorAnd(func(r *onshape.BTCommentInfo) {
			require.NotNil(Tester(), r)
			require.NotEmpty(Tester(), r.GetId())
			require.True(Tester(), r.HasResolvedBy())
			require.True(Tester(), r.HasResolvedAt())
			require.False(Tester(), r.HasReopenedBy())
			require.False(Tester(), r.HasReopenedAt())
		}),
	}.Execute()

	OpenAPITest{
		Call: commentApi.Reopen(ctx, cid),
		Expect: NoAPIErrorAnd(func(r *onshape.BTCommentInfo) {
			require.NotNil(Tester(), r)
			require.NotEmpty(Tester(), r.GetId())
			require.True(Tester(), r.HasReopenedBy())
			require.True(Tester(), r.HasReopenedAt())
			require.True(Tester(), r.GetReopenedAt().After(r.GetResolvedAt().Time))
		}),
	}.Execute()

	OpenAPITest{
		Call: commentApi.AddAttachment(ctx, cid).FileContentLength(int32(fileSize)).File(file),
		Expect: NoAPIErrorAnd(func(r *onshape.BTCommentInfo) {
			require.NotNil(Tester(), r)
			require.NotEmpty(Tester(), r.GetId())
			require.NotNil(Tester(), r.GetAttachment())
			require.NotEmpty(Tester(), r.GetAttachment().Id)
		}),
	}.Execute()

	fdid := Context()["fdid"].(string)
	OpenAPITest{
		Call: commentApi.GetAttachment(ctx, cid, fdid, "jpg"),
		Expect: NoAPIErrorAnd(func(f *onshape.HttpFile) {
			require.NotNil(Tester(), f)
			require.NotNil(Tester(), f.Data)
		}),
	}.Execute()

	OpenAPITest{
		Call:   commentApi.DeleteAttachments(ctx, cid),
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call:   commentApi.DeleteComment(ctx, cid),
		Expect: NoAPIError(),
	}.Execute()
}

// This test creates
// a thread of comments in the order: A, B i.e. comment B is a reply to comment A
// A: references element e1
// B: references element e2
func TestCommentThreadReferenceModels(t *testing.T) {
	documentName := "TestCommentThreadReferenceModels-" + uuid.New().String()
	did, wid, teardownDoc := testhelper.SetupDocument(ctx, client, documentName)

	defer teardownDoc()

	type args struct {
		elementNames []string
	}
	
	elements := args{[]string{"e1.PRT", "e2.PRT"}}
	
	bulkAppElementParams := onshape.NewBTAppElementBulkCreateParams("CollabBulkCreate")
	bulkAppElementParams.SetNames(elements.elementNames)
	bulkAppElementParams.SetDescription("Multiple elements in one thread of comments")

	appElementBulkCreateInfo, rawResp, err :=
		client.AppElementApi.BulkCreateElement(ctx, did, wid).BTAppElementBulkCreateParams(*bulkAppElementParams).Execute()
	require.NoError(t, err, "Error Creating Bulk App Element")
	require.NotNil(t, rawResp, "Response should not be nil")
	require.True(t, rawResp.StatusCode < 300, "Status code should be less than 300")
	require.NotNil(t, appElementBulkCreateInfo, "AppElementBulkCreateInfo should not be nil")
	require.True(t, len(appElementBulkCreateInfo.GetElementIds()) == len(elements.elementNames))

	InitializeTester[*onshape.CommentApiService](t)
	e1id := "e1.PRT"
	e2id := "e2.PRT"
	onshapeAPIClient := Context()["client"].(*onshape.APIClient)
	commentApi := onshapeAPIClient.CommentApi

	//Step 1: create A
	var commentA string
	OpenAPITest{
		Call: onshape.ApiCreateCommentRequest{
			ApiService: commentApi,
		}.BTCommentParams(onshape.BTCommentParams{
			DocumentId:  Ptr(did),
			ObjectId:    Ptr(did),
			WorkspaceId: Ptr(wid),
			ElementId: Ptr(e1id),
			Message:     Ptr("Comment A"),
		}),
		Expect: NoAPIErrorAnd(func(r *onshape.BTCommentInfo) {
			
			require.NotNil(Tester(), r)
			require.NotEmpty(Tester(), r.GetId())
			commentA = r.GetId()
			require.Equal(Tester(), "Comment A", r.GetMessage())
		}),
	}.Execute()
	//Step 2: create B
	OpenAPITest{
		Call: onshape.ApiCreateCommentRequest{
			ApiService: commentApi,
		}.BTCommentParams(onshape.BTCommentParams{
			DocumentId:  Ptr(did),
			ObjectId:    Ptr(did),
			WorkspaceId: Ptr(wid),
			ElementId: Ptr(e2id),
			Message:     Ptr("Comment B as a reply to comment A"),
			ParentId: 	Ptr(commentA),
		}),
		Expect: NoAPIErrorAnd(func(r *onshape.BTCommentInfo) {
			require.NotNil(Tester(), r)
			require.NotEmpty(Tester(), r.GetId())
			require.NotEmpty(Tester(),r.GetElementId())
			require.Equal(Tester(), e2id, r.GetElementId())
			require.Equal(Tester(), "Comment B as a reply to comment A", r.GetMessage())
		}),
	}.Execute()
}