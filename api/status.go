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

// Status initializes the PT with given
// configuration parameters
func (a *API) Status(ctx echo.Context) error {
	var request StatusJSONRequestBody
	if err := ctx.Bind(&request); err != nil {
		return err
	}
	Logger.Info("request Status")

	if err := zvt.PaymentTerminal.Status(); err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(err.Error(), "connect") {
			status = http.StatusServiceUnavailable
		}
		return SendError(ctx, status, fmt.Sprintf("Status returns error: %s", err.Error()))
	}
	authCnt = 0
	return SendStatus(ctx, http.StatusOK, "OK")
}

// StatusCompletion completes the initialization of
// the PT and responses with the registration information
func (a *API) StatusCompletion(ctx echo.Context) error {
	Logger.Info("Status completion")
	var response *zvt.StatusResponse = &zvt.StatusResponse{}
	if err := zvt.PaymentTerminal.Completion(response); err != nil {
		return err
	}
	resp := parseStatusResult(*response)
	ctx.JSON(http.StatusOK, resp)
	// authCnt++
	// jsonResp, _ := json.Marshal(resp)
	// ioutil.WriteFile(fmt.Sprintf("completion%02d", authCnt), jsonResp, 0644)
	Logger.Info(fmt.Sprintf("Status competion %s",
		resp.Transaction.Result))
	return nil
}

func parseStatusResult(result zvt.StatusResponse) *StatusCompletionResponse {
	var response StatusCompletionResponse = StatusCompletionResponse{}
	response.Status = int32(result.Status)
	if len(result.Message) > 0 {
		response.Message = result.Message
	}
	if result.Transaction != nil {
		zvtT := *result.Transaction
		t := StatusEnquiryResponse{}
		switch zvtT.Result {
		case zvt.Result_Success:
			t.Result = PtResult_success
		case zvt.Result_Abort:
			t.Result = PtResult_abort
		case zvt.Result_Pending:
			t.Result = PtResult_pending
			if zvtT.Data != nil {
				log.Println("there is data to be parsed")
			}
			response.Transaction = &t
		default:
			t.Error = "no result"
		}
		response.Transaction = &t
	}
	return &response
}
