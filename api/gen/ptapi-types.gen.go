// Package gen provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package gen

// AuthCompletionResponse defines model for auth_completion_response.
type AuthCompletionResponse struct {
	Message     string             `json:"message"`
	Status      int32              `json:"status"`
	Transaction *AuthoriseResponse `json:"transaction,omitempty"`
}

// AuthoriseResponse defines model for authorise_response.
type AuthoriseResponse struct {
	Data   *AuthoriseResponseData `json:"data,omitempty"`
	Error  string                 `json:"error"`
	Result AuthoriseResult        `json:"result"`
}

// AuthoriseResponseData defines model for authorise_response_data.
type AuthoriseResponseData struct {
	Aid      string `json:"aid"`
	Amount   int64  `json:"amount"`
	Card     Card   `json:"card"`
	CardTech int32  `json:"card_tech"`

	Crypto   string `json:"crypto"`
	Currency int32  `json:"currency"`
	// EMV-print-data (merchant-receipt)
	EmvCustomer string `json:"emv_customer"`
	EmvMerchant string `json:"emv_merchant"`
	Info        string `json:"info"`
	ReceiptNr   int64  `json:"receipt_nr"`
	TerminalId  string `json:"terminal_id"`
	Timestamp   string `json:"timestamp"`
	TurnoverNr  int64  `json:"turnover_nr"`
	VuNr        string `json:"vu_nr"`
}

// AuthoriseResult defines model for authorise_result.
type AuthoriseResult string

// List of AuthoriseResult
const (
	AuthoriseResult_abort   AuthoriseResult = "abort"
	AuthoriseResult_pending AuthoriseResult = "pending"
	AuthoriseResult_success AuthoriseResult = "success"
	AuthoriseResult_timeout AuthoriseResult = "timeout"
)

// Card defines model for card.
type Card struct {

	// YYMM
	Expiry     string `json:"expiry"`
	Name       string `json:"name"`
	PanEfId    string `json:"pan_ef_id"`
	SequenceNr int32  `json:"sequence_nr"`
	Type       int32  `json:"type"`
}

// EndOfDayCompletionResponse defines model for end_of_day_completion_response.
type EndOfDayCompletionResponse struct {
	Message     string            `json:"message"`
	Status      int32             `json:"status"`
	Transaction *EndOfDayResponse `json:"transaction,omitempty"`
}

// EndOfDayResponse defines model for end_of_day_response.
type EndOfDayResponse struct {
	Data   *EndOfDayResponseData `json:"data,omitempty"`
	Error  string                `json:"error"`
	Result EndOfDayResult        `json:"result"`
}

// EndOfDayResponseData defines model for end_of_day_response_data.
type EndOfDayResponseData struct {
	SingleTotals SingleTotals `json:"single_totals"`

	// YYYY-MM-DD hh:mm:ss
	Timestamp string `json:"timestamp"`
	Total     int64  `json:"total"`
	Tracenr   int64  `json:"tracenr"`

	// unix utc timestamp
	UtcTime int64 `json:"utc_time"`
}

// EndOfDayResult defines model for end_of_day_result.
type EndOfDayResult string

// List of EndOfDayResult
const (
	EndOfDayResult_abort   EndOfDayResult = "abort"
	EndOfDayResult_pending EndOfDayResult = "pending"
	EndOfDayResult_success EndOfDayResult = "success"
	EndOfDayResult_timeout EndOfDayResult = "timeout"
)

// RegisterCompletionResponse defines model for register_completion_response.
type RegisterCompletionResponse struct {
	Message     string            `json:"message"`
	Status      int32             `json:"status"`
	Transaction *RegisterResponse `json:"transaction,omitempty"`
}

// RegisterResponse defines model for register_response.
type RegisterResponse struct {
	Error  string         `json:"error"`
	Result RegisterResult `json:"result"`
}

// RegisterResult defines model for register_result.
type RegisterResult string

// List of RegisterResult
const (
	RegisterResult_abort   RegisterResult = "abort"
	RegisterResult_pending RegisterResult = "pending"
	RegisterResult_success RegisterResult = "success"
	RegisterResult_timeout RegisterResult = "timeout"
)

// SingleTotals defines model for single_totals.
type SingleTotals struct {
	CountAmex      *int64 `json:"CountAmex,omitempty"`
	CountDiners    *int64 `json:"CountDiners,omitempty"`
	CountEC        *int64 `json:"CountEC,omitempty"`
	CountEurocard  *int64 `json:"CountEurocard,omitempty"`
	CountJCB       *int64 `json:"CountJCB,omitempty"`
	CountOther     *int64 `json:"CountOther,omitempty"`
	CountVisa      *int64 `json:"CountVisa,omitempty"`
	ReceiptNrEnd   *int64 `json:"ReceiptNrEnd,omitempty"`
	ReceiptNrStart *int64 `json:"ReceiptNrStart,omitempty"`
	TotalAmex      *int64 `json:"TotalAmex,omitempty"`
	TotalDiners    *int64 `json:"TotalDiners,omitempty"`
	TotalEC        *int64 `json:"TotalEC,omitempty"`
	TotalEurocard  *int64 `json:"TotalEurocard,omitempty"`
	TotalJCB       *int64 `json:"TotalJCB,omitempty"`
	TotalOther     *int64 `json:"TotalOther,omitempty"`
	TotalVisa      *int64 `json:"TotalVisa,omitempty"`
}

// Status defines model for status.
type Status struct {

	// Status code
	Code int32 `json:"code"`

	// Status message
	Message string `json:"message"`
}

// Statusresponse defines model for statusresponse.
type Statusresponse Status

// AuthoriseJSONBody defines parameters for Authorise.
type AuthoriseJSONBody struct {

	// amount in cent
	Amount      int64  `json:"amount"`
	ReceiptCode string `json:"receipt_code"`
}

// AuthoriseCompletionJSONBody defines parameters for AuthoriseCompletion.
type AuthoriseCompletionJSONBody struct {
	ReceiptCode string `json:"receipt_code"`
}

// DisplayTextJSONBody defines parameters for DisplayText.
type DisplayTextJSONBody struct {
	Lines *[]string `json:"lines,omitempty"`
}

// RegisterJSONBody defines parameters for Register.
type RegisterJSONBody struct {
	Option *string `json:"option,omitempty"`
}

// AuthoriseJSONRequestBody defines body for Authorise for application/json ContentType.
type AuthoriseJSONRequestBody AuthoriseJSONBody

// AuthoriseCompletionJSONRequestBody defines body for AuthoriseCompletion for application/json ContentType.
type AuthoriseCompletionJSONRequestBody AuthoriseCompletionJSONBody

// DisplayTextJSONRequestBody defines body for DisplayText for application/json ContentType.
type DisplayTextJSONRequestBody DisplayTextJSONBody

// RegisterJSONRequestBody defines body for Register for application/json ContentType.
type RegisterJSONRequestBody RegisterJSONBody
