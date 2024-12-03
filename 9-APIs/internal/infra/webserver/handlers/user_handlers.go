package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/eduardo1520/goexpert/9-APIs/internal/dto"
	"github.com/eduardo1520/goexpert/9-APIs/internal/entity"
	"github.com/eduardo1520/goexpert/9-APIs/internal/infra/database"
	"github.com/go-chi/jwtauth"
	"net/http"
	"time"
)

type ErrorUser struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// GetJWT godoc
// @Summary      Gera um token JWT
// @Description  Endpoint para autenticar um usuário e retornar um token JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.GetJWTInput   true  "user credentials"
// @Success      200    {object} dto.GetJWTOutput
// @Failure      404	{object} ErrorUser
// @Failure      500    {object} ErrorUser
// @Router       /users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	fmt.Printf("%+v", jwt)
	jwtExperiesin := r.Context().Value("experiesin").(int64)
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExperiesin)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create  godoc
// @Summary      Cria um novo usuário
// @Description  Endpoint para criar um novo usuário na aplicação
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure      500 {object} ErrorUser
// @Router       /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
