package api

import (
	"net/http"
	"testing"

	api "github.com/bezahl-online/ptapi/api/gen"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDisplayText(t *testing.T) {
	skipShort(t)
	var request api.DisplayTextJSONBody = api.DisplayTextJSONBody{
		Lines: &[]string{"Textzeile1", "Textzeile2", "Textzeile3", "Textzeile4"},
	}
	result := testutil.NewRequest().Post("/display_text").WithJsonBody(request).WithAcceptJson().Go(t, e)
	assert.Equal(t, http.StatusOK, result.Code())
}
