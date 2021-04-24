package api

import (
	"fmt"
	"net/http"
	"time"

	. "github.com/bezahl-online/ptapi/api/gen"
	zvt "github.com/bezahl-online/zvt/command"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=types.cfg.yaml ptapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml ptapi.yaml

// API is the api interface type
type API struct{}

var e *echo.Echo = echo.New()
var Logger *zap.Logger = zvt.Logger

func init() {
	fmt.Println("init")
	server := &API{}
	RegisterHandlers(e, server)
}

// GetTest returns status ok
func (a *API) GetTest(ctx echo.Context) error {
	if err := SendStatus(ctx, http.StatusOK, "OK"); err != nil {
		return err
	}
	return nil
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

func getUNIXUTCFor(locationName string, Y, M, d, h, m, s int) (Timestamp, error) {
	location, err := time.LoadLocation(locationName)
	if err != nil {
		return 0, err
	}
	_, offset := time.Now().In(location).Zone()
	hours := int(offset / 3600)
	minutes := int((offset - 3600*hours) / 60)
	t, err := time.Parse(time.RFC3339,
		fmt.Sprintf("%4d-%02d-%02dT%02d:%02d:%02d+%s",
			Y, M, d, h, m, s,
			fmt.Sprintf("%02d:%02d", hours, minutes)))
	if err != nil {
		return 0, err
	}
	return Timestamp(t.UTC().Unix()), nil
}
