package mockserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	api "github.com/bezahl-online/ptapi/api/gen"
	"github.com/bezahl-online/ptapi/param"
	"github.com/labstack/echo/v4"
)

// Authorise initiates a payment tranaction given
// a specific amount and receipt code
func (a *API) Authorise(ctx echo.Context) error {
	var err error
	var request api.AuthoriseJSONRequestBody
	authCnt = 9
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

// AuthoriseCompletion completes the payment transaction
// and responses with the transaction's data
func (a *API) AuthoriseCompletion(ctx echo.Context) error {
	authCnt++
	var request api.AuthoriseCompletionJSONRequestBody
	err := ctx.Bind(&request)
	if err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)
	resultJson, _ := ioutil.ReadFile(fmt.Sprintf("mockserver/authorisation/%s/completion%02d", *param.TestDir, authCnt))
	var response *api.AuthCompletionResponse = &api.AuthCompletionResponse{}
	err = json.Unmarshal(resultJson, response)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	ctx.JSON(http.StatusOK, response)
	fmt.Printf("AuthoriseCompetion result: %+v\n", *response.Transaction)
	return nil
}
