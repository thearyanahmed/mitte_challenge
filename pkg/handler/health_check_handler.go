package handler

import (
	"github.com/thearyanahmed/mitte_challenge/pkg/presenter"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type HealthCheckHandler struct {
	db *mongo.Client
}

func NewHealthCheckHandler(db *mongo.Client) *HealthCheckHandler {
	return &HealthCheckHandler{db: db}
}

func (h *HealthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.db.Ping(r.Context(), nil)

	if err != nil {
		_ = presenter.RenderErrorResponse(w, r, presenter.Err(err, http.StatusServiceUnavailable))
		return
	}

	resp := map[string]string{
		"status": "ok",
		"time":   time.Now().String(),
	}

	presenter.RenderResponse(w, r, http.StatusOK, resp)
}
