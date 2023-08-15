package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/osvaldoabel/user-api/internal/adapters/http/handlers"
	"github.com/osvaldoabel/user-api/internal/dto"
	"github.com/osvaldoabel/user-api/internal/entity"
	"github.com/osvaldoabel/user-api/internal/services/user"
)

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

// Create user godoc
//
//	@Summary		Create user
//	@Description	Create user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.CreateUserInput	true	"user request"
//	@Success		201
//	@Failure		500	{object}	handlers.AppError
//	@Router			/users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var userPayload dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUserEntity(userPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := handlers.AppError{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	newUser, err := h.UserService.CreateUser(*user, ctx)
	if err != nil {
		handlers.JsonResponse(w, http.StatusInternalServerError, err)
		return
	}

	handlers.JsonResponse(w, http.StatusCreated, newUser)
}

// ListAccounts godoc
//
//	@Summary		List users
//	@Description	get all users
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			page	query		string	false	"page number"
//	@Param			limit	query		string	false	"limit"
//	@Success		200		{array}		dto.ListUsers
//	@Failure		404		{object}	handlers.AppError
//	@Failure		500		{object}	handlers.AppError
//	@Router			/users [get]
//	@Security		ApiKeyAuth
func (uh *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pagination := handlers.GetPaginationInfo(r)
	result, err := uh.UserService.ListUsers(pagination, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	handlers.JsonResponse(w, http.StatusOK, result)
}

// GetUser Godoc
//
//	@Summary		Get a user
//	@Description	Get a user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"user ID"	Format(uuid)
//	@Success		200	{object}	entity.User
//	@Failure		404
//	@Failure		500	{object}	handlers.AppError
//	@Router			/users/{id} [get]
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

	handlers.JsonResponse(w, http.StatusOK, user)
}

// UpdateUser Godoc
//
//	@Summary		Update a user
//	@Description	Update a user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string				true	"user ID"	Format(uuid)
//	@Param			request	body	dto.UpdateUserInput	true	"user request"
//	@Success		200
//	@Failure		404
//	@Failure		500	{object}	handlers.AppError
//	@Router			/users/{id} [put]
//	@Security		ApiKeyAuth
func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userInput dto.UpdateUserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := uh.UserService.FindUser(entity.ID(id), r.Context())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = user.Update(userInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err = uh.UserService.UpdateUser(user, r.Context())
	if err != nil {
		handlers.JsonResponse(w, http.StatusInternalServerError, user)
		return
	}

	handlers.JsonResponse(w, http.StatusOK, user)
}

// DeleteUser Godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"user ID"	Format(uuid)
//	@Success		200
//	@Failure		404  {object}	handlers.AppError
//	@Failure		500	{object}	handlers.AppError
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
		handlers.JsonResponse(w, http.StatusNotFound, nil)
		return
	}

	err = uh.UserService.DeleteUser(entity.ID(id), ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
