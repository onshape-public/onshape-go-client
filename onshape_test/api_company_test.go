package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestCompanyAPI(t *testing.T) {
    InitializeTester[*onshape.CompanyApiService](t)

    OpenAPITest{
        Call: onshape.ApiFindCompanyRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetCompanyRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetDocumentsByNameRequest{},
        Expect: Todo(),
    }.Execute()
    
}
/*** ADDITIONAL TESTS

OpenAPITest{
    Call: onshape.ApiAddUserToCompanyRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiUpdateCompanyUserRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiRemoveUserFromCompanyRequest{},
    Expect: Todo(),
}.Execute()


***/