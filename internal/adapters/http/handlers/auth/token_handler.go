package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/osvaldoabel/user-api/internal/adapters/http/handlers"
	"github.com/osvaldoabel/user-api/internal/dto"
	"github.com/osvaldoabel/user-api/internal/entity"
	"github.com/osvaldoabel/user-api/internal/services/user"
)

type TokenHandler struct {
	UserService user.UserService
}

func NewTokenHandler(userService user.UserService) *TokenHandler {
	return &TokenHandler{
		UserService: userService,
	}
}

// GetJWT godoc
//
//	@Summary		Get a user JWT
//	@Description	Get a user JWT
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.GetJWTInput	true	"user credentials"
//	@Success		200		{object}	dto.GetJWTOutput
//	@Failure		404		{object}	handlers.AppError
//	@Failure		500		{object}	handlers.AppError
//	@Router			/users/generate_token [post]
func (th *TokenHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExperesIn").(int)
	var input dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		handlers.JsonResponse(w, http.StatusBadRequest, err)
		return
	}

	u, err := th.UserService.FindUserByEmail(entity.Email(input.Email), r.Context())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := handlers.AppError{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if !u.ValidatePassword(input.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	handlers.JsonResponse(w, http.StatusOK, accessToken)
}
