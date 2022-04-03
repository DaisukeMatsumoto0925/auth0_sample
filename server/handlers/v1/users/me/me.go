package me

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DaisukeMatsumoto0925/auth0_sample/middlewares/auth0"
	"github.com/form3tech-oss/jwt-go"
)

type User struct {
	Name string
	Age  int
}

var (
	subToUsers = map[string]User{
		"auth0|61a8178b21127500715968e2": {
			Name: "kourin",
			Age:  15,
		},
	}
)

func getUser(sub string) *User {
	user, ok := subToUsers[sub]
	if !ok {
		return nil
	}

	return &user
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	token := auth0.GetJWT(r.Context())
	fmt.Printf("jwt %+v\n", token)

	claims := token.Claims.(jwt.MapClaims)
	sub := claims["sub"].(string)

	user := getUser(sub)
	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
	}

	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
