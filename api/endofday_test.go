package api

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/bezahl-online/ptapi/api/gen"
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
				_ = 0
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
				SingleTotals: SingleTotals{
					CountAmex:      new(int64),
					CountDiners:    new(int64),
					CountEC:        new(int64),
					CountEurocard:  new(int64),
					CountJCB:       new(int64),
					CountOther:     new(int64),
					CountVisa:      new(int64),
					ReceiptNrEnd:   new(int64),
					ReceiptNrStart: new(int64),
					TotalAmex:      new(int64),
					TotalDiners:    new(int64),
					TotalEC:        new(int64),
					TotalEurocard:  new(int64),
					TotalJCB:       new(int64),
					TotalOther:     new(int64),
					TotalVisa:      new(int64),
				},
				Timestamp: "2021-04-01 12:30:15",
				UtcTime:   0,
				Total:     356600,
				Tracenr:   12345678,
			},
			Error:  "",
			Result: "pending",
		},
	}
	st := want.Transaction.Data.SingleTotals
	*st.CountAmex = int64(3)
	*st.CountDiners = int64(5)
	*st.CountEC = int64(7)
	*st.CountEurocard = int64(2)
	*st.CountJCB = int64(1)
	*st.CountOther = int64(6)
	*st.CountVisa = int64(4)
	*st.ReceiptNrEnd = int64(134)
	*st.ReceiptNrStart = int64(123)
	*st.TotalAmex = int64(789)
	*st.TotalDiners = int64(567)
	*st.TotalEC = int64(3500)
	*st.TotalEurocard = int64(456)
	*st.TotalJCB = int64(123)
	*st.TotalOther = int64(890)
	*st.TotalVisa = int64(234)
	result := zvt.EndOfDayResponse{
		TransactionResponse: zvt.TransactionResponse{
			Status:  1,
			Message: "Test message",
		},
		Transaction: &zvt.EoDResult{
			Error:  "",
			Result: "pending",
			Data: &zvt.EoDResultData{
				TraceNr: 12345678,
				Date:    "0401",
				Time:    "123015",
				Total:   356600,
				Totals: &zvt.SingleTotals{
					ReceiptNrStart: 123,
					ReceiptNrEnd:   134,
					CountEC:        7,
					TotalEC:        3500,
					CountJCB:       1,
					TotalJCB:       123,
					CountEurocard:  2,
					TotalEurocard:  456,
					CountAmex:      3,
					TotalAmex:      789,
					CountVisa:      4,
					TotalVisa:      234,
					CountDiners:    5,
					TotalDiners:    567,
					CountOther:     6,
					TotalOther:     890,
				},
			},
		},
	}
	got, err := parseEndOfDayResult(result)
	if assert.NoError(t, err) {
		assert.EqualValues(t, want, *got)
	}
}
