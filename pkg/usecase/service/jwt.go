package service

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gonzamedrano09/chat/pkg/config"
	"time"
)

type JWTPayload struct {
	Username   string    `json:"username"`
	Expiration time.Time `json:"exp"`
}

type JWTServiceInterface interface {
	GenerateJWT(username string) (string, error)
	CheckJWT(tokenString string) (string, bool)
}

type JWTService struct {
}

func NewJWTService() JWTServiceInterface {
	return &JWTService{}
}

func (js *JWTService) GenerateJWT(username string) (string, error) {
	now := time.Now()
	tokenPayload := JWTPayload{
		Username:   username,
		Expiration: now.Add(time.Hour * 24 * 7),
	}
	tokenPayloadJson, err := json.Marshal(tokenPayload)
	if err != nil {
		return "", err
	}
	var mapClaimsJson map[string]any
	if err := json.Unmarshal(tokenPayloadJson, &mapClaimsJson); err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(mapClaimsJson))
	tokenString, err := token.SignedString([]byte(config.Config.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (js *JWTService) CheckJWT(tokenString string) (string, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.SecretKey), nil
	})
	if err != nil {
		return "", false
	}
	if mapClaims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mapClaimsJson, err := json.Marshal(mapClaims)
		if err != nil {
			return "", false
		}
		var tokenPayload JWTPayload
		if err := json.Unmarshal(mapClaimsJson, &tokenPayload); err != nil {
			return "", false
		}
		if ok := tokenPayload.Expiration.After(time.Now()); !ok {
			return "", false
		}
		return tokenPayload.Username, true
	}
	return "", false
}
