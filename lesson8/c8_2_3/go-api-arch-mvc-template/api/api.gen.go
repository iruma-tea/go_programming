// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for CategoryName.
const (
	Food   CategoryName = "food"
	Music  CategoryName = "music"
	Sports CategoryName = "sports"
)

// AlbumCreateRequest defines model for AlbumCreateRequest.
type AlbumCreateRequest struct {
	Category    Category    `json:"category"`
	ReleaseDate ReleaseDate `json:"releaseDate"`
	Title       string      `json:"title"`
}

// AlbumResponse defines model for AlbumResponse.
type AlbumResponse struct {
	Anniversary Anniversary `json:"anniversary"`
	Category    Category    `json:"category"`
	Id          int         `json:"id"`
	ReleaseDate ReleaseDate `json:"releaseDate"`
	Title       string      `json:"title"`
}

// AlbumUpdateRequest defines model for AlbumUpdateRequest.
type AlbumUpdateRequest struct {
	Category *Category `json:"category,omitempty"`
	Title    *string   `json:"title,omitempty"`
}

// Anniversary defines model for Anniversary.
type Anniversary = int

// Category defines model for Category.
type Category struct {
	Id   *int         `json:"id,omitempty"`
	Name CategoryName `json:"name"`
}

// CategoryName defines model for Category.Name.
type CategoryName string

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Message string `json:"message"`
}

// ReleaseDate defines model for ReleaseDate.
type ReleaseDate = openapi_types.Date

// CreateAlbumJSONRequestBody defines body for CreateAlbum for application/json ContentType.
type CreateAlbumJSONRequestBody = AlbumCreateRequest

// UpdateAlbumByIdJSONRequestBody defines body for UpdateAlbumById for application/json ContentType.
type UpdateAlbumByIdJSONRequestBody = AlbumUpdateRequest

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// CreateAlbumWithBody request with any body
	CreateAlbumWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAlbum(ctx context.Context, body CreateAlbumJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteAlbumById request
	DeleteAlbumById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetAlbumById request
	GetAlbumById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateAlbumByIdWithBody request with any body
	UpdateAlbumByIdWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateAlbumById(ctx context.Context, id int, body UpdateAlbumByIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) CreateAlbumWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateAlbumRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateAlbum(ctx context.Context, body CreateAlbumJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateAlbumRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteAlbumById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteAlbumByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetAlbumById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAlbumByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateAlbumByIdWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateAlbumByIdRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateAlbumById(ctx context.Context, id int, body UpdateAlbumByIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateAlbumByIdRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewCreateAlbumRequest calls the generic CreateAlbum builder with application/json body
func NewCreateAlbumRequest(server string, body CreateAlbumJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateAlbumRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateAlbumRequestWithBody generates requests for CreateAlbum with any type of body
func NewCreateAlbumRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/album")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteAlbumByIdRequest generates requests for DeleteAlbumById
func NewDeleteAlbumByIdRequest(server string, id int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/album/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetAlbumByIdRequest generates requests for GetAlbumById
func NewGetAlbumByIdRequest(server string, id int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/album/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateAlbumByIdRequest calls the generic UpdateAlbumById builder with application/json body
func NewUpdateAlbumByIdRequest(server string, id int, body UpdateAlbumByIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateAlbumByIdRequestWithBody(server, id, "application/json", bodyReader)
}

// NewUpdateAlbumByIdRequestWithBody generates requests for UpdateAlbumById with any type of body
func NewUpdateAlbumByIdRequestWithBody(server string, id int, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/album/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// CreateAlbumWithBodyWithResponse request with any body
	CreateAlbumWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAlbumResponse, error)

	CreateAlbumWithResponse(ctx context.Context, body CreateAlbumJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAlbumResponse, error)

	// DeleteAlbumByIdWithResponse request
	DeleteAlbumByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteAlbumByIdResponse, error)

	// GetAlbumByIdWithResponse request
	GetAlbumByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAlbumByIdResponse, error)

	// UpdateAlbumByIdWithBodyWithResponse request with any body
	UpdateAlbumByIdWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateAlbumByIdResponse, error)

	UpdateAlbumByIdWithResponse(ctx context.Context, id int, body UpdateAlbumByIdJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateAlbumByIdResponse, error)
}

type CreateAlbumResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *AlbumResponse
	JSON400      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r CreateAlbumResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateAlbumResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteAlbumByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *ErrorResponse
	JSON404      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r DeleteAlbumByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteAlbumByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAlbumByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AlbumResponse
	JSON400      *ErrorResponse
	JSON404      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetAlbumByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAlbumByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateAlbumByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AlbumResponse
	JSON400      *ErrorResponse
	JSON404      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r UpdateAlbumByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateAlbumByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// CreateAlbumWithBodyWithResponse request with arbitrary body returning *CreateAlbumResponse
func (c *ClientWithResponses) CreateAlbumWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAlbumResponse, error) {
	rsp, err := c.CreateAlbumWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateAlbumResponse(rsp)
}

func (c *ClientWithResponses) CreateAlbumWithResponse(ctx context.Context, body CreateAlbumJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAlbumResponse, error) {
	rsp, err := c.CreateAlbum(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateAlbumResponse(rsp)
}

// DeleteAlbumByIdWithResponse request returning *DeleteAlbumByIdResponse
func (c *ClientWithResponses) DeleteAlbumByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteAlbumByIdResponse, error) {
	rsp, err := c.DeleteAlbumById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteAlbumByIdResponse(rsp)
}

// GetAlbumByIdWithResponse request returning *GetAlbumByIdResponse
func (c *ClientWithResponses) GetAlbumByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAlbumByIdResponse, error) {
	rsp, err := c.GetAlbumById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAlbumByIdResponse(rsp)
}

