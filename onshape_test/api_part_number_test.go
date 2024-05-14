package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestPartNumberAPI(t *testing.T) {
    InitializeTester[*onshape.NumberingSchemeAPIService](t)

    OpenAPITest{
        Call: onshape.ApiNextNumbersRequest{},
        Expect: Todo(),
    }.Execute()
    
}