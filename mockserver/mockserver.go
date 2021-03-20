package mockserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bezahl-online/ptapi/api"
	"github.com/bezahl-online/ptapi/param"
	"github.com/labstack/echo/v4"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=types.cfg.yaml ptapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml ptapi.yaml

// API is the api interface type
type API struct{}

var e *echo.Echo = echo.New()
var authCnt int

// GetTest returns status ok
func (a *API) GetTest(ctx echo.Context) error {
	var err error
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
	var request api.AuthoriseJSONRequestBody
	authCnt = 0
	err = ctx.Bind(&request)
	if err != nil {
		return err
	}
	fmt.Printf("Mock Authorise amount %d for receipt '%s' OK\n", request.Amount, request.ReceiptCode)
	err = SendStatus(ctx, http.StatusOK, "OK")
	if err != nil {
		return err
	}
	fmt.Printf("Authorise amount %d for receipt '%s' OK\n", request.Amount, request.ReceiptCode)
	return nil
}

// Abort aborts running authorisation process
func (a *API) Abort(ctx echo.Context) error {
	var err error
	var request api.AbortJSONRequestBody
	fmt.Println("Abort incomming...")
	err = ctx.Bind(&request)
	if err != nil {
		return err
	}
	err = SendStatus(ctx, http.StatusOK, "OK")
	if err != nil {
		return err
	}
	return nil
}

// AuthoriseCompletion completes the payment transaction
// and responses with the transaction's data
func (a *API) AuthoriseCompletion(ctx echo.Context) error {
	authCnt++
	var request api.AuthoriseCompletionJSONRequestBody
	err := ctx.Bind(&request)
	if err != nil {
		return err
	}
	fmt.Printf("Mock AuthoriseCompletion for receipt '%s' OK\n", request.ReceiptCode)
	time.Sleep(500 * time.Millisecond)
	resultJson, err := ioutil.ReadFile(fmt.Sprintf("mockserver/%s/completion%02d", *param.TestDir, authCnt))
	var response *api.CompletionResponse = &api.CompletionResponse{}
	err = json.Unmarshal(resultJson, response)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	ctx.JSON(http.StatusOK, response)
	fmt.Printf("AuthoriseCompetion result: %+v\n", *response.Transaction)
	return nil
}

// SendStatus function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func SendStatus(ctx echo.Context, code int, message string) error {
	statusError := api.Status{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, statusError)
	return err
}

// SendError function wraps sending of an error in the Error format
// returns sent error message or new error if sending fails
func SendError(ctx echo.Context, code int, message string) error {
	status := api.Status{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, status)
	if err != nil {
		return err
	}
	return fmt.Errorf(status.Message)
}
