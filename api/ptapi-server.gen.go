// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /abort)
	Abort(ctx echo.Context) error

	// (POST /authorise)
	Authorise(ctx echo.Context) error

	// (POST /authorise_completion)
	AuthoriseCompletion(ctx echo.Context) error

	// (POST /display_text)
	DisplayText(ctx echo.Context) error

	// (POST /endofday)
	EndOfDay(ctx echo.Context) error

	// (POST /endofday_completion)
	EndOfDayCompletion(ctx echo.Context) error

	// (POST /register)
	Register(ctx echo.Context) error

	// (POST /register_completion)
	RegisterCompletion(ctx echo.Context) error

	// (GET /test)
	GetTest(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Abort converts echo context to params.
func (w *ServerInterfaceWrapper) Abort(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Abort(ctx)
	return err
}

// Authorise converts echo context to params.
func (w *ServerInterfaceWrapper) Authorise(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Authorise(ctx)
	return err
}

// AuthoriseCompletion converts echo context to params.
func (w *ServerInterfaceWrapper) AuthoriseCompletion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AuthoriseCompletion(ctx)
	return err
}

// DisplayText converts echo context to params.
func (w *ServerInterfaceWrapper) DisplayText(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DisplayText(ctx)
	return err
}

// EndOfDay converts echo context to params.
func (w *ServerInterfaceWrapper) EndOfDay(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.EndOfDay(ctx)
	return err
}

// EndOfDayCompletion converts echo context to params.
func (w *ServerInterfaceWrapper) EndOfDayCompletion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.EndOfDayCompletion(ctx)
	return err
}

// Register converts echo context to params.
func (w *ServerInterfaceWrapper) Register(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Register(ctx)
	return err
}

// RegisterCompletion converts echo context to params.
func (w *ServerInterfaceWrapper) RegisterCompletion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RegisterCompletion(ctx)
	return err
}

// GetTest converts echo context to params.
func (w *ServerInterfaceWrapper) GetTest(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTest(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/abort", wrapper.Abort)
	router.POST(baseURL+"/authorise", wrapper.Authorise)
	router.POST(baseURL+"/authorise_completion", wrapper.AuthoriseCompletion)
	router.POST(baseURL+"/display_text", wrapper.DisplayText)
	router.POST(baseURL+"/endofday", wrapper.EndOfDay)
	router.POST(baseURL+"/endofday_completion", wrapper.EndOfDayCompletion)
	router.POST(baseURL+"/register", wrapper.Register)
	router.POST(baseURL+"/register_completion", wrapper.RegisterCompletion)
	router.GET(baseURL+"/test", wrapper.GetTest)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
