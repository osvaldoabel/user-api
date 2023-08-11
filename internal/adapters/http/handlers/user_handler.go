package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/osvaldoabel/user-api/internal/dto"
	"github.com/osvaldoabel/user-api/internal/entity"
	"github.com/osvaldoabel/user-api/internal/services/user"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserService  user.UserService
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(userService user.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// GetJWT godoc
//
//	@Summary		Get a user JWT
//	@Description	Get a user JWT
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.GetJWTInput	true	"user credentials"
//	@Success		200		{object}	dto.GetJWTOutput
//	@Failure		404		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/users/generate_token [post]
func (uh *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExperesIn").(int)
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := uh.UserService.FindUserByEmail(entity.Email(user.Email), ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
//
//	@Summary		Create user
//	@Description	Create user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.CreateUserInput	true	"user request"
//	@Success		201
//	@Failure		500	{object}	Error
//	@Router			/users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var userPayload dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(userPayload.Name, userPayload.Email, userPayload.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	newUser, err := h.UserService.CreateUser(*user, ctx)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, err)
		return
	}

	JsonResponse(w, http.StatusCreated, newUser)
}

// ListAccounts godoc
//
//	@Summary		List users
//	@Description	get all users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			page	query		string	false	"page number"
//	@Param			limit	query		string	false	"limit"
//	@Success		200		{array}		dto.ListUsers
//	@Failure		404		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/users [get]
//	@Security		ApiKeyAuth
func (uh *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pagination := GetPaginationInfo(r)
	users, err := uh.UserService.ListUsers(pagination, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	JsonResponse(w, http.StatusOK, users)
}

// GetUser Godoc
//
//	@Summary		Get a user
//	@Description	Get a user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"user ID"	Format(uuid)
//	@Success		200	{object}	entity.User
//	@Failure		404
//	@Failure		500	{object}	Error
//	@Router			/users/{id} [get]
//	@Security		ApiKeyAuth
func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := uh.UserService.FindUser(entity.ID(id), ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	JsonResponse(w, http.StatusOK, user)
}

// UpdateUser Godoc
//
//	@Summary		Update a user
//	@Description	Update a user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string				true	"user ID"	Format(uuid)
//	@Param			request	body	dto.UpdateUserInput	true	"user request"
//	@Success		200
//	@Failure		404
//	@Failure		500	{object}	Error
//	@Router			/users/{id} [put]
//	@Security		ApiKeyAuth
func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userInput dto.UpdateUserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)

	user, err := uh.UserService.UpdateUser(userInput, ctx)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, user)
		return
	}

	JsonResponse(w, http.StatusOK, user)
}

// DeleteUser Godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"user ID"	Format(uuid)
//	@Success		200
//	@Failure		404
//	@Failure		500	{object}	Error
//	@Router			/users/{id} [delete]
//	@Security		ApiKeyAuth
func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := uh.UserService.FindUser(entity.ID(id), ctx)
	if err != nil {
		JsonResponse(w, http.StatusNotFound, nil)
		return
	}

	err = uh.UserService.DeleteUser(entity.ID(id), ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
