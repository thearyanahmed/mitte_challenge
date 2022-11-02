package presenter

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")

	j, _ := json.Marshal(body)

	w.WriteHeader(statusCode)
	w.Write(j)
}

func ErrBadRequest(w http.ResponseWriter) {
	msg := map[string]string{"message": "bad request."}
	Response(w, http.StatusBadRequest, msg)
}

func ErrResponse(w http.ResponseWriter, statusCode int, err error) {
	msg := map[string]string{"message": err.Error()}
	Response(w, statusCode, msg)
}
