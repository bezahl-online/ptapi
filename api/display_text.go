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

// DisplayText initiates a payment tranaction given
// a specific amount and receipt code
func (a *API) DisplayText(ctx echo.Context) error {
	var request DisplayTextJSONRequestBody
	if err := ctx.Bind(&request); err != nil {
		return err
	}
	Logger.Info("request display text")
	if err := zvt.PaymentTerminal.DisplayText(*request.Lines); err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(err.Error(), "connect") {
			status = http.StatusServiceUnavailable
		}
		return SendError(ctx, status, fmt.Sprintf("DisplayText returns error: %s", err.Error()))
	}
	// authCnt = 0
	return SendStatus(ctx, http.StatusOK, "OK")
}
