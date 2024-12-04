package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestDrawingAPI(t *testing.T) {
    InitializeTester[*onshape.DrawingApiService](t)

    OpenAPITest{
        Call: onshape.ApiCreateDrawingAppElementRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetDrawingTranslatorFormatsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiCreateDrawingTranslationRequest{},
        Expect: Todo(),
    }.Execute()
    
}
/*** ADDITIONAL TESTS

OpenAPITest{
    Call: onshape.ApiModifyDrawingRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiGetModificationStatusRequest{},
    Expect: Todo(),
}.Execute()


***/