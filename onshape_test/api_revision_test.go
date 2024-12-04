package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestRevisionAPI(t *testing.T) {
	InitializeTester[*onshape.RevisionApiService](t)

	OpenAPITest{
		Call:   onshape.ApiGetRevisionByPartNumberRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiEnumerateRevisionsRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetRevisionHistoryInCompanyByElementIdRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetRevisionHistoryInCompanyByPartIdRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetRevisionHistoryInCompanyByPartNumberRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiDeleteRevisionHistoryRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetLatestInDocumentOrCompanyRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetAllInDocumentVersionRequest{},
		Expect: Todo(),
	}.Execute()
}

/*** ADDITIONAL TESTS

OpenAPITest{
    Call: onshape.ApiGetAllInDocumentRequest{},
    Expect: Todo(),
}.Execute()


***/