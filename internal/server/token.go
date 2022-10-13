package server

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (sv *Server) GenerateToken(id string) (string, error) {
	accessTokenClaims := jwt.MapClaims{}

	accessTokenClaims["id"] = id
	accessTokenClaims["iat"] = time.Now().Unix()
	accessTokenClaims["exp"] = time.Now().Add(sv.cfg.GetDuration("token.ttl") * time.Minute).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	signedToken, err := accessToken.SignedString([]byte(sv.cfg.GetString("token.accessSecret")))
	if err != nil {
		return "", fmt.Errorf("Sign with secret code failed: %s", err)
	}
	return signedToken, nil
}
