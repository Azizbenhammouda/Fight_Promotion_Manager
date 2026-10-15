package signing

import (
	"errors"
	"net/http"

	"github.com/Azizbenhammouda/Fight_Promotion_Manager/fighters"
	"github.com/Azizbenhammouda/Fight_Promotion_Manager/promotions"
	"github.com/Azizbenhammouda/Fight_Promotion_Manager/users"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Sign(w http.ResponseWriter, req *http.Request) {
	fighterID, err := uuid.Parse(req.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, ok := users.GetUserID(req.Context())
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.service.SignFighter(userID, fighterID)
	if err != nil {
		switch {
		case errors.Is(err, fighters.ErrFighterNotFound),
			errors.Is(err, promotions.ErrPromotionNotFound):
			w.WriteHeader(http.StatusNotFound)
		case errors.Is(err, ErrFighterRetired),
			errors.Is(err, ErrFighterNotFree),
			errors.Is(err, ErrInsufficientBudget):
			w.WriteHeader(http.StatusConflict)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
