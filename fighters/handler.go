package fighters

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type fighterHandler struct {
	service FighterService
}

func NewFighterHandler(s FighterService) fighterHandler {
	return fighterHandler{
		service: s,
	}
}
func (fh fighterHandler) CreateFighter(w http.ResponseWriter, req *http.Request) {
	var input CreateFighterInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	f, err := fh.service.CreateFighter(input)
	if err != nil {
		if errors.Is(err, ErrNameTaken) {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(f)
}
func (fh fighterHandler) GetFighter(w http.ResponseWriter, req *http.Request) {
	idStr := req.PathValue("id")
	fighterID, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	f, err := fh.service.GetFighter(fighterID)
	if err != nil {
		if errors.Is(err, ErrFighterNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(f)
}

func (fh fighterHandler) DeleteFighter(w http.ResponseWriter, req *http.Request) {
	idStr := req.PathValue("id")
	fighterID, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = fh.service.DeleteFighter(fighterID)
	if err != nil {
		if errors.Is(err, ErrFighterNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
