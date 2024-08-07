package token

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	signingKey = "mrbek"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateJWTToken(userID string, email string, role string) *Tokens {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["email"] = email
	claims["role"] = role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(180 * time.Minute).Unix() // Token expires in 3 minutes
	access, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		log.Fatal("error while generating access token : ", err)
	}

	rftClaims := refreshToken.Claims.(jwt.MapClaims)
	rftClaims["user_id"] = userID
	rftClaims["email"] = email
	rftClaims["role"] = role
	rftClaims["iat"] = time.Now().Unix()
	rftClaims["exp"] = time.Now().Add(24 * time.Hour).Unix() // Refresh token expires in 24 hours
	refresh, err := refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		log.Fatal("error while generating refresh token : ", err)
	}

	return &Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, errors.New("parsing token:" + err.Error())
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
