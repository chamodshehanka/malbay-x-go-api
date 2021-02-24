package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shehanka/malbay-x-go-api/db"
	"github.com/Shehanka/malbay-x-go-api/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

var userCollection = db.GetUserCollection()
var jwtKey = []byte("my_secret_key")

func Signup(w http.ResponseWriter, r *http.Request) {
	var creds models.UserDetail

	_ = json.NewDecoder(r.Body).Decode(&creds)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, _ := userCollection.InsertOne(ctx, creds)

	ResponseWithJSON(w, http.StatusOK, result)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds, userCreds models.Credentials

	_ = json.NewDecoder(r.Body).Decode(&creds)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	f := bson.D{{"email", creds.Email}}

	if err := userCollection.FindOne(ctx, f).Decode(&userCreds); err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())

		return
	}

	fmt.Println("Password :: ", userCreds.Password)

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		Username: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			RespondWithError(w, http.StatusUnauthorized, err.Error())

			return
		}

		RespondWithError(w, http.StatusBadRequest, err.Error())

		return
	}

	tknStr := c.Value
	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			RespondWithError(w, http.StatusUnauthorized, err.Error())

			return
		}

		RespondWithError(w, http.StatusBadRequest, err.Error())

		return
	}

	if !tkn.Valid {
		RespondWithError(w, http.StatusUnauthorized, "Token is invalid")

		return
	}

	ResponseWithJSON(w, http.StatusOK, "Welcome to malbay-x")
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			RespondWithError(w, http.StatusUnauthorized, err.Error())

			return
		}

		RespondWithError(w, http.StatusBadRequest, err.Error())

		return
	}

	tknStr := c.Value
	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !tkn.Valid {
		RespondWithError(w, http.StatusUnauthorized, "Token is invalid")

		return
	}

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			RespondWithError(w, http.StatusUnauthorized, err.Error())

			return
		}

		RespondWithError(w, http.StatusBadRequest, err.Error())

		return
	}
	//TODO: Increase the time more than 30s
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		RespondWithError(w, http.StatusBadRequest, "Token is not expired yet")

		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
