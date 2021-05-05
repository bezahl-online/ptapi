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

func TestAuthorise(t *testing.T) {
	skipShort(t)
	config := zvt.AuthConfig{Amount: 1}
	result := testutil.NewRequest().Post("/authorise").WithJsonBody(config).WithAcceptJson().Go(t, e)
	if assert.Equal(t, http.StatusOK, result.Code()) {
		var response *AuthoriseResponse = &AuthoriseResponse{}
		err := result.UnmarshalBodyToObject(&response)
		assert.NoError(t, err)
	}
}

func TestAuthoriseCompletion(t *testing.T) {
	skipShort(t)
	TestAuthorise(t)
	if t.Failed() {
		return
	}
	request := AuthoriseCompletionJSONBody{
		ReceiptCode: "12345",
	}
	for {
		result := testutil.NewRequest().Post("/authorise_completion").WithJsonBody(request).WithAcceptJson().Go(t, e)
		if assert.Equal(t, http.StatusOK, result.Code()) {
			var response *AuthCompletionResponse = &AuthCompletionResponse{}
			err := result.UnmarshalBodyToObject(&response)
			fmt.Printf("SEND response Status: %02X\n Message:'%s'\nResult: %s\nData:\n%+v\n", response.Status,
				response.Message, (*response.Transaction).Result, (*response.Transaction).Data)
			if response != nil &&
				response.Transaction != nil &&
				response.Transaction.Result == PtResult_abort ||
				response.Transaction.Result == PtResult_success {
				break
			}
			if assert.NoError(t, err) {
				_ = 0
			}
		} else {
			break
		}
	}
}

func TestParseAuthResult(t *testing.T) {
	want := AuthCompletionResponse{
		Message: "Test message",
		Status:  1,
		Transaction: &AuthoriseResponse{
			Data: &AuthoriseResponseData{
				Aid:    "",
				Amount: 199,
				Card: Card{
					Name:       "Mastercard",
					PanEfId:    "",
					SequenceNr: 444,
					Type:       0x70,
				},
				CardTech:   3,
				Crypto:     "",
				ReceiptNr:  120,
				TerminalId: "29001006",
				Timestamp:  "0311 173315",
				TurnoverNr: 120,
				VuNr:       "100003045",
			},
			Error:  "",
			Result: "pending",
		},
	}
	result := zvt.AuthorisationResponse{
		TransactionResponse: zvt.TransactionResponse{
			Status:  1,
			Message: "Test message",
		},
		Transaction: &zvt.AuthResult{
			Error:  "",
			Result: "pending",
			Data: &zvt.AuthResultData{
				Amount:     199,
				ReceiptNr:  120,
				TurnoverNr: 120,
				TraceNr:    0,
				Date:       "0311",
				Time:       "173315",
				TID:        "29001006",
				VU:         "100003045",
				AID:        "",
				Card: zvt.CardData{
					Name:  "Mastercard",
					Type:  0x70,
					PAN:   "",
					Tech:  3,
					SeqNr: 444,
				},
			},
		},
	}
	got, err := parseAuthResult(result)
	if assert.NoError(t, err) {
		assert.EqualValues(t, want, *got)
	}
}
