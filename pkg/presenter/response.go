package presenter

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type (
	// ErrorResponse represents the response structure for error resource.
	ErrorResponse struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Details []ErrorDetail `json:"details,omitempty"`
	}

	ErrorDetail struct {
		Path    string      `json:"path"`
		Message string      `json:"message"`
		Param   interface{} `json:"param"`
	}
)

// Render only implements interface contract.
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// ErrNotacceptable returns an error response for non application-json request.
func ErrNotacceptable() *ErrorResponse {
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

// ErrNotFound returns an error response for not found.
func ErrNotFound(err error) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusNotFound,
		Message: err.Error(),
	}
}

// ErrInvalidCredentials
func ErrInvalidCredentials() *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "invalid credentials.",
	}
}

// ErrInvalidCredentials
func ErrUnauthorized() *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "unauthorized.",
	}
}

// ErrorInternal returns an error response for internal server error.
func ErrorInternal(err error) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: err.Error(), //"something went wrong",
	}
}

// ErrorValidationFailed returns an error response for validation failed.
func ErrorValidationFailed(err error) *ErrorResponse {
	details := []ErrorDetail{}
	for _, err := range err.(validator.ValidationErrors) {
		details = append(details, ErrorDetail{
			Path:    err.Field(),
			Message: err.Tag(),
			Param:   err.Value(),
		})
	}
	return &ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "validation failed",
		Details: details,
	}
}

// ErrorConflict returns an error response for unique constraints database error.
func ErrorConflict(err error) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusConflict,
		Message: err.Error(),
	}
}

// RenderErrorResponse renders error response.
func RenderErrorResponse(w http.ResponseWriter, r *http.Request, er *ErrorResponse) error {
	render.Status(r, er.Code)
	return render.Render(w, r, er)
}

func RenderResponse(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	render.Status(r, statusCode)
	render.JSON(w, r, data)
}
