package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestExportRuleAPI(t *testing.T) {
    InitializeTester[*onshape.ExportRuleAPIService](t)

    OpenAPITest{
        Call: onshape.ApiGetValidRuleOptionsRequest{},
        Expect: Todo(),
    }.Execute()
    
}