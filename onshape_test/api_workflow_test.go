package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestWorkflowAPI(t *testing.T) {
    InitializeTester[*onshape.WorkflowApiService](t)

    OpenAPITest{
        Call: onshape.ApiGetActiveWorkflowsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetAllowedApproversRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetAuditLogRequest{},
        Expect: Todo(),
    }.Execute()
    
}
/*** ADDITIONAL TESTS

OpenAPITest{
    Call: onshape.ApiEnumerateObjectWorkflowsRequest{},
    Expect: Todo(),
}.Execute()


***/