package api

import (
	"fmt"
	"net/http"

	. "github.com/bezahl-online/ptapi/api/gen"
	zvt "github.com/bezahl-online/zvt/command"
	"github.com/labstack/echo/v4"
)

// EndOfDay initiates a payment tranaction given
// a specific amount and receipt code
func (a *API) EndOfDay(ctx echo.Context) error {
	Logger.Info("request end of day")
	// authCnt = 0
	if err := zvt.PaymentTerminal.EndOfDay(); err != nil {
		return SendError(ctx, http.StatusNotFound, fmt.Sprintf("EndOfDay returns error: %s", err.Error()))
	}
	return SendStatus(ctx, http.StatusOK, "OK")
}

// EndOfDayCompletion completes the payment transaction
// and responses with the transaction's data
func (a *API) EndOfDayCompletion(ctx echo.Context) error {
	var response *zvt.EndOfDayResponse = &zvt.EndOfDayResponse{}
	if err := zvt.PaymentTerminal.Completion(response); err != nil {
		return err
	}
	resp := parseEndOfDayResult(*response)
	ctx.JSON(http.StatusOK, resp)
	// authCnt++
	// jsonResp, _ := json.Marshal(resp)
	// ioutil.WriteFile(fmt.Sprintf("completion%02d", authCnt), jsonResp, 0644)
	Logger.Info(fmt.Sprintf("end of day competion %s",
		resp.Transaction.Result))
	return nil
}

func parseEndOfDayResult(result zvt.EndOfDayResponse) *EndOfDayCompletionResponse {
	var response EndOfDayCompletionResponse = EndOfDayCompletionResponse{}
	response.Status = int32(result.Status)
	if len(result.Message) > 0 {
		response.Message = result.Message
	}
	if result.Transaction != nil {
		zvtT := *result.Transaction
		t := EndOfDayResponse{}
		switch zvtT.Result {
		case zvt.Result_Success:
			t.Result = EndOfDayResult_success
		case zvt.Result_Abort:
			t.Result = EndOfDayResult_abort
		case zvt.Result_Pending:
			t.Result = EndOfDayResult_pending
			if zvtT.Data != nil {
				t.Data = &EndOfDayResponseData{
					SingleTotals: SingleTotals{
						CountAmex:      new(int64),
						CountDiners:    new(int64),
						CountEC:        new(int64),
						CountEurocard:  new(int64),
						CountJCB:       new(int64),
						CountOther:     new(int64),
						CountVisa:      new(int64),
						ReceiptNrEnd:   new(int64),
						ReceiptNrStart: new(int64),
						TotalAmex:      new(int64),
						TotalDiners:    new(int64),
						TotalEC:        new(int64),
						TotalEurocard:  new(int64),
						TotalJCB:       new(int64),
						TotalOther:     new(int64),
						TotalVisa:      new(int64),
					},
					Date:    zvtT.Data.Date,
					Time:    zvtT.Data.Time,
					Total:   zvtT.Data.Total,
					Tracenr: int64(zvtT.Data.TraceNr),
				}
				if zvtT.Data.Totals != nil {
					totals := t.Data.SingleTotals
					zT := zvtT.Data.Totals
					*totals.ReceiptNrStart = int64(zT.ReceiptNrStart)
					*totals.ReceiptNrEnd = int64(zT.ReceiptNrEnd)
					*totals.CountEC = int64(zT.CountEC)
					*totals.TotalEC = int64(zT.TotalEC)
					*totals.CountJCB = int64(zT.CountJCB)
					*totals.TotalJCB = int64(zT.TotalJCB)
					*totals.CountEurocard = int64(zT.CountEurocard)
					*totals.TotalEurocard = int64(zT.TotalEurocard)
					*totals.CountAmex = int64(zT.CountAmex)
					*totals.TotalAmex = int64(zT.TotalAmex)
					*totals.CountVisa = int64(zT.CountVisa)
					*totals.TotalVisa = int64(zT.TotalVisa)
					*totals.CountDiners = int64(zT.CountDiners)
					*totals.TotalDiners = int64(zT.TotalDiners)
					*totals.CountOther = int64(zT.CountOther)
					*totals.TotalOther = int64(zT.TotalOther)
				}
			}
			response.Transaction = &t
		default:
			t.Error = "no result"
		}
		response.Transaction = &t
	}
	return &response
}
