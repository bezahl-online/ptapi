package api

import (
	"net/http"
	"testing"

	zvt "github.com/bezahl-online/zvt/command"
	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestAuthorise(t *testing.T) {
	wantForAbort := AuthoriseResponse{
		Result: "abort",
	}
	config := zvt.AuthConfig{Amount: 1}
	result := testutil.NewRequest().Post("/authorise").WithJsonBody(config).WithAcceptJson().Go(t, e)
	if assert.Equal(t, http.StatusOK, result.Code()) {
		var response *AuthoriseResponse = &AuthoriseResponse{}
		err := result.UnmarshalBodyToObject(&response)
		if assert.NoError(t, err) {
			assert.EqualValues(t, wantForAbort, response)
		}
	}
}
