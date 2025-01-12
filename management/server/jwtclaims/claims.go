package jwtclaims

import (
	"github.com/golang-jwt/jwt"
)

// AuthorizationClaims stores authorization information from JWTs
type AuthorizationClaims struct {
	UserId         string
	AccountId      string
	Domain         string
	DomainCategory string

	Raw jwt.MapClaims
}
