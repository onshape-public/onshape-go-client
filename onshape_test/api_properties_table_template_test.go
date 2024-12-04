package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestPropertiesTableTemplateAPI(t *testing.T) {
    InitializeTester[*onshape.PropertiesTableTemplateApiService](t)

    OpenAPITest{
        Call: onshape.ApiCreateTableTemplateRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetByCompanyIdRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetByDocumentIdRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetTableTemplateRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiDeleteTableTemplateRequest{},
        Expect: Todo(),
    }.Execute()
    
}