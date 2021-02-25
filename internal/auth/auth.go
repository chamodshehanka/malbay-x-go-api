package auth

import (
	"errors"
	"github.com/Shehanka/malbay-x-go-api/config"
	"github.com/Shehanka/malbay-x-go-api/models"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func TokenValidation(r *http.Request) (bool, int, error) {
	jwtKey := []byte(config.GetEnv("secret.key"))
	c, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			return false, http.StatusUnauthorized, err
		}

		return false, http.StatusBadRequest, err
	}

	tknStr := c.Value
	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, http.StatusUnauthorized, err
		}

		return false, http.StatusBadRequest, err
	}

	if !tkn.Valid {
		err := errors.New("token is invalid")

		return false, http.StatusUnauthorized, err
	}

	return true, http.StatusOK, nil
}
