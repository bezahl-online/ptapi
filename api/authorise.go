package api

import (
	"fmt"
	"net/http"
	"strings"

	. "github.com/bezahl-online/ptapi/api/gen"
	zvt "github.com/bezahl-online/zvt/command"
	"github.com/labstack/echo/v4"
)

// var authCnt int = 0

// Authorise initiates a payment tranaction given
// a specific amount and receipt code
func (a *API) Authorise(ctx echo.Context) error {
	var request AuthoriseJSONRequestBody
	if err := ctx.Bind(&request); err != nil {
		return err
	}
	Logger.Info(fmt.Sprintf("request authorise %d for receipt %s",
		request.Amount, request.ReceiptCode))
	if err := zvt.PaymentTerminal.ReConnect(); err != nil {
		Logger.Error(err.Error())
	}
	if err := zvt.PaymentTerminal.Authorisation(&zvt.AuthConfig{Amount: request.Amount}); err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(err.Error(), "connect") {
			status = http.StatusServiceUnavailable
		}
		if strings.Contains(err.Error(), "84 83") {
			return SendStatus(ctx, http.StatusConflict, "")
		}
		return SendError(ctx, status, fmt.Sprintf("Authorise returns error: %s", err.Error()))
	}
	// authCnt = 0
	return SendStatus(ctx, http.StatusOK, "OK")
}

// AuthoriseCompletion completes the payment transaction
// and responses with the transaction's data
func (a *API) AuthoriseCompletion(ctx echo.Context) error {
	var request AuthoriseCompletionJSONRequestBody
	if err := ctx.Bind(&request); err != nil {
		return err
	}
	Logger.Info(fmt.Sprintf("authorise completion for receipt %s",
		request.ReceiptCode))
	var response *zvt.AuthorisationResponse = &zvt.AuthorisationResponse{}
	// 	TransactionResponse: zvt.TransactionResponse{},
	// 	Transaction: &zvt.AuthResult{
	// 		Error:  "",
	// 		Result: string(PtResult_pending),
	// 	},
	// }
	if err := zvt.PaymentTerminal.Completion(response); err != nil {
		return err
	}
	resp, err := parseAuthResult(*response)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, resp)
	// authCnt++
	// jsonResp, _ := json.Marshal(resp)
	// ioutil.WriteFile(fmt.Sprintf("completion%02d", authCnt), jsonResp, 0644)
	Logger.Info(fmt.Sprintf("authorise competion %s",
		resp.Transaction.Result))
	return nil
}

func parseAuthResult(result zvt.AuthorisationResponse) (*AuthCompletionResponse, error) {
	var response AuthCompletionResponse = AuthCompletionResponse{}
	response.Status = int32(result.Status)
	if len(result.Message) > 0 {
		response.Message = result.Message
	}
	t := AuthoriseResponse{}
	if result.Transaction != nil {
		zvtT := *result.Transaction
		switch zvtT.Result {
		case zvt.Result_Success:
			t.Result = PtResult_success
		case zvt.Result_Abort:
			t.Result = PtResult_abort
		case zvt.Result_Pending:
			response.Transaction = &t
			t.Result = PtResult_pending
			if zvtT.Data == nil || zvtT.Data.Amount == 0 {
				break
			}
			d := *zvtT.Data
			t.Data = &AuthoriseResponseData{
				Aid:         d.AID,
				Currency:    int32(d.Currency),
				Amount:      d.Amount,
				Card:        Card{Name: d.Card.Name, PanEfId: d.Card.PAN, SequenceNr: int32(d.Card.SeqNr), Type: int32(d.Card.Type)},
				CardTech:    int32(d.Card.Tech),
				Crypto:      "",
				EmvCustomer: d.EMVCustomer,
				EmvMerchant: d.EMVMerchant,
				ReceiptNr:   int64(d.ReceiptNr),
				TerminalId:  d.TID,
				Timestamp:   d.Date + " " + d.Time,
				TurnoverNr:  int64(d.TurnoverNr),
				Info:        d.Info,
				VuNr:        d.VU,
			}
		case zvt.Result_Need_EoD:
			t.Result = PtResult_need_end_of_day
		default:
			t.Error = "no result"
		}
	}
	response.Transaction = &t

	return &response, nil
}
