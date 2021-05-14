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

	// (POST /status)
	Status(ctx echo.Context) error

	// (POST /status_completion)
	StatusCompletion(ctx echo.Context) error

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

// Status converts echo context to params.
func (w *ServerInterfaceWrapper) Status(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Status(ctx)
	return err
}

// StatusCompletion converts echo context to params.
func (w *ServerInterfaceWrapper) StatusCompletion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StatusCompletion(ctx)
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
	router.POST(baseURL+"/status", wrapper.Status)
	router.POST(baseURL+"/status_completion", wrapper.StatusCompletion)
	router.GET(baseURL+"/test", wrapper.GetTest)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"qc3onAwIXv1Y/S8AAP//wqXFqXUlAAA=",
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
