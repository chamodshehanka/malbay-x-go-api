package auth

import (
	"context"
	"errors"
	"github.com/Shehanka/malbay-x-go-api/config"
	"github.com/Shehanka/malbay-x-go-api/db"
	"github.com/Shehanka/malbay-x-go-api/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func TokenValidation(r *http.Request) (bool, int, error) {
	userCollection := db.GetUserCollection()
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

	var userCreds models.Credentials

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	f := bson.D{{"email", claims.Email}}

	if err := userCollection.FindOne(ctx, f).Decode(&userCreds); err != nil {
		return false, http.StatusUnauthorized, err
	}

	if userCreds.Password != claims.Password {
		return false, http.StatusUnauthorized, err
	}

	return true, http.StatusOK, nil
}
