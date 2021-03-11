// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

// AuthoriseResponse defines model for AuthoriseResponse.
type AuthoriseResponse struct {
	Aid        *string `json:"aid,omitempty"`
	Amount     *int64  `json:"amount,omitempty"`
	Card       *Card   `json:"card,omitempty"`
	CardTech   *int32  `json:"card_tech,omitempty"`
	Crypto     *string `json:"crypto,omitempty"`
	ReceiptNr  *string `json:"receipt_nr,omitempty"`
	TerminalId *string `json:"terminal_id,omitempty"`
	Timestamp  *string `json:"timestamp,omitempty"`
	VuNr       *string `json:"vuNr,omitempty"`
}

// Card defines model for card.
type Card struct {
	Name       *string `json:"name,omitempty"`
	PanEfId    *string `json:"pan_ef_id,omitempty"`
	SequenceNr *int32  `json:"sequence_nr,omitempty"`
	Type       *string `json:"type,omitempty"`
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
	Amount *int64 `json:"amount,omitempty"`
}

// AuthoriseJSONRequestBody defines body for Authorise for application/json ContentType.
type AuthoriseJSONRequestBody AuthoriseJSONBody
