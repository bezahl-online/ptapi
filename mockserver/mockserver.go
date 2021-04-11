package mockserver

import (
	"fmt"
	"net/http"

	api "github.com/bezahl-online/ptapi/api/gen"
	"github.com/labstack/echo/v4"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=types.cfg.yaml ptapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml ptapi.yaml

// API is the api interface type
type API struct{}

var authCnt int

// GetTest returns status ok
func (a *API) GetTest(ctx echo.Context) error {
	if err := SendStatus(ctx, http.StatusOK, "OK"); err != nil {
		return err
	}
	return nil
}

// Abort aborts running authorisation process
func (a *API) Abort(ctx echo.Context) error {
	var err error
	err = SendStatus(ctx, http.StatusOK, "OK")
	if err != nil {
		return err
	}
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
