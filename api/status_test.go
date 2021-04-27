package api

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/bezahl-online/ptapi/api/gen"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	skipShort(t)
	var request StatusJSONBody = StatusJSONBody{}
	result := testutil.NewRequest().Post("/status").WithJsonBody(request).WithAcceptJson().Go(t, e)
	assert.Equal(t, http.StatusOK, result.Code())
}

func TestStatusCompletion(t *testing.T) {
	skipShort(t)
	TestStatus(t)
	if !t.Failed() {
		for {
			result := testutil.NewRequest().Post("/status_completion").WithAcceptJson().Go(t, e)
			if assert.Equal(t, http.StatusOK, result.Code()) {
				var response *StatusCompletionResponse = &StatusCompletionResponse{}
				err := result.UnmarshalBodyToObject(&response)
				fmt.Printf("SEND response Status: %02X\n Message:'%s'\nResult: %s\n" /*Data:\n%+v\n",*/, response.Status,
					response.Message, (*response.Transaction).Result) //, (*response.Transaction).Data)
				if response != nil &&
					response.Transaction != nil &&
					response.Transaction.Result == PtResult_abort ||
					response.Transaction.Result == PtResult_success {
					break
				}
				if assert.NoError(t, err) {
					_ = 0
				}
			}
		}
	}
}
