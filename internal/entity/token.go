package entity

import "github.com/go-chi/jwtauth"

type Token struct {
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}
