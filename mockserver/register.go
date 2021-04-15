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

// GetTest returns status ok
func (a *API) Register(ctx echo.Context) error {
	if err := SendStatus(ctx, http.StatusOK, "OK"); err != nil {
		return err
	}
	return nil
}

// RegisterCompletion completes the payment transaction
// and responses with the transaction's data
func (a *API) RegisterCompletion(ctx echo.Context) error {
	authCnt++
	var request api.RegisterJSONRequestBody
	err := ctx.Bind(&request)
	if err != nil {
		return err
	}
	time.Sleep(50 * time.Millisecond)
	resultJson, _ := ioutil.ReadFile(fmt.Sprintf("mockserver/register/%s/completion%02d", *param.TestDir, authCnt))
	var response *api.RegisterCompletionResponse = &api.RegisterCompletionResponse{}
	err = json.Unmarshal(resultJson, response)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	ctx.JSON(http.StatusOK, response)
	fmt.Printf("RegisterCompetion result: %+v\n", *response.Transaction)
	return nil
}
