package api

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/bezahl-online/ptapi/api/gen"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	skipShort(t)
	var request RegisterJSONBody = RegisterJSONBody{}
	result := testutil.NewRequest().Post("/register").WithJsonBody(request).WithAcceptJson().Go(t, e)
	assert.Equal(t, http.StatusOK, result.Code())
}

func TestRegisterCompletion(t *testing.T) {
	skipShort(t)
	TestRegister(t)
	for {
		result := testutil.NewRequest().Post("/register_completion").WithAcceptJson().Go(t, e)
		if assert.Equal(t, http.StatusOK, result.Code()) {
			var response *RegisterCompletionResponse = &RegisterCompletionResponse{}
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
