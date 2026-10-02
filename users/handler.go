package users

import (
	"encoding/json"
	"errors"
	"net/http"
)

type userHandler struct {
	service UserService
}

func NewUserHandler(s UserService) userHandler {
	return userHandler{
		service: s,
	}
}

func (h userHandler) Register(w http.ResponseWriter, req *http.Request) {
	var input RegisterInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.service.Register(input)
	if err != nil {
		if errors.Is(err, ErrEmailTaken) {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	//might change this doesnt handle error well
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
func (h userHandler) Login(w http.ResponseWriter, req *http.Request) {
	var input LoginInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	session, err := h.service.Login(input.Email, input.Password)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(session)
}
func (h userHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID, ok := GetUserID(r.Context())
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"user_id": userID.String()})
}
