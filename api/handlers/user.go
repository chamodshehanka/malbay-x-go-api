package handlers

import (
	"context"
	"encoding/json"
	"github.com/Shehanka/malbay-x-go-api/config"
	"github.com/Shehanka/malbay-x-go-api/db"
	"github.com/Shehanka/malbay-x-go-api/internal/auth"
	"github.com/Shehanka/malbay-x-go-api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var userCollection = db.GetUserCollection()
var jwtKey = []byte(config.GetEnv("secret.key"))

func Signup(w http.ResponseWriter, r *http.Request) {
	var creds models.UserDetail

	_ = json.NewDecoder(r.Body).Decode(&creds)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	passwordString, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.MinCost)
	newCreds := models.UserDetail{
		Password: string(passwordString),
		Email:    creds.Email,
		Name:     creds.Name,
		Address:  creds.Address,
	}

	result, _ := userCollection.InsertOne(ctx, newCreds)

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

	err := bcrypt.CompareHashAndPassword([]byte(userCreds.Password), []byte(creds.Password))

	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())

		return
	}

	expirationTime := time.Now().Add(5 * time.Hour)
	claims := &models.Claims{
		Email:    creds.Email,
		Password: userCreds.Password,
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

	ResponseWithJSON(w, http.StatusOK, tokenString)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	v, httpStatus, err := auth.TokenValidation(r)

	if !v {
		RespondWithError(w, httpStatus, err.Error())

		return
	}

	ResponseWithJSON(w, http.StatusOK, "Welcome to malbay-x")
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	claims := &models.Claims{}
	v, httpStatus, err := auth.TokenValidation(r)

	if !v {
		RespondWithError(w, httpStatus, err.Error())

		return
	}

	//TODO: Increase the time more than 30s
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		RespondWithError(w, http.StatusBadRequest, "Token is not expired yet")

		return
	}

	expirationTime := time.Now().Add(5 * time.Hour)
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

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var creds, userCreds models.Credentials

	_ = json.NewDecoder(r.Body).Decode(&creds)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	f := bson.D{{"email", creds.Email}}

	if err := userCollection.FindOne(ctx, f).Decode(&userCreds); err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())

		return
	}

	ResponseWithJSON(w, http.StatusOK, "http://localhost:4000/api/v1/user/resetpassword/"+userCreds.Password)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var userCreds models.Credentials

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	f := bson.D{{"password", id}}

	if err := userCollection.FindOne(ctx, f).Decode(&userCreds); err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())

		return
	}

	ResponseWithJSON(w, http.StatusOK, "Password : "+userCreds.Password)
}
