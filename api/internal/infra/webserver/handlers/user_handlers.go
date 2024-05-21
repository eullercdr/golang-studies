package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/eullercdr/go/api/internal/dto"
	"github.com/eullercdr/go/api/internal/entity"
	"github.com/eullercdr/go/api/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDb      database.UserInterface
	Jwt         *jwtauth.JWTAuth
	JwtExpiries int
}

func NewUserHandler(userDb database.UserInterface, jwt *jwtauth.JWTAuth, JwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDb:      userDb,
		Jwt:         jwt,
		JwtExpiries: JwtExpiresIn,
	}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := h.UserDb.FindByEmail(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !u.ValidatePassword(user.Password) {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}
	_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiries)).Unix(),
	})
	accessToken := struct {
		AcessToken string `json:"access_token"`
	}{
		AcessToken: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.UserDb.Create(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
