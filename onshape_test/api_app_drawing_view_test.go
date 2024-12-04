package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestAppDrawingViewAPI(t *testing.T) {
    InitializeTester[*onshape.AppDrawingViewApiService](t)

    OpenAPITest{
        Call: onshape.ApiGetDrawingViewsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetDrawingViewJsonGeometryRequest{},
        Expect: Todo(),
    }.Execute()
    
}