// UpdateAlbumByIdWithBodyWithResponse request with arbitrary body returning *UpdateAlbumByIdResponse
func (c *ClientWithResponses) UpdateAlbumByIdWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateAlbumByIdResponse, error) {
	rsp, err := c.UpdateAlbumByIdWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateAlbumByIdResponse(rsp)
}

func (c *ClientWithResponses) UpdateAlbumByIdWithResponse(ctx context.Context, id int, body UpdateAlbumByIdJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateAlbumByIdResponse, error) {
	rsp, err := c.UpdateAlbumById(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateAlbumByIdResponse(rsp)
}

// ParseCreateAlbumResponse parses an HTTP response from a CreateAlbumWithResponse call
func ParseCreateAlbumResponse(rsp *http.Response) (*CreateAlbumResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateAlbumResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest AlbumResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseDeleteAlbumByIdResponse parses an HTTP response from a DeleteAlbumByIdWithResponse call
func ParseDeleteAlbumByIdResponse(rsp *http.Response) (*DeleteAlbumByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteAlbumByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseGetAlbumByIdResponse parses an HTTP response from a GetAlbumByIdWithResponse call
func ParseGetAlbumByIdResponse(rsp *http.Response) (*GetAlbumByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAlbumByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AlbumResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseUpdateAlbumByIdResponse parses an HTTP response from a UpdateAlbumByIdWithResponse call
func ParseUpdateAlbumByIdResponse(rsp *http.Response) (*UpdateAlbumByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateAlbumByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AlbumResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new album
	// (POST /album)
	CreateAlbum(c *gin.Context)
	// Delete a album by ID
	// (DELETE /album/{id})
	DeleteAlbumById(c *gin.Context, id int)
	// Find album by ID
	// (GET /album/{id})
	GetAlbumById(c *gin.Context, id int)
	// Update a album by ID
	// (PATCH /album/{id})
	UpdateAlbumById(c *gin.Context, id int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// CreateAlbum operation middleware
func (siw *ServerInterfaceWrapper) CreateAlbum(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateAlbum(c)
}

// DeleteAlbumById operation middleware
func (siw *ServerInterfaceWrapper) DeleteAlbumById(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteAlbumById(c, id)
}

// GetAlbumById operation middleware
func (siw *ServerInterfaceWrapper) GetAlbumById(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAlbumById(c, id)
}

// UpdateAlbumById operation middleware
func (siw *ServerInterfaceWrapper) UpdateAlbumById(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateAlbumById(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/album", wrapper.CreateAlbum)
	router.DELETE(options.BaseURL+"/album/:id", wrapper.DeleteAlbumById)
	router.GET(options.BaseURL+"/album/:id", wrapper.GetAlbumById)
	router.PATCH(options.BaseURL+"/album/:id", wrapper.UpdateAlbumById)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xWW2vbPhT/Kub8/48mdrrCit7adC1h0I2wPZU+qNZJomJdKh13hODvPiQlrR2nF0aX",
	"vZS8GHOk31Vy1lAZZY1GTR7YGny1RMXj42l926iJQ044w/sGPYW31hmLjiTGmYoTLoxbhef/Hc6BwX/F",
	"047FZrtisp1rc3BYI/d4zglfWzbrjLY5kKQ6rqGVRWDgyUm9gDZuet9IhwLY9WYsfyLXx7zJt+vN7R1W",
	"FHaOWmfordEehzK51vIBneevKz3tjLb5HxkkRUej1IQLdAcxTgrI97nXlf9GL39a8d69eUHEkEM/sKGX",
	"kw6BPrXn3NdcRXTUjQpmzY0JdqnGywpy8NY48h1DnvE4brPPti/OGfd8BRV6zxdvCHE7uA9j1m/Q3DjF",
	"CRiErGDIPHRRz03ETObDpclOv0+zH6hsnRYFm6XRwGA8KkdlgDEWNbcSGHwajUcl5GA5LaOMgod2RHkm",
	"FSOI5CSNngpgkO6bWCFIutDTmRExqMpoQh1XcWtrWcV1xZ0P8Nu769UDOrzW2r6H5BqML1IWkfdROX5f",
	"Bo9JR3CBvnLSUjIysRPByuOyfDfcfsP24J5xkT15koNvlIonaMMo45nGX1mKMAykNIu1FG0AF1hjalY/",
	"1PP4Pso+W01F7IPjCgmdB3a9BhnQQ0dge87SXdTPJO/o3D2f7c0gsONEqSvwymSTjZP/1NuAfXw47CtD",
	"2YVptNhJNeWS8ZRodrvKpueB3AL3nMxLpEMnWB7uyH37+tEIYHAhtdhtg+VULYd9SF/4A1TiL30C+v9Q",
	"3vQJOGAfEzvxUcpHL3avqTiD7mFbt8bVwGBJZFlRlKP4YyflSVlwK4uHMbT5zlBtKl4vjaeXx8ZHn+Nu",
	"4/7YTfs7AAD//2dhTAQ8DQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
