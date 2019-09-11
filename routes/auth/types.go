package auth

import "github.com/dgrijalva/jwt-go"

// NamespacedClaims are the claims handled within the hasura.io namespace.
type NamespacedClaims struct {
	AllowedRoles []string `json:"x-hasura-allowed-roles"`
	DefaultRole  string   `json:"x-hasura-default-role"`
	
	UserID       string   `json:"x-hasura-user-id"`
}

// JwtTokenClaims is the token claims to be added to the token.
type JwtTokenClaims struct {
	jwt.StandardClaims
	Name string `json:"name"`

	Response NamespacedClaims `json:"https://hasura.io/jwt/claims"`
}
