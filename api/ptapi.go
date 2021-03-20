package api

import (
	"fmt"
	"net/http"

	zvt "github.com/bezahl-online/zvt/command"
	"github.com/labstack/echo/v4"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=types.cfg.yaml ptapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml ptapi.yaml

// API is the api interface type
type API struct{}

var e *echo.Echo = echo.New()
var authCnt int

// var scanner device.Scanner

func init() {
	fmt.Println("init")
	server := &API{}
	RegisterHandlers(e, server)
}

// GetTest returns status ok
func (a *API) GetTest(ctx echo.Context) error {
	var err error
	err = SendStatus(ctx, http.StatusOK, "OK")
	if err != nil {
		return err
	}
	return nil
}

// Abort aborts running authorisation process
func (a *API) Abort(ctx echo.Context) error {
	var err error
	var request AbortJSONRequestBody
	fmt.Println("Abort incomming...")
	err = ctx.Bind(&request)
	if err != nil {
		return err
	}
	err = zvt.PaymentTerminal.Abort()
	if err != nil {
		return err
	}
	err = SendStatus(ctx, http.StatusOK, "OK")
	if err != nil {
		return err
	}
	return nil
}

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
	err = zvt.PaymentTerminal.Authorisation(&zvt.AuthConfig{
		Amount: request.Amount,
	})
	if err != nil {
		return err
	}
	err = SendStatus(ctx, http.StatusOK, "OK")
	if err != nil {
		return err
	}
	fmt.Printf("Authorise amount %d for receipt '%s' OK\n", request.Amount, request.ReceiptCode)
	return nil
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
	result, err := PT.Completion()
	if err != nil {
		return err
	}
	response := parseResult(result)
	ctx.JSON(http.StatusOK, response)
	// jsonResp, err := json.Marshal(response)
	// ioutil.WriteFile(fmt.Sprintf("mockserver/authorisation/abort/completion%02d", authCnt), jsonResp, 0644)
	fmt.Printf("AuthoriseCompetion result: %+v\n", *response.Transaction)
	return nil
}

func parseResult(result zvt.CompletionResponse) *CompletionResponse {
	var response CompletionResponse = CompletionResponse{}
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
					Aid:        new(string),
					Amount:     d.Amount,
					Card:       Card{Name: d.Card.Name, PanEfId: d.Card.PAN, SequenceNr: int32(d.Card.SeqNr), Type: int32(d.Card.Type)},
					CardTech:   new(int32),
					Crypto:     "",
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

// SendStatus function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func SendStatus(ctx echo.Context, code int, message string) error {
	statusError := Status{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, statusError)
	return err
}

// SendError function wraps sending of an error in the Error format
// returns sent error message or new error if sending fails
func SendError(ctx echo.Context, code int, message string) error {
	status := Status{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, status)
	if err != nil {
		return err
	}
	return fmt.Errorf(status.Message)
}
