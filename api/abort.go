package api

import (
	"net/http"

	zvt "github.com/bezahl-online/zvt/command"
	"github.com/labstack/echo/v4"
)

// Abort aborts running authorisation process
func (a *API) Abort(ctx echo.Context) error {
	var err error
	Logger.Info("Abort incomming...")
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
