package engine

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	engine *Engine
}

func NewHandler(engine *Engine) *Handler {
	return &Handler{engine: engine}
}

func (h *Handler) Advance(w http.ResponseWriter, r *http.Request) {
	result, err := h.engine.Tick()
	if err != nil {
		http.Error(w, "failed to advance", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
