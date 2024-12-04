package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestNumberingSchemeAPI(t *testing.T) {
    InitializeTester[*onshape.NumberingSchemeApiService](t)

    OpenAPITest{
        Call: onshape.ApiNextNumbersRequest{},
        Expect: Todo(),
    }.Execute()
    
}