package api

import (
	"fmt"
	"net/http"

	zvt "github.com/bezahl-online/zvt/command"
	"github.com/labstack/echo/v4"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=types.cfg.yaml ptapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml ptapi.yaml

// API is the api interface type
type API struct{}

var e *echo.Echo = echo.New()

// var scanner device.Scanner

func init() {
	fmt.Println("init")
	testServer := &API{}
	RegisterHandlers(e, testServer)
	// go device.Connect(&scanner)
}

// GetTest returns status ok
func (a *API) GetTest(ctx echo.Context) error {
	var err error
	err = SendStatus(ctx, http.StatusOK, "OK")
	if err != nil {
		return err
	}
	return nil
}

// Authorise returns status ok
func (a *API) Authorise(ctx echo.Context) error {
	var request AuthoriseJSONRequestBody
	err := ctx.Bind(&request)
	if err != nil {
		return err
	}
	PT := zvt.PaymentTerminal
	config := zvt.AuthConfig{
		Amount: int(*request.Amount),
	}
	// result := zvt.AuthResult{
	// 	Success: true,
	// 	Card: zvt.CardData{
	// 		Name: "RalphCard",
	// 		Type: "Nimmmit",
	// 		PAN:  "1293871293578",
	// 		Tech: 3,
	// 	},
	// 	ReceiptNr: "23412",
	// 	TID:       "239487",
	// }
	result, err := PT.Authorisation(&config)
	if err != nil {
		return err
	}
	var response AuthoriseResponse = AuthoriseResponse{
		Card: &Card{
			Name:    &result.Card.Name,
			PanEfId: &result.Card.PAN,
			Type:    &result.Card.Type,
		},
		Amount:     request.Amount,
		Aid:        new(string),
		CardTech:   new(int32),
		Crypto:     new(string),
		ReceiptNr:  &result.ReceiptNr,
		TerminalId: &result.TID,
	}
	*response.CardTech = int32(result.Card.Tech)
	*response.Aid = "1000fake00000333"
	err = ctx.JSON(http.StatusOK, response)
	if err != nil {
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
