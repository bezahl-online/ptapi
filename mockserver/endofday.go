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

// EndOfDay initiates a payment tranaction given
// a specific amount and receipt code
func (a *API) EndOfDay(ctx echo.Context) error {
	var err error
	authCnt = 0
	fmt.Printf("Mock EndOfDay OK\n")
	if err = SendStatus(ctx, http.StatusOK, "OK"); err != nil {
		return err
	}
	return nil
}

// EndOfDayCompletion completes the payment transaction
// and responses with the transaction's data
func (a *API) EndOfDayCompletion(ctx echo.Context) error {
	authCnt++
	fmt.Printf("Mock EndOfDayCompletion OK\n")
	time.Sleep(50 * time.Millisecond)
	resultJson, _ := ioutil.ReadFile(fmt.Sprintf("mockserver/%s/completion%02d", *param.TestDir, authCnt))
	var response *api.AuthCompletionResponse = &api.AuthCompletionResponse{}
	if err := json.Unmarshal(resultJson, response); err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, response)
	return nil
}
