package api

import (
	"fmt"
	"net/http"

	zvt "github.com/bezahl-online/zvt/command"
	"github.com/labstack/echo/v4"
)

// Authorise initiates a payment tranaction given
// a specific amount and receipt code
func (a *API) Authorise(ctx echo.Context) error {
	var err error
	var request AuthoriseJSONRequestBody
	authCnt = 0
	fmt.Println("Authorise incomming...")
	err = ctx.Bind(&request)
	if err != nil {
		return err
	}
	if err = zvt.PaymentTerminal.Authorisation(&zvt.AuthConfig{Amount: request.Amount}); err != nil {
		return SendError(ctx, http.StatusNotFound, fmt.Sprintf("EndOfDay returns error: %s", err.Error()))
	}
	return SendStatus(ctx, http.StatusOK, "OK")
}

// AuthoriseCompletion completes the payment transaction
// and responses with the transaction's data
func (a *API) AuthoriseCompletion(ctx echo.Context) error {
	authCnt++
	var request AuthoriseCompletionJSONRequestBody
	fmt.Println("AuthoriseCompletion incomming...")
	err := ctx.Bind(&request)
	if err != nil {
		return err
	}
	PT := zvt.PaymentTerminal
	var response *zvt.AuthorisationResponse = &zvt.AuthorisationResponse{}
	err = PT.Completion(response)
	if err != nil {
		return err
	}
	resp := parseAuthResult(*response)
	ctx.JSON(http.StatusOK, resp)
	// jsonResp, err := json.Marshal(resp)
	// ioutil.WriteFile(fmt.Sprintf("mockserver/authorisation/abort/completion%02d", authCnt), jsonResp, 0644)
	fmt.Printf("AuthoriseCompetion result: %+v\n", *resp.Transaction)
	return nil
}

func parseAuthResult(result zvt.AuthorisationResponse) *AuthCompletionResponse {
	var response AuthCompletionResponse = AuthCompletionResponse{}
	response.Status = int32(result.Status)
	if len(result.Message) > 0 {
		response.Message = result.Message
	}
	if result.Transaction != nil {
		zvtT := *result.Transaction
		t := AuthoriseResponse{}
		switch zvtT.Result {
		case zvt.Result_Success:
			t.Result = AuthoriseResult_success
		case zvt.Result_Abort:
			t.Result = AuthoriseResult_abort
		case zvt.Result_Pending:
			t.Result = AuthoriseResult_pending
			if zvtT.Data != nil {
				d := *zvtT.Data
				t.Data = &AuthoriseResponseData{
					Aid:        &d.AID, // Gen.Nr.
					Amount:     d.Amount,
					Card:       Card{Name: d.Card.Name, PanEfId: d.Card.PAN, SequenceNr: int32(d.Card.SeqNr), Type: int32(d.Card.Type)},
					CardTech:   new(int32),
					Crypto:     "", // FIXME: find field
					ReceiptNr:  int64(d.ReceiptNr),
					TerminalId: d.TID,
					Timestamp:  d.Date + " " + d.Time,
					TurnoverNr: int64(d.TurnoverNr),
					VuNr:       d.VU,
				}
				*t.Data.CardTech = int32(d.Card.Tech)
			}
			response.Transaction = &t
		default:
			t.Error = "no result"
		}
		response.Transaction = &t
	}
	return &response
}