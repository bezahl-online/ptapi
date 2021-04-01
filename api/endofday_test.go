package api

import (
	"fmt"
	"net/http"
	"testing"

	zvt "github.com/bezahl-online/zvt/command"
	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestEndOfDay(t *testing.T) {
	skipShort(t)
	result := testutil.NewRequest().Post("/endofday").WithAcceptJson().Go(t, e)
	assert.Equal(t, http.StatusOK, result.Code())
}

func TestEndOfDayCompletion(t *testing.T) {
	skipShort(t)
	TestEndOfDay(t)
	for {
		result := testutil.NewRequest().Post("/endofday_completion").WithAcceptJson().Go(t, e)
		if assert.Equal(t, http.StatusOK, result.Code()) {
			var response *EndOfDayCompletionResponse = &EndOfDayCompletionResponse{}
			err := result.UnmarshalBodyToObject(&response)
			fmt.Printf("SEND response Status: %02X\n Message:'%s'\nResult: %s\nData:\n%+v\n", response.Status,
				response.Message, (*response.Transaction).Result, (*response.Transaction).Data)
			if response != nil &&
				response.Transaction != nil &&
				response.Transaction.Result == EndOfDayResult_abort ||
				response.Transaction.Result == EndOfDayResult_success {
				break
			}
			if assert.NoError(t, err) {
			}
		}
	}
}

func TestParseEndOfDayResult(t *testing.T) {
	want := EndOfDayCompletionResponse{
		Message: "Test message",
		Status:  1,
		Transaction: &EndOfDayResponse{
			Data: &EndOfDayResponseData{
				Date:         "",
				SingleTotals: SingleTotals{},
				Time:         "",
				Total:        0,
				Tracenr:      0,
			},
			Error:  "",
			Result: "pending",
		},
	}
	result := zvt.EndOfDayResponse{
		TransactionResponse: zvt.TransactionResponse{
			Status:  1,
			Message: "Test message",
		},
		Transaction: &zvt.EoDResult{
			Error:  "",
			Result: "pending",
			Data: &zvt.EoDResultData{
				TraceNr: 0,
				Date:    "",
				Time:    "",
				Total:   0,
				Totals:  zvt.SingleTotals{},
			},
		},
	}
	got := parseEndOfDayResult(result)
	assert.EqualValues(t, want, *got)
}
