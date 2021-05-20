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

// Status returns status ok
func (a *API) Status(ctx echo.Context) error {
	authCnt = 0
	if err := SendStatus(ctx, http.StatusOK, "OK"); err != nil {
		return err
	}
	return nil
}

// StatusCompletion completes the status enquiry
func (a *API) StatusCompletion(ctx echo.Context) error {
	authCnt++
	var request api.StatusJSONRequestBody
	err := ctx.Bind(&request)
	if err != nil {
		return err
	}
	time.Sleep(50 * time.Millisecond)
	resultJson, _ := ioutil.ReadFile(fmt.Sprintf("mockserver/status/%s/completion%02d", *param.TestDir, authCnt))
	var response *api.StatusCompletionResponse = &api.StatusCompletionResponse{}
	err = json.Unmarshal(resultJson, response)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	ctx.JSON(http.StatusOK, response)
	fmt.Printf("StatusCompetion result: %+v\n", *response.Transaction)
	return nil
}
