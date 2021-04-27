package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	. "github.com/bezahl-online/ptapi/api/gen"
	zvt "github.com/bezahl-online/zvt/command"

	"github.com/labstack/echo/v4"
)

var authCnt int = 0

// Register initializes the PT with given
// configuration parameters
func (a *API) Register(ctx echo.Context) error {
	var request RegisterJSONRequestBody
	if err := ctx.Bind(&request); err != nil {
		return err
	}
	Logger.Info("request register")

	if err := zvt.PaymentTerminal.Register(); err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(err.Error(), "connect") {
			status = http.StatusServiceUnavailable
		}
		return SendError(ctx, status, fmt.Sprintf("Register returns error: %s", err.Error()))
	}
	authCnt = 0
	return SendStatus(ctx, http.StatusOK, "OK")
}

// RegisterCompletion completes the initialization of
// the PT and responses with the registration information
func (a *API) RegisterCompletion(ctx echo.Context) error {
	Logger.Info("register completion")
	var response *zvt.RegisterResponse = &zvt.RegisterResponse{}
	if err := zvt.PaymentTerminal.Completion(response); err != nil {
		return err
	}
	resp := parseRegisterResult(*response)
	ctx.JSON(http.StatusOK, resp)
	// authCnt++
	// jsonResp, _ := json.Marshal(resp)
	// ioutil.WriteFile(fmt.Sprintf("completion%02d", authCnt), jsonResp, 0644)
	Logger.Info(fmt.Sprintf("register competion %s",
		resp.Transaction.Result))
	return nil
}

func parseRegisterResult(result zvt.RegisterResponse) *RegisterCompletionResponse {
	var response RegisterCompletionResponse = RegisterCompletionResponse{}
	response.Status = int32(result.Status)
	if len(result.Message) > 0 {
		response.Message = result.Message
	}
	switch result.TransactionResponse.Status {
	case zvt.Status_initialisation_necessary:
		Logger.Debug("PT status: diagnosis necessary")
	case zvt.Status_diagnosis_necessary:
		Logger.Debug("PT status: OPT action necessary")
	case zvt.Status_OPT_action_necessary:
		Logger.Debug("PT status: fillingstation mode")
	case zvt.Status_fillingstation_mode:
		Logger.Debug("PT status: vendingmachine mode")
	case zvt.Status_vendingmachine_mode:
	}
	if result.Transaction != nil {
		zvtT := *result.Transaction
		t := RegisterResponse{}
		switch zvtT.Result {
		case zvt.Result_Success:
			t.Result = PtResult_success
		case zvt.Result_Abort:
			t.Result = PtResult_abort
		case zvt.Result_Pending:
			t.Result = PtResult_pending
			if zvtT.Data != nil {
				log.Println("there is data to be parsed")
				// d := *zvtT.Data
				// t.Data = &RegisterResponseData{
				// 	// Aid:        &d.AID, // Gen.Nr.
				// 	// Currency:   int32(d.Currency),
				// 	// Amount:     d.Amount,
				// 	// Card:       Card{Name: d.Card.Name, PanEfId: d.Card.PAN, SequenceNr: int32(d.Card.SeqNr), Type: int32(d.Card.Type)},
				// 	// CardTech:   new(int32),
				// 	// Crypto:     "", // FIXME: find field if possible
				// 	// ReceiptNr:  int64(d.ReceiptNr),
				// 	// TerminalId: d.TID,
				// 	// Timestamp:  d.Date + " " + d.Time,
				// 	// TurnoverNr: int64(d.TurnoverNr),
				// 	// Info:       d.Info,
				// 	// VuNr:       d.VU,
				// }
			}
			response.Transaction = &t
		default:
			t.Error = "no result"
		}
		response.Transaction = &t
	}
	return &response
}
