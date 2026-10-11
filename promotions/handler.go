package promotions

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Azizbenhammouda/Fight_Promotion_Manager/users"
)

type promotionHandler struct {
	service PromotionService
}

func NewPromotionHandler(s PromotionService) promotionHandler {
	return promotionHandler{
		service: s,
	}
}
func (ph promotionHandler) CreatePromotion(w http.ResponseWriter, req *http.Request) {
	var input CreatePromotionInput
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, ok := users.GetUserID(req.Context())
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	promo, err := ph.service.CreatePromotion(id, input)
	if err != nil {
		if errors.Is(err, ErrInvalidName) {
			w.WriteHeader(http.StatusBadRequest)
		} else if errors.Is(err, ErrInvalidCountry) {
			w.WriteHeader(http.StatusBadRequest)
		} else if errors.Is(err, ErrNameTaken) {
			w.WriteHeader(http.StatusConflict)
		} else if errors.Is(err, ErrAlreadyHasPromotion) {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(promo)
}
func (ph promotionHandler) GetMyPromotion(w http.ResponseWriter, req *http.Request) {
	id, ok := users.GetUserID(req.Context())
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	promo, err := ph.service.GetMyPromotion(id)
	if err != nil {
		if errors.Is(err, ErrPromotionNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(promo)
}
