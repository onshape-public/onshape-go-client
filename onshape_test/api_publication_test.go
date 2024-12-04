package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestPublicationAPI(t *testing.T) {
    InitializeTester[*onshape.PublicationApiService](t)

    OpenAPITest{
        Call: onshape.ApiGetPublicationItemsRequest{},
        Expect: Todo(),
    }.Execute()
    
}
/*** ADDITIONAL TESTS

OpenAPITest{
    Call: onshape.ApiCreatePublicationRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiUpdatePublicationAttributesRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiDeletePublicationRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiAddItemToPublicationRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiDeletePublicationItemRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiAddItemsToPublicationRequest{},
    Expect: Todo(),
}.Execute()


***/