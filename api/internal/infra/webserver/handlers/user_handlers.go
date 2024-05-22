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
	UserDb        database.UserInterface
	Jwt           *jwtauth.JWTAuth
	JwtExpiriesIn int
}

func NewUserHandler(userDb database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDb: userDb,
	}
}

type Error struct {
	Message string `json:"message"`
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiriesIn := r.Context().Value("jwtExpiriesIn").(int)
	var user dto.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
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
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Minute * time.Duration(jwtExpiriesIn)).Unix(),
	})
	accessToken := struct {
		AcessToken string `json:"access_token"`
	}{
		AcessToken: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
}

// CreateUser user godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param request body dto.CreateUserInput true "user request"
// @Success 201
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /users [post]
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
