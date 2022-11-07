package presenter

import (
	"net/http"
	"net/url"

	"github.com/go-chi/render"
)

type (
	// ErrorResponse represents the response structure for error resource.
	ErrorResponse struct {
		Code    int        `json:"code"`
		Message string     `json:"message"`
		Details url.Values `json:"details,omitempty"`
	}

	ErrorDetail struct {
		Message string      `json:"message"`
		Param   interface{} `json:"param"`
	}
)

// Render only implements interface contract.
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// ErrNotAcceptable returns an error response for non application-json request.
func ErrNotAcceptable() *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusNotAcceptable,
		Message: "content-type must be application/json",
	}
}

// ErrBadRequest returns an error response for bad request.
func ErrBadRequest(err error) *ErrorResponse {
	msg := err.Error()
	if msg == "EOF" || msg == "" {
		msg = "request body can not be empty"
	}
	return &ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

// Err returns an error response for not found.
func Err(err error, code int) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: err.Error(),
	}
}

func ErrInvalidCredentials() *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "invalid credentials.",
	}
}

func ErrUnauthorized() *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "unauthorized.",
	}
}

// ErrorValidationFailed returns an error response for validation failed.
func ErrorValidationFailed(validationErrors url.Values) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "validation failed.",
		Details: validationErrors,
	}
}

// RenderErrorResponse renders error response.
func RenderErrorResponse(w http.ResponseWriter, r *http.Request, er *ErrorResponse) error {
	render.Status(r, er.Code)
	return render.Render(w, r, er)
}

// RenderResponse render generic response
func RenderResponse(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	render.Status(r, statusCode)
	render.JSON(w, r, data)
}
